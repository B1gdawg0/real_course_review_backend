#!/bin/bash

ENV=$1

if [[ "$ENV" != "dev" && "$ENV" != "prod" ]]; then
  echo "Error: Please provide environment parameter: dev or prod"
  exit 1
fi

ENV_FILE=".env.$ENV"

if [[ ! -f "$ENV_FILE" ]]; then
  echo "Error: Environment file $ENV_FILE does not exist."
  exit 1
fi

echo "Starting docker-compose with environment: $ENV"
docker compose --env-file "$ENV_FILE" up -d
