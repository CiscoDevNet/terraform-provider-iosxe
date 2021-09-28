
# Adding NTP information
resource "iosxe_rest" "ntp_example_post" {
  method = "POST"
  path = "/data/Cisco-IOS-XE-native:native/ntp"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-ntp:authenticate": [
        null
      ],
      "Cisco-IOS-XE-ntp:authentication-key": [
        {
          "number": 1,
          "md5-cfg": "15110955002928757E31607615574F53570E5D0B04050F5A1A135C0A0405030102",
          "encryption-type": 7
        }
      ],
      "Cisco-IOS-XE-ntp:trusted-key": [
        {
          "number": 1,
          "hyphen": [
              null
          ],
          "end-key": 3
        }
      ]
    }
  )
}

# Updating the available NTP information
resource "iosxe_rest" "ntp_example_put" {
  method = "PUT"
  path = "/data/Cisco-IOS-XE-native:native/ntp"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-native:ntp": {
        "Cisco-IOS-XE-ntp:authenticate": [
          null
        ],
        "Cisco-IOS-XE-ntp:authentication-key": [
          {
            "number": 1,
            "md5-cfg": "15110955002928757E31607615574F53570E5D0B04050F5A1A135C0A0405030102",
            "encryption-type": 7
          },
          {
            "number": 2,
            "md5-cfg": "262837b2378ff0956a78ec8272673513",
            "encryption-type": 7
          }
        ],
        "Cisco-IOS-XE-ntp:server": {
          "server-list": [
            {
              "ip-address": "10.1.1.3"
            },
            {
              "ip-address": "10.1.1.5"
            },
            {
              "ip-address": "172.16.22.44",
              "key": 2
            }
          ]
        },
        "Cisco-IOS-XE-ntp:trusted-key": [
        {
          "number": 1,
          "hyphen": [
            null
          ],
          "end-key": 3
        }
        ]
      }
    }
  )
}

# Adding/Updating the NTP information
resource "iosxe_rest" "ntp_example_patch" {
  method = "PATCH"
  path = "/data/Cisco-IOS-XE-native:native/ntp"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-native:ntp": {
        "Cisco-IOS-XE-ntp:authenticate": [
          null
        ],
        "Cisco-IOS-XE-ntp:authentication-key": [
          {
            "number": 1,
            "md5-cfg": "15110955002928757E31607615574F53570E5D0B04050F5A1A135C0A0405030102",
            "encryption-type": 7
          }
        ],
        "Cisco-IOS-XE-ntp:server": {
          "server-list": [
            {
              "ip-address": "10.1.1.3"
            },
            {
              "ip-address": "10.1.1.5"
            },
            {
              "ip-address": "172.16.22.44",
              "key": 2
            }
          ]
        },
        "Cisco-IOS-XE-ntp:trusted-key": [
          {
            "number": 1,
            "hyphen": [
                null
            ],
            "end-key": 3
          }
        ]
      }
    }
  )
}

# Fetch the available NTP information
resource "iosxe_rest" "ntp_example_get" {
  method = "GET"
  path = "/data/Cisco-IOS-XE-native:native/ntp"
}

# Deletes all the NTP configuration
resource "iosxe_rest" "ntp_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/ntp"
}
