resource "iosxe_vrf" "example" {
  name                = "VRF22"
  description         = "VRF22 description"
  rd                  = "1342"
  address_family_ipv4 = true
  address_family_ipv6 = true
  vpn_id              = "1342"
  ipv4_route_target_import = [
    {
      value = "1342"
    }
  ]
  ipv4_route_target_import_stitching = [
    {
      value = "1342"
    }
  ]
  ipv4_route_target_export = [
    {
      value = "1342"
    }
  ]
  ipv4_route_target_export_stitching = [
    {
      value = "1342"
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
      value = "1342"
    }
  ]
  ipv6_route_target_import_stitching = [
    {
      value = "1342"
    }
  ]
  ipv6_route_target_export = [
    {
      value = "1342"
    }
  ]
  ipv6_route_target_export_stitching = [
    {
      value = "1342"
    }
  ]
  vnid = [
    {
      vnid_value = 10001
      evpn_instance_vni_vni_num = [
        {
          vni_num   = 20000
          core_vlan = 200
        }
      ]
    }
  ]
}
