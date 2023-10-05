#!/bin/sh

set -e

DB_HOST="peakpal_db_container"
DB_USER="root"
DB_NAME="peakpal_carpool_db"
DB_PASSWORD="secret"  # Replace with your desired password

echo "Check if the database exists"
if !  PGPASSWORD=secret psql -h "$DB_HOST" -U "$DB_USER" -w -lqt | cut -d \| -f 1 | grep -qw "$DB_NAME"; then
    echo "Creating database1: $DB_NAME"
    PGPASSWORD=secret psql -h $DB_HOST -U $DB_USER -d postgres -w -c "CREATE DATABASE peakpal_carpool_db;"
fi

echo "start the app"
exec "$@"
