## Unreleased

- BREAKING CHANGE: Consolidate `iosxe_device_tracking_policy` into `iosxe_device_tracking` as a `policies` list attribute
- BREAKING CHANGE: Consolidate `iosxe_dhcp_pool` and `iosxe_ipv6_dhcp_pool` into `iosxe_dhcp` as `pools` and `ipv6_pools` list attributes
- BREAKING CHANGE: Consolidate `iosxe_evpn_profile` into `iosxe_evpn` as a `profiles` list attribute
- Add `iosxe_interface_vrrp_v2` resource and data source for legacy VRRPv2 (RFC 3768) interface-level configuration, including primary/secondary virtual IPv4 addresses, priority, preempt delay, advertisement timers, plaintext authentication, description, object tracking with priority decrement, and shutdown
- Add `iosxe_eigrp` resource and data source for EIGRP named mode global IPv4 unicast address-family configuration (`router eigrp <name>` / `address-family ipv4 unicast autonomous-system <asn>`)
- Add `iosxe_eigrp_vrf` resource and data source for EIGRP named mode VRF IPv4 unicast address-family configuration (`router eigrp <name>` / `address-family ipv4 unicast vrf <vrf> autonomous-system <asn>`)
- Add `tunnel_key`, `tunnel_mode_gre_multipoint`, `ip_nhrp_authentication`, `ip_nhrp_network_id`, `ip_nhrp_nhs`, `ip_nhrp_maps`, `ip_nhrp_redirect`, `ip_nhrp_shortcut`, and `mpls_nhrp` attributes to `iosxe_interface_tunnel` resource and data source for DMVPN/NHRP configuration
