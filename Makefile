# Common makefile commands & variables between projects
include .make/Makefile.common

# Common Golang makefile commands & variables between projects
include .make/Makefile.go

## Not defined? Use default repo name which is the application
ifeq ($(REPO_NAME),)
	REPO_NAME="go-pipl"
endif

## Not defined? Use default repo owner
ifeq ($(REPO_OWNER),)
	REPO_OWNER="mrz1836"
endif

.PHONY: clean

all: ## Runs lint, test-short and vet
	@$(MAKE) test-short

clean: ## Remove previous builds and any test cache data
	@go clean -cache -testcache -i -r
	@test $(DISTRIBUTIONS_DIR)
	@if [ -d $(DISTRIBUTIONS_DIR) ]; then rm -r $(DISTRIBUTIONS_DIR); fi

release:: ## Runs common.release then runs godocs
	@$(MAKE) godocs