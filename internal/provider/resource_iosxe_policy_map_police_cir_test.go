// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxePolicyMapPoliceCirExceedTransmit(t *testing.T) {
	resourceTransmitChecks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr("iosxe_policy_map.police_cir", "name", "TFACC_COPP_POLICE_TRANSMIT"),
		resource.TestCheckResourceAttr("iosxe_policy_map.police_cir", "classes.0.name", "TFACC_COPP_ICMP"),
		resource.TestCheckResourceAttr("iosxe_policy_map.police_cir", "classes.0.actions.0.type", "police"),
		resource.TestCheckResourceAttr("iosxe_policy_map.police_cir", "classes.0.actions.0.police_cir", "256000"),
		resource.TestCheckResourceAttr("iosxe_policy_map.police_cir", "classes.0.actions.0.police_bc", "8000"),
		resource.TestCheckResourceAttr("iosxe_policy_map.police_cir", "classes.0.actions.0.police_cir_conform_transmit", "true"),
		resource.TestCheckResourceAttr("iosxe_policy_map.police_cir", "classes.0.actions.0.police_cir_exceed_drop", "true"),
		resource.TestCheckResourceAttr("iosxe_policy_map.police_cir", "classes.1.name", "class-default"),
		resource.TestCheckResourceAttr("iosxe_policy_map.police_cir", "classes.1.actions.0.type", "police"),
		resource.TestCheckResourceAttr("iosxe_policy_map.police_cir", "classes.1.actions.0.police_cir", "10000000"),
		resource.TestCheckResourceAttr("iosxe_policy_map.police_cir", "classes.1.actions.0.police_bc", "312500"),
		resource.TestCheckResourceAttr("iosxe_policy_map.police_cir", "classes.1.actions.0.police_cir_conform_transmit", "true"),
		resource.TestCheckResourceAttr("iosxe_policy_map.police_cir", "classes.1.actions.0.police_cir_exceed_transmit", "true"),
	}
	transmitChecks := append([]resource.TestCheckFunc{}, resourceTransmitChecks...)
	transmitChecks = append(transmitChecks, resource.TestCheckResourceAttr("data.iosxe_policy_map.police_cir", "classes.1.actions.0.police_cir_exceed_transmit", "true"))
	dropChecks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr("iosxe_policy_map.police_cir", "classes.1.actions.0.police_cir_exceed_drop", "true"),
		resource.TestCheckResourceAttr("data.iosxe_policy_map.police_cir", "classes.1.actions.0.police_cir_exceed_drop", "true"),
		resource.TestCheckResourceAttr("data.iosxe_policy_map.police_cir", "classes.1.actions.0.police_cir_exceed_transmit", "false"),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxePolicyMapPoliceCirExceedTransmitConfig(true),
				Check:  resource.ComposeTestCheckFunc(transmitChecks...),
			},
			{
				Config: testAccIosxePolicyMapPoliceCirExceedTransmitConfig(false),
				Check:  resource.ComposeTestCheckFunc(dropChecks...),
			},
			{
				Config: testAccIosxePolicyMapPoliceCirExceedTransmitConfig(true),
				Check:  resource.ComposeTestCheckFunc(transmitChecks...),
			},
			{
				ResourceName:      "iosxe_policy_map.police_cir",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     "TFACC_COPP_POLICE_TRANSMIT",
				ImportStateVerifyIgnore: []string{
					"subscriber",
					"classes.0.policy_log",
					"classes.0.actions.0.shape_average_ms",
					"classes.0.actions.0.police_cir_exceed_transmit",
					"classes.0.actions.0.police_target_bitrate_conform_transmit",
					"classes.0.actions.0.police_target_bitrate_exceed_drop",
					"classes.0.actions.0.police_target_bitrate_exceed_transmit",
					"classes.1.policy_log",
					"classes.1.actions.0.shape_average_ms",
					"classes.1.actions.0.police_cir_exceed_drop",
					"classes.1.actions.0.police_target_bitrate_conform_transmit",
					"classes.1.actions.0.police_target_bitrate_exceed_drop",
					"classes.1.actions.0.police_target_bitrate_exceed_transmit",
				},
				Check: resource.ComposeTestCheckFunc(resourceTransmitChecks...),
			},
		},
	})
}

func testAccIosxePolicyMapPoliceCirExceedTransmitConfig(exceedTransmit bool) string {
	exceedAction := "police_cir_exceed_drop = true"
	if exceedTransmit {
		exceedAction = "police_cir_exceed_transmit = true"
	}

	return fmt.Sprintf(`resource "iosxe_yang" "policy_class_map" {
  path = "/Cisco-IOS-XE-native:native/policy/Cisco-IOS-XE-policy:class-map[name=TFACC_COPP_ICMP]"
  attributes = {
    "name" = "TFACC_COPP_ICMP"
    "prematch" = "match-any"
  }
}

resource "iosxe_policy_map" "police_cir" {
  name = "TFACC_COPP_POLICE_TRANSMIT"
  classes = [
    {
      name = "TFACC_COPP_ICMP"
      actions = [
        {
          type = "police"
          police_cir = 256000
          police_bc = 8000
          police_cir_conform_transmit = true
          police_cir_exceed_drop = true
        },
      ]
    },
    {
      name = "class-default"
      actions = [
        {
          type = "police"
          police_cir = 10000000
          police_bc = 312500
          police_cir_conform_transmit = true
          %s
        },
      ]
    },
  ]
  depends_on = [iosxe_yang.policy_class_map]
}

data "iosxe_policy_map" "police_cir" {
  name = iosxe_policy_map.police_cir.name
}
`, exceedAction)
}
