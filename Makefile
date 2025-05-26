# Makefile for vikunja-mcp

.PHONY: lint build

fmt:
	go fmt ./...

# Lint the Go code using golangci-lint
lint:
	golangci-lint run -v -cover -c ./.golangci.yaml ./...

# Build the server binary
build:
	go build -o bin/server ./cmd/server