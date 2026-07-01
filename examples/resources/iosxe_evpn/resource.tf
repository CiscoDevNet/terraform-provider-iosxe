resource "iosxe_evpn" "example" {
  replication_type_ingress                             = false
  replication_type_static                              = true
  replication_type_p2mp                                = false
  replication_type_mp2mp                               = false
  evpn_mac_duplication_limit                           = 10
  evpn_mac_duplication_time                            = 100
  evpn_ip_duplication_limit                            = 10
  evpn_ip_duplication_time                             = 100
  router_id_loopback                                   = 100
  evpn_default_gateway_advertise                       = true
  evpn_logging_peer_state                              = true
  evpn_route_target_auto_vni                           = true
  evpn_flooding_suppression_address_resolution_disable = true
  evpn_multicast_advertise                             = true
}
