include includes.mk

PHONY: install build lint env-up env-down run help

APPS ?= ws-epoll

.DEFAULT_GOAL := help

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

install: ## download dependencies
	@go mod download > /dev/null >&1

build:	install ## build binary
	@$(foreach APP, $(APPS), $(MAKE) -C $(APPS_DIR)/$(APP) build ;)

lint: bootstrap ## run golangci-linter
	$(GOLANGCI_LINT_BIN) run ./...

env-up:
	@docker-compose up -d
env-down:
	@docker-compose down -v

run: ## run
	@$(foreach APP, $(APPS), $(MAKE) -C $(APPS_DIR)/$(APP) run ;)

help:
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'