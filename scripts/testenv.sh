#!/bin/bash

ENV_FILE=.env
docker compose -f docker-compose.yml --env-file .env.test up -d
