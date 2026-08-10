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
	"strings"
	"testing"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-netconf"
)

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

// sequentialDeletePathsXML reproduces the pre-batching behaviour of
// addDeletePathsXML: one helpers.RemoveFromXPath call per VLAN.
func sequentialDeletePathsXML(data *SpanningTree) string {
	b := netconf.NewBody("")
	for i := range data.Vlans {
		if !data.Vlans[i].Priority.IsNull() {
			predicates := fmt.Sprintf("[id='%s']", data.Vlans[i].Id.ValueString())
			b = helpers.RemoveFromXPath(b, fmt.Sprintf(data.getXPath()+"/Cisco-IOS-XE-spanning-tree:vlan%v/priority", predicates))
		}
	}
	b = helpers.CleanupRedundantRemoveOperations(b)
	return b.Res()
}

// TestAddDeletePathsXML_MatchesSequential pins byte-for-byte equivalence with
// the per-item path at sizes below xmldot's MaxWildcardResults (1000) cap,
// where that path is still correct. Kept small because it is quadratic.
func TestAddDeletePathsXML_MatchesSequential(t *testing.T) {
	ctx := context.Background()
	for _, n := range []int{1, 2, 50} {
		got := spanningTreeWithVlans(n).addDeletePathsXML(ctx, "")
		want := sequentialDeletePathsXML(spanningTreeWithVlans(n))
		if got != want {
			t.Errorf("n=%d: batched output differs from sequential\n batched:    %s\n sequential: %s", n, got, want)
		}
	}
}

// TestAddDeletePathsXML_NamespacesAtScale is the Delete-path counterpart to
// TestRemoveFromXPathMulti_NamespacesPastWildcardCap. Every <vlan> element must
// carry an inline xmlns, including past the 1000-element cap, or IOS-XE rejects
// the whole edit-config with unknown-element.
func TestAddDeletePathsXML_NamespacesAtScale(t *testing.T) {
	ctx := context.Background()
	const stNs = `<vlan xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-spanning-tree">`

	for _, n := range []int{999, 1000, 1001, 1500, 4094} {
		res := spanningTreeWithVlans(n).addDeletePathsXML(ctx, "")

		if got := strings.Count(res, "<vlan"); got != n {
			t.Errorf("n=%d: got %d vlan elements, want %d", n, got, n)
		}
		if got := strings.Count(res, stNs); got != n {
			t.Errorf("n=%d: got %d vlan elements carrying xmlns, want %d (%d would be silently rejected by the device)",
				n, got, n, n-got)
		}
		if got := strings.Count(res, `<priority operation="remove">`); got != n {
			t.Errorf("n=%d: got %d priority remove operations, want %d", n, got, n)
		}
	}
}
