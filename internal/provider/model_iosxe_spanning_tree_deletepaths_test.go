// Copyright © 2025 Cisco Systems, Inc. and its affiliates.
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
	"context"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-netconf"
)

// vlanRemovalRe matches a whole-element <vlan ...><id>N</id></vlan> removal
// fragment regardless of attribute order.
var vlanRemovalRe = regexp.MustCompile(`<vlan\s+([^>]*)>\s*<id>(\d+)</id>\s*</vlan>`)

// extractVlanRemovals returns id -> (has correct xmlns AND operation="remove"),
// independent of attribute order. Needed because RemoveFromXPath (used by
// sequentialDeletePathsXML below) and RemoveFromXPathMulti's buildRemoveFragment
// (used by the real addDeletePathsXML) serialize element attributes via
// different code paths -- one goes through Go map iteration, which does not
// guarantee order. The two are semantically identical NETCONF payloads; a
// byte-for-byte string comparison across them is not a sound test once an
// element carries more than one attribute.
func extractVlanRemovals(xmlStr string) map[string]bool {
	result := make(map[string]bool)
	for _, m := range vlanRemovalRe.FindAllStringSubmatch(xmlStr, -1) {
		attrs, id := m[1], m[2]
		hasNs := strings.Contains(attrs, `xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-spanning-tree"`)
		hasRemove := strings.Contains(attrs, `operation="remove"`)
		result[id] = hasNs && hasRemove
	}
	return result
}

func spanningTreeWithVlans(n int) *SpanningTree {
	data := &SpanningTree{}
	for i := 1; i <= n; i++ {
		data.Vlans = append(data.Vlans, SpanningTreeVlans{
			Id:       types.StringValue(fmt.Sprintf("%d", i)),
			Priority: types.Int64Value(32768),
		})
	}
	return data
}

// sequentialDeletePathsXML reproduces the pre-batching, per-item behaviour
// of addDeletePathsXML's WHOLE-ELEMENT removal: one helpers.RemoveFromXPath
// call per VLAN, targeting the <vlan> element itself rather than just its
// priority leaf.
//
// UPDATED 2026-08-10: addDeletePathsXML now removes the whole <vlan>
// element for content-bearing entries rather than just the priority leaf
// (hardware-tested safe -- see the block comment on addDeletePathsXML
// itself). This helper, and the tests below, were updated to match; they
// previously pinned the old leaf-only shape.
func sequentialDeletePathsXML(data *SpanningTree) string {
	b := netconf.NewBody("")
	for i := range data.Vlans {
		if !data.Vlans[i].Priority.IsNull() {
			predicates := fmt.Sprintf("[id='%s']", data.Vlans[i].Id.ValueString())
			b = helpers.RemoveFromXPath(b, fmt.Sprintf(data.getXPath()+"/Cisco-IOS-XE-spanning-tree:vlan%v", predicates))
		}
	}
	b = helpers.CleanupRedundantRemoveOperations(b)
	return b.Res()
}

// TestAddDeletePathsXML_MatchesSequential pins semantic equivalence
// (same VLAN ids, each whole-element-removed with correct xmlns) with the
// per-item path at sizes below xmldot's MaxWildcardResults (1000) cap,
// where that path is still correct. Kept small because it is quadratic.
// Compares semantically rather than byte-for-byte -- see extractVlanRemovals
// for why a literal string comparison isn't sound here.
func TestAddDeletePathsXML_MatchesSequential(t *testing.T) {
	ctx := context.Background()
	for _, n := range []int{1, 2, 50} {
		got := extractVlanRemovals(spanningTreeWithVlans(n).addDeletePathsXML(ctx, ""))
		want := extractVlanRemovals(sequentialDeletePathsXML(spanningTreeWithVlans(n)))

		if len(got) != n {
			t.Errorf("n=%d: batched output has %d correctly-removed vlan elements, want %d", n, len(got), n)
		}
		if len(want) != n {
			t.Errorf("n=%d: sequential reference has %d correctly-removed vlan elements, want %d (test bug if this fails)", n, len(want), n)
		}
		for id, wantOk := range want {
			gotOk, present := got[id]
			if !present {
				t.Errorf("n=%d: vlan id=%s present in sequential but missing from batched output", n, id)
				continue
			}
			if gotOk != wantOk {
				t.Errorf("n=%d: vlan id=%s attribute correctness differs: batched=%v sequential=%v", n, id, gotOk, wantOk)
			}
		}
		for id := range got {
			if _, present := want[id]; !present {
				t.Errorf("n=%d: vlan id=%s present in batched but not in sequential output", n, id)
			}
		}
	}
}

// TestAddDeletePathsXML_NamespacesAtScale is the Delete-path counterpart to
// TestRemoveFromXPathMulti_NamespacesPastWildcardCap. Every <vlan> element must
// carry an inline xmlns, including past the 1000-element cap, or IOS-XE rejects
// the whole edit-config with unknown-element.
func TestAddDeletePathsXML_NamespacesAtScale(t *testing.T) {
	ctx := context.Background()
	const stNsRemove = `<vlan xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-spanning-tree" operation="remove">`

	for _, n := range []int{999, 1000, 1001, 1500, 4094} {
		res := spanningTreeWithVlans(n).addDeletePathsXML(ctx, "")

		if got := strings.Count(res, "<vlan"); got != n {
			t.Errorf("n=%d: got %d vlan elements, want %d", n, got, n)
		}
		if got := strings.Count(res, stNsRemove); got != n {
			t.Errorf("n=%d: got %d vlan elements carrying xmlns+operation=remove, want %d (%d would be silently rejected by the device, or fail to remove the whole element)",
				n, got, n, n-got)
		}
		// Whole-element removal: no <priority> child should be emitted at
		// all -- operation="remove" lives on <vlan> itself, not on a leaf
		// underneath it. This is the behavioral change from the old
		// leaf-only shape.
		if got := strings.Count(res, "<priority"); got != 0 {
			t.Errorf("n=%d: got %d <priority> elements, want 0 -- whole-element removal should not touch the priority leaf directly", n, got)
		}
		if got := strings.Count(res, "<id>"); got != n {
			t.Errorf("n=%d: got %d <id> elements, want %d", n, got, n)
		}
	}
}
