# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod before source for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build
RUN go build -o property-service ./services/property

# Run stage
FROM alpine

WORKDIR /app

COPY --from=builder /app/property-service .

EXPOSE 3002

CMD ["./property-service"]
