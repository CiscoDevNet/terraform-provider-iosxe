resource "iosxe_snmp_server_user" "example" {
  username                         = "USER1"
  grpname                          = "GROUP1"
  v3_auth_algorithm                = "sha"
  v3_auth_password                 = "Cisco123"
  v3_auth_priv_aes_algorithm       = "128"
  v3_auth_priv_aes_password        = "Cisco123"
  v3_auth_priv_aes_access_ipv6_acl = "V6ACL1"
  v3_auth_priv_aes_access_acl_name = "ACL123"
}
