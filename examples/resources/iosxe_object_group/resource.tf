resource "iosxe_object_group" "example" {
  fqdn = [
    {
      name        = "FQDN_GROUP_1"
      description = "My FQDN object group"
      patterns = [
        {
          fqdn_pattern = "test-fqdn-pattern"
        }
      ]
    }
  ]
  network = [
    {
      name        = "NETWORK_GROUP_1"
      description = "My network object group"
      hosts = [
        {
          ipv4_host = "10.1.1.1"
        }
      ]
      network_addresses = [
        {
          ipv4_address = "10.1.2.0"
          ipv4_mask    = "255.255.255.0"
        }
      ]
      address_ranges = [
        {
          start = "10.1.3.1"
          end   = "10.1.3.10"
        }
      ]
    }
  ]
}
