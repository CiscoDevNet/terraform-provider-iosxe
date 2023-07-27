resource "iosxe_aaa_authorization" "example" {
  exec = [
    {
      name                = "TEST"
      a1_local            = false
      a1_if_authenticated = true
    }
  ]
}
