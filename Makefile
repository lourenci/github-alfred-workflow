.DEFAULT_GOAL := help

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## Run all the tests
	@go test ./...

.PHONY: build-workflow
build-workflow: ## Copy all the source files to the workflow
	cp -R ./.assets ./.workflow/
	mkdir -p ./.workflow && rsync -av --exclude='*_test.go' --exclude='.git/' --exclude='.assets/' --exclude='.workflow/' . ./.workflow/
	zip -r "GitHub.alfredworkflow" ./.workflow && rm -rf ./.workflow

