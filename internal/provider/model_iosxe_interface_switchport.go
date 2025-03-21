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
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type InterfaceSwitchport struct {
	Device                     types.String `tfsdk:"device"`
	Id                         types.String `tfsdk:"id"`
	DeleteMode                 types.String `tfsdk:"delete_mode"`
	Type                       types.String `tfsdk:"type"`
	Name                       types.String `tfsdk:"name"`
	ModeAccess                 types.Bool   `tfsdk:"mode_access"`
	ModeDot1qTunnel            types.Bool   `tfsdk:"mode_dot1q_tunnel"`
	ModePrivateVlanTrunk       types.Bool   `tfsdk:"mode_private_vlan_trunk"`
	ModePrivateVlanHost        types.Bool   `tfsdk:"mode_private_vlan_host"`
	ModePrivateVlanPromiscuous types.Bool   `tfsdk:"mode_private_vlan_promiscuous"`
	ModeTrunk                  types.Bool   `tfsdk:"mode_trunk"`
	Nonegotiate                types.Bool   `tfsdk:"nonegotiate"`
	AccessVlan                 types.String `tfsdk:"access_vlan"`
	TrunkAllowedVlans          types.String `tfsdk:"trunk_allowed_vlans"`
	TrunkAllowedVlansNone      types.Bool   `tfsdk:"trunk_allowed_vlans_none"`
	TrunkNativeVlanTag         types.Bool   `tfsdk:"trunk_native_vlan_tag"`
	TrunkNativeVlan            types.Int64  `tfsdk:"trunk_native_vlan"`
	Host                       types.Bool   `tfsdk:"host"`
}

type InterfaceSwitchportData struct {
	Device                     types.String `tfsdk:"device"`
	Id                         types.String `tfsdk:"id"`
	Type                       types.String `tfsdk:"type"`
	Name                       types.String `tfsdk:"name"`
	ModeAccess                 types.Bool   `tfsdk:"mode_access"`
	ModeDot1qTunnel            types.Bool   `tfsdk:"mode_dot1q_tunnel"`
	ModePrivateVlanTrunk       types.Bool   `tfsdk:"mode_private_vlan_trunk"`
	ModePrivateVlanHost        types.Bool   `tfsdk:"mode_private_vlan_host"`
	ModePrivateVlanPromiscuous types.Bool   `tfsdk:"mode_private_vlan_promiscuous"`
	ModeTrunk                  types.Bool   `tfsdk:"mode_trunk"`
	Nonegotiate                types.Bool   `tfsdk:"nonegotiate"`
	AccessVlan                 types.String `tfsdk:"access_vlan"`
	TrunkAllowedVlans          types.String `tfsdk:"trunk_allowed_vlans"`
	TrunkAllowedVlansNone      types.Bool   `tfsdk:"trunk_allowed_vlans_none"`
	TrunkNativeVlanTag         types.Bool   `tfsdk:"trunk_native_vlan_tag"`
	TrunkNativeVlan            types.Int64  `tfsdk:"trunk_native_vlan"`
	Host                       types.Bool   `tfsdk:"host"`
}

func (data InterfaceSwitchport) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/interface/%s=%v/switchport-config/switchport", url.QueryEscape(fmt.Sprintf("%v", data.Type.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

