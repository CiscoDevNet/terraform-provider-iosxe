resource "iosxe_bgp" "example" {
  asn                  = "65000"
  default_ipv4_unicast = false
  log_neighbor_changes = true
  router_id_loopback   = 100
  router_id_ip         = "172.16.255.1"
}
