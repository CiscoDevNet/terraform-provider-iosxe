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

func TestAccIosxeCDP(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_cdp.test", "holdtime", "15"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_cdp.test", "timer", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_cdp.test", "run", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_cdp.test", "filter_tlv_list", "TLIST"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_cdp.test", "tlv_lists.0.name", "TLIST"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_cdp.test", "tlv_lists.0.vtp_mgmt_domain", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_cdp.test", "tlv_lists.0.cos", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_cdp.test", "tlv_lists.0.duplex", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_cdp.test", "tlv_lists.0.trust", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_cdp.test", "tlv_lists.0.version", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeCDPConfig_minimum(),
			},
			{
				Config: testAccIosxeCDPConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:      "iosxe_cdp.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     "Cisco-IOS-XE-native:native/cdp",
				Check:             resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccIosxeCDPConfig_minimum() string {
	config := `resource "iosxe_cdp" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeCDPConfig_all() string {
	config := `resource "iosxe_cdp" "test" {` + "\n"
	config += `	holdtime = 15` + "\n"
	config += `	timer = 5` + "\n"
	config += `	run = true` + "\n"
	config += `	filter_tlv_list = "TLIST"` + "\n"
	config += `	tlv_lists = [{` + "\n"
	config += `		name = "TLIST"` + "\n"
	config += `		vtp_mgmt_domain = true` + "\n"
	config += `		cos = true` + "\n"
	config += `		duplex = true` + "\n"
	config += `		trust = true` + "\n"
	config += `		version = true` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}
