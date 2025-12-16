resource "iosxe_vrf" "example" {
  name                = "VRF22"
  description         = "VRF22 description"
  rd                  = "22:22"
  address_family_ipv4 = true
  address_family_ipv6 = true
  vpn_id              = "22:22"
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
  ipv4_route_replicate = [
    {
      name                  = "VRF1"
      unicast_all           = true
      unicast_all_route_map = "RM1"
    }
  ]
  ipv4_import_map = "IMPORT-MAP-1"
  ipv4_export_map = "EXPORT-MAP-1"
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
  ipv6_import_map = "IMPORT-MAP-1"
  ipv6_export_map = "EXPORT-MAP-1"
}
