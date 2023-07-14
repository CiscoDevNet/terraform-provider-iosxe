resource "iosxe_snmp_server" "example" {
  chassis_id                        = "R1"
  contact                           = "Contact1"
  ifindex_persist                   = true
  location                          = "Location1"
  packetsize                        = 2000
  queue_length                      = 100
  enable_logging_getop              = true
  enable_logging_setop              = true
  enable_traps                      = true
  enable_traps_snmp_authentication  = true
  enable_traps_snmp_coldstart       = true
  enable_traps_snmp_linkdown        = true
  enable_traps_snmp_linkup          = true
  enable_traps_snmp_warmstart       = true
  source_interface_informs_loopback = 1
  source_interface_traps_loopback   = 1
  trap_source_loopback              = 1
  snmp_communities = [
    {
      name             = "COM1"
      view             = "VIEW1"
      permission       = "ro"
      ipv6             = "ACL1"
      access_list_name = "1"
    }
  ]
  contexts = [
    {
      name = "CON1"
    }
  ]
  views = [
    {
      name    = "VIEW1"
      mib     = "interfaces"
      inc_exl = "included"
    }
  ]
}
