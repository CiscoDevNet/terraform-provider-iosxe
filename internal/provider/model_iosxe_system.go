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

type System struct {
	Device                      types.String                 `tfsdk:"device"`
	Id                          types.String                 `tfsdk:"id"`
	Hostname                    types.String                 `tfsdk:"hostname"`
	IpRouting                   types.Bool                   `tfsdk:"ip_routing"`
	Ipv6UnicastRouting          types.Bool                   `tfsdk:"ipv6_unicast_routing"`
	Mtu                         types.Int64                  `tfsdk:"mtu"`
	IpSourceRoute               types.Bool                   `tfsdk:"ip_source_route"`
	IpDomainLookup              types.Bool                   `tfsdk:"ip_domain_lookup"`
	IpDomainName                types.String                 `tfsdk:"ip_domain_name"`
	LoginDelay                  types.Int64                  `tfsdk:"login_delay"`
	LoginOnFailure              types.Bool                   `tfsdk:"login_on_failure"`
	LoginOnFailureLog           types.Bool                   `tfsdk:"login_on_failure_log"`
	LoginOnSuccess              types.Bool                   `tfsdk:"login_on_success"`
	LoginOnSuccessLog           types.Bool                   `tfsdk:"login_on_success_log"`
	MulticastRouting            types.Bool                   `tfsdk:"multicast_routing"`
	MulticastRoutingSwitch      types.Bool                   `tfsdk:"multicast_routing_switch"`
	MulticastRoutingDistributed types.Bool                   `tfsdk:"multicast_routing_distributed"`
	MulticastRoutingVrfs        []SystemMulticastRoutingVrfs `tfsdk:"multicast_routing_vrfs"`
}

type SystemData struct {
	Device                      types.String                 `tfsdk:"device"`
	Id                          types.String                 `tfsdk:"id"`
	Hostname                    types.String                 `tfsdk:"hostname"`
	IpRouting                   types.Bool                   `tfsdk:"ip_routing"`
	Ipv6UnicastRouting          types.Bool                   `tfsdk:"ipv6_unicast_routing"`
	Mtu                         types.Int64                  `tfsdk:"mtu"`
	IpSourceRoute               types.Bool                   `tfsdk:"ip_source_route"`
	IpDomainLookup              types.Bool                   `tfsdk:"ip_domain_lookup"`
	IpDomainName                types.String                 `tfsdk:"ip_domain_name"`
	LoginDelay                  types.Int64                  `tfsdk:"login_delay"`
	LoginOnFailure              types.Bool                   `tfsdk:"login_on_failure"`
	LoginOnFailureLog           types.Bool                   `tfsdk:"login_on_failure_log"`
	LoginOnSuccess              types.Bool                   `tfsdk:"login_on_success"`
	LoginOnSuccessLog           types.Bool                   `tfsdk:"login_on_success_log"`
	MulticastRouting            types.Bool                   `tfsdk:"multicast_routing"`
	MulticastRoutingSwitch      types.Bool                   `tfsdk:"multicast_routing_switch"`
	MulticastRoutingDistributed types.Bool                   `tfsdk:"multicast_routing_distributed"`
	MulticastRoutingVrfs        []SystemMulticastRoutingVrfs `tfsdk:"multicast_routing_vrfs"`
}
type SystemMulticastRoutingVrfs struct {
	Vrf         types.String `tfsdk:"vrf"`
	Distributed types.Bool   `tfsdk:"distributed"`
}

func (data System) getPath() string {
	return "Cisco-IOS-XE-native:native"
}

func (data SystemData) getPath() string {
	return "Cisco-IOS-XE-native:native"
}

