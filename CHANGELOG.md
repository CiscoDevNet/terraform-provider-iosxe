## 0.5.2

- Add `speed` options to `iosxe_interface_ethernet` resource and data source
- Add more source interface type options to `iosxe_aaa` resource and data source
- Add `ipv4_unicast_aggregate_addresses` attribute to `iosxe_bgp_address_family_ipv4` and `iosxe_bgp_address_family_ipv4_vrf` resources and data sources
- Add `soft_reconfiguration` attribute to `iosxe_bgp_ipv4_unicast_neighbor`, `iosxe_bgp_ipv4_unicast_vrf_neighbor`, `iosxe_bgp_ipv6_unicast_neighbor` and `iosxe_bgp_l2vpn_evpn_neighbor` resources and data sources
- Add `fall_over_bfd` attribute to `iosxe_bgp_neighbor` and `iosxe_bgp_ipv4_unicast_vrf_neighbor` resources and data sources
- Add `iosxe_as_path_access_list` resource and data source
- Add `default_originate` and `default_originate_route_map` attributes to `iosxe_bgp_ipv4_unicast_neighbor`, `iosxe_bgp_ipv4_unicast_vrf_neighbor` and `iosxe_bgp_ipv6_unicast_neighbor` resources and data sources
- Add `ip_bgp_community_new_format` attribute to `iosxe_system` resource and data source
- Add `community_list_standard` resource and data source
- Add `community_list_expanded` resource and data source
- Add `authentication`, `mab` and `dot1x` attributes to `iosxe_interface_ethernet` resource and data source
- Add `authorization_exec` and `transport_input` attributes to `iosxe_line` resource and data source
- Add `a2`, `a3` and `a4` attributes to `iosxe_aaa_authorization` resource and data source

## 0.5.1

- Add `iosxe_vlan_filter` resource and data source
- Add `iosxe_vlan_group` resource and data source
- Add `iosxe_save_config` resource
- Add `iosxe_cli` resource

## 0.5.0

- Add `iosxe_errdisable` resource and data source
- Add `iosxe_line` resource and data source
- Add `spanning_tree_link_type` and `ip_dhcp_snooping_trust` attributes to `iosxe_interface_port_channel` resource and data source
- Add `compress_config`, `sequence_numbers` and `call_home` attributes to `iosxe_service` resource and data source
- Add `hosts`, `system_shutdown` and `enable_traps_*` attributes to `iosxe_snmp_server` resource and data source
- BREAKING CHANGE: Rename `advertise_l2vpn_evpn` attribute to `ipv4_unicast_advertise_l2vpn_evpn` of `iosxe_bgp_address_family_ipv4_vrf` resource and data source
- BREAKING CHANGE: Rename `redistribute_connected` attribute to `ipv4_unicast_redistribute_connected` of `iosxe_bgp_address_family_ipv4_vrf` resource and data source
- BREAKING CHANGE: Rename `redistribute_static` attribute to `ipv4_unicast_redistribute_static` of `iosxe_bgp_address_family_ipv4_vrf` resource and data source
- Add `ipv4_unicast_redistribute_connected` attribute to `iosxe_bgp_address_family_ipv4` resource and data source
- Add `ipv4_unicast_redistribute_static` attribute to `iosxe_bgp_address_family_ipv4` resource and data source
- BREAKING CHANGE: Rename `advertise_l2vpn_evpn` attribute to `ipv6_unicast_advertise_l2vpn_evpn` of `iosxe_bgp_address_family_ipv6_vrf` resource and data source
- BREAKING CHANGE: Rename `redistribute_connected` attribute to `ipv6_unicast_redistribute_connected` of `iosxe_bgp_address_family_ipv6_vrf` resource and data source
- BREAKING CHANGE: Rename `redistribute_static` attribute to `ipv6_unicast_redistribute_static` of `iosxe_bgp_address_family_ipv6_vrf` resource and data source
- Add `ipv6_unicast_redistribute_connected` attribute to `iosxe_bgp_address_family_ipv6` resource and data source
- Add `ipv6_unicast_redistribute_static` attribute to `iosxe_bgp_address_family_ipv6` resource and data source
- When removing attributes from a resource (or setting them to `null`) which were previously set, the corresponding configuration will be removed from the device
- Add `ip_radius_source_interface_loopback` attribute to `iosxe_aaa` resources and data sources
- BREAKING CHANGE: Rename `group_tacacsplus` attribute to `group_server_tacacsplus` of `iosxe_aaa` resource and data source
- BREAKING CHANGE: Rename `servers` attribute to `server_names` of `iosxe_aaa` resource and data source
- Add `a1_group` and `a2_local` attributes to `iosxe_aaa_authorization` resources and data sources
- Add `arp_timeout` attribute to `iosxe_interface_tunnel` resources and data sources
- Add `iosxe_service_template` resource and data source
- Add `iosxe_vlan_access_map` resource and data source
- Add `dot1x_timeout_tx_period` and `service_policy_type_control_subscriber` attributes to `iosxe_template` resources and data sources
- When removing elements from a list attribute, the corresponding configuration on the device will be updated accordingly

