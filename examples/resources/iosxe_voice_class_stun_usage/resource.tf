resource "iosxe_voice_class_stun_usage" "example" {
  tag                                    = 1
  stun_usage_firewall_traversal_flowdata = true
  stun_usage_ice_lite                    = true
}
