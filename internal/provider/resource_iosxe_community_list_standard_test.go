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

func TestAccIosxeCommunityListStandard(t *testing.T) {
	if os.Getenv("IOSXE1712") == "" {
		t.Skip("skipping test, set environment variable IOSXE1712")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_community_list_standard.test", "name", "CLS1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_community_list_standard.test", "deny_entries.0", "65000:500"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_community_list_standard.test", "permit_entries.0", "65000:501"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeCommunityListStandardConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:            "iosxe_community_list_standard.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateId:           "Cisco-IOS-XE-native:native/ip/Cisco-IOS-XE-bgp:community-list/standard=CLS1",
				ImportStateVerifyIgnore: []string{},
				Check:                   resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccIosxeCommunityListStandardConfig_minimum() string {
	config := `resource "iosxe_community_list_standard" "test" {` + "\n"
	config += `	name = "CLS1"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeCommunityListStandardConfig_all() string {
	config := `resource "iosxe_community_list_standard" "test" {` + "\n"
	config += `	name = "CLS1"` + "\n"
	config += `	deny_entries = ["65000:500"]` + "\n"
	config += `	permit_entries = ["65000:501"]` + "\n"
	config += `}` + "\n"
	return config
}
