# Variables
BINARY_NAME := colorfinder 
SRC_DIR := ./src
GOBIN ?= $(shell go env GOBIN)
GOPATH := $(shell go env GOPATH)
BIN_DIR := $(if $(GOBIN),$(GOBIN),$(GOPATH)/bin)

# Default target
all: build

# Build the project
build:
	@echo "Building the project..."
	go build -o bin/$(BINARY_NAME) $(SRC_DIR)

# Install the project binary
install:
	@echo "Installing the binary..."
	go install -o $(BIN_DIR)/$(BINARY_NAME) $(SRC_DIR)

# Run linting
lint:
	@echo "Running linter..."
	golangci-lint run

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Clean up generated files
clean:
	@echo "Cleaning up..."
	rm -rf bin/*

# Run all steps: fmt, lint, test, and build
all: fmt lint test build

