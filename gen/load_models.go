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

//go:build ignore

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var models = []string{
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-native.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/cisco-semver.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/ietf-inet-types.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-types.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-features.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-interface-common.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-parser.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-license.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-crypto.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-tunnel.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-line.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-logging.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-ip.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-ipv6.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-cef.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-interfaces.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-isis.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-snmp.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-segment-routing.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-ospf-obsolete.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-vlan.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-vtp.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-aaa.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-device-tracking.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-nd.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-flow.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-multicast.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-l2vpn.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-ethernet.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-mpls.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-ethernet-oam.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-ethernet-cfm-efp.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-pppoe.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-bgp.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-ospfv3.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-ospf.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-switch.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/ietf-yang-types.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-igmp.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-arp.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-dhcp.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-hsrp.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-template.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-policy.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-atm.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-sisf.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-acl.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-object-group.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-route-map.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-nhrp.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-location.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-transceiver-monitor.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-icmp.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-mdt-cfg.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-mdt-common-defs.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-ntp.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-spanning-tree.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-aaa.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-cts.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-crypto.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-tunnel.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-cdp.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-bfd.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-dot1x.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-mdns-gateway.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-udld.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-switch.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-sanet.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-http.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-flow.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-diagnostics.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-pnp.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-nbar.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-device-sensor.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-lldp.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-nat.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-call-home.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-ppp.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-track.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-eem.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-platform.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-mld.yang",
	// "https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/17151/Cisco-IOS-XE-sla.yang", // Disabled: Complex YANG 1.1 multiple augment patterns exceed goyang v1.6.3 capabilities
}

const (
	modelsPath = "./gen/models/"
)

func main() {
	for _, model := range models {
		f := strings.Split(model, "/")
		filename := f[len(f)-1]
		path := filepath.Join(modelsPath, filename)
		if _, err := os.Stat(path); err != nil {
			err := downloadModel(path, model)
			if err != nil {
				panic(err)
			}
			fmt.Println("Downloaded model: " + path)

			// Apply patches to specific models after download
			err = patchModel(path, filename)
			if err != nil {
				panic(err)
			}
		}
	}
}

func downloadModel(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// patchModel applies necessary patches to specific YANG models after download
func patchModel(filepath, filename string) error {
	switch filename {
	case "Cisco-IOS-XE-eem.yang":
		return patchLoggingLevelTypeModel(filepath)
	case "Cisco-IOS-XE-platform.yang":
		return patchLoggingLevelTypeModel(filepath)
	case "Cisco-IOS-XE-sla.yang":
		return patchSLAModel(filepath)
	default:
		// No patch needed for other models
		return nil
	}
}

// extractTypeDefinition extracts a typedef from a YANG model file
func extractTypeDefinition(yangFile, typeName string) (string, error) {
	content, err := os.ReadFile(yangFile)
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(content), "\n")
	var typeDefLines []string
	var inTypeDef bool
	var braceCount int

	typedefPattern := fmt.Sprintf("typedef %s {", typeName)

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		// Start of typedef
		if strings.Contains(trimmedLine, typedefPattern) {
			inTypeDef = true
			// Extract just the type definition part, not the typedef wrapper
			continue
		}

		if inTypeDef {
			// Count braces to know when typedef ends
			braceCount += strings.Count(line, "{")
			braceCount -= strings.Count(line, "}")

			// Add line to type definition
			if braceCount > 0 {
				// Remove the leading "type " from the first line if present
				if len(typeDefLines) == 0 && strings.HasPrefix(strings.TrimSpace(line), "type ") {
					line = strings.TrimPrefix(strings.TrimSpace(line), "type ")
					line = "    " + line // Maintain indentation
				}
				typeDefLines = append(typeDefLines, line)
			} else if braceCount == 0 && strings.Contains(trimmedLine, "}") {
				// This is the final closing brace - add it and break
				typeDefLines = append(typeDefLines, line)
				break
			}
		}
	}

	if len(typeDefLines) == 0 {
		return "", fmt.Errorf("typedef %s not found in %s", typeName, yangFile)
	}

	return strings.Join(typeDefLines, "\n"), nil
}

