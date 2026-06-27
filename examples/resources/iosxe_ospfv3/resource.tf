resource "iosxe_ospfv3" "example" {
  process_id                    = 1
  router_id                     = "1.2.3.4"
  bfd_all_interfaces            = true
  auto_cost_reference_bandwidth = 40000
  shutdown                      = false
  log_adjacency_changes         = true
  log_adjacency_changes_detail  = true
}
