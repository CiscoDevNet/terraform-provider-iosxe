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

type DHCP struct {
	Device                               types.String           `tfsdk:"device"`
	Id                                   types.String           `tfsdk:"id"`
	DeleteMode                           types.String           `tfsdk:"delete_mode"`
	CompatibilitySuboptionLinkSelection  types.String           `tfsdk:"compatibility_suboption_link_selection"`
	CompatibilitySuboptionServerOverride types.String           `tfsdk:"compatibility_suboption_server_override"`
	RelayInformationTrustAll             types.Bool             `tfsdk:"relay_information_trust_all"`
	RelayInformationOptionDefault        types.Bool             `tfsdk:"relay_information_option_default"`
	RelayInformationOptionVpn            types.Bool             `tfsdk:"relay_information_option_vpn"`
	SnoopingVlans                        []DHCPSnoopingVlans    `tfsdk:"snooping_vlans"`
	Snooping                             types.Bool             `tfsdk:"snooping"`
	SnoopingRemoteidHostname             types.Bool             `tfsdk:"snooping_remoteid_hostname"`
	SnoopingVlanList                     []DHCPSnoopingVlanList `tfsdk:"snooping_vlan_list"`
}

type DHCPData struct {
	Device                               types.String           `tfsdk:"device"`
	Id                                   types.String           `tfsdk:"id"`
	CompatibilitySuboptionLinkSelection  types.String           `tfsdk:"compatibility_suboption_link_selection"`
	CompatibilitySuboptionServerOverride types.String           `tfsdk:"compatibility_suboption_server_override"`
	RelayInformationTrustAll             types.Bool             `tfsdk:"relay_information_trust_all"`
	RelayInformationOptionDefault        types.Bool             `tfsdk:"relay_information_option_default"`
	RelayInformationOptionVpn            types.Bool             `tfsdk:"relay_information_option_vpn"`
	SnoopingVlans                        []DHCPSnoopingVlans    `tfsdk:"snooping_vlans"`
	Snooping                             types.Bool             `tfsdk:"snooping"`
	SnoopingRemoteidHostname             types.Bool             `tfsdk:"snooping_remoteid_hostname"`
	SnoopingVlanList                     []DHCPSnoopingVlanList `tfsdk:"snooping_vlan_list"`
}
type DHCPSnoopingVlans struct {
	VlanId types.Int64 `tfsdk:"vlan_id"`
}
type DHCPSnoopingVlanList struct {
	Id types.String `tfsdk:"id"`
}

func (data DHCP) getPath() string {
	return "Cisco-IOS-XE-native:native/ip/dhcp"
}

func (data DHCPData) getPath() string {
	return "Cisco-IOS-XE-native:native/ip/dhcp"
}

// if last path element has a key -> remove it
func (data DHCP) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data DHCP) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.CompatibilitySuboptionLinkSelection.IsNull() && !data.CompatibilitySuboptionLinkSelection.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-dhcp:compatibility.suboption.link-selection", data.CompatibilitySuboptionLinkSelection.ValueString())
	}
	if !data.CompatibilitySuboptionServerOverride.IsNull() && !data.CompatibilitySuboptionServerOverride.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-dhcp:compatibility.suboption.server-override", data.CompatibilitySuboptionServerOverride.ValueString())
	}
	if !data.RelayInformationTrustAll.IsNull() && !data.RelayInformationTrustAll.IsUnknown() {
		if data.RelayInformationTrustAll.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-dhcp:relay.information.trust-all", map[string]string{})
		}
	}
	if !data.RelayInformationOptionDefault.IsNull() && !data.RelayInformationOptionDefault.IsUnknown() {
		if data.RelayInformationOptionDefault.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-dhcp:relay.information.option.option-default", map[string]string{})
		}
	}
	if !data.RelayInformationOptionVpn.IsNull() && !data.RelayInformationOptionVpn.IsUnknown() {
		if data.RelayInformationOptionVpn.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-dhcp:relay.information.option.vpn", map[string]string{})
		}
	}
	if !data.Snooping.IsNull() && !data.Snooping.IsUnknown() {
		if data.Snooping.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-dhcp:snooping", map[string]string{})
		}
	}
	if !data.SnoopingRemoteidHostname.IsNull() && !data.SnoopingRemoteidHostname.IsUnknown() {
		if data.SnoopingRemoteidHostname.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-dhcp:snooping-conf.snooping.information.options.option.format.remote-id.hostname", map[string]string{})
		}
	}
	if len(data.SnoopingVlans) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-dhcp:snooping-conf.snooping.vlan", []interface{}{})
		for index, item := range data.SnoopingVlans {
			if !item.VlanId.IsNull() && !item.VlanId.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-dhcp:snooping-conf.snooping.vlan"+"."+strconv.Itoa(index)+"."+"id", strconv.FormatInt(item.VlanId.ValueInt64(), 10))
			}
		}
	}
	if len(data.SnoopingVlanList) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-dhcp:snooping-conf.snooping.vlan-list", []interface{}{})
		for index, item := range data.SnoopingVlanList {
			if !item.Id.IsNull() && !item.Id.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-dhcp:snooping-conf.snooping.vlan-list"+"."+strconv.Itoa(index)+"."+"id", item.Id.ValueString())
			}
		}
	}
	return body
}

