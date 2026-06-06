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
	"testing"
)

func TestNormalizeCommunityValue(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "decimal 1:1",
			input:    "65537",
			expected: "1:1",
		},
		{
			name:     "decimal 2:2",
			input:    "131074",
			expected: "2:2",
		},
		{
			name:     "decimal max community",
			input:    "4294967295",
			expected: "65535:65535",
		},
		{
			name:     "decimal 1:0 boundary",
			input:    "65536",
			expected: "1:0",
		},
		{
			name:     "decimal 0:1",
			input:    "1",
			expected: "0:1",
		},
		{
			name:     "already colon notation",
			input:    "1:1",
			expected: "1:1",
		},
		{
			name:     "4-byte ASN colon notation passthrough",
			input:    "65540:100",
			expected: "65540:100",
		},
		{
			name:     "well-known no-export",
			input:    "no-export",
			expected: "no-export",
		},
		{
			name:     "well-known no-advertise",
			input:    "no-advertise",
			expected: "no-advertise",
		},
		{
			name:     "well-known internet",
			input:    "internet",
			expected: "internet",
		},
		{
			name:     "well-known local-AS",
			input:    "local-AS",
			expected: "local-AS",
		},
		{
			name:     "well-known gshut",
			input:    "gshut",
			expected: "gshut",
		},
		{
			name:     "overflow uint32 passthrough",
			input:    "4294967296",
			expected: "4294967296",
		},
		{
			name:     "empty string passthrough",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeCommunityValue(tt.input)
			if result != tt.expected {
				t.Errorf("NormalizeCommunityValue(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
