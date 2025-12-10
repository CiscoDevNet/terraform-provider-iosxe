resource "iosxe_interface_isis" "example" {
  type = "Loopback"
  name = "100"
  ipv4_metric_levels = [
    {
      level = "level-1"
      value = 100
    }
  ]
}
