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
	"net/url"
	"regexp"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type CommunityListStandard struct {
	Device        types.String `tfsdk:"device"`
	Id            types.String `tfsdk:"id"`
	Name          types.String `tfsdk:"name"`
	DenyEntries   types.List   `tfsdk:"deny_entries"`
	PermitEntries types.List   `tfsdk:"permit_entries"`
}

type CommunityListStandardData struct {
	Device        types.String `tfsdk:"device"`
	Id            types.String `tfsdk:"id"`
	Name          types.String `tfsdk:"name"`
	DenyEntries   types.List   `tfsdk:"deny_entries"`
	PermitEntries types.List   `tfsdk:"permit_entries"`
}

func (data CommunityListStandard) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/ip/Cisco-IOS-XE-bgp:community-list/standard=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

func (data CommunityListStandardData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/ip/Cisco-IOS-XE-bgp:community-list/standard=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

// if last path element has a key -> remove it
func (data CommunityListStandard) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data CommunityListStandard) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", data.Name.ValueString())
	}
	if !data.DenyEntries.IsNull() && !data.DenyEntries.IsUnknown() {
		var values []string
		data.DenyEntries.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"deny.deny-list", values)
	}
	if !data.PermitEntries.IsNull() && !data.PermitEntries.IsUnknown() {
		var values []string
		data.PermitEntries.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"permit.permit-list", values)
	}
	return body
}

func (data *CommunityListStandard) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(prefix + "deny.deny-list"); value.Exists() && !data.DenyEntries.IsNull() {
		data.DenyEntries = helpers.GetStringList(value.Array())
	} else {
		data.DenyEntries = types.ListNull(types.StringType)
	}
	if value := res.Get(prefix + "permit.permit-list"); value.Exists() && !data.PermitEntries.IsNull() {
		data.PermitEntries = helpers.GetStringList(value.Array())
	} else {
		data.PermitEntries = types.ListNull(types.StringType)
	}
}

func (data *CommunityListStandardData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "deny.deny-list"); value.Exists() {
		data.DenyEntries = helpers.GetStringList(value.Array())
	} else {
		data.DenyEntries = types.ListNull(types.StringType)
	}
	if value := res.Get(prefix + "permit.permit-list"); value.Exists() {
		data.PermitEntries = helpers.GetStringList(value.Array())
	} else {
		data.PermitEntries = types.ListNull(types.StringType)
	}
}

func (data *CommunityListStandard) getDeletedItems(ctx context.Context, state CommunityListStandard) []string {
	deletedItems := make([]string, 0)
	if !state.DenyEntries.IsNull() {
		if data.DenyEntries.IsNull() {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/deny/deny-list", state.getPath()))
		} else {
			var dataValues, stateValues []string
			data.DenyEntries.ElementsAs(ctx, &dataValues, false)
			state.DenyEntries.ElementsAs(ctx, &stateValues, false)
			for _, v := range stateValues {
				found := false
				for _, vv := range dataValues {
					if v == vv {
						found = true
						break
					}
				}
				if !found {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/deny/deny-list=%v", state.getPath(), v))
				}
			}
		}
	}
	if !state.PermitEntries.IsNull() {
		if data.PermitEntries.IsNull() {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/permit/permit-list", state.getPath()))
		} else {
			var dataValues, stateValues []string
			data.PermitEntries.ElementsAs(ctx, &dataValues, false)
			state.PermitEntries.ElementsAs(ctx, &stateValues, false)
			for _, v := range stateValues {
				found := false
				for _, vv := range dataValues {
					if v == vv {
						found = true
						break
					}
				}
				if !found {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/permit/permit-list=%v", state.getPath(), v))
				}
			}
		}
	}
	return deletedItems
}

func (data *CommunityListStandard) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}

func (data *CommunityListStandard) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.DenyEntries.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/deny/deny-list", data.getPath()))
	}
	if !data.PermitEntries.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/permit/permit-list", data.getPath()))
	}
	return deletePaths
}