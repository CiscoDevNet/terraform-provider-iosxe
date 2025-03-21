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

type BGPAddressFamilyIPv6 struct {
	Device                           types.String                              `tfsdk:"device"`
	Id                               types.String                              `tfsdk:"id"`
	DeleteMode                       types.String                              `tfsdk:"delete_mode"`
	Asn                              types.String                              `tfsdk:"asn"`
	AfName                           types.String                              `tfsdk:"af_name"`
	Ipv6UnicastRedistributeConnected types.Bool                                `tfsdk:"ipv6_unicast_redistribute_connected"`
	Ipv6UnicastRedistributeStatic    types.Bool                                `tfsdk:"ipv6_unicast_redistribute_static"`
	Ipv6UnicastNetworks              []BGPAddressFamilyIPv6Ipv6UnicastNetworks `tfsdk:"ipv6_unicast_networks"`
}

type BGPAddressFamilyIPv6Data struct {
	Device                           types.String                              `tfsdk:"device"`
	Id                               types.String                              `tfsdk:"id"`
	Asn                              types.String                              `tfsdk:"asn"`
	AfName                           types.String                              `tfsdk:"af_name"`
	Ipv6UnicastRedistributeConnected types.Bool                                `tfsdk:"ipv6_unicast_redistribute_connected"`
	Ipv6UnicastRedistributeStatic    types.Bool                                `tfsdk:"ipv6_unicast_redistribute_static"`
	Ipv6UnicastNetworks              []BGPAddressFamilyIPv6Ipv6UnicastNetworks `tfsdk:"ipv6_unicast_networks"`
}
type BGPAddressFamilyIPv6Ipv6UnicastNetworks struct {
	Network  types.String `tfsdk:"network"`
	RouteMap types.String `tfsdk:"route_map"`
	Backdoor types.Bool   `tfsdk:"backdoor"`
}

func (data BGPAddressFamilyIPv6) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v/address-family/no-vrf/ipv6=%s", url.QueryEscape(fmt.Sprintf("%v", data.Asn.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.AfName.ValueString())))
}

func (data BGPAddressFamilyIPv6Data) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v/address-family/no-vrf/ipv6=%s", url.QueryEscape(fmt.Sprintf("%v", data.Asn.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.AfName.ValueString())))
}

// if last path element has a key -> remove it
func (data BGPAddressFamilyIPv6) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data BGPAddressFamilyIPv6) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.AfName.IsNull() && !data.AfName.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"af-name", data.AfName.ValueString())
	}
	if !data.Ipv6UnicastRedistributeConnected.IsNull() && !data.Ipv6UnicastRedistributeConnected.IsUnknown() {
		if data.Ipv6UnicastRedistributeConnected.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv6-unicast.redistribute-v6.connected", map[string]string{})
		}
	}
	if !data.Ipv6UnicastRedistributeStatic.IsNull() && !data.Ipv6UnicastRedistributeStatic.IsUnknown() {
		if data.Ipv6UnicastRedistributeStatic.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv6-unicast.redistribute-v6.static", map[string]string{})
		}
	}
	if len(data.Ipv6UnicastNetworks) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv6-unicast.network", []interface{}{})
		for index, item := range data.Ipv6UnicastNetworks {
			if !item.Network.IsNull() && !item.Network.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv6-unicast.network"+"."+strconv.Itoa(index)+"."+"number", item.Network.ValueString())
			}
			if !item.RouteMap.IsNull() && !item.RouteMap.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv6-unicast.network"+"."+strconv.Itoa(index)+"."+"route-map", item.RouteMap.ValueString())
			}
			if !item.Backdoor.IsNull() && !item.Backdoor.IsUnknown() {
				if item.Backdoor.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv6-unicast.network"+"."+strconv.Itoa(index)+"."+"backdoor", map[string]string{})
				}
			}
		}
	}
	return body
}

