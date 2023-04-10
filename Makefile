setup: # Setup dependencies
	go mod tidy
.PHONY: setup

run: # Run
	go run main.go
.PHONY: run

build: # Build
	go build -o hupi
.PHONY: build

dist: # Creates and build dist folder
	goreleaser check
	goreleaser release --clean --snapshot
.PHONY: dist

generate: # Runs go generate
	go generate ./...
.PHONY: generate

format: # Run gofmt
	go fmt ./...
.PHONY: format

lint: # Run linter
	golangci-lint run ./...
.PHONY: lint

excluded := grep -v /res/ | grep -v /mocks/

test: # Test uses race and coverage
	go clean -testcache && go test -race $$(go list ./... | $(excluded)) -coverprofile=coverage.out -covermode=atomic
.PHONY: test

test-v: # Test with -v
	go clean -testcache && go test -race -v $$(go list ./... | $(excluded)) -coverprofile=coverage.out -covermode=atomic
.PHONY: test-v

mock: # Make mocks keeping directory tree
	rm -rf internal/mocks && mockery --all --output ./internal/mocks
.PHONY: mocks

doc: # Run go doc
	godoc -http localhost:8080
.PHONY: doc

all: # Make format, lint and test
	$(MAKE) format
	$(MAKE) lint
	$(MAKE) test
.PHONY: all

todo: # Show to-do items per file
	$(Q) grep \
		--exclude=Makefile.util \
		--exclude-dir=vendor \
		--exclude-dir=.idea \
		--text \
		--color \
		-nRo \
		-E '\S*[^\.]TODO.*' \
		.
.PHONY: todo

lines: # Show line count of Go code
	find . -name '*.go' | xargs wc -l
.PHONY: lines

help: # Display this help
	$(Q) awk 'BEGIN {FS = ":.*#"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?#/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help
