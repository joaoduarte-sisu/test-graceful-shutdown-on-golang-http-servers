#!/bin/sh
protoc --go_out=./service-b/protos --go-grpc_out=./service-b/protos protos/service_b.proto
protoc --go_out=./service-a/protos --go-grpc_out=./service-a/protos protos/service_b.proto
