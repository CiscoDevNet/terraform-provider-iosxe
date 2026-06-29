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

// TestConvertXPathToRestconfPath verifies the XPath-to-RESTCONF-style path
// conversion semantics implemented in xpath.go: the leading slash is trimmed,
// list-key predicates are flattened to "=value" (multi-key predicates joined
// with ","), single and double quotes around predicate values are stripped,
// and no URL encoding is applied — special characters survive verbatim.
func TestConvertXPathToRestconfPath(t *testing.T) {
	tests := []struct {
		name     string
		xPath    string
		expected string
	}{
		{
			name:     "simple element without predicates",
			xPath:    "native/interface",
			expected: "native/interface",
		},
		{
			name:     "leading slash is trimmed",
			xPath:    "/native/interface",
			expected: "native/interface",
		},
		{
			name:     "single predicate (unquoted value)",
			xPath:    "interface[name=GigabitEthernet1]",
			expected: "interface=GigabitEthernet1",
		},
		{
			name:     "single predicate (single-quoted value)",
			xPath:    "interface[name='GigabitEthernet1']",
			expected: "interface=GigabitEthernet1",
		},
		{
			name:     "single predicate (double-quoted value)",
			xPath:    `interface[name="GigabitEthernet1"]`,
			expected: "interface=GigabitEthernet1",
		},
		{
			name:     "leading slash and predicate",
			xPath:    "/native/interface[name='Gi1']",
			expected: "native/interface=Gi1",
		},
		{
			name:     "namespace prefix retained",
			xPath:    "Cisco-IOS-XE-native:native/Cisco-IOS-XE-policy:class-map[name='cm1']",
			expected: "Cisco-IOS-XE-native:native/Cisco-IOS-XE-policy:class-map=cm1",
		},
		{
			name:     "multi-key predicate joined with comma",
			xPath:    "vrf[name='VRF1'][af-name='ipv4']",
			expected: "vrf=VRF1,ipv4",
		},
		{
			name:     "nested path with multiple list keys",
			xPath:    "native/vrf[name='VRF1']/address-family[af-name='ipv4']",
			expected: "native/vrf=VRF1/address-family=ipv4",
		},
		{
			name:     "real-world nested path",
			xPath:    "Cisco-IOS-XE-native:native/policy/Cisco-IOS-XE-policy:class-map[name='cm']/match/activated-service-template",
			expected: "Cisco-IOS-XE-native:native/policy/Cisco-IOS-XE-policy:class-map=cm/match/activated-service-template",
		},
		{
			name:     "slash inside predicate value is preserved (no URL encoding)",
			xPath:    "interface[name='GigabitEthernet1/0/1']",
			expected: "interface=GigabitEthernet1/0/1",
		},
		{
			name:     "double-quoted slash value preserved (no URL encoding)",
			xPath:    `interface[name="GigabitEthernet1/0/1"]`,
			expected: "interface=GigabitEthernet1/0/1",
		},
		{
			name:     "special characters survive verbatim (no URL encoding)",
			xPath:    "policy[name='CPU>80%']",
			expected: "policy=CPU>80%",
		},
		{
			name:     "space in value survives verbatim",
			xPath:    "route-map[name='test map']",
			expected: "route-map=test map",
		},
		{
			name:     "at-sign and bang survive verbatim",
			xPath:    "description[text='hello@world!']",
			expected: "description=hello@world!",
		},
		{
			name:     "namespace + slash inside predicate, no encoding",
			xPath:    "Cisco-IOS-XE-native:native/route-map[name='my/route@map']",
			expected: "Cisco-IOS-XE-native:native/route-map=my/route@map",
		},
		{
			name:     "multi-key predicate with slashes preserved",
			xPath:    "native/vrf[name='V/RF']/address-family[af-name='ipv4']",
			expected: "native/vrf=V/RF/address-family=ipv4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertXPathToRestconfPath(tt.xPath)
			if result != tt.expected {
				t.Errorf("ConvertXPathToRestconfPath(%q) = %q, expected %q", tt.xPath, result, tt.expected)
			}
		})
	}
}
