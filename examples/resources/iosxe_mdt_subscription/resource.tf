resource "iosxe_mdt_subscription" "example" {
  subscription_id         = 101
  stream                  = "yang-notif-native"
  encoding                = "encode-kvgpb"
  source_vrf              = "Mgmt-vrf"
  source_address          = "1.2.3.4"
  update_policy_on_change = true
  filter_xpath            = "/ios-events-ios-xe-oper:ospf-neighbor-state-change"
  receivers = [
    {
      address  = "5.6.7.8"
      port     = 57600
      protocol = "grpc-tcp"
    }
  ]
}
