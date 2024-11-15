#!/bin/sh

echo "Starting the app..."

# Extract host and port from DB_SOURCE
DB_HOST=$(echo "$DB_SOURCE" | awk -F[@:] '{print $(NF-1)}')
DB_PORT=$(echo "$DB_SOURCE" | awk -F[@:] '{print $NF}' | awk -F/ '{print $1}')

echo "Waiting for database at $DB_HOST:$DB_PORT..."

# Verify that migrate is installed
migrate -version

# Wait for the database to be ready
./wait-for.sh "$DB_HOST:$DB_PORT" -- echo "Database is up"

# Run database migrations
migrate -path db/migration -database "$DB_SOURCE" -verbose up

# Start the Go application
./main