// if last path element has a key -> remove it
func (data System) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data System) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Hostname.IsNull() && !data.Hostname.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"hostname", data.Hostname.ValueString())
	}
	if !data.IpRouting.IsNull() && !data.IpRouting.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.routing-conf.routing", data.IpRouting.ValueBool())
	}
	if !data.Ipv6UnicastRouting.IsNull() && !data.Ipv6UnicastRouting.IsUnknown() {
		if data.Ipv6UnicastRouting.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv6.unicast-routing", map[string]string{})
		}
	}
	if !data.Mtu.IsNull() && !data.Mtu.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"system.Cisco-IOS-XE-switch:mtu.size", strconv.FormatInt(data.Mtu.ValueInt64(), 10))
	}
	if !data.IpSourceRoute.IsNull() && !data.IpSourceRoute.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.source-route", data.IpSourceRoute.ValueBool())
	}
	if !data.IpDomainLookup.IsNull() && !data.IpDomainLookup.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.domain.lookup", data.IpDomainLookup.ValueBool())
	}
	if !data.IpDomainName.IsNull() && !data.IpDomainName.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.domain.name", data.IpDomainName.ValueString())
	}
	if !data.LoginDelay.IsNull() && !data.LoginDelay.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"login.delay", strconv.FormatInt(data.LoginDelay.ValueInt64(), 10))
	}
	if !data.LoginOnFailure.IsNull() && !data.LoginOnFailure.IsUnknown() {
		if data.LoginOnFailure.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"login.on-failure", map[string]string{})
		}
	}
	if !data.LoginOnFailureLog.IsNull() && !data.LoginOnFailureLog.IsUnknown() {
		if data.LoginOnFailureLog.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"login.on-failure.log", map[string]string{})
		}
	}
	if !data.LoginOnSuccess.IsNull() && !data.LoginOnSuccess.IsUnknown() {
		if data.LoginOnSuccess.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"login.on-success", map[string]string{})
		}
	}
	if !data.LoginOnSuccessLog.IsNull() && !data.LoginOnSuccessLog.IsUnknown() {
		if data.LoginOnSuccessLog.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"login.on-success.log", map[string]string{})
		}
	}
	if !data.MulticastRouting.IsNull() && !data.MulticastRouting.IsUnknown() {
		if data.MulticastRouting.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.Cisco-IOS-XE-multicast:multicast-routing", map[string]string{})
		}
	}
	if !data.MulticastRoutingSwitch.IsNull() && !data.MulticastRoutingSwitch.IsUnknown() {
		if data.MulticastRoutingSwitch.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.Cisco-IOS-XE-multicast:mcr-conf.multicast-routing", map[string]string{})
		}
	}
	if !data.MulticastRoutingDistributed.IsNull() && !data.MulticastRoutingDistributed.IsUnknown() {
		if data.MulticastRoutingDistributed.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.Cisco-IOS-XE-multicast:multicast-routing.distributed", map[string]string{})
		}
	}
	if len(data.MulticastRoutingVrfs) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.Cisco-IOS-XE-multicast:multicast-routing.vrf", []interface{}{})
		for index, item := range data.MulticastRoutingVrfs {
			if !item.Vrf.IsNull() && !item.Vrf.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.Cisco-IOS-XE-multicast:multicast-routing.vrf"+"."+strconv.Itoa(index)+"."+"name", item.Vrf.ValueString())
			}
			if !item.Distributed.IsNull() && !item.Distributed.IsUnknown() {
				if item.Distributed.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ip.Cisco-IOS-XE-multicast:multicast-routing.vrf"+"."+strconv.Itoa(index)+"."+"distributed", map[string]string{})
				}
			}
		}
	}
	return body
}

