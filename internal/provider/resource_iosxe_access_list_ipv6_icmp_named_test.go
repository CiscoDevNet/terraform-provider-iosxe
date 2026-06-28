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

func TestAccIosxeAccessListIPv6IcmpNamed(t *testing.T) {
	initialChecks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "name", "TFACC_IPV6_ICMP_NAMED"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.0.sequence", "10"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.0.remark", "ipv6 icmp named proof"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.1.sequence", "20"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.1.ace_rule_action", "permit"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.1.ace_rule_protocol", "icmp"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.1.source_any", "true"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.1.destination_any", "true"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.1.icmp_named_msg_type", "packet-too-big"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.2.sequence", "30"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.2.ace_rule_action", "permit"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.2.ace_rule_protocol", "tcp"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.2.source_prefix", "2001:db8:64::/64"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.2.destination_any", "true"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.2.destination_port_equal", "22"),
	}
	updatedChecks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.0.remark", "updated ipv6 choice proof"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.1.ace_rule_protocol", "icmp"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.1.icmp_msg_type", "128"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.1.icmp_msg_code", "0"),
		resource.TestCheckNoResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.1.icmp_named_msg_type"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.2.ace_rule_protocol", "udp"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.2.source_host", "2001:db8::10"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.2.source_port_range_from", "1000"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.2.source_port_range_to", "2000"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.2.destination_prefix", "2001:db8:100::/64"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.2.destination_port_range_from", "3000"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.2.destination_port_range_to", "4000"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.2.dscp", "af11"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.2.log", "true"),
	}
	serviceObjectGroupChecks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "name", "TFACC_IPV6_ICMP_NAMED"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.#", "1"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.0.sequence", "10"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.0.ace_rule_action", "permit"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.0.ace_rule_protocol", "object-group"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.0.service_object_group", "TFACC_IPV6_SERVICE"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.0.source_any", "true"),
		resource.TestCheckResourceAttr("iosxe_access_list_ipv6.named_icmp", "entries.0.destination_any", "true"),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeAccessListIPv6IcmpNamedConfig(),
				Check:  resource.ComposeTestCheckFunc(initialChecks...),
			},
			{
				Config: testAccIosxeAccessListIPv6IcmpNamedUpdatedConfig(),
				Check:  resource.ComposeTestCheckFunc(updatedChecks...),
			},
			{
				Config: testAccIosxeAccessListIPv6ServiceObjectGroupConfig(),
				Check:  resource.ComposeTestCheckFunc(serviceObjectGroupChecks...),
			},
			{
				ResourceName:      "iosxe_access_list_ipv6.named_icmp",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     "TFACC_IPV6_ICMP_NAMED",
				ImportStateVerifyIgnore: []string{
					"entries.0.ack",
					"entries.0.established",
					"entries.0.fin",
					"entries.0.fragments",
					"entries.0.log",
					"entries.0.log_input",
					"entries.0.psh",
					"entries.0.rst",
					"entries.0.syn",
					"entries.0.urg",
				},
				Check: resource.ComposeTestCheckFunc(serviceObjectGroupChecks...),
			},
		},
	})
}

func testAccIosxeAccessListIPv6IcmpNamedConfig() string {
	return testAccIosxeAccessListIPv6ServiceObjectGroupPrerequisite + `resource "iosxe_access_list_ipv6" "named_icmp" {
  name = "TFACC_IPV6_ICMP_NAMED"
  depends_on = [iosxe_yang.ipv6_service_group]
  entries = [
    {
      sequence = 10
      remark = "ipv6 icmp named proof"
    },
    {
      sequence = 20
      ace_rule_action = "permit"
      ace_rule_protocol = "icmp"
      source_any = true
      destination_any = true
      icmp_named_msg_type = "packet-too-big"
    },
    {
      sequence = 30
      ace_rule_action = "permit"
      ace_rule_protocol = "tcp"
      source_prefix = "2001:db8:64::/64"
      destination_any = true
      destination_port_equal = "22"
    },
  ]
}
`
}

func testAccIosxeAccessListIPv6IcmpNamedUpdatedConfig() string {
	return testAccIosxeAccessListIPv6ServiceObjectGroupPrerequisite + `resource "iosxe_access_list_ipv6" "named_icmp" {
  name = "TFACC_IPV6_ICMP_NAMED"
  depends_on = [iosxe_yang.ipv6_service_group]
  entries = [
    {
      sequence = 10
      remark = "updated ipv6 choice proof"
    },
    {
      sequence = 20
      ace_rule_action = "permit"
      ace_rule_protocol = "icmp"
      source_any = true
      destination_any = true
      icmp_msg_type = 128
      icmp_msg_code = 0
    },
    {
      sequence = 30
      ace_rule_action = "permit"
      ace_rule_protocol = "udp"
      source_host = "2001:db8::10"
      source_port_range_from = "1000"
      source_port_range_to = "2000"
      destination_prefix = "2001:db8:100::/64"
      destination_port_range_from = "3000"
      destination_port_range_to = "4000"
      dscp = "af11"
      log = true
    },
  ]
}
`
}

func testAccIosxeAccessListIPv6ServiceObjectGroupConfig() string {
	return testAccIosxeAccessListIPv6ServiceObjectGroupPrerequisite + `resource "iosxe_access_list_ipv6" "named_icmp" {
  name = "TFACC_IPV6_ICMP_NAMED"
  depends_on = [iosxe_yang.ipv6_service_group]
  entries = [
    {
      sequence = 10
      ace_rule_action = "permit"
      ace_rule_protocol = "object-group"
      service_object_group = "TFACC_IPV6_SERVICE"
      source_any = true
      destination_any = true
    },
  ]
}
`
}

const testAccIosxeAccessListIPv6ServiceObjectGroupPrerequisite = `
resource "iosxe_yang" "ipv6_service_group" {
  path = "/Cisco-IOS-XE-native:native/object-group/Cisco-IOS-XE-object-group:v6-service[name=TFACC_IPV6_SERVICE]"
  attributes = {
    "name" = "TFACC_IPV6_SERVICE"
    "ipv6" = ""
  }
}

`
