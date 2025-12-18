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
	"context"
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-netconf"
	"github.com/netascode/xmldot"
)

// Pre-compiled regular expressions for performance
var (
	xpathPredicateRegex = regexp.MustCompile(`\[[^\]]*\]`)
	namespaceRegex      = regexp.MustCompile(`[a-zA-Z0-9\-]+:`)
	predicatePattern    = regexp.MustCompile(`\[([^\]]+)\]`)
)

// Namespace base URL for Cisco YANG models
const namespaceBaseURL = "http://cisco.com/ns/yang/"

// namespaceExceptions maps namespace prefixes to their full namespace URLs
// for modules that don't follow the standard pattern
var namespaceExceptions = map[string]string{
	"Cisco-IOS-XE-template": "http://cisco.com/ns/yang/ios-xe/template",
}

// AcquireNetconfLock acquires the appropriate lock for a NETCONF operation.
//
// Lock strategy based on reuseConnection and operation type:
// - When reuseConnection=false: Lock for ALL operations (serializes everything including Close)
// - When reuseConnection=true && isWrite=true: Lock for WRITE operations (Lock/EditConfig/Commit sequence)
// - When reuseConnection=true && isWrite=false: NO lock for READ operations (concurrent reads allowed)
//
// This prevents issues:
// - When reuse disabled: Prevents concurrent Close() attempts on same connection
// - When reuse enabled: Serializes write sequences but allows concurrent reads
//
// Returns true if lock was acquired, false if not acquired.
//
// Usage:
//
//	locked := helpers.AcquireNetconfLock(opMutex, reuseConnection, isWrite)
//	if locked {
//	    defer opMutex.Unlock()
//	}
//	defer helpers.CloseNetconfConnection(ctx, client, reuseConnection)
func AcquireNetconfLock(opMutex *sync.Mutex, reuseConnection bool, isWrite bool) bool {
	// Serialize all operations when reuse disabled (prevent concurrent Close)
	if !reuseConnection {
		opMutex.Lock()
		return true
	}

	// When reuse enabled, only serialize write operations
	// Read operations can run concurrently
	if isWrite {
		opMutex.Lock()
		return true
	}

	return false
}

// CloseNetconfConnection safely closes a NETCONF connection if reuse is disabled.
//
// IMPORTANT: This must be called with the operation mutex still held (deferred after AcquireNetconfLock)
// to prevent concurrent close attempts.
//
// Usage:
//
//	defer helpers.CloseNetconfConnection(ctx, device.NetconfClient, device.ReuseConnection)
func CloseNetconfConnection(ctx context.Context, client *netconf.Client, reuseConnection bool) {
	if reuseConnection {
		return // Keep connection open for reuse
	}

	// Close the connection (mutex already held by caller)
	if err := client.Close(); err != nil {
		// Log error but don't fail - connection cleanup is best-effort
		tflog.Warn(ctx, fmt.Sprintf("Failed to close NETCONF connection: %s", err))
	}
}

// FormatNetconfError extracts detailed error information from a NETCONF error
func FormatNetconfError(err error) string {
	if netconfErr, ok := err.(*netconf.NetconfError); ok {
		var details strings.Builder
		details.WriteString(netconfErr.Message)

		// Add detailed error information from each ErrorModel
		for i, e := range netconfErr.Errors {
			if i == 0 {
				details.WriteString("\n\nError Details:")
			}
			details.WriteString(fmt.Sprintf("\n  [%d] ", i+1))

			if e.ErrorMessage != "" {
				details.WriteString(e.ErrorMessage)
			}

			if e.ErrorPath != "" {
				details.WriteString(fmt.Sprintf(" (path: %s)", e.ErrorPath))
			}

			if e.ErrorType != "" || e.ErrorTag != "" {
				details.WriteString(fmt.Sprintf(" [type=%s, tag=%s]", e.ErrorType, e.ErrorTag))
			}

			if e.ErrorInfo != "" {
				details.WriteString(fmt.Sprintf("\n      Info: %s", e.ErrorInfo))
			}
		}

		return details.String()
	}
	return err.Error()
}

