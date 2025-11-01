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
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-netconf"
	"github.com/netascode/xmldot"
	"github.com/tidwall/gjson"
)

// Pre-compiled regular expressions for performance
var (
	xpathPredicateRegex = regexp.MustCompile(`\[[^\]]*\]`)
	namespaceRegex      = regexp.MustCompile(`[a-zA-Z0-9\-]+:`)
	predicatePattern    = regexp.MustCompile(`\[([^\]]+)\]`)
)

// Namespace base URL for Cisco YANG models
const namespaceBaseURL = "http://cisco.com/ns/yang/"

// Contains checks if a string exists in a slice of strings.
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// LastElement returns the last element of a YANG path with its namespace prefix.
// Example: "Cisco-IOS-XE-native:native/interface/GigabitEthernet[name='1']" -> "Cisco-IOS-XE-native:GigabitEthernet"
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

// GetValueSlice converts a slice of gjson.Result to a slice of Terraform attr.Value.
func GetValueSlice(result []gjson.Result) []attr.Value {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return v
}

// GetStringList converts a slice of gjson.Result to a Terraform types.List of strings.
func GetStringList(result []gjson.Result) types.List {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return types.ListValueMust(types.StringType, v)
}

// GetInt64List converts a slice of gjson.Result to a Terraform types.List of int64.
func GetInt64List(result []gjson.Result) types.List {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.Int64Value(result[r].Int())
	}
	return types.ListValueMust(types.Int64Type, v)
}

// GetStringSet converts a slice of gjson.Result to a Terraform types.Set of strings.
func GetStringSet(result []gjson.Result) types.Set {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return types.SetValueMust(types.StringType, v)
}

// GetInt64Set converts a slice of gjson.Result to a Terraform types.Set of int64.
func GetInt64Set(result []gjson.Result) types.Set {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.Int64Value(result[r].Int())
	}
	return types.SetValueMust(types.Int64Type, v)
}

// GetStringListXML converts a slice of xmldot.Result to a Terraform types.List of strings.
func GetStringListXML(result []xmldot.Result) types.List {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return types.ListValueMust(types.StringType, v)
}

// GetInt64ListXML converts a slice of xmldot.Result to a Terraform types.List of int64.
func GetInt64ListXML(result []xmldot.Result) types.List {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.Int64Value(result[r].Int())
	}
	return types.ListValueMust(types.Int64Type, v)
}

// GetStringSetXML converts a slice of xmldot.Result to a Terraform types.Set of strings.
func GetStringSetXML(result []xmldot.Result) types.Set {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return types.SetValueMust(types.StringType, v)
}

// GetInt64SetXML converts a slice of xmldot.Result to a Terraform types.Set of int64.
func GetInt64SetXML(result []xmldot.Result) types.Set {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.Int64Value(result[r].Int())
	}
	return types.SetValueMust(types.Int64Type, v)
}

// Must panics if err is not nil, otherwise returns the value v.
// Useful for must-succeed operations in initialization code.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// RemoveEmptyStrings filters out empty strings from a slice.
func RemoveEmptyStrings(s []string) []string {
	var r []string
	for _, v := range s {
		if v != "" {
			r = append(r, v)
		}
	}
	return r
}

// dotPath converts a YANG path to a dot path by removing XPath predicates (keys) and namespace prefixes.
// Example: "Cisco-IOS-XE-native:interface/GigabitEthernet[name]/description" -> "interface.GigabitEthernet.description"
func dotPath(path string) string {
	// Remove XPath predicates like [name='value'] or [name]
	path = xpathPredicateRegex.ReplaceAllString(path, "")

	path = strings.ReplaceAll(path, "/", ".")

	// Remove namespace prefixes like "Cisco-IOS-XE-native:" or "Cisco-IOS-XE-bgp:"
	return namespaceRegex.ReplaceAllString(path, "")
}

// removeNamespacePrefix removes the namespace prefix from a single element or attribute name.
// Example: "Cisco-IOS-XE-native:interface" -> "interface"
func removeNamespacePrefix(name string) string {
	if idx := strings.Index(name, ":"); idx != -1 {
		return name[idx+1:]
	}
	return name
}

