resource "iosxe_device_tracking" "example" {
  logging_theft                          = true
  tracking_auto_source_fallback_ipv4     = "10.0.0.1"
  tracking_auto_source_fallback_mask     = "255.255.255.0"
  tracking_auto_source_fallback_override = true
  tracking_retry_interval                = 10
  binding_reachable_lifetime             = 7200
  policies = [
    {
      name                            = "DT_trunk_policy"
      trusted_port                    = true
      data_glean_recovery_dhcp        = true
      data_glean_recovery_ndp         = true
      prefix_glean                    = true
      destination_glean_recovery_dhcp = true
      protocol_arp                    = false
      protocol_dhcp4                  = false
      protocol_dhcp6                  = false
      protocol_ndp                    = false
      tracking_enable                 = true
      limit_address_count             = 100
      security_level_glean            = true
    }
  ]
}
