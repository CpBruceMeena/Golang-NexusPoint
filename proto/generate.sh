#!/bin/bash

# Install required tools if not already installed
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate code for golang-central
protoc --go_out=../golang-central --go_opt=paths=source_relative \
    --go-grpc_out=../golang-central --go-grpc_opt=paths=source_relative \
    user.proto

# Generate code for golang-grpc-app
protoc --go_out=../golang-grpc-app --go_opt=paths=source_relative \
    --go-grpc_out=../golang-grpc-app --go-grpc_opt=paths=source_relative \
    user.proto 