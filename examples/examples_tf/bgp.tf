# Adding BGP configuration
resource "iosxe_rest" "bgp_example_post" {
  method = "POST"
  path = "/data/Cisco-IOS-XE-native:native/router/bgp"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-bgp:bgp": {
            "id": 45000,
            "bgp": {
                "always-compare-med": [
                    null
                ],
                "bestpath": {
                    "as-path": "ignore",
                    "med": {
                        "confed-leaf": [
                            null
                        ],
                        "missing-as-worst-leaf": [
                            null
                        ]
                    }
                },
                "default": {
                    "local-preference": 200
                },
                "log-neighbor-changes": true
            },
            "neighbor": [
                {
                    "id": "10.106.1.2",
                    "remote-as": 65000
                }
            ]
        }
    }
  )
}

# Updating the available BGP configuration
resource "iosxe_rest" "bgp_example_put" {
  method = "PUT"
  path = "/data/Cisco-IOS-XE-native:native/router/bgp=46000"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-bgp:bgp": {
            "id": 46000,
            "bgp": {
                "always-compare-med": [
                    null
                ],
                "bestpath": {
                    "as-path": "ignore",
                    "med": {
                        "confed-leaf": [
                            null
                        ],
                        "missing-as-worst-leaf": [
                            null
                        ]
                    }
                },
                "default": {
                    "local-preference": 200
                },
                "log-neighbor-changes": true
            },
            "neighbor": [
                {
                    "id": "10.106.1.2",
                    "remote-as": 65000
                },
                {
                    "id": "10.107.1.2",
                    "remote-as": 65200
                }
            ]
        }
    }
  )
}

# Adding/Updating the available BGP configuration
resource "iosxe_rest" "bgp_example_patch" {
  method = "PATCH"
  path = "/data/Cisco-IOS-XE-native:native/router/bgp"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-bgp:bgp": {
            "id": 46000,
            "bgp": {
                "always-compare-med": [
                    null
                ],
                "bestpath": {
                    "as-path": "ignore",
                    "med": {
                        "confed-leaf": [
                            null
                        ],
                        "missing-as-worst-leaf": [
                            null
                        ]
                    }
                },
                "default": {
                    "local-preference": 200
                },
                "log-neighbor-changes": true
            },
            "neighbor": [
                {
                    "id": "10.108.1.2",
                    "remote-as": 65200
                }
            ]
        }
    }
  )
}

# Fetch the available BGP configuration
resource "iosxe_rest" "bgp_example_get" {
  method = "GET"
  path = "/data/Cisco-IOS-XE-native:native/router/bgp"
}

# Fetch the available BGP configuration by id
resource "iosxe_rest" "bgp_example_get_id" {
  method = "GET"
  path = "/data/Cisco-IOS-XE-native:native/router/bgp=46000"
}

# Remove the available BGP configuration by id
resource "iosxe_rest" "bgp_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/router/bgp=46000"
}
