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

package helpers

import (
	"fmt"
	"strings"
	"testing"

	"github.com/netascode/go-netconf"
	"github.com/netascode/xmldot"
)

// --- SetRawFromXPathMulti (write path) ------------------------------------

// TestSetRawFromXPathMulti_MatchesSequential pins byte-for-byte equivalence
// with the pre-fix behaviour of calling SetRawFromXPath once per list item.
// Kept small because the sequential path is quadratic.
func TestSetRawFromXPathMulti_MatchesSequential(t *testing.T) {
	const xPath = "/native/spanning-tree/vlan"

	for _, n := range []int{1, 2, 50} {
		values := make([]string, 0, n)
		for i := 0; i < n; i++ {
			values = append(values, fmt.Sprintf("<id>%d</id><priority>32768</priority>", i))
		}

		var sequential netconf.Body
		for _, v := range values {
			sequential = SetRawFromXPath(sequential, xPath, v)
		}

		var batched netconf.Body
		batched = SetRawFromXPathMulti(batched, xPath, values)

		if sequential.Res() != batched.Res() {
			t.Errorf("n=%d: batched output differs from sequential\n sequential: %s\n batched:    %s",
				n, sequential.Res(), batched.Res())
		}
	}
}

// TestSetRawFromXPathMulti_MatchesSequentialWithNamespaces does the same on
// namespace-prefixed paths, at sizes below xmldot's MaxWildcardResults cap
// where the sequential path still produces correct output.
func TestSetRawFromXPathMulti_MatchesSequentialWithNamespaces(t *testing.T) {
	const xPath = "/Cisco-IOS-XE-native:native/spanning-tree/Cisco-IOS-XE-spanning-tree:vlan"

	for _, n := range []int{1, 2, 50} {
		values := make([]string, 0, n)
		for i := 0; i < n; i++ {
			values = append(values, fmt.Sprintf("<id>%d</id><priority>32768</priority>", i))
		}

		var sequential netconf.Body
		for _, v := range values {
			sequential = SetRawFromXPath(sequential, xPath, v)
		}

		var batched netconf.Body
		batched = SetRawFromXPathMulti(batched, xPath, values)

		if sequential.Res() != batched.Res() {
			t.Errorf("n=%d: batched output differs from sequential\n sequential: %s\n batched:    %s",
				n, sequential.Res(), batched.Res())
		}
	}
}

// TestSetRawFromXPathMulti_EmptyAndSingle covers the trivial fallback paths:
// an empty slice, a slice of only empty strings (which SetRawFromXPath treats
// as a no-op), and a single value.
func TestSetRawFromXPathMulti_EmptyAndSingle(t *testing.T) {
	const xPath = "/native/spanning-tree/vlan"

	var body netconf.Body
	body = SetRawFromXPathMulti(body, xPath, nil)
	if body.Res() != "" {
		t.Errorf("empty input: got non-empty body %q", body.Res())
	}

	body = netconf.Body{}
	body = SetRawFromXPathMulti(body, xPath, []string{"", "", ""})
	if body.Res() != "" {
		t.Errorf("all-empty input: got non-empty body %q", body.Res())
	}

	body = netconf.Body{}
	body = SetRawFromXPathMulti(body, xPath, []string{"<id>10</id>"})
	if got := xmldot.Get(body.Res(), "native.spanning-tree.vlan.id").String(); got != "10" {
		t.Errorf("single-item input: vlan id = %q, want %q", got, "10")
	}
}

// TestSetRawFromXPathMulti_SkipsEmptyValues verifies that empty entries mixed
// in with real ones are dropped rather than emitted as empty elements, so the
// result matches what a sequential run (which no-ops on empty) would produce.
func TestSetRawFromXPathMulti_SkipsEmptyValues(t *testing.T) {
	const xPath = "/native/spanning-tree/vlan"
	values := []string{"<id>1</id>", "", "<id>2</id>", "", "<id>3</id>"}

	var sequential netconf.Body
	for _, v := range values {
		sequential = SetRawFromXPath(sequential, xPath, v)
	}

	var batched netconf.Body
	batched = SetRawFromXPathMulti(batched, xPath, values)

	if sequential.Res() != batched.Res() {
		t.Errorf("batched output differs from sequential\n sequential: %s\n batched:    %s",
			sequential.Res(), batched.Res())
	}
	if got := xmldot.Get(batched.Res(), "native.spanning-tree.vlan.#").Int(); got != 3 {
		t.Errorf("got %d vlan elements, want 3", got)
	}
}

