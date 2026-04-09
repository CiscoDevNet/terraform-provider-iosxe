resource "iosxe_nat" "example" {
  inside_source_interfaces = [
    {
      id = "10"
      interfaces = [
        {
          interface = "GigabitEthernet4"
          overload  = true
        }
      ]
    }
  ]
  inside_source_static_entries = [
    {
      local_ip  = "10.0.0.1"
      global_ip = "203.0.113.1"
    }
  ]
  outside_source_static_entries = [
    {
      global_ip = "198.51.100.1"
      local_ip  = "10.0.0.2"
    }
  ]
}
