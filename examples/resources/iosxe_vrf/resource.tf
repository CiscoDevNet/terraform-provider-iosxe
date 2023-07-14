resource "iosxe_vrf" "example" {
  name                = "VRF22"
  description         = "VRF22 description"
  rd                  = "22:22"
  address_family_ipv4 = true
  address_family_ipv6 = true
  vpn_id              = "22:22"
  route_target_import = [
    {
      value     = "22:22"
      stitching = false
    }
  ]
  route_target_export = [
    {
      value     = "22:22"
      stitching = false
    }
  ]
  ipv4_route_target_import = [
    {
      value = "22:22"
    }
  ]
  ipv4_route_target_import_stitching = [
    {
      value = "22:22"
    }
  ]
  ipv4_route_target_export = [
    {
      value = "22:22"
    }
  ]
  ipv4_route_target_export_stitching = [
    {
      value = "22:22"
    }
  ]
  ipv6_route_target_import = [
    {
      value = "22:22"
    }
  ]
  ipv6_route_target_import_stitching = [
    {
      value = "22:22"
    }
  ]
  ipv6_route_target_export = [
    {
      value = "22:22"
    }
  ]
  ipv6_route_target_export_stitching = [
    {
      value = "22:22"
    }
  ]
}
