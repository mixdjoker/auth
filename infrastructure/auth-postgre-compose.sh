#!/bin/bash
compose_file_dir=${PWD}/infrastructure
compose_file_name=auth-postgre.yml

file_path=${compose_file_dir}/${compose_file_name}

# Startup Postgre in Docker

if [ ! -e "$file_path" ]; then
    echo "Compose file not exist"
    exit 1
fi

cmd_output=$(docker compose -f "${file_path}" up -d 2>&1)

if [[ $cmd_output == *"error"* ]]; then
  echo "Error in Docker Compose"
  echo "$cmd_output"
  exit 1
fi
