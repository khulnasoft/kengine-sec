#!/bin/bash
set -e

# wait for sql database connection
until pg_isready -h "${KHULNASOFT_POSTGRES_USER_DB_HOST}" -p "${KHULNASOFT_POSTGRES_USER_DB_PORT}" -U "${KHULNASOFT_POSTGRES_USER_DB_USER}" -d "${KHULNASOFT_POSTGRES_USER_DB_NAME}";
do
  echo >&2 "Postgres is unavailable - sleeping"
  sleep 5
done

# wait for neo4j to start
until nc -z ${KHULNASOFT_NEO4J_HOST} ${KHULNASOFT_NEO4J_BOLT_PORT};
do 
  echo "neo4j is unavailable - sleeping"
  sleep 5; 
done

# wait for kafka connection
until kcat -L -b ${KHULNASOFT_KAFKA_BROKERS};
do 
  echo "kafka is unavailable - sleeping"
  sleep 5;
done

# wait for file server to start
if [ "$KHULNASOFT_FILE_SERVER_HOST" != "s3.amazonaws.com" ]; then
  until nc -z "${KHULNASOFT_FILE_SERVER_HOST}" "${KHULNASOFT_FILE_SERVER_PORT}";
  do
    echo "file server is unavailable - sleeping"
    sleep 5;
  done
else
  echo "S3 mode skip file server health check"
fi

# for aws s3
fileServerProtocol="http"
if [ "$KHULNASOFT_FILE_SERVER_SECURE" == "true" ]; then
  fileServerProtocol="https"
fi

export GRYPE_DB_UPDATE_URL="${fileServerProtocol}://${KHULNASOFT_FILE_SERVER_HOST}:${KHULNASOFT_FILE_SERVER_PORT}/database/database/vulnerability/listing.json"
if [ "$KHULNASOFT_FILE_SERVER_HOST" == "s3.amazonaws.com" ]; then
  export GRYPE_DB_UPDATE_URL="${fileServerProtocol}://${KHULNASOFT_FILE_SERVER_DB_BUCKET}.s3.${KHULNASOFT_FILE_SERVER_REGION}.amazonaws.com/database/vulnerability/listing.json"
fi

# update vulnerability databae
if [ "$KHULNASOFT_MODE" == "worker" ]; then
  echo "add cron job to update vulnerability database"
  echo "vulnerability database update url $GRYPE_DB_UPDATE_URL"
  # /usr/local/bin/grype db update
  echo "0 */2 * * * export GRYPE_DB_UPDATE_URL=${GRYPE_DB_UPDATE_URL} && /usr/local/bin/grype db update" >> /etc/cron.d/crontab && chmod 0644 /etc/cron.d/crontab
  /usr/sbin/cron
fi

if [[ "${1#-}" != "$1" ]]; then
	set -- /usr/local/bin/khulnasoft_worker "$@"
fi

exec "$@"
