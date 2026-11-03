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
// device-rejecting bug: xmldot's counted/indexed sibling queries cap at
// MaxWildcardResults (1000), so backfilling xmlns via augmentNamespaces left
// every list element past the 1000th without a namespace. IOS-XE then resolved
// those against the parent (native) namespace and failed the whole edit-config
// with "unknown-element" on <vlan>. Fragments now carry an inline xmlns.
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