// patchLoggingLevelTypeModel fixes the ios:logging-level-type reference issue in models
func patchLoggingLevelTypeModel(filepath string) error {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	contentStr := string(content)

	// Check if already patched to avoid double-patching
	if !strings.Contains(contentStr, "type ios:logging-level-type;") {
		return nil // Already patched
	}

	// Dynamically extract the type definition from the logging model
	loggingModelPath := "./gen/models/Cisco-IOS-XE-logging.yang"
	typeDefinition, err := extractTypeDefinition(loggingModelPath, "logging-level-type")
	if err != nil {
		return fmt.Errorf("failed to extract logging-level-type definition: %v", err)
	}

	// Replace the entire type statement with the dynamically extracted type definition
	// Need to adjust indentation: typedef uses 4-space indent, but inline usage needs 2-space
	lines := strings.Split(typeDefinition, "\n")
	var adjustedLines []string
	for _, line := range lines {
		// Remove the typedef-level indentation (4 spaces) and replace with inline-level (2 spaces)
		if strings.HasPrefix(line, "    ") {
			// Remove 4 spaces and add 2 spaces
			adjustedLine := "  " + strings.TrimPrefix(line, "    ")
			adjustedLines = append(adjustedLines, adjustedLine)
		} else {
			adjustedLines = append(adjustedLines, line)
		}
	}
	adjustedDefinition := strings.Join(adjustedLines, "\n")
	// Trim any leading whitespace from the first line
	adjustedDefinition = strings.TrimSpace(adjustedDefinition)
	contentStr = strings.ReplaceAll(contentStr, "type ios:logging-level-type;", "type "+adjustedDefinition)

	// Write the patched content back to file
	err = os.WriteFile(filepath, []byte(contentStr), 0644)
	if err != nil {
		return err
	}

	fmt.Println("Applied patch to: " + filepath)
	return nil
}

// patchSLAModel handles patches for the SLA model to resolve YANG 1.1 multiple augment conflicts
func patchSLAModel(filepath string) error {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	contentStr := string(content)

	// Check if already patched to avoid double-patching
	if strings.Contains(contentStr, "// PATCHED: Multiple augment conflict resolution") {
		return nil // Already patched
	}

	// Apply multiple patches in sequence
	contentStr, err = patchSLAMultipleAugments(contentStr)
	if err != nil {
		return fmt.Errorf("failed to patch multiple augments: %v", err)
	}

	// Also remove problematic if-feature statements that may cause goyang parsing issues
	contentStr = strings.ReplaceAll(contentStr, `if-feature "ios-features:service-performance-ooop";`, `// if-feature "ios-features:service-performance-ooop";`)
	contentStr = strings.ReplaceAll(contentStr, `if-feature "ios-features:service-performance";`, `// if-feature "ios-features:service-performance";`)

	// Add patch marker
	contentStr = "// PATCHED: Multiple augment conflict resolution for goyang v1.6.3 compatibility\n" + contentStr

	// Write the patched content back to file
	err = os.WriteFile(filepath, []byte(contentStr), 0644)
	if err != nil {
		return err
	}

	fmt.Println("Applied SLA model patch (resolved multiple augment conflicts): " + filepath)
	return nil
}

// patchSLAMultipleAugments resolves the YANG 1.1 multiple augment statements issue
// by replacing conflicting uses statements with expanded inline definitions
func patchSLAMultipleAugments(content string) (string, error) {
	// First, handle ip-sla-service-performance-sub-command-grouping conflicts
	baseGrouping, err := extractGroupingDefinition(content, "ip-sla-service-performance-sub-command-grouping")
	if err != nil {
		return "", fmt.Errorf("failed to extract base grouping: %v", err)
	}

	// Replace the first conflicting uses statement (ethernet case)
	content, err = replaceUsesWithInlineEthernet(content, baseGrouping)
	if err != nil {
		return "", fmt.Errorf("failed to replace ethernet uses: %v", err)
	}

	// Replace the second conflicting uses statement (ip case)
	content, err = replaceUsesWithInlineIP(content, baseGrouping)
	if err != nil {
		return "", fmt.Errorf("failed to replace ip uses: %v", err)
	}

	// Second, handle config-ip-sla-mpls-lsp-submode-grouping conflicts
	mplsGrouping, err := extractGroupingDefinition(content, "config-ip-sla-mpls-lsp-submode-grouping")
	if err != nil {
		return "", fmt.Errorf("failed to extract mpls grouping: %v", err)
	}

	// Replace multiple MPLS uses statements with unique history augments
	content, err = replaceMPLSUsesStatements(content, mplsGrouping)
	if err != nil {
		return "", fmt.Errorf("failed to replace mpls uses: %v", err)
	}

	return content, nil
}

