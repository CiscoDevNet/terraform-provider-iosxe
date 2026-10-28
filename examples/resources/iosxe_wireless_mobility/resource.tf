resource "iosxe_wireless_mobility" "example" {
  group_name               = "MOBILITY-1"
  mac_address              = "00:00:00:00:0c:01"
  group_keepalive_interval = 20
  group_keepalive_count    = 5
  dscp                     = 46
  group_members = [
    {
      mac_address          = "00:00:00:00:0a:20"
      ip                   = "192.0.2.20"
      public_ip            = "198.51.100.20"
      group                = "MOBILITY-2"
      ssc_hash             = "0123456789abcdef0123456789abcdef01234567"
      data_link_encryption = true
    }
  ]
}
