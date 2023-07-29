resource "iosxe_aaa_authentication" "example" {
  logins = [
    {
      name     = "test"
      a1_group = "Radius-GROUP"
      a2_none  = true
    }
  ]
  dot1x_default_a1_group = "Radius-GROUP"
  dot1x_default_a2_group = "Radius-GROUP2"
}
