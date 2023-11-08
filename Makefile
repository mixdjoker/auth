include ./config/.local.env
include .credentials

LOCAL_BIN = $(CURDIR)/bin
CONF_DIR = $(CURDIR)/config
SOURCE_DIR = $(CURDIR)/cmd
INFRA_DIR = $(CURDIR)/infrastructure
GO_CMP_ARGS = CGO_ENABLED=0 GOEXPERIMENT="loopvar"

# Tools versions
GOLINT_VER = v1.55.2
PROTOC_GO_VER = v1.28.1
PROTOC_GRPC_VER = v1.2
GOOSE_VER = v3.14.0

# Dockerfiles
DEV_APP_DCFILE= dev-$(APP_NAME)-server.Dockerfile
PROD_APP_DCFILE= prod-$(APP_NAME)-server.Dockerfile
MIGRATION_DCFILE= $(APP_NAME)-migration.Dockerfile

# Docker Images
DEV_APP_IMAGE = dev-$(APP_NAME)-server
PROD_APP_IMAGE = prod-$(APP_NAME)-server
MIGRATION_IMAGE = $(APP_NAME)-migration

# Image tag from hash
GIT_SHORT_HASH := $(shell git rev-parse --short HEAD)

LOCAL_MIGRATION_DIR=$(MIGRATION_DIR)
LOCAL_MIGRATION_DSN="host=$(PG_HOST) port=$(PG_LOCAL_PORT) dbname=$(PG_DATABASE) user=$(PG_USER) password=$(PG_PASSWORD) sslmode=disable"

SILENT = @

# Install dependences
PHONY: install-deps
install-deps:
	$(SILENT) GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLINT_VER)
	$(SILENT) GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GO_VER)
	$(SILENT) GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GRPC_VER)
	$(SILENT) GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@$(GOOSE_VER)

# Download dependences
PHONY: get-deps
get-deps:
	$(SILENT) go get -u google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GO_VER)
	$(SILENT) go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GRPC_VER)

# Base init
PHONY: init
init:
	$(SILENT) rm -rf $(LOCAL_BIN)
	$(SILENT) mkdir -p $(LOCAL_BIN)
	make install-deps
	make get-deps

# API generation
PHONY: generate
generate:
	make generate-note-api

# API V1 generation
PHONY: generate-note-api
generate-note-api:
	protoc --proto_path api/user_v1 \
	--go_out=pkg/user_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go \
	--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc \
	api/user_v1/user.proto

# Local linter run
PHONY: lint
lint:
	$(SILENT) $(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

# Make build
PHONY: build
build:
	$(SILENT) $(GO_CMP_ARGS) go build -o $(LOCAL_BIN)/$(APP_NAME) $(SOURCE_DIR)

# Make run
PHONY: run
run:
	$(SILENT) $(GO_CMP_ARGS) go run $(SOURCE_DIR)

#################
## Git Section ##
#################

git-commit-all:
	git add .
	git commit -m "$(ARGS)"

# Hand deploy
PHONY: copy-to-server
copy-to-server:
	scp -i ~/.ssh/gopher $(LOCAL_BIN)/$(APP_NAME) gopher@course:

# Docker Build and Push
docker-build-and-push-app-dev:
	docker login -u oauth -p $(YA_TOKEN) $(YA_REGISTRY)
	docker buildx build --no-cache --platform linux/amd64 --push --tag $(YA_REGISTRY)/$(DEV_DB_IMAGE):$(GIT_SHORT_HASH) $(INFRA_DIR)/$(DEV_APP_DCFILE)

compose-build:
	docker login -u oauth -p $(YA_TOKEN) $(YA_REGISTRY)
	docker buildx build --no-cache --platform linux/amd64 --push --tag $(YA_REGISTRY)/$(APP_NAME)-server:$(GIT_SHORT_HASH) .

compose-db-config:
	docker compose -f ./infrastructure/auth-postgre.yml config

compose-db-up:
	docker compose -f ./infrastructure/auth-postgre.yml up -d

compose-db-down:
	docker compose -f ./infrastructure/auth-postgre.yml down
	
#################
## DBA Section ##
#################
local-migration-create:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=${LOCAL_MIGRATION_DSN} $(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} create $(ARGS) sql

local-migration-status:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=${LOCAL_MIGRATION_DSN} $(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} status -v

local-migration-up:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=${LOCAL_MIGRATION_DSN} $(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} up -v

local-migration-down:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=${LOCAL_MIGRATION_DSN} $(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} down -v

#################
## Black Magic ##
#################

# ARGS Reading
ARGS = $(filter-out $@,$(MAKECMDGOALS))
%:
	@:
