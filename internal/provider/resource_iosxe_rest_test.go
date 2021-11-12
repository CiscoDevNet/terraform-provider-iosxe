package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// For three aaa i.e. aaa-accounting, aaa-authentication, and aaa-authorization, it would be better to not touch them :). Change it if you know what you are doing! (change the 't' in function's name to 'T' to run)
//
// Both POE and NAT needs a different hardware to work. Configure the environment variables accordingly! (change the 't' in function's name to 'T' to run)
//
// Dev-vrf needs to be created before executing emp feature test

var providerIOSXE *schema.Provider

func testAccResource_aaa_accounting_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_aaa_accounting_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_aaa_accounting_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}

func testAccResource_aaa_authentication_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_aaa_authentication_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_aaa_authentication_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}

func testAccResource_aaa_authorization_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_aaa_authorization_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_aaa_authorization_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}

func TestAccResource_acl_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_acl_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_acl_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}

func TestAccResource_bgp_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_bgp_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_bgp_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}

func TestAccResource_cdp_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_cdp_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_cdp_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}

func TestAccResource_dhcp_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_dhcp_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_dhcp_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}

func TestAccResource_emp_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_emp_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_emp_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}

func TestAccResource_etherChannel_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_etherChannel_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_etherChannel_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}

func TestAccResource_hsrp_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_hsrp_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_hsrp_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}

func TestAccResource_igmp_proxy_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_igmp_proxy_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_igmp_proxy_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}

func TestAccResource_igmp_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_igmp_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_igmp_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}

func TestAccResource_l3_subinterface_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_l3_subinterface_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_l3_subinterface_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}

func TestAccResource_line_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_line_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_line_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}

func TestAccResource_mdt_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_mdt_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_mdt_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}

func testAccResource_nat_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_nat_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_nat_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}

func TestAccResource_ntp_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_ntp_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_ntp_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}

func TestAccResource_ospf_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_ospf_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_ospf_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}

func TestAccResource_pim_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_pim_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_pim_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}

func testAccResource_poe_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_poe_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_poe_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}

func TestAccResource_radius_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_radius_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_radius_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}

func TestAccResource_snmp_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_snmp_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_snmp_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}

func TestAccResource_span_rspan_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_span_rspan_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_span_rspan_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}
func TestAccResource_vlan_trunk_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_vlan_trunk_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_vlan_trunk_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}
func TestAccResource_vlan_voice_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_vlan_voice_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_vlan_voice_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}
func TestAccResource_vlan_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_vlan_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_vlan_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}
func TestAccResource_vtp_IOSXERest(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIOSXE),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIOSXE_vtp_RestConfig_Create,
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccResourceIOSXE_vtp_RestConfig_Delete,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
		CheckDestroy: testAccCheckIOSXEGeneralizeDestroy(providerIOSXE),
	})
}

const testAccResourceIOSXE_aaa_accounting_RestConfig_Create = `
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
`

const testAccResourceIOSXE_aaa_accounting_RestConfig_Delete = `
	resource "iosxe_rest" "aaa_accounting_example_delete" {
		method = "DELETE"
		path = "/data/Cisco-IOS-XE-native:native/aaa/accounting"
	}
`

const testAccResourceIOSXE_aaa_authentication_RestConfig_Create = `
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
`

const testAccResourceIOSXE_aaa_authentication_RestConfig_Delete = `
	resource "iosxe_rest" "aaa_authentication_example_delete" {
		method = "DELETE"
		path = "/data/Cisco-IOS-XE-native:native/aaa/authentication"
	}
`

const testAccResourceIOSXE_aaa_authorization_RestConfig_Create = `
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
							"local": [
								null
							]
						}
					},
					{
						"name": "patch-auth",
						"a1": {
							"local": [
								null
							]
						}
					}
				],
				"network": [
					{
						"id": "default",
						"a1": {
							"local": [
								null
							]
						}
					},
					{
						"id": "patch-auth",
						"a1": {
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
`

const testAccResourceIOSXE_aaa_authorization_RestConfig_Delete = `
	resource "iosxe_rest" "aaa_authorization_example_delete" {
		method = "DELETE"
		path = "/data/Cisco-IOS-XE-native:native/aaa/authorization/exec=put-auth"
	}
`

const testAccResourceIOSXE_acl_RestConfig_Create = `
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
					}
				]
			}
			}
		)
	}
`
const testAccResourceIOSXE_acl_RestConfig_Delete = `
	resource "iosxe_rest" "acl_example_delete" {
		method = "DELETE"
		path = "/data/Cisco-IOS-XE-native:native/ip/access-list/extended=101"
	}
`

const testAccResourceIOSXE_bgp_RestConfig_Create = `
	resource "iosxe_rest" "bgp_example_patch" {
		method = "PATCH"
		path = "/data/Cisco-IOS-XE-native:native/router/bgp"
		payload = jsonencode(
		{
			"Cisco-IOS-XE-bgp:bgp": [
				{
					"id": 46000,
					"neighbor": [
						{
							"id": "10.109.1.4",
							"remote-as": 66000
						}
					]
				}
			]
		}
		)
	}
`
const testAccResourceIOSXE_bgp_RestConfig_Delete = `
	resource "iosxe_rest" "bgp_example_delete" {
		method = "DELETE"
		path = "/data/Cisco-IOS-XE-native:native/router/bgp=46000"
	}
`

const testAccResourceIOSXE_cdp_RestConfig_Create = `
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
`
const testAccResourceIOSXE_cdp_RestConfig_Delete = `
	resource "iosxe_rest" "cdp_example_delete" {
		method = "DELETE"
		path = "/data/Cisco-IOS-XE-native:native/cdp"
	}
`

const testAccResourceIOSXE_dhcp_RestConfig_Create = `
	resource "iosxe_rest" "dhcp_example_patch" {
		method = "PATCH"
		path   = "/data/Cisco-IOS-XE-native:native/ipv6/dhcp"
		payload = jsonencode(
		{
			"Cisco-IOS-XE-native:dhcp" : {
			"Cisco-IOS-XE-dhcp:pool" : [
				{
				"name" : "7",
				"link-address" : [
					{
					"address" : "2001:1002::/64"
					}
				]
				}
			]
			}
		}
		)
	}
`
const testAccResourceIOSXE_dhcp_RestConfig_Delete = `
	resource "iosxe_rest" "dhcp_example_delete" {
		method = "DELETE"
		path   = "/data/Cisco-IOS-XE-native:native/ipv6/dhcp/pool=7"
	}
`

const testAccResourceIOSXE_emp_RestConfig_Create = `
	resource "iosxe_rest" "emp_example_patch" {
		method = "PATCH"
		path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet"
		payload = jsonencode(
		{
			"Cisco-IOS-XE-native:GigabitEthernet": {
				"name": "1/0/1",
				"description": "AP1-new-desc",
				"switchport-conf": {
					"switchport": false
				},
				"vrf": {
					"forwarding": "Dev-vrf"
				}
			}
		}
		)
	}
`

const testAccResourceIOSXE_emp_RestConfig_Delete = `
	resource "iosxe_rest" "emp_example_delete" {
		method = "DELETE"
		path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f1/switchport-conf"
	}
`

const testAccResourceIOSXE_etherChannel_RestConfig_Create = `
	resource "iosxe_rest" "etherChannel_example_patch" {
		method = "PATCH"
		path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f19"
		payload = jsonencode(
		{
			"Cisco-IOS-XE-native:GigabitEthernet" : {
			"name" : "1/0/19",
			"Cisco-IOS-XE-ethernet:channel-group" : {
				"number" : 5,
				"mode" : "desirable"
			},
			"Cisco-IOS-XE-ethernet:lacp" : {
				"port-priority" : 34000,
				"rate" : "fast"
			}
			}
		}
		)
	}
`
const testAccResourceIOSXE_etherChannel_RestConfig_Delete = `
	resource "iosxe_rest" "etherChannel_example_delete" {
		method = "DELETE"
		path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f19/Cisco-IOS-XE-ethernet:channel-group"
	}
`

