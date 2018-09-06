#!/bin/sh

if [ "$1" != "dev" ] && [ "$1" != "prod" ]; then
  echo "error. \$1 undefined env (dev or prod)"
  exit 1
fi

ENV=$1
HOST=127.0.0.1
DB=odpt_$ENV
USER=root
PASS=zaqroot

PORT=8306
if [ "$ENV" = "prod" ]; then
  PORT=5306
fi

mysql -u $USER -p$PASS -h $HOST -P $PORT -e "DROP DATABASE $DB"
mysql -u $USER -p$PASS -h $HOST -P $PORT -e "CREATE DATABASE $DB"

ENV=$ENV go run main.go
