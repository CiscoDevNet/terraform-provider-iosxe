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

type InterfaceLoopback struct {
	Device                       types.String                              `tfsdk:"device"`
	Id                           types.String                              `tfsdk:"id"`
	DeleteMode                   types.String                              `tfsdk:"delete_mode"`
	Name                         types.Int64                               `tfsdk:"name"`
	Description                  types.String                              `tfsdk:"description"`
	Shutdown                     types.Bool                                `tfsdk:"shutdown"`
	IpProxyArp                   types.Bool                                `tfsdk:"ip_proxy_arp"`
	IpRedirects                  types.Bool                                `tfsdk:"ip_redirects"`
	Unreachables                 types.Bool                                `tfsdk:"unreachables"`
	VrfForwarding                types.String                              `tfsdk:"vrf_forwarding"`
	Ipv4Address                  types.String                              `tfsdk:"ipv4_address"`
	Ipv4AddressMask              types.String                              `tfsdk:"ipv4_address_mask"`
	IpAccessGroupIn              types.String                              `tfsdk:"ip_access_group_in"`
	IpAccessGroupInEnable        types.Bool                                `tfsdk:"ip_access_group_in_enable"`
	IpAccessGroupOut             types.String                              `tfsdk:"ip_access_group_out"`
	IpAccessGroupOutEnable       types.Bool                                `tfsdk:"ip_access_group_out_enable"`
	Ipv6Enable                   types.Bool                                `tfsdk:"ipv6_enable"`
	Ipv6Mtu                      types.Int64                               `tfsdk:"ipv6_mtu"`
	RaSuppressAll                types.Bool                                `tfsdk:"ra_suppress_all"`
	Ipv6AddressAutoconfigDefault types.Bool                                `tfsdk:"ipv6_address_autoconfig_default"`
	Ipv6AddressDhcp              types.Bool                                `tfsdk:"ipv6_address_dhcp"`
	Ipv6LinkLocalAddresses       []InterfaceLoopbackIpv6LinkLocalAddresses `tfsdk:"ipv6_link_local_addresses"`
	Ipv6AddressPrefixLists       []InterfaceLoopbackIpv6AddressPrefixLists `tfsdk:"ipv6_address_prefix_lists"`
}

type InterfaceLoopbackData struct {
	Device                       types.String                              `tfsdk:"device"`
	Id                           types.String                              `tfsdk:"id"`
	Name                         types.Int64                               `tfsdk:"name"`
	Description                  types.String                              `tfsdk:"description"`
	Shutdown                     types.Bool                                `tfsdk:"shutdown"`
	IpProxyArp                   types.Bool                                `tfsdk:"ip_proxy_arp"`
	IpRedirects                  types.Bool                                `tfsdk:"ip_redirects"`
	Unreachables                 types.Bool                                `tfsdk:"unreachables"`
	VrfForwarding                types.String                              `tfsdk:"vrf_forwarding"`
	Ipv4Address                  types.String                              `tfsdk:"ipv4_address"`
	Ipv4AddressMask              types.String                              `tfsdk:"ipv4_address_mask"`
	IpAccessGroupIn              types.String                              `tfsdk:"ip_access_group_in"`
	IpAccessGroupInEnable        types.Bool                                `tfsdk:"ip_access_group_in_enable"`
	IpAccessGroupOut             types.String                              `tfsdk:"ip_access_group_out"`
	IpAccessGroupOutEnable       types.Bool                                `tfsdk:"ip_access_group_out_enable"`
	Ipv6Enable                   types.Bool                                `tfsdk:"ipv6_enable"`
	Ipv6Mtu                      types.Int64                               `tfsdk:"ipv6_mtu"`
	RaSuppressAll                types.Bool                                `tfsdk:"ra_suppress_all"`
	Ipv6AddressAutoconfigDefault types.Bool                                `tfsdk:"ipv6_address_autoconfig_default"`
	Ipv6AddressDhcp              types.Bool                                `tfsdk:"ipv6_address_dhcp"`
	Ipv6LinkLocalAddresses       []InterfaceLoopbackIpv6LinkLocalAddresses `tfsdk:"ipv6_link_local_addresses"`
	Ipv6AddressPrefixLists       []InterfaceLoopbackIpv6AddressPrefixLists `tfsdk:"ipv6_address_prefix_lists"`
}
type InterfaceLoopbackIpv6LinkLocalAddresses struct {
	Address   types.String `tfsdk:"address"`
	LinkLocal types.Bool   `tfsdk:"link_local"`
}
type InterfaceLoopbackIpv6AddressPrefixLists struct {
	Prefix types.String `tfsdk:"prefix"`
	Eui64  types.Bool   `tfsdk:"eui_64"`
}

