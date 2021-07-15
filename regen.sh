#!/bin/bash

if [ -z "$PROTOC_INSTALL" ]; then
	echo "PROTOC_INSTALL not set"
	exit 1
fi

basepath=$GOPATH/src
pb_package=github.com/EricTusk/template-http-grpc/pb
proto_install="$PROTOC_INSTALL"
go_package=api
rm -rf $go_package

cd $basepath
for i in $(ls $basepath/$pb_package/*.proto); do
	echo $i
	fn=$pb_package/$(basename "$i")
	$proto_install/bin/protoc -I$proto_install/include -I. \
		-I$GOPATH/src \
		-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis\
		-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway\
		--go_out=plugins=grpc:. "$fn"
	$proto_install/bin/protoc -I$proto_install/include -I. \
		-I$GOPATH/src \
		-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis\
		-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway\
		--grpc-gateway_out=logtostderr=true:. "$fn"
	$proto_install/bin/protoc -I$proto_install/include -I. \
		-I$GOPATH/src \
		-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis\
		-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway\
		--swagger_out=logtostderr=true:. "$fn"
done
