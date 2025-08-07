default: testall

# Run all acceptance tests
.PHONY: testall
testall:
	TF_ACC=1 go test -v $(TESTARGS) -timeout 60m ./internal/provider

# Run a subset of tests by specifying a regex (NAME).
# Usage: make test NAME=IosxeLogging
.PHONY: test
test:
	TF_ACC=1 go test ./... -v -run $(NAME) -timeout 60m

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