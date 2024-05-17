#!/bin/bash

set +e

if [ ! -v KE_PROG_NAME ]; then
    echo "Environment variable KE_PROG_NAME not set. Set it to any string"
    exit 0
fi

if [ -d "/data" ]; then
    mkdir -p /data/$KE_PROG_NAME/data
#    rm -rf /var/log/postgresql/data
#    chown -R postgres:postgres /data/$KE_PROG_NAME
#    ln -s /data/$KE_PROG_NAME/data /var/lib/postgresql/
#    chown -R postgres:postgres /var/lib/postgresql
fi
