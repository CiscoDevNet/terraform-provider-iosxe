resource "iosxe_spanning_tree" "example" {
  mode                       = "mst"
  logging                    = true
  loopguard_default          = true
  portfast_default           = true
  portfast_bpduguard_default = true
  extend_system_id           = true
  mst_instances = [
    {
      id       = 1
      vlan_ids = [10]
    }
  ]
  vlans = [
    {
      id       = "10"
      priority = 32768
    }
  ]
  disabled_vlans = [
    {
      id = "20"
    }
  ]
}
