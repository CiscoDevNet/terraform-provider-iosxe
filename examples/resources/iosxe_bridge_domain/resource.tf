resource "iosxe_bridge_domain" "example" {
  bridge_domain_id = 100
  member_vni       = 10100
  member_interface = [
    {
      interface = "GigabitEthernet2"
      service_instance = [
        {
          instance_id = 100
        }
      ]
    }
  ]
}
