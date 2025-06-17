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
	go build -o bin/server ./cmd/mcp/main.go
	docker build -t vikunja-mcp:latest .

ci-local: fmt goimports lint

ci-test: fmt goimports lint test