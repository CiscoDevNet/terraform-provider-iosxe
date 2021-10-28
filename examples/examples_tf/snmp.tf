# Adding SNMP configuration
resource "iosxe_rest" "snmp_example_post" {
  method = "POST"
  path   = "/data/Cisco-IOS-XE-native:native/snmp-server"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-snmp:community-config": [
            {
                "name": "public",
                "permission": "ro"
            }
        ],
        "Cisco-IOS-XE-snmp:contact": "Dial System Operator at beeper 21555",
        "Cisco-IOS-XE-snmp:group": [
            {
                "id": "authgroup",
                "v3": {
                    "security-level-list": [
                        {
                            "security-level": "auth"
                        }
                    ]
                }
            }
        ],
        "Cisco-IOS-XE-snmp:location": "Building 3/Room 222"
    }
  )
}

# Updating the available SNMP configuration 
resource "iosxe_rest" "snmp_example_put" {
  method = "PUT"
  path   = "/data/Cisco-IOS-XE-native:native/snmp-server"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-native:snmp-server": {
            "Cisco-IOS-XE-snmp:community-config": [
                {
                    "name": "public",
                    "permission": "ro"
                }
            ],
            "Cisco-IOS-XE-snmp:contact": "Dial System Operator at beeper 21555",
            "Cisco-IOS-XE-snmp:engineID": {
                "local": "1234567890",
                "remote-conf": {
                    "host-list": [
                        {
                            "ip-address": "192.180.1.27",
                            "engine-id": "00000063000100A1C0B4011B"
                        }
                    ]
                }
            },
            "Cisco-IOS-XE-snmp:group": [
                {
                    "id": "authgroup",
                    "v3": {
                        "security-level-list": [
                            {
                                "security-level": "auth"
                            }
                        ]
                    }
                }
            ],
            "Cisco-IOS-XE-snmp:host-config": {
                "ip-community": [
                    {
                        "ip-address": "192.180.1.27",
                        "community-or-user": "authuser",
                        "informs": [
                            null
                        ],
                        "version": "3",
                        "security-level": "auth",
                        "trap-enable": {
                            "config": [
                                null
                            ]
                        }
                    },
                    {
                        "ip-address": "192.180.1.27",
                        "community-or-user": "public",
                        "version": "2c"
                    },
                    {
                        "ip-address": "192.180.1.33",
                        "community-or-user": "public"
                    },
                    {
                        "ip-address": "192.180.1.111",
                        "community-or-user": "public"
                    }
                ]
            },
            "Cisco-IOS-XE-snmp:location": "Building 3/Room 222"
        }
    }
  )
}

