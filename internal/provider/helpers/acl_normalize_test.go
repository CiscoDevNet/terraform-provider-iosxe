// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package helpers

import "testing"

func TestNormalizePort(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"www", "80"},
		{"bgp", "179"},
		{"https", "443"},
		{"telnet", "23"},
		{"ftp", "21"},
		{"domain", "53"},
		{"smtp", "25"},
		{"ntp", "123"},
		{"snmp", "161"},
		{"sip", "5060"},
		// Numeric passthrough
		{"80", "80"},
		{"179", "179"},
		{"443", "443"},
		{"8080", "8080"},
		// Unknown string passthrough
		{"unknown", "unknown"},
		{"custom-app", "custom-app"},
		// Empty string passthrough
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := NormalizePort(tt.input)
			if result != tt.expected {
				t.Errorf("NormalizePort(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestNormalizeDscp(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"ef", "46"},
		{"af41", "34"},
		{"af11", "10"},
		{"cs1", "8"},
		{"cs6", "48"},
		{"default", "0"},
		// Numeric passthrough
		{"46", "46"},
		{"34", "34"},
		{"0", "0"},
		// Unknown string passthrough
		{"unknown", "unknown"},
		// Empty string passthrough
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := NormalizeDscp(tt.input)
			if result != tt.expected {
				t.Errorf("NormalizeDscp(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