func (data *DHCP) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-dhcp:compatibility.suboption.link-selection"); value.Exists() && !data.CompatibilitySuboptionLinkSelection.IsNull() {
		data.CompatibilitySuboptionLinkSelection = types.StringValue(value.String())
	} else {
		data.CompatibilitySuboptionLinkSelection = types.StringNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-dhcp:compatibility.suboption.server-override"); value.Exists() && !data.CompatibilitySuboptionServerOverride.IsNull() {
		data.CompatibilitySuboptionServerOverride = types.StringValue(value.String())
	} else {
		data.CompatibilitySuboptionServerOverride = types.StringNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-dhcp:relay.information.trust-all"); !data.RelayInformationTrustAll.IsNull() {
		if value.Exists() {
			data.RelayInformationTrustAll = types.BoolValue(true)
		} else {
			data.RelayInformationTrustAll = types.BoolValue(false)
		}
	} else {
		data.RelayInformationTrustAll = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-dhcp:relay.information.option.option-default"); !data.RelayInformationOptionDefault.IsNull() {
		if value.Exists() {
			data.RelayInformationOptionDefault = types.BoolValue(true)
		} else {
			data.RelayInformationOptionDefault = types.BoolValue(false)
		}
	} else {
		data.RelayInformationOptionDefault = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-dhcp:relay.information.option.vpn"); !data.RelayInformationOptionVpn.IsNull() {
		if value.Exists() {
			data.RelayInformationOptionVpn = types.BoolValue(true)
		} else {
			data.RelayInformationOptionVpn = types.BoolValue(false)
		}
	} else {
		data.RelayInformationOptionVpn = types.BoolNull()
	}
	for i := range data.SnoopingVlans {
		keys := [...]string{"id"}
		keyValues := [...]string{strconv.FormatInt(data.SnoopingVlans[i].VlanId.ValueInt64(), 10)}

		var r gjson.Result
		res.Get(prefix + "Cisco-IOS-XE-dhcp:snooping-conf.snooping.vlan").ForEach(
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
		if value := r.Get("id"); value.Exists() && !data.SnoopingVlans[i].VlanId.IsNull() {
			data.SnoopingVlans[i].VlanId = types.Int64Value(value.Int())
		} else {
			data.SnoopingVlans[i].VlanId = types.Int64Null()
		}
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-dhcp:snooping"); !data.Snooping.IsNull() {
		if value.Exists() {
			data.Snooping = types.BoolValue(true)
		} else {
			data.Snooping = types.BoolValue(false)
		}
	} else {
		data.Snooping = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-dhcp:snooping-conf.snooping.information.options.option.format.remote-id.hostname"); !data.SnoopingRemoteidHostname.IsNull() {
		if value.Exists() {
			data.SnoopingRemoteidHostname = types.BoolValue(true)
		} else {
			data.SnoopingRemoteidHostname = types.BoolValue(false)
		}
	} else {
		data.SnoopingRemoteidHostname = types.BoolNull()
	}
	for i := range data.SnoopingVlanList {
		keys := [...]string{"id"}
		keyValues := [...]string{data.SnoopingVlanList[i].Id.ValueString()}

		var r gjson.Result
		res.Get(prefix + "Cisco-IOS-XE-dhcp:snooping-conf.snooping.vlan-list").ForEach(
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
		if value := r.Get("id"); value.Exists() && !data.SnoopingVlanList[i].Id.IsNull() {
			data.SnoopingVlanList[i].Id = types.StringValue(value.String())
		} else {
			data.SnoopingVlanList[i].Id = types.StringNull()
		}
	}
}

func (data *DHCPData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-dhcp:compatibility.suboption.link-selection"); value.Exists() {
		data.CompatibilitySuboptionLinkSelection = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-dhcp:compatibility.suboption.server-override"); value.Exists() {
		data.CompatibilitySuboptionServerOverride = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-dhcp:relay.information.trust-all"); value.Exists() {
		data.RelayInformationTrustAll = types.BoolValue(true)
	} else {
		data.RelayInformationTrustAll = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-dhcp:relay.information.option.option-default"); value.Exists() {
		data.RelayInformationOptionDefault = types.BoolValue(true)
	} else {
		data.RelayInformationOptionDefault = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-dhcp:relay.information.option.vpn"); value.Exists() {
		data.RelayInformationOptionVpn = types.BoolValue(true)
	} else {
		data.RelayInformationOptionVpn = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-dhcp:snooping-conf.snooping.vlan"); value.Exists() {
		data.SnoopingVlans = make([]DHCPSnoopingVlans, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := DHCPSnoopingVlans{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.VlanId = types.Int64Value(cValue.Int())
			}
			data.SnoopingVlans = append(data.SnoopingVlans, item)
			return true
		})
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-dhcp:snooping"); value.Exists() {
		data.Snooping = types.BoolValue(true)
	} else {
		data.Snooping = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-dhcp:snooping-conf.snooping.information.options.option.format.remote-id.hostname"); value.Exists() {
		data.SnoopingRemoteidHostname = types.BoolValue(true)
	} else {
		data.SnoopingRemoteidHostname = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-dhcp:snooping-conf.snooping.vlan-list"); value.Exists() {
		data.SnoopingVlanList = make([]DHCPSnoopingVlanList, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := DHCPSnoopingVlanList{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			}
			data.SnoopingVlanList = append(data.SnoopingVlanList, item)
			return true
		})
	}
}

