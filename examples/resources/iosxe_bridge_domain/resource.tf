resource "iosxe_bridge_domain" "example" {
  bridge_domain_id = 836
  member_vni       = 10800
  member_interfaces = [
    {
      interface = "GigabitEthernet2"
      service_instances = [
        {
          instance_id = 836
        }
      ]
    }
  ]
}
