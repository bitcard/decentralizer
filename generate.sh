#!/bin/bash
go-bindata -pkg app -o app/static.go static/


CPP="././sdk/libdn/dependencies/include/pb/"

echo "Compiling protocol buffers for API";
protoc -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:./ pb/*.proto
#Generate reverse proxy: https://github.com/grpc-ecosystem/grpc-gateway
protoc -I/usr/local/include -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. pb/platform.proto

echo "Compiling protocol buffers for windows SDK";
protoc -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --error_format=msvs --cpp_out=. pb/*.proto
protoc -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --error_format=msvs --grpc_out=. --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` pb/*.proto
rm -rf "${CPP}/*"
mv -f pb/*.cc "${CPP}"
mv -f pb/*.h "${CPP}"