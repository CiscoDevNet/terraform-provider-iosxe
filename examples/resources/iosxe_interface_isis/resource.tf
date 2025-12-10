resource "iosxe_interface_isis" "example" {
  type                   = "Loopback"
  name                   = "100"
  network_point_to_point = true
  ipv4_metric_levels = [
    {
      level = "level-1"
      value = 100
    }
  ]
}
