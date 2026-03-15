#!/bin/bash

ENV=$1
TARGET=$2

if [[ "$ENV" != "dev" && "$ENV" != "prod" ]]; then
  echo "Error: Please provide environment parameter: dev or prod"
  exit 1
fi

ENV_FILE=".env.$ENV"

if [[ ! -f "$ENV_FILE" ]]; then
  echo "Error: Environment file $ENV_FILE does not exist."
  exit 1
fi

COMPOSE_FILE="docker-compose.yml"

if [[ "$TARGET" == "local" ]]; then
  COMPOSE_FILE="docker-compose.local.yml"
fi

echo "Starting docker compose with:"
echo "Environment: $ENV"
echo "Compose file: $COMPOSE_FILE"

docker compose --env-file "$ENV_FILE" -f "$COMPOSE_FILE" up -d