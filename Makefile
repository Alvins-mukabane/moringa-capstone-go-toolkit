.PHONY: all build run-cli run-api test tidy

# Default target
all: tidy test build

# Tidy dependencies
tidy:
	go mod tidy

# Build the binaries
build:
	@mkdir -p bin
	go build -o bin/joke-cli ./cmd/joke-cli
	go build -o bin/joke-api ./cmd/joke-api

# Run the CLI app
run-cli:
	go run ./cmd/joke-cli

# Run the API server
run-api:
	go run ./cmd/joke-api

# Run unit tests
test:
	go test -v ./...
