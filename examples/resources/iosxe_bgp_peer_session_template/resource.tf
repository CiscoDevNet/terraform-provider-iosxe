resource "iosxe_bgp_peer_session_template" "example" {
  asn                              = "65000"
  template_name                    = "PEER_SESSION_TEMPLATE_1"
  remote_as                        = "65001"
  description                      = "Peer Session Template Description"
  disable_connected_check          = true
  ebgp_multihop                    = true
  ebgp_multihop_max_hop            = 10
  update_source_interface_loopback = 100
}
