resource "iosxe_nat" "example" {
  inside_source_interface_lists = [
    {
      id = "10"
      interfaces = [
        {
          interface = "GigabitEthernet10"
          overload  = true
        }
      ]
    }
  ]
}
