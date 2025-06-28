# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

# Install git and build tools
RUN apk add --no-cache git

WORKDIR /app
COPY . .

# Download Go modules
RUN go mod download

# Build statically linked binary (important for Alpine!)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp .

# Stage 2: Minimal Alpine image
FROM alpine:latest

# Create working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/myapp .

# Make sure it's executable
RUN chmod +x ./myapp

# Optional: expose port
EXPOSE 33500

# Run the binary
CMD ["./myapp"]