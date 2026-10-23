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

// Hand-written acceptance and unit tests for the AP tag assignment resources (not generated).

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/netascode/go-netconf"
	"github.com/netascode/xmldot"
)

const (
	testAPTagsResource = "iosxe_wireless_ap_tags.test"
	testAPTagsMAC1     = "00:53:00:03:00:11"
	testAPTagsMAC2     = "00:53:00:03:00:12"
	testAPTagsMAC3     = "00:53:00:03:00:13"
	testAPTagsMACOOB   = "00:53:00:03:00:1f"
)

// testAccIosxeWirelessAPTagsUnmanagedConfig creates an ap-tag entry outside of
// iosxe_wireless_ap_tags, standing in for an entry configured out of band.
func testAccIosxeWirelessAPTagsUnmanagedConfig() string {
	return fmt.Sprintf(`
resource "iosxe_yang" "unmanaged" {
  path = "/Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-tags/ap-tag[ap-mac=%s]"
  attributes = {
    ap-mac   = "%s"
    site-tag = "S3-ST-UNMANAGED"
  }
}
`, testAPTagsMACOOB, testAPTagsMACOOB)
}

func testAccIosxeWirelessAPTagsBulkConfig(items string) string {
	return testAccIosxeWirelessAPTagsUnmanagedConfig() + fmt.Sprintf(`
resource "iosxe_wireless_ap_tags" "test" {
  items = {
%s
  }
  depends_on = [iosxe_yang.unmanaged]
}
`, items)
}

const testAPTagsItemsAll = `
    "00:53:00:03:00:11" = {
      policy_tag = "S3-PT-TEST-1"
      site_tag   = "S3-ST-TEST-1"
      rf_tag     = "S3-RT-TEST-1"
    }
    "00:53:00:03:00:12" = {
      policy_tag = "S3-PT-TEST-2"
      site_tag   = "S3-ST-TEST-2"
      rf_tag     = "S3-RT-TEST-2"
    }
    "00:53:00:03:00:13" = {
      site_tag = "S3-ST-TEST-3"
    }
`

// MAC 2 removed, rf_tag removed from MAC 1.
const testAPTagsItemsReduced = `
    "00:53:00:03:00:11" = {
      policy_tag = "S3-PT-TEST-1"
      site_tag   = "S3-ST-TEST-1"
    }
    "00:53:00:03:00:13" = {
      site_tag = "S3-ST-TEST-3"
    }
`

// testAccCheckYangAttr verifies the attributes the iosxe_yang resource read back from the device.
func testAccCheckYangAttr(name, key, value string) resource.TestCheckFunc {
	return resource.TestCheckResourceAttr(name, "attributes."+key, value)
}

func TestAccIosxeWirelessAPTagsBulk(t *testing.T) {
	if os.Getenv("C9800") == "" {
		t.Skip("skipping test, set environment variable C9800")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckWirelessAPTagsDestroyed(testAPTagsMAC1, testAPTagsMAC2, testAPTagsMAC3),
		Steps: []resource.TestStep{
			// Create three items in one apply, next to an entry managed outside the bulk resource.
			{
				Config: testAccIosxeWirelessAPTagsBulkConfig(testAPTagsItemsAll),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(testAPTagsResource, "items.%", "3"),
					resource.TestCheckResourceAttr(testAPTagsResource, "items."+testAPTagsMAC1+".rf_tag", "S3-RT-TEST-1"),
					resource.TestCheckResourceAttr(testAPTagsResource, "items."+testAPTagsMAC2+".policy_tag", "S3-PT-TEST-2"),
					resource.TestCheckResourceAttr(testAPTagsResource, "items."+testAPTagsMAC3+".site_tag", "S3-ST-TEST-3"),
					resource.TestCheckNoResourceAttr(testAPTagsResource, "items."+testAPTagsMACOOB+".site_tag"),
				),
			},
			// The unmanaged entry causes no drift.
			{
				Config:   testAccIosxeWirelessAPTagsBulkConfig(testAPTagsItemsAll),
				PlanOnly: true,
			},
			// Remove one item and one attribute of another item; nothing else changes.
			{
				Config: testAccIosxeWirelessAPTagsBulkConfig(testAPTagsItemsReduced),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(testAPTagsResource, plancheck.ResourceActionUpdate),
						plancheck.ExpectResourceAction("iosxe_yang.unmanaged", plancheck.ResourceActionNoop),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(testAPTagsResource, "items.%", "2"),
					resource.TestCheckNoResourceAttr(testAPTagsResource, "items."+testAPTagsMAC1+".rf_tag"),
					resource.TestCheckResourceAttr(testAPTagsResource, "items."+testAPTagsMAC1+".site_tag", "S3-ST-TEST-1"),
					resource.TestCheckResourceAttr(testAPTagsResource, "items."+testAPTagsMAC3+".site_tag", "S3-ST-TEST-3"),
					testAccCheckWirelessAPTagsDestroyed(testAPTagsMAC2),
					testAccCheckWirelessAPTagLeafAbsent(testAPTagsMAC1, "rf-tag"),
				),
			},
			// The reduced configuration is stable.
			{
				Config:   testAccIosxeWirelessAPTagsBulkConfig(testAPTagsItemsReduced),
				PlanOnly: true,
			},
			// Import reads every entry below ap-tags, including the unmanaged one.
			{
				ResourceName:      testAPTagsResource,
				ImportState:       true,
				ImportStateIdFunc: func(*terraform.State) (string, error) { return "", nil },
				ImportStateCheck: func(states []*terraform.InstanceState) error {
					for _, mac := range []string{testAPTagsMAC1, testAPTagsMAC3, testAPTagsMACOOB} {
						if _, ok := states[0].Attributes["items."+mac+".site_tag"]; !ok {
							return fmt.Errorf("imported state is missing item %s", mac)
						}
					}
					return nil
				},
			},
			// Destroy only the bulk resource: the unmanaged entry survives.
			{
				Config: testAccIosxeWirelessAPTagsUnmanagedConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWirelessAPTagsDestroyed(testAPTagsMAC1, testAPTagsMAC3),
					testAccCheckYangAttr("iosxe_yang.unmanaged", "site-tag", "S3-ST-UNMANAGED"),
				),
			},
		},
	})
}

