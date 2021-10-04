
# Adding a new MDT subscription
resource "iosxe_rest" "mdt_example_post" {
  method = "POST"
  path   = "/data/Cisco-IOS-XE-mdt-cfg:mdt-config-data"
  payload = jsonencode(
    {
      "mdt-subscription" : [
        {
          "subscription-id" : 105,
          "base" : {
            "stream" : "yang-push",
            "encoding" : "encode-kvgpb",
            "source-vrf" : "Mgmt-intf",
            "period" : 6000,
            "xpath" : "/memory-ios-xe-oper:memory-statisticts/memory-statistic"
          },
          "mdt-receivers" : [
            {
              "address" : "10.28.35.45",
              "port" : 57555,
              "protocol" : "grpc-tcp"
            }
          ]
        }
      ]
    }
  )
}

# Updating the available configuration of MDT subscription
resource "iosxe_rest" "mdt_example_put" {
  method = "PUT"
  path   = "/data/Cisco-IOS-XE-mdt-cfg:mdt-config-data"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-mdt-cfg:mdt-config-data" : {
        "mdt-subscription" : [
          {
            "subscription-id" : 105,
            "base" : {
              "stream" : "yang-push",
              "encoding" : "encode-kvgpb",
              "source-vrf" : "Mgmt-intf",
              "period" : 6500,
              "xpath" : "/memory-ios-xe-oper:memory-statisticts/memory-statistic"
            },
            "mdt-receivers" : [
              {
                "address" : "10.28.35.45",
                "port" : 57555,
                "protocol" : "grpc-tcp"
              }
            ]
          },
          {
            "subscription-id" : 1200,
            "base" : {
              "stream" : "yang-push",
              "encoding" : "encode-kvgpb",
              "no-synch-on-start" : false,
              "xpath" : "/iosxe-oper:ios-oper-db/hwidb-table"
            }
          }
        ],
        "mdt-named-protocol-rcvrs" : {
          "mdt-named-protocol-rcvr" : [
            {
              "name" : "yangsuite",
              "protocol" : "grpc-tcp",
              "host" : {
                "hostname" : "yangsuite-telemetry.cisco.com"
              },
              "port" : 57500
            }
          ]
        }
      }
    }
  )
}

# Adding/Updating the configuration of MDT subscription
resource "iosxe_rest" "mdt_example_patch" {
  method = "PATCH"
  path   = "/data/Cisco-IOS-XE-mdt-cfg:mdt-config-data"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-mdt-cfg:mdt-config-data" : {
        "mdt-subscription" : [
          {
            "subscription-id" : 105,
            "base" : {
              "stream" : "yang-push",
              "encoding" : "encode-kvgpb",
              "source-vrf" : "Mgmt-intf",
              "period" : 6000,
              "xpath" : "/cdp-ios-xe-oper:cdp-neighbor-details/cdp-neighbor-detail"
            },
            "mdt-receivers" : [
              {
                "address" : "10.28.35.45",
                "port" : 57555,
                "protocol" : "grpc-tcp"
              }
            ]
          }
        ]
      }
    }
  )
}

# Fetch all the available MDT configuration
resource "iosxe_rest" "mdt_example_get" {
  method = "GET"
  path   = "/data/Cisco-IOS-XE-mdt-cfg:mdt-config-data"
}

# Fetch the MDT subscription details with ID 1200
resource "iosxe_rest" "mdt_example_get_id" {
  method = "GET"
  path   = "/data/Cisco-IOS-XE-mdt-cfg:mdt-config-data/mdt-subscription=1200"
}

# Remove the MDT subscription with ID 1200
resource "iosxe_rest" "mdt_example_delete" {
  method = "DELETE"
  path   = "/data/Cisco-IOS-XE-mdt-cfg:mdt-config-data/mdt-subscription=1200"
}
