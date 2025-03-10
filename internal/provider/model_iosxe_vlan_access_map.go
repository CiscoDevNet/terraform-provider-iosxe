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
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type VLANAccessMap struct {
	Device           types.String `tfsdk:"device"`
	Id               types.String `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	Sequence         types.Int64  `tfsdk:"sequence"`
	MatchIpv6Address types.List   `tfsdk:"match_ipv6_address"`
	MatchIpAddress   types.List   `tfsdk:"match_ip_address"`
	Action           types.String `tfsdk:"action"`
}

type VLANAccessMapData struct {
	Device           types.String `tfsdk:"device"`
	Id               types.String `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	Sequence         types.Int64  `tfsdk:"sequence"`
	MatchIpv6Address types.List   `tfsdk:"match_ipv6_address"`
	MatchIpAddress   types.List   `tfsdk:"match_ip_address"`
	Action           types.String `tfsdk:"action"`
}

func (data VLANAccessMap) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/vlan/Cisco-IOS-XE-vlan:access-map=%v,%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Sequence.ValueInt64())))
}

func (data VLANAccessMapData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/vlan/Cisco-IOS-XE-vlan:access-map=%v,%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Sequence.ValueInt64())))
}

// if last path element has a key -> remove it
func (data VLANAccessMap) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data VLANAccessMap) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", data.Name.ValueString())
	}
	if !data.Sequence.IsNull() && !data.Sequence.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"value", strconv.FormatInt(data.Sequence.ValueInt64(), 10))
	}
	if !data.MatchIpv6Address.IsNull() && !data.MatchIpv6Address.IsUnknown() {
		var values []string
		data.MatchIpv6Address.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match.ipv6.address", values)
	}
	if !data.MatchIpAddress.IsNull() && !data.MatchIpAddress.IsUnknown() {
		var values []string
		data.MatchIpAddress.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"match.ip.address", values)
	}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"action", data.Action.ValueString())
	}
	return body
}

func (data *VLANAccessMap) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(prefix + "value"); value.Exists() && !data.Sequence.IsNull() {
		data.Sequence = types.Int64Value(value.Int())
	} else {
		data.Sequence = types.Int64Null()
	}
	if value := res.Get(prefix + "match.ipv6.address"); value.Exists() && !data.MatchIpv6Address.IsNull() {
		data.MatchIpv6Address = helpers.GetStringList(value.Array())
	} else {
		data.MatchIpv6Address = types.ListNull(types.StringType)
	}
	if value := res.Get(prefix + "match.ip.address"); value.Exists() && !data.MatchIpAddress.IsNull() {
		data.MatchIpAddress = helpers.GetStringList(value.Array())
	} else {
		data.MatchIpAddress = types.ListNull(types.StringType)
	}
	if value := res.Get(prefix + "action"); value.Exists() && !data.Action.IsNull() {
		data.Action = types.StringValue(value.String())
	} else {
		data.Action = types.StringNull()
	}
}

func (data *VLANAccessMap) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "match.ipv6.address"); value.Exists() {
		data.MatchIpv6Address = helpers.GetStringList(value.Array())
	} else {
		data.MatchIpv6Address = types.ListNull(types.StringType)
	}
	if value := res.Get(prefix + "match.ip.address"); value.Exists() {
		data.MatchIpAddress = helpers.GetStringList(value.Array())
	} else {
		data.MatchIpAddress = types.ListNull(types.StringType)
	}
	if value := res.Get(prefix + "action"); value.Exists() {
		data.Action = types.StringValue(value.String())
	}
}

func (data *VLANAccessMapData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "match.ipv6.address"); value.Exists() {
		data.MatchIpv6Address = helpers.GetStringList(value.Array())
	} else {
		data.MatchIpv6Address = types.ListNull(types.StringType)
	}
	if value := res.Get(prefix + "match.ip.address"); value.Exists() {
		data.MatchIpAddress = helpers.GetStringList(value.Array())
	} else {
		data.MatchIpAddress = types.ListNull(types.StringType)
	}
	if value := res.Get(prefix + "action"); value.Exists() {
		data.Action = types.StringValue(value.String())
	}
}

func (data *VLANAccessMap) getDeletedItems(ctx context.Context, state VLANAccessMap) []string {
	deletedItems := make([]string, 0)
	if !state.MatchIpv6Address.IsNull() {
		if data.MatchIpv6Address.IsNull() {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/match/ipv6/address", state.getPath()))
		} else {
			var dataValues, stateValues []string
			data.MatchIpv6Address.ElementsAs(ctx, &dataValues, false)
			state.MatchIpv6Address.ElementsAs(ctx, &stateValues, false)
			for _, v := range stateValues {
				found := false
				for _, vv := range dataValues {
					if v == vv {
						found = true
						break
					}
				}
				if !found {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/match/ipv6/address=%v", state.getPath(), v))
				}
			}
		}
	}
	if !state.MatchIpAddress.IsNull() {
		if data.MatchIpAddress.IsNull() {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/match/ip/address", state.getPath()))
		} else {
			var dataValues, stateValues []string
			data.MatchIpAddress.ElementsAs(ctx, &dataValues, false)
			state.MatchIpAddress.ElementsAs(ctx, &stateValues, false)
			for _, v := range stateValues {
				found := false
				for _, vv := range dataValues {
					if v == vv {
						found = true
						break
					}
				}
				if !found {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/match/ip/address=%v", state.getPath(), v))
				}
			}
		}
	}
	if !state.Action.IsNull() && data.Action.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/action", state.getPath()))
	}
	return deletedItems
}

func (data *VLANAccessMap) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}

func (data *VLANAccessMap) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.MatchIpv6Address.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/match/ipv6/address", data.getPath()))
	}
	if !data.MatchIpAddress.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/match/ip/address", data.getPath()))
	}
	if !data.Action.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/action", data.getPath()))
	}
	return deletePaths
}
