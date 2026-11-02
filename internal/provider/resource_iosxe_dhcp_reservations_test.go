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
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// Scenario tests for DHCP manual bindings (pools with host + client-identifier
// or hardware-address) and ip dhcp excluded-address. Every removal is verified
// against the device through the iosxe_dhcp data source, not only state.

const dhcpLiveDataSource = "data.iosxe_dhcp.live"

// dhcpLiveOnlyConfig reads the device without managing anything.
const dhcpLiveOnlyConfig = "data \"iosxe_dhcp\" \"live\" {}\n"

type dhcpTestPool struct {
	name, network, mask, host, clientID, hwAddr string
}

func (p dhcpTestPool) hcl() string {
	var b strings.Builder
	b.WriteString("    {\n")
	fmt.Fprintf(&b, "      name = %q\n", p.name)
	if p.network != "" {
		fmt.Fprintf(&b, "      network_number = %q\n      network_mask = %q\n", p.network, p.mask)
	}
	if p.host != "" {
		fmt.Fprintf(&b, "      host_number = %q\n      host_mask = %q\n", p.host, p.mask)
	}
	if p.clientID != "" {
		fmt.Fprintf(&b, "      client_identifier = %q\n", p.clientID)
	}
	if p.hwAddr != "" {
		fmt.Fprintf(&b, "      hardware_address = %q\n", p.hwAddr)
	}
	b.WriteString("    },\n")
	return b.String()
}

func testAccIosxeDHCPReservationsConfig(pools ...dhcpTestPool) string {
	var b strings.Builder
	b.WriteString("resource \"iosxe_dhcp\" \"test\" {\n  pools = [\n")
	for _, p := range pools {
		b.WriteString(p.hcl())
	}
	b.WriteString("  ]\n}\n\n")
	b.WriteString("data \"iosxe_dhcp\" \"live\" {\n  depends_on = [iosxe_dhcp.test]\n}\n")
	return b.String()
}

// liveAttrs returns the attributes the data source read from the device.
func liveAttrs(s *terraform.State) (map[string]string, error) {
	rs, ok := s.RootModule().Resources[dhcpLiveDataSource]
	if !ok {
		return nil, fmt.Errorf("%s not found in state", dhcpLiveDataSource)
	}
	return rs.Primary.Attributes, nil
}

// findLiveListItem returns the index prefix (e.g. "pools.2.") of the first
// item in list whose attributes match all of want, or "" if none does.
func findLiveListItem(attrs map[string]string, list string, want map[string]string) string {
	for i := 0; ; i++ {
		prefix := fmt.Sprintf("%s.%d.", list, i)
		found := false
		for k := range attrs {
			if strings.HasPrefix(k, prefix) {
				found = true
				break
			}
		}
		if !found {
			return ""
		}
		match := true
		for k, v := range want {
			if attrs[prefix+k] != v {
				match = false
				break
			}
		}
		if match {
			return prefix
		}
	}
}

// testCheckDeviceDHCPPool asserts that the device has a pool with the given
// name, and that each attribute in want matches. An empty value in want
// asserts the attribute is unset on the device.
func testCheckDeviceDHCPPool(name string, want map[string]string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		attrs, err := liveAttrs(s)
		if err != nil {
			return err
		}
		prefix := findLiveListItem(attrs, "pools", map[string]string{"name": name})
		if prefix == "" {
			return fmt.Errorf("pool %q not found on device", name)
		}
		for k, v := range want {
			if got := attrs[prefix+k]; got != v {
				return fmt.Errorf("pool %q on device: %s = %q, want %q", name, k, got, v)
			}
		}
		return nil
	}
}

// testCheckDeviceListItemAbsent asserts the device has no item in list
// matching want.
func testCheckDeviceListItemAbsent(list string, want map[string]string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		attrs, err := liveAttrs(s)
		if err != nil {
			return err
		}
		if prefix := findLiveListItem(attrs, list, want); prefix != "" {
			return fmt.Errorf("%s item %v still present on device at %s", list, want, prefix)
		}
		return nil
	}
}

