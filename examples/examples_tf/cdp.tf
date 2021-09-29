
# Adding CDP information
resource "iosxe_rest" "cdp_example_post" {
  method = "POST"
  path = "/data/Cisco-IOS-XE-native:native/cdp"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-cdp:holdtime": 60,
      "Cisco-IOS-XE-cdp:timer": 20,
      "Cisco-IOS-XE-cdp:run-enable": true,
      "Cisco-IOS-XE-cdp:tlv-list": [
        {
          "name": "CDP-1",
          "version": "",
          "vtp-mgmt-domain": "",
          "cos": "",
          "duplex": "",
          "trust": ""
        }
      ],
      "Cisco-IOS-XE-cdp:filter-tlv-list": "CDP-1"
    }
  )
}

# Updating the available CDP information
resource "iosxe_rest" "cdp_example_put" {
  method = "PUT"
  path = "/data/Cisco-IOS-XE-native:native/cdp"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-native:cdp": {
        "Cisco-IOS-XE-cdp:holdtime": 50,
        "Cisco-IOS-XE-cdp:timer": 30,
        "Cisco-IOS-XE-cdp:run-enable": true,
        "Cisco-IOS-XE-cdp:tlv-list": [
          {
            "name": "CDP-1",
            "version": "",
            "vtp-mgmt-domain": "",
            "cos": "",
            "duplex": "",
            "trust": ""
          }
        ],
        "Cisco-IOS-XE-cdp:filter-tlv-list": "CDP-1"
      }
    }
  )
}

# Adding/Updating the CDP information
resource "iosxe_rest" "cdp_example_patch" {
  method = "PATCH"
  path = "/data/Cisco-IOS-XE-native:native/cdp"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-native:cdp": {
        "Cisco-IOS-XE-cdp:holdtime": 60,
        "Cisco-IOS-XE-cdp:timer": 20,
        "Cisco-IOS-XE-cdp:run-enable": true,
        "Cisco-IOS-XE-cdp:tlv-list": [
          {
            "name": "CDP-1",
            "version": "",
            "vtp-mgmt-domain": "",
            "cos": "",
            "duplex": "",
            "trust": ""
          },            
          {
            "name": "CDP-2",
            "version": "",
            "vtp-mgmt-domain": "",
            "cos": "",
            "duplex": "",
            "trust": ""
          }
        ],
        "Cisco-IOS-XE-cdp:filter-tlv-list": "CDP-2"
      }
    }
  )
}

# Fetch the available CDP information
resource "iosxe_rest" "cdp_example_get" {
  method = "GET"
  path = "/data/Cisco-IOS-XE-native:native/cdp"
}

# Deletes all the CDP configuration
resource "iosxe_rest" "cdp_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/cdp"
}
