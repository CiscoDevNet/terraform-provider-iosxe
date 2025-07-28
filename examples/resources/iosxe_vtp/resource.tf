resource "iosxe_vtp" "example" {
  file             = "TEST"
  version          = 3
  interface        = "GigabitEthernet1/0/1"
  password         = "test123"
  password_hidden  = false
  domain           = "TESTDOMAIN"
  mode_transparent = true
}
