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

func TestAccIosxeAccessListStandard(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_access_list_standard.test", "name", "SACL1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_access_list_standard.test", "entries.0.sequence", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_access_list_standard.test", "entries.0.remark", "Description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_access_list_standard.test", "entries.0.deny_prefix", "10.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_access_list_standard.test", "entries.0.deny_prefix_mask", "0.0.0.255"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_access_list_standard.test", "entries.0.deny_log", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeAccessListStandardConfig_minimum(),
			},
			{
				Config: testAccIosxeAccessListStandardConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:            "iosxe_access_list_standard.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateId:           "Cisco-IOS-XE-native:native/ip/access-list/Cisco-IOS-XE-acl:standard=SACL1",
				ImportStateVerifyIgnore: []string{"entries.0.deny_any", "entries.0.permit_any", "entries.0.permit_log"},
				Check:                   resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccIosxeAccessListStandardConfig_minimum() string {
	config := `resource "iosxe_access_list_standard" "test" {` + "\n"
	config += `	name = "SACL1"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeAccessListStandardConfig_all() string {
	config := `resource "iosxe_access_list_standard" "test" {` + "\n"
	config += `	name = "SACL1"` + "\n"
	config += `	entries = [{` + "\n"
	config += `		sequence = 10` + "\n"
	config += `		remark = "Description"` + "\n"
	config += `		deny_prefix = "10.0.0.0"` + "\n"
	config += `		deny_prefix_mask = "0.0.0.255"` + "\n"
	config += `		deny_log = true` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}
