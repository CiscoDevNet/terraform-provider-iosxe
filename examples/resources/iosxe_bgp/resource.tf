resource "iosxe_bgp" "example" {
  asn                  = "65000"
  default_ipv4_unicast = false
  log_neighbor_changes = true
  router_id_ip         = "172.16.255.1"
  bgp_graceful_restart = true
  bgp_update_delay     = 200
}
