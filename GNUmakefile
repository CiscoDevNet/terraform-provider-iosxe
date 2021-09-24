TEST=./internal/provider
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
PKG_NAME=iosxe

default: build

tools: 
	go mod vendor

build: fmtcheck
	go install

test: fmtcheck
	cd $(TEST)
	go test ./... || exit 1
	echo ./... | \
		xargs -t -n4 go test $(TESTARGS) -timeout=2m -parallel=4

testacc: fmtcheck
	cd $(TEST)
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 2m

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	go test -v -c $(TEST) $(TESTARGS)

.PHONY: build test testacc vet fmt fmtcheck errcheck test-compile tools