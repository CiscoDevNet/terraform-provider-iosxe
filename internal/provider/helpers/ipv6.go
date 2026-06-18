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
	"net"
	"strings"
)

// NormalizeIPv6Address normalizes an IPv6 address or prefix to its canonical form.
// Handles both bare addresses (e.g., "2001:DB8::1") and CIDR notation (e.g., "2001:DB8::/32").
// Returns the input unchanged if it cannot be parsed as a valid IPv6 address/prefix.
func NormalizeIPv6Address(val string) string {
	if val == "" {
		return val
	}

	// Check if the value contains a prefix length (CIDR notation)
	if strings.Contains(val, "/") {
		ip, network, err := net.ParseCIDR(val)
		if err != nil {
			return val
		}
		// Use the parsed IP (which preserves host bits) with the mask length
		ones, _ := network.Mask.Size()
		// If the original had host bits set (e.g., "2001:db8::1/64"), preserve them
		if !ip.Equal(network.IP) {
			return ip.String() + "/" + strings.Split(val, "/")[1]
		}
		// Pure network prefix — use net.IPNet canonical form
		_ = ones
		return network.String()
	}

	// Bare IPv6 address
	ip := net.ParseIP(val)
	if ip == nil {
		return val
	}
	// Ensure it's actually IPv6 (net.ParseIP also parses IPv4)
	if ip.To4() != nil {
		return val
	}
	return ip.String()
}
