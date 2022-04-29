SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

.PHONY: help
help_spacing := 20
help: ## List all available targets with help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-$(help_spacing)s\033[0m %s\n", $$1, $$2}'

.PHONY: tidy
tidy: ## Tidying all project go modules
	@find . -name go.mod -execdir sh -c 'echo "$(shell pwd): go mod tidy" && go mod tidy' \;

.PHONY: generate
generate: ## Run code generation
	go generate ./...

.PHONY: outdated
outdated: ## Print outdated dependencies (`go install github.com/psampaz/go-mod-outdated@latest` required)
	@find . -name go.mod -execdir sh -c 'echo "$(shell pwd): outdated:" && go list -u -m -json all | go-mod-outdated -update -direct' \;

.PHONY: .update
.update:
	@find . -name go.mod -execdir sh -c 'echo "$(shell pwd): go get -u ./..." && go get -u ./...' \;

.PHONY: update
update: .update tidy generate ## Update go.mod dependencies

.PHONY: lint
lint: ## Run golangci-linter
	@find . -name go.mod -execdir sh -c 'echo "$(shell pwd): lint:" && golangci-lint run' \;

.PHONY: test
test: lint ## Run tests
	go test -short -v ./...

.PHONY: docker-builddocker images
docker-build: ## Build docker image
	docker build --file build/go/Dockerfile --tag training-scheduler --build-arg appName=training_scheduler --build-arg appVersion=0.0.1 .

.PHONY: docker-build-frontend
docker-build-frontend:
	docker build --file build/frontend/Dockerfile --tag training-scheduler-frontend .

.PHONY: migration
migration:
	goose -dir ./internal/migrations create $(name) go