// EditConfig edits the configuration on the device
// If the server supports the candidate capability, it will edit the configuration in the candidate datastore
// and commit it to the running datastore if commit is true.
// If the server does not support the candidate capability, it will edit the configuration in the running datastore.
//
// IMPORTANT: When connection reuse is enabled, callers MUST serialize calls to EditConfig using an
// application-level mutex that also covers ManageNetconfConnection(). This prevents concurrent goroutines
// from attempting to acquire NETCONF datastore locks simultaneously on the same session.
//
// Parameters:
//   - ctx: context.Context
//   - client: *netconf.Client
//   - body: string
//   - commit: bool
func EditConfig(ctx context.Context, client *netconf.Client, body string, commit bool) error {
	// Ensure connection is open before checking capabilities
	// With lazy connections, Open() is idempotent and safe to call multiple times
	if err := client.Open(); err != nil {
		return fmt.Errorf("failed to open NETCONF connection: %w", err)
	}

	candidate := client.ServerHasCapability("urn:ietf:params:netconf:capability:candidate:1.0")

	if candidate {
		if commit {
			// Lock running datastore
			if _, err := client.Lock(ctx, "running"); err != nil {
				return fmt.Errorf("failed to lock running datastore: %s", FormatNetconfError(err))
			}
			defer client.Unlock(ctx, "running")

			// Lock candidate datastore
			if _, err := client.Lock(ctx, "candidate"); err != nil {
				return fmt.Errorf("failed to lock candidate datastore: %s", FormatNetconfError(err))
			}
			defer client.Unlock(ctx, "candidate")
		}

		if _, err := client.EditConfig(ctx, "candidate", body); err != nil {
			return fmt.Errorf("failed to edit config: %s", FormatNetconfError(err))
		}

		if commit {
			if _, err := client.Commit(ctx); err != nil {
				return fmt.Errorf("failed to commit config: %s", FormatNetconfError(err))
			}
		}
	} else {
		// Lock running datastore
		if _, err := client.Lock(ctx, "running"); err != nil {
			return fmt.Errorf("failed to lock running datastore: %s", FormatNetconfError(err))
		}
		defer client.Unlock(ctx, "running")

		if _, err := client.EditConfig(ctx, "running", body); err != nil {
			return fmt.Errorf("failed to edit config: %s", FormatNetconfError(err))
		}
	}
	return nil
}

// Commit commits the candidate datastore to the running datastore
func Commit(ctx context.Context, client *netconf.Client) error {
	if err := client.Open(); err != nil {
		return fmt.Errorf("failed to open NETCONF connection: %w", err)
	}

	candidate := client.ServerHasCapability("urn:ietf:params:netconf:capability:candidate:1.0")

	if candidate {
		// Lock running datastore
		if _, err := client.Lock(ctx, "running"); err != nil {
			return fmt.Errorf("failed to lock running datastore: %s", FormatNetconfError(err))
		}
		defer client.Unlock(ctx, "running")

		if _, err := client.Commit(ctx); err != nil {
			return fmt.Errorf("failed to commit config: %s", FormatNetconfError(err))
		}
	}
	return nil
}

