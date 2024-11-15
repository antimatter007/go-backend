# Build stage
FROM golang:1.21-alpine3.18 AS builder

# Set environment variables for build
ENV GO111MODULE=on

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to leverage Docker layer caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN go build -o main main.go

# Run stage
FROM alpine:3.18

# Install necessary packages
RUN apk add --no-cache bash netcat-openbsd

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Copy necessary scripts and files
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration

# Make sure the scripts are executable
RUN chmod +x start.sh wait-for.sh

# Expose the application ports
EXPOSE 8081 9090

# Set ENTRYPOINT to the start script
ENTRYPOINT ["/app/start.sh"]

# Default command to run the application
CMD ["./main"]
