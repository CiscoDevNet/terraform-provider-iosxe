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

func TestAccIosxePrefixList(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_prefix_list.test", "prefixes.0.name", "PREFIX_LIST_1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_prefix_list.test", "prefixes.0.seq", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_prefix_list.test", "prefixes.0.action", "permit"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_prefix_list.test", "prefixes.0.ip", "10.0.0.0/8"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_prefix_list.test", "prefixes.0.ge", "24"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_prefix_list.test", "prefixes.0.le", "32"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_prefix_list.test", "prefix_list_description.0.name", "PREFIX_LIST_1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_prefix_list.test", "prefix_list_description.0.description", "My prefix list"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxePrefixListConfig_minimum(),
			},
			{
				Config: testAccIosxePrefixListConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:      "iosxe_prefix_list.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     "Cisco-IOS-XE-native:native/ip/prefix-lists",
				Check:             resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccIosxePrefixListConfig_minimum() string {
	config := `resource "iosxe_prefix_list" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxePrefixListConfig_all() string {
	config := `resource "iosxe_prefix_list" "test" {` + "\n"
	config += `	prefixes = [{` + "\n"
	config += `		name = "PREFIX_LIST_1"` + "\n"
	config += `		seq = 10` + "\n"
	config += `		action = "permit"` + "\n"
	config += `		ip = "10.0.0.0/8"` + "\n"
	config += `		ge = 24` + "\n"
	config += `		le = 32` + "\n"
	config += `	}]` + "\n"
	config += `	prefix_list_description = [{` + "\n"
	config += `		name = "PREFIX_LIST_1"` + "\n"
	config += `		description = "My prefix list"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}