func (data *DHCP) getDeletedListItems(ctx context.Context, state DHCP) []string {
	deletedListItems := make([]string, 0)
	for i := range state.SnoopingVlans {
		stateKeyValues := [...]string{strconv.FormatInt(state.SnoopingVlans[i].VlanId.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.SnoopingVlans[i].VlanId.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.SnoopingVlans {
			found = true
			if state.SnoopingVlans[i].VlanId.ValueInt64() != data.SnoopingVlans[j].VlanId.ValueInt64() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/Cisco-IOS-XE-dhcp:snooping-conf/snooping/vlan=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	for i := range state.SnoopingVlanList {
		stateKeyValues := [...]string{state.SnoopingVlanList[i].Id.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.SnoopingVlanList[i].Id.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.SnoopingVlanList {
			found = true
			if state.SnoopingVlanList[i].Id.ValueString() != data.SnoopingVlanList[j].Id.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/Cisco-IOS-XE-dhcp:snooping-conf/snooping/vlan-list=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedListItems
}

func (data *DHCP) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.RelayInformationTrustAll.IsNull() && !data.RelayInformationTrustAll.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-dhcp:relay/information/trust-all", data.getPath()))
	}
	if !data.RelayInformationOptionDefault.IsNull() && !data.RelayInformationOptionDefault.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-dhcp:relay/information/option/option-default", data.getPath()))
	}
	if !data.RelayInformationOptionVpn.IsNull() && !data.RelayInformationOptionVpn.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-dhcp:relay/information/option/vpn", data.getPath()))
	}

	if !data.Snooping.IsNull() && !data.Snooping.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-dhcp:snooping", data.getPath()))
	}
	if !data.SnoopingRemoteidHostname.IsNull() && !data.SnoopingRemoteidHostname.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-dhcp:snooping-conf/snooping/information/options/option/format/remote-id/hostname", data.getPath()))
	}

	return emptyLeafsDelete
}

func (data *DHCP) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.CompatibilitySuboptionLinkSelection.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-dhcp:compatibility/suboption/link-selection", data.getPath()))
	}
	if !data.CompatibilitySuboptionServerOverride.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-dhcp:compatibility/suboption/server-override", data.getPath()))
	}
	if !data.RelayInformationTrustAll.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-dhcp:relay/information/trust-all", data.getPath()))
	}
	if !data.RelayInformationOptionDefault.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-dhcp:relay/information/option/option-default", data.getPath()))
	}
	if !data.RelayInformationOptionVpn.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-dhcp:relay/information/option/vpn", data.getPath()))
	}
	for i := range data.SnoopingVlans {
		keyValues := [...]string{strconv.FormatInt(data.SnoopingVlans[i].VlanId.ValueInt64(), 10)}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-dhcp:snooping-conf/snooping/vlan=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	if !data.Snooping.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-dhcp:snooping", data.getPath()))
	}
	if !data.SnoopingRemoteidHostname.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-dhcp:snooping-conf/snooping/information/options/option/format/remote-id/hostname", data.getPath()))
	}
	for i := range data.SnoopingVlanList {
		keyValues := [...]string{data.SnoopingVlanList[i].Id.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-dhcp:snooping-conf/snooping/vlan-list=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	return deletePaths
}
