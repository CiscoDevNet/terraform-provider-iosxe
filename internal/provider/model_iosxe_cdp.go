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
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type CDP struct {
	Device        types.String  `tfsdk:"device"`
	Id            types.String  `tfsdk:"id"`
	Holdtime      types.Int64   `tfsdk:"holdtime"`
	Timer         types.Int64   `tfsdk:"timer"`
	Run           types.Bool    `tfsdk:"run"`
	FilterTlvList types.String  `tfsdk:"filter_tlv_list"`
	TlvLists      []CDPTlvLists `tfsdk:"tlv_lists"`
}

type CDPData struct {
	Device        types.String  `tfsdk:"device"`
	Id            types.String  `tfsdk:"id"`
	Holdtime      types.Int64   `tfsdk:"holdtime"`
	Timer         types.Int64   `tfsdk:"timer"`
	Run           types.Bool    `tfsdk:"run"`
	FilterTlvList types.String  `tfsdk:"filter_tlv_list"`
	TlvLists      []CDPTlvLists `tfsdk:"tlv_lists"`
}
type CDPTlvLists struct {
	Name          types.String `tfsdk:"name"`
	VtpMgmtDomain types.Bool   `tfsdk:"vtp_mgmt_domain"`
	Cos           types.Bool   `tfsdk:"cos"`
	Duplex        types.Bool   `tfsdk:"duplex"`
	Trust         types.Bool   `tfsdk:"trust"`
	Version       types.Bool   `tfsdk:"version"`
}

func (data CDP) getPath() string {
	return "Cisco-IOS-XE-native:native/cdp"
}

func (data CDPData) getPath() string {
	return "Cisco-IOS-XE-native:native/cdp"
}

// if last path element has a key -> remove it
func (data CDP) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data CDP) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Holdtime.IsNull() && !data.Holdtime.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-cdp:holdtime", strconv.FormatInt(data.Holdtime.ValueInt64(), 10))
	}
	if !data.Timer.IsNull() && !data.Timer.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-cdp:timer", strconv.FormatInt(data.Timer.ValueInt64(), 10))
	}
	if !data.Run.IsNull() && !data.Run.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-cdp:run-enable", data.Run.ValueBool())
	}
	if !data.FilterTlvList.IsNull() && !data.FilterTlvList.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-cdp:filter-tlv-list", data.FilterTlvList.ValueString())
	}
	if len(data.TlvLists) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-cdp:tlv-list", []interface{}{})
		for index, item := range data.TlvLists {
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-cdp:tlv-list"+"."+strconv.Itoa(index)+"."+"name", item.Name.ValueString())
			}
			if !item.VtpMgmtDomain.IsNull() && !item.VtpMgmtDomain.IsUnknown() {
				if item.VtpMgmtDomain.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-cdp:tlv-list"+"."+strconv.Itoa(index)+"."+"vtp-mgmt-domain", map[string]string{})
				}
			}
			if !item.Cos.IsNull() && !item.Cos.IsUnknown() {
				if item.Cos.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-cdp:tlv-list"+"."+strconv.Itoa(index)+"."+"cos", map[string]string{})
				}
			}
			if !item.Duplex.IsNull() && !item.Duplex.IsUnknown() {
				if item.Duplex.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-cdp:tlv-list"+"."+strconv.Itoa(index)+"."+"duplex", map[string]string{})
				}
			}
			if !item.Trust.IsNull() && !item.Trust.IsUnknown() {
				if item.Trust.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-cdp:tlv-list"+"."+strconv.Itoa(index)+"."+"trust", map[string]string{})
				}
			}
			if !item.Version.IsNull() && !item.Version.IsUnknown() {
				if item.Version.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-cdp:tlv-list"+"."+strconv.Itoa(index)+"."+"version", map[string]string{})
				}
			}
		}
	}
	return body
}