const testAccResourceIOSXE_hsrp_RestConfig_Create = `
	resource "iosxe_rest" "hsrp_example_patch" {
		method = "PATCH"
		path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet"
		payload = jsonencode(
		{
			"Cisco-IOS-XE-native:GigabitEthernet" : {
				"name" : "1/0/13",
				"switchport-conf" : {
				"switchport" : false
				},
				"ip" : {
					"address" : {
						"primary" : {
						"address" : "10.0.0.1",
						"mask" : "255.255.255.0"
						}
					}
				},
				"standby" : {
					"standby-list" : [
						{
						"group-number" : 1,
						"ip" : {
							"address" : "10.0.0.3"
						},
						"priority" : 110
						},
						{
						"group-number" : 2,
						"ip" : {
							"address" : "10.0.0.4"
						},
						"priority" : 115
						}
					]
				}
			}
		}
		)
	}
`
const testAccResourceIOSXE_hsrp_RestConfig_Delete = `
	resource "iosxe_rest" "hsrp_example_delete" {
		method = "DELETE"
		path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f13/standby"
	}
`

const testAccResourceIOSXE_igmp_proxy_RestConfig_Create = `
	resource "iosxe_rest" "igmp_proxy_example_patch" {
		method = "PATCH"
		path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet"
		payload = jsonencode(
		{
			"Cisco-IOS-XE-native:GigabitEthernet": {
				"name": "1/0/18",
				"switchport-conf": {
					"switchport": false
				},
				"ip": {
					"Cisco-IOS-XE-igmp:igmp": {
						"mroute-proxy": {
							"Loopback": 0
						},
						"unidirectional-link-leaf": [
							null
						],
						"upstream-proxy": "proxymap",
						"iif-starg": [
							null
						],
						"proxy-report-interval": 120
					}
				}
			}
		}
		)
	}
`

const testAccResourceIOSXE_igmp_proxy_RestConfig_Delete = `
	resource "iosxe_rest" "igmp_proxy_example_delete" {
		method = "DELETE"
		path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f18/ip/igmp"
	}
`

const testAccResourceIOSXE_igmp_RestConfig_Create = `
	resource "iosxe_rest" "igmp_example_patch" {
		method = "PATCH"
		path   = "/data/Cisco-IOS-XE-native:native/ip/igmp"
		payload = jsonencode(
		{
			"Cisco-IOS-XE-igmp:igmp" : {
			"profile" : [
				{
				"id" : "3",
				"action" : "permit",
				"range" : [
					{
					"low_ip" : "229.9.9.1"
					},
					{
					"low_ip" : "229.9.9.2"
					}
				]
				}
			]
			}
		}
		)
	}
	
`

const testAccResourceIOSXE_igmp_RestConfig_Delete = `
	resource "iosxe_rest" "igmp_example_delete" {
		method = "DELETE"
		path   = "/data/Cisco-IOS-XE-native:native/ip/igmp/profile=3"
	}
`

const testAccResourceIOSXE_l3_subinterface_RestConfig_Create = `
	resource "iosxe_rest" "l3_subinterface_example_patch" {
		method = "PATCH"
		path   = "/data/Cisco-IOS-XE-native:native/interface/Port-channel-subinterface/Port-channel"
		payload = jsonencode(
		{
			"Port-channel": {
				"name": "2.10",
				"encapsulation": {
					"dot1Q": {
						"vlan-id": 10
					}
				},
				"ip": {
					"address": {
						"primary": {
							"address": "10.10.10.10",
							"mask": "255.255.255.0"
						}
					}
				}
			}
		}
		)
	}
`