// TestSetRawFromXPathMulti_PreservesExistingContent verifies the single
// combined write appends to whatever is already at the parent path rather than
// replacing it — the sequential path appends, so the batched one must too.
func TestSetRawFromXPathMulti_PreservesExistingContent(t *testing.T) {
	var body netconf.Body
	body = SetRawFromXPath(body, "/native/spanning-tree/mode", "rapid-pvst")
	body = SetRawFromXPathMulti(body, "/native/spanning-tree/vlan",
		[]string{"<id>1</id>", "<id>2</id>"})

	res := body.Res()
	if got := xmldot.Get(res, "native.spanning-tree.mode").String(); got != "rapid-pvst" {
		t.Errorf("pre-existing sibling lost: mode = %q, want %q", got, "rapid-pvst")
	}
	if got := xmldot.Get(res, "native.spanning-tree.vlan.#").Int(); got != 2 {
		t.Errorf("got %d vlan elements, want 2", got)
	}
}

// TestSetRawFromXPathMulti_SingleSegmentFallback checks the documented
// fallback: single-segment paths are not batched and must still behave
// exactly like a sequential run.
func TestSetRawFromXPathMulti_SingleSegmentFallback(t *testing.T) {
	const xPath = "/vlan"
	values := []string{"<id>1</id>", "<id>2</id>", "<id>3</id>"}

	var sequential netconf.Body
	for _, v := range values {
		sequential = SetRawFromXPath(sequential, xPath, v)
	}

	var batched netconf.Body
	batched = SetRawFromXPathMulti(batched, xPath, values)

	if sequential.Res() != batched.Res() {
		t.Errorf("single-segment fallback differs from sequential\n sequential: %s\n batched:    %s",
			sequential.Res(), batched.Res())
	}
}

// TestSetRawFromXPathMulti_NamespacesPastWildcardCap is a regression test for
// the device-rejecting bug described in #619: augmentNamespaces discovers
// siblings through xmldot's counted/indexed queries, which cap at
// MaxWildcardResults (1000), so backfilling left every element past the
// 1000th without an xmlns. IOS-XE resolved those against the parent (native)
// namespace and failed the whole edit-config with "unknown-element".
// SetRawFromXPathMulti embeds the xmlns into each fragment instead.
//
// This only reproduces with namespace-prefixed xPaths, which is why the
// bare-path tests above do not catch it.
func TestSetRawFromXPathMulti_NamespacesPastWildcardCap(t *testing.T) {
	const xPath = "/Cisco-IOS-XE-native:native/spanning-tree/Cisco-IOS-XE-spanning-tree:vlan"
	const stNs = `<vlan xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-spanning-tree">`

	// 4094 is the real-world worst case: a full VLAN range being declared.
	for _, n := range []int{999, 1000, 1001, 2000, 4094} {
		values := make([]string, 0, n)
		for i := 0; i < n; i++ {
			values = append(values, fmt.Sprintf("<id>%d</id><priority>32768</priority>", i))
		}

		var body netconf.Body
		body = SetRawFromXPathMulti(body, xPath, values)
		res := body.Res()

		if got := strings.Count(res, "<vlan"); got != n {
			t.Errorf("n=%d: got %d vlan elements, want %d", n, got, n)
		}
		if got := strings.Count(res, stNs); got != n {
			t.Errorf("n=%d: got %d vlan elements carrying xmlns, want %d (%d would be silently rejected by the device)",
				n, got, n, n-got)
		}
	}
}

// --- CollectListItemsXML (read path) --------------------------------------

