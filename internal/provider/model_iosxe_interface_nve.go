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

type InterfaceNVE struct {
	Device                      types.String          `tfsdk:"device"`
	Id                          types.String          `tfsdk:"id"`
	DeleteMode                  types.String          `tfsdk:"delete_mode"`
	Name                        types.Int64           `tfsdk:"name"`
	Description                 types.String          `tfsdk:"description"`
	Shutdown                    types.Bool            `tfsdk:"shutdown"`
	HostReachabilityProtocolBgp types.Bool            `tfsdk:"host_reachability_protocol_bgp"`
	SourceInterfaceLoopback     types.Int64           `tfsdk:"source_interface_loopback"`
	VniVrfs                     []InterfaceNVEVniVrfs `tfsdk:"vni_vrfs"`
	Vnis                        []InterfaceNVEVnis    `tfsdk:"vnis"`
}

type InterfaceNVEData struct {
	Device                      types.String          `tfsdk:"device"`
	Id                          types.String          `tfsdk:"id"`
	Name                        types.Int64           `tfsdk:"name"`
	Description                 types.String          `tfsdk:"description"`
	Shutdown                    types.Bool            `tfsdk:"shutdown"`
	HostReachabilityProtocolBgp types.Bool            `tfsdk:"host_reachability_protocol_bgp"`
	SourceInterfaceLoopback     types.Int64           `tfsdk:"source_interface_loopback"`
	VniVrfs                     []InterfaceNVEVniVrfs `tfsdk:"vni_vrfs"`
	Vnis                        []InterfaceNVEVnis    `tfsdk:"vnis"`
}
type InterfaceNVEVniVrfs struct {
	VniRange types.String `tfsdk:"vni_range"`
	Vrf      types.String `tfsdk:"vrf"`
}
type InterfaceNVEVnis struct {
	VniRange           types.String `tfsdk:"vni_range"`
	Ipv4MulticastGroup types.String `tfsdk:"ipv4_multicast_group"`
	IngressReplication types.Bool   `tfsdk:"ingress_replication"`
}

func (data InterfaceNVE) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/interface/nve=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueInt64())))
}

func (data InterfaceNVEData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/interface/nve=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueInt64())))
}

// if last path element has a key -> remove it
func (data InterfaceNVE) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data InterfaceNVE) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", strconv.FormatInt(data.Name.ValueInt64(), 10))
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"description", data.Description.ValueString())
	}
	if !data.Shutdown.IsNull() && !data.Shutdown.IsUnknown() {
		if data.Shutdown.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"shutdown", map[string]string{})
		}
	}
	if !data.HostReachabilityProtocolBgp.IsNull() && !data.HostReachabilityProtocolBgp.IsUnknown() {
		if data.HostReachabilityProtocolBgp.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"host-reachability.protocol.bgp", map[string]string{})
		}
	}
	if !data.SourceInterfaceLoopback.IsNull() && !data.SourceInterfaceLoopback.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"source-interface.Loopback", strconv.FormatInt(data.SourceInterfaceLoopback.ValueInt64(), 10))
	}
	if len(data.VniVrfs) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member-in-one-line.member.vni", []interface{}{})
		for index, item := range data.VniVrfs {
			if !item.VniRange.IsNull() && !item.VniRange.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member-in-one-line.member.vni"+"."+strconv.Itoa(index)+"."+"vni-range", item.VniRange.ValueString())
			}
			if !item.Vrf.IsNull() && !item.Vrf.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member-in-one-line.member.vni"+"."+strconv.Itoa(index)+"."+"vrf", item.Vrf.ValueString())
			}
		}
	}
	if len(data.Vnis) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.vni", []interface{}{})
		for index, item := range data.Vnis {
			if !item.VniRange.IsNull() && !item.VniRange.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.vni"+"."+strconv.Itoa(index)+"."+"vni-range", item.VniRange.ValueString())
			}
			if !item.Ipv4MulticastGroup.IsNull() && !item.Ipv4MulticastGroup.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.vni"+"."+strconv.Itoa(index)+"."+"mcast-group.multicast-group-min", item.Ipv4MulticastGroup.ValueString())
			}
			if !item.IngressReplication.IsNull() && !item.IngressReplication.IsUnknown() {
				if item.IngressReplication.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"member.vni"+"."+strconv.Itoa(index)+"."+"ir-cp-config.ingress-replication", map[string]string{})
				}
			}
		}
	}
	return body
}

