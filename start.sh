#!/bin/sh

echo "Starting the app..."

# Extract host and port from DB_SOURCE
DB_HOST=$(echo "$DB_SOURCE" | awk -F[@:] '{print $(NF-1)}')
DB_PORT=$(echo "$DB_SOURCE" | awk -F[@:] '{print $NF}' | awk -F/ '{print $1}')

echo "Waiting for database at $DB_HOST:$DB_PORT..."

# Verify that migrate is installed
if ! command -v migrate &> /dev/null
then
    echo "migrate could not be found"
    exit 1
fi

echo "migrate version: $(migrate -version)"

# Wait for the database to be ready
./wait-for.sh "$DB_HOST:$DB_PORT" -- echo "Database is up"

# Run database migrations
echo "Running database migrations..."
migrate -path db/migration -database "$DB_SOURCE" -verbose up

# Debugging: Print non-sensitive environment variables
echo "ENVIRONMENT: $ENVIRONMENT"
echo "HTTP_SERVER_ADDRESS: $HTTP_SERVER_ADDRESS"
echo "GRPC_SERVER_ADDRESS: $GRPC_SERVER_ADDRESS"

# Start the Go application
./main
