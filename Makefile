LOCAL_BIN = $(CURDIR)/bin
GOLINT_VER = 1.53.3
APP_NAME = auth
# APP_BIN_DIR = $(LOCAL_BIN)/$(app)
SOURCE_DIR = $(CURDIR)/cmd
GO_CMP_ARGS = CGO_ENABLED=0 GOEXPERIMENT="loopvar"

SILENT = @

# Linter installation
PHONY: install-golangci-lint
install-golangci-lint:
	$(SILENT) GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLINT_VER)

# Protoc local installation
PHONY: install-deps
install-deps:
	$(SILENT) GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	$(SILENT) GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2


PHONY: get-deps
get-deps:
	$(SILENT) go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	$(SILENT) go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

# Base init
PHONY: init
init:
	$(SILENT) rm -rf $(LOCAL_BIN)
	$(SILENT) mkdir -p $(LOCAL_BIN)
	make install-golangci-lint
	make install-deps
	make get-deps

PHONY: generate
generate:
	make generate-note-api

PHONY: generate-note-api
generate-note-api:
	rm -rf ./pkg/note_v1
	mkdir -p pkg/note_v1
	protoc --proto_path api/user_v1 \
	--go_out=pkg/user_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/user_v1/user.proto

# Build & Run

clean-bin:
	$(SILENT) rm -f $(LOCAL_BIN)/$(APP_NAME)

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