func testCheckDeviceDHCPPoolAbsent(name string) resource.TestCheckFunc {
	return testCheckDeviceListItemAbsent("pools", map[string]string{"name": name})
}

func TestAccIosxeDHCPReservations(t *testing.T) {
	network := dhcpTestPool{name: "TFACC-RES-NET", network: "10.81.1.0", mask: "255.255.255.0"}
	cid := dhcpTestPool{name: "TFACC-RES-CID", host: "10.81.1.11", mask: "255.255.255.0", clientID: "01b0.e45c.d9d6.55"}
	hw := dhcpTestPool{name: "TFACC-RES-HW", host: "10.81.1.12", mask: "255.255.255.0", hwAddr: "a8b1.3ba9.f06e"}

	cidNewID := cid
	cidNewID.clientID = "01a4.307a.f331.12"

	cidNewIP := cidNewID
	cidNewIP.host = "10.81.1.21"

	renamed := cidNewIP
	renamed.name = "TFACC-RES-CID-RENAMED"

	// Swap host addresses between the two reservations in a single apply.
	// IOS XE applies the commit as CLI and refuses a host address that is
	// still bound to another pool, so this is expected to fail atomically.
	swapA := renamed
	swapA.host = hw.host
	swapB := hw
	swapB.host = renamed.host

	// Switch the renamed pool from client-identifier to hardware-address.
	switched := renamed
	switched.clientID = ""
	switched.hwAddr = "b8b4.0980.3b76"

	// And back from hardware-address to client-identifier.
	switchedBack := switched
	switchedBack.hwAddr = ""
	switchedBack.clientID = "01b8.b409.803b.76"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Create a network pool and one reservation of each kind.
				Config: testAccIosxeDHCPReservationsConfig(network, cid, hw),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_dhcp.test", "pools.1.client_identifier", cid.clientID),
					resource.TestCheckResourceAttr("iosxe_dhcp.test", "pools.2.hardware_address", hw.hwAddr),
					testCheckDeviceDHCPPool(network.name, map[string]string{"network_number": network.network, "client_identifier": "", "hardware_address": ""}),
					testCheckDeviceDHCPPool(cid.name, map[string]string{"host_number": cid.host, "host_mask": cid.mask, "client_identifier": cid.clientID, "hardware_address": ""}),
					testCheckDeviceDHCPPool(hw.name, map[string]string{"host_number": hw.host, "host_mask": hw.mask, "hardware_address": hw.hwAddr, "client_identifier": ""}),
				),
			},
			{
				// Re-applying the same config must be a no-op.
				Config:   testAccIosxeDHCPReservationsConfig(network, cid, hw),
				PlanOnly: true,
			},
			{
				// Change the client identifier in place.
				Config: testAccIosxeDHCPReservationsConfig(network, cidNewID, hw),
				Check:  testCheckDeviceDHCPPool(cid.name, map[string]string{"client_identifier": cidNewID.clientID, "host_number": cid.host}),
			},
			{
				// Change the reserved address in place.
				Config: testAccIosxeDHCPReservationsConfig(network, cidNewIP, hw),
				Check:  testCheckDeviceDHCPPool(cid.name, map[string]string{"host_number": cidNewIP.host, "client_identifier": cidNewID.clientID}),
			},
			{
				// Rename: the old pool must be removed from the device and the
				// new one created with the same binding, in one apply.
				Config: testAccIosxeDHCPReservationsConfig(network, renamed, hw),
				Check: resource.ComposeTestCheckFunc(
					testCheckDeviceDHCPPoolAbsent(cid.name),
					testCheckDeviceDHCPPool(renamed.name, map[string]string{"host_number": renamed.host, "client_identifier": renamed.clientID}),
				),
			},
			{
				// Swapping addresses between two reservations in one apply is
				// refused by the device.
				Config:      testAccIosxeDHCPReservationsConfig(network, swapA, swapB),
				ExpectError: regexp.MustCompile(`A binding already exists for`),
			},
			{
				// The refused commit must leave the device unchanged.
				Config: testAccIosxeDHCPReservationsConfig(network, renamed, hw),
				Check: resource.ComposeTestCheckFunc(
					testCheckDeviceDHCPPool(renamed.name, map[string]string{"host_number": renamed.host, "client_identifier": renamed.clientID}),
					testCheckDeviceDHCPPool(hw.name, map[string]string{"host_number": hw.host, "hardware_address": hw.hwAddr}),
				),
			},
			{
				// Switch a binding from client-identifier to hardware-address.
				Config: testAccIosxeDHCPReservationsConfig(network, switched, hw),
				Check:  testCheckDeviceDHCPPool(switched.name, map[string]string{"hardware_address": switched.hwAddr, "client_identifier": ""}),
			},
			{
				// And from hardware-address back to client-identifier. Pools
				// are listed in the order the device returns them (by name) so
				// the import below compares like for like.
				Config: testAccIosxeDHCPReservationsConfig(switchedBack, hw, network),
				Check:  testCheckDeviceDHCPPool(switched.name, map[string]string{"client_identifier": switchedBack.clientID, "hardware_address": ""}),
			},
			{
				ResourceName:            "iosxe_dhcp.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       iosxeDHCPImportStateIdFunc("iosxe_dhcp.test"),
				ImportStateVerifyIgnore: dhcpImportStateVerifyIgnore(3),
			},
			{
				// Remove one reservation from the list.
				Config: testAccIosxeDHCPReservationsConfig(network, switchedBack),
				Check: resource.ComposeTestCheckFunc(
					testCheckDeviceDHCPPoolAbsent(hw.name),
					testCheckDeviceDHCPPool(switchedBack.name, map[string]string{"client_identifier": switchedBack.clientID}),
				),
			},
			{
				// Destroy the resource. The data source is read during this
				// plan, before the destroy, so the device is checked next step.
				Config: dhcpLiveOnlyConfig,
			},
			{
				// Confirm nothing is left behind on the device.
				Config: dhcpLiveOnlyConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckDeviceDHCPPoolAbsent(network.name),
					testCheckDeviceDHCPPoolAbsent(switchedBack.name),
				),
			},
		},
	})
}

