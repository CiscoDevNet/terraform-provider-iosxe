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

func TestAccIosxeEVPNInstance(t *testing.T) {
	if os.Getenv("C9000V") == "" {
		t.Skip("skipping test, set environment variable C9000V")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_evpn_instance.test", "evpn_instance_num", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_evpn_instance.test", "vlan_based_replication_type_ingress", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_evpn_instance.test", "vlan_based_replication_type_static", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_evpn_instance.test", "vlan_based_replication_type_p2mp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_evpn_instance.test", "vlan_based_replication_type_mp2mp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_evpn_instance.test", "vlan_based_encapsulation", "vxlan"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_evpn_instance.test", "vlan_based_auto_route_target", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_evpn_instance.test", "vlan_based_rd", "10:10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_evpn_instance.test", "vlan_based_route_target", "10:10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_evpn_instance.test", "vlan_based_route_target_both", "10:10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_evpn_instance.test", "vlan_based_route_target_import", "10:10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_evpn_instance.test", "vlan_based_route_target_export", "10:10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_evpn_instance.test", "vlan_based_ip_local_learning_disable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_evpn_instance.test", "vlan_based_ip_local_learning_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_evpn_instance.test", "vlan_based_default_gateway_advertise", "enable"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_evpn_instance.test", "vlan_based_re_originate_route_type5", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeEVPNInstanceConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_evpn_instance.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/l2vpn/Cisco-IOS-XE-l2vpn:evpn_cont/evpn-instance/evpn/instance/instance=10",
			},
		},
	})
}

func testAccIosxeEVPNInstanceConfig_minimum() string {
	config := `resource "iosxe_evpn_instance" "test" {` + "\n"
	config += `	evpn_instance_num = 10` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeEVPNInstanceConfig_all() string {
	config := `resource "iosxe_evpn_instance" "test" {` + "\n"
	config += `	evpn_instance_num = 10` + "\n"
	config += `	vlan_based_replication_type_ingress = false` + "\n"
	config += `	vlan_based_replication_type_static = true` + "\n"
	config += `	vlan_based_replication_type_p2mp = false` + "\n"
	config += `	vlan_based_replication_type_mp2mp = false` + "\n"
	config += `	vlan_based_encapsulation = "vxlan"` + "\n"
	config += `	vlan_based_auto_route_target = false` + "\n"
	config += `	vlan_based_rd = "10:10"` + "\n"
	config += `	vlan_based_route_target = "10:10"` + "\n"
	config += `	vlan_based_route_target_both = "10:10"` + "\n"
	config += `	vlan_based_route_target_import = "10:10"` + "\n"
	config += `	vlan_based_route_target_export = "10:10"` + "\n"
	config += `	vlan_based_ip_local_learning_disable = false` + "\n"
	config += `	vlan_based_ip_local_learning_enable = true` + "\n"
	config += `	vlan_based_default_gateway_advertise = "enable"` + "\n"
	config += `	vlan_based_re_originate_route_type5 = true` + "\n"
	config += `}` + "\n"
	return config
}
