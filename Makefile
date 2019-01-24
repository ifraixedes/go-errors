GOLANGCI_LINT_VERSION := 1.12.3
GOLANGCI_LINT_BIN := $(realpath .bin/golangci-lint)

define HELP_MSG
Execute one of the following targets:

endef

export HELP_MSG

.PHONY: help
help: ## Show this help
	@echo "$$HELP_MSG"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/:.*##/:##/' | column -t -s '##'

.PHONY: lint
lint: ## Lint the code
	@$(GOLANGCI_LINT_BIN) run --enable-all --exclude-use-default=false

.PHONY: test
test: ## Execute the tests
	@go test $(TARGS) ./...

.PHONY: ci
ci: ## Simulate the same checks that the CI runs
	@make lint
	@make test

.PHONY: go-tools-install
go-tools-install: .gti-golangci-lint ## Install Go tools

.PHONY: .go-tools-install-ci
.go-tools-install-ci: .gti-golangci-lint

.PHONY: .gti-golangci-lint
.gti-golangci-lint:
	@mkdir -p .bin
	@curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b .bin v$(GOLANGCI_LINT_VERSION)
