resource "iosxe_ospf" "example" {
  process_id                           = 1
  bfd_all_interfaces                   = true
  default_information_originate        = true
  default_information_originate_always = true
  default_metric                       = 21
  distance                             = 120
  domain_tag                           = 10
  neighbors = [
    {
      ip       = "2.2.2.2"
      priority = 10
      cost     = 100
    }
  ]
  networks = [
    {
      ip       = "3.3.3.0"
      wildcard = "0.0.0.255"
      area     = "0"
    }
  ]
  priority  = 100
  router_id = "1.2.3.4"
  shutdown  = false
  summary_addresses = [
    {
      ip   = "3.3.3.0"
      mask = "255.255.255.0"
    }
  ]
  areas = [
    {
      area_id                                        = "5"
      authentication_message_digest                  = true
      nssa                                           = true
      nssa_default_information_originate             = true
      nssa_default_information_originate_metric      = 100
      nssa_default_information_originate_metric_type = 1
      nssa_no_summary                                = true
      nssa_no_redistribution                         = true
    }
  ]
  passive_interface_default = true
}
