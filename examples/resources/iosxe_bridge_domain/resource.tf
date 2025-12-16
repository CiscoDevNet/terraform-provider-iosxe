resource "iosxe_bridge_domain" "example" {
  bridge_domain_id = 100
  member_vni       = 10100
  member_interfaces = [
    {
      interface = "GigabitEthernet2"
      service_instances = [
        {
          instance_id = 100
        }
      ]
    }
  ]
}
