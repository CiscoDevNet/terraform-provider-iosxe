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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeRouteMap(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "name", "RM1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.seq", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.operation", "permit"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.description", "Entry 10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.continue", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_interfaces.0", "Loopback1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_ip_address_access_lists.0", "ACL1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_ip_next_hop_access_lists.0", "ACL1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_ipv6_address_access_lists", "ACL1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_ipv6_next_hop_access_lists", "ACL1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_route_type_external", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_route_type_external_type_1", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_route_type_external_type_2", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_route_type_internal", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_route_type_level_1", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_route_type_level_2", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_route_type_local", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_source_protocol_bgp.0", "65000"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_source_protocol_connected", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_source_protocol_eigrp.0", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_source_protocol_isis", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_source_protocol_lisp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_source_protocol_ospf.0", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_source_protocol_ospfv3.0", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_source_protocol_rip", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_source_protocol_static", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_tags.0", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_track", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_as_paths_legacy.0", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_community_lists_legacy.0", "COMM1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_extcommunity_lists_legacy.0", "EXTCOMM1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.match_local_preferences_legacy.0", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_default_interfaces.0", "Loopback1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_global", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_interfaces.0", "Loopback1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_ip_address", "PFL1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_ip_default_global_next_hop_address.0", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_ip_default_next_hop_address.0", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_ip_global_next_hop_address.0", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_ip_next_hop_address.0", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_ip_qos_group", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_ipv6_address.0", "PFL2"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_ipv6_default_global_next_hop", "2001::1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_ipv6_default_next_hop.0", "2001::1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_ipv6_next_hop.0", "2001::1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_level_1", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_metric_value", "110"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_metric_delay", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_metric_reliability", "90"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_metric_loading", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_metric_mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_metric_type", "external"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_tag", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_as_path_prepend_as_legacy", "65001 65001"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_as_path_prepend_last_as_legacy", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_as_path_tag_legacy", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_communities_legacy.0", "1:2"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_communities_additive_legacy", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_community_list_delete_legacy", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_community_list_name_legacy", "COMML1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_extcomunity_rt_legacy.0", "10:10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_extcomunity_soo_legacy", "10:10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_extcomunity_vpn_distinguisher_legacy", "10:10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_local_preference_legacy", "110"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_route_map.test", "entries.0.set_weight_legacy", "10000"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeRouteMapPrerequisitesConfig + testAccIosxeRouteMapConfig_minimum(),
			},
			{
				Config: testAccIosxeRouteMapPrerequisitesConfig + testAccIosxeRouteMapConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:      "iosxe_route_map.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     "Cisco-IOS-XE-native:native/route-map=RM1",
				Check:             resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

const testAccIosxeRouteMapPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/interface/Loopback=1"
	attributes = {
		"name" = "1"
	}
}

`

func testAccIosxeRouteMapConfig_minimum() string {
	config := `resource "iosxe_route_map" "test" {` + "\n"
	config += `	name = "RM1"` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeRouteMapConfig_all() string {
	config := `resource "iosxe_route_map" "test" {` + "\n"
	config += `	name = "RM1"` + "\n"
	config += `	entries = [{` + "\n"
	config += `		seq = 10` + "\n"
	config += `		operation = "permit"` + "\n"
	config += `		description = "Entry 10"` + "\n"
	config += `		continue = false` + "\n"
	config += `		match_interfaces = ["Loopback1"]` + "\n"
	config += `		match_ip_address_access_lists = ["ACL1"]` + "\n"
	config += `		match_ip_next_hop_access_lists = ["ACL1"]` + "\n"
	config += `		match_ipv6_address_access_lists = "ACL1"` + "\n"
	config += `		match_ipv6_next_hop_access_lists = "ACL1"` + "\n"
	config += `		match_route_type_external = true` + "\n"
	config += `		match_route_type_external_type_1 = true` + "\n"
	config += `		match_route_type_external_type_2 = true` + "\n"
	config += `		match_route_type_internal = true` + "\n"
	config += `		match_route_type_level_1 = true` + "\n"
	config += `		match_route_type_level_2 = true` + "\n"
	config += `		match_route_type_local = true` + "\n"
	config += `		match_source_protocol_bgp = ["65000"]` + "\n"
	config += `		match_source_protocol_connected = true` + "\n"
	config += `		match_source_protocol_eigrp = ["10"]` + "\n"
	config += `		match_source_protocol_isis = true` + "\n"
	config += `		match_source_protocol_lisp = true` + "\n"
	config += `		match_source_protocol_ospf = ["10"]` + "\n"
	config += `		match_source_protocol_ospfv3 = ["10"]` + "\n"
	config += `		match_source_protocol_rip = true` + "\n"
	config += `		match_source_protocol_static = true` + "\n"
	config += `		match_tags = [100]` + "\n"
	config += `		match_track = 1` + "\n"
	config += `		match_as_paths_legacy = [10]` + "\n"
	config += `		match_community_lists_legacy = ["COMM1"]` + "\n"
	config += `		match_extcommunity_lists_legacy = ["EXTCOMM1"]` + "\n"
	config += `		match_local_preferences_legacy = [100]` + "\n"
	config += `		set_default_interfaces = ["Loopback1"]` + "\n"
	config += `		set_global = false` + "\n"
	config += `		set_interfaces = ["Loopback1"]` + "\n"
	config += `		set_ip_address = "PFL1"` + "\n"
	config += `		set_ip_default_global_next_hop_address = ["1.2.3.4"]` + "\n"
	config += `		set_ip_default_next_hop_address = ["1.2.3.4"]` + "\n"
	config += `		set_ip_global_next_hop_address = ["1.2.3.4"]` + "\n"
	config += `		set_ip_next_hop_address = ["1.2.3.4"]` + "\n"
	config += `		set_ip_qos_group = 1` + "\n"
	config += `		set_ipv6_address = ["PFL2"]` + "\n"
	config += `		set_ipv6_default_global_next_hop = "2001::1"` + "\n"
	config += `		set_ipv6_default_next_hop = ["2001::1"]` + "\n"
	config += `		set_ipv6_next_hop = ["2001::1"]` + "\n"
	config += `		set_level_1 = true` + "\n"
	config += `		set_metric_value = 110` + "\n"
	config += `		set_metric_delay = "10"` + "\n"
	config += `		set_metric_reliability = 90` + "\n"
	config += `		set_metric_loading = 10` + "\n"
	config += `		set_metric_mtu = 1500` + "\n"
	config += `		set_metric_type = "external"` + "\n"
	config += `		set_tag = 100` + "\n"
	config += `		set_as_path_prepend_as_legacy = "65001 65001"` + "\n"
	config += `		set_as_path_prepend_last_as_legacy = 5` + "\n"
	config += `		set_as_path_tag_legacy = true` + "\n"
	config += `		set_communities_legacy = ["1:2"]` + "\n"
	config += `		set_communities_additive_legacy = true` + "\n"
	config += `		set_community_list_delete_legacy = true` + "\n"
	config += `		set_community_list_name_legacy = "COMML1"` + "\n"
	config += `		set_extcomunity_rt_legacy = ["10:10"]` + "\n"
	config += `		set_extcomunity_soo_legacy = "10:10"` + "\n"
	config += `		set_extcomunity_vpn_distinguisher_legacy = "10:10"` + "\n"
	config += `		set_local_preference_legacy = 110` + "\n"
	config += `		set_weight_legacy = 10000` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}
