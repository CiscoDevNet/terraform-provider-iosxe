resource "iosxe_aaa_authentication" "example" {
  login = [
    {
      name = "test"
    }
  ]
  login_a1_auth_login_choice_group_group = "Radius-GROUP"
  login_a2_auth_login_choice_none_none   = true
  dot1x_default_group_                   = "Radius-GROUP"
}
