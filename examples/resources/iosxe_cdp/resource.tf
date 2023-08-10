resource "iosxe_cdp" "example" {
  holdtime        = 15
  timer           = 5
  run_enable      = true
  filter_tlv_list = "TLIST"
  tlv_list = [
    {
      name            = "TLIST"
      vtp_mgmt_domain = true
      cos             = true
      duplex          = true
      trust           = true
      version         = true
    }
  ]
}
