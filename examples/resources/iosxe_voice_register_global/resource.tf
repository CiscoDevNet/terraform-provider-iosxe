resource "iosxe_voice_register_global" "example" {
  max_dn                 = 100
  max_pool               = 50
  system_message         = "CUBE SRST"
  security_policy_secure = false
}
