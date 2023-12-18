resource "iosxe_line" "example" {
  console = [
    {
      first                = "0"
      exec_timeout_minutes = 45
      exec_timeout_seconds = 25
      privilege_level      = 15
      stopbits             = "1"
      password_type        = "0"
      password             = "testpasswd"
    }
  ]
  vty = [
    {
      first = 10
      last  = 27
      access_classes = [
        {
          direction   = "in"
          access_list = "2"
          vrf_also    = true
        }
      ]
      exec_timeout_minutes         = 45
      exec_timeout_seconds         = 25
      password_type                = "0"
      password                     = "testpasswd"
      transport_preferred_protocol = "none"
      escape_character             = "27"
      authorization_exec_default   = true
      transport_input              = "ssh"
    }
  ]
}
