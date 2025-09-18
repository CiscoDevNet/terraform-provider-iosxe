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
		return patchEEMModel(filepath)
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

// patchEEMModel fixes the ios:logging-level-type reference issue in the EEM model
func patchEEMModel(filepath string) error {
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
