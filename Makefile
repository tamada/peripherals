GO=go
SHELL=/bin/bash
VERSION := 1.0.0
DIST := $(NAME)-$(VERSION)

all: test build

setup:
	mkdir -p bin

test: setup
	$(GO) test --tags $$(go env GOOS) -covermode=count -coverprofile=coverage.out $$(go list ./...)

define _buildSubcommand
	$(GO) build --tags $$(go env GOOS) -o bin/$(1) cmd/$(1)/*.go
endef

build: setup
	@$(call _buildSubcommand,ptake)
	@$(call _buildSubcommand,puniq)
	@$(call _buildSubcommand,ptest)

lint: setup format
	$(GO) vet $$(go list ./...)
	for pkg in $$(go list ./...); do \
		golint -set_exit_status $$pkg || exit $$?; \
	done

format: setup
# $(go list -f '{{.Name}}' ./...) outputs the list of package name.
# However, goimports could not accept package name 'main'.
# Therefore, we replace 'main' to the go source code name 'rrh.go'
# Other packages are no problem, their have the same name with directories.
	goimports -w $$(go list ./... | sed 's/github.com\/tamada\/peripherals//g' | sed 's/^\///g')

clean:
	rm -rf bin