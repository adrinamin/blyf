# Stage 1: Build the Go application
FROM golang:1-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

COPY main.go ./

# Copy go.mod and go.sum files
# COPY go.mod go.sum ./

# Download dependencies
# RUN go mod download

# Copy the source code
# COPY . .

# Build the Go application
RUN go build main.go

# Stage 2: Create a minimal image with the compiled binary
FROM alpine:latest

# Install CA certificates for HTTPS support
# RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the application's port (adjust if necessary)
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
