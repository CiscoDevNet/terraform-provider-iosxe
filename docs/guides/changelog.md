---
subcategory: "Guides"
page_title: "Changelog"
description: |-
    Changelog
---

# Changelog

## Unreleased

- Add `mld_snooping` and `mld_snooping_querier` attributes to the `iosxe_system` resource and data sources
- Add `iosxe_interface_bdi` resource and data sources
- Add `service_instances` attribute to `iosxe_interface_ethernet` resource and data sources
- Add `iosxe_bridge_domain` resource and data sources
- Add `trunk_allowed_vlans_all`, `trunk_allowed_vlans_none`, `trunk_allowed_vlans_add`, `trunk_allowed_vlans_except`, and `trunk_allowed_vlans_remove` attributes to `iosxe_interface_switchport` resource and data source using non-deprecated YANG vlan-v2 paths
- Add `trunk_allowed_vlans_legacy` and `trunk_allowed_vlans_none_legacy` attributes to `iosxe_interface_switchport` resource and data source for backward compatibility with deprecated YANG paths
- Add `icmp_echo_frequency` attribute to `iosxe_sla` resource and data source
- Add `iosxe_commit` action
- Add `iosxe_save_config` action
- Add `igmp_snooping_querier`, `igmp_snooping_querier_version`, `igmp_snooping_querier_max_response_time`, and `igmp_snooping_querier_timer_expiry` attributes to `iosxe_system` resource and data source
- BREAKING CHANGE: Rename `buffered_size` attribute of `iosxe_logging` resource to `buffered_size_legacy`
- BREAKING CHANGE: Rename `buffered_severity` attribute of `iosxe_logging` resource to `buffered_severity_legacy`
- Add `buffered_size` and `buffered_severity` attributes to `iosxe_logging` resource and data source
- Add `vnids` attribute to `iosxe_vrf` resource and data source
- Add `iosxe_evpn_profile` resource and data source
- Add `ip_tcp_path_mtu_discovery`, `ip_tcp_mss` and `ip_tcp_window_size` attributes to `iosxe_system` resource and data source

## 0.13.0


- Add `stopbits` attribute to `iosxe_line` resource and data source
- Add `ipv4_evpn_mcast_*` and `ipv6_evpn_mcast_*` attributes to `iosxe_vrf` resource and data source
- Add `ipv4_import_map`, `ipv4_export_map`, `ipv6_import_map` and `ipv6_export_map` attributes to `iosxe_vrf` resource and data source
- Add `inherit_peer_policy` attribute to `iosxe_bgp_ipv4_unicast_neighbor` and `iosxe_bgp_l2vpn_evpn_neighbor` resources and data source
- Add `network_point_to_point` attribute to `iosxe_interface_isis` resource and data source
- Add device level `protocol` provider attribute
- Add `enable_traps_bgp_cbgp2_state_changes` and `enable_traps_bgp_cbgp2_threshold_prefix` attributes to `iosxe_snmp_server` resource and data source
- Add `tunnel_protection_ipsec_profile_legacy` attribute to `iosxe_interface_tunnel` resource and data source
- Add `destination_ip_vrf` attribute to `iosxe_flow_exporter` resource and data source
- Add `collect_interface_input` attribute to `iosxe_flow_record` resource and data source
- Add `ip_tcp_adjust_mss` attribute to `iosxe_interface_tunnel` resource and data source
- Add `match_cos` attribute to `iosxe_class_map` resource and data source

## 0.12.0

- Add `iosxe_l2_vfi` resource and data source
- Add `ip_local_proxy_arp` attribute to `iosxe_interface_vlan` resource and data source
- Add `iosxe_crypto` resource and data source
- Add `rd_auto` attribute to `iosxe_vrf` resource and data source
- Add `ipv6_multicast_routing` attribute to `iosxe_system` resource and data source
- Add `iosxe_mpls` resource and data source
- Add `iosxe_bgp_address_family_vpnv4` resource and data source
- Add `iosxe_bgp_address_family_vpnv6` resource and data source
- Add `ipv4_mdt_*` attributes to `iosxe_vrf` resource and data source
- Add `event_syslog_*` attributes to `iosxe_eem` resource and data source

## 0.11.0

- Fix `iosxe_yang` resource payload ordering with NETCONF, [link](https://github.com/CiscoDevNet/terraform-provider-iosxe/issues/372)
- Add `ip_igmp_version` attribute to `iosxe_interface_ethernet`, `iosxe_interface_loopback`, `iosxe_interface_port_channel`, `iosxe_interface_port_channel_subinterface`, `iosxe_interface_tunnel`, and `iosxe_interface_vlan` resources and data sources
- Add `ip_default_gateway` attribute to `iosxe_system` resource and data source for default gateway configuration on non-routing devices

