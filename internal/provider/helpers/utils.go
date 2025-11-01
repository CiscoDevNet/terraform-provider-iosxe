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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/xmldot"
	"github.com/tidwall/gjson"
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
