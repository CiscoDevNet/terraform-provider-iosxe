# Adding VLAN-VOICE configuration
#
# POST method is not available for VLAN-VOICE
# resource "iosxe_rest" "vlan_voice_example_post" {
#   method = "POST"
#   path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1{{/}}0{{/}}12/switchport-config"
#   payload = jsonencode(
#   )
# }

# Updating the available VLAN-VOICE configuration for interface Gi1/0/12
resource "iosxe_rest" "vlan_voice_example_put" {
  method = "PUT"
  path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1{{/}}0{{/}}12/switchport-config"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-native:switchport-config": {
            "switchport": {
                "Cisco-IOS-XE-switch:voice": {
                    "vlan": {
                        "vlan": "dot1p"
                    }
                }
            }
        }
    }
  )
}

# Adding/Updating the available VLAN-VOICE configuration for interface Gi1/0/12
resource "iosxe_rest" "vlan_voice_example_patch" {
  method = "PATCH"
  path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-native:GigabitEthernet": {
            "name": "1/0/12",
            "description": "AP1 updated desc",
            "switchport-config": {
                "switchport": {
                    "Cisco-IOS-XE-switch:voice": {
                        "vlan": {
                            "vlan": "dot1p"
                        }
                    }
                }
            }
        }
    }
    )
}

# Fetch the available VLAN-VOICE configuration for interface Gi1/0/12
resource "iosxe_rest" "vlan_voice_example_get" {
  method = "GET"
  path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1{{/}}0{{/}}12"
}

# Remove the available VLAN-VOICE configuration for interface Gi1/0/12
resource "iosxe_rest" "vlan_voice_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1{{/}}0{{/}}12/switchport-config"
}