func (data *System) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "hostname"); value.Exists() && !data.Hostname.IsNull() {
		data.Hostname = types.StringValue(value.String())
	} else {
		data.Hostname = types.StringNull()
	}
	if value := res.Get(prefix + "ip.routing-conf.routing"); !data.IpRouting.IsNull() {
		if value.Exists() {
			data.IpRouting = types.BoolValue(value.Bool())
		}
	} else {
		data.IpRouting = types.BoolNull()
	}
	if value := res.Get(prefix + "ipv6.unicast-routing"); !data.Ipv6UnicastRouting.IsNull() {
		if value.Exists() {
			data.Ipv6UnicastRouting = types.BoolValue(true)
		} else {
			data.Ipv6UnicastRouting = types.BoolValue(false)
		}
	} else {
		data.Ipv6UnicastRouting = types.BoolNull()
	}
	if value := res.Get(prefix + "system.Cisco-IOS-XE-switch:mtu.size"); value.Exists() && !data.Mtu.IsNull() {
		data.Mtu = types.Int64Value(value.Int())
	} else {
		data.Mtu = types.Int64Null()
	}
	if value := res.Get(prefix + "ip.source-route"); !data.IpSourceRoute.IsNull() {
		if value.Exists() {
			data.IpSourceRoute = types.BoolValue(value.Bool())
		}
	} else {
		data.IpSourceRoute = types.BoolNull()
	}
	if value := res.Get(prefix + "ip.domain.lookup"); !data.IpDomainLookup.IsNull() {
		if value.Exists() {
			data.IpDomainLookup = types.BoolValue(value.Bool())
		}
	} else {
		data.IpDomainLookup = types.BoolNull()
	}
	if value := res.Get(prefix + "ip.domain.name"); value.Exists() && !data.IpDomainName.IsNull() {
		data.IpDomainName = types.StringValue(value.String())
	} else {
		data.IpDomainName = types.StringNull()
	}
	if value := res.Get(prefix + "login.delay"); value.Exists() && !data.LoginDelay.IsNull() {
		data.LoginDelay = types.Int64Value(value.Int())
	} else {
		data.LoginDelay = types.Int64Null()
	}
	if value := res.Get(prefix + "login.on-failure"); !data.LoginOnFailure.IsNull() {
		if value.Exists() {
			data.LoginOnFailure = types.BoolValue(true)
		} else {
			data.LoginOnFailure = types.BoolValue(false)
		}
	} else {
		data.LoginOnFailure = types.BoolNull()
	}
	if value := res.Get(prefix + "login.on-failure.log"); !data.LoginOnFailureLog.IsNull() {
		if value.Exists() {
			data.LoginOnFailureLog = types.BoolValue(true)
		} else {
			data.LoginOnFailureLog = types.BoolValue(false)
		}
	} else {
		data.LoginOnFailureLog = types.BoolNull()
	}
	if value := res.Get(prefix + "login.on-success"); !data.LoginOnSuccess.IsNull() {
		if value.Exists() {
			data.LoginOnSuccess = types.BoolValue(true)
		} else {
			data.LoginOnSuccess = types.BoolValue(false)
		}
	} else {
		data.LoginOnSuccess = types.BoolNull()
	}
	if value := res.Get(prefix + "login.on-success.log"); !data.LoginOnSuccessLog.IsNull() {
		if value.Exists() {
			data.LoginOnSuccessLog = types.BoolValue(true)
		} else {
			data.LoginOnSuccessLog = types.BoolValue(false)
		}
	} else {
		data.LoginOnSuccessLog = types.BoolNull()
	}
	if value := res.Get(prefix + "ip.Cisco-IOS-XE-multicast:multicast-routing"); !data.MulticastRouting.IsNull() {
		if value.Exists() {
			data.MulticastRouting = types.BoolValue(true)
		} else {
			data.MulticastRouting = types.BoolValue(false)
		}
	} else {
		data.MulticastRouting = types.BoolNull()
	}
	if value := res.Get(prefix + "ip.Cisco-IOS-XE-multicast:mcr-conf.multicast-routing"); !data.MulticastRoutingSwitch.IsNull() {
		if value.Exists() {
			data.MulticastRoutingSwitch = types.BoolValue(true)
		} else {
			data.MulticastRoutingSwitch = types.BoolValue(false)
		}
	} else {
		data.MulticastRoutingSwitch = types.BoolNull()
	}
	if value := res.Get(prefix + "ip.Cisco-IOS-XE-multicast:multicast-routing.distributed"); !data.MulticastRoutingDistributed.IsNull() {
		if value.Exists() {
			data.MulticastRoutingDistributed = types.BoolValue(true)
		} else {
			data.MulticastRoutingDistributed = types.BoolValue(false)
		}
	} else {
		data.MulticastRoutingDistributed = types.BoolNull()
	}
	for i := range data.MulticastRoutingVrfs {
		keys := [...]string{"name"}
		keyValues := [...]string{data.MulticastRoutingVrfs[i].Vrf.ValueString()}

		var r gjson.Result
		res.Get(prefix + "ip.Cisco-IOS-XE-multicast:multicast-routing.vrf").ForEach(
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
		if value := r.Get("name"); value.Exists() && !data.MulticastRoutingVrfs[i].Vrf.IsNull() {
			data.MulticastRoutingVrfs[i].Vrf = types.StringValue(value.String())
		} else {
			data.MulticastRoutingVrfs[i].Vrf = types.StringNull()
		}
		if value := r.Get("distributed"); !data.MulticastRoutingVrfs[i].Distributed.IsNull() {
			if value.Exists() {
				data.MulticastRoutingVrfs[i].Distributed = types.BoolValue(true)
			} else {
				data.MulticastRoutingVrfs[i].Distributed = types.BoolValue(false)
			}
		} else {
			data.MulticastRoutingVrfs[i].Distributed = types.BoolNull()
		}
	}
}