func (data InterfaceSwitchportData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/interface/%s=%v/switchport-config/switchport", url.QueryEscape(fmt.Sprintf("%v", data.Type.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

// if last path element has a key -> remove it
func (data InterfaceSwitchport) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data InterfaceSwitchport) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.ModeAccess.IsNull() && !data.ModeAccess.IsUnknown() {
		if data.ModeAccess.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-switch:mode.access", map[string]string{})
		}
	}
	if !data.ModeDot1qTunnel.IsNull() && !data.ModeDot1qTunnel.IsUnknown() {
		if data.ModeDot1qTunnel.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-switch:mode.dot1q-tunnel", map[string]string{})
		}
	}
	if !data.ModePrivateVlanTrunk.IsNull() && !data.ModePrivateVlanTrunk.IsUnknown() {
		if data.ModePrivateVlanTrunk.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-switch:mode.private-vlan.trunk", map[string]string{})
		}
	}
	if !data.ModePrivateVlanHost.IsNull() && !data.ModePrivateVlanHost.IsUnknown() {
		if data.ModePrivateVlanHost.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-switch:mode.private-vlan.host", map[string]string{})
		}
	}
	if !data.ModePrivateVlanPromiscuous.IsNull() && !data.ModePrivateVlanPromiscuous.IsUnknown() {
		if data.ModePrivateVlanPromiscuous.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-switch:mode.private-vlan.promiscuous", map[string]string{})
		}
	}
	if !data.ModeTrunk.IsNull() && !data.ModeTrunk.IsUnknown() {
		if data.ModeTrunk.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-switch:mode.trunk", map[string]string{})
		}
	}
	if !data.Nonegotiate.IsNull() && !data.Nonegotiate.IsUnknown() {
		if data.Nonegotiate.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-switch:nonegotiate", map[string]string{})
		}
	}
	if !data.AccessVlan.IsNull() && !data.AccessVlan.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-switch:access.vlan.vlan", data.AccessVlan.ValueString())
	}
	if !data.TrunkAllowedVlans.IsNull() && !data.TrunkAllowedVlans.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-switch:trunk.allowed.vlan.vlans", data.TrunkAllowedVlans.ValueString())
	}
	if !data.TrunkAllowedVlansNone.IsNull() && !data.TrunkAllowedVlansNone.IsUnknown() {
		if data.TrunkAllowedVlansNone.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-switch:trunk.allowed.vlan.none", map[string]string{})
		}
	}
	if !data.TrunkNativeVlanTag.IsNull() && !data.TrunkNativeVlanTag.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-switch:trunk.native.vlan.tag", data.TrunkNativeVlanTag.ValueBool())
	}
	if !data.TrunkNativeVlan.IsNull() && !data.TrunkNativeVlan.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-switch:trunk.native.vlan.vlan-id", strconv.FormatInt(data.TrunkNativeVlan.ValueInt64(), 10))
	}
	if !data.Host.IsNull() && !data.Host.IsUnknown() {
		if data.Host.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"host", map[string]string{})
		}
	}
	return body
}

