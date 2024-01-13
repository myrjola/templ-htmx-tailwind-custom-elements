#!/bin/sh

# Run the bundler
go generate -run bundler ./...

# Run the server
go run ./cmd/web