resource "iosxe_vrrp" "example" {
  type                    = "Vlan"
  name                    = "123"
  group_id                = 1
  address_primary_address = "192.0.2.254"
  address_primary         = true
  secondary_addresses = [
    {
      address   = "192.0.2.253"
      secondary = true
    }
  ]
  priority              = 110
  preempt_delay_minimum = 30
  timers_advertise      = 3000
  description           = "VRRP-GROUP-1"
  shutdown              = false
}
