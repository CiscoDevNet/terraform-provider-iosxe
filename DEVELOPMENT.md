# Development Environment Setup

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) (0.14, 0.15, 1.0.0)
- [Go](https://golang.org/doc/install) 1.17 (to build the provider plugin)

## Quick Start

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (please check the [requirements](#requirements) before proceeding).

*Note:* This project uses [Go Modules](https://blog.golang.org/using-go-modules) making it safe to work with it outside of your existing [GOPATH](http://golang.org/doc/code.html#GOPATH). The instructions that follow assume a directory in your home directory outside of the standard GOPATH (i.e `$HOME/providers/`).

Clone repository to: `$HOME/providers/`

```sh
$ mkdir -p $HOME/providers/; cd $HOME/providers/
$ git clone github.com/CiscoDevNet/terraform-provider-iosxe
...
```

Enter the provider directory and run `make tools`. This will install the needed tools for the provider.

```sh
$ make tools
```

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-iosxe
...
```

# Using the Provider

Move the generated binary from the build step to the [plugin directory](https://www.terraform.io/docs/cli/config/config-file.html#implied-local-mirror-directories)/registry.terraform.io/CiscoDevNet/iosxe/`<version>`/`<os>_<arch>`. Examples for `<os>_<arch>` are `windows_amd64`, `linux_arm`, `darwin_amd64`, etc. After placing it into your plugins directory, run `terraform init` to initialize it.

*Note:* Make sure `HOST_IOSXE`, `DEVICE_USERNAME_IOSXE`, and `DEVICE_PASSWORD_IOSXE` variables are set. For windows, copy env.example.bat to env.bat and replace dummy values with credentials and then execute the bat file with command promt.

Example
```hcl
terraform {
    required_providers {
        iosxe = {
        version = "0.1"
        source  = "CiscoDevNet/iosxe"
        }
    }
}

provider "iosxe" {
    request_timeout = 30
    insecure = true
}

resource "iosxe_rest" "vlan_example" {
    method = "POST"
    path = "/data/Cisco-IOS-XE-native:native/vlan"
    payload = jsonencode(
        {
            "Cisco-IOS-XE-vlan:vlan-list": {
            "id": "1",
            "name": "VLAN1"
            }
        }
    )
}
```

## Testing the Provider

In order to test the provider, you can run `make test`.

*Note:* Make sure `HOST_IOSXE`, `DEVICE_USERNAME_IOSXE`, and `DEVICE_PASSWORD_IOSXE` variables are set. For windows, copy env.example.bat to env.bat and replace dummy values with credentials and then execute the bat file with command promt.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

```sh
$ make testacc
```

## Debugging and Troubleshooting

- Set environment variable `TF_LOG` to one of the log levels `TRACE`, `DEBUG`, `INFO`, `WARN` or `ERROR`
- Set environment variable `TF_LOG_PATH` to write logs in a file. e.g. `TF_LOG_PATH=tf.log`

For more details visit - [Terraform Debugging](https://www.terraform.io/docs/internals/debugging.html)