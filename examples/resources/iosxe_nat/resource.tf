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
}
