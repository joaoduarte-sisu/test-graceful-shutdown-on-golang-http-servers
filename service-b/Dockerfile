FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code and proto files
COPY . .

# Build the Go binary
RUN go build -o service-b .

# Final stage: small image with only the binary and necessary libs
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/service-b .

CMD ["./service-b"]
