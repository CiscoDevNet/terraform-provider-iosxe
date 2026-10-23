resource "iosxe_voice_class_sip_profiles" "example" {
  tag         = 100
  description = "Header normalization"
  rules = [
    {
      rule                              = 10
      request_method                    = "ANY"
      request_sip_header                = "Contact"
      request_sip_header_modify_pattern = "@.*:"
      request_sip_header_modify_replace = "@cube.example.com:"
    }
  ]
}
