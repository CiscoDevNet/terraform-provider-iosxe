resource "iosxe_interface_vrrp_v2" "example" {
  type               = "GigabitEthernet"
  name               = "1"
  group_id           = 1
  ip_primary_address = "10.0.0.254"
  ip_secondary_addresses = [
    {
      address = "10.0.0.253"
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
