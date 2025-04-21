#!/bin/bash

# Install buf if not already installed
if ! command -v buf &> /dev/null; then
    echo "Installing buf..."
    go install github.com/bufbuild/buf/cmd/buf@latest
fi

# Clean existing generated files
echo "Cleaning existing generated files..."
rm -rf gen/go/*

# Generate protobuf files
echo "Generating protobuf files..."
buf generate

echo "Protobuf files generated successfully!" 