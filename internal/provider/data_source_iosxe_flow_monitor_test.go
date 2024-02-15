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

func TestAccDataSourceIosxeFlowMonitor(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_record.0.name", "FNF-input"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_record.0.description", "IPv4-NetFlow"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_record.0.match_ipv4_source_address", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_record.0.match_ipv4_destination_address", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_record.0.match_ipv4_protocol", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_record.0.match_ipv4_tos", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_record.0.match_transport_source_port", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_record.0.match_transport_destination_port", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_record.0.match_interface_input", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_record.0.match_flow_direction", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_record.0.collect_interface_output", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_record.0.collect_counter_bytes_long", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_record.0.collect_counter_packets_long", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_record.0.collect_transport_tcp_flags", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_record.0.collect_timestamp_absolute_first", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_record.0.collect_timestamp_absolute_last", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_exporter.0.name", "Scrutinizer"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_exporter.0.description", "Export to Scrutinizer"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_exporter.0.destination_destination_choice_ipdest_ipdest_ip", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_exporter.0.source_loopback", "1474"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_exporter.0.transport_udp", "655"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_exporter.0.template_data_timeout", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_monitor.0.name", "Scrut_mon_input"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_monitor.0.description", "IPv4 FNF ingress exports"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_monitor.0.flow_monitor_exporter.0.name", "Scrutinizer"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_monitor.0.cache_timeout_active", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_monitor.0.record_type", "FNF_Input"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeFlowMonitorConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxeFlowMonitorConfig() string {
	config := `resource "iosxe_flow_monitor" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	flow_record = [{` + "\n"
	config += `		name = "FNF-input"` + "\n"
	config += `		description = "IPv4-NetFlow"` + "\n"
	config += `		match_ipv4_source_address = true` + "\n"
	config += `		match_ipv4_destination_address = true` + "\n"
	config += `		match_ipv4_protocol = true` + "\n"
	config += `		match_ipv4_tos = true` + "\n"
	config += `		match_transport_source_port = true` + "\n"
	config += `		match_transport_destination_port = true` + "\n"
	config += `		match_interface_input = true` + "\n"
	config += `		match_flow_direction = true` + "\n"
	config += `		collect_interface_output = true` + "\n"
	config += `		collect_counter_bytes_long = true` + "\n"
	config += `		collect_counter_packets_long = true` + "\n"
	config += `		collect_transport_tcp_flags = true` + "\n"
	config += `		collect_timestamp_absolute_first = true` + "\n"
	config += `		collect_timestamp_absolute_last = true` + "\n"
	config += `	}]` + "\n"
	config += `	flow_exporter = [{` + "\n"
	config += `		name = "Scrutinizer"` + "\n"
	config += `		description = "Export to Scrutinizer"` + "\n"
	config += `		destination_destination_choice_ipdest_ipdest_ip = "1.1.1.1"` + "\n"
	config += `		source_loopback = 1474` + "\n"
	config += `		transport_udp = 655` + "\n"
	config += `		template_data_timeout = 60` + "\n"
	config += `	}]` + "\n"
	config += `	flow_monitor = [{` + "\n"
	config += `		name = "Scrut_mon_input"` + "\n"
	config += `		description = "IPv4 FNF ingress exports"` + "\n"
	config += `		flow_monitor_exporter = [{` + "\n"
	config += `			name = "Scrutinizer"` + "\n"
	config += `		}]` + "\n"
	config += `		cache_timeout_active = 60` + "\n"
	config += `		record_type = "FNF_Input"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxe_flow_monitor" "test" {
			depends_on = [iosxe_flow_monitor.test]
		}
	`
	return config
}