// setWithNamespaces sets a value in the netconf body and automatically adds
// namespace declarations for any prefixes found in the path.
func setWithNamespaces(body netconf.Body, fullPath string, value any) netconf.Body {
	// Set the value
	body = body.Set(dotPath(fullPath), value)

	// Extract and add namespace declarations
	body = augmentNamespaces(body, fullPath)

	return body
}

// augmentNamespaces walks through the path and adds namespace declarations
// to elements where prefixes appear, checking the current body to avoid duplicates.
func augmentNamespaces(body netconf.Body, path string) netconf.Body {
	segments := strings.Split(path, ".")
	pathWithoutPrefix := make([]string, 0, len(segments))

	for _, segment := range segments {
		// Strip prefix from segment and add to path
		cleanSegment := removeNamespacePrefix(segment)
		// Also strip XPath predicates like [key='value']
		// This prevents malformed paths like "standard[name=SACL1].@xmlns"
		if idx := strings.IndexByte(cleanSegment, '['); idx != -1 {
			cleanSegment = cleanSegment[:idx]
		}
		pathWithoutPrefix = append(pathWithoutPrefix, cleanSegment)

		// If this segment has a namespace prefix, add xmlns declaration
		if idx := strings.Index(segment, ":"); idx != -1 {
			prefix := segment[:idx]
			currentPath := strings.Join(pathWithoutPrefix, ".")
			namespace := namespaceBaseURL + prefix

			xmlnsPath := currentPath + ".@xmlns"
			if !xmldot.Get(body.Res(), xmlnsPath).Exists() {
				body = body.Set(xmlnsPath, namespace)
			}
		}
	}

	return body
}

