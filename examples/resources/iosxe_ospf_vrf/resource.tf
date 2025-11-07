resource "iosxe_ospf_vrf" "example" {
  process_id                                    = 2
  vrf                                           = "VRF1"
  bfd_all_interfaces                            = true
  default_information_originate                 = true
  default_information_originate_always          = true
  default_metric                                = 21
  distance                                      = 120
  domain_tag                                    = 10
  log_adjacency_changes                         = true
  log_adjacency_changes_detail                  = true
  nsf_cisco                                     = true
  nsf_cisco_enforce_global                      = true
  nsf_ietf                                      = true
  nsf_ietf_restart_interval                     = 120
  max_metric_router_lsa                         = true
  max_metric_router_lsa_summary_lsa_metric      = 16711680
  max_metric_router_lsa_external_lsa_metric     = 16711680
  max_metric_router_lsa_include_stub            = true
  max_metric_router_lsa_on_startup_time         = 60
  max_metric_router_lsa_on_startup_wait_for_bgp = true
  redistribute_static_subnets                   = true
  redistribute_connected_subnets                = true
  neighbor = [
    {
      ip       = "2.2.2.2"
      priority = 10
      cost     = 100
    }
  ]
  network = [
    {
      ip       = "3.3.3.0"
      wildcard = "0.0.0.255"
      area     = "0"
    }
  ]
  priority  = 100
  router_id = "1.2.3.4"
  shutdown  = false
  summary_address = [
    {
      ip   = "3.3.3.0"
      mask = "255.255.255.0"
    }
  ]
  areas = [
    {
      area_id                                        = "5"
      authentication_message_digest                  = true
      nssa                                           = true
      nssa_default_information_originate             = true
      nssa_default_information_originate_metric      = 100
      nssa_default_information_originate_metric_type = 1
      nssa_no_summary                                = true
      nssa_no_redistribution                         = true
    }
  ]
  auto_cost_reference_bandwidth = 40000
  passive_interface_default     = true
}
