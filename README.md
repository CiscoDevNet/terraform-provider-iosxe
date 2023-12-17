[![Tests](https://github.com/CiscoDevNet/terraform-provider-iosxe/actions/workflows/test.yml/badge.svg)](https://github.com/CiscoDevNet/terraform-provider-iosxe/actions/workflows/test.yml)

# Terraform Provider IOS-XE

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.20

## Building The Provider

1. Clone the repository
2. Enter the repository directory
3. Build the provider using the Go `install` command:

```shell
go install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```shell
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

This Terraform Provider is available to install automatically via `terraform init`. If you're building the provider, follow the instructions to
[install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin)
After placing it into your plugins directory,  run `terraform init` to initialize it.

Additional documentation, including available resources and their arguments/attributes can be found on the [Terraform documentation website](https://registry.terraform.io/providers/CiscoDevNet/iosxe/latest/docs).

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`. Make sure the respective environment variables are set (e.g., `IOSXE_USERNAME`, `IOSXE_PASSWORD`, `IOSXE_URL`).

*Note:* Acceptance tests create real resources.

```shell
make testacc
```