func (data *InterfaceSwitchport) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:mode.access"); !data.ModeAccess.IsNull() {
		if value.Exists() {
			data.ModeAccess = types.BoolValue(true)
		} else {
			data.ModeAccess = types.BoolValue(false)
		}
	} else {
		data.ModeAccess = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:mode.dot1q-tunnel"); !data.ModeDot1qTunnel.IsNull() {
		if value.Exists() {
			data.ModeDot1qTunnel = types.BoolValue(true)
		} else {
			data.ModeDot1qTunnel = types.BoolValue(false)
		}
	} else {
		data.ModeDot1qTunnel = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:mode.private-vlan.trunk"); !data.ModePrivateVlanTrunk.IsNull() {
		if value.Exists() {
			data.ModePrivateVlanTrunk = types.BoolValue(true)
		} else {
			data.ModePrivateVlanTrunk = types.BoolValue(false)
		}
	} else {
		data.ModePrivateVlanTrunk = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:mode.private-vlan.host"); !data.ModePrivateVlanHost.IsNull() {
		if value.Exists() {
			data.ModePrivateVlanHost = types.BoolValue(true)
		} else {
			data.ModePrivateVlanHost = types.BoolValue(false)
		}
	} else {
		data.ModePrivateVlanHost = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:mode.private-vlan.promiscuous"); !data.ModePrivateVlanPromiscuous.IsNull() {
		if value.Exists() {
			data.ModePrivateVlanPromiscuous = types.BoolValue(true)
		} else {
			data.ModePrivateVlanPromiscuous = types.BoolValue(false)
		}
	} else {
		data.ModePrivateVlanPromiscuous = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:mode.trunk"); !data.ModeTrunk.IsNull() {
		if value.Exists() {
			data.ModeTrunk = types.BoolValue(true)
		} else {
			data.ModeTrunk = types.BoolValue(false)
		}
	} else {
		data.ModeTrunk = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:nonegotiate"); !data.Nonegotiate.IsNull() {
		if value.Exists() {
			data.Nonegotiate = types.BoolValue(true)
		} else {
			data.Nonegotiate = types.BoolValue(false)
		}
	} else {
		data.Nonegotiate = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:access.vlan.vlan"); value.Exists() && !data.AccessVlan.IsNull() {
		data.AccessVlan = types.StringValue(value.String())
	} else {
		data.AccessVlan = types.StringNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:trunk.allowed.vlan.vlans"); value.Exists() && !data.TrunkAllowedVlans.IsNull() {
		data.TrunkAllowedVlans = types.StringValue(value.String())
	} else {
		data.TrunkAllowedVlans = types.StringNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:trunk.allowed.vlan.none"); !data.TrunkAllowedVlansNone.IsNull() {
		if value.Exists() {
			data.TrunkAllowedVlansNone = types.BoolValue(true)
		} else {
			data.TrunkAllowedVlansNone = types.BoolValue(false)
		}
	} else {
		data.TrunkAllowedVlansNone = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:trunk.native.vlan.tag"); !data.TrunkNativeVlanTag.IsNull() {
		if value.Exists() {
			data.TrunkNativeVlanTag = types.BoolValue(value.Bool())
		}
	} else {
		data.TrunkNativeVlanTag = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:trunk.native.vlan.vlan-id"); value.Exists() && !data.TrunkNativeVlan.IsNull() {
		data.TrunkNativeVlan = types.Int64Value(value.Int())
	} else {
		data.TrunkNativeVlan = types.Int64Null()
	}
	if value := res.Get(prefix + "host"); !data.Host.IsNull() {
		if value.Exists() {
			data.Host = types.BoolValue(true)
		} else {
			data.Host = types.BoolValue(false)
		}
	} else {
		data.Host = types.BoolNull()
	}
}

func (data *InterfaceSwitchport) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:mode.access"); value.Exists() {
		data.ModeAccess = types.BoolValue(true)
	} else {
		data.ModeAccess = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:mode.dot1q-tunnel"); value.Exists() {
		data.ModeDot1qTunnel = types.BoolValue(true)
	} else {
		data.ModeDot1qTunnel = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:mode.private-vlan.trunk"); value.Exists() {
		data.ModePrivateVlanTrunk = types.BoolValue(true)
	} else {
		data.ModePrivateVlanTrunk = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:mode.private-vlan.host"); value.Exists() {
		data.ModePrivateVlanHost = types.BoolValue(true)
	} else {
		data.ModePrivateVlanHost = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:mode.private-vlan.promiscuous"); value.Exists() {
		data.ModePrivateVlanPromiscuous = types.BoolValue(true)
	} else {
		data.ModePrivateVlanPromiscuous = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:mode.trunk"); value.Exists() {
		data.ModeTrunk = types.BoolValue(true)
	} else {
		data.ModeTrunk = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:nonegotiate"); value.Exists() {
		data.Nonegotiate = types.BoolValue(true)
	} else {
		data.Nonegotiate = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:access.vlan.vlan"); value.Exists() {
		data.AccessVlan = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:trunk.allowed.vlan.vlans"); value.Exists() {
		data.TrunkAllowedVlans = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:trunk.allowed.vlan.none"); value.Exists() {
		data.TrunkAllowedVlansNone = types.BoolValue(true)
	} else {
		data.TrunkAllowedVlansNone = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:trunk.native.vlan.tag"); value.Exists() {
		data.TrunkNativeVlanTag = types.BoolValue(value.Bool())
	} else {
		data.TrunkNativeVlanTag = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:trunk.native.vlan.vlan-id"); value.Exists() {
		data.TrunkNativeVlan = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "host"); value.Exists() {
		data.Host = types.BoolValue(true)
	} else {
		data.Host = types.BoolValue(false)
	}
}

