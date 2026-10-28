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

// TestRemoveFromXPathMulti_MatchesSequential verifies that batching many
// RemoveFromXPath calls into a single RemoveFromXPathMulti call produces the
// same set of VLAN/priority elements, each still marked operation="remove",
// as calling RemoveFromXPath once per item (the pre-fix behavior).
func TestRemoveFromXPathMulti_MatchesSequential(t *testing.T) {
	const n = 50

	var sequential netconf.Body
	var xPaths []string
	for i := 0; i < n; i++ {
		xPath := fmt.Sprintf("/native/spanning-tree/vlan[id='%d']/priority", i)
		xPaths = append(xPaths, xPath)
		sequential = RemoveFromXPath(sequential, xPath)
	}

	var batched netconf.Body
	batched = RemoveFromXPathMulti(batched, xPaths)

	seqXML := sequential.Res()
	batchedXML := batched.Res()

	seqCount := xmldot.Get(seqXML, "native.spanning-tree.vlan.#").Int()
	batchedCount := xmldot.Get(batchedXML, "native.spanning-tree.vlan.#").Int()

	if seqCount != int64(n) {
		t.Fatalf("sequential: got %d vlan elements, want %d", seqCount, n)
	}
	if batchedCount != int64(n) {
		t.Fatalf("batched: got %d vlan elements, want %d", batchedCount, n)
	}

	seqIds := map[string]bool{}
	batchedIds := map[string]bool{}
	for i := int64(0); i < seqCount; i++ {
		id := xmldot.Get(seqXML, fmt.Sprintf("native.spanning-tree.vlan.%d.id", i)).String()
		op := xmldot.Get(seqXML, fmt.Sprintf("native.spanning-tree.vlan.%d.priority.@operation", i)).String()
		if op != "remove" {
			t.Errorf("sequential: vlan id=%s priority operation = %q, want %q", id, op, "remove")
		}
		seqIds[id] = true
	}
	for i := int64(0); i < batchedCount; i++ {
		id := xmldot.Get(batchedXML, fmt.Sprintf("native.spanning-tree.vlan.%d.id", i)).String()
		op := xmldot.Get(batchedXML, fmt.Sprintf("native.spanning-tree.vlan.%d.priority.@operation", i)).String()
		if op != "remove" {
			t.Errorf("batched: vlan id=%s priority operation = %q, want %q", id, op, "remove")
		}
		batchedIds[id] = true
	}

	for id := range seqIds {
		if !batchedIds[id] {
			t.Errorf("batched output missing vlan id=%s present in sequential output", id)
		}
	}
	for id := range batchedIds {
		if !seqIds[id] {
			t.Errorf("batched output has extra vlan id=%s not present in sequential output", id)
		}
	}

	t.Logf("✓ %d-item batched removal matches %d sequential RemoveFromXPath calls", n, n)
}

// TestRemoveFromXPathMulti_EmptyAndSingle checks the trivial fallback paths.
func TestRemoveFromXPathMulti_EmptyAndSingle(t *testing.T) {
	var body netconf.Body

	body = RemoveFromXPathMulti(body, nil)
	if body.Res() != "" {
		t.Errorf("empty input: got non-empty body %q", body.Res())
	}

	body = RemoveFromXPathMulti(body, []string{"/native/spanning-tree/vlan[id='10']/priority"})
	op := xmldot.Get(body.Res(), "native.spanning-tree.vlan.priority.@operation").String()
	if op != "remove" {
		t.Errorf("single-item input: priority operation = %q, want %q", op, "remove")
	}
}

// TestRemoveFromXPathMulti_MixedShapes checks that xPaths without a list-key
// predicate (nothing to batch) are still handled correctly alongside ones
// that do share a batchable parent.
func TestRemoveFromXPathMulti_MixedShapes(t *testing.T) {
	var body netconf.Body

	xPaths := []string{
		"/native/spanning-tree/vlan[id='10']/priority",
		"/native/spanning-tree/vlan[id='20']/priority",
		"/native/spanning-tree/logging",
	}
	body = RemoveFromXPathMulti(body, xPaths)
	xml := body.Res()

	vlanCount := xmldot.Get(xml, "native.spanning-tree.vlan.#").Int()
	if vlanCount != 2 {
		t.Errorf("got %d vlan elements, want 2", vlanCount)
	}
	loggingOp := xmldot.Get(xml, "native.spanning-tree.logging.@operation").String()
	if loggingOp != "remove" {
		t.Errorf("logging operation = %q, want %q", loggingOp, "remove")
	}
}

