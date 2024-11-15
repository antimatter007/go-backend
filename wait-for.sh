#!/bin/sh

# wait-for.sh
# Waits for the specified host and port to become available

set -e

host_port="$1"
shift
cmd="$@"

host=$(echo "$host_port" | cut -d':' -f1)
port=$(echo "$host_port" | cut -d':' -f2)

until nc -z "$host" "$port"; do
  echo "Waiting for $host:$port to be available..."
  sleep 1
done

echo "$host:$port is available, executing command."
exec $cmd
