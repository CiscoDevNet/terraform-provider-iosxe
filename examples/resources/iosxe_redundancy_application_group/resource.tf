resource "iosxe_redundancy_application_group" "example" {
  group_id           = 1
  name               = "cube-rg-1"
  priority           = 100
  failover_threshold = 75
  timers_delay       = 30
  timers_reload      = 60
  preempt            = true
  shutdown           = false
  tracks = [
    {
      object_number = "1"
      shutdown      = true
    }
  ]
}
