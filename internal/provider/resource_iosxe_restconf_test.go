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

func TestAccIosxeRestconf(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeRestconfConfig_empty(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_restconf.test", "id", "Cisco-IOS-XE-native:native/banner/login"),
				),
			},
			{
				Config: testAccIosxeRestconfConfig_banner("My Banner"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_restconf.test", "id", "Cisco-IOS-XE-native:native/banner/login"),
					resource.TestCheckResourceAttr("iosxe_restconf.test", "attributes.banner", "My Banner"),
				),
			},
			{
				ResourceName:  "iosxe_restconf.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/banner/login",
			},
			{
				Config: testAccIosxeRestconfConfig_banner("My New Banner"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_restconf.test", "attributes.banner", "My New Banner"),
				),
			},
			{
				Config: testAccIosxeRestconfConfig_nested(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_restconf.nested", "attributes.hostname", "R1"),
					resource.TestCheckResourceAttr("iosxe_restconf.nested", "lists.0.name", "route-map"),
					resource.TestCheckResourceAttr("iosxe_restconf.nested", "lists.0.items.0.name", "test123"),
				),
			},
		},
	})
}

func testAccIosxeRestconfConfig_empty() string {
	return `
	resource "iosxe_restconf" "test" {
		path = "Cisco-IOS-XE-native:native/banner/login"
	}
	`
}

func testAccIosxeRestconfConfig_banner(message string) string {
	return fmt.Sprintf(`
	resource "iosxe_restconf" "test" {
		path = "Cisco-IOS-XE-native:native/banner/login"
		attributes = {
			banner = "%s"
		}
	}
	`, message)
}

func testAccIosxeRestconfConfig_nested() string {
	return `
	resource "iosxe_restconf" "nested" {
		path = "Cisco-IOS-XE-native:native"
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
