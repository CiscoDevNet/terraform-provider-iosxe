resource "iosxe_flow_exporter" "example" {
  name                  = "EXPORTER1"
  description           = "My exporter"
  destination_ip        = "1.1.1.1"
  source_loopback       = 123
  transport_udp         = 655
  ttl                   = 64
  template_data_timeout = 60
}
