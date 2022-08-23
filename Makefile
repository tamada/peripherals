GO=go
SHELL=/bin/bash
VERSION := 0.9.6
NAME := peripherals
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
	@$(call _buildSubcommand,pskip)
	@$(call _buildSubcommand,ptake)
	@$(call _buildSubcommand,ptest)
	@$(call _buildSubcommand,puniq)

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

define _createDist
	echo -n "creating $(DIST)-$(1)-$(2).tar.gz..."
	mkdir -p dist/$(1)-$(2)/$(DIST)/bin
	cp README.md LICENSE dist/$(1)-$(2)/peripherals-${VERSION}/
	GOOS=$(1) GOARCH=$(2) go build -tags $(1) -o dist/$(1)-$(2)/$(DIST)/bin/pskip$(3) cmd/pskip/*.go
	GOOS=$(1) GOARCH=$(2) go build -tags $(1) -o dist/$(1)-$(2)/$(DIST)/bin/ptake$(3) cmd/ptake/*.go
	GOOS=$(1) GOARCH=$(2) go build -tags $(1) -o dist/$(1)-$(2)/$(DIST)/bin/ptest$(3) cmd/ptest/*.go
	GOOS=$(1) GOARCH=$(2) go build -tags $(1) -o dist/$(1)-$(2)/$(DIST)/bin/puniq$(3) cmd/puniq/*.go
	tar cfz dist/$(DIST)-$(1)-$(2).tar.gz -C dist/$(1)-$(2) $(DIST)
	echo "done"
endef

dist:
	@$(call _createDist,darwin,arm64,)
	@$(call _createDist,darwin,amd64,)
	@$(call _createDist,linux,arm64,)
	@$(call _createDist,linux,amd64,)
	@$(call _createDist,windows,amd64,.exe)
	@$(call _createDist,windows,386,.exe)

clean:
	$(GO) clean
	rm -rf bin dist