type dhcpTestExclusions struct {
	singles, ranges, vrfSingles, vrfRanges [][]string
}

func testAccIosxeDHCPExclusionsConfig(e dhcpTestExclusions) string {
	var b strings.Builder
	b.WriteString(`resource "iosxe_yang" "vrf" {
  path = "/Cisco-IOS-XE-native:native/vrf/definition[name=TFACC_DHCP]"
  attributes = {
    "name"                = "TFACC_DHCP"
    "address-family/ipv4" = ""
  }
}

resource "iosxe_dhcp" "test" {
  depends_on = [iosxe_yang.vrf]
`)
	list := func(name string, items [][]string, keys ...string) {
		if items == nil {
			return
		}
		fmt.Fprintf(&b, "  %s = [\n", name)
		for _, it := range items {
			b.WriteString("    {")
			for i, k := range keys {
				fmt.Fprintf(&b, " %s = %q,", k, it[i])
			}
			b.WriteString(" },\n")
		}
		b.WriteString("  ]\n")
	}
	list("excluded_addresses", e.singles, "address")
	list("excluded_address_ranges", e.ranges, "low_address", "high_address")
	list("excluded_addresses_vrf", e.vrfSingles, "vrf", "address")
	list("excluded_address_ranges_vrf", e.vrfRanges, "vrf", "low_address", "high_address")
	b.WriteString("}\n\n")
	b.WriteString("data \"iosxe_dhcp\" \"live\" {\n  depends_on = [iosxe_dhcp.test]\n}\n")
	return b.String()
}

func testCheckDeviceExclusion(present bool, list string, kv ...string) resource.TestCheckFunc {
	want := map[string]string{}
	for i := 0; i+1 < len(kv); i += 2 {
		want[kv[i]] = kv[i+1]
	}
	if !present {
		return testCheckDeviceListItemAbsent(list, want)
	}
	return func(s *terraform.State) error {
		attrs, err := liveAttrs(s)
		if err != nil {
			return err
		}
		if findLiveListItem(attrs, list, want) == "" {
			return fmt.Errorf("%s item %v not found on device", list, want)
		}
		return nil
	}
}

