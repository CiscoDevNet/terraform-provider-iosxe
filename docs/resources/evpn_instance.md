---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_evpn_instance Resource - terraform-provider-iosxe"
subcategory: "EVPN"
description: |-
  This resource can manage the EVPN Instance configuration.
---

# iosxe_evpn_instance (Resource)

This resource can manage the EVPN Instance configuration.

## Example Usage

```terraform
resource "iosxe_evpn_instance" "example" {
  evpn_instance_num                    = 10
  vlan_based_replication_type_ingress  = false
  vlan_based_replication_type_static   = true
  vlan_based_replication_type_p2mp     = false
  vlan_based_replication_type_mp2mp    = false
  vlan_based_encapsulation             = "vxlan"
  vlan_based_auto_route_target         = false
  vlan_based_rd                        = "10:10"
  vlan_based_route_target_import       = "10:10"
  vlan_based_route_target_export       = "10:10"
  vlan_based_ip_local_learning_disable = false
  vlan_based_ip_local_learning_enable  = true
  vlan_based_default_gateway_advertise = "enable"
  vlan_based_re_originate_route_type5  = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `evpn_instance_num` (Number) evpn instance number
  - Range: `1`-`65535`

### Optional

- `device` (String) A device name from the provider configuration.
- `vlan_based_auto_route_target` (Boolean) Automatically set a route-target (OBSOLETE, use auto-route-target-boolean)
- `vlan_based_default_gateway_advertise` (String) Advertise Default Gateway MAC/IP routes
  - Choices: `disable`, `enable`
- `vlan_based_encapsulation` (String) Data encapsulation method
  - Choices: `mpls`, `vxlan`
- `vlan_based_ip_local_learning_disable` (Boolean) Disable IP local learning from dataplane
- `vlan_based_ip_local_learning_enable` (Boolean) Enable IP local learning from dataplane
- `vlan_based_rd` (String) ASN:nn or IP-address:nn
- `vlan_based_re_originate_route_type5` (Boolean) Re-originate route-type 5
- `vlan_based_replication_type_ingress` (Boolean) Ingress replication
- `vlan_based_replication_type_mp2mp` (Boolean) mp2mp replication
- `vlan_based_replication_type_p2mp` (Boolean) p2mp replication
- `vlan_based_replication_type_static` (Boolean) Static replication
- `vlan_based_route_target` (String) ASN:nn or IP-address:nn
- `vlan_based_route_target_both` (String) ASN:nn or IP-address:nn
- `vlan_based_route_target_export` (String) ASN:nn or IP-address:nn (Obsolete, use rt-value-entry)
- `vlan_based_route_target_import` (String) ASN:nn or IP-address:nn (Obsolete, use rt-value-entry)

### Read-Only

- `id` (String) The path of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import iosxe_evpn_instance.example "Cisco-IOS-XE-native:native/l2vpn/Cisco-IOS-XE-l2vpn:evpn_cont/evpn-instance/evpn/instance/instance=10"
```
