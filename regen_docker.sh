#!/usr/bin/env bash

code_root=github.com/EricTusk/$(basename $PWD)

docker run -it --rm \
	-v $PWD:/go/src/${code_root} \
	-e PROTOC_INSTALL=/go \
	-w /go/src/${code_root} \
	proto-tools:3.6 ./regen.sh
