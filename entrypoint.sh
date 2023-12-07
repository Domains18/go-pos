#!/bin/sh

# Exit immediately if a command exits with a non-zero status.
set -e

#check environment variables
if [ -z "$DB_CONNECTION" ] || [ -z "$DB_USERNAME" ] || [ -z "$DB_PASSWORD" ] || [ -z "$DB_HOST" ] || [ -z "$DB_PORT" ] || [ -z "$DB_DATABASE" ]; then
  echo "Error: One or more required environment variables are not set."
  exit 1
fi

#construct the database connection string
DB_URL="$DB_CONNECTION://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_DATABASE?sslmode=disable"

#wait for the database to be ready
until psql "$DB_URL" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

# sleep 5

#run the migrations
echo "Running migrations"
migrate -path=/app/migrations -database "$DB_URL" -verbose up

#run the application
echo "Running application"
exec ./gopos