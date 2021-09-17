# Examples

This directory contains examples that are mostly used for documentation, but can also be run/tested manually via the Terraform CLI.

The document generation tool looks for files in the following locations by default. All other *.tf files besides the ones mentioned below are ignored by the documentation tool. This is useful for creating examples that can run and/or ar testable even if some parts are not relevant for the documentation.

* **examples_tf** example tf files in HCL
* **provider/provider.tf** example file for the provider index page
* **resources/<full resource name>/resource.tf** example file for the named data source page
* **yang_to_json** JSON-converted schema from the YANG modules
