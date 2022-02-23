terraform {
  required_providers {
    iosxe = {
      source = "local.plugin/ciscodevnet/iosxe"
    }
  }
}

provider "iosxe" {
  host = "https://128.107.251.88"
  insecure = true
  device_username = "netadmin"
  device_password = "C1sc0dna"
}

# Configure a Crypto IPsec tunnel
resource "iosxe_rest" "crypto_example_post" {
  method = "PATCH"
  path = "/data/Cisco-IOS-XE-native:native/crypto"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-native:crypto": {
            "Cisco-IOS-XE-crypto:ikev2": {
                "keyring": [
                    {
                    "name": "aws_tgw_bgp_2_backup",
                    "peer": [
                            {
                                "name": "aws_tgw_bgp_2_backup",
                                "address": {
                                    "ipv4": {
                                        "ipv4-address": "0.0.0.0",
                                        "ipv4-mask": "0.0.0.0"
                                    }
                                },
                                "pre-shared-key": {
                                    "key": "uNZptlnyDbRUFZxXRBImilyYouoDmLVb"
                                }
                            }
                        ]
                    }
                ],
                "policy": [
                    {
                        "name": "aws_tgw_bgp_2_backup",
                        "match": {
                            "fvrf": {
                                "any": [null]
                            }
                        },
                        "proposal": [
                            {
                                "proposals": "aws_tgw_bgp_2_backup"
                            }
                        ]
                    }
                ],
                "profile": [
                    {
                        "name": "aws_tgw_bgp_2_backup",
                        "authentication": {
                            "local": {
                                "pre-share": {
                                }
                            },
                            "remote": {
                                "pre-share": {
                                    }
                            }
                        },
                        "config-exchange": {
                            "request-1": false
                        },
                        "dpd": {
                            "interval": 10,
                            "retry": 2,
                            "query": "periodic"
                        },
                        "identity": {
                            "local": {
                                "address": "128.107.251.88"
                            }
                        },
                        "keyring": {
                            "local": {
                                "name": "aws_tgw_bgp_2_backup"
                            }
                        },
                        "match": {
                            "identity": {
                                "remote": {
                                    "address": {
                                        "ipv4": [
                                            {
                                                "ipv4-address": "52.52.2.74",
                                                "ipv4-mask": "255.255.255.0"
                                            }
                                        ]
                                    }
                                }
                            }
                        }
                    }
                ],
                "proposal": [
                    {
                        "name": "aws_tgw_bgp_2_backup",
                        "encryption": {
                            "aes-cbc-256": [null]
                        },
                        "group": {
                            "fourteen": [null],
                            "nineteen": [null],
                            "twenty": [null]
                        },
                        "integrity": {
                            "sha1": [null]
                        }
                    }
                ]
            },
            "Cisco-IOS-XE-crypto:ipsec": {
                "transform-set": [
                    {
                        "tag": "aws_tgw_bgp_2_backup",
                        "esp": "esp-aes",
                        "esp-hmac": "esp-sha-hmac",
                        "mode": {
                            "tunnel-choice": [null]
                        }
                    }
                ],
                "profile": [
                    {
                        "name": "aws_tgw_bgp_2_backup",
                        "set": {
                            "ikev2-profile": "aws_tgw_bgp_2_backup"
                        }
                    }
                ]
            }
        }
    }
  )
}


# Create Tunnel 303
resource "iosxe_rest" "tunnel_example_post" {
  method = "POST"
  path = "/data/Cisco-IOS-XE-native:native/interface"
  payload = jsonencode(
    {
        "Cisco-IOS-XE-native:Tunnel": {
            "name": 303,
            "description": "##Tunnel to AWS TGW##",
            "ip": {
                "address": {
                    "primary": {
                    "address": "169.254.26.254",
                    "mask": "255.255.255.252"
                    }
                }
            },
            "Cisco-IOS-XE-tunnel:tunnel": {
                "source": "Vlan2",
                "destination-config": {
                    "ipv4": "52.52.2.74"
                },
                "mode": {
                    "ipsec": {
                        "ipv4": {
                        }
                    }
                },
                "protection": {
                    "Cisco-IOS-XE-crypto:ipsec": {
                        "profile-option": {
                            "name": "aws_tgw_bgp_2_backup"
                        }
                    }
                }
            }
        }
    }
  )
}