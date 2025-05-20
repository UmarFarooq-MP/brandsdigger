# Stage 1: Build the Go binary
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go binary and name it explicitly
RUN CGO_ENABLED=0 GOOS=linux go build -o brandsdigger ./cmd

# Stage 2: Run the binary in a minimal image
FROM alpine:latest

# Install CA certificates (optional, but required for HTTPS)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/brandsdigger .

# Expose port if needed
EXPOSE 8080

# Run the binary
ENTRYPOINT ["./brandsdigger"]
