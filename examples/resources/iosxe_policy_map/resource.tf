resource "iosxe_policy_map" "example" {
  name       = "dot1x_policy"
  type       = "control"
  subscriber = true
}
