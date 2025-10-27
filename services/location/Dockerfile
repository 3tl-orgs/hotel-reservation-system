# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o app ./services/location

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS and timezone data
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Create non-root user
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

# Copy binary from builder
COPY --from=builder /app/app .

# Change ownership
RUN chown -R appuser:appuser /app

# Switch to non-root user
USER appuser

# GIN port
EXPOSE 3001

# Run the binary
CMD ["./app"]