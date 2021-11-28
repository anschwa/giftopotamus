# GNUmakefile
SHELL := bash
.ONESHELL:
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
MAKEFLAGS += --warn-undefined-variables

HEROKU = /opt/heroku/bin/heroku

build = $(shell git describe --always --dirty)

.PHONY: help
help: ## Display this help section
	@echo -e "Usage: make <command>\n\nAvailable commands are:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "  %-38s %s\n", $$1, $$2}' ${MAKEFILE_LIST}
.DEFAULT_GOAL := help

.PHONY: test
test: ## Run all tests
	go test -cover ./...

.PHONY: deploy
deploy: ## Deploy production build to Heroku
	$(HEROKU) config:set APP_ENV="production"
	$(HEROKU) config:set GO_LINKER_SYMBOL="main.Build"
	$(HEROKU) config:set GO_LINKER_VALUE="$(build)"

	git push heroku main
	$(HEROKU) ps:restart
	$(HEROKU) logs --tail

.PHONY: build
build: ## Build website for production
	npm run prod
	go build -ldflags "-X main.Build=$(build)" -o cmd/app/app cmd/app/main.go

.PHONY: serve
serve: ## Start development server with tailwindcss
	npm run dev
	go run -ldflags "-X main.Build=$(build)" cmd/app/main.go

.PHONY: clean
clean:  ## Remove executable and tailwind.css
	-rm -f static/css/tailwind.css cmd/app/app cmd/mailer/mailer
