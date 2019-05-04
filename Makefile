# Project
BINARY := versionbump
VERSION ?= $(shell cat VERSION)
COMMIT := $(shell git rev-parse --short HEAD)
BUILD_INFO := $(COMMIT)-$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
BRANCH := $(shell git status -b -u no | awk 'NR==1{print $3;}')

# Utilities
DOTENV := godotenv -f $(HOME)/.env,.env

# Variables
OS = $(word 1, $@)

# Go
GOMODOPTS = GO111MODULE=on
GOGETOPTS = GO111MODULE=off
GODIRS = $(shell find . -type d -maxdepth 1 -mindepth 1 | egrep 'cmd|internal|pkg|api')

.PHONY: build _build _build_xcompile browsetest cattest clean deps _deps depsdev deploy lint release \
        test _test _test_setup

cmd/versionbump/version.go: VERSION
	@echo "// This file is generated do not edit" > $@
	@echo "package main" >> $@
	@echo "var Version = \"$(VERSION)-$(COMMIT)\"" >> $@

build: cmd/versionbump/version.go
	@$(DOTENV) make _build

_build:
	go fmt ./...
ifneq ($(XCOMPILE),true)
	@make _build_simple
else
	@make _build_xcompile
endif

_build_simple:
	go build -ldflags "-X main.Version=$(VERSION)-$(BUILD_INFO)" ./cmd/$(BINARY)

_build_xcompile:
	make $(PLATFORMS)

PLATFORMS := linux darwin
.PHONY: $(PLATFORMS)
$(PLATFORMS):
	@mkdir -p dist
	GOOS=$(OS) GOARCH=amd64 CGO_ENABLED=0 go build \
	-o dist/$(BINARY)-$(OS)-$(VERSION) \
	-a -installsuffix cgo \
	-ldflags "-X main.Version=$(VERSION)-$(BUILD_INFO)" \
		./cmd/$(BINARY)

clean:
	rm -rf dist reports tmp vendor versionbump

deps:
	@$(DOTENV) make _deps

_deps:
	$(GOMODOPTS) go mod tidy
	$(GOMODOPTS) go mod vendor
	$(GOGETOPTS) make depsdev

depsdev:
	@$(DOTENV) make $(GOGETS)

GOGETS := github.com/jstemmer/go-junit-report github.com/golangci/golangci-lint/cmd/golangci-lint \
		  github.com/ains/go-test-html github.com/fzipp/gocyclo github.com/joho/godotenv/cmd/godotenv \
		  github.com/stretchr/testify
.PHONY: $(GOGETS)
$(GOGETS):
	go get -u $@

test:
	@$(DOTENV) make _test

_test:
	make _test_setup
	@mkdir -p reports/html
	@echo "### Unit Tests"; \
	go test -covermode atomic -coverprofile=./reports/coverage.out -v ./... 2>&1 | tee reports/test.txt; \
	EXITCODE="$${PIPESTATUS[0]}"; \
	cat ./reports/test.txt | go-junit-report > reports/junit.xml; \
	echo "### Code Coverage"; \
	go tool cover -func=./reports/coverage.out | tee ./reports/coverage.txt; \
	go tool cover -html=reports/coverage.out -o reports/html/coverage.html; \
	echo "### Cyclomatix Complexity Report"; \
	gocyclo -avg $(GODIRS) | grep -v _test.go | tee reports/cyclocomplexity.txt; \
	exit $(EXITCODE)

_test_setup:
	@mkdir -p tmp
	@cp -r test/fixtures tmp/testfiles
	@-chmod 600 tmp/testfiles/version_no_access.txt 2>&1 > /dev/null
	@cp tmp/testfiles/{version.txt,version_no_access.txt}
	@chmod 000 tmp/testfiles/version_no_access.txt

deploy:
	@echo TODO

release:
ifeq ($(BRANCH),master)
	git fetch --tags
	git tag $(VERSION) && git push origin :refs/tags/$(VERSION)
endif

lint:
	golangci-lint run --enable=gocyclo

REPORTS := reports/html/coverage.html
.PHONY: $(REPORTS)
$(REPORTS):
	@test -f $@ && open $@

browsetest: $(REPORTS)

cattest:
	@echo "### Unit Tests"
	@cat reports/test.txt
	@echo "### Code Coverage"
	@cat ./reports/coverage.txt
	@echo "### Cyclomatix Complexity Report"
	@cat reports/cyclocomplexity.txt
