# Adding AAA-ACCOUNTING configuration
#
# POST method is not available for AAA-ACCOUNTING
# resource "iosxe_rest" "aaa_accounting_example_post" {
#   method = "POST"
#   path   = "/data/Cisco-IOS-XE-native:native/aaa/accounting"
#   payload = jsonencode(
#   )
# }

# Updating the available AAA-ACCOUNTING configuration
resource "iosxe_rest" "aaa_accounting_example_put" {
  method = "PUT"
  path   = "/data/Cisco-IOS-XE-native:native/aaa/accounting"
  payload = jsonencode(
    {
    "Cisco-IOS-XE-aaa:accounting": {
        "network": [
            {
                "id": "default-put"
            },
            {
                "id": "network2",
                "start-stop": {
                    "group-config": {
                        "group1": {
                            "group": "radius"
                        },
                        "group2": {
                            "group": "tacacs+"
                        }
                    }
                }
            }
        ],
        "system": {
            "guarantee-first": false
        }
    }
}
  )
}

# Adding/Updating the available AAA-ACCOUNTING configuration
resource "iosxe_rest" "aaa_accounting_example_patch" {
  method = "PATCH"
  path   = "/data/Cisco-IOS-XE-native:native/aaa/accounting"
  payload = jsonencode(
    {
    "Cisco-IOS-XE-aaa:accounting": {
        "network": [
            {
                "id": "default-patch"
            },
            {
                "id": "network2",
                "start-stop": {
                    "group-config": {
                        "group1": {
                            "group": "radius"
                        },
                        "group2": {
                            "group": "tacacs+"
                        }
                    }
                }
            }
        ],
        "system": {
            "guarantee-first": false
        }
    }
}
    )
}

# Fetch the available AAA-ACCOUNTING configuration 
resource "iosxe_rest" "aaa_accounting_example_get" {
  method = "GET"
  path   = "/data/Cisco-IOS-XE-native:native/aaa/accounting"
}

# Remove the available AAA-ACCOUNTING configuration 
resource "iosxe_rest" "aaa_accounting_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/aaa/accounting"
}