// extractGroupingDefinition extracts a complete grouping definition from the YANG content
func extractGroupingDefinition(content, groupingName string) (string, error) {
	lines := strings.Split(content, "\n")
	var groupingLines []string
	var inGrouping bool
	var braceCount int

	groupingPattern := fmt.Sprintf("grouping %s {", groupingName)

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		// Start of grouping
		if strings.Contains(trimmedLine, groupingPattern) {
			inGrouping = true
			groupingLines = append(groupingLines, line)
			braceCount = 1
			continue
		}

		if inGrouping {
			// Count braces to know when grouping ends
			braceCount += strings.Count(line, "{")
			braceCount -= strings.Count(line, "}")

			groupingLines = append(groupingLines, line)

			if braceCount == 0 {
				// End of grouping reached
				break
			}
		}
	}

	if len(groupingLines) == 0 {
		return "", fmt.Errorf("grouping %s not found", groupingName)
	}

	return strings.Join(groupingLines, "\n"), nil
}

// replaceUsesWithInlineEthernet replaces the ethernet case uses statement with inline content
func replaceUsesWithInlineEthernet(content, baseGrouping string) (string, error) {
	// Find the problematic uses statement in ethernet case
	ethernetUsesStart := `uses ip-sla-service-performance-sub-command-grouping {`
	ethernetContext := `augment "packet" {
                              leaf dscp {` // Unique identifier for ethernet case

	startIdx, endIdx, usesBlock, err := findUsesBlock(content, ethernetUsesStart, ethernetContext)
	if err != nil {
		return content, nil // Already patched or pattern not found
	}

	// Extract the augment contents dynamically
	profileAugment, err := extractAugmentContent(usesBlock, "profile")
	if err != nil {
		return "", fmt.Errorf("failed to extract profile augment: %v", err)
	}

	measurementAugment, err := extractAugmentContent(usesBlock, "measurement-type")
	if err != nil {
		return "", fmt.Errorf("failed to extract measurement-type augment: %v", err)
	}

	// Create the replacement inline content dynamically
	replacement, err := createInlineExpansion(content, baseGrouping, "ethernet", profileAugment, measurementAugment)
	if err != nil {
		return "", fmt.Errorf("failed to create inline expansion: %v", err)
	}

	// Replace the uses block with inline content
	newContent := content[:startIdx] + replacement + content[endIdx:]

	return newContent, nil
}

// replaceUsesWithInlineIP replaces the ip case uses statement with inline content
func replaceUsesWithInlineIP(content, baseGrouping string) (string, error) {
	// Find the problematic uses statement in ip case
	ipUsesStart := `uses ip-sla-service-performance-sub-command-grouping {`
	ipContext := `augment "packet" {
                              container tunnel {` // Unique identifier for ip case

	startIdx, endIdx, usesBlock, err := findUsesBlock(content, ipUsesStart, ipContext)
	if err != nil {
		return content, nil // Already patched or pattern not found
	}

	// Extract the augment contents dynamically
	profileAugment, err := extractAugmentContent(usesBlock, "profile")
	if err != nil {
		return "", fmt.Errorf("failed to extract profile augment: %v", err)
	}

	measurementAugment, err := extractAugmentContent(usesBlock, "measurement-type")
	if err != nil {
		return "", fmt.Errorf("failed to extract measurement-type augment: %v", err)
	}

	// Create the replacement inline content dynamically
	replacement, err := createInlineExpansion(content, baseGrouping, "ip", profileAugment, measurementAugment)
	if err != nil {
		return "", fmt.Errorf("failed to create inline expansion: %v", err)
	}

	// Replace the uses block with inline content
	newContent := content[:startIdx] + replacement + content[endIdx:]

	return newContent, nil
}