func TestAccIosxeDHCPExcludedAddresses(t *testing.T) {
	all := dhcpTestExclusions{
		singles:    [][]string{{"10.82.1.1"}, {"10.82.1.5"}},
		ranges:     [][]string{{"10.82.1.10", "10.82.1.20"}, {"10.82.2.10", "10.82.2.20"}},
		vrfSingles: [][]string{{"TFACC_DHCP", "10.83.1.1"}},
		vrfRanges:  [][]string{{"TFACC_DHCP", "10.83.1.10", "10.83.1.20"}},
	}
	allChecks := resource.ComposeTestCheckFunc(
		testCheckDeviceExclusion(true, "excluded_addresses", "address", "10.82.1.1"),
		testCheckDeviceExclusion(true, "excluded_addresses", "address", "10.82.1.5"),
		testCheckDeviceExclusion(true, "excluded_address_ranges", "low_address", "10.82.1.10", "high_address", "10.82.1.20"),
		testCheckDeviceExclusion(true, "excluded_address_ranges", "low_address", "10.82.2.10", "high_address", "10.82.2.20"),
		testCheckDeviceExclusion(true, "excluded_addresses_vrf", "vrf", "TFACC_DHCP", "address", "10.83.1.1"),
		testCheckDeviceExclusion(true, "excluded_address_ranges_vrf", "vrf", "TFACC_DHCP", "low_address", "10.83.1.10", "high_address", "10.83.1.20"),
	)

	// Widen one range. low and high are both list keys, so this is a
	// remove + add of the entry, not an in-place update.
	widened := all
	widened.ranges = [][]string{{"10.82.1.10", "10.82.1.30"}, {"10.82.2.10", "10.82.2.20"}}

	// Drop one single and one range, add a new range.
	reshuffled := dhcpTestExclusions{
		singles:    [][]string{{"10.82.1.5"}},
		ranges:     [][]string{{"10.82.1.10", "10.82.1.30"}, {"10.82.3.10", "10.82.3.20"}},
		vrfSingles: all.vrfSingles,
		vrfRanges:  all.vrfRanges,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeDHCPExclusionsConfig(all),
				Check:  allChecks,
			},
			{
				Config:   testAccIosxeDHCPExclusionsConfig(all),
				PlanOnly: true,
			},
			{
				ResourceName:            "iosxe_dhcp.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       iosxeDHCPImportStateIdFunc("iosxe_dhcp.test"),
				ImportStateVerifyIgnore: dhcpImportStateVerifyIgnore(0),
			},
			{
				Config: testAccIosxeDHCPExclusionsConfig(widened),
				Check: resource.ComposeTestCheckFunc(
					testCheckDeviceExclusion(false, "excluded_address_ranges", "low_address", "10.82.1.10", "high_address", "10.82.1.20"),
					testCheckDeviceExclusion(true, "excluded_address_ranges", "low_address", "10.82.1.10", "high_address", "10.82.1.30"),
				),
			},
			{
				Config: testAccIosxeDHCPExclusionsConfig(reshuffled),
				Check: resource.ComposeTestCheckFunc(
					testCheckDeviceExclusion(false, "excluded_addresses", "address", "10.82.1.1"),
					testCheckDeviceExclusion(true, "excluded_addresses", "address", "10.82.1.5"),
					testCheckDeviceExclusion(false, "excluded_address_ranges", "low_address", "10.82.2.10", "high_address", "10.82.2.20"),
					testCheckDeviceExclusion(true, "excluded_address_ranges", "low_address", "10.82.3.10", "high_address", "10.82.3.20"),
				),
			},
			{
				// Remove every exclusion list from the config.
				Config: testAccIosxeDHCPExclusionsConfig(dhcpTestExclusions{}),
				Check: resource.ComposeTestCheckFunc(
					testCheckDeviceExclusion(false, "excluded_addresses", "address", "10.82.1.5"),
					testCheckDeviceExclusion(false, "excluded_address_ranges", "low_address", "10.82.1.10", "high_address", "10.82.1.30"),
					testCheckDeviceExclusion(false, "excluded_address_ranges", "low_address", "10.82.3.10", "high_address", "10.82.3.20"),
					testCheckDeviceExclusion(false, "excluded_addresses_vrf", "vrf", "TFACC_DHCP", "address", "10.83.1.1"),
					testCheckDeviceExclusion(false, "excluded_address_ranges_vrf", "vrf", "TFACC_DHCP", "low_address", "10.83.1.10", "high_address", "10.83.1.20"),
				),
			},
			{
				// Re-add everything, then destroy and confirm the device is clean.
				Config: testAccIosxeDHCPExclusionsConfig(all),
				Check:  allChecks,
			},
			{
				// Destroy, then confirm the device is clean in the next step.
				Config: dhcpLiveOnlyConfig,
			},
			{
				Config: dhcpLiveOnlyConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckDeviceExclusion(false, "excluded_addresses", "address", "10.82.1.1"),
					testCheckDeviceExclusion(false, "excluded_address_ranges", "low_address", "10.82.1.10", "high_address", "10.82.1.20"),
					testCheckDeviceExclusion(false, "excluded_addresses_vrf", "vrf", "TFACC_DHCP", "address", "10.83.1.1"),
					testCheckDeviceExclusion(false, "excluded_address_ranges_vrf", "vrf", "TFACC_DHCP", "low_address", "10.83.1.10", "high_address", "10.83.1.20"),
				),
			},
		},
	})
}

