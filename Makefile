# Makefile for vikunja-mcp

.PHONY: lint build

fmt:
	go fmt ./...

goimports:
	goimports -w .

# Lint the Go code using golangci-lint
lint:
	golangci-lint run -v -cover -c ./.golangci.yaml ./...

test:
	go test -v -cover ./...

# Build the server binary
build:
	go build -o bin/server ./cmd/server

ci-local: fmt goimports lint test