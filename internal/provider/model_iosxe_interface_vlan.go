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

type InterfaceVLAN struct {
	Device                      types.String                   `tfsdk:"device"`
	Id                          types.String                   `tfsdk:"id"`
	DeleteMode                  types.String                   `tfsdk:"delete_mode"`
	Name                        types.Int64                    `tfsdk:"name"`
	Autostate                   types.Bool                     `tfsdk:"autostate"`
	Description                 types.String                   `tfsdk:"description"`
	Shutdown                    types.Bool                     `tfsdk:"shutdown"`
	IpProxyArp                  types.Bool                     `tfsdk:"ip_proxy_arp"`
	IpRedirects                 types.Bool                     `tfsdk:"ip_redirects"`
	Unreachables                types.Bool                     `tfsdk:"unreachables"`
	VrfForwarding               types.String                   `tfsdk:"vrf_forwarding"`
	Ipv4Address                 types.String                   `tfsdk:"ipv4_address"`
	Ipv4AddressMask             types.String                   `tfsdk:"ipv4_address_mask"`
	Unnumbered                  types.String                   `tfsdk:"unnumbered"`
	IpDhcpRelaySourceInterface  types.String                   `tfsdk:"ip_dhcp_relay_source_interface"`
	IpAccessGroupIn             types.String                   `tfsdk:"ip_access_group_in"`
	IpAccessGroupInEnable       types.Bool                     `tfsdk:"ip_access_group_in_enable"`
	IpAccessGroupOut            types.String                   `tfsdk:"ip_access_group_out"`
	IpAccessGroupOutEnable      types.Bool                     `tfsdk:"ip_access_group_out_enable"`
	HelperAddresses             []InterfaceVLANHelperAddresses `tfsdk:"helper_addresses"`
	Template                    types.String                   `tfsdk:"template"`
	Enable                      types.Bool                     `tfsdk:"enable"`
	LocalAddress                types.String                   `tfsdk:"local_address"`
	IntervalInterfaceMsecs      types.Int64                    `tfsdk:"interval_interface_msecs"`
	IntervalInterfaceMinRx      types.Int64                    `tfsdk:"interval_interface_min_rx"`
	IntervalInterfaceMultiplier types.Int64                    `tfsdk:"interval_interface_multiplier"`
	Echo                        types.Bool                     `tfsdk:"echo"`
}

type InterfaceVLANData struct {
	Device                      types.String                   `tfsdk:"device"`
	Id                          types.String                   `tfsdk:"id"`
	Name                        types.Int64                    `tfsdk:"name"`
	Autostate                   types.Bool                     `tfsdk:"autostate"`
	Description                 types.String                   `tfsdk:"description"`
	Shutdown                    types.Bool                     `tfsdk:"shutdown"`
	IpProxyArp                  types.Bool                     `tfsdk:"ip_proxy_arp"`
	IpRedirects                 types.Bool                     `tfsdk:"ip_redirects"`
	Unreachables                types.Bool                     `tfsdk:"unreachables"`
	VrfForwarding               types.String                   `tfsdk:"vrf_forwarding"`
	Ipv4Address                 types.String                   `tfsdk:"ipv4_address"`
	Ipv4AddressMask             types.String                   `tfsdk:"ipv4_address_mask"`
	Unnumbered                  types.String                   `tfsdk:"unnumbered"`
	IpDhcpRelaySourceInterface  types.String                   `tfsdk:"ip_dhcp_relay_source_interface"`
	IpAccessGroupIn             types.String                   `tfsdk:"ip_access_group_in"`
	IpAccessGroupInEnable       types.Bool                     `tfsdk:"ip_access_group_in_enable"`
	IpAccessGroupOut            types.String                   `tfsdk:"ip_access_group_out"`
	IpAccessGroupOutEnable      types.Bool                     `tfsdk:"ip_access_group_out_enable"`
	HelperAddresses             []InterfaceVLANHelperAddresses `tfsdk:"helper_addresses"`
	Template                    types.String                   `tfsdk:"template"`
	Enable                      types.Bool                     `tfsdk:"enable"`
	LocalAddress                types.String                   `tfsdk:"local_address"`
	IntervalInterfaceMsecs      types.Int64                    `tfsdk:"interval_interface_msecs"`
	IntervalInterfaceMinRx      types.Int64                    `tfsdk:"interval_interface_min_rx"`
	IntervalInterfaceMultiplier types.Int64                    `tfsdk:"interval_interface_multiplier"`
	Echo                        types.Bool                     `tfsdk:"echo"`
}
type InterfaceVLANHelperAddresses struct {
	Address types.String `tfsdk:"address"`
	Global  types.Bool   `tfsdk:"global"`
	Vrf     types.String `tfsdk:"vrf"`
}

