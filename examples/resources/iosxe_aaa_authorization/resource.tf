resource "iosxe_aaa_authorization" "example" {
  execs = [
    {
      name                = "TEST"
      a1_local            = false
      a1_if_authenticated = true
    }
  ]
}
