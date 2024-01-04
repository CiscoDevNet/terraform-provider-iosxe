resource "iosxe_aaa_authentication" "example" {
  logins = [
    {
      name     = "test"
      a1_group = "Radius-GROUP"
      a2_none  = true
    }
  ]
  dot1x = [
    {
      name      = "test"
      a1_group  = "GROUP1"
      a2_cache  = "GROUP2"
      a3_radius = true
      a4_local  = true
    }
  ]
  dot1x_default_a1_group = "Radius-GROUP"
  dot1x_default_a2_group = "Radius-GROUP2"
}