func (data InterfaceVLAN) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/interface/Vlan=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueInt64())))
}

func (data InterfaceVLANData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/interface/Vlan=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueInt64())))
}

// if last path element has a key -> remove it
func (data InterfaceVLAN) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data InterfaceVLAN) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", strconv.FormatInt(data.Name.ValueInt64(), 10))
	}
	if !data.Autostate.IsNull() && !data.Autostate.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"autostate", data.Autostate.ValueBool())
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"description", data.Description.ValueString())
	}
	if !data.Shutdown.IsNull() && !data.Shutdown.IsUnknown() {
		if data.Shutdown.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"shutdown", map[string]string{})
		}
	}
	if !data.IpProxyArp.IsNull() && !data.IpProxyArp.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.proxy-arp", data.IpProxyArp.ValueBool())
	}
	if !data.IpRedirects.IsNull() && !data.IpRedirects.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.redirects", data.IpRedirects.ValueBool())
	}
	if !data.Unreachables.IsNull() && !data.Unreachables.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.Cisco-IOS-XE-icmp:unreachables", data.Unreachables.ValueBool())
	}
	if !data.VrfForwarding.IsNull() && !data.VrfForwarding.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vrf.forwarding", data.VrfForwarding.ValueString())
	}
	if !data.Ipv4Address.IsNull() && !data.Ipv4Address.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.address.primary.address", data.Ipv4Address.ValueString())
	}
	if !data.Ipv4AddressMask.IsNull() && !data.Ipv4AddressMask.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.address.primary.mask", data.Ipv4AddressMask.ValueString())
	}
	if !data.Unnumbered.IsNull() && !data.Unnumbered.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.unnumbered", data.Unnumbered.ValueString())
	}
	if !data.IpDhcpRelaySourceInterface.IsNull() && !data.IpDhcpRelaySourceInterface.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.dhcp.Cisco-IOS-XE-dhcp:relay.source-interface", data.IpDhcpRelaySourceInterface.ValueString())
	}
	if !data.IpAccessGroupIn.IsNull() && !data.IpAccessGroupIn.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.access-group.in.acl.acl-name", data.IpAccessGroupIn.ValueString())
	}
	if !data.IpAccessGroupInEnable.IsNull() && !data.IpAccessGroupInEnable.IsUnknown() {
		if data.IpAccessGroupInEnable.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.access-group.in.acl.in", map[string]string{})
		}
	}
	if !data.IpAccessGroupOut.IsNull() && !data.IpAccessGroupOut.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.access-group.out.acl.acl-name", data.IpAccessGroupOut.ValueString())
	}
	if !data.IpAccessGroupOutEnable.IsNull() && !data.IpAccessGroupOutEnable.IsUnknown() {
		if data.IpAccessGroupOutEnable.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.access-group.out.acl.out", map[string]string{})
		}
	}
	if !data.Template.IsNull() && !data.Template.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"bfd.Cisco-IOS-XE-bfd:template", data.Template.ValueString())
	}
	if !data.Enable.IsNull() && !data.Enable.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"bfd.Cisco-IOS-XE-bfd:enable", data.Enable.ValueBool())
	}
	if !data.LocalAddress.IsNull() && !data.LocalAddress.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"bfd.Cisco-IOS-XE-bfd:local-address", data.LocalAddress.ValueString())
	}
	if !data.IntervalInterfaceMsecs.IsNull() && !data.IntervalInterfaceMsecs.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"bfd.Cisco-IOS-XE-bfd:interval-interface.msecs", strconv.FormatInt(data.IntervalInterfaceMsecs.ValueInt64(), 10))
	}
	if !data.IntervalInterfaceMinRx.IsNull() && !data.IntervalInterfaceMinRx.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"bfd.Cisco-IOS-XE-bfd:interval-interface.min_rx", strconv.FormatInt(data.IntervalInterfaceMinRx.ValueInt64(), 10))
	}
	if !data.IntervalInterfaceMultiplier.IsNull() && !data.IntervalInterfaceMultiplier.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"bfd.Cisco-IOS-XE-bfd:interval-interface.multiplier", strconv.FormatInt(data.IntervalInterfaceMultiplier.ValueInt64(), 10))
	}
	if !data.Echo.IsNull() && !data.Echo.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"bfd.Cisco-IOS-XE-bfd:echo", data.Echo.ValueBool())
	}
	if len(data.HelperAddresses) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.helper-address", []interface{}{})
		for index, item := range data.HelperAddresses {
			if !item.Address.IsNull() && !item.Address.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.helper-address"+"."+strconv.Itoa(index)+"."+"address", item.Address.ValueString())
			}
			if !item.Global.IsNull() && !item.Global.IsUnknown() {
				if item.Global.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.helper-address"+"."+strconv.Itoa(index)+"."+"global", map[string]string{})
				}
			}
			if !item.Vrf.IsNull() && !item.Vrf.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.helper-address"+"."+strconv.Itoa(index)+"."+"vrf", item.Vrf.ValueString())
			}
		}
	}
	return body
}