// testAccIosxeWirelessAPTagReadConfig reads an ap-tag entry through the data source.
func testAccCheckWirelessAPTagsDestroyed(macs ...string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, mac := range macs {
			exists, _, err := testAccWirelessAPTagGet(mac)
			if err != nil {
				return err
			}
			if exists {
				return fmt.Errorf("ap-tag %s still exists on the device", mac)
			}
		}
		return nil
	}
}

func testAccCheckWirelessAPTagLeafAbsent(mac, leaf string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		exists, leaves, err := testAccWirelessAPTagGet(mac)
		if err != nil {
			return err
		}
		if !exists {
			return fmt.Errorf("ap-tag %s does not exist on the device", mac)
		}
		if _, ok := leaves[leaf]; ok {
			return fmt.Errorf("ap-tag %s still has %s on the device", mac, leaf)
		}
		return nil
	}
}

const testTSPResource = "iosxe_wireless_ap_tag_source_priority.test"

func testAccIosxeWirelessAPTagSourcePriorityOrderConfig(source2, source3 string) string {
	return fmt.Sprintf(`
resource "iosxe_wireless_ap_tag_source_priority" "test" {
  priorities = [
    { priority = 2, source = "%s" },
    { priority = 3, source = "%s" },
  ]
}
`, source2, source3)
}

func TestAccIosxeWirelessAPTagSourcePrioritySwap(t *testing.T) {
	if os.Getenv("C9800") == "" {
		t.Skip("skipping test, set environment variable C9800")
	}
	factory := map[int64]string{2: "tag-source-filter", 3: "tag-source-ap-pnp"}
	swapped := map[int64]string{2: "tag-source-ap-pnp", 3: "tag-source-filter"}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Destroy restores the factory order rather than deleting the priorities.
		CheckDestroy: testAccCheckWirelessAPTagSourcePriorities(factory),
		Steps: []resource.TestStep{
			// AP (PnP) before filter: requires the two-step write.
			{
				Config: testAccIosxeWirelessAPTagSourcePriorityOrderConfig("tag-source-ap-pnp", "tag-source-filter"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(testTSPResource, "priorities.#", "2"),
					testAccCheckWirelessAPTagSourcePriorities(swapped),
				),
			},
			{
				Config:   testAccIosxeWirelessAPTagSourcePriorityOrderConfig("tag-source-ap-pnp", "tag-source-filter"),
				PlanOnly: true,
			},
			// Import reads only the configurable priorities 2 and 3.
			{
				ResourceName:      testTSPResource,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(*terraform.State) (string, error) { return "", nil },
			},
			// Swap back to the factory order.
			{
				Config: testAccIosxeWirelessAPTagSourcePriorityOrderConfig("tag-source-filter", "tag-source-ap-pnp"),
				Check:  testAccCheckWirelessAPTagSourcePriorities(factory),
			},
			{
				Config:   testAccIosxeWirelessAPTagSourcePriorityOrderConfig("tag-source-filter", "tag-source-ap-pnp"),
				PlanOnly: true,
			},
			// Swap again, so that destroy has to restore the factory order with the two-step write.
			{
				Config: testAccIosxeWirelessAPTagSourcePriorityOrderConfig("tag-source-ap-pnp", "tag-source-filter"),
				Check:  testAccCheckWirelessAPTagSourcePriorities(swapped),
			},
		},
	})
}