func (data *CDP) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-cdp:holdtime"); value.Exists() && !data.Holdtime.IsNull() {
		data.Holdtime = types.Int64Value(value.Int())
	} else {
		data.Holdtime = types.Int64Null()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-cdp:timer"); value.Exists() && !data.Timer.IsNull() {
		data.Timer = types.Int64Value(value.Int())
	} else {
		data.Timer = types.Int64Null()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-cdp:run-enable"); !data.Run.IsNull() {
		if value.Exists() {
			data.Run = types.BoolValue(value.Bool())
		}
	} else {
		data.Run = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-cdp:filter-tlv-list"); value.Exists() && !data.FilterTlvList.IsNull() {
		data.FilterTlvList = types.StringValue(value.String())
	} else {
		data.FilterTlvList = types.StringNull()
	}
	for i := range data.TlvLists {
		keys := [...]string{"name"}
		keyValues := [...]string{data.TlvLists[i].Name.ValueString()}

		var r gjson.Result
		res.Get(prefix + "Cisco-IOS-XE-cdp:tlv-list").ForEach(
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
		if value := r.Get("name"); value.Exists() && !data.TlvLists[i].Name.IsNull() {
			data.TlvLists[i].Name = types.StringValue(value.String())
		} else {
			data.TlvLists[i].Name = types.StringNull()
		}
		if value := r.Get("vtp-mgmt-domain"); !data.TlvLists[i].VtpMgmtDomain.IsNull() {
			if value.Exists() {
				data.TlvLists[i].VtpMgmtDomain = types.BoolValue(true)
			} else {
				data.TlvLists[i].VtpMgmtDomain = types.BoolValue(false)
			}
		} else {
			data.TlvLists[i].VtpMgmtDomain = types.BoolNull()
		}
		if value := r.Get("cos"); !data.TlvLists[i].Cos.IsNull() {
			if value.Exists() {
				data.TlvLists[i].Cos = types.BoolValue(true)
			} else {
				data.TlvLists[i].Cos = types.BoolValue(false)
			}
		} else {
			data.TlvLists[i].Cos = types.BoolNull()
		}
		if value := r.Get("duplex"); !data.TlvLists[i].Duplex.IsNull() {
			if value.Exists() {
				data.TlvLists[i].Duplex = types.BoolValue(true)
			} else {
				data.TlvLists[i].Duplex = types.BoolValue(false)
			}
		} else {
			data.TlvLists[i].Duplex = types.BoolNull()
		}
		if value := r.Get("trust"); !data.TlvLists[i].Trust.IsNull() {
			if value.Exists() {
				data.TlvLists[i].Trust = types.BoolValue(true)
			} else {
				data.TlvLists[i].Trust = types.BoolValue(false)
			}
		} else {
			data.TlvLists[i].Trust = types.BoolNull()
		}
		if value := r.Get("version"); !data.TlvLists[i].Version.IsNull() {
			if value.Exists() {
				data.TlvLists[i].Version = types.BoolValue(true)
			} else {
				data.TlvLists[i].Version = types.BoolValue(false)
			}
		} else {
			data.TlvLists[i].Version = types.BoolNull()
		}
	}
}

func (data *CDP) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-cdp:holdtime"); value.Exists() {
		data.Holdtime = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-cdp:timer"); value.Exists() {
		data.Timer = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-cdp:run-enable"); value.Exists() {
		data.Run = types.BoolValue(value.Bool())
	} else {
		data.Run = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-cdp:filter-tlv-list"); value.Exists() {
		data.FilterTlvList = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-cdp:tlv-list"); value.Exists() {
		data.TlvLists = make([]CDPTlvLists, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CDPTlvLists{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			if cValue := v.Get("vtp-mgmt-domain"); cValue.Exists() {
				item.VtpMgmtDomain = types.BoolValue(true)
			} else {
				item.VtpMgmtDomain = types.BoolValue(false)
			}
			if cValue := v.Get("cos"); cValue.Exists() {
				item.Cos = types.BoolValue(true)
			} else {
				item.Cos = types.BoolValue(false)
			}
			if cValue := v.Get("duplex"); cValue.Exists() {
				item.Duplex = types.BoolValue(true)
			} else {
				item.Duplex = types.BoolValue(false)
			}
			if cValue := v.Get("trust"); cValue.Exists() {
				item.Trust = types.BoolValue(true)
			} else {
				item.Trust = types.BoolValue(false)
			}
			if cValue := v.Get("version"); cValue.Exists() {
				item.Version = types.BoolValue(true)
			} else {
				item.Version = types.BoolValue(false)
			}
			data.TlvLists = append(data.TlvLists, item)
			return true
		})
	}
}

