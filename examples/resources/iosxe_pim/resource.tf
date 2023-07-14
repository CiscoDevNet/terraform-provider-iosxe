resource "iosxe_pim" "example" {
  autorp                 = false
  autorp_listener        = false
  bsr_candidate_loopback = 100
  bsr_candidate_mask     = 30
  bsr_candidate_priority = 10
  ssm_range              = "10"
  ssm_default            = false
  rp_address             = "9.9.9.9"
  rp_address_override    = false
  rp_address_bidir       = false
  rp_addresses = [
    {
      access_list = "10"
      rp_address  = "10.10.10.10"
      override    = false
      bidir       = false
    }
  ]
  rp_candidates = [
    {
      interface = "Loopback100"
      interval  = 100
      priority  = 10
      bidir     = false
    }
  ]
}