func (data *InterfaceVLAN) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.Int64Value(value.Int())
	} else {
		data.Name = types.Int64Null()
	}
	if value := res.Get(prefix + "autostate"); !data.Autostate.IsNull() {
		if value.Exists() {
			data.Autostate = types.BoolValue(value.Bool())
		}
	} else {
		data.Autostate = types.BoolNull()
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
	if value := res.Get(prefix + "ip.proxy-arp"); !data.IpProxyArp.IsNull() {
		if value.Exists() {
			data.IpProxyArp = types.BoolValue(value.Bool())
		}
	} else {
		data.IpProxyArp = types.BoolNull()
	}
	if value := res.Get(prefix + "ip.redirects"); !data.IpRedirects.IsNull() {
		if value.Exists() {
			data.IpRedirects = types.BoolValue(value.Bool())
		}
	} else {
		data.IpRedirects = types.BoolNull()
	}
	if value := res.Get(prefix + "ip.Cisco-IOS-XE-icmp:unreachables"); !data.Unreachables.IsNull() {
		if value.Exists() {
			data.Unreachables = types.BoolValue(value.Bool())
		}
	} else {
		data.Unreachables = types.BoolNull()
	}
	if value := res.Get(prefix + "vrf.forwarding"); value.Exists() && !data.VrfForwarding.IsNull() {
		data.VrfForwarding = types.StringValue(value.String())
	} else {
		data.VrfForwarding = types.StringNull()
	}
	if value := res.Get(prefix + "ip.address.primary.address"); value.Exists() && !data.Ipv4Address.IsNull() {
		data.Ipv4Address = types.StringValue(value.String())
	} else {
		data.Ipv4Address = types.StringNull()
	}
	if value := res.Get(prefix + "ip.address.primary.mask"); value.Exists() && !data.Ipv4AddressMask.IsNull() {
		data.Ipv4AddressMask = types.StringValue(value.String())
	} else {
		data.Ipv4AddressMask = types.StringNull()
	}
	if value := res.Get(prefix + "ip.unnumbered"); value.Exists() && !data.Unnumbered.IsNull() {
		data.Unnumbered = types.StringValue(value.String())
	} else {
		data.Unnumbered = types.StringNull()
	}
	if value := res.Get(prefix + "ip.dhcp.Cisco-IOS-XE-dhcp:relay.source-interface"); value.Exists() && !data.IpDhcpRelaySourceInterface.IsNull() {
		data.IpDhcpRelaySourceInterface = types.StringValue(value.String())
	} else {
		data.IpDhcpRelaySourceInterface = types.StringNull()
	}
	if value := res.Get(prefix + "ip.access-group.in.acl.acl-name"); value.Exists() && !data.IpAccessGroupIn.IsNull() {
		data.IpAccessGroupIn = types.StringValue(value.String())
	} else {
		data.IpAccessGroupIn = types.StringNull()
	}
	if value := res.Get(prefix + "ip.access-group.in.acl.in"); !data.IpAccessGroupInEnable.IsNull() {
		if value.Exists() {
			data.IpAccessGroupInEnable = types.BoolValue(true)
		} else {
			data.IpAccessGroupInEnable = types.BoolValue(false)
		}
	} else {
		data.IpAccessGroupInEnable = types.BoolNull()
	}
	if value := res.Get(prefix + "ip.access-group.out.acl.acl-name"); value.Exists() && !data.IpAccessGroupOut.IsNull() {
		data.IpAccessGroupOut = types.StringValue(value.String())
	} else {
		data.IpAccessGroupOut = types.StringNull()
	}
	if value := res.Get(prefix + "ip.access-group.out.acl.out"); !data.IpAccessGroupOutEnable.IsNull() {
		if value.Exists() {
			data.IpAccessGroupOutEnable = types.BoolValue(true)
		} else {
			data.IpAccessGroupOutEnable = types.BoolValue(false)
		}
	} else {
		data.IpAccessGroupOutEnable = types.BoolNull()
	}
	for i := range data.HelperAddresses {
		keys := [...]string{"address"}
		keyValues := [...]string{data.HelperAddresses[i].Address.ValueString()}

		var r gjson.Result
		res.Get(prefix + "ip.helper-address").ForEach(
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
		if value := r.Get("address"); value.Exists() && !data.HelperAddresses[i].Address.IsNull() {
			data.HelperAddresses[i].Address = types.StringValue(value.String())
		} else {
			data.HelperAddresses[i].Address = types.StringNull()
		}
		if value := r.Get("global"); !data.HelperAddresses[i].Global.IsNull() {
			if value.Exists() {
				data.HelperAddresses[i].Global = types.BoolValue(true)
			} else {
				data.HelperAddresses[i].Global = types.BoolValue(false)
			}
		} else {
			data.HelperAddresses[i].Global = types.BoolNull()
		}
		if value := r.Get("vrf"); value.Exists() && !data.HelperAddresses[i].Vrf.IsNull() {
			data.HelperAddresses[i].Vrf = types.StringValue(value.String())
		} else {
			data.HelperAddresses[i].Vrf = types.StringNull()
		}
	}
	if value := res.Get(prefix + "bfd.Cisco-IOS-XE-bfd:template"); value.Exists() && !data.Template.IsNull() {
		data.Template = types.StringValue(value.String())
	} else {
		data.Template = types.StringNull()
	}
	if value := res.Get(prefix + "bfd.Cisco-IOS-XE-bfd:enable"); !data.Enable.IsNull() {
		if value.Exists() {
			data.Enable = types.BoolValue(value.Bool())
		}
	} else {
		data.Enable = types.BoolNull()
	}
	if value := res.Get(prefix + "bfd.Cisco-IOS-XE-bfd:local-address"); value.Exists() && !data.LocalAddress.IsNull() {
		data.LocalAddress = types.StringValue(value.String())
	} else {
		data.LocalAddress = types.StringNull()
	}
	if value := res.Get(prefix + "bfd.Cisco-IOS-XE-bfd:interval-interface.msecs"); value.Exists() && !data.IntervalInterfaceMsecs.IsNull() {
		data.IntervalInterfaceMsecs = types.Int64Value(value.Int())
	} else {
		data.IntervalInterfaceMsecs = types.Int64Null()
	}
	if value := res.Get(prefix + "bfd.Cisco-IOS-XE-bfd:interval-interface.min_rx"); value.Exists() && !data.IntervalInterfaceMinRx.IsNull() {
		data.IntervalInterfaceMinRx = types.Int64Value(value.Int())
	} else {
		data.IntervalInterfaceMinRx = types.Int64Null()
	}
	if value := res.Get(prefix + "bfd.Cisco-IOS-XE-bfd:interval-interface.multiplier"); value.Exists() && !data.IntervalInterfaceMultiplier.IsNull() {
		data.IntervalInterfaceMultiplier = types.Int64Value(value.Int())
	} else {
		data.IntervalInterfaceMultiplier = types.Int64Null()
	}
	if value := res.Get(prefix + "bfd.Cisco-IOS-XE-bfd:echo"); !data.Echo.IsNull() {
		if value.Exists() {
			data.Echo = types.BoolValue(value.Bool())
		}
	} else {
		data.Echo = types.BoolNull()
	}
}

func (data *InterfaceVLANData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "autostate"); value.Exists() {
		data.Autostate = types.BoolValue(value.Bool())
	} else {
		data.Autostate = types.BoolValue(false)
	}
	if value := res.Get(prefix + "description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "shutdown"); value.Exists() {
		data.Shutdown = types.BoolValue(true)
	} else {
		data.Shutdown = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ip.proxy-arp"); value.Exists() {
		data.IpProxyArp = types.BoolValue(value.Bool())
	} else {
		data.IpProxyArp = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ip.redirects"); value.Exists() {
		data.IpRedirects = types.BoolValue(value.Bool())
	} else {
		data.IpRedirects = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ip.Cisco-IOS-XE-icmp:unreachables"); value.Exists() {
		data.Unreachables = types.BoolValue(value.Bool())
	} else {
		data.Unreachables = types.BoolValue(false)
	}
	if value := res.Get(prefix + "vrf.forwarding"); value.Exists() {
		data.VrfForwarding = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "ip.address.primary.address"); value.Exists() {
		data.Ipv4Address = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "ip.address.primary.mask"); value.Exists() {
		data.Ipv4AddressMask = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "ip.unnumbered"); value.Exists() {
		data.Unnumbered = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "ip.dhcp.Cisco-IOS-XE-dhcp:relay.source-interface"); value.Exists() {
		data.IpDhcpRelaySourceInterface = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "ip.access-group.in.acl.acl-name"); value.Exists() {
		data.IpAccessGroupIn = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "ip.access-group.in.acl.in"); value.Exists() {
		data.IpAccessGroupInEnable = types.BoolValue(true)
	} else {
		data.IpAccessGroupInEnable = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ip.access-group.out.acl.acl-name"); value.Exists() {
		data.IpAccessGroupOut = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "ip.access-group.out.acl.out"); value.Exists() {
		data.IpAccessGroupOutEnable = types.BoolValue(true)
	} else {
		data.IpAccessGroupOutEnable = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ip.helper-address"); value.Exists() {
		data.HelperAddresses = make([]InterfaceVLANHelperAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := InterfaceVLANHelperAddresses{}
			if cValue := v.Get("address"); cValue.Exists() {
				item.Address = types.StringValue(cValue.String())
			}
			if cValue := v.Get("global"); cValue.Exists() {
				item.Global = types.BoolValue(true)
			} else {
				item.Global = types.BoolValue(false)
			}
			if cValue := v.Get("vrf"); cValue.Exists() {
				item.Vrf = types.StringValue(cValue.String())
			}
			data.HelperAddresses = append(data.HelperAddresses, item)
			return true
		})
	}
	if value := res.Get(prefix + "bfd.Cisco-IOS-XE-bfd:template"); value.Exists() {
		data.Template = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "bfd.Cisco-IOS-XE-bfd:enable"); value.Exists() {
		data.Enable = types.BoolValue(value.Bool())
	} else {
		data.Enable = types.BoolValue(false)
	}
	if value := res.Get(prefix + "bfd.Cisco-IOS-XE-bfd:local-address"); value.Exists() {
		data.LocalAddress = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "bfd.Cisco-IOS-XE-bfd:interval-interface.msecs"); value.Exists() {
		data.IntervalInterfaceMsecs = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "bfd.Cisco-IOS-XE-bfd:interval-interface.min_rx"); value.Exists() {
		data.IntervalInterfaceMinRx = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "bfd.Cisco-IOS-XE-bfd:interval-interface.multiplier"); value.Exists() {
		data.IntervalInterfaceMultiplier = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "bfd.Cisco-IOS-XE-bfd:echo"); value.Exists() {
		data.Echo = types.BoolValue(value.Bool())
	} else {
		data.Echo = types.BoolValue(false)
	}
}

