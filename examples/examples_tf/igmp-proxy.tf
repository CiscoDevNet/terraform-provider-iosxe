# Adding IGMP_PROXY configuration
#
# POST method is not available for IGMP_PROXY
# resource "iosxe_rest" "igmp_proxy_example_post" {
#   method = "POST"
#   path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f18"
#   payload = jsonencode(
#   )
# }

# Updating the available IGMP_PROXY configuration for interface Gi1/0/18
resource "iosxe_rest" "igmp_proxy_example_put" {
  method = "PUT"
  path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f18"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-native:GigabitEthernet": {
            "name": "1/0/18",
            "switchport-conf": {
                "switchport": false
            },
            "ip": {
                "Cisco-IOS-XE-igmp:igmp": {
                    "mroute-proxy": {
                        "Loopback": 0
                    },
                    "unidirectional-link-leaf": [
                        null
                    ],
                    "upstream-proxy": "proxymap",
                    "iif-starg": [
                        null
                    ],
                    "proxy-report-interval": 130
                }
            }
        }
    }
  )
}

# Adding/Updating the available IGMP_PROXY configuration for interface Gi1/0/18
resource "iosxe_rest" "igmp_proxy_example_patch" {
  method = "PATCH"
  path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-native:GigabitEthernet": {
            "name": "1/0/18",
            "switchport-conf": {
                "switchport": false
            },
            "ip": {
                "Cisco-IOS-XE-igmp:igmp": {
                    "mroute-proxy": {
                        "Loopback": 0
                    },
                    "unidirectional-link-leaf": [
                        null
                    ],
                    "upstream-proxy": "proxymap",
                    "iif-starg": [
                        null
                    ],
                    "proxy-report-interval": 120
                }
            }
        }
    }
    )
}

# Fetch the available IGMP_PROXY configuration for interface Gi1/0/18
resource "iosxe_rest" "igmp_proxy_example_get" {
  method = "GET"
  path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f18"
}

# Remove the available IGMP_PROXY configuration for interface Gi1/0/18
resource "iosxe_rest" "igmp_proxy_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f18/ip/igmp"
}
