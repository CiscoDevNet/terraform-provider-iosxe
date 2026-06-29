resource "iosxe_monitor_session" "example" {
  session_id = 1
  destination_interface = [
    {
      name          = "GigabitEthernet1/0/1"
      encapsulation = "replicate"
    }
  ]
  source_interface = [
    {
      name      = "GigabitEthernet1/0/2"
      direction = "both"
    }
  ]
}