func (data *CDPData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-cdp:holdtime"); value.Exists() {
		data.Holdtime = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-cdp:timer"); value.Exists() {
		data.Timer = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-cdp:run-enable"); value.Exists() {
		data.Run = types.BoolValue(value.Bool())
	} else {
		data.Run = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-cdp:filter-tlv-list"); value.Exists() {
		data.FilterTlvList = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-cdp:tlv-list"); value.Exists() {
		data.TlvLists = make([]CDPTlvLists, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CDPTlvLists{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			if cValue := v.Get("vtp-mgmt-domain"); cValue.Exists() {
				item.VtpMgmtDomain = types.BoolValue(true)
			} else {
				item.VtpMgmtDomain = types.BoolValue(false)
			}
			if cValue := v.Get("cos"); cValue.Exists() {
				item.Cos = types.BoolValue(true)
			} else {
				item.Cos = types.BoolValue(false)
			}
			if cValue := v.Get("duplex"); cValue.Exists() {
				item.Duplex = types.BoolValue(true)
			} else {
				item.Duplex = types.BoolValue(false)
			}
			if cValue := v.Get("trust"); cValue.Exists() {
				item.Trust = types.BoolValue(true)
			} else {
				item.Trust = types.BoolValue(false)
			}
			if cValue := v.Get("version"); cValue.Exists() {
				item.Version = types.BoolValue(true)
			} else {
				item.Version = types.BoolValue(false)
			}
			data.TlvLists = append(data.TlvLists, item)
			return true
		})
	}
}

func (data *CDP) getDeletedItems(ctx context.Context, state CDP) []string {
	deletedItems := make([]string, 0)
	if !state.Holdtime.IsNull() && data.Holdtime.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-cdp:holdtime", state.getPath()))
	}
	if !state.Timer.IsNull() && data.Timer.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-cdp:timer", state.getPath()))
	}
	if !state.Run.IsNull() && data.Run.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-cdp:run-enable", state.getPath()))
	}
	if !state.FilterTlvList.IsNull() && data.FilterTlvList.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-cdp:filter-tlv-list", state.getPath()))
	}
	for i := range state.TlvLists {
		stateKeyValues := [...]string{state.TlvLists[i].Name.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.TlvLists[i].Name.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.TlvLists {
			found = true
			if state.TlvLists[i].Name.ValueString() != data.TlvLists[j].Name.ValueString() {
				found = false
			}
			if found {
				if !state.TlvLists[i].VtpMgmtDomain.IsNull() && data.TlvLists[j].VtpMgmtDomain.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-cdp:tlv-list=%v/vtp-mgmt-domain", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.TlvLists[i].Cos.IsNull() && data.TlvLists[j].Cos.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-cdp:tlv-list=%v/cos", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.TlvLists[i].Duplex.IsNull() && data.TlvLists[j].Duplex.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-cdp:tlv-list=%v/duplex", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.TlvLists[i].Trust.IsNull() && data.TlvLists[j].Trust.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-cdp:tlv-list=%v/trust", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.TlvLists[i].Version.IsNull() && data.TlvLists[j].Version.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-cdp:tlv-list=%v/version", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-cdp:tlv-list=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedItems
}

func (data *CDP) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	for i := range data.TlvLists {
		keyValues := [...]string{data.TlvLists[i].Name.ValueString()}
		if !data.TlvLists[i].VtpMgmtDomain.IsNull() && !data.TlvLists[i].VtpMgmtDomain.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-cdp:tlv-list=%v/vtp-mgmt-domain", data.getPath(), strings.Join(keyValues[:], ",")))
		}
		if !data.TlvLists[i].Cos.IsNull() && !data.TlvLists[i].Cos.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-cdp:tlv-list=%v/cos", data.getPath(), strings.Join(keyValues[:], ",")))
		}
		if !data.TlvLists[i].Duplex.IsNull() && !data.TlvLists[i].Duplex.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-cdp:tlv-list=%v/duplex", data.getPath(), strings.Join(keyValues[:], ",")))
		}
		if !data.TlvLists[i].Trust.IsNull() && !data.TlvLists[i].Trust.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-cdp:tlv-list=%v/trust", data.getPath(), strings.Join(keyValues[:], ",")))
		}
		if !data.TlvLists[i].Version.IsNull() && !data.TlvLists[i].Version.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-cdp:tlv-list=%v/version", data.getPath(), strings.Join(keyValues[:], ",")))
		}
	}
	return emptyLeafsDelete
}

func (data *CDP) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Holdtime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-cdp:holdtime", data.getPath()))
	}
	if !data.Timer.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-cdp:timer", data.getPath()))
	}
	if !data.Run.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-cdp:run-enable", data.getPath()))
	}
	if !data.FilterTlvList.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-cdp:filter-tlv-list", data.getPath()))
	}
	for i := range data.TlvLists {
		keyValues := [...]string{data.TlvLists[i].Name.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-cdp:tlv-list=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	return deletePaths
}
