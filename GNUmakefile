default: testall

# Load environment variables from .env file if it exists
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Run all acceptance tests across all devices and versions
.PHONY: testall
testall: gen test-1715-router test-1715-switch test-1712-router test-1712-switch
	@echo ""
	@echo "All multi-device tests completed!"

# Run a subset of tests by specifying a regex (NAME).
# Usage: make test NAME=IosxeLogging
.PHONY: test
test:
	TF_ACC=1 go test ./... -v -run $(NAME) -count 1 -timeout 60m

# Test against 17.15.x Router (C8000V)
# Usage: make test-1715-router [NAME=TestName] [DEBUG=1]
.PHONY: test-1715-router
test-1715-router:
	@echo "========================================="
	@echo "Testing against 17.15.x Router..."
	@echo "========================================="
	@if [ -z "$(IOSXE_1715_ROUTER_HOST)" ]; then \
		echo "SKIPPED: IOSXE_1715_ROUTER_HOST is not configured"; \
		echo "To enable this test, configure IOSXE_1715_ROUTER_HOST in your .env file"; \
	else \
		$(if $(DEBUG),echo "Debug mode enabled - logs will be written to test-output-1715-router.log";) \
		$(if $(NAME),echo "Running tests matching: $(NAME)";) \
		TF_ACC=1 \
		IOSXE_HOST=$(IOSXE_1715_ROUTER_HOST) \
		IOSXE_USERNAME=$(or $(IOSXE_1715_ROUTER_USERNAME),$(IOSXE_USERNAME)) \
		IOSXE_PASSWORD=$(or $(IOSXE_1715_ROUTER_PASSWORD),$(IOSXE_PASSWORD)) \
		IOSXE1715=1 \
		C8000V=1 \
		$(if $(DEBUG),TF_LOG=Trace) \
		go test -v $(if $(NAME),-run $(NAME)) $(TESTARGS) -count 1 -timeout 60m ./... $(if $(DEBUG),2>&1 | tee test-output-1715-router.log); \
	fi

# Test against 17.15.x Switch (C9000V)
# Usage: make test-1715-switch [NAME=TestName] [DEBUG=1]
.PHONY: test-1715-switch
test-1715-switch:
	@echo "========================================="
	@echo "Testing against 17.15.x Switch..."
	@echo "========================================="
	@if [ -z "$(IOSXE_1715_SWITCH_HOST)" ]; then \
		echo "SKIPPED: IOSXE_1715_SWITCH_HOST is not configured"; \
		echo "To enable this test, configure IOSXE_1715_SWITCH_HOST in your .env file"; \
	else \
		$(if $(DEBUG),echo "Debug mode enabled - logs will be written to test-output-1715-switch.log";) \
		$(if $(NAME),echo "Running tests matching: $(NAME)";) \
		TF_ACC=1 \
		IOSXE_HOST=$(IOSXE_1715_SWITCH_HOST) \
		IOSXE_USERNAME=$(or $(IOSXE_1715_SWITCH_USERNAME),$(IOSXE_USERNAME)) \
		IOSXE_PASSWORD=$(or $(IOSXE_1715_SWITCH_PASSWORD),$(IOSXE_PASSWORD)) \
		IOSXE1715=1 \
		C9000V=1 \
		$(if $(DEBUG),TF_LOG=Trace) \
		go test -v $(if $(NAME),-run $(NAME)) $(TESTARGS) -count 1 -timeout 60m ./... $(if $(DEBUG),2>&1 | tee test-output-1715-switch.log); \
	fi

# Test against 17.12.x Router (C8000V)
# Usage: make test-1712-router [NAME=TestName] [DEBUG=1]
.PHONY: test-1712-router
test-1712-router:
	@echo "========================================="
	@echo "Testing against 17.12.x Router..."
	@echo "========================================="
	@if [ -z "$(IOSXE_1712_ROUTER_HOST)" ]; then \
		echo "SKIPPED: IOSXE_1712_ROUTER_HOST is not configured"; \
		echo "To enable this test, configure IOSXE_1712_ROUTER_HOST in your .env file"; \
	else \
		$(if $(DEBUG),echo "Debug mode enabled - logs will be written to test-output-1712-router.log";) \
		$(if $(NAME),echo "Running tests matching: $(NAME)";) \
		TF_ACC=1 \
		IOSXE_HOST=$(IOSXE_1712_ROUTER_HOST) \
		IOSXE_USERNAME=$(or $(IOSXE_1712_ROUTER_USERNAME),$(IOSXE_USERNAME)) \
		IOSXE_PASSWORD=$(or $(IOSXE_1712_ROUTER_PASSWORD),$(IOSXE_PASSWORD)) \
		IOSXE1712=1 \
		C8000V=1 \
		$(if $(DEBUG),TF_LOG=Trace) \
		go test -v $(if $(NAME),-run $(NAME)) $(TESTARGS) -count 1 -timeout 60m ./... $(if $(DEBUG),2>&1 | tee test-output-1712-router.log); \
	fi

# Test against 17.12.x Switch (C9000V)
# Usage: make test-1712-switch [NAME=TestName] [DEBUG=1]
.PHONY: test-1712-switch
test-1712-switch:
	@echo "========================================="
	@echo "Testing against 17.12.x Switch..."
	@echo "========================================="
	@if [ -z "$(IOSXE_1712_SWITCH_HOST)" ]; then \
		echo "SKIPPED: IOSXE_1712_SWITCH_HOST is not configured"; \
		echo "To enable this test, configure IOSXE_1712_SWITCH_HOST in your .env file"; \
	else \
		$(if $(DEBUG),echo "Debug mode enabled - logs will be written to test-output-1712-switch.log";) \
		$(if $(NAME),echo "Running tests matching: $(NAME)";) \
		TF_ACC=1 \
		IOSXE_HOST=$(IOSXE_1712_SWITCH_HOST) \
		IOSXE_USERNAME=$(or $(IOSXE_1712_SWITCH_USERNAME),$(IOSXE_USERNAME)) \
		IOSXE_PASSWORD=$(or $(IOSXE_1712_SWITCH_PASSWORD),$(IOSXE_PASSWORD)) \
		IOSXE1712=1 \
		C9000V=1 \
		$(if $(DEBUG),TF_LOG=Trace) \
		go test -v $(if $(NAME),-run $(NAME)) $(TESTARGS) -count 1 -timeout 60m ./... $(if $(DEBUG),2>&1 | tee test-output-1712-switch.log); \
	fi

# Test all 17.15.x devices (router and switch)
.PHONY: test-1715
test-1715: gen test-1715-router test-1715-switch
	@echo ""
	@echo "All 17.15.x tests completed!"

# Test all 17.12.x devices (router and switch)
.PHONY: test-1712
test-1712: gen test-1712-router test-1712-switch
	@echo ""
	@echo "All 17.12.x tests completed!"

# Update a files from a single definition
# Usage: make gen NAME="Logging"
# NAME: The name of the definition, e.g. "Logging"
.PHONY: gen
gen:
	go run gen/load_models.go
	go run ./gen/generator.go "$(NAME)"
	go run golang.org/x/tools/cmd/goimports -w internal/provider/
	terraform fmt -recursive ./examples/
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
	go run gen/doc_category.go

# Update all files
.PHONY: genall
genall:
	go run gen/load_models.go
	go run ./gen/generator.go
	go run golang.org/x/tools/cmd/goimports -w internal/provider/
	terraform fmt -recursive ./examples/
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
	go run gen/doc_category.go