// splitXPathSegments splits an XPath into segments while respecting bracket boundaries.
// This prevents splitting on forward slashes inside predicates like [name='GigabitEthernet1/0/1']
func splitXPathSegments(xPath string) []string {
	segments := []string{}
	var currentSegment strings.Builder
	bracketDepth := 0

	for _, char := range xPath {
		switch char {
		case '[':
			bracketDepth++
			currentSegment.WriteRune(char)
		case ']':
			bracketDepth--
			currentSegment.WriteRune(char)
		case '/':
			if bracketDepth == 0 {
				// We're not inside brackets, so this is a segment separator
				if currentSegment.Len() > 0 {
					segments = append(segments, currentSegment.String())
					currentSegment.Reset()
				}
			} else {
				// We're inside brackets, so this is part of a predicate value
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

// buildXPathStructure is a helper that creates all elements in an XPath, including keys and namespaces.
// Returns the body with the structure created and the path segments for further processing.
func buildXPathStructure(body netconf.Body, xPath string) (netconf.Body, []string) {
	// Remove leading slash if present
	xPath = strings.TrimPrefix(xPath, "/")

	// Split into segments while respecting bracket boundaries
	segments := splitXPathSegments(xPath)

	// Build path incrementally, creating each element
	pathSegments := make([]string, 0, len(segments))

	for i, segment := range segments {
		// Parse segment: element[key='value'][key2='value2'] -> element, map[key:value, key2:value2]
		elementName, keys := parseXPathSegment(segment)

		// Add element name to path (without predicates)
		pathSegments = append(pathSegments, elementName)
		fullPath := strings.Join(pathSegments[:i+1], ".")

		// If this segment has keys, set all key values
		if len(keys) > 0 {
			for keyName, keyValue := range keys {
				keyPath := fullPath + "." + keyName
				body = setWithNamespaces(body, keyPath, keyValue)
			}
		}
	}

	// Ensure the complete path structure exists, including non-predicate elements.
	// Only create the structure if it doesn't already exist to avoid overwriting key values.
	if len(pathSegments) > 0 {
		fullPath := strings.Join(pathSegments, ".")
		existingContent := xmldot.Get(body.Res(), dotPath(fullPath)).String()
		if existingContent == "" {
			// Path doesn't exist yet, create it with SetWithNamespaces
			body = setWithNamespaces(body, fullPath, "")
		}
	}

	return body, pathSegments
}

// parseXPathSegment parses an XPath segment with single or multiple keys
// Supports formats:
//   - element[key='value']
//   - element[key1='value1'][key2='value2']
//   - element[key1='value1' and key2='value2']
//
// Returns: (elementName, map[keyName]keyValue)
func parseXPathSegment(segment string) (string, map[string]string) {
	// Check for predicate: element[...]
	if idx := strings.Index(segment, "["); idx != -1 {
		elementName := segment[:idx]
		keys := make(map[string]string)

		// Extract all predicates - handle both [key1='val1'][key2='val2'] and [key1='val1' and key2='val2']
		remainingPredicates := segment[idx:]

		// Use pre-compiled pattern to match predicates: [anything]
		predicates := predicatePattern.FindAllStringSubmatch(remainingPredicates, -1)

		for _, match := range predicates {
			if len(match) > 1 {
				predicate := match[1]

				// Split by 'and' to handle combined predicates
				conditions := strings.Split(predicate, " and ")
				for _, condition := range conditions {
					// Parse each condition: key='value' or key="value"
					if eqIdx := strings.Index(condition, "="); eqIdx != -1 {
						keyName := condition[:eqIdx]
						value := condition[eqIdx+1:]
						// Remove quotes
						keyValue := strings.Trim(value, `'"`)
						keys[keyName] = keyValue
					}
				}
			}
		}

		return elementName, keys
	}

	// No predicate
	return segment, nil
}

// RemoveFromXPath creates all elements in an XPath with an operation="remove" attribute
// on the last element for NETCONF delete operations.
// Supports the same XPath formats as SetFromXPath.
//
// Example: /native/interface[name='Gi1']/ip/address
// Creates: <native><interface><name>Gi1</name><ip><address operation="remove"/></ip></interface></native>
func RemoveFromXPath(body netconf.Body, xPath string) netconf.Body {
	body, pathSegments := buildXPathStructure(body, xPath)

	// Set operation="remove" on the last element
	if len(pathSegments) > 0 {
		targetPath := strings.Join(pathSegments, ".")
		operationPath := targetPath + ".@operation"
		body = setWithNamespaces(body, operationPath, "remove")
	}

	return body
}

// SetFromXPath creates all elements in an XPath, including keys and namespaces,
// and optionally sets a value at the final path location.
// Supports single and composite keys:
//   - Single: /interface[name='GigabitEthernet1']
//   - Multiple predicates: /interface[name='GigabitEthernet1'][vrf='VRF1']
//   - Combined with 'and': /interface[name='GigabitEthernet1' and vrf='VRF1']
//   - Values with slashes: /interface[name='GigabitEthernet1/0/1']
//
// If value is nil or empty string, only the structure is created without setting a value.
// If value is non-empty, it's set at the final path location.
//
// Multi-Root Support:
// SetFromXPath properly handles multiple root-level sibling elements (e.g., both <deny> and <permit>
// at the same level). The underlying xmldot library automatically detects when different root paths
// are used and creates sibling elements instead of nesting them.
//
// Example (multi-root XML):
//
//	body := netconf.Body{}
//	body = SetFromXPath(body, "sequence", "10")
//	body = SetFromXPath(body, "deny/std-ace/prefix", "10.0.0.0")
//	body = SetFromXPath(body, "permit/std-ace/prefix", "192.168.0.0")
//	// Result: <sequence>10</sequence><deny>...</deny><permit>...</permit>
func SetFromXPath(body netconf.Body, xPath string, value any) netconf.Body {
	body, pathSegments := buildXPathStructure(body, xPath)

	// Only set the value if it's not nil and not empty string
	// This prevents overwriting key children when no value is needed
	if len(pathSegments) > 0 && value != nil && value != "" {
		fullPath := strings.Join(pathSegments, ".")
		body = setWithNamespaces(body, fullPath, value)
	}

	return body
}

// SetRawFromXPath creates all elements in an XPath, including keys and namespace declarations,
// then inserts raw XML content at the final path location. This is useful when you have
// pre-formatted XML that needs to be inserted as child elements.
//
// The value parameter should contain raw XML content (child elements, attributes, etc.) that will
// be parsed and inserted at the target path. The content is wrapped in the final element tag
// specified by the xPath.
//
// Multi-root Support:
// When called multiple times with the same path, this function appends the new XML content
// as an additional sibling element, creating a multi-root XML fragment at the parent level.
// The underlying xmldot library automatically handles multi-root fragments, making this safe for
// creating multiple sibling elements (e.g., multiple <interface> elements in a list).
//
// Example (single call):
//
//	xPath: /Cisco-IOS-XE-native:native/interface[name='Gi1']
//	value: "<description>Management</description><shutdown/>"
//	Result: <native xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-native">
//	          <interface>
//	            <name>Gi1</name>
//	            <description>Management</description>
//	            <shutdown/>
//	          </interface>
//	        </native>
//
// Example (multiple calls - multi-root):
//
//	First call:  SetRawFromXPath(body, "/native/interface", "<name>Gi1</name>")
//	Second call: SetRawFromXPath(body, "/native/interface", "<name>Gi2</name>")
//	Result: <native>
//	          <interface><name>Gi1</name></interface>
//	          <interface><name>Gi2</name></interface>
//	        </native>
//
// Unlike SetFromXPath, this function:
//   - Adds xmlns declarations for namespace prefixes in the path (via buildXPathStructure)
//   - Inserts the value as raw XML (parsed as child elements) rather than as text content
//   - Uses body.SetRaw() instead of body.Set() for XML insertion
//   - Supports appending multiple elements at the same path (multi-root fragments)
func SetRawFromXPath(body netconf.Body, xPath string, value string) netconf.Body {
	if len(value) == 0 {
		return body
	}

	// Remove leading slash if present
	xPath = strings.TrimPrefix(xPath, "/")

	// Split into segments while respecting bracket boundaries
	segments := splitXPathSegments(xPath)
	if len(segments) == 0 {
		return body
	}

	// Extract the final element name to wrap the content
	finalSegment := segments[len(segments)-1]
	finalElement, keys := parseXPathSegment(finalSegment)
	finalElementClean := removeNamespacePrefix(finalElement)
	wrappedContent := "<" + finalElementClean + ">" + value + "</" + finalElementClean + ">"

	// Build parent structure (everything except the final element)
	if len(segments) > 1 {
		parentXPath := "/" + strings.Join(segments[:len(segments)-1], "/")
		body, _ = buildXPathStructure(body, parentXPath)

		// Get parent path for setting content
		parentPathSegments := make([]string, 0, len(segments)-1)
		for _, segment := range segments[:len(segments)-1] {
			elementName, _ := parseXPathSegment(segment)
			parentPathSegments = append(parentPathSegments, elementName)
		}
		parentPath := dotPath(strings.Join(parentPathSegments, "."))

		// Check if content already exists at parent path
		existingXML := xmldot.Get(body.Res(), parentPath).Raw

		if existingXML != "" {
			// Append wrapped element as sibling to existing content
			// xmldot now supports multi-root XML fragments at the parent level
			combinedXML := existingXML + wrappedContent
			body = body.SetRaw(parentPath, combinedXML)
		} else {
			// First element, set the wrapped content at parent
			body = body.SetRaw(parentPath, wrappedContent)
		}
	} else {
		// No parent path - the element is at root level
		// Build any keys if present
		if len(keys) > 0 {
			tempBody := netconf.Body{}
			for keyName, keyValue := range keys {
				tempBody = setWithNamespaces(tempBody, keyName, keyValue)
			}
			wrappedContent = "<" + finalElementClean + ">" + tempBody.Res() + value + "</" + finalElementClean + ">"
		}

		// Check if root already has content
		existingXML := body.Res()
		if existingXML != "" {
			// Append as sibling
			body = body.SetRaw("", existingXML+wrappedContent)
		} else {
			// First element
			body = body.SetRaw("", wrappedContent)
		}
	}

	// Add namespace declarations if present in the original path
	// Convert XPath segments back to dotPath format for augmentNamespaces
	if len(segments) > 0 {
		dotPathForNamespaces := strings.Join(segments, ".")
		body = augmentNamespaces(body, dotPathForNamespaces)
	}

	return body
}

// GetFromXPath converts an XPath expression to a xmldot path with filters and retrieves the result.
// Uses xmldot's native filter syntax #() for single predicates, manual filtering for multiple.
// Supports the same XPath formats as SetFromXPath:
//   - Single: /interface[name='GigabitEthernet1']
//   - Multiple predicates: /interface[name='GigabitEthernet1'][vrf='VRF1']
//   - Combined with 'and': /interface[name='GigabitEthernet1' and vrf='VRF1']
//   - Values with slashes: /interface[name='GigabitEthernet1/0/1']
//   - Nested paths: /native/interface[name='Gi1']/ip/address
//
// Example: /native/interface[name='Gi1']/ip/address
// Converts to: native.interface.#(name==Gi1).ip.address
func GetFromXPath(res xmldot.Result, xPath string) xmldot.Result {
	// Remove leading slash if present
	xPath = strings.TrimPrefix(xPath, "/")

	// Split into segments while respecting bracket boundaries
	segments := splitXPathSegments(xPath)

	// Check if we need manual filtering:
	// 1. Any segment has multiple predicates
	// 2. Multiple segments have predicates (nested filters not supported by xmldot)
	needsManualFiltering := false
	segmentsWithKeys := 0
	for _, segment := range segments {
		_, keys := parseXPathSegment(segment)
		if len(keys) > 1 {
			needsManualFiltering = true
			break
		}
		if len(keys) > 0 {
			segmentsWithKeys++
		}
	}
	if segmentsWithKeys > 1 {
		needsManualFiltering = true
	}

	// If only single predicates, use xmldot's native filter syntax (fast path)
	if !needsManualFiltering {
		pathParts := make([]string, 0, len(segments))
		for _, segment := range segments {
			elementName, keys := parseXPathSegment(segment)

			// Remove namespace prefix from element name only
			elementName = removeNamespacePrefix(elementName)

			if len(keys) == 1 {
				// Single predicate - use xmldot's native filter syntax
				// Also remove namespace prefix from key name
				for keyName, keyValue := range keys {
					keyName = removeNamespacePrefix(keyName)
					pathParts = append(pathParts, elementName+".#("+keyName+"=="+keyValue+")")
				}
			} else {
				// No predicates
				pathParts = append(pathParts, elementName)
			}
		}

		dotPath := strings.Join(pathParts, ".")
		return res.Get(dotPath)
	}

	// Slow path: manual filtering for multiple predicates
	// Since we need to work with absolute paths for counting, get the XML from res.Raw
	xml := res.Raw
	pathSoFar := make([]string, 0, len(segments))

	for _, segment := range segments {
		elementName, keys := parseXPathSegment(segment)

		// Remove namespace prefix from element name
		elementName = removeNamespacePrefix(elementName)
		pathSoFar = append(pathSoFar, elementName)

		// Build current path
		currentPath := strings.Join(pathSoFar, ".")

		// Check if there are multiple sibling elements using absolute path
		countPath := currentPath + ".#"
		count := xmldot.Get(xml, countPath).Int()

		// Apply filtering if keys exist
		if len(keys) > 0 {
			found := false
			if count > 1 {
				// Multiple elements - iterate using numeric indices
				for idx := 0; idx < int(count); idx++ {
					indexedPath := fmt.Sprintf("%s.%d", currentPath, idx)
					item := xmldot.Get(xml, indexedPath)

					allMatch := true
					for keyName, keyValue := range keys {
						// Remove namespace prefix from key name
						keyName = removeNamespacePrefix(keyName)
						keyResult := item.Get(keyName)
						if !keyResult.Exists() || keyResult.String() != keyValue {
							allMatch = false
							break
						}
					}

					if allMatch {
						// Update currentPath to point to the matched item
						pathSoFar[len(pathSoFar)-1] = fmt.Sprintf("%s.%d", elementName, idx)
						found = true
						break
					}
				}
			} else {
				// Single element - check directly
				currentResult := xmldot.Get(xml, currentPath)
				allMatch := true
				for keyName, keyValue := range keys {
					// Remove namespace prefix from key name
					keyName = removeNamespacePrefix(keyName)
					keyResult := currentResult.Get(keyName)
					if !keyResult.Exists() || keyResult.String() != keyValue {
						allMatch = false
						break
					}
				}
				found = allMatch
			}

			if !found {
				return xmldot.Result{}
			}
		}
	}

	// Return the final result
	finalPath := strings.Join(pathSoFar, ".")
	return xmldot.Get(xml, finalPath)
}

// GetXpathFilter creates a NETCONF XPath filter with namespace prefixes removed.
// It processes the XPath expression to strip namespace prefixes from both element names
// and predicate key names, preserving the path structure.
//
// Supports the same XPath formats as SetFromXPath and GetFromXPath:
//   - Paths with namespace prefixes: /Cisco-IOS-XE-native:native/interface
//   - Single predicates: /native/interface[name='GigabitEthernet1']
//   - Multiple predicates: /native/interface[name='Gi1'][vrf='VRF1']
//   - Nested paths: /native/interface[name='Gi1']/ip/address
//   - Values with slashes: /native/interface[name='GigabitEthernet1/0/1']
//   - Predicates with namespace prefixes: /Cisco-IOS-XE-native:interface[Cisco-IOS-XE-native:name='Gi1']
//
// Example transformations:
//
//	Input:  "/Cisco-IOS-XE-native:native/aaa"
//	Output: netconf.Filter{Type: "xpath", Content: "/native/aaa"}
//
//	Input:  "/Cisco-IOS-XE-native:native/interface[Cisco-IOS-XE-native:name='Gi1']/Cisco-IOS-XE-native:ip"
//	Output: netconf.Filter{Type: "xpath", Content: "/native/interface[name='Gi1']/ip"}
//
//	Input:  "/native/interface[name='GigabitEthernet1/0/1']"
//	Output: netconf.Filter{Type: "xpath", Content: "/native/interface[name='GigabitEthernet1/0/1']"}
func GetXpathFilter(xPath string) netconf.Filter {
	// Remove leading slash if present
	xPath = strings.TrimPrefix(xPath, "/")

	// Split into segments while respecting bracket boundaries
	segments := splitXPathSegments(xPath)

	// Process each segment to remove namespace prefixes
	processedSegments := make([]string, 0, len(segments))
	for _, segment := range segments {
		elementName, keys := parseXPathSegment(segment)

		// Remove namespace prefix from element name
		elementName = removeNamespacePrefix(elementName)

		// Reconstruct segment with predicates
		if len(keys) > 0 {
			// Build predicates
			predicates := make([]string, 0, len(keys))
			for keyName, keyValue := range keys {
				// Remove namespace prefix from key name
				keyName = removeNamespacePrefix(keyName)
				predicates = append(predicates, fmt.Sprintf("%s='%s'", keyName, keyValue))
			}
			// Reconstruct segment with all predicates
			reconstructed := elementName
			for _, pred := range predicates {
				reconstructed += "[" + pred + "]"
			}
			processedSegments = append(processedSegments, reconstructed)
		} else {
			processedSegments = append(processedSegments, elementName)
		}
	}

	// Join segments back with slashes
	cleanedPath := "/" + strings.Join(processedSegments, "/")
	return netconf.XPathFilter(cleanedPath)
}

// EditConfig edits the configuration on the device
// If the server supports the candidate capability, it will edit the configuration in the candidate datastore
// and commit it to the running datastore if commit is true.
// If the server does not support the candidate capability, it will edit the configuration in the running datastore.
//
// IMPORTANT: IOS-XE devices with writable-running capability may not respond to Lock/Unlock operations.
// This function skips locking when only writable-running is available, as locking is optional for that capability.
//
// Parameters:
//   - ctx: context.Context
//   - client: *netconf.Client
//   - body: string
//   - commit: bool
func EditConfig(ctx context.Context, client *netconf.Client, body string, commit bool) error {
	candidate := client.ServerHasCapability("urn:ietf:params:netconf:capability:candidate:1.0")
	writableRunning := client.ServerHasCapability("urn:ietf:params:netconf:capability:writable-running:1.0")

	if candidate {
		// Lock running datastore
		if _, err := client.Lock(ctx, "running"); err != nil {
			return fmt.Errorf("failed to lock running datastore: %w", err)
		}
		defer client.Unlock(ctx, "running")

		// Lock candidate datastore
		if _, err := client.Lock(ctx, "candidate"); err != nil {
			return fmt.Errorf("failed to lock candidate datastore: %w", err)
		}
		defer client.Unlock(ctx, "candidate")

		if _, err := client.EditConfig(ctx, "candidate", body); err != nil {
			return fmt.Errorf("failed to edit config: %w", err)
		}

		if commit {
			if _, err := client.Commit(ctx); err != nil {
				return fmt.Errorf("failed to commit config: %w", err)
			}
		}
	} else if writableRunning {
		// IOS-XE devices support writable-running but Lock operations may hang
		// Edit running config directly without locking (optional for writable-running)
		if _, err := client.EditConfig(ctx, "running", body); err != nil {
			return fmt.Errorf("failed to edit config: %w", err)
		}
	} else {
		// Fallback: try with lock (standard NETCONF behavior)
		// Lock running datastore
		if _, err := client.Lock(ctx, "running"); err != nil {
			return fmt.Errorf("failed to lock running datastore: %w", err)
		}
		defer client.Unlock(ctx, "running")

		if _, err := client.EditConfig(ctx, "running", body); err != nil {
			return fmt.Errorf("failed to edit config: %w", err)
		}
	}
	return nil
}

// ConvertRestconfPathToXPath converts a RESTCONF-style path to XPath-style with predicates.
// Key names are always converted to %v placeholders for generic path handling.
//
// RESTCONF uses: element=value
// XPath uses: element[%v=value]
//
// Examples:
//   - "interface/%s=%v" → "interface/%s[%v=%v]"
//   - "interface/GigabitEthernet=1" → "interface/GigabitEthernet[%v=1]"
//   - "vrf=VRF1,af=ipv4" → "vrf[%v=VRF1][%v=ipv4]"
//   - "native/vrf=VRF1/address-family=ipv4" → "native/vrf[%v=VRF1]/address-family[%v=ipv4]"
func ConvertRestconfPathToXPath(path string) string {
	// Replace =placeholder patterns with [%v=placeholder]
	// This handles: %s=%v, %d=%v, %v=%v, etc.
	re := regexp.MustCompile(`=(%[sdv])`)
	path = re.ReplaceAllString(path, "[%v=$1]")

	// Handle concrete paths with actual key values
	// Split path into segments
	segments := strings.Split(path, "/")
	result := make([]string, 0, len(segments))

	for _, segment := range segments {
		if !strings.Contains(segment, "=") || strings.Contains(segment, "[") {
			// No key in this segment, or already processed, keep as-is
			result = append(result, segment)
			continue
		}

		// Extract element name and key values
		parts := strings.SplitN(segment, "=", 2)
		if len(parts) != 2 {
			result = append(result, segment)
			continue
		}

		elementName := parts[0]
		keyValues := strings.Split(parts[1], ",")

		// Build XPath predicates with %v placeholders for key names
		predicates := ""
		for _, keyValue := range keyValues {
			predicates += fmt.Sprintf("[%%v=%s]", keyValue)
		}

		result = append(result, elementName+predicates)
	}

	return strings.Join(result, "/")
}