const testAccResourceIOSXE_l3_subinterface_RestConfig_Delete = `
	resource "iosxe_rest" "l3_subinterface_example_delete" {
		method = "DELETE"
		path = "/data/Cisco-IOS-XE-native:native/interface/Port-channel-subinterface/Port-channel=2.10"
	}
`

const testAccResourceIOSXE_line_RestConfig_Create = `
	resource "iosxe_rest" "line_example_patch" {
		method = "PATCH"
		path   = "/data/Cisco-IOS-XE-native:native/line"
		payload = jsonencode(
		{
			"Cisco-IOS-XE-native:line": {
			"auto-consolidation": false,
			"console": [
				{
					"first": "0",
					"exec-timeout": {
						"minutes": 0,
						"seconds": 0
					},
					"login": {
						"authentication": "NOAUTH"
					},
					"stopbits": "1"
				}
			],
			"vty": [
				{
					"first": 0,
					"exec-timeout": {
						"minutes": 0,
						"seconds": 0
					},
					"length": 0,
					"password": {
						"secret": "login"
					},
					"transport": {
						"input": {
							"all": [
								null
							]
						}
					}
				},
				{
					"first": 1,
					"last": 4,
					"exec-timeout": {
						"minutes": 0,
						"seconds": 0
					},
					"transport": {
						"input": {
							"all": [
								null
							]
						}
					}
				},
				{
					"first": 5,
					"last": 31,
					"exec-timeout": {
						"minutes": 0,
						"seconds": 0
					},
					"transport": {
						"input": {
							"all": [
								null
							]
						}
					}
				},
				{
					"first": 32,
					"exec-timeout": {
						"minutes": 0,
						"seconds": 0
					},
					"transport": {
						"input": {
							"all": [
								null
							]
						}
					}
				}
			]
		}
		}
		)
	}
`

const testAccResourceIOSXE_line_RestConfig_Delete = `
	resource "iosxe_rest" "line_example_delete" {
		method = "DELETE"
		path = "/data/Cisco-IOS-XE-native:native/line/vty=32"
	}
`

const testAccResourceIOSXE_mdt_RestConfig_Create = `
	resource "iosxe_rest" "mdt_example_patch" {
		method = "PATCH"
		path   = "/data/Cisco-IOS-XE-mdt-cfg:mdt-config-data"
		payload = jsonencode(
		{
			"Cisco-IOS-XE-mdt-cfg:mdt-config-data" : {
			"mdt-subscription" : [
				{
				"subscription-id" : 1200,
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
`

const testAccResourceIOSXE_mdt_RestConfig_Delete = `
	resource "iosxe_rest" "mdt_example_delete" {
		method = "DELETE"
		path   = "/data/Cisco-IOS-XE-mdt-cfg:mdt-config-data/mdt-subscription=1200"
	}
`

const testAccResourceIOSXE_nat_RestConfig_Create = `
	resource "iosxe_rest" "nat_example_patch" {
		method = "PATCH"
		path   = "/data/Cisco-IOS-XE-native:native/ip/nat"
		payload = jsonencode(
		{
			"Cisco-IOS-XE-nat:nat": {
				"pool": [
					{
						"id": "net-208",
						"prefix-length": 27
					}
				],
				"inside": {
					"source": {
						"list": [
							{
								"id": 1,
								"pool": "net-208"
							}
						]
					}
				}
			}
		}
		)
	}
`
const testAccResourceIOSXE_nat_RestConfig_Delete = `
	resource "iosxe_rest" "nat_example_delete" {
		method = "DELETE"
		path = "/data/Cisco-IOS-XE-native:native/ip/nat/inside"
	}
`

