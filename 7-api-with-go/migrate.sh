#!/bin/bash

# Ensure a name is provided
if [ -z "$1" ]; then
  echo "Usage: ./migrate.sh <migration_name>"
  exit 1
fi

# Run golang-migrate create command
migrate create -ext sql -dir internals/db/migrations -seq "$1"