func (data *SystemData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "hostname"); value.Exists() {
		data.Hostname = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "ip.routing-conf.routing"); value.Exists() {
		data.IpRouting = types.BoolValue(value.Bool())
	} else {
		data.IpRouting = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ipv6.unicast-routing"); value.Exists() {
		data.Ipv6UnicastRouting = types.BoolValue(true)
	} else {
		data.Ipv6UnicastRouting = types.BoolValue(false)
	}
	if value := res.Get(prefix + "system.Cisco-IOS-XE-switch:mtu.size"); value.Exists() {
		data.Mtu = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "ip.source-route"); value.Exists() {
		data.IpSourceRoute = types.BoolValue(value.Bool())
	} else {
		data.IpSourceRoute = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ip.domain.lookup"); value.Exists() {
		data.IpDomainLookup = types.BoolValue(value.Bool())
	} else {
		data.IpDomainLookup = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ip.domain.name"); value.Exists() {
		data.IpDomainName = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "login.delay"); value.Exists() {
		data.LoginDelay = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "login.on-failure"); value.Exists() {
		data.LoginOnFailure = types.BoolValue(true)
	} else {
		data.LoginOnFailure = types.BoolValue(false)
	}
	if value := res.Get(prefix + "login.on-failure.log"); value.Exists() {
		data.LoginOnFailureLog = types.BoolValue(true)
	} else {
		data.LoginOnFailureLog = types.BoolValue(false)
	}
	if value := res.Get(prefix + "login.on-success"); value.Exists() {
		data.LoginOnSuccess = types.BoolValue(true)
	} else {
		data.LoginOnSuccess = types.BoolValue(false)
	}
	if value := res.Get(prefix + "login.on-success.log"); value.Exists() {
		data.LoginOnSuccessLog = types.BoolValue(true)
	} else {
		data.LoginOnSuccessLog = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ip.Cisco-IOS-XE-multicast:multicast-routing"); value.Exists() {
		data.MulticastRouting = types.BoolValue(true)
	} else {
		data.MulticastRouting = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ip.Cisco-IOS-XE-multicast:mcr-conf.multicast-routing"); value.Exists() {
		data.MulticastRoutingSwitch = types.BoolValue(true)
	} else {
		data.MulticastRoutingSwitch = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ip.Cisco-IOS-XE-multicast:multicast-routing.distributed"); value.Exists() {
		data.MulticastRoutingDistributed = types.BoolValue(true)
	} else {
		data.MulticastRoutingDistributed = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ip.Cisco-IOS-XE-multicast:multicast-routing.vrf"); value.Exists() {
		data.MulticastRoutingVrfs = make([]SystemMulticastRoutingVrfs, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemMulticastRoutingVrfs{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Vrf = types.StringValue(cValue.String())
			}
			if cValue := v.Get("distributed"); cValue.Exists() {
				item.Distributed = types.BoolValue(true)
			} else {
				item.Distributed = types.BoolValue(false)
			}
			data.MulticastRoutingVrfs = append(data.MulticastRoutingVrfs, item)
			return true
		})
	}
}

