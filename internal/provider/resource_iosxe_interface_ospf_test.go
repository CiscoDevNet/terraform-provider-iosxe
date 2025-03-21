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

func TestAccIosxeInterfaceOSPF(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_ospf.test", "cost", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_ospf.test", "dead_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_ospf.test", "hello_interval", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_ospf.test", "mtu_ignore", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_ospf.test", "network_type_broadcast", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_ospf.test", "network_type_non_broadcast", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_ospf.test", "network_type_point_to_multipoint", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_ospf.test", "network_type_point_to_point", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_ospf.test", "priority", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_ospf.test", "ttl_security_hops", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_ospf.test", "process_ids.0.id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_ospf.test", "process_ids.0.areas.0.area_id", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_ospf.test", "message_digest_keys.0.id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_ospf.test", "message_digest_keys.0.md5_auth_key", "mykey"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_ospf.test", "message_digest_keys.0.md5_auth_type", "0"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeInterfaceOSPFPrerequisitesConfig + testAccIosxeInterfaceOSPFConfig_minimum(),
			},
			{
				Config: testAccIosxeInterfaceOSPFPrerequisitesConfig + testAccIosxeInterfaceOSPFConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:            "iosxe_interface_ospf.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateId:           "Cisco-IOS-XE-native:native/interface/Loopback=1/ip/Cisco-IOS-XE-ospf:router-ospf/ospf",
				ImportStateVerifyIgnore: []string{},
				Check:                   resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

const testAccIosxeInterfaceOSPFPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-ospf:router-ospf/ospf/process-id=1"
	attributes = {
		"id" = "1"
	}
}

resource "iosxe_restconf" "PreReq1" {
	path = "Cisco-IOS-XE-native:native/interface/Loopback=1"
	attributes = {
		"name" = "1"
	}
}

`

func testAccIosxeInterfaceOSPFConfig_minimum() string {
	config := `resource "iosxe_interface_ospf" "test" {` + "\n"
	config += `	type = "Loopback"` + "\n"
	config += `	name = "1"` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeInterfaceOSPFConfig_all() string {
	config := `resource "iosxe_interface_ospf" "test" {` + "\n"
	config += `	type = "Loopback"` + "\n"
	config += `	name = "1"` + "\n"
	config += `	cost = 10` + "\n"
	config += `	dead_interval = 30` + "\n"
	config += `	hello_interval = 5` + "\n"
	config += `	mtu_ignore = false` + "\n"
	config += `	network_type_broadcast = false` + "\n"
	config += `	network_type_non_broadcast = false` + "\n"
	config += `	network_type_point_to_multipoint = false` + "\n"
	config += `	network_type_point_to_point = true` + "\n"
	config += `	priority = 10` + "\n"
	config += `	ttl_security_hops = 2` + "\n"
	config += `	process_ids = [{` + "\n"
	config += `		id = 1` + "\n"
	config += `		areas = [{` + "\n"
	config += `			area_id = "0"` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `	message_digest_keys = [{` + "\n"
	config += `		id = 1` + "\n"
	config += `		md5_auth_key = "mykey"` + "\n"
	config += `		md5_auth_type = 0` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]` + "\n"
	config += `}` + "\n"
	return config
}
