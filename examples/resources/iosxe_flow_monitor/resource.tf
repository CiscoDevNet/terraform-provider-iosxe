resource "iosxe_flow_monitor" "example" {
  name        = "MON1"
  description = "My monitor"
  flow_monitor_exporter = [
    {
      name = "EXPORTER1"
    }
  ]
  cache_timeout_active = 60
  record               = "FNF1"
}