// SaveConfig saves the running configuration to startup configuration.
// This is equivalent to 'copy running-config startup-config' in the CLI.
// Uses the cisco-ia:save-config RPC operation.
func SaveConfig(ctx context.Context, client *netconf.Client) error {
	// Ensure connection is open
	if err := client.Open(); err != nil {
		return fmt.Errorf("failed to open NETCONF connection: %w", err)
	}

	// Build NETCONF body for save-config RPC
	body := netconf.Body{}
	body = SetFromXPath(body, "/cisco-ia:save-config", "")
	body = body.SetAttr("save-config", "xmlns", "http://cisco.com/yang/cisco-ia")

	// Execute the save-config RPC
	if _, err := client.RPC(ctx, body.Res()); err != nil {
		return fmt.Errorf("failed to save config: %s", FormatNetconfError(err))
	}

	return nil
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
			// Build predicates in order
			predicates := make([]string, 0, len(keys))
			for _, kv := range keys {
				// Remove namespace prefix from key name
				keyName := removeNamespacePrefix(kv.Key)
				predicates = append(predicates, fmt.Sprintf("%s='%s'", keyName, kv.Value))
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

// KeyValue represents a key-value pair with preserved order
type KeyValue struct {
	Key   string
	Value string
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
// When multiple sibling elements exist at the same path, adds namespace to all of them.
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

			// Check for namespace exceptions first
			namespace, ok := namespaceExceptions[prefix]
			if !ok {
				// Use default pattern if no exception exists
				namespace = namespaceBaseURL + prefix
			}

			// Check if there are multiple elements at this path
			countPath := currentPath + ".#"
			count := xmldot.Get(body.Res(), countPath).Int()

			if count > 1 {
				// Multiple elements - add namespace to all that don't have it
				for i := 0; i < int(count); i++ {
					indexedXmlnsPath := fmt.Sprintf("%s.%d.@xmlns", currentPath, i)
					if !xmldot.Get(body.Res(), indexedXmlnsPath).Exists() {
						body = body.Set(indexedXmlnsPath, namespace)
					}
				}
			} else {
				// Single element or first element
				xmlnsPath := currentPath + ".@xmlns"
				if !xmldot.Get(body.Res(), xmlnsPath).Exists() {
					body = body.Set(xmlnsPath, namespace)
				}
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
// The ensureStructure parameter controls whether an empty element should be created at the final path
// if it doesn't already exist. Set to false when a value will be immediately set afterward.
func buildXPathStructure(body netconf.Body, xPath string, ensureStructure bool) (netconf.Body, []string) {
	// Remove leading slash if present
	xPath = strings.TrimPrefix(xPath, "/")

	// Split into segments while respecting bracket boundaries
	segments := splitXPathSegments(xPath)

	// Build path incrementally, creating each element
	pathSegments := make([]string, 0, len(segments))

	for i, segment := range segments {
		// Parse segment: element[key='value'][key2='value2'] -> element, []KeyValue
		elementName, keys := parseXPathSegment(segment)

		// Add element name to path (without predicates)
		pathSegments = append(pathSegments, elementName)
		fullPath := strings.Join(pathSegments[:i+1], ".")

		// If this segment has keys, set all key values in order
		if len(keys) > 0 {
			for _, kv := range keys {
				keyPath := fullPath + "." + kv.Key
				body = setWithNamespaces(body, keyPath, kv.Value)
			}
		}
	}

	// Optionally ensure the complete path structure exists, including non-predicate elements.
	// Only create the structure if requested and if it doesn't already exist.
	if ensureStructure && len(pathSegments) > 0 {
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
// Returns: (elementName, []KeyValue) - order is preserved from the XPath
func parseXPathSegment(segment string) (string, []KeyValue) {
	// Check for predicate: element[...]
	if idx := strings.Index(segment, "["); idx != -1 {
		elementName := segment[:idx]
		keys := make([]KeyValue, 0)

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
						keyName := strings.TrimSpace(condition[:eqIdx])
						value := condition[eqIdx+1:]
						// Remove quotes
						keyValue := strings.Trim(value, `'"`)
						keys = append(keys, KeyValue{Key: keyName, Value: keyValue})
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
	// Build the structure and set operation="remove" on the last element
	body, pathSegments := buildXPathStructure(body, xPath, false)

	// Set operation="remove" on the last element
	if len(pathSegments) > 0 {
		targetPath := strings.Join(pathSegments, ".")
		operationPath := targetPath + ".@operation"
		body = setWithNamespaces(body, operationPath, "remove")
	}

	// Clean up any redundant child operations
	body = CleanupRedundantRemoveOperations(body)

	return body
}

// CleanupRedundantRemoveOperations removes redundant operation="remove" attributes from child elements
// when a parent element already has operation="remove". This is called after the full NETCONF payload
// has been built to sanitize the XML before sending to the device.
//
// Example:
//
//	Input:  <trap operation="remove"><severity operation="remove"></severity></trap>
//	Output: <trap operation="remove"></trap>
//
// This prevents NETCONF errors like '"" is not a valid value' for empty child elements.
func CleanupRedundantRemoveOperations(body netconf.Body) netconf.Body {
	if body.Res() == "" {
		return body
	}

	// Start recursive cleanup from root level
	body = cleanupRedundantRemoveRecursive(body, "", make(map[string]bool))

	return body
}

// cleanupRedundantRemoveRecursive walks through the XML tree and removes redundant child operations.
// It processes the tree level by level, tracking which elements have operation="remove" to avoid
// checking the same elements multiple times.
func cleanupRedundantRemoveRecursive(body netconf.Body, basePath string, visited map[string]bool) netconf.Body {
	xml := body.Res()
	if xml == "" {
		return body
	}

	// Skip if we've already processed this path
	if visited[basePath] {
		return body
	}
	visited[basePath] = true

	// Get content at current path
	var currentResult xmldot.Result
	if basePath == "" {
		currentResult = xmldot.Get(xml, "")
	} else {
		currentResult = xmldot.Get(xml, basePath)
		if !currentResult.Exists() {
			return body
		}
	}

	// Find immediate child elements at this level by parsing Raw content
	// We need to find child element names, not search for operation="remove" yet
	rawXML := xml
	if basePath != "" {
		rawXML = currentResult.Raw
	}

	// Extract child element names from the raw XML
	childPattern := regexp.MustCompile(`<([a-zA-Z0-9\-_]+)[\s>]`)
	matches := childPattern.FindAllStringSubmatch(rawXML, -1)

	seen := make(map[string]bool)
	for _, match := range matches {
		if len(match) > 1 {
			elementName := match[1]
			if seen[elementName] {
				continue
			}
			seen[elementName] = true

			// Build path to this child element
			var childPath string
			if basePath == "" {
				childPath = elementName
			} else {
				childPath = basePath + "." + elementName
			}

			// Check if this child has operation="remove"
			operationPath := childPath + ".@operation"
			hasRemove := xmldot.Get(body.Res(), operationPath).String() == "remove"

			if hasRemove {
				// This element has operation="remove", clean up its children
				body = removeAllChildOperations(body, childPath)
			}

			// Recursively process this child's descendants
			body = cleanupRedundantRemoveRecursive(body, childPath, visited)
		}
	}

	return body
}

// removeAllChildOperations removes all child elements that have operation="remove" from under the given parent path.
// When a parent has operation="remove", all its children will be removed by the device anyway, so we delete
// child elements entirely to avoid redundant operations and prevent empty element errors.
// However, we preserve key elements (list identifiers) as they may be needed for identification.
func removeAllChildOperations(body netconf.Body, parentPath string) netconf.Body {
	xml := body.Res()

	parentResult := xmldot.Get(xml, parentPath)
	if !parentResult.Exists() {
		return body
	}

	rawXML := parentResult.Raw
	if rawXML == "" {
		return body
	}

	// Find all immediate child elements at this level
	childPattern := regexp.MustCompile(`<([a-zA-Z0-9\-_]+)[\s>]`)
	matches := childPattern.FindAllStringSubmatch(rawXML, -1)

	seen := make(map[string]bool)
	for _, match := range matches {
		if len(match) > 1 {
			childName := match[1]
			if seen[childName] {
				continue
			}
			seen[childName] = true

			childPath := parentPath + "." + childName

			// Check if this child exists and what it contains
			childResult := xmldot.Get(body.Res(), childPath)
			if !childResult.Exists() {
				continue
			}

			// Check if this child has operation="remove"
			operationPath := childPath + ".@operation"
			hasRemove := xmldot.Get(body.Res(), operationPath).String() == "remove"

			if hasRemove {
				// This child has operation="remove" under a parent that also has operation="remove"
				// Check if it contains keys - if so, keep it but remove the operation attribute
				// If it doesn't contain keys (empty or only has other operations), delete it entirely
				childContent := childResult.Raw

				if shouldDeleteEntireElement(childContent) || !containsKeys(childContent) {
					// No keys found, safe to delete entirely
					body = body.Delete(childPath)
				} else {
					// Has keys, just remove the operation attribute but keep the element
					body = body.Delete(operationPath)

					// Now recursively clean this element's children
					// Since the parent (host) has operation="remove", ALL descendants are redundant
					body = removeAllDescendantOperations(body, childPath)
				}
			} else {
				// Child doesn't have operation="remove", but we still need to check its descendants
				// because they might have redundant operation="remove" attributes
				body = removeAllDescendantOperations(body, childPath)
			}
		}
	}

	return body
}

// shouldDeleteEntireElement determines if an element should be deleted entirely or just have
// its operation attribute removed. Elements with only keys and operation="remove" should be
// deleted entirely to avoid empty elements like <severity operation="remove"></severity>.
func shouldDeleteEntireElement(content string) bool {
	if content == "" || strings.TrimSpace(content) == "" {
		return true
	}

	// Check if content only contains simple key elements (elements with text content only)
	// Pattern: <elementName>text</elementName>
	keyPattern := regexp.MustCompile(`^\s*(?:<[a-zA-Z0-9\-_]+>[^<]+</[a-zA-Z0-9\-_]+>\s*)+$`)
	return keyPattern.MatchString(content)
}

// containsKeys checks if the content contains key elements (simple elements with text content).
// Returns true if there are any key-like elements: <elementName>value</elementName>
func containsKeys(content string) bool {
	if content == "" || strings.TrimSpace(content) == "" {
		return false
	}

	// Pattern matches key elements: <elementName>text</elementName>
	keyPattern := regexp.MustCompile(`<[a-zA-Z0-9\-_]+>[^<]+</[a-zA-Z0-9\-_]+>`)
	return keyPattern.MatchString(content)
}

// removeAllDescendantOperations removes all operation="remove" attributes and elements from
// descendants of the given path. This is used when a parent has operation="remove" and we've
// decided to keep a child (because it has keys), but need to clean up all of that child's descendants.
func removeAllDescendantOperations(body netconf.Body, parentPath string) netconf.Body {
	xml := body.Res()

	parentResult := xmldot.Get(xml, parentPath)
	if !parentResult.Exists() {
		return body
	}

	rawXML := parentResult.Raw
	if rawXML == "" {
		return body
	}

	// Find all child elements at this level
	childPattern := regexp.MustCompile(`<([a-zA-Z0-9\-_]+)[\s>]`)
	matches := childPattern.FindAllStringSubmatch(rawXML, -1)

	seen := make(map[string]bool)
	for _, match := range matches {
		if len(match) > 1 {
			elementName := match[1]
			if seen[elementName] {
				continue
			}
			seen[elementName] = true

			childPath := parentPath + "." + elementName

			// Check if this child exists
			childResult := xmldot.Get(body.Res(), childPath)
			if !childResult.Exists() {
				continue
			}

			// Check if this child has operation="remove"
			operationPath := childPath + ".@operation"
			hasRemove := xmldot.Get(body.Res(), operationPath).String() == "remove"

			if hasRemove {
				// This descendant has operation="remove", delete it entirely
				// (we don't need to preserve it since the ancestor will be removed)
				body = body.Delete(childPath)
			} else {
				// No operation attribute, but recurse to check its children
				body = removeAllDescendantOperations(body, childPath)
			}
		}
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
	// Determine if we need to create empty structure
	// Only create empty structure if no value will be set
	hasValue := value != nil && value != ""
	ensureStructure := !hasValue

	body, pathSegments := buildXPathStructure(body, xPath, ensureStructure)

	// Only set the value if it's not nil and not empty string
	// This prevents overwriting key children when no value is needed
	if hasValue && len(pathSegments) > 0 {
		fullPath := strings.Join(pathSegments, ".")
		body = setWithNamespaces(body, fullPath, value)
	}

	return body
}

// AppendFromXPath creates all elements in an XPath and appends a value to a list by using
// the ".-1" syntax. This is useful for adding multiple items to a list without keys.
// The function automatically appends ".-1" to the final element in the path.
//
// Example:
//
//	body := netconf.Body{}
//	body = AppendFromXPath(body, "native/route-map/rule/match/ip/address", "10")
//	body = AppendFromXPath(body, "native/route-map/rule/match/ip/address", "20")
//	// Result: <native><route-map><rule><match><ip>
//	//           <address>10</address>
//	//           <address>20</address>
//	//         </ip></match></rule></route-map></native>
//
// Note: This function is designed for simple list items without keys. For lists with keys,
// use SetFromXPath with predicates instead.
func AppendFromXPath(body netconf.Body, xPath string, value any) netconf.Body {
	// Determine if we need to create empty structure
	hasValue := value != nil && value != ""
	ensureStructure := !hasValue

	body, pathSegments := buildXPathStructure(body, xPath, ensureStructure)

	// Append to the list using .-1 syntax
	if hasValue && len(pathSegments) > 0 {
		fullPath := strings.Join(pathSegments, ".") + ".-1"
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

	// Extract the final element name
	finalSegment := segments[len(segments)-1]
	finalElement, keys := parseXPathSegment(finalSegment)
	finalElementClean := removeNamespacePrefix(finalElement)

	// Build parent structure (everything except the final element)
	if len(segments) > 1 {
		// Multi-segment path: wrap the content with the final element tag
		wrappedContent := "<" + finalElementClean + ">" + value + "</" + finalElementClean + ">"
		parentXPath := "/" + strings.Join(segments[:len(segments)-1], "/")
		// Build structure but don't create empty elements (may already have content)
		body, _ = buildXPathStructure(body, parentXPath, false)

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
		// Single-segment path - SetRaw will wrap the content for us
		// Just prepare the inner content (with keys if present)
		innerContent := value
		if len(keys) > 0 {
			tempBody := netconf.Body{}
			for _, kv := range keys {
				tempBody = setWithNamespaces(tempBody, kv.Key, kv.Value)
			}
			innerContent = tempBody.Res() + value
		}

		// Use the element name as the dotPath for SetRaw
		// SetRaw will create <finalElementClean>innerContent</finalElementClean>
		// Check if this element already exists at the root
		existingXML := xmldot.Get(body.Res(), finalElementClean).Raw
		if existingXML != "" {
			// Append as sibling - need to wrap each occurrence
			wrappedNew := "<" + finalElementClean + ">" + innerContent + "</" + finalElementClean + ">"
			combinedXML := existingXML + wrappedNew
			body = body.SetRaw(finalElementClean, combinedXML)
		} else {
			// First element at this path - SetRaw will wrap it
			body = body.SetRaw(finalElementClean, innerContent)
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

// GetFromXPath converts an XPath expression to a xmldot path and retrieves the result.
// Uses manual filtering to correctly handle both single and multiple elements with predicates.
// Supports the same XPath formats as SetFromXPath:
//   - Single: /interface[name='GigabitEthernet1']
//   - Multiple predicates: /interface[name='GigabitEthernet1'][vrf='VRF1']
//   - Combined with 'and': /interface[name='GigabitEthernet1' and vrf='VRF1']
//   - Values with slashes: /interface[name='GigabitEthernet1/0/1']
//   - Nested paths: /native/interface[name='Gi1']/ip/address
//
// Example: /native/interface[name='Gi1']/ip/address
// Processes path segments, validates predicates, and constructs the final xmldot path
func GetFromXPath(res xmldot.Result, xPath string) xmldot.Result {
	// Remove leading slash if present
	xPath = strings.TrimPrefix(xPath, "/")

	// Split into segments while respecting bracket boundaries
	segments := splitXPathSegments(xPath)

	// Use manual filtering to handle both single and multiple elements with predicates.
	// xmldot's #() filter syntax doesn't work reliably with single elements, so we
	// always use the manual path which correctly handles both cases.
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
					for _, kv := range keys {
						// Remove namespace prefix from key name
						keyName := removeNamespacePrefix(kv.Key)
						keyResult := item.Get(keyName)
						if !keyResult.Exists() || keyResult.String() != kv.Value {
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
				for _, kv := range keys {
					// Remove namespace prefix from key name
					keyName := removeNamespacePrefix(kv.Key)
					keyResult := currentResult.Get(keyName)
					if !keyResult.Exists() || keyResult.String() != kv.Value {
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

	// Check if the final path points to multiple elements (array)
	// If so, use the #. syntax to return an array result that ForEach can iterate over
	countPath := finalPath + ".#"
	count := xmldot.Get(xml, countPath).Int()
	if count > 1 {
		// Multiple elements exist - use #. syntax to get array result
		// Split the path to insert #. before the last segment
		if len(pathSoFar) >= 2 {
			// Build parent path and use #. syntax: parent.#.child
			parentPath := strings.Join(pathSoFar[:len(pathSoFar)-1], ".")
			childName := pathSoFar[len(pathSoFar)-1]
			arrayPath := parentPath + ".#." + childName
			return xmldot.Get(xml, arrayPath)
		}
	}

	return xmldot.Get(xml, finalPath)
}

// IsListPath checks if an XPath represents a list item (ends with a predicate).
// List items have predicates like [name='value'] or [name=value] at the end of their XPath,
// while containers/singletons don't end with predicates.
//
// This is useful for determining if an empty GetConfig response should be
// interpreted as "resource not found" (for list items) vs other semantics.
//
// Parameters:
//   - xPath: The XPath to check
//
// Returns:
//   - true if the path ends with a predicate (is a list item)
//   - false if the path does not end with a predicate (is a container/singleton)
//
// Examples:
//   - "/native/interface/Vlan[name=10]" → true (ends with predicate)
//   - "/native/interface/GigabitEthernet[name='1/0/1']" → true (ends with predicate)
//   - "/native/router/bgp[id=65000]/neighbor" → false (has predicate but doesn't end with one)
//   - "/native/clock" → false (container)
//   - "/native/hostname" → false (singleton)
func IsListPath(xPath string) bool {
	// Trim whitespace
	xPath = strings.TrimSpace(xPath)

	// Check if the path ends with a closing bracket (predicate)
	// A list path will end with something like [name='value']
	return strings.HasSuffix(xPath, "]")
}

// IsGetConfigResponseEmpty checks if a GetConfig response has an empty <data> element.
// Returns true if the response contains <data></data> with no child elements,
// indicating that the requested configuration does not exist on the device.
//
// This is useful for determining if a resource exists before attempting to parse
// its attributes, particularly during Read operations or import.
//
// IMPORTANT: This should typically be combined with IsListPath() to only treat
// empty responses as "not found" for list items:
//
//	if helpers.IsGetConfigResponseEmpty(&res) && helpers.IsListPath(state.getXPath()) {
//	    // List item does not exist
//	    resp.State.RemoveResource(ctx)
//	    return
//	}
//
// Parameters:
//   - res: The NETCONF response from GetConfig operation
//
// Returns:
//   - true if the data element is empty (no child elements)
//   - false if the data element contains configuration
//
// Example usage:
//
//	res, err := device.NetconfClient.GetConfig(ctx, "running", filter)
//	if err != nil {
//	    return err
//	}
//	if helpers.IsGetConfigResponseEmpty(&res) && helpers.IsListPath(state.getXPath()) {
//	    // Resource does not exist (list item)
//	    resp.State.RemoveResource(ctx)
//	    return
//	}
//	// Parse the configuration
//	state.fromBodyXML(ctx, res.Res)
func IsGetConfigResponseEmpty(res *netconf.Res) bool {
	if res == nil {
		return true
	}

	// Get the data element from the response
	dataResult := res.Res.Get("data")

	// If data element doesn't exist, consider it empty
	if !dataResult.Exists() {
		return true
	}

	// Check if data element has any children using Map()
	// An empty <data></data> element will have an empty map
	children := dataResult.Map()

	// If the map is empty (no child elements), the response is empty
	// The "%" key represents direct text content, which we also want to ignore
	// since whitespace-only content is not meaningful configuration
	for key := range children {
		if key != "%" {
			return false
		}
	}

	return true
}
