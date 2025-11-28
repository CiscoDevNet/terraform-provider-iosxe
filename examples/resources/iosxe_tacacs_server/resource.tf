resource "iosxe_tacacs_server" "example" {
  timeout                      = 5
  directed_request             = true
  directed_request_restricted  = true
  directed_request_no_truncate = true
  encryption                   = "0"
  key                          = "123"
  attribute_allow_unknown      = true
}
