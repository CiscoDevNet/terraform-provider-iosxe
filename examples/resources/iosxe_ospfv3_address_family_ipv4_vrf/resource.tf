resource "iosxe_ospfv3_address_family_ipv4_vrf" "example" {
  process_id                                           = 2
  vrf                                                  = "VRF1"
  unicast                                              = true
  capability_vrf                                       = "vrf-lite"
  bfd_all_interfaces                                   = true
  default_information_originate                        = true
  default_information_originate_always                 = true
  default_information_originate_metric                 = 1000
  distance                                             = 120
  log_adjacency_changes                                = true
  log_adjacency_changes_detail                         = true
  router_id                                            = "1.2.3.4"
  shutdown                                             = false
  auto_cost_reference_bandwidth                        = 40000
  timers_lsa_arrival                                   = 15
  timers_pacing_lsa_group                              = 15
  timers_throttle_lsa_all_delay                        = 0
  timers_throttle_lsa_all_min_delay                    = 50
  timers_throttle_lsa_all_max_delay                    = 5000
  timers_throttle_spf_delay                            = 10
  timers_throttle_spf_min_delay                        = 50
  timers_throttle_spf_max_delay                        = 5000
  max_metric_router_lsa_config                         = true
  max_metric_router_lsa_config_stub_prefix_lsa         = true
  max_metric_router_lsa_config_inter_area_lsas_metric  = 16711680
  max_metric_router_lsa_config_external_lsa_metric     = 16711680
  max_metric_router_config_lsa_on_startup_wait_for_bgp = true
  redistribute_static                                  = true
  redistribute_connected                               = true
  summary_prefix = [
    {
      prefix = "3.3.3.0/24"
    }
  ]
  areas = [
    {
      area_id                                        = "5"
      nssa                                           = true
      nssa_default_information_originate             = true
      nssa_default_information_originate_metric      = 100
      nssa_default_information_originate_metric_type = 1
      nssa_no_summary                                = true
      nssa_no_redistribution                         = true
    }
  ]
  passive_interface_default = true
}
