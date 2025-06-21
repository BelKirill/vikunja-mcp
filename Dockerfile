# Stage 1: Build the Go binary
FROM golang:1.24.4-alpine AS builder

# Install security updates and build dependencies
RUN apk update && apk add --no-cache git ca-certificates tzdata && \
    update-ca-certificates

# Create a non-root user for building
RUN adduser -D -g '' appuser

WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build with security flags and optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o vikunja-mcp ./cmd/mcp

# Stage 2: Create minimal runtime image
FROM scratch

# Copy timezone data and certificates from builder
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

# Copy the binary from builder
COPY --from=builder /app/vikunja-mcp /vikunja-mcp

# Use non-root user
USER appuser

# Expose port (document only, actual port binding done at runtime)
EXPOSE 8080

# Environment variables with sensible defaults
ENV SERVER_PORT="8080"
ENV LOG_LEVEL="info"
ENV READ_TIMEOUT="30s"
ENV WRITE_TIMEOUT="30s"
ENV OPENAI_MODEL="gpt-4o-mini"
ENV OPENAI_BASE_URL="https://api.openai.com/v1"

# Add health check metadata
LABEL org.opencontainers.image.title="vikunja-mcp"
LABEL org.opencontainers.image.description="ADHD-optimized task management MCP server"
LABEL org.opencontainers.image.version="0.1.0"

# Run the binary
ENTRYPOINT ["/vikunja-mcp"]