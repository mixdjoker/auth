#!/bin/bash
compose_file_dir=${PWD}/infrastructure
compose_file_name=auth-postgre.yml

file_path=${compose_file_dir}/${compose_file_name}

# Startup Postgre in Docker

if [ -e "$file_path" ]; then
    docker-compose -f "${file_path}" up -d
else
    exit 1
fi
