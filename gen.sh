#!/usr/bin/env bash

protoc -I=./proto \
    --go_out . \
    ./proto/api/*.proto

protoc -I=./proto \
    --go_out . \
    ./proto/core/*.proto

protoc -I=./proto \
    --go_out . \
    ./proto/core/contract/*.proto

protoc -I=./proto \
    --go-grpc_out . \
    ./proto/api/*.proto

# move proto files to the right places
cp -r github.com/zhang0125/go-tron-proto/* .
rm -rf github.com