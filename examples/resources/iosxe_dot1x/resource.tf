resource "iosxe_dot1x" "example" {
  auth_fail_eapol = true
  dot1x_credentials = [
    {
      profile_name    = "profile1"
      description     = "credential_profile_name"
      username        = "username1"
      password_secret = "password123"
      pki_trustpoint  = "trustpoint1"
      anonymous_id    = "1"
    }
  ]
  critical_eapol_config_block     = true
  test_timeout                    = 1000
  logging_verbose                 = true
  supplicant_controlled_transient = true
  supplicant_force_multicast      = true
  system_auth_control             = true
}
