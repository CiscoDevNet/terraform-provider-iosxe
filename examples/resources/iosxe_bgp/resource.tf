resource "iosxe_bgp" "example" {
  asn                  = "65000"
  default_ipv4_unicast = false
  log_neighbor_changes = true
  bgp_graceful_restart = true
  bgp_update_delay     = 200
}
