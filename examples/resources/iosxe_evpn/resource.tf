resource "iosxe_evpn" "example" {
  replication_type_ingress  = false
  replication_type_static   = true
  replication_type_p2mp     = false
  replication_type_mp2mp    = false
  mac_duplication_limit     = 10
  mac_duplication_time      = 100
  ip_duplication_limit      = 10
  ip_duplication_time       = 100
  router_id_loopback        = 100
  default_gateway_advertise = true
  logging_peer_state        = true
  route_target_auto_vni     = true
}
