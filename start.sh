#!/bin/sh

echo "Starting the app..."

# Extract host and port from DB_SOURCE
# Assumes DB_SOURCE is in the format postgresql://user:password@host:port/dbname
DB_HOST=$(echo "$DB_SOURCE" | awk -F[@:] '{print $(NF-1)}')
DB_PORT=$(echo "$DB_SOURCE" | awk -F[@:] '{print $NF}' | awk -F/ '{print $1}')

echo "Waiting for database at $DB_HOST:$DB_PORT..."

# Wait for the database to be ready
./wait-for.sh "$DB_HOST:$DB_PORT" -- echo "Database is up"

# Run database migrations
migrate -path db/migration -database "$DB_SOURCE" -verbose up

# Start the Go application
./main
