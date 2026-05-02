resource "iosxe_device_tracking_policy" "example" {
  name                                       = "DT_trunk_policy"
  trusted_port                               = true
  data_glean_recovery_dhcp                   = true
  data_glean_recovery_ndp                    = true
  prefix_glean                               = true
  destination_glean_recovery_dhcp            = true
  protocol_arp                               = false
  protocol_dhcp4                             = false
  protocol_dhcp6                             = false
  protocol_ndp                               = false
  tracking_enable                            = true
  tracking_enable_reachable_lifetime_seconds = 300
  limit_address_count                        = 100
  security_level_glean                       = true
}
