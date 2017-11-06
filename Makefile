PKGS=$(shell go list ./... | grep -vF /vendor/)
GIT_SHA=$(shell git rev-parse --verify HEAD)
VERSION=$(shell git describe --tags --dirty)
DATE=$(shell date "+%Y-%m-%d")

.PHONY: all build install clean fmt vet lint test install-help dependencies help
.DEFAULT_GOAL := help

all: fmt vet lint megacheck test build ## Runs all the code checks, tests and build.

build: ## Create the github-cli executable in ./bin directory.
	go build -o bin/github-cli -ldflags "-X github.com/alok87/github-cli/cmd.gitSha=${GIT_SHA} -X github.com/alok87/github-cli/cmd.version=${VERSION} -X github.com/alok87/github-cli/cmd.buildDate=${DATE}"

install: build ## Build and create github-cli executable in $GOPATH/bin directory.
	install -m 0755 bin/github-cli ${GOPATH}/bin/github-cli

clean: ## Clean the project tree from binary files.
	rm -rf bin/*

fmt: ## Run go fmt ./...
	@if [ "`go fmt ./... | tee /dev/stderr`" ]; then \
		echo "^ improperly formatted go files" && echo && exit 1; \
	fi

vet: ## Apply go vet to all the Go files.
	@if [ "`go vet $(PKGS) | tee /dev/stderr`" ]; then \
		echo "^ go ver errors!" && echo && exit 1; \
	fi

megacheck: install-tools ## Apply megacheck to all the Go files.
	@if [ "`megacheck $(PKGS) | tee /dev/stderr`" ]; then \
		echo "^ megacheck errors!" && echo && exit 1; \
	fi

lint: install-tools ## Check for style mistakes in all the Go files using golint.
	@if [ "`golint $(PKGS) | tee /dev/stderr`" ]; then \
		echo "^ golint errors!" && echo && exit 1; \
	fi

test: ## Run the tests.
	go test -v $(PKGS)

install-tools:
	@GOLINT_CMD=$(shell command -v golint 2> /dev/null)
ifndef GOLINT_CMD
	go get github.com/golang/lint/golint
endif

	@MEGACHECK_CMD=$(shell command -v megacheck 2> /dev/null)
ifndef MEGACHECK_CMD
	go get honnef.co/go/tools/cmd/megacheck
endif

dependencies: ## Install all the dependencies.
	go get -u github.com/golang/dep/cmd/dep
	dep ensure -vendor-only

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
