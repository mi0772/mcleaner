# Makefile for mcleaner

.PHONY: build test clean install uninstall release help

# Build variables
BINARY_NAME=mcleaner
VERSION?=$(shell git describe --tags --always --dirty)
BUILD_TIME=$(shell date +%FT%T%z)
LDFLAGS=-ldflags "-s -w -X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)"

# Default target
all: build

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	go build $(LDFLAGS) -o $(BINARY_NAME) .

# Build for multiple platforms
build-all:
	@echo "Building for multiple platforms..."
	@mkdir -p dist
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-arm64 .
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-arm64 .
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-windows-amd64.exe .

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -race -buildvcs -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Lint the code
lint:
	@echo "Running linter..."
	golangci-lint run

# Format the code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Vet the code
vet:
	@echo "Vetting code..."
	go vet ./...

# Run staticcheck
staticcheck:
	@echo "Running staticcheck..."
	staticcheck ./...

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod verify

# Tidy dependencies
tidy:
	@echo "Tidying dependencies..."
	go mod tidy

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -f $(BINARY_NAME)
	rm -rf dist/
	rm -f coverage.out coverage.html

# Install the application
install: build
	@echo "Installing $(BINARY_NAME)..."
	sudo cp $(BINARY_NAME) /usr/local/bin/

# Uninstall the application
uninstall:
	@echo "Uninstalling $(BINARY_NAME)..."
	sudo rm -f /usr/local/bin/$(BINARY_NAME)

# Create a new release (requires clean git state)
release:
	@echo "Creating release..."
	@if [ -n "$$(git status --porcelain)" ]; then \
		echo "Error: Git working directory is dirty. Please commit or stash changes."; \
		exit 1; \
	fi
	@read -p "Enter version (e.g., v1.0.0): " version; \
	git tag -a $$version -m "Release $$version"; \
	git push origin $$version

# Run all checks (fmt, vet, lint, test)
check: fmt vet lint test

# Run the application
run:
	go run . $(ARGS)

# Show help
help:
	@echo "Available targets:"
	@echo "  build         - Build the application"
	@echo "  build-all     - Build for multiple platforms"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage"
	@echo "  lint          - Run linter"
	@echo "  fmt           - Format code"
	@echo "  vet           - Vet code"
	@echo "  staticcheck   - Run staticcheck"
	@echo "  deps          - Download dependencies"
	@echo "  tidy          - Tidy dependencies"
	@echo "  clean         - Clean build artifacts"
	@echo "  install       - Install the application"
	@echo "  uninstall     - Uninstall the application"
	@echo "  release       - Create a new release"
	@echo "  check         - Run all checks"
	@echo "  run           - Run the application (use ARGS='...' for arguments)"
	@echo "  help          - Show this help"