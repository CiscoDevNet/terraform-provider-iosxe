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

func TestAccIosxeSystem(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_system.test", "hostname", "ROUTER-1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_system.test", "ip_bgp_community_new_format", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_system.test", "ipv6_unicast_routing", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_system.test", "ip_source_route", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_system.test", "ip_domain_lookup", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_system.test", "ip_domain_name", "test.com"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_system.test", "login_delay", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_system.test", "login_on_failure", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_system.test", "login_on_failure_log", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_system.test", "login_on_success", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_system.test", "login_on_success_log", "true"))
	if os.Getenv("C8000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_system.test", "ip_multicast_routing", "true"))
	}
	if os.Getenv("C9000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_system.test", "multicast_routing_switch", "true"))
	}
	if os.Getenv("C8000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_system.test", "ip_multicast_routing_distributed", "true"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_system.test", "multicast_routing_vrfs.0.vrf", "VRF1"))
	if os.Getenv("C8000V") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("iosxe_system.test", "multicast_routing_vrfs.0.distributed", "true"))
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeSystemPrerequisitesConfig + testAccIosxeSystemConfig_minimum(),
			},
			{
				Config: testAccIosxeSystemPrerequisitesConfig + testAccIosxeSystemConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:            "iosxe_system.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateId:           "Cisco-IOS-XE-native:native",
				ImportStateVerifyIgnore: []string{"ip_multicast_routing", "multicast_routing_switch", "ip_multicast_routing_distributed", "multicast_routing_vrfs.0.distributed", "ip_http_authentication_aaa", "ip_http_authentication_local"},
				Check:                   resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

const testAccIosxeSystemPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
	delete = false
	attributes = {
		"name" = "VRF1"
		"address-family/ipv4" = ""
	}
}

`

func testAccIosxeSystemConfig_minimum() string {
	config := `resource "iosxe_system" "test" {` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeSystemConfig_all() string {
	config := `resource "iosxe_system" "test" {` + "\n"
	config += `	hostname = "ROUTER-1"` + "\n"
	config += `	ip_bgp_community_new_format = true` + "\n"
	config += `	ipv6_unicast_routing = true` + "\n"
	config += `	ip_source_route = false` + "\n"
	config += `	ip_domain_lookup = false` + "\n"
	config += `	ip_domain_name = "test.com"` + "\n"
	config += `	login_delay = 10` + "\n"
	config += `	login_on_failure = true` + "\n"
	config += `	login_on_failure_log = true` + "\n"
	config += `	login_on_success = true` + "\n"
	config += `	login_on_success_log = true` + "\n"
	if os.Getenv("C8000V") != "" {
		config += `	ip_multicast_routing = true` + "\n"
	}
	if os.Getenv("C9000V") != "" {
		config += `	multicast_routing_switch = true` + "\n"
	}
	if os.Getenv("C8000V") != "" {
		config += `	ip_multicast_routing_distributed = true` + "\n"
	}
	config += `	multicast_routing_vrfs = [{` + "\n"
	config += `		vrf = "VRF1"` + "\n"
	if os.Getenv("C8000V") != "" {
		config += `		distributed = true` + "\n"
	}
	config += `	}]` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}
