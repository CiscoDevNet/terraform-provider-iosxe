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

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeYang(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeYangConfig_empty(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_yang.test", "id", "/Cisco-IOS-XE-native:native/banner/login"),
				),
			},
			{
				Config: testAccIosxeYangConfig_banner("My Banner"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_yang.test", "id", "/Cisco-IOS-XE-native:native/banner/login"),
					resource.TestCheckResourceAttr("iosxe_yang.test", "attributes.banner", "My Banner"),
				),
			},
			{
				ResourceName:  "iosxe_yang.test",
				ImportState:   true,
				ImportStateId: "/Cisco-IOS-XE-native:native/banner/login",
			},
			{
				Config: testAccIosxeYangConfig_banner("My New Banner"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_yang.test", "attributes.banner", "My New Banner"),
				),
			},
			{
				Config: testAccIosxeYangConfig_nested(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_yang.nested", "attributes.hostname", "R1"),
					resource.TestCheckResourceAttr("iosxe_yang.nested", "lists.0.name", "route-map"),
					resource.TestCheckResourceAttr("iosxe_yang.nested", "lists.0.items.0.name", "test123"),
				),
			},
		},
	})
}

func testAccIosxeYangConfig_empty() string {
	return `
	resource "iosxe_yang" "test" {
		path = "/Cisco-IOS-XE-native:native/banner/login"
	}
	`
}

func testAccIosxeYangConfig_banner(message string) string {
	return fmt.Sprintf(`
	resource "iosxe_yang" "test" {
		path = "/Cisco-IOS-XE-native:native/banner/login"
		attributes = {
			banner = "%s"
		}
	}
	`, message)
}

func testAccIosxeYangConfig_nested() string {
	return `
	resource "iosxe_yang" "nested" {
		path = "/Cisco-IOS-XE-native:native"
		delete = false
		attributes = {
			hostname = "R1"
		}
		lists = [
			{
				name = "route-map"
				key = "name"
				items = [
					{
						name = "test123"
					}
				]
			}
		]
	}
	`
}
