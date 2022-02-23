<a href="https://terraform.io">
    <img src=".github/terraform_logo.svg" alt="Terraform logo" title="Terraform" align="right" height="50" />
</a>

# Terraform Provider for Cisco IOS XE

The terraform-provider-iosxe is a plugin for Terraform that one can use with Terraform to work with Cisco IOS XE.

The following versions of Cisco IOS XE are supported:
- Cisco IOS XE 17.6.1

## Getting Started

- [Using the provider](docs/index.md)
- [Provider development](./DEVELOPMENT.md)
- [Intro to IOS XE slides](docs/resources/intro_to_terraform_video.pdf)


The primary use-case for the Cisco IOS XE provider is managing the following features:
1. [aaa-authentication](./examples/examples_tf/aaa-authentication.tf)
1. [aaa-authorization](./examples/examples_tf/aaa-authorization.tf)
1. [aaa-accounting](./examples/examples_tf/aaa-accounting.tf)
1. [acl](./examples/examples_tf/acl.tf)
1. [bgp](./examples/examples_tf/bgp.tf)
1. [cdp](./examples/examples_tf/cdp.tf)
1. [dhcp](./examples/examples_tf/dhcp.tf)
1. [emp](./examples/examples_tf/emp.tf)
1. [etherChannel](./examples/examples_tf/etherChannel.tf)
1. [hsrp](./examples/examples_tf/hsrp.tf)
1. [igmp](./examples/examples_tf/igmp.tf)
1. [igmp-proxy](./examples/examples_tf/igmp-proxy.tf)
1. [l3-subinterface](./examples/examples_tf/l3-subinterface.tf)
1. [line](./examples/examples_tf/line.tf)
1. [mdt](./examples/examples_tf/mdt.tf)
1. [nat](./examples/examples_tf/nat.tf)
1. [ntp](./examples/examples_tf/ntp.tf)
1. [ospf](./examples/examples_tf/ospf.tf)
1. [pim](./examples/examples_tf/pim.tf)
1. [poe](./examples/examples_tf/poe.tf)
1. [radius](./examples/examples_tf/radius.tf)
1. [snmp](./examples/examples_tf/snmp.tf)
1. [span-rspan](./examples/examples_tf/span-rspan.tf)
1. [vlan](./examples/examples_tf/vlan.tf)
1. [vlan-trunk](./examples/examples_tf/vlan-trunk.tf)
1. [vlan-voice](./examples/examples_tf/vlan-voice.tf)
1. [vtp](./examples/examples_tf/vtp.tf)

## Creating Additional Terraform Resources
Any feature or Remote Procedure Call (RPC) supported by RESTCONF & YANG is also supported by Terraform. If a particular feature is not yet in this GitHub repository, you can create the necessary Terraform file using these steps
1.	Configure the feature as per the CLI config guide, if needed.
1.	Execute the feature's RESTCONF XPATH and provide any necessary JSON. 
    -	You can find the JSON for features currently configured device using `show run | format restconf-json`
    -	An alternate approach to find the RESTCONF JSON can be done using [YANG Suite](https://github.com/CiscoDevNet/yangsuite), a tool to visualize and understand YANG models
    ![](restconf_with_yang_suite.gif)
1.  The resulting JSON found from executing RESTCONF can be used to create the .tf file. For example, replace each of the values in angle brackets (<>) in the example Terraform file below with the corresponding Xpath and JSON:

    example.tf
    ```
    resource "iosxe_rest" "feature_put" {
        method = "PUT"
        path   = <RESTCONF_XPATH>
        payload = jsonencode(    
            {
                <JSON_RESPONSE>
            }
        )
    }
    ```
