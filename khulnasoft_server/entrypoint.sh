#!/bin/sh
set -e

if [ -z "$KHULNASOFT_POSTGRES_USER_DB_PASSWORD" ]; then
  export KHULNASOFT_POSTGRES_USER_DB_PASSWORD="khulnasoft"
fi

if [ -z "$KHULNASOFT_POSTGRES_USER_DB_USER" ]; then
  export KHULNASOFT_POSTGRES_USER_DB_USER="khulnasoft"
fi

until pg_isready -h "${KHULNASOFT_POSTGRES_USER_DB_HOST}" -p "${KHULNASOFT_POSTGRES_USER_DB_PORT}" -U "${KHULNASOFT_POSTGRES_USER_DB_USER}" -d "${KHULNASOFT_POSTGRES_USER_DB_NAME}"; 
do
  echo >&2 "Postgres is unavailable - sleeping"
  sleep 5
done

# check migrations complete
# psql -U ${KHULNASOFT_POSTGRES_USER_DB_USER} -d ${KHULNASOFT_POSTGRES_USER_DB_NAME} -t -c "SELECT EXISTS(SELECT name FROM role WHERE name = 'admin')"
export PGPASSWORD=${KHULNASOFT_POSTGRES_USER_DB_PASSWORD}
until psql -h "${KHULNASOFT_POSTGRES_USER_DB_HOST}" -U ${KHULNASOFT_POSTGRES_USER_DB_USER} -p "${KHULNASOFT_POSTGRES_USER_DB_PORT}" "${KHULNASOFT_POSTGRES_USER_DB_NAME}" -c '\q'; 
do
  echo >&2 "Database is unavailable - sleeping"
  sleep 5
done
echo >&2 "Database is available"

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

sed -i "s/https:\/\/petstore.swagger.io\/v2\/swagger.json/\/khulnasoft\/openapi.json/g" /usr/local/share/swagger-ui/swagger-initializer.js

exec "$@"
