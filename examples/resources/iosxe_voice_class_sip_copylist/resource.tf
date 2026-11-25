resource "iosxe_voice_class_sip_copylist" "example" {
  tag                    = 100
  sip_header_req_uri     = true
  sip_header_status_line = true
  headers                = ["History-Info"]
}