const testAccResourceIOSXE_ntp_RestConfig_Create = `

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
`
const testAccResourceIOSXE_ntp_RestConfig_Delete = `

	resource "iosxe_rest" "ntp_example_delete" {
		method = "DELETE"
		path = "/data/Cisco-IOS-XE-native:native/ntp"
	}
`

const testAccResourceIOSXE_ospf_RestConfig_Create = `
	resource "iosxe_rest" "ospf_example_patch" {
		method = "PATCH"
		path = "/data/Cisco-IOS-XE-native:native/router/router-ospf/ospf"
		payload = jsonencode(
		{
			"ospf": {
				"process-id": [
					{
						"id": 15,
						"network": [
							{
								"ip": "10.1.1.1",
								"wildcard": "255.240.0.0",
								"area": 20
							}
						],
						"nsf": {
							"nsf-cisco": {
								"enforce": {
									"global": [
										null
									]
								}
							}
						}
					}
				]
			}
		}
		)
	}
`
const testAccResourceIOSXE_ospf_RestConfig_Delete = `
	resource "iosxe_rest" "ospf_example_delete" {
		method = "DELETE"
		path = "/data/Cisco-IOS-XE-native:native/router/router-ospf/ospf/process-id=15"
	}
`

const testAccResourceIOSXE_pim_RestConfig_Create = `
	resource "iosxe_rest" "pim_example_patch" {
		method = "PATCH"
		path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f15/ip/pim"
		payload = jsonencode(
		{
			"Cisco-IOS-XE-native:pim": {
				"Cisco-IOS-XE-multicast:pim-mode-choice-cfg": {
					"passive": [
						null
					]
				}
			}
		}
		)
	}
`
const testAccResourceIOSXE_pim_RestConfig_Delete = `
	resource "iosxe_rest" "pim_example_delete" {
		method = "DELETE"
		path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f15/ip/pim"
	} 
`

const testAccResourceIOSXE_poe_RestConfig_Create = `
	resource "iosxe_rest" "poe_example_patch" {
		method = "PATCH"
		path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f12/Cisco-IOS-XE-power:power/inline"
		payload = jsonencode(
		{
			"Cisco-IOS-XE-power:inline": {
			"auto-choice": {
				"max": 5000
			}
			}
		}
		)
	}
`
const testAccResourceIOSXE_poe_RestConfig_Delete = `
	resource "iosxe_rest" "poe_example_delete" {
		method = "DELETE"
		path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f12/Cisco-IOS-XE-power:power/inline"
	}
`

const testAccResourceIOSXE_radius_RestConfig_Create = `
	resource "iosxe_rest" "radius_example_patch" {
		method = "PATCH"
		path = "/data/Cisco-IOS-XE-native:native/radius"
		payload = jsonencode(
		{
			"Cisco-IOS-XE-native:radius": {
				"Cisco-IOS-XE-aaa:server": [
					{
					"id": "rsim",
					"address": {
						"ipv4": "124.2.2.12",
						"auth-port": 1612,
						"acct-port": 1813
					},
					"timeout": 60,
					"key": {
						"key": "rad123"
					},
					"retransmit": 10
					}
				]
			}
		}
		)
	}
`
const testAccResourceIOSXE_radius_RestConfig_Delete = `
	resource "iosxe_rest" "radius_example_delete" {
		method = "DELETE"
		path = "/data/Cisco-IOS-XE-native:native/radius/server=rsim"
	}
`

const testAccResourceIOSXE_snmp_RestConfig_Create = `
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
`

const testAccResourceIOSXE_snmp_RestConfig_Delete = `
	resource "iosxe_rest" "snmp_example_delete" {
		method = "DELETE"
		path = "/data/Cisco-IOS-XE-native:native/snmp-server"
	}
`