func (data *InterfaceSwitchportData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:mode.access"); value.Exists() {
		data.ModeAccess = types.BoolValue(true)
	} else {
		data.ModeAccess = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:mode.dot1q-tunnel"); value.Exists() {
		data.ModeDot1qTunnel = types.BoolValue(true)
	} else {
		data.ModeDot1qTunnel = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:mode.private-vlan.trunk"); value.Exists() {
		data.ModePrivateVlanTrunk = types.BoolValue(true)
	} else {
		data.ModePrivateVlanTrunk = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:mode.private-vlan.host"); value.Exists() {
		data.ModePrivateVlanHost = types.BoolValue(true)
	} else {
		data.ModePrivateVlanHost = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:mode.private-vlan.promiscuous"); value.Exists() {
		data.ModePrivateVlanPromiscuous = types.BoolValue(true)
	} else {
		data.ModePrivateVlanPromiscuous = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:mode.trunk"); value.Exists() {
		data.ModeTrunk = types.BoolValue(true)
	} else {
		data.ModeTrunk = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:nonegotiate"); value.Exists() {
		data.Nonegotiate = types.BoolValue(true)
	} else {
		data.Nonegotiate = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:access.vlan.vlan"); value.Exists() {
		data.AccessVlan = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:trunk.allowed.vlan.vlans"); value.Exists() {
		data.TrunkAllowedVlans = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:trunk.allowed.vlan.none"); value.Exists() {
		data.TrunkAllowedVlansNone = types.BoolValue(true)
	} else {
		data.TrunkAllowedVlansNone = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:trunk.native.vlan.tag"); value.Exists() {
		data.TrunkNativeVlanTag = types.BoolValue(value.Bool())
	} else {
		data.TrunkNativeVlanTag = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-switch:trunk.native.vlan.vlan-id"); value.Exists() {
		data.TrunkNativeVlan = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "host"); value.Exists() {
		data.Host = types.BoolValue(true)
	} else {
		data.Host = types.BoolValue(false)
	}
}

func (data *InterfaceSwitchport) getDeletedItems(ctx context.Context, state InterfaceSwitchport) []string {
	deletedItems := make([]string, 0)
	if !state.ModeAccess.IsNull() && data.ModeAccess.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-switch:mode/access", state.getPath()))
	}
	if !state.ModeDot1qTunnel.IsNull() && data.ModeDot1qTunnel.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-switch:mode/dot1q-tunnel", state.getPath()))
	}
	if !state.ModePrivateVlanTrunk.IsNull() && data.ModePrivateVlanTrunk.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-switch:mode/private-vlan/trunk", state.getPath()))
	}
	if !state.ModePrivateVlanHost.IsNull() && data.ModePrivateVlanHost.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-switch:mode/private-vlan/host", state.getPath()))
	}
	if !state.ModePrivateVlanPromiscuous.IsNull() && data.ModePrivateVlanPromiscuous.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-switch:mode/private-vlan/promiscuous", state.getPath()))
	}
	if !state.ModeTrunk.IsNull() && data.ModeTrunk.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-switch:mode/trunk", state.getPath()))
	}
	if !state.Nonegotiate.IsNull() && data.Nonegotiate.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-switch:nonegotiate", state.getPath()))
	}
	if !state.AccessVlan.IsNull() && data.AccessVlan.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-switch:access/vlan/vlan", state.getPath()))
	}
	if !state.TrunkAllowedVlans.IsNull() && data.TrunkAllowedVlans.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-switch:trunk/allowed/vlan/vlans", state.getPath()))
	}
	if !state.TrunkAllowedVlansNone.IsNull() && data.TrunkAllowedVlansNone.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-switch:trunk/allowed/vlan/none", state.getPath()))
	}
	if !state.TrunkNativeVlanTag.IsNull() && data.TrunkNativeVlanTag.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-switch:trunk/native/vlan/tag", state.getPath()))
	}
	if !state.TrunkNativeVlan.IsNull() && data.TrunkNativeVlan.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/Cisco-IOS-XE-switch:trunk/native/vlan/vlan-id", state.getPath()))
	}
	if !state.Host.IsNull() && data.Host.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/host", state.getPath()))
	}
	return deletedItems
}

