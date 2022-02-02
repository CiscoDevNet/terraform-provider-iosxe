# Adding VLAN & extended ACL
terraform {
  required_providers {
    iosxe = {
      source = "local.plugin/ciscodevnet/iosxe"
#      version = ">= 0.1"
    }
  }
}

provider "iosxe" {
  host = "https://10.1.1.5"
  insecure = true
  device_username = "admin"
  device_password = "Cisco123"
}

resource "iosxe_rest" "acl_example_post" {
  method = "POST"
  path = "/data/Cisco-IOS-XE-native:native/ip/access-list"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-acl:extended": [
          {
              "name": 102,
              "access-list-seq-rule": [
                  {
                      "sequence": "10",
                      "ace-rule": {
                          "action": "permit",
                          "protocol": "ip",
                          "host-address": "10.1.1.3",
                          "dst-any": [
                              null
                          ],
                          "precedence": "routine",
                          "tos": "normal",
                          "log": [
                              null
                          ]
                      }
                  },
                  {
                      "sequence": "20",
                      "ace-rule": {
                          "action": "permit",
                          "protocol": "tcp",
                          "any": [
                              null
                          ],
                          "dst-any": [
                              null
                          ],
                          "dst-eq": 600
                      }
                  },
                  {
                      "sequence": "30",
                      "ace-rule": {
                          "action": "permit",
                          "protocol": "udp",
                          "any": [
                              null
                          ],
                          "dst-any": [
                              null
                          ],
                          "dst-eq": 200
                      }
                  },
                  {
                      "sequence": "40",
                      "ace-rule": {
                          "action": "permit",
                          "protocol": "icmp",
                          "any": [
                              null
                          ],
                          "dst-any": [
                              null
                          ],
                          "dst-eq-port2": 250
                      }
                  },
                  {
                      "sequence": "50",
                      "ace-rule": {
                          "action": "permit",
                          "protocol": "igmp",
                          "any": [
                              null
                          ],
                          "dst-any": [
                              null
                          ],
                          "dst-eq-port2": 15
                      }
                  }
              ]
          }
      ]
  }
  )
}
