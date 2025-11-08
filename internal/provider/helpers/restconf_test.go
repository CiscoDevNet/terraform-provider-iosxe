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

// TestConvertXPathToRestconfPath tests the XPath to RESTCONF path conversion
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
			name:     "single predicate with concrete value",
			xPath:    "interface[name=GigabitEthernet1]",
			expected: "interface=GigabitEthernet1",
		},
		{
			name:     "single predicate with placeholder key",
			xPath:    "interface[%v=GigabitEthernet1]",
			expected: "interface=%v",
		},
		{
			name:     "single predicate with placeholder value",
			xPath:    "interface[name=%v]",
			expected: "interface=%v",
		},
		{
			name:     "double placeholders",
			xPath:    "interface[%v=%v]",
			expected: "interface=%v",
		},
		{
			name:     "multiple predicates (composite key)",
			xPath:    "vrf[name=VRF1][af-name=ipv4]",
			expected: "vrf=VRF1,ipv4",
		},
		{
			name:     "multiple predicates with placeholders",
			xPath:    "vrf[%v=VRF1][%v=ipv4]",
			expected: "vrf=%v,%v",
		},
		{
			name:     "namespace prefix retained",
			xPath:    "Cisco-IOS-XE-native:native/Cisco-IOS-XE-policy:class-map[%v=%v]",
			expected: "Cisco-IOS-XE-native:native/Cisco-IOS-XE-policy:class-map=%v",
		},
		{
			name:     "nested path with multiple keys",
			xPath:    "native/vrf[name=VRF1]/address-family[af-name=ipv4]",
			expected: "native/vrf=VRF1/address-family=ipv4",
		},
		{
			name:     "value with special characters (quoted)",
			xPath:    "interface[name='GigabitEthernet1/0/1']",
			expected: "interface=GigabitEthernet1%2F0%2F1",
		},
		{
			name:     "value with special characters (double quotes)",
			xPath:    `interface[name="GigabitEthernet1/0/1"]`,
			expected: "interface=GigabitEthernet1%2F0%2F1",
		},
		{
			name:     "leading slash",
			xPath:    "/native/interface[name=Gi1]",
			expected: "native/interface=Gi1",
		},
		{
			name:     "complex real-world example",
			xPath:    "Cisco-IOS-XE-native:native/policy/Cisco-IOS-XE-policy:class-map[%v=%v]/match/activated-service-template",
			expected: "Cisco-IOS-XE-native:native/policy/Cisco-IOS-XE-policy:class-map=%v/match/activated-service-template",
		},
		{
			name:     "placeholder with %s format",
			xPath:    "interface[name=%s]",
			expected: "interface=%s",
		},
		{
			name:     "placeholder with %d format",
			xPath:    "vlan[id=%d]",
			expected: "vlan=%d",
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

// TestConvertXPathToRestconfPath_URLEncoding tests URL encoding of special characters
func TestConvertXPathToRestconfPath_URLEncoding(t *testing.T) {
	tests := []struct {
		name     string
		xPath    string
		expected string
	}{
		{
			name:     "slash in value",
			xPath:    "interface[name='GigabitEthernet1/0/1']",
			expected: "interface=GigabitEthernet1%2F0%2F1",
		},
		{
			name:     "space in value",
			xPath:    "route-map[name='test map']",
			expected: "route-map=test+map",
		},
		{
			name:     "special characters",
			xPath:    "description[text='hello@world!']",
			expected: "description=hello%40world%21",
		},
		{
			name:     "equals sign in value",
			xPath:    "config[setting='key=value']",
			expected: "config=key%3Dvalue",
		},
		{
			name:     "percent in value (not placeholder)",
			xPath:    "policy[name='CPU>80%']",
			expected: "policy=CPU%3E80%25",
		},
		{
			name:     "placeholder not encoded",
			xPath:    "interface[name=%v]",
			expected: "interface=%v",
		},
		{
			name:     "multiple special chars with namespace",
			xPath:    "Cisco-IOS-XE-native:native/route-map[name='my/route@map']",
			expected: "Cisco-IOS-XE-native:native/route-map=my%2Froute%40map",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertXPathToRestconfPath(tt.xPath)
			if result != tt.expected {
				t.Errorf("ConvertXPathToRestconfPath(%q) = %q, expected %q", tt.xPath, result, tt.expected)
			} else {
				t.Logf("✓ %s correctly encoded", tt.name)
			}
		})
	}
}
