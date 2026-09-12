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

import "testing"

func TestNormalizeASN(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: ""},
		{name: "asplain 2-byte", input: "65000", expected: "65000"},
		{name: "asplain 4-byte", input: "4201634865", expected: "4201634865"},
		{name: "asdot to asplain", input: "64111.56369", expected: "4201634865"},
		{name: "asdot second example", input: "65000.14902", expected: "4259854902"},
		{name: "asdot zero high", input: "0.65000", expected: "65000"},
		{name: "asdot one dot one", input: "1.1", expected: "65537"},
		{name: "asdot max", input: "65535.65535", expected: "4294967295"},
		{name: "invalid string", input: "not-a-number", expected: "not-a-number"},
		{name: "invalid asdot", input: "abc.def", expected: "abc.def"},
		{name: "too many dots", input: "1.2.3", expected: "1.2.3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizeASN(tt.input)
			if got != tt.expected {
				t.Errorf("NormalizeASN(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