func (data InterfaceLoopback) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/interface/Loopback=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueInt64())))
}

func (data InterfaceLoopbackData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/interface/Loopback=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueInt64())))
}

// if last path element has a key -> remove it
func (data InterfaceLoopback) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data InterfaceLoopback) toBody(ctx context.Context) string {
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
	if !data.Ipv6Enable.IsNull() && !data.Ipv6Enable.IsUnknown() {
		if data.Ipv6Enable.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv6.enable", map[string]string{})
		}
	}
	if !data.Ipv6Mtu.IsNull() && !data.Ipv6Mtu.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv6.mtu", strconv.FormatInt(data.Ipv6Mtu.ValueInt64(), 10))
	}
	if !data.RaSuppressAll.IsNull() && !data.RaSuppressAll.IsUnknown() {
		if data.RaSuppressAll.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv6.nd.Cisco-IOS-XE-nd:ra.suppress.all", map[string]string{})
		}
	}
	if !data.Ipv6AddressAutoconfigDefault.IsNull() && !data.Ipv6AddressAutoconfigDefault.IsUnknown() {
		if data.Ipv6AddressAutoconfigDefault.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv6.address.autoconfig.default", map[string]string{})
		}
	}
	if !data.Ipv6AddressDhcp.IsNull() && !data.Ipv6AddressDhcp.IsUnknown() {
		if data.Ipv6AddressDhcp.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv6.address.dhcp", map[string]string{})
		}
	}
	if len(data.Ipv6LinkLocalAddresses) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv6.address.link-local-address", []interface{}{})
		for index, item := range data.Ipv6LinkLocalAddresses {
			if !item.Address.IsNull() && !item.Address.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv6.address.link-local-address"+"."+strconv.Itoa(index)+"."+"address", item.Address.ValueString())
			}
			if !item.LinkLocal.IsNull() && !item.LinkLocal.IsUnknown() {
				if item.LinkLocal.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv6.address.link-local-address"+"."+strconv.Itoa(index)+"."+"link-local", map[string]string{})
				}
			}
		}
	}
	if len(data.Ipv6AddressPrefixLists) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv6.address.prefix-list", []interface{}{})
		for index, item := range data.Ipv6AddressPrefixLists {
			if !item.Prefix.IsNull() && !item.Prefix.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv6.address.prefix-list"+"."+strconv.Itoa(index)+"."+"prefix", item.Prefix.ValueString())
			}
			if !item.Eui64.IsNull() && !item.Eui64.IsUnknown() {
				if item.Eui64.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv6.address.prefix-list"+"."+strconv.Itoa(index)+"."+"eui-64", map[string]string{})
				}
			}
		}
	}
	return body
}

