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
	"strconv"
	"strings"
)

// NormalizeASN converts a BGP AS number to its asplain (uint32) string form.
// Handles both asplain ("4201634865") and asdot ("64111.56369") notations.
// Returns the input unchanged if it cannot be parsed.
func NormalizeASN(val string) string {
	if val == "" {
		return val
	}

	// If it contains a dot, it's asdot notation: "high.low"
	if strings.Contains(val, ".") {
		parts := strings.SplitN(val, ".", 2)
		if len(parts) != 2 {
			return val
		}
		high, err := strconv.ParseUint(parts[0], 10, 16)
		if err != nil {
			return val
		}
		low, err := strconv.ParseUint(parts[1], 10, 16)
		if err != nil {
			return val
		}
		return fmt.Sprintf("%d", high*65536+low)
	}

	// Already asplain — validate it parses as a number and return as-is
	if _, err := strconv.ParseUint(val, 10, 32); err != nil {
		return val
	}
	return val
}

// AlternateASN returns the same BGP AS number in the other notation: asdot
// ("1.1") becomes asplain ("65537") and a 4-byte asplain value becomes asdot.
// IOS-XE keys the BGP process by the AS in whichever notation `bgp asnotation
// dot` currently selects, so toggling it renames the list entry. 2-byte AS
// numbers look the same in both notations, and unparseable input is returned
// unchanged.
func AlternateASN(val string) string {
	if strings.Contains(val, ".") {
		return NormalizeASN(val)
	}
	n, err := strconv.ParseUint(val, 10, 32)
	if err != nil || n < 65536 {
		return val
	}
	return fmt.Sprintf("%d.%d", n/65536, n%65536)
}
