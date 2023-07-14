resource "iosxe_snmp_server_group" "example" {
  name = "GROUP1"
  v3_security = [
    {
      security_level  = "priv"
      context_node    = "CON1"
      match_node      = "exact"
      read_node       = "VIEW1"
      write_node      = "VIEW2"
      notify_node     = "VIEW3"
      access_ipv6_acl = "V6ACL1"
      access_acl_name = "ACL1"
    }
  ]
}
