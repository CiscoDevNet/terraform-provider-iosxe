resource "iosxe_vrf" "example" {
  name                = "VRF22"
  description         = "VRF22 description"
  rd_auto             = true
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
  ipv4_mdt_default_address               = "239.1.1.1"
  ipv4_mdt_auto_discovery_vxlan          = true
  ipv4_mdt_auto_discovery_vxlan_inter_as = true
  ipv4_mdt_overlay_use_bgp               = true
  ipv4_mdt_overlay_use_bgp_spt_only      = true
  ipv4_mdt_data_multicast = [
    {
      address  = "239.1.2.0"
      wildcard = "0.0.0.255"
    }
  ]
  ipv4_mdt_data_threshold = 50
}
