#!/bin/bash

env PGPASSWORD="$KHULNASOFT_POSTGRES_USER_DB_PASSWORD" \
    pg_dump "host=$KHULNASOFT_POSTGRES_USER_DB_HOST port=$KHULNASOFT_POSTGRES_USER_DB_PORT user=$KHULNASOFT_POSTGRES_USER_DB_USER dbname=$KHULNASOFT_POSTGRES_USER_DB_NAME sslmode=$KHULNASOFT_POSTGRES_USER_DB_SSLMODE" \
    -f /data/pg_data.dump

echo "Postgres backup saved to the file /data/pg_data.dump"