func (data *System) getDeletedListItems(ctx context.Context, state System) []string {
	deletedListItems := make([]string, 0)
	for i := range state.MulticastRoutingVrfs {
		stateKeyValues := [...]string{state.MulticastRoutingVrfs[i].Vrf.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.MulticastRoutingVrfs[i].Vrf.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.MulticastRoutingVrfs {
			found = true
			if state.MulticastRoutingVrfs[i].Vrf.ValueString() != data.MulticastRoutingVrfs[j].Vrf.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/ip/Cisco-IOS-XE-multicast:multicast-routing/vrf=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedListItems
}

func (data *System) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Ipv6UnicastRouting.IsNull() && !data.Ipv6UnicastRouting.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ipv6/unicast-routing", data.getPath()))
	}
	if !data.LoginOnFailure.IsNull() && !data.LoginOnFailure.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/login/on-failure", data.getPath()))
	}
	if !data.LoginOnFailureLog.IsNull() && !data.LoginOnFailureLog.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/login/on-failure/log", data.getPath()))
	}
	if !data.LoginOnSuccess.IsNull() && !data.LoginOnSuccess.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/login/on-success", data.getPath()))
	}
	if !data.LoginOnSuccessLog.IsNull() && !data.LoginOnSuccessLog.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/login/on-success/log", data.getPath()))
	}
	if !data.MulticastRouting.IsNull() && !data.MulticastRouting.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ip/Cisco-IOS-XE-multicast:multicast-routing", data.getPath()))
	}
	if !data.MulticastRoutingSwitch.IsNull() && !data.MulticastRoutingSwitch.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ip/Cisco-IOS-XE-multicast:mcr-conf/multicast-routing", data.getPath()))
	}
	if !data.MulticastRoutingDistributed.IsNull() && !data.MulticastRoutingDistributed.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ip/Cisco-IOS-XE-multicast:multicast-routing/distributed", data.getPath()))
	}

	for i := range data.MulticastRoutingVrfs {
		keyValues := [...]string{data.MulticastRoutingVrfs[i].Vrf.ValueString()}
		if !data.MulticastRoutingVrfs[i].Distributed.IsNull() && !data.MulticastRoutingVrfs[i].Distributed.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ip/Cisco-IOS-XE-multicast:multicast-routing/vrf=%v/distributed", data.getPath(), strings.Join(keyValues[:], ",")))
		}
	}
	return emptyLeafsDelete
}

func (data *System) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Hostname.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/hostname", data.getPath()))
	}
	if !data.IpRouting.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/routing-conf/routing", data.getPath()))
	}
	if !data.Ipv6UnicastRouting.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ipv6/unicast-routing", data.getPath()))
	}
	if !data.Mtu.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/system/Cisco-IOS-XE-switch:mtu/size", data.getPath()))
	}
	if !data.IpSourceRoute.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/source-route", data.getPath()))
	}
	if !data.IpDomainLookup.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/domain/lookup", data.getPath()))
	}
	if !data.IpDomainName.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/domain/name", data.getPath()))
	}
	if !data.LoginDelay.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/login/delay", data.getPath()))
	}
	if !data.LoginOnFailure.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/login/on-failure", data.getPath()))
	}
	if !data.LoginOnFailureLog.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/login/on-failure/log", data.getPath()))
	}
	if !data.LoginOnSuccess.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/login/on-success", data.getPath()))
	}
	if !data.LoginOnSuccessLog.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/login/on-success/log", data.getPath()))
	}
	if !data.MulticastRouting.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/Cisco-IOS-XE-multicast:multicast-routing", data.getPath()))
	}
	if !data.MulticastRoutingSwitch.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/Cisco-IOS-XE-multicast:mcr-conf/multicast-routing", data.getPath()))
	}
	if !data.MulticastRoutingDistributed.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/Cisco-IOS-XE-multicast:multicast-routing/distributed", data.getPath()))
	}
	for i := range data.MulticastRoutingVrfs {
		keyValues := [...]string{data.MulticastRoutingVrfs[i].Vrf.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/ip/Cisco-IOS-XE-multicast:multicast-routing/vrf=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	return deletePaths
}