// buildVlanScope returns the inner content of a <spanning-tree> element
// holding n vlan entries — the shape updateFromBodyXML works with, since
// GetFromXPath(...).Raw yields an element's children, not the element itself.
func buildVlanScope(n int) string {
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteString(fmt.Sprintf("<vlan><id>%d</id><priority>32768</priority></vlan>", i))
	}
	return b.String()
}

// vlanParentResult wraps buildVlanScope output so it can be queried the way
// generated read code queries a get-config reply.
func vlanParentResult(n int) xmldot.Result {
	return xmldot.Get("<spanning-tree>"+buildVlanScope(n)+"</spanning-tree>", "spanning-tree")
}

// TestCollectListItemsXML_MatchesForEachScan compares the map built by a
// single walk against the per-item ForEach linear scan the generated read code
// used to perform, at a size where that scan is still correct.
func TestCollectListItemsXML_MatchesForEachScan(t *testing.T) {
	const n = 50
	parent := vlanParentResult(n)
	items := CollectListItemsXML(parent.Raw, "vlan", []string{"id"})

	if len(items) != n {
		t.Fatalf("collected %d items, want %d", len(items), n)
	}

	for i := 0; i < n; i++ {
		want := fmt.Sprintf("%d", i)

		// Legacy behaviour: a fresh linear scan for every item in state.
		var legacy xmldot.Result
		GetFromXPath(parent, "vlan").ForEach(func(_ int, v xmldot.Result) bool {
			if v.Get("id").String() == want {
				legacy = v
				return false
			}
			return true
		})

		got := items[CompositeKey(want)]
		if got.Get("id").String() != want {
			t.Errorf("id=%s: collected item id = %q, want %q", want, got.Get("id").String(), want)
		}
		if got.Get("priority").String() != legacy.Get("priority").String() {
			t.Errorf("id=%s: collected priority = %q, legacy scan priority = %q",
				want, got.Get("priority").String(), legacy.Get("priority").String())
		}
	}
}

// TestCollectListItemsXML_PastWildcardCap is the read-path counterpart to the
// namespace regression on the write path. GetFromXPath has to collect matching
// siblings into an array before ForEach can iterate them, and that collection
// caps at MaxWildcardResults (1000). For a list longer than 1000, every item
// past the 1000th is invisible to the scan — not an error, just silently
// absent — so state for those entries is dropped and the resource reports a
// perpetual diff. Walking the document directly has no array to cap.
func TestCollectListItemsXML_PastWildcardCap(t *testing.T) {
	for _, n := range []int{999, 1000, 1001, 1500, 4094} {
		parent := vlanParentResult(n)

		legacyCount := 0
		GetFromXPath(parent, "vlan").ForEach(func(_ int, v xmldot.Result) bool {
			legacyCount++
			return true
		})

		items := CollectListItemsXML(parent.Raw, "vlan", []string{"id"})
		if len(items) != n {
			t.Errorf("n=%d: collected %d items, want %d", n, len(items), n)
		}

		// Spot-check the last entry, which is the one the cap would drop.
		last := fmt.Sprintf("%d", n-1)
		if got := items[CompositeKey(last)].Get("id").String(); got != last {
			t.Errorf("n=%d: last item id = %q, want %q", n, got, last)
		}

		if n > MaxLegacyScanResults && legacyCount >= n {
			t.Errorf("n=%d: legacy scan saw %d items; expected it to be capped, "+
				"so this test no longer demonstrates the regression", n, legacyCount)
		}
	}
}

// MaxLegacyScanResults mirrors xmldot's MaxWildcardResults, restated here so
// the assertion above documents the threshold it depends on.
const MaxLegacyScanResults = 1000

// TestCollectListItemsXML_CompositeKeys verifies that lists keyed on more than
// one leaf are indexed on the combination, not just the first key.
func TestCollectListItemsXML_CompositeKeys(t *testing.T) {
	raw := `<host>` +
		`<ip-address>10.0.0.1</ip-address><community-or-user>public</community-or-user><version>2c</version>` +
		`</host><host>` +
		`<ip-address>10.0.0.1</ip-address><community-or-user>private</community-or-user><version>3</version>` +
		`</host>`

	items := CollectListItemsXML(raw, "host", []string{"ip-address", "community-or-user"})

	if len(items) != 2 {
		t.Fatalf("collected %d items, want 2", len(items))
	}
	if got := items[CompositeKey("10.0.0.1", "public")].Get("version").String(); got != "2c" {
		t.Errorf("public host version = %q, want %q", got, "2c")
	}
	if got := items[CompositeKey("10.0.0.1", "private")].Get("version").String(); got != "3" {
		t.Errorf("private host version = %q, want %q", got, "3")
	}
}

