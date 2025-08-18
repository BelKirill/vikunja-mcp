# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Build and Test
- `make build` - Build MCP server binary and Docker image
- `make build-binary` - Build only the MCP server binary to `./bin/server`
- `make test` - Run all Go tests with coverage
- `make ci-test` - Full CI pipeline: format, imports, lint, and test
- `make ci-local` - Local development checks: format, imports, and lint

### Code Quality
- `make fmt` - Format Go code using `go fmt`
- `make goimports` - Fix imports using `goimports`
- `make lint` - Run golangci-lint with project configuration
- `go test ./... -v` - Run tests with verbose output
- `go test -v -cover ./...` - Run tests with coverage reports

### Docker
- `make docker` - Build Docker image using buildx
- `docker-compose.yaml` - Available for local Vikunja integration

### Binary Execution
- `./bin/server` - Run the built MCP server
- Main entry points: `cmd/mcp/main.go`, `cmd/rest-server/main.go`, `cmd/feel/main.go`

## Architecture Overview

This is an AI-powered task management system that integrates Vikunja with Claude via MCP (Model Context Protocol). The system uses multi-tier LLM architecture for intelligent task recommendations optimized for ADHD workflows.

### Core Components

**MCP Server** (`pkg/mcp/`)
- Implements Model Context Protocol for Claude integration
- Registers tools: `daily-focus`, `get-focus-recommendation`, `get-task-metadata`, `upsert_task`
- Entry point: `cmd/mcp/main.go`

**Focus Engine** (`internal/focus/`)
- Business logic for AI-powered task recommendations
- MCP handlers in `mcp_handlers.go`, REST handlers in `rest_handlers.go`
- Service layer coordinates Vikunja API and OpenAI decision engine

**Vikunja Integration** (`pkg/vikunja/client/`, `internal/vikunja/`)
- HTTP client for Vikunja API operations
- Task CRUD operations with metadata embedding
- Authentication and retry logic

**AI Decision Engine** (`internal/openai/`)
- OpenAI GPT integration for contextual task analysis
- Multi-model approach for task ranking and selection
- Explainable recommendations with reasoning

**Data Models** (`models/`)
- `FocusOptions`: Energy, mode, time constraints for task selection
- `HyperFocusMetadata`: ADHD-optimized task attributes (energy, mode, duration, hyperfocus compatibility)
- `FocusResult`: Enriched tasks with AI-calculated focus scores

### Key Design Patterns

**Metadata Embedding**: Rich ADHD-specific metadata stored as JSON in Vikunja task descriptions, extracted and enriched by the focus engine.

**Energy-Aware Matching**: Tasks matched to user's current cognitive state (low/medium/high/social energy, deep/quick/admin modes).

**Hyperfocus Optimization**: 25-minute base units with extension capabilities, compatibility scoring 1-5.

**Multi-Tier LLM**: Claude for natural language interface, OpenAI for contextual task analysis and ranking.

## Configuration

### Environment Variables
- `VIKUNJA_URL`: Your Vikunja instance URL
- `VIKUNJA_TOKEN`: API token for Vikunja authentication
- `OPENAI_API_KEY`: OpenAI API key for AI decisions
- `OPENAI_MODEL`: Model to use (default: gpt-4o-mini)
- `LOG_LEVEL`: Logging level (info, debug, etc.)
- `SERVER_PORT`: REST server port (default: 8080)

### Dependencies
- Go 1.23.5+ required
- golangci-lint for code quality
- swag for API documentation generation
- Active Vikunja instance
- OpenAI API access

## Testing Strategy

Test files follow `*_test.go` pattern:
- `pkg/vikunja/client/*_test.go` - HTTP client and API integration tests
- `internal/focus/*_test.go` - Business logic and handler tests
- `pkg/mcp/server_test.go` - MCP protocol implementation tests

Use `make test` to run full suite with coverage reporting.