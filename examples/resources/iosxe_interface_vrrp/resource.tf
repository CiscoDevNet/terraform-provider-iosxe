resource "iosxe_interface_vrrp" "example" {
  type                 = "Vlan"
  name                 = "123"
  group_id             = 1
  ipv4_primary_address = "192.0.2.254"
  ipv4_secondary_addresses = [
    {
      address   = "192.0.2.253"
      secondary = true
    }
  ]
  ipv4_priority              = 110
  ipv4_preempt_delay_minimum = 30
  ipv4_timers_advertise      = 3000
  ipv4_description           = "VRRP-GROUP-1"
  ipv4_tracks = [
    {
      object_id = "1"
      decrement = 20
      shutdown  = false
    }
  ]
  ipv4_shutdown        = false
  ipv6_primary_address = "fe80::1"
  ipv6_prefixes = [
    {
      prefix = "2001:DB8::FFFF/64"
    }
  ]
  ipv6_priority              = 120
  ipv6_preempt_delay_minimum = 30
  ipv6_timers_advertise      = 3000
  ipv6_description           = "VRRP-GROUP-1-IPV6"
  ipv6_tracks = [
    {
      object_id = "1"
      decrement = 20
      shutdown  = false
    }
  ]
  ipv6_shutdown = false
}
