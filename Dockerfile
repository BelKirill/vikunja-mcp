# Stage 1: Build the Go binary
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o vikunja-mcp ./cmd/server

# Stage 2: Run the binary using a minimal image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/vikunja-mcp .
EXPOSE 8080
CMD ["./vikunja-mcp"]