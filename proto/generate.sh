#!/bin/bash

# Install required tools
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Create necessary directories
mkdir -p ../golang-central/proto
mkdir -p ../golang-grpc-app/proto

# Generate code
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/user.proto

# Copy generated files to both services
cp -r proto/*.pb.go ../golang-central/proto/
cp -r proto/*.pb.go ../golang-grpc-app/proto/

echo "Proto files generated and copied successfully!" 