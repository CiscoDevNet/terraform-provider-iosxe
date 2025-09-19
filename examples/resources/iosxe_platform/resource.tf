resource "iosxe_platform" "example" {
  punt_keepalive_disable_kernel_core        = true
  punt_keepalive_settings_fatal_count       = 15
  punt_keepalive_settings_transmit_interval = 15
  punt_keepalive_settings_warning_count     = 15
}
