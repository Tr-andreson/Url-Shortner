# ==========================================
# STAGE 1: Build the Go application binary
# ==========================================
FROM golang:1.27-alpine AS builder

# Install build dependencies (like git or certificates if needed)
RUN apk add --no-cache git ca-certificates && update-ca-certificates

# Set the working directory inside the container
WORKDIR /app

# Copy dependency manifests first to leverage Docker caching
COPY go.mod go.sum ./
RUN go-mod-download.sh || go mod download

# Copy the rest of the application source code
COPY . .

# Compile the binary with optimization flags for containerization:
# - CGO_ENABLED=0 creates a statically linked binary (no dependencies on C libraries)
# - GOOS=linux targets the Linux operating system
# - ldflags "-s -w" strips debug/logging symbols to shrink the file size
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o my-go-app ./cmd/server/main.go

# ==========================================
# STAGE 2: Final minimal runtime environment
# ==========================================
FROM gcr.io/distroless/static-debian12:latest

# Set up a secure, non-root user (built into distroless, UID 65532)
USER nonroot:nonroot

# Copy the compiled binary from the builder stage
COPY --from=builder /app/my-go-app /my-go-app

# Copy system certificates to allow your Go app to make HTTPS external API requests
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Expose the port your Go application listens on
EXPOSE 8080

# Run the app
ENTRYPOINT ["/my-go-app"]
