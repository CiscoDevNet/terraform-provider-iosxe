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

func TestAccDataSourceIosxeInterfaceOSPF(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ospf.test", "cost", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ospf.test", "dead_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ospf.test", "hello_interval", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ospf.test", "mtu_ignore", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ospf.test", "network_type_broadcast", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ospf.test", "network_type_non_broadcast", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ospf.test", "network_type_point_to_multipoint", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ospf.test", "network_type_point_to_point", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ospf.test", "priority", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ospf.test", "process_ids.0.id", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_interface_ospf.test", "process_ids.0.areas.0.area_id", "0"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeInterfaceOSPFPrerequisitesConfig + testAccDataSourceIosxeInterfaceOSPFConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

const testAccDataSourceIosxeInterfaceOSPFPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/interface/Loopback=1"
	attributes = {
		"name" = "1"
	}
}

`

func testAccDataSourceIosxeInterfaceOSPFConfig() string {
	config := `resource "iosxe_interface_ospf" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
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
	config += `	process_ids = [{` + "\n"
	config += `		id = 10` + "\n"
	config += `		areas = [{` + "\n"
	config += `			area_id = "0"` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxe_interface_ospf" "test" {
			type = "Loopback"
			name = "1"
			depends_on = [iosxe_interface_ospf.test]
		}
	`
	return config
}
