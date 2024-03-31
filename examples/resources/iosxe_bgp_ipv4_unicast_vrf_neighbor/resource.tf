resource "iosxe_bgp_ipv4_unicast_vrf_neighbor" "example" {
  asn                                       = "65000"
  vrf                                       = "VRF1"
  ip                                        = "3.3.3.3"
  remote_as                                 = "65000"
  description                               = "BGP Neighbor Description"
  shutdown                                  = false
  log_neighbor_changes_disable              = true
  password_type                             = 1
  password                                  = "LINE"
  timers_keepalive_interval                 = 30
  timers_holdtime                           = 40
  timers_minimum_neighbor_hold              = 30
  version                                   = 4
  fall_over_default_route_map               = "RMAP"
  fall_over_bfd                             = true
  fall_over_bfd_single_hop                  = true
  fall_over_bfd_check_control_plane_failure = true
  fall_over_bfd_strict_mode                 = true
  fall_over_maximum_metric_route_map        = "ROUTEMAP"
  update_source_loopback                    = "100"
  activate                                  = true
  send_community                            = "both"
  route_reflector_client                    = false
  soft_reconfiguration                      = "inbound"
  default_originate                         = true
  default_originate_route_map               = "RM1"
  route_maps = [
    {
      in_out         = "in"
      route_map_name = "RM1"
    }
  ]
  ha_mode_graceful_restart = true
  next_hop_self            = true
  next_hop_self_all        = true
  advertisement_interval   = 300
}