func testAccCheckWirelessAPTagSourcePriorities(want map[int64]string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		got, err := testAccWirelessAPTagSourcePrioritiesGet()
		if err != nil {
			return err
		}
		for priority, source := range want {
			if got[priority] != source {
				return fmt.Errorf("tag source priority %d: expected %q, got %q (table %v)", priority, source, got[priority], got)
			}
		}
		// The fixed priorities are never touched.
		fixed := map[int64]string{0: "tag-source-static", 1: "tag-source-location", 4: "tag-source-default"}
		for priority, source := range fixed {
			if got[priority] != source {
				return fmt.Errorf("fixed tag source priority %d changed: expected %q, got %q", priority, source, got[priority])
			}
		}
		return nil
	}
}

func TestWirelessAPTagSourcePriorityTarget(t *testing.T) {
	item := func(p int64, s string) WirelessAPTagSourcePriorityPriorities {
		return WirelessAPTagSourcePriorityPriorities{Priority: types.Int64Value(p), Source: types.StringValue(s)}
	}
	cases := []struct {
		name  string
		items []WirelessAPTagSourcePriorityPriorities
		want  map[int64]string
	}{
		{"empty", nil, map[int64]string{2: "tag-source-filter", 3: "tag-source-ap-pnp"}},
		{"factory", []WirelessAPTagSourcePriorityPriorities{item(2, "tag-source-filter"), item(3, "tag-source-ap-pnp")}, map[int64]string{2: "tag-source-filter", 3: "tag-source-ap-pnp"}},
		{"swapped", []WirelessAPTagSourcePriorityPriorities{item(2, "tag-source-ap-pnp"), item(3, "tag-source-filter")}, map[int64]string{2: "tag-source-ap-pnp", 3: "tag-source-filter"}},
		{"only 2 swapped", []WirelessAPTagSourcePriorityPriorities{item(2, "tag-source-ap-pnp")}, map[int64]string{2: "tag-source-ap-pnp", 3: "tag-source-filter"}},
		{"only 3 swapped", []WirelessAPTagSourcePriorityPriorities{item(3, "tag-source-filter")}, map[int64]string{2: "tag-source-ap-pnp", 3: "tag-source-filter"}},
	}
	for _, c := range cases {
		if got := tagSourcePriorityTarget(c.items); !reflect.DeepEqual(got, c.want) {
			t.Errorf("%s: got %v, want %v", c.name, got, c.want)
		}
	}
}

func TestWirelessAPTagSourcePrioritiesToRemove(t *testing.T) {
	factory := map[int64]string{0: "tag-source-static", 1: "tag-source-location", 2: "tag-source-filter", 3: "tag-source-ap-pnp", 4: "tag-source-default"}
	swapped := map[int64]string{2: "tag-source-ap-pnp", 3: "tag-source-filter"}
	if got := tagSourcePrioritiesToRemove(factory, swapped); !reflect.DeepEqual(got, []int64{2, 3}) {
		t.Errorf("swap: got %v", got)
	}
	if got := tagSourcePrioritiesToRemove(factory, map[int64]string{2: "tag-source-filter", 3: "tag-source-ap-pnp"}); got != nil {
		t.Errorf("no change: got %v", got)
	}
	if got := tagSourcePrioritiesToRemove(map[int64]string{2: "tag-source-filter"}, swapped); !reflect.DeepEqual(got, []int64{2}) {
		t.Errorf("partial: got %v", got)
	}
}