// dhcpImportStateVerifyIgnore lists attributes that import reads back as
// false when the config leaves them unset, mirroring the generated test.
func dhcpImportStateVerifyIgnore(pools int) []string {
	ignore := []string{
		"relay_bootp_ignore",
		"relay_information_option_default",
		"relay_information_option_vpn",
		"relay_information_trust_all",
		"snooping",
		"snooping_information_option",
		"snooping_information_option_allow_untrusted",
		"snooping_information_option_format_remote_id_hostname",
	}
	for i := 0; i < pools; i++ {
		for _, a := range []string{"lease_infinite", "utilization_mark_high_log", "utilization_mark_low_log"} {
			ignore = append(ignore, fmt.Sprintf("pools.%d.%s", i, a))
		}
	}
	return ignore
}

// testAccIosxeDHCPInvalidPool runs one config that must fail and then checks
// the failed apply left nothing on the device.
func testAccIosxeDHCPInvalidPool(t *testing.T, pool dhcpTestPool, expect *regexp.Regexp) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccIosxeDHCPReservationsConfig(pool),
				ExpectError: expect,
			},
			{
				// Destroy anything the step above may have created.
				Config: dhcpLiveOnlyConfig,
			},
			{
				Config: dhcpLiveOnlyConfig,
				Check:  testCheckDeviceDHCPPoolAbsent(pool.name),
			},
		},
	})
}

func TestAccIosxeDHCPReservationsInvalidNetworkPool(t *testing.T) {
	// client-identifier is not allowed on a network pool.
	testAccIosxeDHCPInvalidPool(t,
		dhcpTestPool{name: "TFACC-INV-NET", network: "10.84.1.0", mask: "255.255.255.0", clientID: "01b0.e45c.d9d6.55"},
		regexp.MustCompile(`(?is)client-identifier|network`))
}

func TestAccIosxeDHCPReservationsInvalidBoth(t *testing.T) {
	// client-identifier and hardware-address are mutually exclusive.
	testAccIosxeDHCPInvalidPool(t,
		dhcpTestPool{name: "TFACC-INV-BOTH", host: "10.84.2.11", mask: "255.255.255.0", clientID: "01b0.e45c.d9d6.55", hwAddr: "a8b1.3ba9.f06e"},
		regexp.MustCompile(`Conflicting DHCP pool attributes`))
}

func TestAccIosxeDHCPReservationsInvalidFormat(t *testing.T) {
	// Malformed client identifier.
	testAccIosxeDHCPInvalidPool(t,
		dhcpTestPool{name: "TFACC-INV-FMT", host: "10.84.3.11", mask: "255.255.255.0", clientID: "not-a-client-id"},
		regexp.MustCompile(`(?is)client-identifier|invalid|pattern`))
}
