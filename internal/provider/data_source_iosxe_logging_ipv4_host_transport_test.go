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

func TestAccDataSourceIosxeLoggingIPv4HostTransport(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_logging_ipv4_host_transport.test", "transport_udp_ports.0.port_number", "10000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_logging_ipv4_host_transport.test", "transport_tcp_ports.0.port_number", "10001"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_logging_ipv4_host_transport.test", "transport_tls_ports.0.port_number", "10002"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeLoggingIPv4HostTransportPrerequisitesConfig + testAccDataSourceIosxeLoggingIPv4HostTransportConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

const testAccDataSourceIosxeLoggingIPv4HostTransportPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/logging/host/ipv4-host-list=2.2.2.2"
	attributes = {
		"ipv4-host" = "2.2.2.2"
	}
}

`

func testAccDataSourceIosxeLoggingIPv4HostTransportConfig() string {
	config := `resource "iosxe_logging_ipv4_host_transport" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	ipv4_host = "2.2.2.2"` + "\n"
	config += `	transport_udp_ports = [{` + "\n"
	config += `		port_number = 10000` + "\n"
	config += `	}]` + "\n"
	config += `	transport_tcp_ports = [{` + "\n"
	config += `		port_number = 10001` + "\n"
	config += `	}]` + "\n"
	config += `	transport_tls_ports = [{` + "\n"
	config += `		port_number = 10002` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, ]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxe_logging_ipv4_host_transport" "test" {
			ipv4_host = "2.2.2.2"
			depends_on = [iosxe_logging_ipv4_host_transport.test]
		}
	`
	return config
}
