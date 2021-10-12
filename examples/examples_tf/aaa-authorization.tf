# Adding AAA-AUTHORIZATION configuration
#
# POST method is not available for AAA-AUTHORIZATION
# resource "iosxe_rest" "aaa_authorization_example_post" {
#   method = "POST"
#   path   = "/data/Cisco-IOS-XE-native:native/aaa/authorization"
#   payload = jsonencode(
      
#   )
# }

# Updating the available AAA-AUTHORIZATION configuration
resource "iosxe_rest" "aaa_authorization_example_put" {
  method = "PUT"
  path   = "/data/Cisco-IOS-XE-native:native/aaa/authorization"
  payload = jsonencode(
    {
    "Cisco-IOS-XE-aaa:authorization": {
        "exec": [
            {
                "name": "default",
                "a1": {
                    "group": "ISE"
                },
                "a2": {
                    "local": [
                        null
                    ]
                }
            }
        ]
    }
}
  )
}

# Adding/Updating the available AAA-AUTHORIZATION configuration
resource "iosxe_rest" "aaa_authorization_example_patch" {
  method = "PATCH"
  path   = "/data/Cisco-IOS-XE-native:native/aaa/authorization"
  payload = jsonencode(
    {
    "Cisco-IOS-XE-aaa:authorization": {
        "exec": [
            {
                "name": "default",
                "a1": {
                    "group": "ISE"
                },
                "a2": {
                    "local": [
                        null
                    ]
                }
            }
        ]
    }
}
    )
}

# Fetch the available AAA-AUTHORIZATION configuration 
resource "iosxe_rest" "aaa_authorization_example_get" {
  method = "GET"
  path   = "/data/Cisco-IOS-XE-native:native/aaa/authorization"
}

# Remove the available AAA-AUTHORIZATION configuration 
resource "iosxe_rest" "aaa_authorization_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/aaa/authorization"
}
