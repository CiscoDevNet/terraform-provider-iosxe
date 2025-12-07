resource "iosxe_l2vpn_evpn_profile" "example" {
  name       = "MY_EVPN_PROFILE"
  evi_base   = 1000
  l2vni_base = 10000
}
