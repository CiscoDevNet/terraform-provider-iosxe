// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeInterfaceEthernet(t *testing.T) {
	if os.Getenv("C8000V") == "" {
		t.Skip("skipping test, set environment variable C8000V")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "bandwidth", "1000000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "description", "My Interface Description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ip_proxy_arp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ip_redirects", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ip_unreachables", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ipv4_address", "15.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ipv4_address_mask", "255.255.255.252"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ip_dhcp_relay_source_interface", "Loopback100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ip_access_group_in", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ip_access_group_in_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ip_access_group_out", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ip_access_group_out_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "helper_addresses.0.address", "10.10.10.10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "helper_addresses.0.global", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "helper_addresses.0.vrf", "VRF1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "source_template.0.template_name", "TEMP1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "source_template.0.merge", "false"))
	if os.Getenv("IOSXE179") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "bfd_template", "bfd_template1"))
	}
	if os.Getenv("IOSXE179") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "bfd_enable", "false"))
	}
	if os.Getenv("IOSXE179") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "bfd_local_address", "1.2.3.4"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ipv6_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ipv6_mtu", "1300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ipv6_nd_ra_suppress_all", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ipv6_address_dhcp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ipv6_link_local_addresses.0.address", "fe80::9656:d028:8652:66b6"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ipv6_link_local_addresses.0.link_local", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ipv6_addresses.0.prefix", "2001:DB8::/32"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ipv6_addresses.0.eui_64", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "arp_timeout", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "spanning_tree_link_type", "point-to-point"))
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ip_arp_inspection_trust", "true"))
	}
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ip_arp_inspection_limit_rate", "1000"))
	}
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ip_dhcp_snooping_trust", "true"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "negotiation_auto", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "service_policy_input", "POLICY1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "service_policy_output", "POLICY1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ip_flow_monitors.0.name", "MON1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "ip_flow_monitors.0.direction", "input"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "load_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "snmp_trap_link_status", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ethernet.test", "logging_event_link_status_enable", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeInterfaceEthernetPrerequisitesConfig + testAccDataSourceIosxeInterfaceEthernetConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

const testAccDataSourceIosxeInterfaceEthernetPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
	delete = false
	attributes = {
		"name" = "VRF1"
		"address-family/ipv4" = ""
	}
}

resource "iosxe_restconf" "PreReq1" {
	path = "Cisco-IOS-XE-native:native/policy/Cisco-IOS-XE-policy:policy-map=POLICY1"
	attributes = {
		"name" = "POLICY1"
	}
}

resource "iosxe_restconf" "PreReq2" {
	path = "Cisco-IOS-XE-native:native/flow/Cisco-IOS-XE-flow:record=REC1"
	attributes = {
		"name" = "REC1"
		"match/ipv4/source/address" = ""
		"collect/interface/output" = ""
	}
}

resource "iosxe_restconf" "PreReq3" {
	path = "Cisco-IOS-XE-native:native/flow/Cisco-IOS-XE-flow:monitor=MON1"
	attributes = {
		"name" = "MON1"
		"record/type" = "REC1"
	}
	depends_on = [iosxe_restconf.PreReq2, ]
}

`

func testAccDataSourceIosxeInterfaceEthernetConfig() string {
	config := `resource "iosxe_interface_ethernet" "test" {` + "\n"
	config += `	type = "GigabitEthernet"` + "\n"
	config += `	name = "3"` + "\n"
	config += `	bandwidth = 1000000` + "\n"
	config += `	description = "My Interface Description"` + "\n"
	config += `	shutdown = false` + "\n"
	config += `	ip_proxy_arp = false` + "\n"
	config += `	ip_redirects = false` + "\n"
	config += `	ip_unreachables = false` + "\n"
	config += `	ipv4_address = "15.1.1.1"` + "\n"
	config += `	ipv4_address_mask = "255.255.255.252"` + "\n"
	config += `	ip_dhcp_relay_source_interface = "Loopback100"` + "\n"
	config += `	ip_access_group_in = "1"` + "\n"
	config += `	ip_access_group_in_enable = true` + "\n"
	config += `	ip_access_group_out = "1"` + "\n"
	config += `	ip_access_group_out_enable = true` + "\n"
	config += `	helper_addresses = [{` + "\n"
	config += `		address = "10.10.10.10"` + "\n"
	config += `		global = false` + "\n"
	config += `		vrf = "VRF1"` + "\n"
	config += `	}]` + "\n"
	config += `	source_template = [{` + "\n"
	config += `		template_name = "TEMP1"` + "\n"
	config += `		merge = false` + "\n"
	config += `	}]` + "\n"
	if os.Getenv("IOSXE179") != "" {
		config += `	bfd_template = "bfd_template1"` + "\n"
	}
	if os.Getenv("IOSXE179") != "" {
		config += `	bfd_enable = false` + "\n"
	}
	if os.Getenv("IOSXE179") != "" {
		config += `	bfd_local_address = "1.2.3.4"` + "\n"
	}
	config += `	ipv6_enable = true` + "\n"
	config += `	ipv6_mtu = 1300` + "\n"
	config += `	ipv6_nd_ra_suppress_all = true` + "\n"
	config += `	ipv6_address_dhcp = true` + "\n"
	config += `	ipv6_link_local_addresses = [{` + "\n"
	config += `		address = "fe80::9656:d028:8652:66b6"` + "\n"
	config += `		link_local = true` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_addresses = [{` + "\n"
	config += `		prefix = "2001:DB8::/32"` + "\n"
	config += `		eui_64 = true` + "\n"
	config += `	}]` + "\n"
	config += `	arp_timeout = 300` + "\n"
	config += `	spanning_tree_link_type = "point-to-point"` + "\n"
	if os.Getenv("C9000V") != "" {
		config += `	ip_arp_inspection_trust = true` + "\n"
	}
	if os.Getenv("C9000V") != "" {
		config += `	ip_arp_inspection_limit_rate = 1000` + "\n"
	}
	if os.Getenv("C9000V") != "" {
		config += `	ip_dhcp_snooping_trust = true` + "\n"
	}
	config += `	negotiation_auto = false` + "\n"
	config += `	service_policy_input = "POLICY1"` + "\n"
	config += `	service_policy_output = "POLICY1"` + "\n"
	config += `	ip_flow_monitors = [{` + "\n"
	config += `		name = "MON1"` + "\n"
	config += `		direction = "input"` + "\n"
	config += `	}]` + "\n"
	config += `	load_interval = 30` + "\n"
	config += `	snmp_trap_link_status = true` + "\n"
	config += `	logging_event_link_status_enable = true` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, iosxe_restconf.PreReq2, iosxe_restconf.PreReq3, ]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxe_interface_ethernet" "test" {
			type = "GigabitEthernet"
			name = "3"
			depends_on = [iosxe_interface_ethernet.test]
		}
	`
	return config
}
