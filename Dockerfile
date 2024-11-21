FROM golang:1.21-alpine3.18 AS builder


# Set environment variables for Go
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Install git (required for some Go modules)
RUN apk add --no-cache git

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

# =========================
# Run Stage
# =========================
FROM alpine:3.18

# Install necessary packages
RUN apk add --no-cache \
    bash \
    netcat-openbsd \
    curl \
    tar \
    gzip

# Set the working directory
WORKDIR /app

# Install migrate CLI
ENV MIGRATE_VERSION=v4.15.2

RUN curl -fsSL "https://github.com/golang-migrate/migrate/releases/download/${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz" \
    | tar xzv \
    && mv migrate /usr/local/bin/ \
    && chmod +x /usr/local/bin/migrate \
    && rm -rf migrate.linux-amd64.tar.gz

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Copy necessary scripts and files
COPY start.sh wait-for.sh ./
COPY db/migration ./db/migration

# Ensure scripts are executable
RUN chmod +x start.sh wait-for.sh

# Expose the application ports
EXPOSE 8081 9090

# Set ENTRYPOINT to the start script
ENTRYPOINT ["/app/start.sh"]

# Default command to run the application
CMD ["./main"]