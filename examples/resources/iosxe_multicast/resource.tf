resource "iosxe_multicast" "example" {
  multipath          = true
  multipath_s_g_hash = "next-hop-based"
  vrfs = [
    {
      vrf                = "VRF1"
      multipath          = true
      multipath_s_g_hash = "next-hop-based"
    }
  ]
}