func (data *InterfaceSwitchport) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.ModeAccess.IsNull() && !data.ModeAccess.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-switch:mode/access", data.getPath()))
	}
	if !data.ModeDot1qTunnel.IsNull() && !data.ModeDot1qTunnel.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-switch:mode/dot1q-tunnel", data.getPath()))
	}
	if !data.ModePrivateVlanTrunk.IsNull() && !data.ModePrivateVlanTrunk.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-switch:mode/private-vlan/trunk", data.getPath()))
	}
	if !data.ModePrivateVlanHost.IsNull() && !data.ModePrivateVlanHost.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-switch:mode/private-vlan/host", data.getPath()))
	}
	if !data.ModePrivateVlanPromiscuous.IsNull() && !data.ModePrivateVlanPromiscuous.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-switch:mode/private-vlan/promiscuous", data.getPath()))
	}
	if !data.ModeTrunk.IsNull() && !data.ModeTrunk.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-switch:mode/trunk", data.getPath()))
	}
	if !data.Nonegotiate.IsNull() && !data.Nonegotiate.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-switch:nonegotiate", data.getPath()))
	}
	if !data.TrunkAllowedVlansNone.IsNull() && !data.TrunkAllowedVlansNone.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-switch:trunk/allowed/vlan/none", data.getPath()))
	}
	if !data.Host.IsNull() && !data.Host.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/host", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *InterfaceSwitchport) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.ModeAccess.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-switch:mode/access", data.getPath()))
	}
	if !data.ModeDot1qTunnel.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-switch:mode/dot1q-tunnel", data.getPath()))
	}
	if !data.ModePrivateVlanTrunk.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-switch:mode/private-vlan/trunk", data.getPath()))
	}
	if !data.ModePrivateVlanHost.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-switch:mode/private-vlan/host", data.getPath()))
	}
	if !data.ModePrivateVlanPromiscuous.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-switch:mode/private-vlan/promiscuous", data.getPath()))
	}
	if !data.ModeTrunk.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-switch:mode/trunk", data.getPath()))
	}
	if !data.Nonegotiate.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-switch:nonegotiate", data.getPath()))
	}
	if !data.AccessVlan.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-switch:access/vlan/vlan", data.getPath()))
	}
	if !data.TrunkAllowedVlans.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-switch:trunk/allowed/vlan/vlans", data.getPath()))
	}
	if !data.TrunkAllowedVlansNone.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-switch:trunk/allowed/vlan/none", data.getPath()))
	}
	if !data.TrunkNativeVlanTag.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-switch:trunk/native/vlan/tag", data.getPath()))
	}
	if !data.TrunkNativeVlan.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-switch:trunk/native/vlan/vlan-id", data.getPath()))
	}
	if !data.Host.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/host", data.getPath()))
	}
	return deletePaths
}

func (data *InterfaceSwitchport) getIdsFromPath() {
	reString := strings.ReplaceAll("Cisco-IOS-XE-native:native/interface/%s=%v/switchport-config/switchport", "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Id.ValueString())
	data.Type = types.StringValue(matches[1])
	data.Name = types.StringValue(matches[2])
}
