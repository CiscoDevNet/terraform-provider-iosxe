resource "iosxe_interface_hsrp" "example" {
  type                    = "Vlan"
  name                    = "123"
  version                 = "2"
  bfd                     = true
  delay_minimum           = 100
  delay_reload            = 200
  mac_refresh             = 30
  use_bia                 = true
  use_bia_scope_interface = false
  standby_list = [
    {
      group_number        = 10
      authentication_text = "MySecret"
      ip                  = true
      ip_address          = "192.0.2.254"
      ip_secondary_addresses = [
        {
          address   = "192.0.2.253"
          secondary = true
        }
      ]
      preempt                       = true
      preempt_delay_minimum         = 30
      preempt_delay_reload          = 60
      preempt_delay_sync            = 120
      priority                      = 110
      timers_hello_interval_seconds = 3
      timers_hold_time_seconds      = 10
      tracks = [
        {
          number    = 1
          decrement = 20
          shutdown  = false
        }
      ]
    }
  ]
}
