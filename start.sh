#!/bin/sh

set -e
echo "load environment variables from env file"
source app.env                                                                                                                                                                                                                                                                      

echo "run db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"