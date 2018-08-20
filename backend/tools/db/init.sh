#!/bin/sh

if [ "$1" != "dev" ] && [ "$1" != "prod" ]; then
  echo "error. \$1 undefined env (dev or prod)"
  exit 1
fi

ENV=$1
HOST=127.0.0.1
PORT=8306
DB=odpt_$ENV
USER=root
PASS=zaqroot

# mysql -u $USER -p$PASS -h $HOST -P $PORT -e "DROP DATABASE $DB"
# mysql -u $USER -p$PASS -h $HOST -P $PORT -e "CREATE DATABASE $DB"

go run main.go