func (data *BGPAddressFamilyIPv6) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "af-name"); value.Exists() && !data.AfName.IsNull() {
		data.AfName = types.StringValue(value.String())
	} else {
		data.AfName = types.StringNull()
	}
	if value := res.Get(prefix + "ipv6-unicast.redistribute-v6.connected"); !data.Ipv6UnicastRedistributeConnected.IsNull() {
		if value.Exists() {
			data.Ipv6UnicastRedistributeConnected = types.BoolValue(true)
		} else {
			data.Ipv6UnicastRedistributeConnected = types.BoolValue(false)
		}
	} else {
		data.Ipv6UnicastRedistributeConnected = types.BoolNull()
	}
	if value := res.Get(prefix + "ipv6-unicast.redistribute-v6.static"); !data.Ipv6UnicastRedistributeStatic.IsNull() {
		if value.Exists() {
			data.Ipv6UnicastRedistributeStatic = types.BoolValue(true)
		} else {
			data.Ipv6UnicastRedistributeStatic = types.BoolValue(false)
		}
	} else {
		data.Ipv6UnicastRedistributeStatic = types.BoolNull()
	}
	for i := range data.Ipv6UnicastNetworks {
		keys := [...]string{"number"}
		keyValues := [...]string{data.Ipv6UnicastNetworks[i].Network.ValueString()}

		var r gjson.Result
		res.Get(prefix + "ipv6-unicast.network").ForEach(
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
		if value := r.Get("number"); value.Exists() && !data.Ipv6UnicastNetworks[i].Network.IsNull() {
			data.Ipv6UnicastNetworks[i].Network = types.StringValue(value.String())
		} else {
			data.Ipv6UnicastNetworks[i].Network = types.StringNull()
		}
		if value := r.Get("route-map"); value.Exists() && !data.Ipv6UnicastNetworks[i].RouteMap.IsNull() {
			data.Ipv6UnicastNetworks[i].RouteMap = types.StringValue(value.String())
		} else {
			data.Ipv6UnicastNetworks[i].RouteMap = types.StringNull()
		}
		if value := r.Get("backdoor"); !data.Ipv6UnicastNetworks[i].Backdoor.IsNull() {
			if value.Exists() {
				data.Ipv6UnicastNetworks[i].Backdoor = types.BoolValue(true)
			} else {
				data.Ipv6UnicastNetworks[i].Backdoor = types.BoolValue(false)
			}
		} else {
			data.Ipv6UnicastNetworks[i].Backdoor = types.BoolNull()
		}
	}
}

func (data *BGPAddressFamilyIPv6) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "ipv6-unicast.redistribute-v6.connected"); value.Exists() {
		data.Ipv6UnicastRedistributeConnected = types.BoolValue(true)
	} else {
		data.Ipv6UnicastRedistributeConnected = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ipv6-unicast.redistribute-v6.static"); value.Exists() {
		data.Ipv6UnicastRedistributeStatic = types.BoolValue(true)
	} else {
		data.Ipv6UnicastRedistributeStatic = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ipv6-unicast.network"); value.Exists() {
		data.Ipv6UnicastNetworks = make([]BGPAddressFamilyIPv6Ipv6UnicastNetworks, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := BGPAddressFamilyIPv6Ipv6UnicastNetworks{}
			if cValue := v.Get("number"); cValue.Exists() {
				item.Network = types.StringValue(cValue.String())
			}
			if cValue := v.Get("route-map"); cValue.Exists() {
				item.RouteMap = types.StringValue(cValue.String())
			}
			if cValue := v.Get("backdoor"); cValue.Exists() {
				item.Backdoor = types.BoolValue(true)
			} else {
				item.Backdoor = types.BoolValue(false)
			}
			data.Ipv6UnicastNetworks = append(data.Ipv6UnicastNetworks, item)
			return true
		})
	}
}

