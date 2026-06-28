// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeAccessListIPv6IPv4Parity(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeAccessListIPv6IPv4ParityConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_access_list_ipv6.parity", "name", "TFACC_IPV6_PARITY"),
					resource.TestCheckResourceAttr("iosxe_access_list_ipv6.parity", "entries.#", "3"),
					resource.TestCheckResourceAttr("iosxe_access_list_ipv6.parity", "entries.0.source_fqdn_group", "TFACC_IPV6_FQDN"),
					resource.TestCheckResourceAttr("iosxe_access_list_ipv6.parity", "entries.0.source_port_equal", "1000"),
					resource.TestCheckResourceAttr("iosxe_access_list_ipv6.parity", "entries.0.source_port_equal_2", "1002"),
					resource.TestCheckResourceAttr("iosxe_access_list_ipv6.parity", "entries.0.source_port_equal_10", "1010"),
					resource.TestCheckResourceAttr("iosxe_access_list_ipv6.parity", "entries.0.ack", "true"),
					resource.TestCheckResourceAttr("iosxe_access_list_ipv6.parity", "entries.0.fin", "true"),
					resource.TestCheckResourceAttr("iosxe_access_list_ipv6.parity", "entries.0.psh", "true"),
					resource.TestCheckResourceAttr("iosxe_access_list_ipv6.parity", "entries.1.destination_fqdn_group", "TFACC_IPV6_FQDN"),
					resource.TestCheckResourceAttr("iosxe_access_list_ipv6.parity", "entries.1.destination_port_equal", "2000"),
					resource.TestCheckResourceAttr("iosxe_access_list_ipv6.parity", "entries.1.destination_port_equal_2", "2002"),
					resource.TestCheckResourceAttr("iosxe_access_list_ipv6.parity", "entries.1.destination_port_equal_10", "2010"),
					resource.TestCheckResourceAttr("iosxe_access_list_ipv6.parity", "entries.1.rst", "true"),
					resource.TestCheckResourceAttr("iosxe_access_list_ipv6.parity", "entries.1.syn", "true"),
					resource.TestCheckResourceAttr("iosxe_access_list_ipv6.parity", "entries.1.urg", "true"),
					resource.TestCheckResourceAttr("iosxe_access_list_ipv6.parity", "entries.2.established", "true"),
				),
			},
			{
				ResourceName:      "iosxe_access_list_ipv6.parity",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     "TFACC_IPV6_PARITY",
				ImportStateVerifyIgnore: []string{
					"entries.0.established",
					"entries.0.fragments",
					"entries.0.log",
					"entries.0.log_input",
					"entries.0.rst",
					"entries.0.source_any",
					"entries.0.syn",
					"entries.0.urg",
					"entries.1.ack",
					"entries.1.destination_any",
					"entries.1.established",
					"entries.1.fin",
					"entries.1.fragments",
					"entries.1.log",
					"entries.1.log_input",
					"entries.1.psh",
					"entries.2.ack",
					"entries.2.fin",
					"entries.2.fragments",
					"entries.2.log",
					"entries.2.log_input",
					"entries.2.psh",
					"entries.2.rst",
					"entries.2.syn",
					"entries.2.urg",
				},
			},
		},
	})
}

func testAccIosxeAccessListIPv6IPv4ParityConfig() string {
	return `resource "iosxe_yang" "fqdn_group" {
  path = "/Cisco-IOS-XE-native:native/object-group/Cisco-IOS-XE-object-group:fqdn[name=TFACC_IPV6_FQDN]"
  attributes = {
    "name" = "TFACC_IPV6_FQDN"
  }
  lists = [
    {
      name = "pattern"
      key = "fqdn-pattern"
      items = [
        {
          "fqdn-pattern" = ".*\\.example\\.com"
        },
      ]
    },
  ]
}

resource "iosxe_access_list_ipv6" "parity" {
  name = "TFACC_IPV6_PARITY"
  entries = [
    {
      sequence = 10
      ace_rule_action = "permit"
      ace_rule_protocol = "tcp"
      source_fqdn_group = "TFACC_IPV6_FQDN"
      source_port_equal = "1000"
      source_port_equal_2 = "1002"
      source_port_equal_3 = "1003"
      source_port_equal_4 = "1004"
      source_port_equal_5 = "1005"
      source_port_equal_6 = "1006"
      source_port_equal_7 = "1007"
      source_port_equal_8 = "1008"
      source_port_equal_9 = "1009"
      source_port_equal_10 = "1010"
      destination_any = true
      ack = true
      fin = true
      psh = true
    },
    {
      sequence = 20
      ace_rule_action = "permit"
      ace_rule_protocol = "tcp"
      source_any = true
      destination_fqdn_group = "TFACC_IPV6_FQDN"
      destination_port_equal = "2000"
      destination_port_equal_2 = "2002"
      destination_port_equal_3 = "2003"
      destination_port_equal_4 = "2004"
      destination_port_equal_5 = "2005"
      destination_port_equal_6 = "2006"
      destination_port_equal_7 = "2007"
      destination_port_equal_8 = "2008"
      destination_port_equal_9 = "2009"
      destination_port_equal_10 = "2010"
      rst = true
      syn = true
      urg = true
    },
    {
      sequence = 30
      ace_rule_action = "permit"
      ace_rule_protocol = "tcp"
      source_any = true
      destination_any = true
      established = true
    },
  ]
  depends_on = [iosxe_yang.fqdn_group]
}
`
}