func (data *InterfaceNVE) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.Int64Value(value.Int())
	} else {
		data.Name = types.Int64Null()
	}
	if value := res.Get(prefix + "description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get(prefix + "shutdown"); !data.Shutdown.IsNull() {
		if value.Exists() {
			data.Shutdown = types.BoolValue(true)
		} else {
			data.Shutdown = types.BoolValue(false)
		}
	} else {
		data.Shutdown = types.BoolNull()
	}
	if value := res.Get(prefix + "host-reachability.protocol.bgp"); !data.HostReachabilityProtocolBgp.IsNull() {
		if value.Exists() {
			data.HostReachabilityProtocolBgp = types.BoolValue(true)
		} else {
			data.HostReachabilityProtocolBgp = types.BoolValue(false)
		}
	} else {
		data.HostReachabilityProtocolBgp = types.BoolNull()
	}
	if value := res.Get(prefix + "source-interface.Loopback"); value.Exists() && !data.SourceInterfaceLoopback.IsNull() {
		data.SourceInterfaceLoopback = types.Int64Value(value.Int())
	} else {
		data.SourceInterfaceLoopback = types.Int64Null()
	}
	for i := range data.VniVrfs {
		keys := [...]string{"vni-range"}
		keyValues := [...]string{data.VniVrfs[i].VniRange.ValueString()}

		var r gjson.Result
		res.Get(prefix + "member-in-one-line.member.vni").ForEach(
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
		if value := r.Get("vni-range"); value.Exists() && !data.VniVrfs[i].VniRange.IsNull() {
			data.VniVrfs[i].VniRange = types.StringValue(value.String())
		} else {
			data.VniVrfs[i].VniRange = types.StringNull()
		}
		if value := r.Get("vrf"); value.Exists() && !data.VniVrfs[i].Vrf.IsNull() {
			data.VniVrfs[i].Vrf = types.StringValue(value.String())
		} else {
			data.VniVrfs[i].Vrf = types.StringNull()
		}
	}
	for i := range data.Vnis {
		keys := [...]string{"vni-range"}
		keyValues := [...]string{data.Vnis[i].VniRange.ValueString()}

		var r gjson.Result
		res.Get(prefix + "member.vni").ForEach(
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
		if value := r.Get("vni-range"); value.Exists() && !data.Vnis[i].VniRange.IsNull() {
			data.Vnis[i].VniRange = types.StringValue(value.String())
		} else {
			data.Vnis[i].VniRange = types.StringNull()
		}
		if value := r.Get("mcast-group.multicast-group-min"); value.Exists() && !data.Vnis[i].Ipv4MulticastGroup.IsNull() {
			data.Vnis[i].Ipv4MulticastGroup = types.StringValue(value.String())
		} else {
			data.Vnis[i].Ipv4MulticastGroup = types.StringNull()
		}
		if value := r.Get("ir-cp-config.ingress-replication"); !data.Vnis[i].IngressReplication.IsNull() {
			if value.Exists() {
				data.Vnis[i].IngressReplication = types.BoolValue(true)
			} else {
				data.Vnis[i].IngressReplication = types.BoolValue(false)
			}
		} else {
			data.Vnis[i].IngressReplication = types.BoolNull()
		}
	}
}

func (data *InterfaceNVE) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "shutdown"); value.Exists() {
		data.Shutdown = types.BoolValue(true)
	} else {
		data.Shutdown = types.BoolValue(false)
	}
	if value := res.Get(prefix + "host-reachability.protocol.bgp"); value.Exists() {
		data.HostReachabilityProtocolBgp = types.BoolValue(true)
	} else {
		data.HostReachabilityProtocolBgp = types.BoolValue(false)
	}
	if value := res.Get(prefix + "source-interface.Loopback"); value.Exists() {
		data.SourceInterfaceLoopback = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "member-in-one-line.member.vni"); value.Exists() {
		data.VniVrfs = make([]InterfaceNVEVniVrfs, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := InterfaceNVEVniVrfs{}
			if cValue := v.Get("vni-range"); cValue.Exists() {
				item.VniRange = types.StringValue(cValue.String())
			}
			if cValue := v.Get("vrf"); cValue.Exists() {
				item.Vrf = types.StringValue(cValue.String())
			}
			data.VniVrfs = append(data.VniVrfs, item)
			return true
		})
	}
	if value := res.Get(prefix + "member.vni"); value.Exists() {
		data.Vnis = make([]InterfaceNVEVnis, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := InterfaceNVEVnis{}
			if cValue := v.Get("vni-range"); cValue.Exists() {
				item.VniRange = types.StringValue(cValue.String())
			}
			if cValue := v.Get("mcast-group.multicast-group-min"); cValue.Exists() {
				item.Ipv4MulticastGroup = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ir-cp-config.ingress-replication"); cValue.Exists() {
				item.IngressReplication = types.BoolValue(true)
			} else {
				item.IngressReplication = types.BoolValue(false)
			}
			data.Vnis = append(data.Vnis, item)
			return true
		})
	}
}

