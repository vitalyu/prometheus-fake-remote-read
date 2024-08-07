dummy := $(shell touch .env)
include .env
export

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Prepare

.PHONY: tidy
tidy: ## Pull up dependecies
	go mod tidy -v

.PHONY: lint
lint: ## Run linting
	golangci-lint run --tests=false
	golangci-lint run --disable-all -E golint,goimports,misspell

##@ Prepare

.PHONY: test
test: ## Run tests
	go test -race -cover -v ./...

.PHONY: run
run: ## Run for current platform
	CGO_ENABLED=0 go run ./cmd/prometheus-fake-remote-read --config configs/example.config.json

.PHONY: demo
demo: ## Run docker compose and checkout http://127.0.0.1:9090
	docker compose --project-directory ./demo/ up
	docker compose --project-directory ./demo/ rm --stop --volumes --force

##@ Build

.PHONY: build
build: ## Build for current platform
	mkdir ./bin/ || true
	CGO_ENABLED=0 go build -o ./dist/ ./cmd/prometheus-fake-remote-read 

.PHONY: goreleaser
goreleaser: ## Build via goreleaser
	goreleaser release --snapshot  --clean 
