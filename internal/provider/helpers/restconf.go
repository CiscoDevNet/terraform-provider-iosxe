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
	"net/url"
	"regexp"
	"strings"
)

// ConvertXPathToRestconfPath converts an XPath-style path to RESTCONF-style path.
// It retains placeholders (%v, %s, %d) and namespace prefixes.
// Values are URL-encoded as required by RESTCONF (except placeholders).
//
// XPath uses: element[key=value] or element[%v=value]
// RESTCONF uses: element=value (URL-encoded)
//
// Examples:
//   - "interface[name=GigabitEthernet1]" → "interface=GigabitEthernet1"
//   - "interface[%v=GigabitEthernet1]" → "interface=%v"
//   - "interface[%v=%v]" → "interface=%v"
//   - "vrf[name=VRF1][af-name=ipv4]" → "vrf=VRF1,ipv4"
//   - "Cisco-IOS-XE-native:native/vrf[%v=%v]" → "Cisco-IOS-XE-native:native/vrf=%v"
//   - "native/interface[name='GigabitEthernet1/0/1']" → "native/interface=GigabitEthernet1%2F0%2F1"
//   - "route-map[name='test map']" → "route-map=test+map"
func ConvertXPathToRestconfPath(xPath string) string {
	// Remove leading slash if present
	xPath = strings.TrimPrefix(xPath, "/")

	// Use splitXPathSegments to handle slashes inside predicates correctly
	segments := splitXPathSegmentsForConversion(xPath)
	result := make([]string, 0, len(segments))

	for _, segment := range segments {
		if !strings.Contains(segment, "[") {
			// No predicates, keep as-is
			result = append(result, segment)
			continue
		}

		// Extract element name and predicates
		elementName := segment
		predicates := []string{}

		// Find the element name (everything before first '[')
		bracketIdx := strings.IndexByte(segment, '[')
		if bracketIdx != -1 {
			elementName = segment[:bracketIdx]

			// Extract all predicates
			predicateStr := segment[bracketIdx:]

			// Parse predicates - handle quotes for values with special chars
			predicateRegex := regexp.MustCompile(`\[([^\]]+)\]`)
			matches := predicateRegex.FindAllStringSubmatch(predicateStr, -1)

			for _, match := range matches {
				if len(match) > 1 {
					predicate := match[1]

					// Parse key=value
					eqIdx := strings.Index(predicate, "=")
					if eqIdx != -1 {
						key := strings.TrimSpace(predicate[:eqIdx])
						value := strings.TrimSpace(predicate[eqIdx+1:])

						// Remove quotes from value if present
						value = strings.Trim(value, `'"`)

						// If key is a placeholder (%v), use placeholder for value
						// If value is a placeholder, keep it (don't encode placeholders)
						if strings.HasPrefix(key, "%") {
							predicates = append(predicates, "%v")
						} else if strings.HasPrefix(value, "%") {
							predicates = append(predicates, value)
						} else {
							// URL-encode the value for RESTCONF paths
							encodedValue := url.QueryEscape(value)
							predicates = append(predicates, encodedValue)
						}
					}
				}
			}
		}

		// Build RESTCONF-style segment
		if len(predicates) > 0 {
			result = append(result, elementName+"="+strings.Join(predicates, ","))
		} else {
			result = append(result, elementName)
		}
	}

	return strings.Join(result, "/")
}

// splitXPathSegmentsForConversion splits an XPath into segments while respecting
// bracket boundaries. This prevents splitting on slashes inside predicates.
func splitXPathSegmentsForConversion(xPath string) []string {
	segments := []string{}
	var currentSegment strings.Builder
	bracketDepth := 0
	inQuote := false
	var quoteChar rune

	for _, char := range xPath {
		switch char {
		case '\'', '"':
			if !inQuote {
				inQuote = true
				quoteChar = char
			} else if char == quoteChar {
				inQuote = false
			}
			currentSegment.WriteRune(char)
		case '[':
			if !inQuote {
				bracketDepth++
			}
			currentSegment.WriteRune(char)
		case ']':
			if !inQuote {
				bracketDepth--
			}
			currentSegment.WriteRune(char)
		case '/':
			if bracketDepth == 0 && !inQuote {
				// We're not inside brackets or quotes, so this is a segment separator
				if currentSegment.Len() > 0 {
					segments = append(segments, currentSegment.String())
					currentSegment.Reset()
				}
			} else {
				// We're inside brackets or quotes, so this is part of a predicate value
				currentSegment.WriteRune(char)
			}
		default:
			currentSegment.WriteRune(char)
		}
	}

	// Add the last segment if there's anything left
	if currentSegment.Len() > 0 {
		segments = append(segments, currentSegment.String())
	}

	return segments
}

// LastElement returns the last element of a YANG path with its namespace prefix.
// Example: "Cisco-IOS-XE-native:native/interface/GigabitEthernet=1" -> "Cisco-IOS-XE-native:GigabitEthernet"
func LastElement(path string) string {
	pes := strings.Split(path, "/")
	var prefix, element string
	for _, pe := range pes {
		// remove key
		if strings.Contains(pe, "=") {
			pe = pe[:strings.Index(pe, "=")]
		}
		if strings.Contains(pe, ":") {
			prefix = strings.Split(pe, ":")[0]
			element = strings.Split(pe, ":")[1]
		} else {
			element = pe
		}
	}
	return prefix + ":" + element
}