func TestWirelessAPTagSourcePriorityConfigErrors(t *testing.T) {
	item := func(p int64, s string) WirelessAPTagSourcePriorityPriorities {
		return WirelessAPTagSourcePriorityPriorities{Priority: types.Int64Value(p), Source: types.StringValue(s)}
	}
	cases := []struct {
		name  string
		items []WirelessAPTagSourcePriorityPriorities
		want  int
	}{
		{"empty", nil, 0},
		{"factory", []WirelessAPTagSourcePriorityPriorities{item(2, "tag-source-filter"), item(3, "tag-source-ap-pnp")}, 0},
		{"swapped", []WirelessAPTagSourcePriorityPriorities{item(2, "tag-source-ap-pnp"), item(3, "tag-source-filter")}, 0},
		{"single", []WirelessAPTagSourcePriorityPriorities{item(2, "tag-source-ap-pnp")}, 0},
		{"duplicate source", []WirelessAPTagSourcePriorityPriorities{item(2, "tag-source-filter"), item(3, "tag-source-filter")}, 1},
		{"duplicate priority", []WirelessAPTagSourcePriorityPriorities{item(2, "tag-source-filter"), item(2, "tag-source-ap-pnp")}, 1},
		{"duplicate both", []WirelessAPTagSourcePriorityPriorities{item(2, "tag-source-filter"), item(2, "tag-source-filter")}, 2},
		{"unknown skipped", []WirelessAPTagSourcePriorityPriorities{
			{Priority: types.Int64Value(2), Source: types.StringUnknown()},
			{Priority: types.Int64Unknown(), Source: types.StringUnknown()},
		}, 0},
	}
	for _, c := range cases {
		if got := tagSourcePriorityConfigErrors(c.items); len(got) != c.want {
			t.Errorf("%s: got %d errors %v, want %d", c.name, len(got), got, c.want)
		}
	}
}

// TestAccIosxeWirelessAPTagSourcePriorityDuplicate checks that a duplicate source or priority is
// rejected at plan time, before any edit-config is sent. It needs no device.
func TestAccIosxeWirelessAPTagSourcePriorityDuplicate(t *testing.T) {
	if os.Getenv("C9800") == "" {
		t.Skip("skipping test, set environment variable C9800")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccIosxeWirelessAPTagSourcePriorityOrderConfig("tag-source-filter", "tag-source-filter"),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`assigned to more than one priority`),
			},
			{
				Config: `
resource "iosxe_wireless_ap_tag_source_priority" "test" {
  priorities = [
    { priority = 2, source = "tag-source-filter" },
    { priority = 2, source = "tag-source-ap-pnp" },
  ]
}
`,
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`Priority 2 is configured more than once`),
			},
		},
	})
}

// TestWirelessAPTagsListKeysOver1000 documents how the bulk import read enumerates more than 1000
// ap-tag entries. ListXPathKeys counts siblings with xmldot, whose match collection is capped at
// 1000 results, so only the first 1000 entries are returned. This is the known bulk limitation
// tracked in https://github.com/CiscoDevNet/terraform-provider-iosxe/issues/619; update the
// expectation when it is fixed.
func TestWirelessAPTagsListKeysOver1000(t *testing.T) {
	const n = 1001
	var sb strings.Builder
	sb.WriteString(`<root><data><ap-cfg-data xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-wireless-ap-cfg"><ap-tags>`)
	for i := 0; i < n; i++ {
		fmt.Fprintf(&sb, "<ap-tag><ap-mac>00:53:00:%02x:%02x:01</ap-mac><site-tag>S3-ST</site-tag></ap-tag>", i/256, i%256)
	}
	sb.WriteString(`</ap-tags></ap-cfg-data></data></root>`)
	res := xmldot.Get(sb.String(), "root")
	var data WirelessAPTags
	keys := helpers.ListXPathKeys(res, "data"+data.getXPath(), "ap-tag", []string{"ap-mac"})
	if len(keys) != 1000 {
		t.Fatalf("got %d of %d keys; the 1000-entry cap (issue 619) changed, update this test and the iosxe_wireless_ap_tags description", len(keys), n)
	}
}

// testAccIosxeWirelessAPFilterPriorityOOBConfig configures two filters and, through iosxe_yang, a
// filter priority standing in for one configured outside Terraform. With managed set, the
// iosxe_wireless_ap_filter_priority resource manages a second priority.
func testAccIosxeWirelessAPFilterPriorityOOBConfig(managed bool) string {
	config := `
resource "iosxe_wireless_ap_filter" "managed" {
  name          = "S3-AF-PRIO-MANAGED"
  ap_name_regex = "^S3-AP-PRIO-M-.*"
  tag_site      = "default-site-tag"
}

resource "iosxe_wireless_ap_filter" "oob" {
  name          = "S3-AF-PRIO-OOB"
  ap_name_regex = "^S3-AP-PRIO-O-.*"
  tag_site      = "default-site-tag"
}

resource "iosxe_yang" "oob_priority" {
  path = "/Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-rule-priority-cfg-entries/ap-rule-priority-cfg-entry[priority=1011]"
  attributes = {
    priority    = "1011"
    filter-name = "S3-AF-PRIO-OOB"
  }
  depends_on = [iosxe_wireless_ap_filter.oob]
}
`
	if managed {
		config += `
resource "iosxe_wireless_ap_filter_priority" "test" {
  priorities = [{ priority = 1010, filter_name = "S3-AF-PRIO-MANAGED" }]
  depends_on = [iosxe_wireless_ap_filter.managed, iosxe_yang.oob_priority]
}
`
	}
	return config
}

