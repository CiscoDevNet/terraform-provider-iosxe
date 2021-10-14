# Adding HSRP configuration
#
# POST method is not available for HSRP
# resource "iosxe_rest" "hsrp_example_post" {
#   method = "POST"
#   path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f13"
#   payload = jsonencode(
#   )
# }

# Updating the available HSRP configuration for interface Gi1/0/13
resource "iosxe_rest" "hsrp_example_put" {
  method = "PUT"
  path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f13"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-native:GigabitEthernet": {
            "name": "1/0/13",
            "switchport-conf": {
                "switchport": false
            },
            "ip": {
                "address": {
                    "primary": {
                        "address": "10.0.0.1",
                        "mask": "255.255.255.0"
                    }
                }
            },
            "standby": {
                "standby-list": [
                    {
                        "group-number": 1,
                        "ip": {
                            "address": "10.0.0.3"
                        },
                        "priority": 110
                    }
                ]
            }
        }
    }
  )
}

# Adding/Updating the available HSRP configuration for interface Gi1/0/13
resource "iosxe_rest" "hsrp_example_patch" {
  method = "PATCH"
  path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet"
  payload = jsonencode(
   {
        "Cisco-IOS-XE-native:GigabitEthernet": {
            "name": "1/0/13",
            "switchport-conf": {
                "switchport": false
            },
            "ip": {
                "address": {
                    "primary": {
                        "address": "10.0.0.1",
                        "mask": "255.255.255.0"
                    }
                }
            },
            "standby": {
                "standby-list": [
                    {
                        "group-number": 1,
                        "ip": {
                            "address": "10.0.0.3"
                        },
                        "priority": 110
                    },
                    {
                        "group-number": 2,
                        "ip": {
                            "address": "10.0.0.4"
                        },
                        "priority": 115
                    }
                ]
            }
        }
    }
    )
}

# Fetch the available HSRP configuration for interface Gi1/0/13
resource "iosxe_rest" "hsrp_example_get" {
  method = "GET"
  path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f13"
}

# Remove the available HSRP configuration for interface Gi1/0/13
resource "iosxe_rest" "hsrp_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f13/standby"
}
