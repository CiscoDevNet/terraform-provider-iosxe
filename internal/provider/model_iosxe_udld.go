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

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type UDLD struct {
	Device           types.String `tfsdk:"device"`
	Id               types.String `tfsdk:"id"`
	DeleteMode       types.String `tfsdk:"delete_mode"`
	Aggressive       types.Bool   `tfsdk:"aggressive"`
	Enable           types.Bool   `tfsdk:"enable"`
	MessageTime      types.Int64  `tfsdk:"message_time"`
	RecoveryInterval types.Int64  `tfsdk:"recovery_interval"`
}

type UDLDData struct {
	Device           types.String `tfsdk:"device"`
	Id               types.String `tfsdk:"id"`
	Aggressive       types.Bool   `tfsdk:"aggressive"`
	Enable           types.Bool   `tfsdk:"enable"`
	MessageTime      types.Int64  `tfsdk:"message_time"`
	RecoveryInterval types.Int64  `tfsdk:"recovery_interval"`
}

func (data UDLD) getPath() string {
	return "Cisco-IOS-XE-native:native/udld"
}

func (data UDLDData) getPath() string {
	return "Cisco-IOS-XE-native:native/udld"
}

// if last path element has a key -> remove it
func (data UDLD) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data UDLD) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Aggressive.IsNull() && !data.Aggressive.IsUnknown() {
		if data.Aggressive.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-udld:aggressive", map[string]string{})
		}
	}
	if !data.Enable.IsNull() && !data.Enable.IsUnknown() {
		if data.Enable.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-udld:enable", map[string]string{})
		}
	}
	if !data.MessageTime.IsNull() && !data.MessageTime.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-udld:message.time", strconv.FormatInt(data.MessageTime.ValueInt64(), 10))
	}
	if !data.RecoveryInterval.IsNull() && !data.RecoveryInterval.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-udld:recovery.interval", strconv.FormatInt(data.RecoveryInterval.ValueInt64(), 10))
	}
	return body
}

func (data *UDLD) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-udld:aggressive"); !data.Aggressive.IsNull() {
		if value.Exists() {
			data.Aggressive = types.BoolValue(true)
		} else {
			data.Aggressive = types.BoolValue(false)
		}
	} else {
		data.Aggressive = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-udld:enable"); !data.Enable.IsNull() {
		if value.Exists() {
			data.Enable = types.BoolValue(true)
		} else {
			data.Enable = types.BoolValue(false)
		}
	} else {
		data.Enable = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-udld:message.time"); value.Exists() && !data.MessageTime.IsNull() {
		data.MessageTime = types.Int64Value(value.Int())
	} else {
		data.MessageTime = types.Int64Null()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-udld:recovery.interval"); value.Exists() && !data.RecoveryInterval.IsNull() {
		data.RecoveryInterval = types.Int64Value(value.Int())
	} else {
		data.RecoveryInterval = types.Int64Null()
	}
}

func (data *UDLD) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-udld:aggressive"); value.Exists() {
		data.Aggressive = types.BoolValue(true)
	} else {
		data.Aggressive = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-udld:enable"); value.Exists() {
		data.Enable = types.BoolValue(true)
	} else {
		data.Enable = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-udld:message.time"); value.Exists() {
		data.MessageTime = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-udld:recovery.interval"); value.Exists() {
		data.RecoveryInterval = types.Int64Value(value.Int())
	}
}

func (data *UDLDData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-udld:aggressive"); value.Exists() {
		data.Aggressive = types.BoolValue(true)
	} else {
		data.Aggressive = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-udld:enable"); value.Exists() {
		data.Enable = types.BoolValue(true)
	} else {
		data.Enable = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-udld:message.time"); value.Exists() {
		data.MessageTime = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-udld:recovery.interval"); value.Exists() {
		data.RecoveryInterval = types.Int64Value(value.Int())
	}
}

func (data *UDLD) getDeletedItems(ctx context.Context, state UDLD) []string {
	deletedItems := make([]string, 0)
	if !state.Aggressive.IsNull() && data.Aggressive.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-udld:aggressive", state.getPath()))
	}
	if !state.Enable.IsNull() && data.Enable.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-udld:enable", state.getPath()))
	}
	if !state.MessageTime.IsNull() && data.MessageTime.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-udld:message/time", state.getPath()))
	}
	if !state.RecoveryInterval.IsNull() && data.RecoveryInterval.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-udld:recovery/interval", state.getPath()))
	}
	return deletedItems
}

func (data *UDLD) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Aggressive.IsNull() && !data.Aggressive.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-udld:aggressive", data.getPath()))
	}
	if !data.Enable.IsNull() && !data.Enable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-udld:enable", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *UDLD) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Aggressive.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-udld:aggressive", data.getPath()))
	}
	if !data.Enable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-udld:enable", data.getPath()))
	}
	if !data.MessageTime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-udld:message/time", data.getPath()))
	}
	if !data.RecoveryInterval.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-udld:recovery/interval", data.getPath()))
	}
	return deletePaths
}
