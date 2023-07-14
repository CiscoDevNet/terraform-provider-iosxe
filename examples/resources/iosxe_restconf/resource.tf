resource "iosxe_restconf" "simple" {
  path = "Cisco-IOS-XE-native:native/banner/login"
  attributes = {
    banner = "My Banner"
  }
}

resource "iosxe_restconf" "nested_list" {
  path = "Cisco-IOS-XE-native:native/ip"
  attributes = {
    source-route = "true"
  }
  lists = [{
    name = "vrf"
    key  = "name"
    items = [
      {
        name = "VRF1"
      }
    ]
  }]
}
