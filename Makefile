# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# CONTAINER_TOOL defines the container tool to be used for building images.
# Be aware that the target commands are only tested with Docker which is
# scaffolded by default. However, you might want to replace it to use other
# tools. (i.e. podman)
CONTAINER_TOOL ?= docker

# Setting SHELL to bash allows bash commands to be executed by recipes.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk command is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

##@ Development

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: test
test: ## Run tests.
	bash scripts/test.sh

.PHONY: lint
lint: golangci-lint ## Run golangci-lint linter
	$(GOLANGCI_LINT) run

.PHONY: lint-fix
lint-fix: golangci-lint ## Run golangci-lint linter and perform fixes
	$(GOLANGCI_LINT) run --fix

##@ CLI

.PHONY: install-cli
build-cli: build ## Build CLI binary.
	ln -sf $(LOCALBIN)/quoxy-cli quoxy

.PHONY: uninstall-cli
uninstall-cli: ## Uninstall CLI binary.
	rm quoxy

##@ Build

.PHONY: clean-build
clean-build: install fmt vet ## Build manager binary.
	# Build all binaries
	@for d in cmd/*; do \
		if [ -d $$d ]; then \
			if [ ! -f $(LOCALBIN)/quoxy-$$(basename $$d) ]; then \
				go build -o $(LOCALBIN)/quoxy-$$(basename $$d) $$d/main.go; \
			fi; \
		fi; \
	done

.PHONY: build
build: clean ## Remove build artifacts.
	@for d in cmd/*; do \
    	go build -o $(LOCALBIN)/quoxy-$$(basename $$d) $$d/main.go; \
    done

.PHONY: clean
clean: fmt vet test lint ## Remove build artifacts.
	rm -rf $(LOCALBIN)/quoxy-*

.PHONY: run
run: clean-build ## Run controllers from your host.
	@if [ ! -f /tmp/quoxy-logs ]; then \
		mkfifo /tmp/quoxy-logs; \
	fi
	@export LOG_LEVEL=INFO; \
	cat /tmp/quoxy-logs & \
	for d in $(LOCALBIN)/quoxy-*; do \
		if [ $$(basename $$d) = "quoxy-cli" ]; then \
			continue; \
		fi; \
		( $$d 2>&1 | sed "s/^/[$$(basename $$d)] /" ) >> /tmp/quoxy-logs & \
	done

##@ Build Docker

IMG_BASE ?= quoxy
POSSIBLE_SERVICES = "proxy rest-api token-handler"

.PHONY: check-service
check-service:
	@if [ -z "$(service)" ]; then \
		echo "service is not defined"; \
		exit 1; \
	fi

	@if ! echo "$(POSSIBLE_SERVICES)" | grep -q "\b$$service\b"; then \
		echo "$(service) is not in the list of possible services"; \
		exit 1; \
	fi

.PHONY: docker-build
docker-build: check-service ## Build docker image with the manager.
	@echo "Building docker image for $(service)"
	$(CONTAINER_TOOL) build  --no-cache -t ${IMG_BASE}-$(service) --target $(service) .

.PHONY: docker-build-cache
docker-build-cache: check-service ## Build docker image with the manager.
	@echo "Building docker image for $(service)"
	$(CONTAINER_TOOL) build -t ${IMG_BASE}-$(service) --target $(service) .

.PHONY: docker-push
docker-push: docker-build ## Push docker image with the manager.
	$(CONTAINER_TOOL) push ${IMG_BASE}-$(service)

.PHONY: docker-run
docker-run: docker-build-cache
	docker rm ${IMG_BASE}-$(service) || true
	$(CONTAINER_TOOL) run --name ${IMG_BASE}-$(service) ${IMG_BASE}-$(service):latest

##@ Deployment

ifndef ignore-not-found
  ignore-not-found = false
endif

##@ Dependencies

.PHONY: install
install: ## Install dependencies
	go mod vendor

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Tool Binaries
ENVTEST ?= $(LOCALBIN)/setup-envtest-$(ENVTEST_VERSION)
GOLANGCI_LINT = $(LOCALBIN)/golangci-lint-$(GOLANGCI_LINT_VERSION)

## Tool Versions
ENVTEST_VERSION ?= release-0.18
GOLANGCI_LINT_VERSION ?= v1.57.2


.PHONY: golangci-lint
golangci-lint: $(GOLANGCI_LINT) ## Download golangci-lint locally if necessary.
$(GOLANGCI_LINT): $(LOCALBIN)
	$(call go-install-tool,$(GOLANGCI_LINT),github.com/golangci/golangci-lint/cmd/golangci-lint,${GOLANGCI_LINT_VERSION})

# go-install-tool will 'go install' any package with custom target and name of binary, if it doesn't exist
# $1 - target path with name of binary (ideally with version)
# $2 - package url which can be installed
# $3 - specific version of package
define go-install-tool
@[ -f $(1) ] || { \
set -e; \
package=$(2)@$(3) ;\
echo "Downloading $${package}" ;\
GOBIN=$(LOCALBIN) go install $${package} ;\
mv "$$(echo "$(1)" | sed "s/-$(3)$$//")" $(1) ;\
}
endef
