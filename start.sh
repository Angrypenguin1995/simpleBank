#!/bin/sh

set -e

echo "load environment variables from env file "

echo "start the app"
exec "$@"