func (data *InterfaceNVEData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "shutdown"); value.Exists() {
		data.Shutdown = types.BoolValue(true)
	} else {
		data.Shutdown = types.BoolValue(false)
	}
	if value := res.Get(prefix + "host-reachability.protocol.bgp"); value.Exists() {
		data.HostReachabilityProtocolBgp = types.BoolValue(true)
	} else {
		data.HostReachabilityProtocolBgp = types.BoolValue(false)
	}
	if value := res.Get(prefix + "source-interface.Loopback"); value.Exists() {
		data.SourceInterfaceLoopback = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "member-in-one-line.member.vni"); value.Exists() {
		data.VniVrfs = make([]InterfaceNVEVniVrfs, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := InterfaceNVEVniVrfs{}
			if cValue := v.Get("vni-range"); cValue.Exists() {
				item.VniRange = types.StringValue(cValue.String())
			}
			if cValue := v.Get("vrf"); cValue.Exists() {
				item.Vrf = types.StringValue(cValue.String())
			}
			data.VniVrfs = append(data.VniVrfs, item)
			return true
		})
	}
	if value := res.Get(prefix + "member.vni"); value.Exists() {
		data.Vnis = make([]InterfaceNVEVnis, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := InterfaceNVEVnis{}
			if cValue := v.Get("vni-range"); cValue.Exists() {
				item.VniRange = types.StringValue(cValue.String())
			}
			if cValue := v.Get("mcast-group.multicast-group-min"); cValue.Exists() {
				item.Ipv4MulticastGroup = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ir-cp-config.ingress-replication"); cValue.Exists() {
				item.IngressReplication = types.BoolValue(true)
			} else {
				item.IngressReplication = types.BoolValue(false)
			}
			data.Vnis = append(data.Vnis, item)
			return true
		})
	}
}

func (data *InterfaceNVE) getDeletedItems(ctx context.Context, state InterfaceNVE) []string {
	deletedItems := make([]string, 0)
	if !state.Description.IsNull() && data.Description.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/description", state.getPath()))
	}
	if !state.Shutdown.IsNull() && data.Shutdown.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/shutdown", state.getPath()))
	}
	if !state.SourceInterfaceLoopback.IsNull() && data.SourceInterfaceLoopback.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/source-interface/Loopback", state.getPath()))
	}
	for i := range state.VniVrfs {
		stateKeyValues := [...]string{state.VniVrfs[i].VniRange.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.VniVrfs[i].VniRange.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.VniVrfs {
			found = true
			if state.VniVrfs[i].VniRange.ValueString() != data.VniVrfs[j].VniRange.ValueString() {
				found = false
			}
			if found {
				if !state.VniVrfs[i].Vrf.IsNull() && data.VniVrfs[j].Vrf.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/member-in-one-line/member/vni=%v/vrf", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/member-in-one-line/member/vni=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	for i := range state.Vnis {
		stateKeyValues := [...]string{state.Vnis[i].VniRange.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Vnis[i].VniRange.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Vnis {
			found = true
			if state.Vnis[i].VniRange.ValueString() != data.Vnis[j].VniRange.ValueString() {
				found = false
			}
			if found {
				if !state.Vnis[i].Ipv4MulticastGroup.IsNull() && data.Vnis[j].Ipv4MulticastGroup.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/member/vni=%v/mcast-group/multicast-group-min", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Vnis[i].IngressReplication.IsNull() && data.Vnis[j].IngressReplication.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/member/vni=%v/ir-cp-config/ingress-replication", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/member/vni=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedItems
}

func (data *InterfaceNVE) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Shutdown.IsNull() && !data.Shutdown.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/shutdown", data.getPath()))
	}
	if !data.HostReachabilityProtocolBgp.IsNull() && !data.HostReachabilityProtocolBgp.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/host-reachability/protocol/bgp", data.getPath()))
	}

	for i := range data.Vnis {
		keyValues := [...]string{data.Vnis[i].VniRange.ValueString()}
		if !data.Vnis[i].IngressReplication.IsNull() && !data.Vnis[i].IngressReplication.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/member/vni=%v/ir-cp-config/ingress-replication", data.getPath(), strings.Join(keyValues[:], ",")))
		}
	}
	return emptyLeafsDelete
}

func (data *InterfaceNVE) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Description.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/description", data.getPath()))
	}
	if !data.Shutdown.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/shutdown", data.getPath()))
	}
	if !data.SourceInterfaceLoopback.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/source-interface/Loopback", data.getPath()))
	}
	for i := range data.VniVrfs {
		keyValues := [...]string{data.VniVrfs[i].VniRange.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/member-in-one-line/member/vni=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	for i := range data.Vnis {
		keyValues := [...]string{data.Vnis[i].VniRange.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/member/vni=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	return deletePaths
}

func (data *InterfaceNVE) getIdsFromPath() {
	reString := strings.ReplaceAll("Cisco-IOS-XE-native:native/interface/nve=%v", "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Id.ValueString())
	data.Name = types.Int64Value(helpers.Must(strconv.ParseInt(matches[1], 10, 0)))
}
