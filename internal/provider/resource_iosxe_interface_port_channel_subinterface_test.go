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

func TestAccIosxeInterfacePortChannelSubinterface(t *testing.T) {
	if os.Getenv("C8000V") == "" {
		t.Skip("skipping test, set environment variable C8000V")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "name", "10.666"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "description", "My Interface Description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ip_proxy_arp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ip_redirects", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ip_unreachables", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "vrf_forwarding", "VRF1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ipv4_address", "192.0.2.2"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ipv4_address_mask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "encapsulation_dot1q_vlan_id", "666"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ip_access_group_in", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ip_access_group_in_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ip_access_group_out", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ip_access_group_out_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "helper_addresses.0.address", "10.10.10.10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "helper_addresses.0.global", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "bfd_template", "bfd_template1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "bfd_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "bfd_local_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ipv6_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ipv6_mtu", "1300"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ipv6_nd_ra_suppress_all", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ipv6_address_dhcp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ipv6_link_local_addresses.0.address", "fe80::9656:d028:8652:66b8"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ipv6_link_local_addresses.0.link_local", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ipv6_addresses.0.prefix", "2003:DB8::/32"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ipv6_addresses.0.eui_64", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "arp_timeout", "2147"))
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ip_arp_inspection_trust", "true"))
	}
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_port_channel_subinterface.test", "ip_arp_inspection_limit_rate", "1000"))
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeInterfacePortChannelSubinterfacePrerequisitesConfig + testAccIosxeInterfacePortChannelSubinterfaceConfig_minimum(),
			},
			{
				Config: testAccIosxeInterfacePortChannelSubinterfacePrerequisitesConfig + testAccIosxeInterfacePortChannelSubinterfaceConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_interface_port_channel_subinterface.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/interface/Port-channel-subinterface/Port-channel=10.666",
			},
		},
	})
}

const testAccIosxeInterfacePortChannelSubinterfacePrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
	delete = false
	attributes = {
		"name" = "VRF1"
		"address-family/ipv4" = ""
	}
}

resource "iosxe_restconf" "PreReq1" {
	path = "Cisco-IOS-XE-native:native/interface/Port-channel=10"
	attributes = {
		"name" = "10"
	}
}

resource "iosxe_restconf" "PreReq2" {
	path = "Cisco-IOS-XE-native:native/interface/Port-channel-subinterface/Port-channel=10.666"
	attributes = {
		"name" = "10.666"
	}
	depends_on = [iosxe_restconf.PreReq1, ]
}

`

func testAccIosxeInterfacePortChannelSubinterfaceConfig_minimum() string {
	config := `resource "iosxe_interface_port_channel_subinterface" "test" {` + "\n"
	config += `	name = "10.666"` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, iosxe_restconf.PreReq2, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeInterfacePortChannelSubinterfaceConfig_all() string {
	config := `resource "iosxe_interface_port_channel_subinterface" "test" {` + "\n"
	config += `	name = "10.666"` + "\n"
	config += `	description = "My Interface Description"` + "\n"
	config += `	shutdown = false` + "\n"
	config += `	ip_proxy_arp = false` + "\n"
	config += `	ip_redirects = false` + "\n"
	config += `	ip_unreachables = false` + "\n"
	config += `	vrf_forwarding = "VRF1"` + "\n"
	config += `	ipv4_address = "192.0.2.2"` + "\n"
	config += `	ipv4_address_mask = "255.255.255.0"` + "\n"
	config += `	encapsulation_dot1q_vlan_id = 666` + "\n"
	config += `	ip_access_group_in = "1"` + "\n"
	config += `	ip_access_group_in_enable = true` + "\n"
	config += `	ip_access_group_out = "1"` + "\n"
	config += `	ip_access_group_out_enable = true` + "\n"
	config += `	helper_addresses = [{` + "\n"
	config += `		address = "10.10.10.10"` + "\n"
	config += `		global = false` + "\n"
	config += `	}]` + "\n"
	config += `	bfd_template = "bfd_template1"` + "\n"
	config += `	bfd_enable = true` + "\n"
	config += `	bfd_local_address = "1.2.3.4"` + "\n"
	config += `	ipv6_enable = true` + "\n"
	config += `	ipv6_mtu = 1300` + "\n"
	config += `	ipv6_nd_ra_suppress_all = true` + "\n"
	config += `	ipv6_address_dhcp = true` + "\n"
	config += `	ipv6_link_local_addresses = [{` + "\n"
	config += `		address = "fe80::9656:d028:8652:66b8"` + "\n"
	config += `		link_local = true` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_addresses = [{` + "\n"
	config += `		prefix = "2003:DB8::/32"` + "\n"
	config += `		eui_64 = true` + "\n"
	config += `	}]` + "\n"
	config += `	arp_timeout = 2147` + "\n"
	if os.Getenv("C9000V") != "" {
		config += `	ip_arp_inspection_trust = true` + "\n"
	}
	if os.Getenv("C9000V") != "" {
		config += `	ip_arp_inspection_limit_rate = 1000` + "\n"
	}
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, iosxe_restconf.PreReq2, ]` + "\n"
	config += `}` + "\n"
	return config
}
