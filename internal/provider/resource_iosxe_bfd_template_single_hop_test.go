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

func TestAccIosxeBFDTemplateSingleHop(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bfd_template_single_hop.test", "name", "SH-TEMPLATE-1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bfd_template_single_hop.test", "authentication_md5_keychain", "KEYC1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bfd_template_single_hop.test", "interval_milliseconds_min_tx", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bfd_template_single_hop.test", "interval_milliseconds_min_rx", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bfd_template_single_hop.test", "interval_milliseconds_multiplier", "4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bfd_template_single_hop.test", "echo", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bfd_template_single_hop.test", "dampening_half_time", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bfd_template_single_hop.test", "dampening_unsuppress_time", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bfd_template_single_hop.test", "dampening_suppress_time", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bfd_template_single_hop.test", "dampening_max_suppressing_time", "60"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeBFDTemplateSingleHopConfig_minimum(),
			},
			{
				Config: testAccIosxeBFDTemplateSingleHopConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:            "iosxe_bfd_template_single_hop.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateId:           "Cisco-IOS-XE-native:native/bfd-template/Cisco-IOS-XE-bfd:single-hop=SH-TEMPLATE-1",
				ImportStateVerifyIgnore: []string{},
				Check:                   resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccIosxeBFDTemplateSingleHopConfig_minimum() string {
	config := `resource "iosxe_bfd_template_single_hop" "test" {` + "\n"
	config += `	name = "SH-TEMPLATE-1"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeBFDTemplateSingleHopConfig_all() string {
	config := `resource "iosxe_bfd_template_single_hop" "test" {` + "\n"
	config += `	name = "SH-TEMPLATE-1"` + "\n"
	config += `	authentication_md5_keychain = "KEYC1"` + "\n"
	config += `	interval_milliseconds_min_tx = 200` + "\n"
	config += `	interval_milliseconds_min_rx = 200` + "\n"
	config += `	interval_milliseconds_multiplier = 4` + "\n"
	config += `	echo = true` + "\n"
	config += `	dampening_half_time = 30` + "\n"
	config += `	dampening_unsuppress_time = 30` + "\n"
	config += `	dampening_suppress_time = 100` + "\n"
	config += `	dampening_max_suppressing_time = 60` + "\n"
	config += `}` + "\n"
	return config
}
