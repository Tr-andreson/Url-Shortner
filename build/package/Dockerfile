# Stage 1: Build the Go binary
FROM golang:1.27-alpine AS builder
WORKDIR /app

# Copy dependency files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build from the cmd/server directory
RUN go build -o /app/bin/server ./cmd/server/main.go

# Stage 2: Run the binary in a minimal image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/bin/server .

EXPOSE 8080
CMD ["./server"]