// TestCollectListItemsXML_NestedPath covers the multi-segment form, where the
// element is reached through one or more container segments first.
func TestCollectListItemsXML_NestedPath(t *testing.T) {
	raw := "<mst><configuration>" +
		"<instance><id>1</id><vlan-ids>1-100</vlan-ids></instance>" +
		"<instance><id>2</id><vlan-ids>101-200</vlan-ids></instance>" +
		"</configuration></mst>"

	items := CollectListItemsXML(raw, "mst/configuration/instance", []string{"id"})
	if len(items) != 2 {
		t.Fatalf("collected %d items, want 2", len(items))
	}
	if got := items[CompositeKey("2")].Get("vlan-ids").String(); got != "101-200" {
		t.Errorf("instance 2 vlan-ids = %q, want %q", got, "101-200")
	}
}

// TestCollectListItemsXML_AbsentAndEmpty checks the not-found paths return an
// empty map rather than panicking or looping forever.
func TestCollectListItemsXML_AbsentAndEmpty(t *testing.T) {
	cases := []struct {
		name string
		raw  string
		path string
	}{
		{"missing element", buildVlanScope(3), "mst"},
		{"missing container", buildVlanScope(3), "mst/configuration/instance"},
		{"empty document", "", "vlan"},
		{"empty path", buildVlanScope(3), ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			items := CollectListItemsXML(tc.raw, tc.path, []string{"id"})
			if len(items) != 0 {
				t.Errorf("collected %d items, want 0", len(items))
			}
		})
	}
}

// TestCompositeKey checks that distinct key tuples do not collide when their
// concatenations would otherwise be ambiguous.
func TestCompositeKey(t *testing.T) {
	if CompositeKey("ab", "c") == CompositeKey("a", "bc") {
		t.Error("CompositeKey collides on ambiguous splits")
	}
	if CompositeKey("10") != CompositeKey("10") {
		t.Error("CompositeKey is not deterministic")
	}
	if strings.ContainsAny(CompositeKey("a", "b"), "/.:-") {
		t.Error("CompositeKey separator overlaps characters valid in list key values")
	}
}

// --- Benchmarks -----------------------------------------------------------

// BenchmarkSetRawFromXPath_Sequential and BenchmarkSetRawFromXPathMulti_Batched
// demonstrate the O(n^2) vs O(n) difference on the write path. Run with:
//
//	go test ./internal/provider/helpers/... -bench SetRawFromXPath -benchtime=1x -run ^$
func BenchmarkSetRawFromXPath_Sequential(b *testing.B) {
	for _, n := range []int{100, 500, 1000} {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			values := make([]string, n)
			for j := 0; j < n; j++ {
				values[j] = fmt.Sprintf("<id>%d</id><priority>32768</priority>", j)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var body netconf.Body
				for _, v := range values {
					body = SetRawFromXPath(body, "/native/spanning-tree/vlan", v)
				}
			}
		})
	}
}

func BenchmarkSetRawFromXPathMulti_Batched(b *testing.B) {
	for _, n := range []int{100, 500, 1000, 4094} {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			values := make([]string, n)
			for j := 0; j < n; j++ {
				values[j] = fmt.Sprintf("<id>%d</id><priority>32768</priority>", j)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var body netconf.Body
				body = SetRawFromXPathMulti(body, "/native/spanning-tree/vlan", values)
			}
		})
	}
}

func BenchmarkCollectListItemsXML(b *testing.B) {
	for _, n := range []int{100, 500, 1000, 4094} {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			raw := buildVlanScope(n)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				CollectListItemsXML(raw, "vlan", []string{"id"})
			}
		})
	}
}
