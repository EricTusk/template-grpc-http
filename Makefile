arch ?=amd64
device ?=none

include .makerc/amd64_none

IMAGE_FLAGS :=
IMAGE_FLAGS := $(IMAGE_FLAGS) --build-arg baseBuilderImage=$(baseBuilderImage)
IMAGE_FLAGS := $(IMAGE_FLAGS) --build-arg baseImage=$(baseImage)
IMAGE_FLAGS := $(IMAGE_FLAGS) --build-arg arch=$(arch)
IMAGE_FLAGS := $(IMAGE_FLAGS) --build-arg device=$(device)

COMMIT := `git rev-parse --short HEAD`
BUILD := $(COMMIT)

ifneq ($(arch), )
BUILD := $(BUILD)-$(arch)
endif

ifneq ($(device), none)
BUILD := $(BUILD)-$(device)
endif

GO=go

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

VERSION := v0.1.0
TARGETS := template-http-grpc
TEST_TARGETS := 
project=github.com/EricTusk/template-http-grpc

LDFLAGS += -X "$(project)/version.BuildTS=$(shell date -u '+%Y-%m-%d %I:%M:%S')"
LDFLAGS += -X "$(project)/version.GitHash=$(shell git rev-parse HEAD)"
LDFLAGS += -X "$(project)/version.Version=$(VERSION)"
LDFLAGS += -X "$(project)/version.GitBranch=$(shell git rev-parse --abbrev-ref HEAD)"

.PHONY: all
all: clean lint check build

.PHONY: build
build: $(TARGETS) $(TEST_TARGETS)

.PHONY: tar
tar: build
	@tar cvf build.tar $(TARGETS)

$(TARGETS): $(SRC)
	$(GO) build -mod=vendor -ldflags '$(LDFLAGS)' $(project)/cmd/$@

.PHONY: image
image: $(TARGETS)
	docker build $(IMAGE_FLAGS) -f ./Dockerfile -t  template-http-grpc:$(VERSION)-$(BUILD) .

.PHONY: push
push: image
	docker push template-http-grpc:$(VERSION)-$(BUILD)

.PHONY: tag
tag:
	@git tag $(VERSION)-$(COMMIT)
	@git push origin --tags

.PHONY: lint
lint:
	@golangci-lint run --deadline=10m

packages = $(shell go list ./...|grep -v /vendor/)
.PHONY: test
test: check
	$(GO) test -v -race ${packages}

.PHONY: cov
cov: check
	gocov test -v -race $(packages) | gocov-html > coverage.html
    @cat coverage.html |grep "<code>Report Total</code>" | perl -nle 'print  "Total Coverage: $$1" if /([0-9.]+%)/g'

.PHONY: check
check:
	@echo skip go vet

.PHONY: clean
clean:
	rm -f $(TARGETS) $(TEST_TARGETS)
