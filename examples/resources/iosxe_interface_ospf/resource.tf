resource "iosxe_interface_ospf" "example" {
  type                             = "Loopback"
  name                             = "1"
  cost                             = 10
  dead_interval                    = 30
  hello_interval                   = 5
  mtu_ignore                       = false
  network_type_broadcast           = false
  network_type_non_broadcast       = false
  network_type_point_to_multipoint = false
  network_type_point_to_point      = true
  priority                         = 10
  ttl_security_hops                = 2
  process_ids = [
    {
      id = 1
      areas = [
        {
          area_id = "0"
        }
      ]
    }
  ]
  message_digest_keys = [
    {
      id            = 1
      md5_auth_key  = "mykey"
      md5_auth_type = 0
    }
  ]
}