func (data *InterfaceVLAN) getDeletedListItems(ctx context.Context, state InterfaceVLAN) []string {
	deletedListItems := make([]string, 0)
	for i := range state.HelperAddresses {
		stateKeyValues := [...]string{state.HelperAddresses[i].Address.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.HelperAddresses[i].Address.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.HelperAddresses {
			found = true
			if state.HelperAddresses[i].Address.ValueString() != data.HelperAddresses[j].Address.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/ip/helper-address=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedListItems
}

func (data *InterfaceVLAN) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Shutdown.IsNull() && !data.Shutdown.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/shutdown", data.getPath()))
	}
	if !data.IpAccessGroupInEnable.IsNull() && !data.IpAccessGroupInEnable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ip/access-group/in/acl/in", data.getPath()))
	}
	if !data.IpAccessGroupOutEnable.IsNull() && !data.IpAccessGroupOutEnable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ip/access-group/out/acl/out", data.getPath()))
	}

	for i := range data.HelperAddresses {
		keyValues := [...]string{data.HelperAddresses[i].Address.ValueString()}
		if !data.HelperAddresses[i].Global.IsNull() && !data.HelperAddresses[i].Global.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ip/helper-address=%v/global", data.getPath(), strings.Join(keyValues[:], ",")))
		}
	}
	return emptyLeafsDelete
}

