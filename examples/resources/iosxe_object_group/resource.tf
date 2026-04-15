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
}
