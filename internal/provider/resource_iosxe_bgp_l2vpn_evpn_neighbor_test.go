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

func TestAccIosxeBGPL2VPNEVPNNeighbor(t *testing.T) {
	if os.Getenv("C9000V") == "" {
		t.Skip("skipping test, set environment variable C9000V")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_l2vpn_evpn_neighbor.test", "ip", "3.3.3.3"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_l2vpn_evpn_neighbor.test", "activate", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_l2vpn_evpn_neighbor.test", "send_community", "both"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_l2vpn_evpn_neighbor.test", "route_reflector_client", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_l2vpn_evpn_neighbor.test", "soft_reconfiguration", "inbound"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeBGPL2VPNEVPNNeighborPrerequisitesConfig + testAccIosxeBGPL2VPNEVPNNeighborConfig_minimum(),
			},
			{
				Config: testAccIosxeBGPL2VPNEVPNNeighborPrerequisitesConfig + testAccIosxeBGPL2VPNEVPNNeighborConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:      "iosxe_bgp_l2vpn_evpn_neighbor.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000/address-family/no-vrf/l2vpn=evpn/l2vpn-evpn/neighbor=3.3.3.3",
				Check:             resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

const testAccIosxeBGPL2VPNEVPNNeighborPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000"
	attributes = {
		"id" = "65000"
	}
}

resource "iosxe_restconf" "PreReq1" {
	path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000/address-family/no-vrf/l2vpn=evpn"
	attributes = {
		"af-name" = "evpn"
	}
	depends_on = [iosxe_restconf.PreReq0, ]
}

resource "iosxe_restconf" "PreReq2" {
	path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000/neighbor=3.3.3.3"
	attributes = {
		"id" = "3.3.3.3"
		"remote-as" = "65000"
	}
	depends_on = [iosxe_restconf.PreReq0, ]
}

`

func testAccIosxeBGPL2VPNEVPNNeighborConfig_minimum() string {
	config := `resource "iosxe_bgp_l2vpn_evpn_neighbor" "test" {` + "\n"
	config += `	asn = "65000"` + "\n"
	config += `	ip = "3.3.3.3"` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, iosxe_restconf.PreReq2, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeBGPL2VPNEVPNNeighborConfig_all() string {
	config := `resource "iosxe_bgp_l2vpn_evpn_neighbor" "test" {` + "\n"
	config += `	asn = "65000"` + "\n"
	config += `	ip = "3.3.3.3"` + "\n"
	config += `	activate = true` + "\n"
	config += `	send_community = "both"` + "\n"
	config += `	route_reflector_client = false` + "\n"
	config += `	soft_reconfiguration = "inbound"` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, iosxe_restconf.PreReq2, ]` + "\n"
	config += `}` + "\n"
	return config
}
