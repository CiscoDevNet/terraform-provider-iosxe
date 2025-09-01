---
subcategory: "Guides"
page_title: "Changelog"
description: |-
    Changelog
---

# Changelog

## 0.8.1 (unreleased)

- Add `ip_nbar_classification_dns_classify_by_domain` attribute to `iosxe_system` resource and data source

## 0.8.0

- BREAKING CHANGE: Combine `iosxe_msdp` and `iosxe_msdp_vrf` resources and data sources into single `iosxe_msdp` resource and data source
- BREAKING CHANGE: Combine `iosxe_pim` and `iosxe_pim_vrf` resources and data sources into single `iosxe_pim` resource and data source
- BREAKING CHANGE: Rename `iosxe_static_route_vrf` resource and data source to `iosxe_static_routes_vrf`
- BREAKING CHANGE: Combine `iosxe_snmp_server`, `iosxe_snmp_server_group` and `iosxe_snmp_server_user` resources and data sources into single `iosxe_snmp_server` resource and data source
- Change`permit_entries` and `deny_entries` attributes of `iosxe_community_list_standard` resource and data source from Lists to Sets
- Add `accept_agreement`, `accept_end`, `accept_user`, `udi_pid`, `udi_sn`, `feature_name`, `feature_port_bulk`, `feature_port_onegig`, `feature_port_b_6xonegig`, and `feature_port_tengig` attributes to `iosxe_license` resource and data source
- Add `enable_traps_bgp`, `enable_traps_cbgp2`, `enable_traps_ospfv3_errors`, `enable_traps_ospfv3_state_change`, `vrf_hosts`, `ip_address`, `vrf`, `community_or_user`, `version`, `encryption`, and `security_level` attributes to `iosxe_snmp_server` resource and data source
- Add `level`, `list_name`, `action_type`, `broadcast`, `group_broadcast`, `group_logger`, `group1_group`, `group2_group`, `group3_group`, `group4_group`, `name`, `default`, `none`, `start_stop_broadcast`, `start_stop_logger`, `start_stop_group1`, `start_stop_group2`, `start_stop_group3`, `start_stop_group4`, `stop_only_broadcast`, `stop_only_logger`, `stop_only_group1`, `stop_only_group2`, `stop_only_group3`, `stop_only_group4`, `wait_start_broadcast`, `wait_start_logger`, `wait_start_group1`, `wait_start_group2`, `wait_start_group3`, `wait_start_group4`, `name`, `none`, `start_stop_broadcast`, `start_stop_logger`, `start_stop_group1`, `start_stop_group2`, `start_stop_group3`, `start_stop_group4`, `stop_only_broadcast`, `stop_only_logger`, `stop_only_group1`, `stop_only_group2`, `stop_only_group3`, `stop_only_group4`, `wait_start_broadcast`, `wait_start_logger`, `wait_start_group1`, `wait_start_group2`, `wait_start_group3`, and `wait_start_group4` attributes to `iosxe_aaa_accounting` resource and data source
- Add `enable_default_group1_cache`, `enable_default_group1_enable`, `enable_default_group1_group`, `enable_default_group1_line`, `enable_default_group1_none`, `enable_default_group2_cache`, `enable_default_group2_enable`, `enable_default_group2_group`, `enable_default_group2_line`, `enable_default_group2_none`, `enable_default_group3_cache`, `enable_default_group3_enable`, `enable_default_group3_group`, `enable_default_group3_line`, `enable_default_group3_none`, `enable_default_group4_cache`, `enable_default_group4_enable`, `enable_default_group4_group`, `enable_default_group4_line`, and `enable_default_group4_none` attributes to `iosxe_aaa_authentication` resource and data source
- Add `level`, `list_name`, `a1_group`, `a1_local`, `a1_if_authenticated`, `a1_none`, `a1_radius`, `a1_tacacs`, `a2_group`, `a2_local`, `a2_if_authenticated`, `a2_none`, `a2_radius`, `a2_tacacs`, `a3_group`, `a3_local`, `a3_if_authenticated`, `a3_none`, `a3_radius`, `a3_tacacs`, `a4_group`, `a4_local`, `a4_if_authenticated`, `a4_none`, `a4_radius`, `a4_tacacs`, `name`, `group1_cache`, `group1_group`, `group1_radius`, and `group1_tacacs` attributes to `iosxe_aaa_authorization` resource and data source
- Add `vrf`, `local_authentication_type`, `local_authorization`, and `local_auth_max_fail_attempts` attributes to `iosxe_aaa` resrouce and data source 
- Add `icmp_named_msg_type`, `destination_port_equal_2`, `destination_port_equal_3`, `destination_port_equal_4`, `destination_port_equal_5`, `destination_port_equal_6`, `destination_port_equal_7`, `destination_port_equal_8`, `destination_port_equal_9`, `destination_port_equal_10`, `icmp_msg_type`, and `icmp_msg_code` attributes to `iosxe_access_list_extended` resource and data source
- Add `iosxe_access_list_role_based` resource and data source
- Add `filter_lists_cdp` attribute to `iosxe_device_sensor` resource and data source
- Add `match_access_group_name`, `match_ip_dscp`, and `match_ip_precedence` attributes to `iosxe_class_map` resource and data source
- Add `police_target_bitrate_conform_transmit`, `police_target_bitrate_exceed_transmit`, `police_target_bitrate`, `police_target_bitrate_conform_burst_byte`, and `police_target_bitrate_excess_burst_byte` attributes to `iosxe_policy_map` resource and data source
- Add `cdp_enable`, `cdp_tlv_app`, `cdp_tlv_location`, `cdp_tlv_server_location`, `ip_nat_inside`, and `ip_nat_outside` attributes to `iosxe_interface_ethernet` resource and data source
- Add `negotiation_auto` attribute to `iosxe_interface_port_channel` resource and data source
- Add `ip_hosts`, `ip_hosts_vrf`, `name`, `ips`, `vrf`, `hosts`, `subscriber_templating`, `call_home_contact_email`, `call_home_cisco_tac_1_profile_active`, `call_home_cisco_tac_1_destination_transport_method`, `ip_ftp_passive`, `tftp_source_interface_gigabit_ethernet`, `tftp_source_interface_loopback`, `version`, and `multilink_ppp_bundle_name` attributes to `iosxe_system` resource and data source
- Add `role_based_enforcement_logging_interval`, `role_based_enforcement_vlan_lists`, `role_based_permissions_default_acl_name`, `sgt`, `sxp_connection_peers_ipv4`, `sxp_connection_peers_ipv4_vrf`, `sxp_default_password`, `sxp_default_password_type`, `sxp_enable`, `sxp_listener_hold_max_time`, `sxp_listener_hold_min_time`, `sxp_retry_period` and `sxp_speaker_hold_time` attributes to `iosxe_cts` resource and data source
- Add `http_url_cert` attribute to `iosxe_crypto_ikev2` resource and data source