// findUsesBlock finds a uses statement block and returns its start/end positions and content
func findUsesBlock(content, usesStart, uniqueContext string) (int, int, string, error) {
	// Find the uses statement by locating the unique context within it
	contextIdx := strings.Index(content, uniqueContext)
	if contextIdx == -1 {
		return 0, 0, "", fmt.Errorf("unique context not found")
	}

	// Search backwards to find the start of the uses statement
	searchStart := contextIdx - 1000 // Search back up to 1000 characters
	if searchStart < 0 {
		searchStart = 0
	}

	usesIdx := strings.LastIndex(content[searchStart:contextIdx], usesStart)
	if usesIdx == -1 {
		return 0, 0, "", fmt.Errorf("uses statement not found")
	}
	startIdx := searchStart + usesIdx

	// Find the matching closing brace for this uses statement
	braceCount := 0
	endIdx := startIdx
	inUses := false

	for i, char := range content[startIdx:] {
		if char == '{' {
			braceCount++
			inUses = true
		} else if char == '}' {
			braceCount--
			if inUses && braceCount == 0 {
				endIdx = startIdx + i + 1
				break
			}
		}
	}

	if braceCount != 0 {
		return 0, 0, "", fmt.Errorf("unmatched braces in uses statement")
	}

	usesBlock := content[startIdx:endIdx]
	return startIdx, endIdx, usesBlock, nil
}

// extractAugmentContent extracts the content of a specific augment statement
func extractAugmentContent(usesBlock, augmentPath string) (string, error) {
	augmentPattern := fmt.Sprintf(`augment "%s" {`, augmentPath)
	startIdx := strings.Index(usesBlock, augmentPattern)
	if startIdx == -1 {
		return "", fmt.Errorf("augment %s not found", augmentPath)
	}

	// Find the matching closing brace for this augment
	braceCount := 0
	endIdx := startIdx
	inAugment := false

	for i, char := range usesBlock[startIdx:] {
		if char == '{' {
			braceCount++
			inAugment = true
		} else if char == '}' {
			braceCount--
			if inAugment && braceCount == 0 {
				endIdx = startIdx + i + 1
				break
			}
		}
	}

	if braceCount != 0 {
		return "", fmt.Errorf("unmatched braces in augment %s", augmentPath)
	}

	// Extract the content inside the augment (excluding the augment wrapper)
	lines := strings.Split(usesBlock[startIdx:endIdx], "\n")
	var contentLines []string

	for i, line := range lines {
		if i == 0 {
			// Skip the augment declaration line
			continue
		}
		if i == len(lines)-1 && strings.TrimSpace(line) == "}" {
			// Skip the final closing brace
			continue
		}
		contentLines = append(contentLines, line)
	}

	return strings.Join(contentLines, "\n"), nil
}

// createInlineExpansion creates the complete inline expansion by merging base grouping with augments
func createInlineExpansion(content, baseGrouping, caseType, profileAugment, measurementAugment string) (string, error) {
	// Extract the base grouping content (without the grouping wrapper)
	baseContent, err := extractGroupingContent(baseGrouping)
	if err != nil {
		return "", fmt.Errorf("failed to extract base grouping content: %v", err)
	}

	// Create the inline expansion
	expansion := fmt.Sprintf("// Inline expansion of ip-sla-service-performance-sub-command-grouping for %s case\n", caseType)

	// Add the base content
	expansion += baseContent

	// Add the expanded profile container with unique naming and recursively expand nested uses
	expansion += fmt.Sprintf("\n                      container profile-%s {\n", caseType)
	expansion += "                        description\n"
	expansion += fmt.Sprintf("                          \"Profile configuration for %s case\";\n", caseType)

	// Recursively expand any nested uses statements in the profile augment
	expandedProfileAugment, err := expandNestedUsesStatements(content, profileAugment, caseType+"-profile")
	if err != nil {
		return "", fmt.Errorf("failed to expand nested uses in profile: %v", err)
	}
	expansion += indentContent(expandedProfileAugment, "                        ")
	expansion += "\n                      }"

	// Add the expanded measurement-type container with unique naming and recursively expand nested uses
	expansion += fmt.Sprintf("\n                      container measurement-type-%s {\n", caseType)
	expansion += "                        description\n"
	expansion += fmt.Sprintf("                          \"Measurement type configuration for %s case\";\n", caseType)

	// Recursively expand any nested uses statements in the measurement augment
	expandedMeasurementAugment, err := expandNestedUsesStatements(content, measurementAugment, caseType+"-measurement")
	if err != nil {
		return "", fmt.Errorf("failed to expand nested uses in measurement: %v", err)
	}
	expansion += indentContent(expandedMeasurementAugment, "                        ")
	expansion += "\n                      }"

	return expansion, nil
}

