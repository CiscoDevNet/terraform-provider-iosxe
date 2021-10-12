# Adding AAA-AUTHENTICATION configuration
resource "iosxe_rest" "aaa_authentication_example_post" {
  method = "POST"
  path   = "/data/Cisco-IOS-XE-native:native/aaa/authentication"
  payload = jsonencode(
      {
        "login": [
            {
                "name": "default",
                "a1": {
                    "group": "tacacs+"
                },
                "a2": {
                    "none": [
                        null
                    ]
                }
            }
        ],
        "ppp": [
            {
                "id": "server-group1",
                "a1": {
                    "group": "radius"
                },
                "a2": {
                    "group": "tacacs+"
                },
                "a3": {
                    "local": [
                        null
                    ]
                },
                "a4": {
                    "none": [
                        null
                    ]
                }
            }
        ]
    }
  )
}

# Updating the available AAA-AUTHENTICATION configuration
resource "iosxe_rest" "aaa_authentication_example_put" {
  method = "PUT"
  path   = "/data/Cisco-IOS-XE-native:native/aaa/authentication"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-aaa:authentication": {
            "login": [
                {
                    "name": "default-put",
                    "a1": {
                        "group": "tacacs+"
                    },
                    "a2": {
                        "none": [
                            null
                        ]
                    }
                }
            ],
            "ppp": [
                {
                    "id": "server-group1",
                    "a1": {
                        "group": "radius"
                    },
                    "a2": {
                        "group": "tacacs+"
                    },
                    "a3": {
                        "local": [
                            null
                        ]
                    },
                    "a4": {
                        "none": [
                            null
                        ]
                    }
                }
            ]
        }
    }
  )
}

# Adding/Updating the available AAA-AUTHENTICATION configuration
resource "iosxe_rest" "aaa_authentication_example_patch" {
  method = "PATCH"
  path   = "/data/Cisco-IOS-XE-native:native/aaa/authentication"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-aaa:authentication": {
            "login": [
                {
                    "name": "default-patch",
                    "a1": {
                        "group": "tacacs+"
                    },
                    "a2": {
                        "none": [
                            null
                        ]
                    }
                }
            ],
            "ppp": [
                {
                    "id": "server-group1",
                    "a1": {
                        "group": "radius"
                    },
                    "a2": {
                        "group": "tacacs+"
                    },
                    "a3": {
                        "local": [
                            null
                        ]
                    },
                    "a4": {
                        "none": [
                            null
                        ]
                    }
                }
            ]
        }
    }
    )
}

# Fetch the available AAA-AUTHENTICATION configuration 
resource "iosxe_rest" "aaa_authentication_example_get" {
  method = "GET"
  path   = "/data/Cisco-IOS-XE-native:native/aaa/authentication"
}

# Remove the available AAA-AUTHENTICATION configuration 
resource "iosxe_rest" "aaa_authentication_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/aaa/authentication"
}
