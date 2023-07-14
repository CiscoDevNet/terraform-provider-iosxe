resource "iosxe_bgp_neighbor" "example" {
  asn                    = "65000"
  ip                     = "3.3.3.3"
  remote_as              = "65000"
  description            = "BGP Neighbor Description"
  shutdown               = false
  update_source_loopback = "100"
}