func (data *InterfaceVLAN) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Autostate.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/autostate", data.getPath()))
	}
	if !data.Description.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/description", data.getPath()))
	}
	if !data.Shutdown.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/shutdown", data.getPath()))
	}
	if !data.IpProxyArp.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/proxy-arp", data.getPath()))
	}
	if !data.IpRedirects.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/redirects", data.getPath()))
	}
	if !data.Unreachables.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/Cisco-IOS-XE-icmp:unreachables", data.getPath()))
	}
	if !data.VrfForwarding.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/vrf/forwarding", data.getPath()))
	}
	if !data.Ipv4Address.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/address/primary", data.getPath()))
	}
	if !data.Ipv4AddressMask.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/address/primary", data.getPath()))
	}
	if !data.Unnumbered.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/unnumbered", data.getPath()))
	}
	if !data.IpDhcpRelaySourceInterface.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/dhcp/Cisco-IOS-XE-dhcp:relay/source-interface", data.getPath()))
	}
	if !data.IpAccessGroupIn.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/access-group/in/acl", data.getPath()))
	}
	if !data.IpAccessGroupInEnable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/access-group/in/acl/in", data.getPath()))
	}
	if !data.IpAccessGroupOut.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/access-group/out/acl", data.getPath()))
	}
	if !data.IpAccessGroupOutEnable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/access-group/out/acl/out", data.getPath()))
	}
	for i := range data.HelperAddresses {
		keyValues := [...]string{data.HelperAddresses[i].Address.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/helper-address=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	if !data.Template.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/bfd/Cisco-IOS-XE-bfd:template", data.getPath()))
	}
	if !data.Enable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/bfd/Cisco-IOS-XE-bfd:enable", data.getPath()))
	}
	if !data.LocalAddress.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/bfd/Cisco-IOS-XE-bfd:local-address", data.getPath()))
	}
	if !data.IntervalInterfaceMsecs.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/bfd/Cisco-IOS-XE-bfd:interval-interface/msecs", data.getPath()))
	}
	if !data.IntervalInterfaceMinRx.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/bfd/Cisco-IOS-XE-bfd:interval-interface/min_rx", data.getPath()))
	}
	if !data.IntervalInterfaceMultiplier.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/bfd/Cisco-IOS-XE-bfd:interval-interface/multiplier", data.getPath()))
	}
	if !data.Echo.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/bfd/Cisco-IOS-XE-bfd:echo", data.getPath()))
	}
	return deletePaths
}
