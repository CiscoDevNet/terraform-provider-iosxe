resource "iosxe_bgp_bmp_server" "example" {
  asn                    = "65000"
  server_id              = 1
  address                = "10.10.10.10"
  port_number            = 1790
  activate               = true
  failure_retry_delay    = 30
  flapping_delay         = 60
  initial_delay          = 10
  stats_reporting_period = 30
  update_source_loopback = 100
}
