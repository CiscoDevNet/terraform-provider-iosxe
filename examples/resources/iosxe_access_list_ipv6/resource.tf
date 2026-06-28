resource "iosxe_access_list_ipv6" "example" {
  name = "V6ACL1"
  entries = [
    {
      sequence          = 10
      ace_rule_action   = "permit"
      ace_rule_protocol = "tcp"
      source_any        = true
      destination_any   = true
    }
  ]
}
