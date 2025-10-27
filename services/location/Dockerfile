# Stage 1: Build Go app
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o app .

# Stage 2: Run the app
FROM alpine
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 3001
CMD ["./app"]
