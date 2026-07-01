resource "iosxe_interface_vrrp_v2" "example" {
  type               = "GigabitEthernet"
  name               = "2"
  group_id           = 1
  ip_primary_address = "192.0.2.254"
  ip_secondary_addresses = [
    {
      address   = "192.0.2.253"
      secondary = true
    }
  ]
  priority                  = 110
  preempt                   = true
  preempt_delay_minimum     = 30
  timers_advertise_interval = 3
  authentication_text       = "SECRET"
  description               = "VRRP-GROUP-1"
  tracks = [
    {
      object_id = 1
      decrement = 20
    }
  ]
  shutdown = false
}
