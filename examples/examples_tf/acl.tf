# Adding extended ACL
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

# Updating the extended ACL with name 102
resource "iosxe_rest" "acl_example_put" {
  method = "PUT"
  path = "/data/Cisco-IOS-XE-native:native/ip/access-list/extended=102"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-acl:extended": {
          "name": 102,
          "access-list-seq-rule": [
              {
                  "sequence": "10",
                  "ace-rule": {
                      "action": "permit",
                      "protocol": "ip",
                      "host-address": "10.1.1.5",
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
                      "dst-eq": 400
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
    }
  )
}

# Adding/Updating the extended ACLs
resource "iosxe_rest" "acl_example_patch" {
  method = "PATCH"
  path = "/data/Cisco-IOS-XE-native:native/ip/access-list"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-native:access-list": {
          "Cisco-IOS-XE-acl:extended": [
              {
                  "name": 101,
                  "access-list-seq-rule": [
                      {
                          "sequence": "10",
                          "ace-rule": {
                              "action": "permit",
                              "protocol": "ip",
                              "host-address": "10.1.1.2",
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
                              "dst-eq": 500
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
                              "dst-eq": 100
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
                              "dst-eq-port2": 200
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
                              "dst-eq-port2": 14
                          }
                      }
                  ]
              },
              {
                  "name": 102,
                  "access-list-seq-rule": [
                      {
                          "sequence": "10",
                          "ace-rule": {
                              "action": "permit",
                              "protocol": "ip",
                              "host-address": "10.1.1.4",
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
                              "dst-eq": 400
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
    }
  )
}

# Fetch all the access lists available
resource "iosxe_rest" "acl_example_get" {
  method = "GET"
  path = "/data/Cisco-IOS-XE-native:native/ip/access-list"
}

# Fetch a single extended ACL with name 102
resource "iosxe_rest" "acl_example_get_id" {
  method = "GET"
  path = "/data/Cisco-IOS-XE-native:native/ip/access-list/extended=102"
}

# Delete the extended ACL with name 102
resource "iosxe_rest" "acl_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/ip/access-list/extended=102"
}
