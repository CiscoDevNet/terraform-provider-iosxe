resource "iosxe_vrrp_ipv6" "example" {
  type            = "Vlan"
  name            = "123"
  group_id        = 2
  ipv6_link_local = "FE80::1"
  ipv6_primary    = true
  ipv6_prefixes = [
    {
      prefix = "2001:DB8::FFFF/64"
    }
  ]
  priority              = 110
  preempt_delay_minimum = 30
  timers_advertise      = 3000
  description           = "VRRP-IPV6-GROUP-2"
  tracks = [
    {
      object_id = "1"
      decrement = 20
      shutdown  = false
    }
  ]
  shutdown = false
}
