// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
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
)

// ConvertXPathToRestconfPath converts an XPath like
// "/Cisco-IOS-XE-native:native/interface/GigabitEthernet[name='%s']"
// to a RESTCONF-style path like
// "Cisco-IOS-XE-native:native/interface/GigabitEthernet=%s".
// This is used by the code generator to produce resource ID paths.
func ConvertXPathToRestconfPath(xpath string) string {
	xpath = strings.TrimPrefix(xpath, "/")

	segments := splitXPathSegmentsForConversion(xpath)
	var result []string

	for _, seg := range segments {
		if idx := strings.Index(seg, "["); idx != -1 {
			element := seg[:idx]
			predicates := seg[idx:]
			var keys []string
			for predicates != "" {
				start := strings.Index(predicates, "[")
				end := strings.Index(predicates, "]")
				if start == -1 || end == -1 {
					break
				}
				pred := predicates[start+1 : end]
				if eqIdx := strings.Index(pred, "="); eqIdx != -1 {
					val := pred[eqIdx+1:]
					val = strings.Trim(val, "'\"")
					keys = append(keys, val)
				}
				predicates = predicates[end+1:]
			}
			if len(keys) > 0 {
				result = append(result, fmt.Sprintf("%s=%s", element, strings.Join(keys, ",")))
			} else {
				result = append(result, element)
			}
		} else {
			result = append(result, seg)
		}
	}

	return strings.Join(result, "/")
}

// splitXPathSegmentsForConversion splits an XPath into segments by '/',
// but only at slashes that are outside of brackets.
func splitXPathSegmentsForConversion(xpath string) []string {
	var segments []string
	depth := 0
	start := 0
	for i, ch := range xpath {
		switch ch {
		case '[':
			depth++
		case ']':
			depth--
		case '/':
			if depth == 0 {
				if i > start {
					segments = append(segments, xpath[start:i])
				}
				start = i + 1
			}
		}
	}
	if start < len(xpath) {
		segments = append(segments, xpath[start:])
	}
	return segments
}
