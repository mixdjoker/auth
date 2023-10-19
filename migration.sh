#!/bin/bash
source .env

export GOOSE_DBSTRING="host=${DB_HOST} port=5432 dbname=${DB_NAME} user=${DB_USER} password=${DB_PASSWORD} sslmode=disable"
export GOOSE_DRIVER=postgres

sleep 2 && goose -dir ${MIGRATION_DIR} up -v
