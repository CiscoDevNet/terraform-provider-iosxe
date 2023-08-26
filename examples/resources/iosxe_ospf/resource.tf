resource "iosxe_ospf" "example" {
  process_id                           = 1
  bfd_all_interfaces                   = true
  default_information_originate        = true
  default_information_originate_always = true
  default_metric                       = 21
  distance                             = 120
  domain_tag                           = 10
  neighbor = [
    {
      ip       = "2.2.2.2"
      priority = 10
      cost     = 100
    }
  ]
  priority  = 100
  router_id = "1.2.3.4"
  shutdown  = false
  summary_address = [
    {
      ip   = "3.3.3.0"
      mask = "255.255.255.0"
    }
  ]
  area_id = [
    {
      area_id                       = "233"
      authentication_message_digest = true
      nssa                          = true
    }
  ]
  passive_interface_passive_interface_choice_default_default = false
  networks = [
    {
      ip       = "10.1.1.1"
      wildcard = "255.255.255.248"
      area     = "3"
    }
  ]
}
