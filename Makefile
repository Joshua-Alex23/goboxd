.PHONY: build run test integration lint load

COMPOSE ?= docker compose
TOOLS   := $(COMPOSE) --profile tools run --rm tools

build:
	$(COMPOSE) build goboxd

run:
	$(COMPOSE) up goboxd

test:
	$(TOOLS) go test ./...

integration:
	$(COMPOSE) up -d goboxd
	$(TOOLS) go test -tags=integration ./tests/...

lint:
	$(TOOLS) golangci-lint run ./...

load:
	$(TOOLS) go test -run Load ./tests/...
