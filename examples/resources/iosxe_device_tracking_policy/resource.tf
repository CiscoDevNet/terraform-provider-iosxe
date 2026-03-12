resource "iosxe_device_tracking_policy" "example" {
  name         = "DT_trunk_policy"
  trusted_port = true
}
