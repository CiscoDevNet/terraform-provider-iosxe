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

func TestAccIosxeCryptoIKEv2Keyring(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_keyring.test", "name", "keyring1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_keyring.test", "peers.0.name", "peer1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_keyring.test", "peers.0.description", "My description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_keyring.test", "peers.0.ipv4_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_keyring.test", "peers.0.ipv4_mask", "255.255.255.248"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_keyring.test", "peers.0.identity_key_id", "key1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_keyring.test", "peers.0.pre_shared_key_local_encryption", "6"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_keyring.test", "peers.0.pre_shared_key_local", "cisco123"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_keyring.test", "peers.0.pre_shared_key_remote_encryption", "6"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_crypto_ikev2_keyring.test", "peers.0.pre_shared_key_remote", "cisco123"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeCryptoIKEv2KeyringConfig_minimum(),
			},
			{
				Config: testAccIosxeCryptoIKEv2KeyringConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:            "iosxe_crypto_ikev2_keyring.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateId:           "Cisco-IOS-XE-native:native/crypto/Cisco-IOS-XE-crypto:ikev2/keyring=keyring1",
				ImportStateVerifyIgnore: []string{},
				Check:                   resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccIosxeCryptoIKEv2KeyringConfig_minimum() string {
	config := `resource "iosxe_crypto_ikev2_keyring" "test" {` + "\n"
	config += `	name = "keyring1"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeCryptoIKEv2KeyringConfig_all() string {
	config := `resource "iosxe_crypto_ikev2_keyring" "test" {` + "\n"
	config += `	name = "keyring1"` + "\n"
	config += `	peers = [{` + "\n"
	config += `		name = "peer1"` + "\n"
	config += `		description = "My description"` + "\n"
	config += `		ipv4_address = "1.2.3.4"` + "\n"
	config += `		ipv4_mask = "255.255.255.248"` + "\n"
	config += `		identity_key_id = "key1"` + "\n"
	config += `		pre_shared_key_local_encryption = "6"` + "\n"
	config += `		pre_shared_key_local = "cisco123"` + "\n"
	config += `		pre_shared_key_remote_encryption = "6"` + "\n"
	config += `		pre_shared_key_remote = "cisco123"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}
