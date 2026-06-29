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

func TestNormalizeIPv6Address(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "uppercase to lowercase",
			input:    "2001:DB8::1",
			expected: "2001:db8::1",
		},
		{
			name:     "mixed case to lowercase",
			input:    "2001:Db8:CAFE::1",
			expected: "2001:db8:cafe::1",
		},
		{
			name:     "already canonical",
			input:    "2001:db8::1",
			expected: "2001:db8::1",
		},
		{
			name:     "full expansion to compressed",
			input:    "2001:0db8:0000:0000:0000:0000:0000:0001",
			expected: "2001:db8::1",
		},
		{
			name:     "CIDR uppercase",
			input:    "2001:DB8::/32",
			expected: "2001:db8::/32",
		},
		{
			name:     "CIDR mixed case",
			input:    "2001:Db8:CAFE::/48",
			expected: "2001:db8:cafe::/48",
		},
		{
			name:     "CIDR already canonical",
			input:    "2001:db8::/32",
			expected: "2001:db8::/32",
		},
		{
			name:     "CIDR full expansion",
			input:    "2001:0DB8:0000:0000::/64",
			expected: "2001:db8::/64",
		},
		{
			name:     "CIDR with host bits",
			input:    "2001:DB8::1/128",
			expected: "2001:db8::1/128",
		},
		{
			name:     "loopback",
			input:    "::1",
			expected: "::1",
		},
		{
			name:     "loopback expanded",
			input:    "0000:0000:0000:0000:0000:0000:0000:0001",
			expected: "::1",
		},
		{
			name:     "link-local uppercase",
			input:    "FE80::1",
			expected: "fe80::1",
		},
		{
			name:     "empty string passthrough",
			input:    "",
			expected: "",
		},
		{
			name:     "invalid string passthrough",
			input:    "not-an-ipv6",
			expected: "not-an-ipv6",
		},
		{
			name:     "IPv4 passthrough",
			input:    "192.168.1.1",
			expected: "192.168.1.1",
		},
		{
			name:     "IPv4 CIDR passthrough",
			input:    "10.0.0.0/8",
			expected: "10.0.0.0/8",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeIPv6Address(tt.input)
			if result != tt.expected {
				t.Errorf("NormalizeIPv6Address(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
