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
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/xmldot"
)

// Contains checks if a string exists in a slice of strings.
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
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

// NormalizeCommunityValue converts a decimal BGP community integer to AA:NN notation.
// Standard communities are 32-bit: upper 16 bits = AS number, lower 16 bits = value.
// Values already in colon notation or well-known names pass through unchanged.
func NormalizeCommunityValue(val string) string {
	if strings.Contains(val, ":") {
		return val
	}
	n, err := strconv.ParseUint(val, 10, 32)
	if err != nil {
		return val
	}
	return fmt.Sprintf("%d:%d", n/65536, n%65536)
}

// GetNormalizedCommunityList converts a slice of gjson.Result to a Terraform types.List,
// normalizing decimal BGP community values to AA:NN notation.
func GetNormalizedCommunityList(result []gjson.Result) types.List {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(NormalizeCommunityValue(result[r].String()))
	}
	return types.ListValueMust(types.StringType, v)
}

// GetNormalizedCommunityListXML converts a slice of xmldot.Result to a Terraform types.List,
// normalizing decimal BGP community values to AA:NN notation.
func GetNormalizedCommunityListXML(result []xmldot.Result) types.List {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(NormalizeCommunityValue(result[r].String()))
	}
	return types.ListValueMust(types.StringType, v)
}