func (data *InterfaceLoopback) updateFromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(prefix + "ipv6.enable"); !data.Ipv6Enable.IsNull() {
		if value.Exists() {
			data.Ipv6Enable = types.BoolValue(true)
		} else {
			data.Ipv6Enable = types.BoolValue(false)
		}
	} else {
		data.Ipv6Enable = types.BoolNull()
	}
	if value := res.Get(prefix + "ipv6.mtu"); value.Exists() && !data.Ipv6Mtu.IsNull() {
		data.Ipv6Mtu = types.Int64Value(value.Int())
	} else {
		data.Ipv6Mtu = types.Int64Null()
	}
	if value := res.Get(prefix + "ipv6.nd.Cisco-IOS-XE-nd:ra.suppress.all"); !data.RaSuppressAll.IsNull() {
		if value.Exists() {
			data.RaSuppressAll = types.BoolValue(true)
		} else {
			data.RaSuppressAll = types.BoolValue(false)
		}
	} else {
		data.RaSuppressAll = types.BoolNull()
	}
	if value := res.Get(prefix + "ipv6.address.autoconfig.default"); !data.Ipv6AddressAutoconfigDefault.IsNull() {
		if value.Exists() {
			data.Ipv6AddressAutoconfigDefault = types.BoolValue(true)
		} else {
			data.Ipv6AddressAutoconfigDefault = types.BoolValue(false)
		}
	} else {
		data.Ipv6AddressAutoconfigDefault = types.BoolNull()
	}
	if value := res.Get(prefix + "ipv6.address.dhcp"); !data.Ipv6AddressDhcp.IsNull() {
		if value.Exists() {
			data.Ipv6AddressDhcp = types.BoolValue(true)
		} else {
			data.Ipv6AddressDhcp = types.BoolValue(false)
		}
	} else {
		data.Ipv6AddressDhcp = types.BoolNull()
	}
	for i := range data.Ipv6LinkLocalAddresses {
		keys := [...]string{"address"}
		keyValues := [...]string{data.Ipv6LinkLocalAddresses[i].Address.ValueString()}

		var r gjson.Result
		res.Get(prefix + "ipv6.address.link-local-address").ForEach(
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
		if value := r.Get("address"); value.Exists() && !data.Ipv6LinkLocalAddresses[i].Address.IsNull() {
			data.Ipv6LinkLocalAddresses[i].Address = types.StringValue(value.String())
		} else {
			data.Ipv6LinkLocalAddresses[i].Address = types.StringNull()
		}
		if value := r.Get("link-local"); !data.Ipv6LinkLocalAddresses[i].LinkLocal.IsNull() {
			if value.Exists() {
				data.Ipv6LinkLocalAddresses[i].LinkLocal = types.BoolValue(true)
			} else {
				data.Ipv6LinkLocalAddresses[i].LinkLocal = types.BoolValue(false)
			}
		} else {
			data.Ipv6LinkLocalAddresses[i].LinkLocal = types.BoolNull()
		}
	}
	for i := range data.Ipv6AddressPrefixLists {
		keys := [...]string{"prefix"}
		keyValues := [...]string{data.Ipv6AddressPrefixLists[i].Prefix.ValueString()}

		var r gjson.Result
		res.Get(prefix + "ipv6.address.prefix-list").ForEach(
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
		if value := r.Get("prefix"); value.Exists() && !data.Ipv6AddressPrefixLists[i].Prefix.IsNull() {
			data.Ipv6AddressPrefixLists[i].Prefix = types.StringValue(value.String())
		} else {
			data.Ipv6AddressPrefixLists[i].Prefix = types.StringNull()
		}
		if value := r.Get("eui-64"); !data.Ipv6AddressPrefixLists[i].Eui64.IsNull() {
			if value.Exists() {
				data.Ipv6AddressPrefixLists[i].Eui64 = types.BoolValue(true)
			} else {
				data.Ipv6AddressPrefixLists[i].Eui64 = types.BoolValue(false)
			}
		} else {
			data.Ipv6AddressPrefixLists[i].Eui64 = types.BoolNull()
		}
	}
}

func (data *InterfaceLoopbackData) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(prefix + "ipv6.enable"); value.Exists() {
		data.Ipv6Enable = types.BoolValue(true)
	} else {
		data.Ipv6Enable = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ipv6.mtu"); value.Exists() {
		data.Ipv6Mtu = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "ipv6.nd.Cisco-IOS-XE-nd:ra.suppress.all"); value.Exists() {
		data.RaSuppressAll = types.BoolValue(true)
	} else {
		data.RaSuppressAll = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ipv6.address.autoconfig.default"); value.Exists() {
		data.Ipv6AddressAutoconfigDefault = types.BoolValue(true)
	} else {
		data.Ipv6AddressAutoconfigDefault = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ipv6.address.dhcp"); value.Exists() {
		data.Ipv6AddressDhcp = types.BoolValue(true)
	} else {
		data.Ipv6AddressDhcp = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ipv6.address.link-local-address"); value.Exists() {
		data.Ipv6LinkLocalAddresses = make([]InterfaceLoopbackIpv6LinkLocalAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := InterfaceLoopbackIpv6LinkLocalAddresses{}
			if cValue := v.Get("address"); cValue.Exists() {
				item.Address = types.StringValue(cValue.String())
			}
			if cValue := v.Get("link-local"); cValue.Exists() {
				item.LinkLocal = types.BoolValue(true)
			} else {
				item.LinkLocal = types.BoolValue(false)
			}
			data.Ipv6LinkLocalAddresses = append(data.Ipv6LinkLocalAddresses, item)
			return true
		})
	}
	if value := res.Get(prefix + "ipv6.address.prefix-list"); value.Exists() {
		data.Ipv6AddressPrefixLists = make([]InterfaceLoopbackIpv6AddressPrefixLists, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := InterfaceLoopbackIpv6AddressPrefixLists{}
			if cValue := v.Get("prefix"); cValue.Exists() {
				item.Prefix = types.StringValue(cValue.String())
			}
			if cValue := v.Get("eui-64"); cValue.Exists() {
				item.Eui64 = types.BoolValue(true)
			} else {
				item.Eui64 = types.BoolValue(false)
			}
			data.Ipv6AddressPrefixLists = append(data.Ipv6AddressPrefixLists, item)
			return true
		})
	}
}

