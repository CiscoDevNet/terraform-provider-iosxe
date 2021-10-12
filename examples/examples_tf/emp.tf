# Adding EMP configuration
#
# POST method is not available for EMP
# resource "iosxe_rest" "emp_example_post" {
#   method = "POST"
#   path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=0%2f0"
#   payload = jsonencode(
#   )
# }

# Updating the available EMP configuration
resource "iosxe_rest" "emp_example_put" {
  method = "PUT"
  path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=0%2f0"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-native:GigabitEthernet": {
            "name": "0/0",
            "description": "AP1-new-desc3-put",
            "vrf": {
                "forwarding": "Mgmt-vrf"
            }
        }
    }
  )
}

# Adding/Updating the available EMP configuration
resource "iosxe_rest" "emp_example_patch" {
  method = "PATCH"
  path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=0%2f0"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-native:GigabitEthernet": {
            "name": "0/0",
            "description": "AP1-new-desc3-patch",
            "vrf": {
                "forwarding": "Mgmt-vrf"
            },
            "ip": {
                "address": {
                    "primary": {
                        "address": "192.168.247.10",
                        "mask": "255.255.0.0"
                    }
                }
            },
            "Cisco-IOS-XE-ethernet:negotiation": {
                "auto": true
            }
        }
    }
    )
}

# Fetch the available EMP configuration 
resource "iosxe_rest" "emp_example_get" {
  method = "GET"
  path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=0%2f0"
}

# Remove the available EMP configuration 
resource "iosxe_rest" "emp_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=0%2f0/vrf"
}