func (data *BGPAddressFamilyIPv6Data) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "ipv6-unicast.redistribute-v6.connected"); value.Exists() {
		data.Ipv6UnicastRedistributeConnected = types.BoolValue(true)
	} else {
		data.Ipv6UnicastRedistributeConnected = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ipv6-unicast.redistribute-v6.static"); value.Exists() {
		data.Ipv6UnicastRedistributeStatic = types.BoolValue(true)
	} else {
		data.Ipv6UnicastRedistributeStatic = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ipv6-unicast.network"); value.Exists() {
		data.Ipv6UnicastNetworks = make([]BGPAddressFamilyIPv6Ipv6UnicastNetworks, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := BGPAddressFamilyIPv6Ipv6UnicastNetworks{}
			if cValue := v.Get("number"); cValue.Exists() {
				item.Network = types.StringValue(cValue.String())
			}
			if cValue := v.Get("route-map"); cValue.Exists() {
				item.RouteMap = types.StringValue(cValue.String())
			}
			if cValue := v.Get("backdoor"); cValue.Exists() {
				item.Backdoor = types.BoolValue(true)
			} else {
				item.Backdoor = types.BoolValue(false)
			}
			data.Ipv6UnicastNetworks = append(data.Ipv6UnicastNetworks, item)
			return true
		})
	}
}

func (data *BGPAddressFamilyIPv6) getDeletedItems(ctx context.Context, state BGPAddressFamilyIPv6) []string {
	deletedItems := make([]string, 0)
	if !state.Ipv6UnicastRedistributeConnected.IsNull() && data.Ipv6UnicastRedistributeConnected.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/ipv6-unicast/redistribute-v6/connected", state.getPath()))
	}
	if !state.Ipv6UnicastRedistributeStatic.IsNull() && data.Ipv6UnicastRedistributeStatic.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/ipv6-unicast/redistribute-v6/static", state.getPath()))
	}
	for i := range state.Ipv6UnicastNetworks {
		stateKeyValues := [...]string{state.Ipv6UnicastNetworks[i].Network.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Ipv6UnicastNetworks[i].Network.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Ipv6UnicastNetworks {
			found = true
			if state.Ipv6UnicastNetworks[i].Network.ValueString() != data.Ipv6UnicastNetworks[j].Network.ValueString() {
				found = false
			}
			if found {
				if !state.Ipv6UnicastNetworks[i].RouteMap.IsNull() && data.Ipv6UnicastNetworks[j].RouteMap.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/ipv6-unicast/network=%v/route-map", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Ipv6UnicastNetworks[i].Backdoor.IsNull() && data.Ipv6UnicastNetworks[j].Backdoor.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/ipv6-unicast/network=%v/backdoor", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/ipv6-unicast/network=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedItems
}

func (data *BGPAddressFamilyIPv6) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Ipv6UnicastRedistributeConnected.IsNull() && !data.Ipv6UnicastRedistributeConnected.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ipv6-unicast/redistribute-v6/connected", data.getPath()))
	}
	if !data.Ipv6UnicastRedistributeStatic.IsNull() && !data.Ipv6UnicastRedistributeStatic.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ipv6-unicast/redistribute-v6/static", data.getPath()))
	}

	for i := range data.Ipv6UnicastNetworks {
		keyValues := [...]string{data.Ipv6UnicastNetworks[i].Network.ValueString()}
		if !data.Ipv6UnicastNetworks[i].Backdoor.IsNull() && !data.Ipv6UnicastNetworks[i].Backdoor.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ipv6-unicast/network=%v/backdoor", data.getPath(), strings.Join(keyValues[:], ",")))
		}
	}
	return emptyLeafsDelete
}

func (data *BGPAddressFamilyIPv6) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Ipv6UnicastRedistributeConnected.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ipv6-unicast/redistribute-v6/connected", data.getPath()))
	}
	if !data.Ipv6UnicastRedistributeStatic.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ipv6-unicast/redistribute-v6/static", data.getPath()))
	}
	for i := range data.Ipv6UnicastNetworks {
		keyValues := [...]string{data.Ipv6UnicastNetworks[i].Network.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/ipv6-unicast/network=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	return deletePaths
}

func (data *BGPAddressFamilyIPv6) getIdsFromPath() {
	reString := strings.ReplaceAll("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v/address-family/no-vrf/ipv6=%s", "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Id.ValueString())
	data.Asn = types.StringValue(matches[1])
	data.AfName = types.StringValue(matches[2])
}
