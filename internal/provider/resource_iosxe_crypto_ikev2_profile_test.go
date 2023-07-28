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

func TestAccIosxeCryptoIKEv2Profile(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "name", "aws_vpg_1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "match_identity_remote_address_ipv4.0.ipv4_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "match_identity_remote_address_ipv4.0.ipv4_mask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "authentication_remote_pre_share", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "authentication_local_pre_share", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "keyring_local_name", "test"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "dpd_interval", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "dpd_retry", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "dpd_query", "periodic"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "config_exchange_request_1", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeCryptoIKEv2ProfilePrerequisitesConfig + testAccIosxeCryptoIKEv2ProfileConfig_minimum(),
			},
			{
				Config: testAccIosxeCryptoIKEv2ProfilePrerequisitesConfig + testAccIosxeCryptoIKEv2ProfileConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_crypto_ikev2_profile.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/crypto/Cisco-IOS-XE-crypto:ikev2/profile=aws_vpg_1",
			},
		},
	})
}

const testAccIosxeCryptoIKEv2ProfilePrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/crypto/Cisco-IOS-XE-crypto:ikev2/keyring=test"
	attributes = {
		"name" = "test"
	}
}

`

func testAccIosxeCryptoIKEv2ProfileConfig_minimum() string {
	config := `resource "iosxe_crypto_ikev2_profile" "test" {` + "\n"
	config += `	name = "aws_vpg_1"` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeCryptoIKEv2ProfileConfig_all() string {
	config := `resource "iosxe_crypto_ikev2_profile" "test" {` + "\n"
	config += `	name = "aws_vpg_1"` + "\n"
	config += `	match_identity_remote_address_ipv4 = [{` + "\n"
	config += `		ipv4_address = "1.2.3.4"` + "\n"
	config += `		ipv4_mask = "255.255.255.0"` + "\n"
	config += `	}]` + "\n"
	config += `	authentication_remote_pre_share = true` + "\n"
	config += `	authentication_local_pre_share = true` + "\n"
	config += `	keyring_local_name = "test"` + "\n"
	config += `	dpd_interval = 10` + "\n"
	config += `	dpd_retry = 2` + "\n"
	config += `	dpd_query = "periodic"` + "\n"
	config += `	config_exchange_request_1 = false` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"
	return config
}
