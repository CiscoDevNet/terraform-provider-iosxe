resource "iosxe_rest" "vlan_example_post" {
  method = "POST"
  path = "/data/Cisco-IOS-XE-native:native/vlan"
  payload = jsonencode(
    {
    "Cisco-IOS-XE-vlan:vlan-list": {
          "id": "51",
          "name": "VLAN51"
      }
    }
  )
}

resource "iosxe_rest" "vlan_example_put" {
  method = "PUT"
  path = "/data/Cisco-IOS-XE-native:native/vlan/vlan-list=51"
  payload = jsonencode(
    {
    "Cisco-IOS-XE-vlan:vlan-list": {
          "id": "51",
          "name": "VLAN51-PUT"
      }
    }
  )
}

resource "iosxe_rest" "vlan_example_patch" {
  method = "PATCH"
  path = "/data/Cisco-IOS-XE-native:native/vlan"
  payload = jsonencode(
    {
    "Cisco-IOS-XE-native:vlan": {
        "Cisco-IOS-XE-vlan:vlan-list": [
            {
                "id": 41,
                "name": "VLAN41"
            },
            {
                "id": 42,
                "name": "VLAN42"
            },
            {
                "id": 51,
                "name": "VLAN51-PATCH"
            }
        ]
      }
  }
  )
}

resource "iosxe_rest" "vlan_example_get" {
  method = "GET"
  path = "/data/Cisco-IOS-XE-native:native/vlan"
}

resource "iosxe_rest" "vlan_example_get" {
  method = "GET"
  path = "/data/Cisco-IOS-XE-native:native/vlan/vlan-list=51"
}

resource "iosxe_rest" "vlan_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/vlan/vlan-list=51"
}
