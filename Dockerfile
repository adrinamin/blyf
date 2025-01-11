# Stage 1: Build the Go application
FROM golang:1-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and source files to container 
COPY go.mod .
COPY cmd ./cmd
COPY api ./api
# Download dependencies
RUN go mod download

# Build the Go application
RUN go build -o blyf ./cmd/server/ 

# Stage 2: Create a minimal image with the compiled binary
FROM alpine:latest

# Install CA certificates for HTTPS support
# RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/blyf .

# Expose the application's port (adjust if necessary)
EXPOSE 8080

# Command to run the executable
CMD ["./blyf"]
