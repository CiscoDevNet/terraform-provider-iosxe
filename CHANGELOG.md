## 0.3.3 (unreleased)

- Add `iosxe_aaa` resource and data source
- Add `iosxe_aaa_accounting` resource and data source
- Add `iosxe_aaa_authorization` resource and data source
- Add `iosxe_interface_mpls` resource and data source
- Add `iosxe_interface_ospfv3` resource and data source
- Add `iosxe_interface_tunnel` resource and data source
- Add `iosxe_crypto_ipsec_transform_set` resource and data source
- Add `iosxe_aaa_authentication` resource and data source
- Add `iosxe_crypto_ikev2` resource and data source
- Add `iosxe_crypto_ikev2_proposal` resource and data source
- Add `iosxe_crypto_ipsec_profile` resource and data source
- Add `iosxe_radius_server` resource and data source
- Add `iosxe_crypto_ikev2_keyring` resource and data source
- Add `iosxe_crypto_ikev2_policy` resource and data source

## 0.3.2

- Add `auto_qos` attributes to `iosxe_interface_ethernet`, `iosxe_interface_port_channel` and `iosxe_interface_port_channel_subinterface` resources and data sources
- Add `spanning_tree_guard` attribute to `iosxe_interface_ethernet` and `iosxe_interface_port_channel` resources and data sources
- Add `trust_device` attribute to `iosxe_interface_ethernet`, `iosxe_interface_port_channel` and `iosxe_interface_port_channel_subinterface` resources and data sources
- Add `trunk_allowed_vlans_none` attribute to `iosxe_interface_switchport` and `iosxe_template` resources and data sources
- Add `trunk_allowed_vlans_all` attribute to `iosxe_template` resource and data source
- Add `ebgp_multihop` attributes to `iosxe_bgp_neighbor` and `iosxe_bgp_ipv4_unicast_vrf_neighbor` resources and data sources

## 0.3.1

- Fix issue with deletion of servers and peers of `iosxe_ntp` resource

## 0.3.0

- BREAKING CHANGE: Completely revamped the provider based on `github.com/netascode/terraform-provider-iosxe` codebase, replacing all existing resources and data sources
- BREAKING CHANGE: Remove `attributes` map of list items in `iosxe_restconf` resource

## 0.1.1

## 0.1.0

- Initial release
