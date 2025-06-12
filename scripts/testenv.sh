#!/bin/bash

ENV_FILE=.env
docker compose -f docker/docker-compose.yml --env-file .env.test up -d