// TestAccIosxeWirelessAPFilterPriorityKeepsUnmanaged checks that destroying
// iosxe_wireless_ap_filter_priority removes only its own priorities and leaves a priority configured
// outside Terraform in place.
func TestAccIosxeWirelessAPFilterPriorityKeepsUnmanaged(t *testing.T) {
	if os.Getenv("C9800") == "" {
		t.Skip("skipping test, set environment variable C9800")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeWirelessAPFilterPriorityOOBConfig(true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWirelessAPFilterPriority(1010, "S3-AF-PRIO-MANAGED"),
					testAccCheckWirelessAPFilterPriority(1011, "S3-AF-PRIO-OOB"),
				),
			},
			// Destroy only iosxe_wireless_ap_filter_priority.
			{
				Config: testAccIosxeWirelessAPFilterPriorityOOBConfig(false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWirelessAPFilterPriority(1010, ""),
					testAccCheckWirelessAPFilterPriority(1011, "S3-AF-PRIO-OOB"),
				),
			},
		},
	})
}

// testAccCheckWirelessAPFilterPriority checks the filter of an AP filter priority on the device. An
// empty filter means the priority must not exist.
func testAccCheckWirelessAPFilterPriority(priority int64, filter string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r, err := testAccWirelessAPCfgGet(fmt.Sprintf("/Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-rule-priority-cfg-entries/ap-rule-priority-cfg-entry[priority=%d]", priority))
		if err != nil {
			return err
		}
		if filter == "" {
			if r.Exists() {
				return fmt.Errorf("AP filter priority %d still exists", priority)
			}
			return nil
		}
		if got := r.Get("filter-name").String(); got != filter {
			return fmt.Errorf("AP filter priority %d: expected filter %q, got %q", priority, filter, got)
		}
		return nil
	}
}

// testAccWirelessAPCfgGet reads a subtree of ap-cfg-data from the device with a dedicated NETCONF
// session, independent of the provider under test.
func testAccWirelessAPCfgGet(xPath string) (xmldot.Result, error) {
	host, port := parseHostPort(os.Getenv("IOSXE_HOST"), 830)
	opts := []func(*netconf.Client){
		netconf.Username(os.Getenv("IOSXE_USERNAME")),
		netconf.Password(os.Getenv("IOSXE_PASSWORD")),
		netconf.Port(port),
		netconf.InsecureSkipHostKeyVerification(),
	}
	c, err := netconf.NewClient(host, opts...)
	if err != nil {
		return xmldot.Result{}, err
	}
	defer c.Close()
	if err := c.Open(); err != nil {
		return xmldot.Result{}, err
	}
	res, err := c.GetConfig(context.Background(), "running", helpers.GetXpathFilter(xPath))
	if err != nil {
		return xmldot.Result{}, err
	}
	return helpers.GetFromXPath(res.Res, "data"+xPath), nil
}

// testAccWirelessAPTagGet returns whether the ap-tag entry exists and the names of its leaves.
func testAccWirelessAPTagGet(mac string) (bool, map[string]bool, error) {
	r, err := testAccWirelessAPCfgGet(fmt.Sprintf("/Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-tags/ap-tag[ap-mac=%s]", mac))
	if err != nil || !r.Exists() {
		return false, nil, err
	}
	leaves := map[string]bool{}
	for _, leaf := range []string{"policy-tag", "site-tag", "rf-tag"} {
		if r.Get(leaf).Exists() {
			leaves[leaf] = true
		}
	}
	return true, leaves, nil
}

// testAccWirelessAPTagSourcePrioritiesGet returns the complete priority -> source table.
func testAccWirelessAPTagSourcePrioritiesGet() (map[int64]string, error) {
	r, err := testAccWirelessAPCfgGet("/Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/tag-source-priority-configs/tag-source-priority-config")
	if err != nil {
		return nil, err
	}
	got := map[int64]string{}
	r.ForEach(func(_ int, v xmldot.Result) bool {
		got[v.Get("priority").Int()] = v.Get("tag-src").String()
		return true
	})
	return got, nil
}