// extractGroupingContent extracts the content inside a grouping (without the grouping wrapper)
func extractGroupingContent(groupingDef string) (string, error) {
	lines := strings.Split(groupingDef, "\n")
	var contentLines []string

	for i, line := range lines {
		if i == 0 {
			// Skip the grouping declaration line
			continue
		}
		if i == len(lines)-1 && strings.TrimSpace(line) == "}" {
			// Skip the final closing brace
			continue
		}
		contentLines = append(contentLines, line)
	}

	return strings.Join(contentLines, "\n"), nil
}

// indentContent adds consistent indentation to content
func indentContent(content, indent string) string {
	lines := strings.Split(content, "\n")
	var indentedLines []string

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			indentedLines = append(indentedLines, "")
		} else {
			indentedLines = append(indentedLines, indent+line)
		}
	}

	return strings.Join(indentedLines, "\n")
}

// replaceMPLSUsesStatements handles multiple MPLS uses statements with history augment conflicts
func replaceMPLSUsesStatements(content, mplsGrouping string) (string, error) {
	// Find all occurrences of the MPLS uses statement
	usesPattern := `uses config-ip-sla-mpls-lsp-submode-grouping {`

	// We need to replace each occurrence with a unique inline expansion
	counter := 0

	for {
		// Find the next occurrence
		startIdx := strings.Index(content, usesPattern)
		if startIdx == -1 {
			break // No more occurrences
		}

		// Find the end of this uses block
		braceCount := 0
		endIdx := startIdx
		inUses := false

		for i, char := range content[startIdx:] {
			if char == '{' {
				braceCount++
				inUses = true
			} else if char == '}' {
				braceCount--
				if inUses && braceCount == 0 {
					endIdx = startIdx + i + 1
					break
				}
			}
		}

		if braceCount != 0 {
			return "", fmt.Errorf("unmatched braces in MPLS uses statement")
		}

		usesBlock := content[startIdx:endIdx]

		// Extract the history augment content if it exists
		historyAugment := ""
		if strings.Contains(usesBlock, `augment "history"`) {
			var err error
			historyAugment, err = extractAugmentContent(usesBlock, "history")
			if err != nil {
				return "", fmt.Errorf("failed to extract history augment: %v", err)
			}
		}

		// Create unique inline expansion for this occurrence
		replacement, err := createMPLSInlineExpansion(mplsGrouping, counter, historyAugment)
		if err != nil {
			return "", fmt.Errorf("failed to create MPLS inline expansion: %v", err)
		}

		// Replace this occurrence
		content = content[:startIdx] + replacement + content[endIdx:]
		counter++

		// Safety check to prevent infinite loops
		if counter > 10 {
			return "", fmt.Errorf("too many MPLS uses statements found (>10), possible infinite loop")
		}
	}

	return content, nil
}

// createMPLSInlineExpansion creates inline expansion for MPLS grouping with unique naming
func createMPLSInlineExpansion(mplsGrouping string, counter int, historyAugment string) (string, error) {
	// Extract the base MPLS grouping content
	baseContent, err := extractGroupingContent(mplsGrouping)
	if err != nil {
		return "", fmt.Errorf("failed to extract MPLS grouping content: %v", err)
	}

	// Create the inline expansion with unique container names
	expansion := fmt.Sprintf("// Inline expansion of config-ip-sla-mpls-lsp-submode-grouping (occurrence %d)\n", counter)

	// Process the base content to add unique names to containers that would conflict
	processedContent := replaceMPLSContainerNames(baseContent, counter)
	expansion += processedContent

	// If there's a history augment, add it with unique naming
	if historyAugment != "" {
		expansion += fmt.Sprintf("\n      container history-mpls-%d {\n", counter)
		expansion += "        description\n"
		expansion += fmt.Sprintf("          \"History configuration for MPLS case %d\";\n", counter)
		expansion += indentContent(historyAugment, "        ")
		expansion += "\n      }"
	}

	return expansion, nil
}