func (data *InterfaceLoopback) getDeletedListItems(ctx context.Context, state InterfaceLoopback) []string {
	deletedListItems := make([]string, 0)
	for i := range state.Ipv6LinkLocalAddresses {
		stateKeyValues := [...]string{state.Ipv6LinkLocalAddresses[i].Address.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Ipv6LinkLocalAddresses[i].Address.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Ipv6LinkLocalAddresses {
			found = true
			if state.Ipv6LinkLocalAddresses[i].Address.ValueString() != data.Ipv6LinkLocalAddresses[j].Address.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/ipv6/address/link-local-address=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	for i := range state.Ipv6AddressPrefixLists {
		stateKeyValues := [...]string{state.Ipv6AddressPrefixLists[i].Prefix.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Ipv6AddressPrefixLists[i].Prefix.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Ipv6AddressPrefixLists {
			found = true
			if state.Ipv6AddressPrefixLists[i].Prefix.ValueString() != data.Ipv6AddressPrefixLists[j].Prefix.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/ipv6/address/prefix-list=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedListItems
}

func (data *InterfaceLoopback) getEmptyLeafsDelete(ctx context.Context) []string {
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
	if !data.Ipv6Enable.IsNull() && !data.Ipv6Enable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ipv6/enable", data.getPath()))
	}
	if !data.RaSuppressAll.IsNull() && !data.RaSuppressAll.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ipv6/nd/Cisco-IOS-XE-nd:ra/suppress/all", data.getPath()))
	}
	if !data.Ipv6AddressAutoconfigDefault.IsNull() && !data.Ipv6AddressAutoconfigDefault.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ipv6/address/autoconfig/default", data.getPath()))
	}
	if !data.Ipv6AddressDhcp.IsNull() && !data.Ipv6AddressDhcp.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ipv6/address/dhcp", data.getPath()))
	}

	for i := range data.Ipv6LinkLocalAddresses {
		keyValues := [...]string{data.Ipv6LinkLocalAddresses[i].Address.ValueString()}
		if !data.Ipv6LinkLocalAddresses[i].LinkLocal.IsNull() && !data.Ipv6LinkLocalAddresses[i].LinkLocal.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ipv6/address/link-local-address=%v/link-local", data.getPath(), strings.Join(keyValues[:], ",")))
		}
	}

	for i := range data.Ipv6AddressPrefixLists {
		keyValues := [...]string{data.Ipv6AddressPrefixLists[i].Prefix.ValueString()}
		if !data.Ipv6AddressPrefixLists[i].Eui64.IsNull() && !data.Ipv6AddressPrefixLists[i].Eui64.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ipv6/address/prefix-list=%v/eui-64", data.getPath(), strings.Join(keyValues[:], ",")))
		}
	}
	return emptyLeafsDelete
}

func (data *InterfaceLoopback) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
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
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/address/primary/address", data.getPath()))
	}
	if !data.Ipv4AddressMask.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/address/primary/mask", data.getPath()))
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
	if !data.Ipv6Enable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ipv6/enable", data.getPath()))
	}
	if !data.Ipv6Mtu.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ipv6/mtu", data.getPath()))
	}
	if !data.RaSuppressAll.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ipv6/nd/Cisco-IOS-XE-nd:ra/suppress/all", data.getPath()))
	}
	if !data.Ipv6AddressAutoconfigDefault.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ipv6/address/autoconfig/default", data.getPath()))
	}
	if !data.Ipv6AddressDhcp.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ipv6/address/dhcp", data.getPath()))
	}
	for i := range data.Ipv6LinkLocalAddresses {
		keyValues := [...]string{data.Ipv6LinkLocalAddresses[i].Address.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/ipv6/address/link-local-address=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	for i := range data.Ipv6AddressPrefixLists {
		keyValues := [...]string{data.Ipv6AddressPrefixLists[i].Prefix.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/ipv6/address/prefix-list=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	return deletePaths
}
