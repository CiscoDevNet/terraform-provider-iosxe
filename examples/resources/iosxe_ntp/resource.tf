resource "iosxe_ntp" "example" {
  authenticate                = true
  logging                     = false
  access_group_peer_acl       = "SACL1"
  access_group_query_only_acl = "SACL1"
  access_group_serve_acl      = "SACL1"
  access_group_serve_only_acl = "SACL1"
  authentication_keys = [
    {
      number          = 1
      md5             = "060506324F41584B564347"
      encryption_type = 7
    }
  ]
  master               = true
  master_stratum       = 5
  passive              = false
  update_calendar      = false
  trap_source_loopback = 1
  servers = [
    {
      ip_address = "1.2.3.4"
      source     = "Loopback1"
      key        = 1
      prefer     = true
      version    = 2
    }
  ]
  server_vrfs = [
    {
      name = "VRF1"
      servers = [
        {
          ip_address = "3.4.5.6"
          key        = 1
          prefer     = true
          version    = 2
        }
      ]
    }
  ]
  peers = [
    {
      ip_address = "5.2.3.4"
      source     = "Loopback1"
      key        = 1
      prefer     = true
      version    = 2
    }
  ]
  peer_vrfs = [
    {
      name = "VRF1"
      peers = [
        {
          ip_address = "5.4.5.6"
          key        = 1
          prefer     = true
          version    = 2
        }
      ]
    }
  ]
}
