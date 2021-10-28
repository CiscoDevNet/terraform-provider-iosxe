# Adding EMP configuration
#
# POST method is not available for EMP
# resource "iosxe_rest" "emp_example_post" {
#   method = "POST"
#   path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f1"
#   payload = jsonencode(
#   )
# }

# Updating the available EMP configuration for interface Gi1/0/1
resource "iosxe_rest" "emp_example_put" {
  method = "PUT"
  path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f1"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-native:GigabitEthernet": {
            "name": "1/0/1",
            "description": "AP1-new-desc",
            "switchport-conf": {
                "switchport": false
            },
            "vrf": {
                "forwarding": "Dev-vrf"
            }
        }
    }
  )
}

# Adding/Updating the available EMP configuration for interface Gi1/0/1
resource "iosxe_rest" "emp_example_patch" {
  method = "PATCH"
  path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-native:GigabitEthernet": {
            "name": "1/0/1",
            "description": "AP1-new-desc",
            "switchport-conf": {
                "switchport": false
            },
            "vrf": {
                "forwarding": "Dev-vrf"
            }
        }
    }
    )
}

# Fetch the available EMP configuration for interface Gi1/0/1
resource "iosxe_rest" "emp_example_get" {
  method = "GET"
  path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f1"
}

# Remove the available EMP configuration for interface Gi1/0/1
resource "iosxe_rest" "emp_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f1/switchport-conf"
}
