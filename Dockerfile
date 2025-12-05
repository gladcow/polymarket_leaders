# Build stage
FROM golang:1.24-alpine AS builder

# Install build dependencies
# Update package index first to avoid hangs
RUN apk update && apk add --no-cache git make

# Set working directory
WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary
# CGO_ENABLED=0 creates a statically linked binary
# -ldflags="-w -s" strips debug info to reduce binary size
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o polymarket_leaders \
    ./main.go

# Run stage - minimal image
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk update && apk --no-cache add ca-certificates tzdata

# Create non-root user for security
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /build/polymarket_leaders .

# Copy config directory if it exists (optional, can be mounted as volume)
# COPY --from=builder /build/config ./config

# Change ownership to non-root user
RUN chown -R appuser:appuser /app

# Switch to non-root user
USER appuser

# Expose port (adjust if your dashboard uses a different port)
EXPOSE 8080

# Run the binary
CMD ["./polymarket_leaders"]

