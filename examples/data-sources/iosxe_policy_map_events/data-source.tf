data "iosxe_policy_map_events" "example" {
  name       = "dot1x_policy"
  event_type = "authentication-success"
}
