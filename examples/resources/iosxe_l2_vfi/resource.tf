resource "iosxe_l2_vfi" "example" {
  name   = "TENANT-A"
  mode   = "manual"
  vpn_id = 20001
  neighbors = [
    {
      ip_address    = "172.16.255.2"
      encapsulation = "mpls"
    }
  ]
}
