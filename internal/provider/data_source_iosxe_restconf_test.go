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

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeRestconf(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeRestconfConfigInterface,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxe_restconf.test", "id", "Cisco-IOS-XE-native:native/banner/login"),
					resource.TestCheckResourceAttr("data.iosxe_restconf.test", "attributes.banner", "My Banner"),
				),
			},
		},
	})
}

const testAccDataSourceIosxeRestconfConfigInterface = `
resource "iosxe_restconf" "test" {
	path = "Cisco-IOS-XE-native:native/banner/login"
	attributes = {
		banner = "My Banner"
	}
}

data "iosxe_restconf" "test" {
	path = "Cisco-IOS-XE-native:native/banner/login"
	depends_on = [iosxe_restconf.test]
}
`