## 0.7.1

- Change type of `transport_input` attribute of `iosxe_line` resource and data source from `String` to `List of String`
- Add `ip_ssh_version_legacy`, `ip_ssh_time_out` and `ip_ssh_authentication_retries` attributes to `iosxe_system` resource and data source
- Add subinterface option to `iosxe_interface_mpls`, `iosxe_interface_ospf`, `iosxe_interface_ospfv3` and `iosxe_interface_pim` resources and data sources

## 0.7.0

- BREAKING CHANGE: Integrate `iosxe_logging_ipv4_host_transport` resource and data source into `iosxe_logging` resource and data source
- BREAKING CHANGE: Integrate `iosxe_logging_ipv6_host_transport` resource and data source into `iosxe_logging` resource and data source
- BREAKING CHANGE: Integrate `iosxe_logging_ipv4_host_vrf_transport` resource and data source into `iosxe_logging` resource and data source
- BREAKING CHANGE: Integrate `iosxe_logging_ipv6_host_vrf_transport` resource and data source into `iosxe_logging` resource and data source
- BREAKING CHANGE: Refactor import logic of resources to use list of attributes instead of YANG paths
- Add support for specifying device name in import identifier
- Add `filter_spec_dhcp_excludes`, `filter_spec_lldp_excludes`, `filter_spec_cdp_includes`, `tlv_name_system_capabilities` attributes to `iosxe_device_sensor` resource and data source
- Extend timeout to wait for device configuration database lock release, [link](https://github.com/CiscoDevNet/terraform-provider-iosxe/issues/201)
- BREAKING CHANGE: Rename `set_isakmp_profile_ikev2_profile_ikev2_profile_case_ikev2_profile` to `set_ikev2_profile` of `iosxe_crypto_ipsec_profile` resource and data source
- BREAKING CHANGE: Rename `set_isakmp_profile_ikev2_profile_isakmp_profile_case_isakmp_profile` to `set_isakmp_profile` of `iosxe_crypto_ipsec_profile` resource and data source
- Add `timezone`, `timezone_offset_hours`, `timezone_offset_minutes` attributes to `iosxe_clock` resource and data source
- Change `update_source_loopback` attribute of `iosxe_bgp_neighbor` resource and data source from `String` to `Number`
- Change `update_source_loopback` attribute of `iosxe_bgp_ipv4_unicast_vrf_neighbor` resource and data source from `String` to `Number`
- Add `enable_traps_aaa_server`, `enable_traps_vdsl2line`, `enable_traps_adslline`, `enable_traps_pki`, `enable_traps_alarm_type`, `enable_traps_casa`, `enable_traps_cnpd`, `enable_traps_dial`, `enable_traps_dlsw`, `enable_traps_ds1`, `enable_traps_dsp_card_status`, `enable_traps_dsp_oper_state`, `enable_traps_entity_sensor`, `enable_traps_entity_state`, `enable_traps_entity_qfp_mem_res_thresh`, `enable_traps_entity_qfp_throughput_notif`, `enable_traps_ether_oam`, `enable_traps_ethernet_cfm_alarm`, `enable_traps_ethernet_cfm_cc_config`, `enable_traps_ethernet_cfm_cc_cross_connect`, `enable_traps_ethernet_cfm_cc_loop`, `enable_traps_ethernet_cfm_cc_mep_down`, `enable_traps_ethernet_cfm_cc_mep_up`, `enable_traps_ethernet_cfm_cc_crosscheck_mep_missing`, `enable_traps_ethernet_cfm_cc_crosscheck_mep_unknown`, `enable_traps_ethernet_cfm_cc_crosscheck_service_up`, `enable_traps_ethernet_evc_create`, `enable_traps_ethernet_evc_delete`, `enable_traps_ethernet_evc_status`, `enable_traps_firewall_serverstatus`, `enable_traps_frame_relay_config_only`, `enable_traps_frame_relay_config_subif_configs`, `enable_traps_frame_relay_subif_count`, `enable_traps_frame_relay_subif_interval`, `enable_traps_frame_relay_config_bundle_mismatch`, `enable_traps_frame_relay_multilink_bundle_mismatch`, `enable_traps_ip_local_pool`, `enable_traps_isdn_call_information`, `enable_traps_isdn_chan_not_avail`, `enable_traps_isdn_ietf`, `enable_traps_isdn_layer2`, `enable_traps_l2tun_session`, `enable_traps_l2tun_tunnel`, `enable_traps_l2tun_pseudowire_status`, `enable_traps_pimstdmib_neighbor_loss`, `enable_traps_pimstdmib_invalid_register`, `enable_traps_pimstdmib_invalid_join_prune`, `enable_traps_pimstdmib_rp_mapping_change`, `enable_traps_pimstdmib_interface_election`, `enable_traps_pfr`, `enable_traps_pppoe`, `enable_traps_resource_policy`, `enable_traps_rsvp`, `enable_traps_vrrp`, `enable_traps_sonet`, `enable_traps_srp`, `enable_traps_voice` attributes to `iosxe_snmp_server` resource and data source
- Add `next_hops_with_track` attribute to `iosxe_static_route` and `iosxe_static_route_vrf` resources and data sources
- Fix idempotency issue with `key` and `pac_key` attributes of `iosxe_radius` resource
- BREAKING CHANGE: Rename `trap_source_x` attributes of `iosxe_ntp` resource and data source to `source_x`

## 0.6.1

- Add `cisp_enable`, `access_session_mac_move_deny`, `diagnostic_bootup_level`, `memory_free_low_watermark_processor`, `archive_path`, `archive_maximum`, `archive_write_memory`, `archive_time_period`, `archive_log_config_logging_enable`, `archive_log_config_logging_size`, `redundancy`, `redundancy_mode`, `transceiver_type_all_monitoring`, `ip_forward_protocol_nd`, `ip_scp_server_enable`, `ip_ssh_version`, `control_plane_service_policy_input`, `pnp_profiles` attributes to `iosxe_system` resource and data source
- Add `ip_ssh_source_interface_x` attributes to `iosxe_system` resource and data source
- Add `ip_domain_lookup_source_interface_x` attributes to `iosxe_system` resource and data source
- Add `recovery_cause_oam_remote_failure`, `recovery_cause_mrp_miscabling` attributes to `iosxe_errdisable` resource and data source
- Add `enable_traps_mvpn`, `enable_traps_lisp`, `enable_traps_mpls`, `enable_traps_mpls_rfc` attributes to `iosxe_snmp_server` resource and data source
- Add `anycast_gateway_mac_auto` attribute to `iosxe_evpn` resource and data source
- Add `mtu`, `ipv6_flow_monitors`, `ip_nbar_protocol_discovery` attributes to `iosxe_interface_ethernet` resource and data source
- Add `auto_cost_reference_bandwidth`, `passive_interface` attributes to `iosxe_ospf` and `iosxe_ospf_vrf` resources and data sources
- Add `match_application_name`, `match_flow_observation_point`, `match_ipv4_version`, `collect_connection_initiator`, `collect_connection_new_connections`, `collect_connection_server_counter_bytes_network_long`, `collect_connection_server_counter_packets_long`, `collect_datalink_mac_source_address_input`, `collect_flow_direction`, `match_ipv6_destination_address`, `match_ipv6_source_address`, `match_ipv6_version`, `match_ipv6_protocol`, `match_connection_client_ipv4_address`, `match_connection_server_ipv4_address`, `match_connection_client_ipv6_address`, `match_connection_server_ipv6_address`, `match_connection_server_transport_port` attributes to `iosxe_flow_record` resource and data source
- Add `export_protocol`, `option_interface_table_timeout`, `option_vrf_table_timeout`, `option_sampler_table`, `option_application_table_timeout`, `option_application_attributes_timeout` attributes to `iosxe_flow_exporter` resource and data source
- Add `cache_timeout_inactive` attribute to `iosxe_flow_monitor` resource and data source
- Add `device_tracking`, `device_tracking_attached_policies` attributes to `iosxe_interface_ethernet` and `iosxe_interface_port_channel` resources and data sources
- Add `ip_tacacs_source_interface_x`, `ip_radius_source_interface_x` attributes to `iosxe_system` resource and data source
- Add `iosxe_spanning_tree` resource and data source
- Add `iosxe_device_sensor` resource and data source
- Add `iosxe_crypto_pki` resource and data source
- Add `boot_system_flash_files`, `boot_system_bootfiles`, `enable_secret` attributes to `iosxe_system` resource and data source
- Add `iosxe_license` resource and data source
- Add `iosxe_lldp` resource and data source

## 0.6.0

- BREAKING_CHANGE: Drop support for IOS-XE 17.9
- Add `esp-192-aes` and `esp-256-aes` options to `iosxe_crypto_ipsec_transform_set` resource and data source
- Make `v3_auth_algorithm` attribute of `iosxe_snmp_server_user` resource and data source optional
- Add `TLSv1.3` option to `ip_http_tls_version` attribute of `iosxe_system` resource and data source
- BREAKING_CHANGE: Rename `snooping_vlans` to `snooping_vlans_legacy` and add `snooping_vlans` attribute to `iosxe_dhcp` resource and data source to support versions >= `17.14`
- BREAKING_CHANGE: Rename `match_route_type_local` to `match_route_type_local_legacy` and add `match_route_type_local` attribute to `iosxe_route_map` resource and data source to support versions >= `17.15`
- BREAKING_CHANGE: Remove `delete_mode` attribute from `iosxe_snmp_server_user` resource
- BREAKING_CHANGE: Rename `inspection_filters.vlan` attribute of `iosxe_arp` resource and data source to `inspection_filters.vlans`
- BREAKING CHANGE: Rename `vlan_based_auto_route_target` to `vlan_based_auto_route_target_legacy` and add `vlan_based_auto_route_target` attribute to `iosxe_evpn_instance` resource and data source to support versions >= `17.15`
- BREAKING CHANGE: Rename `vlan_based_route_target_import` to `vlan_based_route_target_import_legacy` and add `vlan_based_route_target_import` attribute to `iosxe_evpn_instance` resource and data source to support versions >= `17.15`
- BREAKING CHANGE: Rename `vlan_based_route_target` attribute to `vlan_based_route_target_legacy` of `iosxe_evpn_instance` resource and data source
- BREAKING CHANGE: Rename `vlan_based_route_target_both` attribute to `vlan_based_route_target_both_legacy` of `iosxe_evpn_instance` resource and data source
- BREAKING CHANGE: Rename `vlan_based_route_target_export` attribute to `vlan_based_route_target_export_legacy` of `iosxe_evpn_instance` resource and data source
- BREAKING CHANGE: Rename `vlan_based_route_target_import` attribute to `vlan_based_route_target_import_legacy` of `iosxe_evpn_instance` resource and data source
- Add `vlan_based_route_target_exports` attribute to `iosxe_evpn_instance` resource and data source to support versions >= `17.15`
- Add `vlan_based_route_target_imports` attribute to `iosxe_evpn_instance` resource and data source to support versions >= `17.15`
- Add `vtp_mode_client`, `vtp_mode_off`, `vtp_mode_server`, `vtp_mode_transparent` attributes to `iosxe_vtp` resource and data source
- Add `bpduguard_enable`, `bpduguard_disable`, `spanning_tree_portfast`, `spanning_tree_portfast_disable` and `spanning_tree_portfast_edge` attributes to `iosxe_interface_ethernet` resource and data source
- BREAKING CHANGE: Rename `trap_source_hundred_gig_e` attribute to `trap_source_hundred_gigabit_ethernet` of `iosxe_ntp` resource and data source
- Add `ip_name_servers` and `ip_name_servers_vrf` attributes to `iosxe_system` resource and data source

## 0.5.10

- Fix update issue with `iosxe_restconf` resource and empty `attributes` map
- Fix issue with `iosxe_restconf` validation, [link](https://github.com/CiscoDevNet/terraform-provider-iosxe/issues/223)
- Fix issue with `iosxe_interface_ethernet` resource blocking the device database, [link](https://github.com/CiscoDevNet/terraform-provider-iosxe/issues/201)

## 0.5.9

- Fix idempotency issue with `secret` attribute of `iosxe_username` resource
- Add `tunnel_vrf` attribute to `iosxe_interface_tunnel` resource and data source
- Add `ip_http_secure_active_session_modules`, `ip_http_max_connections` and `ip_http_active_session_modules` attributes to `iosxe_system` resource and data source
- Fix idempotency issue with `key` attribute of `iosxe_tacacs_server` resource
- Add `encryption` attribute to `iosxe_tacacs_server` resource and data source
- Add `trusted_keys` attribute to `iosxe_ntp` resource and data source

## 0.5.8

- Fix import of resources
- Add `managed` flag to provider device configuration to allow temporarily skipping a device due to maintenance

## 0.5.7

- Add `Tunnel` interface type to `iosxe_interface_pim` resource and data source
- Add `ip_mtu` attribute to `iosxe_interface_tunnel` resources and data sources

## 0.5.6

- Add support for descriptions to `iosxe_prefix_list` resource and data source
- Add `advertisement_interval` attribute to `iosxe_bgp_ipv4_unicast_vrf_neighbor` resource and data source
- Add `iosxe_flow_record` resource and data source
- Add `iosxe_flow_exporter` resource and data source
- Add `iosxe_flow_monitor` resource and data source
- Add `ip_flow_monitors` attribute to `iosxe_interface_ethernet` resource and data source
- Add `negotiation_auto` attribute to `iosxe_interface_ethernet` resource and data source
- BREAKING CHANGE: Rename `ianctivity_timer` attribute to `inactivity_timer` of `iosxe_service_template` resource and data source
- Add `Tunnel` option to `iosxe_interface_ospf` and `iosxe_interface_ospfv3` resource and data source
- Add `logging_event_link_status_enable` and `snmp_trap_link_status` attributes to `iosxe_interface_ethernet`, `iosxe_interface_port_channel`, `iosxe_interface_tunnel` resources and data sources
- Add `load_interval` attribute to `iosxe_interface_ethernet`, `iosxe_interface_port_channel`, `iosxe_interface_tunnel`, `iosxe_interface_vlan` resources and data sources

## 0.5.5

- Add `bandwidth` attribute to `iosxe_interface_ethernet` resource and data source
- Add `speed_nonegotiate` attribute to `iosxe_interface_ethernet` resource and data source
- Add `service_policy_input` and `service_policy_output` attributes to `iosxe_interface_ethernet` resource and data source
- Add `classes` and `description` attributes to `iosxe_policy_map` resource and data source
- Add `match_dscp` attribute to `iosxe_class_map` resource and data source

## 0.5.4

- Add `next_hop_self` and `next_hop_self_all` attributes to `iosxe_bgp_ipv4_unicast_vrf_neighbor` resource and data source
- Add `set_as_path_replace_any` and `set_as_path_replace_as` attributes to `iosxe_route_map` resource and data source
- Add `ip_http` attributes to `iosxe_system` resource and data source

## 0.5.3

- Add `ipv4_unicast_router_id_loopback` attribute to `iosxe_bgp_address_family_ipv4_vrf` resource and data source
- Add `ha_mode_graceful_restart` attribute to `iosxe_bgp_ipv4_unicast_vrf_neighbor` resource and data source
- Add `dot1x` attribute to `iosxe_aaa_authentication` resource and data source
- Add `identities` attribute to `iosxe_aaa_accounting` resource and data source

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