## 0.4.0

- Fix issue when using `tunnel_destination_ipv4` or `tunnel_mode_ipsec_ipv4` attributes of `iosxe_interface_tunnel` resource
- Add `iosxe_static_route_vrf` resource and data source
- Make `lists.key` attribute of `iosxe_restconf` resource mandatory
- Fix issue with nested `lists.items` attributes of `iosxe_restconf` resource
- Add `iosxe_radius` resource and data source
- Add `iosxe_bfd_template_single_hop` resource and data source
- Add `iosxe_cdp` resource and data source
- Add `cluster_id` ,`fall_over` ,`disable_connected_check`, `local_as`, `log_neighbor_changes`, `password`, `timers`, `ttl_security` attributes to `iosxe_bgp_ipv4_unicast_vrf_neighbor` resource and data source
- Add `iosxe_tacacs_server` resource and data source
- Add `iosxe_bfd` resource and data source
- Add BFD attributes to `iosxe_interface_ethernet`, `iosxe_interface_port_channel`, `iosxe_interface_port_channel_subinterface`, `iosxe_interface_tunnel` and `iosxe_interface_vlan` resources and data sources
- BREAKING CHANGE: Rename `unreachables` attribute to `ip_unreachables` of `iosxe_interface_ethernet`, `iosxe_interface_port_channel`, `iosxe_interface_port_channel_subinterface`, `iosxe_interface_tunnel`, `iosxe_interface_loopback` and `iosxe_interface_vlan` resources and data sources
- BREAKING CHANGE: Rename `multicast_routing` attribute to `ip_multicast_routing` of `iosxe_system` resource and data source
- BREAKING CHANGE: Rename `multicast_routing_distributed` attribute to `ip_multicast_routing_distributed` of `iosxe_system` resource and data source
- Add `ipv6` attributes to `iosxe_interface_ethernet`, `iosxe_interface_port_channel` `iosxe_interface_vlan` `iosxe_interface_loopback` and `iosxe_interface_port_channel_subinterface` resources and data sources
- Add `iosxe_bfd_template_multi_hop` resource and data source
- Add `disable_connected_check`, `fall_over`, `local_as`, `log_neighbor_changes`, `password`, `timers` and `ttl_security` attributes to `iosxe_bgp_neighbor` resource and data source
- Add `ttl_security` and `process_ids` attributes to `iosxe_interface_ospf` resource and data source
- BREAKING CHANGE: Remove `iosxe_interface_ospf_process` resource and data source, functionality moved to `iosxe_interface_ospf` resource and data source
- Add `iosxe_arp` resource and data source
- Add `iosxe_class_map` resource and data source
- BREAKING CHANGE: Rename `ipv6_address_prefix_list` attribute to `ipv6_addresses` of `iosxe_interface_tunnel` resource and data source
- Add `snooping_information_option_format_remote_id_hostname` attribute to `iosxe_dhcp` resource and data source
- Add `iosxe_dot1x` resource and data source
- Add `arp_timeout`, `spanning_tree_link_type` and `spanning_tree_portfast_trunk` attribute to `iosxe_interface_ethernet` resource and data source
- Add `message_digest_keys` attribute to `iosxe_interface_ospf` resource and data source
- Add `areas` and `passive_interface_default` attributes to `iosxe_ospf_vrf` resource and data source
- Add `iosxe_policy_map` resource and data source
- Add `iosxe_policy_map_event` resource and data source
- Add `accounting_port`, `pac_key` and `automate_tester` attributes to `iosxe_radius` resource and data source
- Add `arp_timeout` attribute to `iosxe_interface_loopback` resource and data source
- Add `ip_arp_inspection` and `ip_dhcp_snooping` attributes to `iosxe_interface_ethernet` resource and data source
- Add `arp_timeout` and `ip_arp_inspection` attributes to `iosxe_interface_port_channel` and `iosxe_interface_port_channel_subinterface` resources and data sources
- Add `areas` attribute to `iosxe_ospf` resource and data source
- Add `iosxe_udld` resource and data source
- Add `iosxe_vtp` resource and data source
- BREAKING CHANGE: Rename `neighbor` attribute to `neighbors` of `iosxe_ospf` resource and data source
- BREAKING CHANGE: Rename `network` attribute to `networks` of `iosxe_ospf` resource and data source
- BREAKING CHANGE: Rename `summary_address` attribute to `summary_addresses` of `iosxe_ospf` resource and data source
- Add `ipv4_unicast_networks_mask` and `ipv4_unicast_networks` attribute to `iosxe_bgp_address_family_ipv4` and `iosxe_bgp_address_family_ipv4_vrf` resources and data sources
- Add `ipv6_unicast_networks` attribute to `iosxe_bgp_address_family_ipv6` and `iosxe_bgp_address_family_ipv6_vrf` resources and data sources

## 0.3.3

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