# Adding/Updating the available SNMP configuration 
resource "iosxe_rest" "snmp_example_patch" {
  method = "PATCH"
  path   = "/data/Cisco-IOS-XE-native:native/snmp-server"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-native:snmp-server": {
            "Cisco-IOS-XE-snmp:community-config": [
                {
                    "name": "public",
                    "permission": "ro"
                }
            ],
            "Cisco-IOS-XE-snmp:contact": "Dial System Operator at beeper 21555",
            "Cisco-IOS-XE-snmp:enable": {
                "enable-choice": {
                    "traps": {
                        "ike": {
                            "policy": {
                                "add": [
                                    null
                                ],
                                "delete": [
                                    null
                                ]
                            },
                            "tunnel": {
                                "start": [
                                    null
                                ],
                                "stop": [
                                    null
                                ]
                            }
                        },
                        "local-auth": [
                            null
                        ],
                        "auth-framework": {
                            "sec-violation": [
                                null
                            ]
                        },
                        "bfd": [
                            null
                        ],
                        "bgp-traps": {
                            "cbgp2": [
                                null
                            ]
                        },
                        "bridge": {
                            "newroot": [
                                null
                            ],
                            "topologychange": [
                                null
                            ]
                        },
                        "bulkstat": {
                            "collection": [
                                null
                            ],
                            "transfer": [
                                null
                            ]
                        },
                        "call-home": {
                            "message-send-fail": [
                                null
                            ],
                            "server-fail": [
                                null
                            ]
                        },
                        "cef": {
                            "resource-failure": [
                                null
                            ],
                            "peer-state-change": [
                                null
                            ],
                            "peer-fib-state-change": [
                                null
                            ],
                            "inconsistency": [
                                null
                            ]
                        },
                        "config": [
                            null
                        ],
                        "config-copy": [
                            null
                        ],
                        "config-ctid": [
                            null
                        ],
                        "cpu": {
                            "threshold": [
                                null
                            ]
                        },
                        "dhcp": [
                            null
                        ],
                        "eigrp": [
                            null
                        ],
                        "energywise": {},
                        "entity": [
                            null
                        ],
                        "entity-diag": {
                            "boot-up-fail": [
                                null
                            ],
                            "hm-test-recover": [
                                null
                            ],
                            "hm-thresh-reached": [
                                null
                            ],
                            "scheduled-test-fail": [
                                null
                            ]
                        },
                        "envmon": {},
                        "errdisable": {},
                        "event-manager": [
                            null
                        ],
                        "flowmon": [
                            null
                        ],
                        "fru-ctrl": [
                            null
                        ],
                        "hsrp": [
                            null
                        ],
                        "ipmulticast": [
                            null
                        ],
                        "ipsec": {
                            "cryptomap": {
                                "add": [
                                    null
                                ],
                                "attach": [
                                    null
                                ],
                                "delete": [
                                    null
                                ],
                                "detach": [
                                    null
                                ]
                            },
                            "tunnel": {
                                "start": [
                                    null
                                ],
                                "stop": [
                                    null
                                ]
                            },
                            "too-many-sas": [
                                null
                            ]
                        },
                        "ipsla": [
                            null
                        ],
                        "isis": [
                            null
                        ],
                        "license": {},
                        "mac-notification": {
                            "change": [
                                null
                            ],
                            "move": [
                                null
                            ],
                            "threshold": [
                                null
                            ]
                        },
                        "memory": {
                            "bufferpeak": [
                                null
                            ]
                        },
                        "mpls": {
                            "traffic-eng": {},
                            "fast-reroute": {
                                "protected": [
                                    null
                                ]
                            },
                            "rfc": {
                                "ldp": [
                                    null
                                ]
                            },
                            "ldp": {},
                            "vpn": [
                                null
                            ]
                        },
                        "msdp": [
                            null
                        ],
                        "mvpn": [
                            null
                        ],
                        "nhrp": {
                            "nhs": {},
                            "nhc": {},
                            "nhp": {},
                            "quota-exceeded": {}
                        },
                        "pim": {
                            "invalid-pim-message": [
                                null
                            ],
                            "neighbor-change": [
                                null
                            ],
                            "rp-mapping-change": [
                                null
                            ]
                        },
                        "port-security": {},
                        "power-ethernet": {
                            "police": [
                                null
                            ]
                        },
                        "pw": {
                            "vc": [
                                null
                            ]
                        },
                        "rep": [
                            null
                        ],
                        "rf": [
                            null
                        ],
                        "smart-licenseing": {
                            "smart-license": [
                                null
                            ]
                        },
                        "snmp": {
                            "authentication": [
                                null
                            ],
                            "coldstart": [
                                null
                            ],
                            "linkdown": [
                                null
                            ],
                            "linkup": [
                                null
                            ],
                            "warmstart": [
                                null
                            ]
                        },
                        "stackwise": {},
                        "stpx": {
                            "inconsistency": [
                                null
                            ],
                            "root-inconsistency": [
                                null
                            ],
                            "loop-inconsistency": [
                                null
                            ]
                        },
                        "syslog": [
                            null
                        ],
                        "transceiver": {
                            "all": [
                                null
                            ]
                        },
                        "tty": [
                            null
                        ],
                        "udld": {
                            "link-fail-rpt": [
                                null
                            ],
                            "status-change": [
                                null
                            ]
                        },
                        "vlancreate": [
                            null
                        ],
                        "vlandelete": [
                            null
                        ],
                        "vlan-membership": [
                            null
                        ],
                        "vrfmib": {
                            "vrf-up": [
                                null
                            ],
                            "vrf-down": [
                                null
                            ],
                            "vnet-trunk-up": [
                                null
                            ],
                            "vnet-trunk-down": [
                                null
                            ]
                        },
                        "vtp": [
                            null
                        ],
                        "Cisco-IOS-XE-ospf:ospf-config": {
                            "state-change": {
                                "enable": [
                                    null
                                ]
                            },
                            "errors": {
                                "enable": [
                                    null
                                ]
                            },
                            "retransmit": {
                                "enable": [
                                    null
                                ]
                            },
                            "lsa": {
                                "enable": [
                                    null
                                ]
                            },
                            "cisco-specific": {
                                "state-change": {
                                    "nssa-trans-change": [
                                        null
                                    ],
                                    "shamlink": {
                                        "interface": [
                                            null
                                        ],
                                        "neighbor": [
                                            null
                                        ]
                                    }
                                },
                                "errors": {
                                    "enable": [
                                        null
                                    ]
                                },
                                "retransmit": {
                                    "enable": [
                                        null
                                    ]
                                },
                                "lsa": {
                                    "enable": [
                                        null
                                    ]
                                }
                            }
                        },
                        "Cisco-IOS-XE-ospfv3:ospfv3-config": {
                            "state-change": {
                                "enable": [
                                    null
                                ]
                            },
                            "errors": {
                                "enable": [
                                    null
                                ]
                            }
                        }
                    }
                }
            },
            "Cisco-IOS-XE-snmp:engineID": {
                "local": "1234567890",
                "remote-conf": {
                    "host-list": [
                        {
                            "ip-address": "192.180.1.27",
                            "engine-id": "00000063000100A1C0B4011B"
                        }
                    ]
                }
            },
            "Cisco-IOS-XE-snmp:group": [
                {
                    "id": "authgroup",
                    "v3": {
                        "security-level-list": [
                            {
                                "security-level": "auth"
                            }
                        ]
                    }
                }
            ],
            "Cisco-IOS-XE-snmp:host-config": {
                "ip-community": [
                    {
                        "ip-address": "192.180.1.27",
                        "community-or-user": "authuser",
                        "informs": [
                            null
                        ],
                        "version": "3",
                        "security-level": "auth",
                        "trap-enable": {
                            "config": [
                                null
                            ]
                        }
                    },
                    {
                        "ip-address": "192.180.1.27",
                        "community-or-user": "public",
                        "version": "2c"
                    },
                    {
                        "ip-address": "192.180.1.33",
                        "community-or-user": "public"
                    },
                    {
                        "ip-address": "192.180.1.111",
                        "community-or-user": "public"
                    }
                ]
            },
            "Cisco-IOS-XE-snmp:location": "Building 3/Room 222"
        }
    }
    )
}

# Fetch the available SNMP configuration 
resource "iosxe_rest" "snmp_example_get" {
  method = "GET"
  path   = "/data/Cisco-IOS-XE-native:native/snmp-server"
}

# Remove the available SNMP configuration 
resource "iosxe_rest" "snmp_example_delete" {
  method = "DELETE"
  path = "/data/Cisco-IOS-XE-native:native/snmp-server"
}
