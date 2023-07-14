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
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type BGPIPv6UnicastNeighbor struct {
	Device               types.String                      `tfsdk:"device"`
	Id                   types.String                      `tfsdk:"id"`
	DeleteMode           types.String                      `tfsdk:"delete_mode"`
	Asn                  types.String                      `tfsdk:"asn"`
	Ip                   types.String                      `tfsdk:"ip"`
	Activate             types.Bool                        `tfsdk:"activate"`
	SendCommunity        types.String                      `tfsdk:"send_community"`
	RouteReflectorClient types.Bool                        `tfsdk:"route_reflector_client"`
	RouteMaps            []BGPIPv6UnicastNeighborRouteMaps `tfsdk:"route_maps"`
}

type BGPIPv6UnicastNeighborData struct {
	Device               types.String                      `tfsdk:"device"`
	Id                   types.String                      `tfsdk:"id"`
	Asn                  types.String                      `tfsdk:"asn"`
	Ip                   types.String                      `tfsdk:"ip"`
	Activate             types.Bool                        `tfsdk:"activate"`
	SendCommunity        types.String                      `tfsdk:"send_community"`
	RouteReflectorClient types.Bool                        `tfsdk:"route_reflector_client"`
	RouteMaps            []BGPIPv6UnicastNeighborRouteMaps `tfsdk:"route_maps"`
}
type BGPIPv6UnicastNeighborRouteMaps struct {
	InOut        types.String `tfsdk:"in_out"`
	RouteMapName types.String `tfsdk:"route_map_name"`
}

func (data BGPIPv6UnicastNeighbor) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v/address-family/no-vrf/ipv6=unicast/ipv6-unicast/neighbor=%s", url.QueryEscape(fmt.Sprintf("%v", data.Asn.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Ip.ValueString())))
}

func (data BGPIPv6UnicastNeighborData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v/address-family/no-vrf/ipv6=unicast/ipv6-unicast/neighbor=%s", url.QueryEscape(fmt.Sprintf("%v", data.Asn.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Ip.ValueString())))
}

// if last path element has a key -> remove it
func (data BGPIPv6UnicastNeighbor) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data BGPIPv6UnicastNeighbor) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Ip.IsNull() && !data.Ip.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"id", data.Ip.ValueString())
	}
	if !data.Activate.IsNull() && !data.Activate.IsUnknown() {
		if data.Activate.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"activate", map[string]string{})
		}
	}
	if !data.SendCommunity.IsNull() && !data.SendCommunity.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"send-community.send-community-where", data.SendCommunity.ValueString())
	}
	if !data.RouteReflectorClient.IsNull() && !data.RouteReflectorClient.IsUnknown() {
		if data.RouteReflectorClient.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"route-reflector-client", map[string]string{})
		}
	}
	if len(data.RouteMaps) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"route-map", []interface{}{})
		for index, item := range data.RouteMaps {
			if !item.InOut.IsNull() && !item.InOut.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"route-map"+"."+strconv.Itoa(index)+"."+"inout", item.InOut.ValueString())
			}
			if !item.RouteMapName.IsNull() && !item.RouteMapName.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"route-map"+"."+strconv.Itoa(index)+"."+"route-map-name", item.RouteMapName.ValueString())
			}
		}
	}
	return body
}

func (data *BGPIPv6UnicastNeighbor) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "id"); value.Exists() && !data.Ip.IsNull() {
		data.Ip = types.StringValue(value.String())
	} else {
		data.Ip = types.StringNull()
	}
	if value := res.Get(prefix + "activate"); !data.Activate.IsNull() {
		if value.Exists() {
			data.Activate = types.BoolValue(true)
		} else {
			data.Activate = types.BoolValue(false)
		}
	} else {
		data.Activate = types.BoolNull()
	}
	if value := res.Get(prefix + "send-community.send-community-where"); value.Exists() && !data.SendCommunity.IsNull() {
		data.SendCommunity = types.StringValue(value.String())
	} else {
		data.SendCommunity = types.StringNull()
	}
	if value := res.Get(prefix + "route-reflector-client"); !data.RouteReflectorClient.IsNull() {
		if value.Exists() {
			data.RouteReflectorClient = types.BoolValue(true)
		} else {
			data.RouteReflectorClient = types.BoolValue(false)
		}
	} else {
		data.RouteReflectorClient = types.BoolNull()
	}
	for i := range data.RouteMaps {
		keys := [...]string{"inout"}
		keyValues := [...]string{data.RouteMaps[i].InOut.ValueString()}

		var r gjson.Result
		res.Get(prefix + "route-map").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("inout"); value.Exists() && !data.RouteMaps[i].InOut.IsNull() {
			data.RouteMaps[i].InOut = types.StringValue(value.String())
		} else {
			data.RouteMaps[i].InOut = types.StringNull()
		}
		if value := r.Get("route-map-name"); value.Exists() && !data.RouteMaps[i].RouteMapName.IsNull() {
			data.RouteMaps[i].RouteMapName = types.StringValue(value.String())
		} else {
			data.RouteMaps[i].RouteMapName = types.StringNull()
		}
	}
}

func (data *BGPIPv6UnicastNeighborData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "activate"); value.Exists() {
		data.Activate = types.BoolValue(true)
	} else {
		data.Activate = types.BoolValue(false)
	}
	if value := res.Get(prefix + "send-community.send-community-where"); value.Exists() {
		data.SendCommunity = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "route-reflector-client"); value.Exists() {
		data.RouteReflectorClient = types.BoolValue(true)
	} else {
		data.RouteReflectorClient = types.BoolValue(false)
	}
	if value := res.Get(prefix + "route-map"); value.Exists() {
		data.RouteMaps = make([]BGPIPv6UnicastNeighborRouteMaps, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := BGPIPv6UnicastNeighborRouteMaps{}
			if cValue := v.Get("inout"); cValue.Exists() {
				item.InOut = types.StringValue(cValue.String())
			}
			if cValue := v.Get("route-map-name"); cValue.Exists() {
				item.RouteMapName = types.StringValue(cValue.String())
			}
			data.RouteMaps = append(data.RouteMaps, item)
			return true
		})
	}
}

func (data *BGPIPv6UnicastNeighbor) getDeletedListItems(ctx context.Context, state BGPIPv6UnicastNeighbor) []string {
	deletedListItems := make([]string, 0)
	for i := range state.RouteMaps {
		stateKeyValues := [...]string{state.RouteMaps[i].InOut.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.RouteMaps[i].InOut.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.RouteMaps {
			found = true
			if state.RouteMaps[i].InOut.ValueString() != data.RouteMaps[j].InOut.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/route-map=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedListItems
}

func (data *BGPIPv6UnicastNeighbor) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Activate.IsNull() && !data.Activate.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/activate", data.getPath()))
	}
	if !data.RouteReflectorClient.IsNull() && !data.RouteReflectorClient.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/route-reflector-client", data.getPath()))
	}

	return emptyLeafsDelete
}

func (data *BGPIPv6UnicastNeighbor) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.SendCommunity.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/send-community/send-community-where", data.getPath()))
	}
	if !data.RouteReflectorClient.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/route-reflector-client", data.getPath()))
	}
	for i := range data.RouteMaps {
		keyValues := [...]string{data.RouteMaps[i].InOut.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/route-map=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	return deletePaths
}
