#!/bin/bash

while IFS= read -r line || [[ -n "$line" ]]; do
  if [[ "$line" =~ ^[[:alpha:]_][[:alnum:]_]*= ]]; then
    export "$line"
  fi
done < .env

if [ "$1" == "up" ]; then
  migrate -database "${DATABASE_URL}" -path database/migrations up
elif [ "$1" == "down" ]; then
  migrate -database "${DATABASE_URL}" -path database/migrations down $2
elif [ "$1" == "force" ]; then
  if [ -z "$2" ]; then
    echo "Version is missing. Usage ./migrate.sh force [version]"
    exit 1
  fi
  migrate -database "${DATABASE_URL}" -path database/migrations force "$2"
elif [ "$1" == "create" ]; then
  if [ -z "$2" ]; then
    echo "File name is missing. Usage: ./migrate.sh create [file_name]"
    exit 1
  fi
  migrate create -ext sql -dir database/migrations -seq "$2"
else
  echo "Invalid command. Usage: ./migrate.sh [up|down|create [name]|force [version]]"
fi