const testAccResourceIOSXE_span_rspan_RestConfig_Create = `
	resource "iosxe_rest" "span_rspan_example_patch" {
		method = "PATCH"
		path   = "/data/Cisco-IOS-XE-native:native/monitor"
		payload = jsonencode(
		{
			"Cisco-IOS-XE-native:monitor": {
				"session": [
					{
						"id": 1,
						"destination": {
							"interface": [
								{
									"name": "AppGigabitEthernet1/0/1"
								}
							]
						},
						"source": {
							"interface": [
								{
									"name": "GigabitEthernet1/0/1"
								}
							]
						}
					}
				]
			}
		}
		)
	}
`
const testAccResourceIOSXE_span_rspan_RestConfig_Delete = `
	resource "iosxe_rest" "span_rspan_example_delete" {
		method = "DELETE"
		path = "/data/Cisco-IOS-XE-native:native/monitor/session=1"
	}
`

const testAccResourceIOSXE_vlan_trunk_RestConfig_Create = `
	resource "iosxe_rest" "vlan_trunk_example_patch" {
		method = "PATCH"
		path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet"
		payload = jsonencode(
		{
			"Cisco-IOS-XE-native:GigabitEthernet": {
				"name": "1/0/11",
				"switchport-config": {
					"switchport": {
						"Cisco-IOS-XE-switch:trunk": {
							"allowed": {
								"vlan": {
									"vlans": "20,30,40"
								}
							},
							"native": {
								"vlan": {
									"vlan-id": 100
								}
							}
						}
					}
				}
			}
		}
		)
	}
`
const testAccResourceIOSXE_vlan_trunk_RestConfig_Delete = `
	resource "iosxe_rest" "vlan_trunk_example_delete" {
		method = "DELETE"
		path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f11/switchport-config"
	}
`

const testAccResourceIOSXE_vlan_voice_RestConfig_Create = `
	resource "iosxe_rest" "vlan_voice_example_patch" {
		method = "PATCH"
		path   = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet"
		payload = jsonencode(
		{
			"Cisco-IOS-XE-native:GigabitEthernet": {
				"name": "1/0/12",
				"description": "AP1 updated desc",
				"switchport-config": {
					"switchport": {
						"Cisco-IOS-XE-switch:voice": {
							"vlan": {
								"vlan": "dot1p"
							}
						}
					}
				}
			}
		}
		)
	}
`
const testAccResourceIOSXE_vlan_voice_RestConfig_Delete = `
	resource "iosxe_rest" "vlan_voice_example_delete" {
		method = "DELETE"
		path = "/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet=1%2f0%2f12/switchport-config"
	}
`

const testAccResourceIOSXE_vlan_RestConfig_Create = `
 
	resource "iosxe_rest" "vlan_example_patch" {
		method = "PATCH"
		path = "/data/Cisco-IOS-XE-native:native/vlan"
		payload = jsonencode(
		{
		"Cisco-IOS-XE-native:vlan": {
			"Cisco-IOS-XE-vlan:vlan-list": [
				{
					"id": 51,
					"name": "VLAN51-acc"
				}
			]
			}
		}
		)
	}
`

const testAccResourceIOSXE_vlan_RestConfig_Delete = `
	resource "iosxe_rest" "vlan_example_delete" {
		method = "DELETE"
		path = "/data/Cisco-IOS-XE-native:native/vlan/vlan-list=51"
	}
`

const testAccResourceIOSXE_vtp_RestConfig_Create = `
  resource "iosxe_rest" "vtp_example_patch" {
		method = "PATCH"
		path = "/data/Cisco-IOS-XE-native:native/vtp"
		payload = jsonencode(
			{
				"Cisco-IOS-XE-native:vtp": {
					"Cisco-IOS-XE-vtp:password": {
						"password": "some_password"
					},
					"Cisco-IOS-XE-vtp:domain": "domain_updated"
				}
			}
		)
	}
`
const testAccResourceIOSXE_vtp_RestConfig_Delete = `
	resource "iosxe_rest" "vtp_example_delete" {
		method = "DELETE"
		path = "/data/Cisco-IOS-XE-native:native/vtp/password"
	}
`
