# Setup
setup:
	go mod tidy
.PHONY: setup

# Run gofmt
format:
	go fmt ./...
.PHONY: format

# Run linter
lint:
	golangci-lint run ./...
.PHONY: lint

# Test uses race and coverage
test:
	go clean -testcache && go test -race $$(go list ./... | grep -v tests | grep -v mocks) -coverprofile=coverage.out -covermode=atomic
.PHONY: test

# Test with -v
test-v:
	go clean -testcache && go test -race -v $$(go list ./... | grep -v tests | grep -v mocks) -coverprofile=coverage.out -covermode=atomic
.PHONY: test-v

# Run all the tests and opens the coverage report
cover: test
	go tool cover -html=coverage.out
.PHONY: cover

# Make mocks keeping directory tree
mocks:
	rm -rf internal/mocks && mockery --all --keeptree --output ./internal/mocks
.PHONY: mocks

# Run go doc
doc:
	godoc -http localhost:8080
.PHONY: doc

# Make format, lint and test
all:
	$(MAKE) format
	$(MAKE) lint
	$(MAKE) test
