# Adding Line configuration
#
# POST method is not available for Line
# resource "iosxe_rest" "line_example_post" {
#   method = "POST"
#   path   = "/data/Cisco-IOS-XE-native:native/line"
#   payload = jsonencode(
#   )
# }

# Updating the available Line configuration
resource "iosxe_rest" "line_example_put" {
  method = "PUT"
  path   = "/data/Cisco-IOS-XE-native:native/line"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-native:line": {
        "auto-consolidation": false,
        "console": [
            {
                "first": "0",
                "exec-timeout": {
                    "minutes": 0,
                    "seconds": 0
                },
                "login": {
                    "authentication": "NOAUTH"
                },
                "stopbits": "1"
            }
        ],
        "vty": [
            {
                "first": 0,
                "exec-timeout": {
                    "minutes": 0,
                    "seconds": 0
                },
                "length": 0,
                "password": {
                    "secret": "login"
                },
                "transport": {
                    "input": {
                        "all": [
                            null
                        ]
                    }
                }
            },
            {
                "first": 1,
                "last": 4,
                "exec-timeout": {
                    "minutes": 0,
                    "seconds": 0
                },
                "transport": {
                    "input": {
                        "all": [
                            null
                        ]
                    }
                }
            },
            {
                "first": 5,
                "last": 31,
                "exec-timeout": {
                    "minutes": 0,
                    "seconds": 0
                },
                "transport": {
                    "input": {
                        "all": [
                            null
                        ]
                    }
                }
            },
            {
                "first": 32,
                "exec-timeout": {
                    "minutes": 0,
                    "seconds": 0
                },
                "transport": {
                    "input": {
                        "all": [
                            null
                        ]
                    }
                }
            }
        ]
    }
  }
  )
}

# Adding/Updating the available Line configuration
resource "iosxe_rest" "line_example_patch" {
  method = "PATCH"
  path   = "/data/Cisco-IOS-XE-native:native/line"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-native:line": {
        "auto-consolidation": false,
        "console": [
            {
                "first": "0",
                "exec-timeout": {
                    "minutes": 0,
                    "seconds": 0
                },
                "login": {
                    "authentication": "NOAUTH"
                },
                "stopbits": "1"
            }
        ],
        "vty": [
            {
                "first": 0,
                "exec-timeout": {
                    "minutes": 0,
                    "seconds": 0
                },
                "length": 0,
                "password": {
                    "secret": "login"
                },
                "transport": {
                    "input": {
                        "all": [
                            null
                        ]
                    }
                }
            },
            {
                "first": 1,
                "last": 4,
                "exec-timeout": {
                    "minutes": 0,
                    "seconds": 0
                },
                "transport": {
                    "input": {
                        "all": [
                            null
                        ]
                    }
                }
            },
            {
                "first": 5,
                "last": 31,
                "exec-timeout": {
                    "minutes": 0,
                    "seconds": 0
                },
                "transport": {
                    "input": {
                        "all": [
                            null
                        ]
                    }
                }
            },
            {
                "first": 32,
                "exec-timeout": {
                    "minutes": 0,
                    "seconds": 0
                },
                "transport": {
                    "input": {
                        "all": [
                            null
                        ]
                    }
                }
            }
        ]
    }
  }
  )
}

# Fetch the available Line configuration
resource "iosxe_rest" "line_example_get" {
  method = "GET"
  path   = "/data/Cisco-IOS-XE-native:native/line"
}

# Fetch the available vty by id in Line configuration
resource "iosxe_rest" "line_example_get_by_id" {
  method = "GET"
  path   = "/data/Cisco-IOS-XE-native:native/line/vty=5"
}


# Remove the available Line configuration
resource "iosxe_rest" "line_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/line/vty=32"
}
