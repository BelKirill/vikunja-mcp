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
build-binary:
	go build -o bin/server ./cmd/mcp/main.go

docker:
	docker buildx build . -f Dockerfile -t vikunja-mcp:latest

build: build-binary docker

ci-local: fmt goimports lint

ci-test: fmt goimports lint test