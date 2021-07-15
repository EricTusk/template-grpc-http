ARG baseBuilderImage=golang:1.15.11-gomod-cuda10-gcc49
ARG baseImage=debian:stretch
ARG arch=amd64
ARG device=none

FROM ${baseBuilderImage} as builder


WORKDIR /go/src/github.com/EricTusk/template-http-grpc

COPY . ./

ARG arch=${arch}
ARG device=${device}
RUN make clean && make build arch=${arch} device=${device}


FROM ${baseImage}


WORKDIR /

COPY --from=builder /go/src/github.com/EricTusk/template-http-grpc/example /config
COPY --from=builder /go/src/github.com/EricTusk/template-http-grpc/template-http-grpc /template-http-grpc

CMD ["/bin/sh"]
