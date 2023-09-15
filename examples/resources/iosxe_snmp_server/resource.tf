resource "iosxe_snmp_server" "example" {
  chassis_id                                = "R1"
  contact                                   = "Contact1"
  ifindex_persist                           = true
  location                                  = "Location1"
  packetsize                                = 2000
  queue_length                              = 100
  enable_logging_getop                      = true
  enable_logging_setop                      = true
  enable_traps                              = true
  enable_traps_snmp_authentication          = true
  enable_traps_snmp_coldstart               = true
  enable_traps_snmp_linkdown                = true
  enable_traps_snmp_linkup                  = true
  enable_traps_snmp_warmstart               = true
  system_shutdown                           = true
  enable_traps_flowmon                      = true
  enable_traps_entity_perf_throughput_notif = true
  enable_traps_call_home_message_send_fail  = true
  enable_traps_call_home_server_fail        = true
  enable_traps_tty                          = true
  enable_traps_ospfv3_config_state_change   = true
  enable_traps_ospfv3_config_errors         = true
  enable_traps_ospf_config_retransmit       = true
  enable_traps_ospf_config_lsa              = true
  enable_traps_ospf_nssa_trans_change       = true
  enable_traps_ospf_shamlink_interface      = true
  enable_traps_ospf_shamlink_neighbor       = true
  enable_traps_ospf_errors_enable           = true
  enable_traps_ospf_retransmit_enable       = true
  enable_traps_ospf_lsa_enable              = true
  enable_traps_eigrp                        = true
  enable_traps_auth_framework_sec_violation = true
  enable_traps_vtp                          = true
  enable_traps_vlancreate                   = true
  enable_traps_vlandelete                   = true
  enable_traps_port_security                = true
  enable_traps_license                      = true
  enable_traps_smart_license                = true
  enable_traps_cpu_threshold                = true
  enable_traps_memory_bufferpeak            = true
  enable_traps_fru_ctrl                     = true
  enable_traps_flash_insertion              = true
  enable_traps_flash_removal                = true
  enable_traps_flash_lowspace               = true
  enable_traps_entity                       = true
  enable_traps_pw_vc                        = true
  enable_traps_ipsla                        = true
  enable_traps_bfd                          = true
  enable_traps_ike_policy_add               = true
  enable_traps_ike_policy_delete            = true
  enable_traps_ike_tunnel_start             = true
  enable_traps_ike_tunnel_stop              = true
  enable_traps_ipsec_cryptomap_add          = true
  enable_traps_ipsec_cryptomap_attach       = true
  enable_traps_ipsec_cryptomap_delete       = true
  enable_traps_ipsec_cryptomap_detach       = true
  enable_traps_ipsec_tunnel_start           = true
  enable_traps_ipsec_tunnel_stop            = true
  enable_traps_ipsec_too_many_sas           = true
  enable_traps_config_copy                  = true
  enable_traps_config                       = true
  enable_traps_config_ctid                  = true
  enable_traps_dhcp                         = true
  enable_traps_event_manager                = true
  enable_traps_ipmulticast                  = true
  enable_traps_msdp                         = true
  enable_traps_ospf_config_state_change     = true
  enable_traps_ospf_config_errors           = true
  enable_traps_pim_invalid_pim_message      = true
  enable_traps_pim_neighbor_change          = true
  enable_traps_pim_rp_mapping_change        = true
  enable_traps_syslog                       = true
  enable_traps_rf                           = true
  enable_traps_transceiver_all              = true
  enable_traps_bulkstat_collection          = true
  enable_traps_bulkstat_transfer            = true
  enable_traps_vrfmib_vrf_up                = true
  enable_traps_vrfmib_vrf_down              = true
  enable_traps_vrfmib_vnet_trunk_up         = true
  enable_traps_vrfmib_vnet_trunk_down       = true
  source_interface_informs_loopback         = 1
  source_interface_traps_loopback           = 1
  trap_source_loopback                      = 1
  snmp_communities = [
    {
      name             = "COM1"
      view             = "VIEW1"
      permission       = "ro"
      ipv6             = "ACL1"
      access_list_name = "1"
    }
  ]
  contexts = [
    {
      name = "CON1"
    }
  ]
  views = [
    {
      name    = "VIEW1"
      mib     = "interfaces"
      inc_exl = "included"
    }
  ]
}
