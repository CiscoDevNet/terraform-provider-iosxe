resource "iosxe_system" "example" {
  hostname                    = "ROUTER-1"
  ip_bgp_community_new_format = true
  ipv6_unicast_routing        = true
  ip_source_route             = false
  ip_domain_lookup            = false
  ip_domain_name              = "test.com"
  login_delay                 = 10
  login_on_failure            = true
  login_on_failure_log        = true
  login_on_success            = true
  login_on_success_log        = true
  multicast_routing_vrfs = [
    {
      vrf = "VRF1"
    }
  ]
  ip_name_servers = ["1.2.3.4"]
  ip_name_servers_vrf = [
    {
      vrf     = "VRF1"
      servers = ["2.3.4.5"]
    }
  ]
  ip_domain_lookup_nsap      = true
  ip_domain_lookup_recursive = true
  ip_domain_lookup_vrfs = [
    {
      vrf                               = "VRF1"
      source_interface_gigabit_ethernet = "1/0/1"
    }
  ]
  diagnostic_bootup_level                                   = "minimal"
  memory_free_low_watermark_processor                       = 203038
  ip_ssh_time_out                                           = 120
  ip_ssh_authentication_retries                             = 3
  ip_ssh_bulk_mode                                          = true
  ip_ssh_bulk_mode_window_size                              = 262144
  call_home_contact_email                                   = "email@test.com"
  call_home_cisco_tac_1_profile_active                      = true
  call_home_cisco_tac_1_destination_transport_method        = "email"
  ip_nbar_classification_dns_classify_by_domain             = true
  ip_multicast_route_limit                                  = 200000
  ip_domain_list_vrf_domain                                 = "example.com"
  ip_domain_list_vrf                                        = "VRF1"
  ip_routing_protocol_purge_interface                       = true
  ip_cef_load_sharing_algorithm_include_ports_source        = true
  ip_cef_load_sharing_algorithm_include_ports_destination   = true
  ipv6_cef_load_sharing_algorithm_include_ports_source      = true
  ipv6_cef_load_sharing_algorithm_include_ports_destination = true
  port_channel_load_balance                                 = "src-dst-mixed-ip-port"
}
