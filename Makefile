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

# Base init
PHONY: init
init:
	$(SILENT) rm -rf $(LOCAL_BIN)
	$(SILENT) mkdir -p $(LOCAL_BIN)
	$(SILENT) make install-golangci-lint

# Locals

clean:
	$(SILENT) rm -rf $(LOCAL_BIN)

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
