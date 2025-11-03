resource "iosxe_bgp_neighbor" "example" {
  asn                                       = "65000"
  ip                                        = "3.3.3.3"
  remote_as                                 = "65000"
  description                               = "BGP Neighbor Description"
  shutdown                                  = false
  cluster_id                                = "1234"
  version                                   = 4
  disable_connected_check                   = false
  fall_over_default_enable                  = false
  fall_over_default_route_map               = "RMAP"
  fall_over_bfd                             = true
  fall_over_bfd_single_hop                  = true
  fall_over_bfd_check_control_plane_failure = true
  fall_over_bfd_strict_mode                 = true
  fall_over_maximum_metric_route_map        = "RMAP"
  log_neighbor_changes                      = true
  password_type                             = 1
  password                                  = "test1234"
  timers_keepalive_interval                 = 655
  timers_holdtime                           = 866
  timers_minimum_neighbor_hold              = 222
  update_source_interface_loopback          = 100
}
