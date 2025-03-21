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
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "name", "profile1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "description", "My description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "authentication_remote_pre_share", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "authentication_local_pre_share", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "identity_local_key_id", "key1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "match_address_local_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "match_fvrf_any", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "match_identity_remote_ipv4_addresses.0.address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "match_identity_remote_ipv4_addresses.0.mask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "match_identity_remote_keys.0", "key1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "keyring_local", "test"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "ivrf", "VRF1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "dpd_interval", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "dpd_retry", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "dpd_query", "periodic"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_profile.test", "config_exchange_request", "false"))
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
				ResourceName:      "iosxe_crypto_ikev2_profile.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     "Cisco-IOS-XE-native:native/crypto/Cisco-IOS-XE-crypto:ikev2/profile=profile1",
				Check:             resource.ComposeTestCheckFunc(checks...),
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

resource "iosxe_restconf" "PreReq1" {
	path = "Cisco-IOS-XE-native:native/vrf/definition=VRF1"
	delete = false
	attributes = {
		"name" = "VRF1"
		"address-family/ipv4" = ""
	}
}

`

func testAccIosxeCryptoIKEv2ProfileConfig_minimum() string {
	config := `resource "iosxe_crypto_ikev2_profile" "test" {` + "\n"
	config += `	name = "profile1"` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeCryptoIKEv2ProfileConfig_all() string {
	config := `resource "iosxe_crypto_ikev2_profile" "test" {` + "\n"
	config += `	name = "profile1"` + "\n"
	config += `	description = "My description"` + "\n"
	config += `	authentication_remote_pre_share = true` + "\n"
	config += `	authentication_local_pre_share = true` + "\n"
	config += `	identity_local_key_id = "key1"` + "\n"
	config += `	match_address_local_ip = "1.2.3.4"` + "\n"
	config += `	match_fvrf_any = true` + "\n"
	config += `	match_identity_remote_ipv4_addresses = [{` + "\n"
	config += `		address = "1.2.3.4"` + "\n"
	config += `		mask = "255.255.255.0"` + "\n"
	config += `	}]` + "\n"
	config += `	match_identity_remote_keys = ["key1"]` + "\n"
	config += `	keyring_local = "test"` + "\n"
	config += `	ivrf = "VRF1"` + "\n"
	config += `	dpd_interval = 10` + "\n"
	config += `	dpd_retry = 2` + "\n"
	config += `	dpd_query = "periodic"` + "\n"
	config += `	config_exchange_request = false` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]` + "\n"
	config += `}` + "\n"
	return config
}
