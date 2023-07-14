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

func TestAccDataSourceIosxeService(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_service.test", "pad", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_service.test", "password_encryption", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_service.test", "password_recovery", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_service.test", "timestamps", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_service.test", "timestamps_debug", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_service.test", "timestamps_debug_datetime", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_service.test", "timestamps_debug_datetime_msec", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_service.test", "timestamps_debug_datetime_localtime", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_service.test", "timestamps_debug_datetime_show_timezone", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_service.test", "timestamps_debug_datetime_year", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_service.test", "timestamps_debug_uptime", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_service.test", "timestamps_log", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_service.test", "timestamps_log_datetime", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_service.test", "timestamps_log_datetime_msec", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_service.test", "timestamps_log_datetime_localtime", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_service.test", "timestamps_log_datetime_show_timezone", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_service.test", "timestamps_log_datetime_year", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_service.test", "timestamps_log_uptime", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_service.test", "dhcp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_service.test", "tcp_keepalives_in", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_service.test", "tcp_keepalives_out", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeServiceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxeServiceConfig() string {
	config := `resource "iosxe_service" "test" {` + "\n"
	config += `	pad = true` + "\n"
	config += `	password_encryption = true` + "\n"
	config += `	password_recovery = true` + "\n"
	config += `	timestamps = true` + "\n"
	config += `	timestamps_debug = true` + "\n"
	config += `	timestamps_debug_datetime = true` + "\n"
	config += `	timestamps_debug_datetime_msec = true` + "\n"
	config += `	timestamps_debug_datetime_localtime = true` + "\n"
	config += `	timestamps_debug_datetime_show_timezone = true` + "\n"
	config += `	timestamps_debug_datetime_year = true` + "\n"
	config += `	timestamps_debug_uptime = true` + "\n"
	config += `	timestamps_log = true` + "\n"
	config += `	timestamps_log_datetime = true` + "\n"
	config += `	timestamps_log_datetime_msec = true` + "\n"
	config += `	timestamps_log_datetime_localtime = true` + "\n"
	config += `	timestamps_log_datetime_show_timezone = true` + "\n"
	config += `	timestamps_log_datetime_year = true` + "\n"
	config += `	timestamps_log_uptime = true` + "\n"
	config += `	dhcp = true` + "\n"
	config += `	tcp_keepalives_in = true` + "\n"
	config += `	tcp_keepalives_out = true` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxe_service" "test" {
			depends_on = [iosxe_service.test]
		}
	`
	return config
}