// TestRemoveFromXPathMulti_NamespacesPastWildcardCap is a regression test for a
// device-rejecting bug: xmldot's counted/indexed sibling queries capped at
// 1000 results before xmldot v0.6.0, so backfilling xmlns via
// augmentNamespaces left every list element past the 1000th without a
// namespace. IOS-XE then resolved those against the parent (native) namespace
// and failed the whole edit-config with "unknown-element" on <vlan>. Fragments
// now carry an inline xmlns.
//
// Note this only reproduces with namespace-prefixed xPaths, which is why the
// bare-path tests above did not catch it.
func TestRemoveFromXPathMulti_NamespacesPastWildcardCap(t *testing.T) {
	const stNs = `<vlan xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-spanning-tree">`

	// 4094 is the real-world worst case: a full VLAN range being un-declared.
	for _, n := range []int{999, 1000, 1001, 2000, 4094} {
		xPaths := make([]string, 0, n)
		for i := 0; i < n; i++ {
			xPaths = append(xPaths,
				fmt.Sprintf("/Cisco-IOS-XE-native:native/spanning-tree/Cisco-IOS-XE-spanning-tree:vlan[id='%d']/priority", i))
		}

		var body netconf.Body
		body = RemoveFromXPathMulti(body, xPaths)
		res := body.Res()

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

// TestRemoveFromXPathMulti_MatchesSequentialWithNamespaces pins byte-for-byte
// equivalence with the legacy per-item path on namespaced input, at sizes below
// the 1000 cap where that path is still correct. Kept small because the legacy
// path is quadratic.
func TestRemoveFromXPathMulti_MatchesSequentialWithNamespaces(t *testing.T) {
	for _, n := range []int{1, 2, 50} {
		xPaths := make([]string, 0, n)
		for i := 0; i < n; i++ {
			xPaths = append(xPaths,
				fmt.Sprintf("/Cisco-IOS-XE-native:native/spanning-tree/Cisco-IOS-XE-spanning-tree:vlan[id='%d']/priority", i))
		}

		var sequential netconf.Body
		for _, p := range xPaths {
			sequential = RemoveFromXPath(sequential, p)
		}
		var batched netconf.Body
		batched = RemoveFromXPathMulti(batched, xPaths)

		if sequential.Res() != batched.Res() {
			t.Errorf("n=%d: batched output differs from sequential\n sequential: %s\n batched:    %s",
				n, sequential.Res(), batched.Res())
		}
	}
}

// stVlanBody returns a body with spanning-tree VLAN entries for ids, as written by toBodyXML for
// VLANs that stay declared.
func stVlanBody(ids []int) netconf.Body {
	var b strings.Builder
	b.WriteString(`<native xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-native"><spanning-tree><mode xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-spanning-tree">pvst</mode>`)
	for _, id := range ids {
		fmt.Fprintf(&b, `<vlan xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-spanning-tree"><id>%d</id><max-age>10</max-age></vlan>`, id)
	}
	b.WriteString(`</spanning-tree></native>`)
	return netconf.NewBody(b.String())
}

// TestRemoveFromXPathMulti_MergesIntoExistingEntries checks that removals below a list entry the
// body already sets (a VLAN that stays declared while its priority is unset) are merged into that
// element instead of appending a second element with the same key, matching RemoveFromXPath.
func TestRemoveFromXPathMulti_MergesIntoExistingEntries(t *testing.T) {
	existing := []int{}
	for i := 0; i < 50; i++ {
		existing = append(existing, i)
	}
	// Half of the removals target existing entries, half are new.
	var xPaths []string
	for i := 25; i < 75; i++ {
		xPaths = append(xPaths, fmt.Sprintf("/Cisco-IOS-XE-native:native/spanning-tree/Cisco-IOS-XE-spanning-tree:vlan[id='%d']/priority", i))
	}

	sequential := stVlanBody(existing)
	for _, p := range xPaths {
		sequential = RemoveFromXPath(sequential, p)
	}
	batched := RemoveFromXPathMulti(stVlanBody(existing), xPaths)

	for name, body := range map[string]netconf.Body{"sequential": sequential, "batched": batched} {
		res := body.Res()
		ids := map[string]int{}
		GetFromXPath(xmldot.Get(res, "native.spanning-tree"), "vlan").ForEach(func(_ int, v xmldot.Result) bool {
			ids[v.Get("id").String()]++
			return true
		})
		if len(ids) != 75 {
			t.Errorf("%s: got %d distinct vlan ids, want 75", name, len(ids))
		}
		for id, count := range ids {
			if count != 1 {
				t.Errorf("%s: vlan id=%s appears %d times, want 1", name, id, count)
			}
		}
		if got := strings.Count(res, "<max-age>10</max-age>"); got != 50 {
			t.Errorf("%s: got %d max-age leaves, want 50 (existing content must be kept)", name, got)
		}
		if got := strings.Count(res, `<priority operation="remove">`) + strings.Count(res, `<priority operation="remove"/>`); got != 50 {
			t.Errorf("%s: got %d priority remove operations, want 50\n%s", name, got, res)
		}
	}
}

// TestRemoveFromXPathMulti_MergesWithinBatch checks that several removals below the same list entry
// share one element, e.g. removing multiple vlan-ids from one MST instance.
func TestRemoveFromXPathMulti_MergesWithinBatch(t *testing.T) {
	const prefix = "/Cisco-IOS-XE-native:native/spanning-tree/Cisco-IOS-XE-spanning-tree:mst/configuration/instance"
	xPaths := []string{
		prefix + "[id='1']/vlan-ids[.=10]",
		prefix + "[id='2']/vlan-ids[.=20]",
		prefix + "[id='1']/vlan-ids[.=11]",
		prefix + "[id='1']/vlan-ids[.=12]",
	}

	res := RemoveFromXPathMulti(netconf.Body{}, xPaths).Res()

	instances := map[string][]string{}
	GetFromXPath(xmldot.Get(res, "native.spanning-tree.mst.configuration"), "instance").ForEach(func(_ int, inst xmldot.Result) bool {
		id := inst.Get("id").String()
		GetFromXPath(inst, "vlan-ids").ForEach(func(_ int, v xmldot.Result) bool {
			instances[id] = append(instances[id], v.String())
			return true
		})
		return true
	})
	if got := strings.Count(res, "<instance"); got != 2 {
		t.Errorf("got %d instance elements, want 2\n%s", got, res)
	}
	if got := strings.Join(instances["1"], ","); got != "10,11,12" {
		t.Errorf("instance 1 vlan-ids = %s, want 10,11,12\n%s", got, res)
	}
	if got := strings.Join(instances["2"], ","); got != "20" {
		t.Errorf("instance 2 vlan-ids = %s, want 20\n%s", got, res)
	}
	if got := strings.Count(res, `<vlan-ids operation="remove">`); got != 4 {
		t.Errorf("got %d vlan-ids remove operations, want 4\n%s", got, res)
	}
}

// TestRemoveFromXPathMulti_WholeEntryAlreadyInBody checks that removing a whole list entry that
// the body also contains still marks the existing element, without adding a duplicate.
func TestRemoveFromXPathMulti_WholeEntryAlreadyInBody(t *testing.T) {
	const prefix = "/Cisco-IOS-XE-native:native/spanning-tree/Cisco-IOS-XE-spanning-tree:vlan"
	res := RemoveFromXPathMulti(stVlanBody([]int{10, 20}), []string{
		prefix + "[id='10']",
		prefix + "[id='30']",
	}).Res()

	if got := strings.Count(res, "<vlan"); got != 3 {
		t.Errorf("got %d vlan elements, want 3\n%s", got, res)
	}
	if got := strings.Count(res, `operation="remove"`); got != 2 {
		t.Errorf("got %d remove operations, want 2\n%s", got, res)
	}
}

// TestRemoveFromXPathMulti_MergesIntoLargeList checks the merge past 1000 entries and that it stays
// linear: every one of 4094 declared VLANs has its priority removed.
func TestRemoveFromXPathMulti_MergesIntoLargeList(t *testing.T) {
	const n = 4094
	ids := make([]int, n)
	xPaths := make([]string, n)
	for i := range ids {
		ids[i] = i + 1
		xPaths[i] = fmt.Sprintf("/Cisco-IOS-XE-native:native/spanning-tree/Cisco-IOS-XE-spanning-tree:vlan[id='%d']/priority", i+1)
	}

	res := RemoveFromXPathMulti(stVlanBody(ids), xPaths).Res()

	if got := strings.Count(res, "<vlan "); got != n {
		t.Errorf("got %d vlan elements, want %d", got, n)
	}
	if got := strings.Count(res, `<max-age>10</max-age><priority operation="remove"></priority></vlan>`); got != n {
		t.Errorf("got %d vlan elements with the priority removal merged in, want %d", got, n)
	}
}

// BenchmarkRemoveFromXPath_Sequential and BenchmarkRemoveFromXPathMulti_Batched
// demonstrate the O(n^2) vs O(n) difference directly. Run with:
//
//	go test ./internal/provider/helpers/... -bench RemoveFromXPath -benchtime=1x -run ^$
func BenchmarkRemoveFromXPath_Sequential(b *testing.B) {
	sizes := []int{100, 500, 1000, 2000}
	for _, n := range sizes {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var body netconf.Body
				for j := 0; j < n; j++ {
					body = RemoveFromXPath(body, fmt.Sprintf("/native/spanning-tree/vlan[id='%d']/priority", j))
				}
			}
		})
	}
}

func BenchmarkRemoveFromXPathMulti_Batched(b *testing.B) {
	sizes := []int{100, 500, 1000, 2000}
	for _, n := range sizes {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			xPaths := make([]string, n)
			for j := 0; j < n; j++ {
				xPaths[j] = fmt.Sprintf("/native/spanning-tree/vlan[id='%d']/priority", j)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var body netconf.Body
				body = RemoveFromXPathMulti(body, xPaths)
			}
		})
	}
}