// replaceMPLSContainerNames replaces container names in MPLS content to make them unique
func replaceMPLSContainerNames(content string, counter int) string {
	// Add unique suffix to containers that might conflict
	content = strings.ReplaceAll(content, "container history", fmt.Sprintf("container history-mpls-%d", counter))

	// Add other potential conflict resolutions as needed
	return content
}

// expandNestedUsesStatements recursively expands any nested uses statements to eliminate conflicts
func expandNestedUsesStatements(fullContent, augmentContent, uniqueSuffix string) (string, error) {
	// Find all uses statements within the augment content
	usesPattern := regexp.MustCompile(`uses\s+([a-zA-Z0-9\-_]+)\s*\{`)
	matches := usesPattern.FindAllStringSubmatch(augmentContent, -1)

	if len(matches) == 0 {
		// No nested uses statements, return as-is
		return augmentContent, nil
	}

	expandedContent := augmentContent

	for _, match := range matches {
		groupingName := match[1]

		// Extract the nested grouping definition
		nestedGrouping, err := extractGroupingDefinition(fullContent, groupingName)
		if err != nil {
			// If we can't find the grouping, skip expansion (it might be from another module)
			fmt.Printf("Warning: Could not find nested grouping %s, skipping expansion\n", groupingName)
			continue
		}

		// Find the specific uses block in the augment content
		usesBlockPattern := fmt.Sprintf(`uses\s+%s\s*\{[^}]*\}`, regexp.QuoteMeta(groupingName))
		usesBlockRegex := regexp.MustCompile(usesBlockPattern)

		// For each occurrence of this uses statement
		usesBlockMatches := usesBlockRegex.FindAllString(expandedContent, -1)

		for i, usesBlock := range usesBlockMatches {
			// Create a unique expansion for this occurrence
			inlineExpansion, err := createNestedInlineExpansion(nestedGrouping, groupingName, usesBlock, fmt.Sprintf("%s-%d", uniqueSuffix, i))
			if err != nil {
				return "", fmt.Errorf("failed to create nested inline expansion for %s: %v", groupingName, err)
			}

			// Replace the uses block with the inline expansion
			expandedContent = strings.Replace(expandedContent, usesBlock, inlineExpansion, 1)
		}
	}

	return expandedContent, nil
}

// createNestedInlineExpansion creates an inline expansion for a nested uses statement
func createNestedInlineExpansion(groupingDef, groupingName, usesBlock, uniqueId string) (string, error) {
	// Extract the base grouping content
	baseContent, err := extractGroupingContent(groupingDef)
	if err != nil {
		return "", fmt.Errorf("failed to extract nested grouping content: %v", err)
	}

	// Create comment header
	expansion := fmt.Sprintf("// Inline expansion of %s (%s)\n", groupingName, uniqueId)

	// Check if the uses block has augments and handle them
	if strings.Contains(usesBlock, "augment") {
		// Extract augment contents and merge with base content
		expansion += processNestedAugments(baseContent, usesBlock, uniqueId)
	} else {
		// Simple expansion without augments
		expansion += baseContent
	}

	return expansion, nil
}

// processNestedAugments processes augments within nested uses statements
func processNestedAugments(baseContent, usesBlock, uniqueId string) string {
	// For now, we'll do a simplified approach:
	// Replace any container names in the base content to make them unique
	processedContent := baseContent

	// Add unique suffixes to container names to prevent conflicts
	containerPattern := regexp.MustCompile(`container\s+([a-zA-Z0-9\-_]+)`)
	processedContent = containerPattern.ReplaceAllStringFunc(processedContent, func(match string) string {
		parts := strings.Fields(match)
		if len(parts) >= 2 {
			return fmt.Sprintf("container %s-%s", parts[1], uniqueId)
		}
		return match
	})

	return processedContent
}
