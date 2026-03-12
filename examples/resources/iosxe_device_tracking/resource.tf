resource "iosxe_device_tracking" "example" {
  logging_theft                          = true
  tracking_auto_source_fallback_ipv4     = "10.0.0.1"
  tracking_auto_source_fallback_mask     = "255.255.255.0"
  tracking_auto_source_fallback_override = true
  tracking_retry_interval                = 10
}
