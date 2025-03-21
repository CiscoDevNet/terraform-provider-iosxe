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

func TestAccIosxeCryptoIPSecProfile(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ipsec_profile.test", "name", "vpn200"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ipsec_profile.test", "set_transform_set.0", "TS1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ipsec_profile.test", "set_isakmp_profile_ikev2_profile_ikev2_profile_case_ikev2_profile", "vpn300"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeCryptoIPSecProfilePrerequisitesConfig + testAccIosxeCryptoIPSecProfileConfig_minimum(),
			},
			{
				Config: testAccIosxeCryptoIPSecProfilePrerequisitesConfig + testAccIosxeCryptoIPSecProfileConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:            "iosxe_crypto_ipsec_profile.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateId:           "Cisco-IOS-XE-native:native/crypto/Cisco-IOS-XE-crypto:ipsec/profile=vpn200",
				ImportStateVerifyIgnore: []string{},
				Check:                   resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

const testAccIosxeCryptoIPSecProfilePrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/crypto/Cisco-IOS-XE-crypto:ipsec/transform-set=TS1"
	attributes = {
		"tag" = "TS1"
		"esp" = "esp-aes"
		"key-bit" = "192"
	}
}

resource "iosxe_restconf" "PreReq1" {
	path = "Cisco-IOS-XE-native:native/crypto/Cisco-IOS-XE-crypto:ikev2/profile=vpn300"
	attributes = {
		"name" = "vpn300"
		"aaa/authentication/anyconnect-eap" = "attached"
	}
}

`

func testAccIosxeCryptoIPSecProfileConfig_minimum() string {
	config := `resource "iosxe_crypto_ipsec_profile" "test" {` + "\n"
	config += `	name = "vpn200"` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeCryptoIPSecProfileConfig_all() string {
	config := `resource "iosxe_crypto_ipsec_profile" "test" {` + "\n"
	config += `	name = "vpn200"` + "\n"
	config += `	set_transform_set = ["TS1"]` + "\n"
	config += `	set_isakmp_profile_ikev2_profile_ikev2_profile_case_ikev2_profile = "vpn300"` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]` + "\n"
	config += `}` + "\n"
	return config
}
