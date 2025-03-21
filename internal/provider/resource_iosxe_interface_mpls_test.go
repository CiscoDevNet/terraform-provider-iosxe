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

func TestAccIosxeInterfaceMPLS(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_mpls.test", "ip", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_interface_mpls.test", "mtu", "1200"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeInterfaceMPLSPrerequisitesConfig + testAccIosxeInterfaceMPLSConfig_minimum(),
			},
			{
				Config: testAccIosxeInterfaceMPLSPrerequisitesConfig + testAccIosxeInterfaceMPLSConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:      "iosxe_interface_mpls.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     "Cisco-IOS-XE-native:native/interface/Loopback=1/mpls",
				Check:             resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

const testAccIosxeInterfaceMPLSPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/interface/Loopback=1"
	attributes = {
		"name" = "1"
	}
}

`

func testAccIosxeInterfaceMPLSConfig_minimum() string {
	config := `resource "iosxe_interface_mpls" "test" {` + "\n"
	config += `	type = "Loopback"` + "\n"
	config += `	name = "1"` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeInterfaceMPLSConfig_all() string {
	config := `resource "iosxe_interface_mpls" "test" {` + "\n"
	config += `	type = "Loopback"` + "\n"
	config += `	name = "1"` + "\n"
	config += `	ip = true` + "\n"
	config += `	mtu = "1200"` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}
