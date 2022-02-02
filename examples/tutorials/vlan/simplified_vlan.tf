# Adding VLAN & extended ACL
terraform {
  required_providers {
    iosxe = {
      source = "local.plugin/ciscodevnet/iosxe"
#      version = ">= 0.1"
    }
  }
}

provider "iosxe" {
  host = "https://10.1.1.5"
  insecure = true
  device_username = "admin"
  device_password = "Cisco123"
}

resource "iosxe_rest" "vlan_example_put" {
  method = "PUT"
  path = "/data/Cisco-IOS-XE-native:native/vlan/vlan-list=511"
  payload = jsonencode(
    {
    "Cisco-IOS-XE-vlan:vlan-list": {
          "id": "511",
          "name": "VLAN511"
      }
    }
  )
}