<a href="https://terraform.io">
    <img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" alt="Terraform logo" title="Terraform" align="right" height="50" />
</a>

# Terraform Provider for Cisco IOS XE

The terraform-provider-iosxe is a plugin for Terraform that one can use with Terraform to work with Cisco IOS XE.

The following versions of Cisco IOS XE are supported:
- Cisco IOS XE 17.6.1

## Getting Started

- [Using the provider](docs/index.md)
- [Provider development](./DEVELOPMENT.md)

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