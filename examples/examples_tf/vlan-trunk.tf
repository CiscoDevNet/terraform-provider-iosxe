# Adding VLAN-TRUNK configuration
#
# POST method is not available for VLAN-TRUNK
# resource "iosxe_rest" "vlan_trunk_example_post" {
#   method = "POST"
#   path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1{{/}}0{{/}}11/switchport-config"
#   payload = jsonencode(
#   )
# }

# Updating the available VLAN-TRUNK configuration for interface Gi1/0/11
resource "iosxe_rest" "vlan_trunk_example_put" {
  method = "PUT"
  path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1{{/}}0{{/}}11/switchport-config"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-native:switchport-config": {
            "switchport": {
                "Cisco-IOS-XE-switch:trunk": {
                    "allowed": {
                        "vlan": {
                            "vlans": "20,30"
                        }
                    },
                    "native": {
                        "vlan": {
                            "vlan-id": 100
                        }
                    }
                }
            }
        }
    }
  )
}

# Adding/Updating the available VLAN-TRUNK configuration for interface Gi1/0/11
resource "iosxe_rest" "vlan_trunk_example_patch" {
  method = "PATCH"
  path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-native:GigabitEthernet": {
            "name": "1/0/11",
            "switchport-config": {
                "switchport": {
                    "Cisco-IOS-XE-switch:trunk": {
                        "allowed": {
                            "vlan": {
                                "vlans": "20,30,40"
                            }
                        },
                        "native": {
                            "vlan": {
                                "vlan-id": 100
                            }
                        }
                    }
                }
            }
        }
    }
    )
}

# Fetch the available VLAN-TRUNK configuration for interface Gi1/0/11
resource "iosxe_rest" "vlan_trunk_example_get" {
  method = "GET"
  path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1{{/}}0{{/}}11"
}

# Remove the available VLAN-TRUNK configuration for interface Gi1/0/11
resource "iosxe_rest" "vlan_trunk_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1{{/}}0{{/}}11/switchport-config"
}
