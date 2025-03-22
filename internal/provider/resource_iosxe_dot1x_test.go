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

func TestAccIosxeDot1x(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_dot1x.test", "auth_fail_eapol", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_dot1x.test", "credentials.0.profile_name", "profile1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_dot1x.test", "credentials.0.description", "credential_profile_name"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_dot1x.test", "credentials.0.username", "username1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_dot1x.test", "credentials.0.password_type", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_dot1x.test", "credentials.0.password", "password123"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_dot1x.test", "credentials.0.pki_trustpoint", "trustpoint1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_dot1x.test", "credentials.0.anonymous_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_dot1x.test", "critical_eapol_config_block", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_dot1x.test", "test_timeout", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_dot1x.test", "logging_verbose", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_dot1x.test", "supplicant_controlled_transient", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_dot1x.test", "supplicant_force_multicast", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_dot1x.test", "system_auth_control", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeDot1xConfig_minimum(),
			},
			{
				Config: testAccIosxeDot1xConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:            "iosxe_dot1x.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateId:           "Cisco-IOS-XE-native:native/dot1x",
				ImportStateVerifyIgnore: []string{},
				Check:                   resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccIosxeDot1xConfig_minimum() string {
	config := `resource "iosxe_dot1x" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeDot1xConfig_all() string {
	config := `resource "iosxe_dot1x" "test" {` + "\n"
	config += `	auth_fail_eapol = true` + "\n"
	config += `	credentials = [{` + "\n"
	config += `		profile_name = "profile1"` + "\n"
	config += `		description = "credential_profile_name"` + "\n"
	config += `		username = "username1"` + "\n"
	config += `		password_type = "0"` + "\n"
	config += `		password = "password123"` + "\n"
	config += `		pki_trustpoint = "trustpoint1"` + "\n"
	config += `		anonymous_id = "1"` + "\n"
	config += `	}]` + "\n"
	config += `	critical_eapol_config_block = true` + "\n"
	config += `	test_timeout = 1000` + "\n"
	config += `	logging_verbose = true` + "\n"
	config += `	supplicant_controlled_transient = true` + "\n"
	config += `	supplicant_force_multicast = true` + "\n"
	config += `	system_auth_control = true` + "\n"
	config += `}` + "\n"
	return config
}
