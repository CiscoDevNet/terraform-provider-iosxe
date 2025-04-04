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

type EVPNInstance struct {
	Device                           types.String `tfsdk:"device"`
	Id                               types.String `tfsdk:"id"`
	EvpnInstanceNum                  types.Int64  `tfsdk:"evpn_instance_num"`
	VlanBasedReplicationTypeIngress  types.Bool   `tfsdk:"vlan_based_replication_type_ingress"`
	VlanBasedReplicationTypeStatic   types.Bool   `tfsdk:"vlan_based_replication_type_static"`
	VlanBasedReplicationTypeP2mp     types.Bool   `tfsdk:"vlan_based_replication_type_p2mp"`
	VlanBasedReplicationTypeMp2mp    types.Bool   `tfsdk:"vlan_based_replication_type_mp2mp"`
	VlanBasedEncapsulation           types.String `tfsdk:"vlan_based_encapsulation"`
	VlanBasedAutoRouteTarget         types.Bool   `tfsdk:"vlan_based_auto_route_target"`
	VlanBasedRd                      types.String `tfsdk:"vlan_based_rd"`
	VlanBasedRouteTarget             types.String `tfsdk:"vlan_based_route_target"`
	VlanBasedRouteTargetBoth         types.String `tfsdk:"vlan_based_route_target_both"`
	VlanBasedRouteTargetImport       types.String `tfsdk:"vlan_based_route_target_import"`
	VlanBasedRouteTargetExport       types.String `tfsdk:"vlan_based_route_target_export"`
	VlanBasedIpLocalLearningDisable  types.Bool   `tfsdk:"vlan_based_ip_local_learning_disable"`
	VlanBasedIpLocalLearningEnable   types.Bool   `tfsdk:"vlan_based_ip_local_learning_enable"`
	VlanBasedDefaultGatewayAdvertise types.String `tfsdk:"vlan_based_default_gateway_advertise"`
	VlanBasedReOriginateRouteType5   types.Bool   `tfsdk:"vlan_based_re_originate_route_type5"`
}

type EVPNInstanceData struct {
	Device                           types.String `tfsdk:"device"`
	Id                               types.String `tfsdk:"id"`
	EvpnInstanceNum                  types.Int64  `tfsdk:"evpn_instance_num"`
	VlanBasedReplicationTypeIngress  types.Bool   `tfsdk:"vlan_based_replication_type_ingress"`
	VlanBasedReplicationTypeStatic   types.Bool   `tfsdk:"vlan_based_replication_type_static"`
	VlanBasedReplicationTypeP2mp     types.Bool   `tfsdk:"vlan_based_replication_type_p2mp"`
	VlanBasedReplicationTypeMp2mp    types.Bool   `tfsdk:"vlan_based_replication_type_mp2mp"`
	VlanBasedEncapsulation           types.String `tfsdk:"vlan_based_encapsulation"`
	VlanBasedAutoRouteTarget         types.Bool   `tfsdk:"vlan_based_auto_route_target"`
	VlanBasedRd                      types.String `tfsdk:"vlan_based_rd"`
	VlanBasedRouteTarget             types.String `tfsdk:"vlan_based_route_target"`
	VlanBasedRouteTargetBoth         types.String `tfsdk:"vlan_based_route_target_both"`
	VlanBasedRouteTargetImport       types.String `tfsdk:"vlan_based_route_target_import"`
	VlanBasedRouteTargetExport       types.String `tfsdk:"vlan_based_route_target_export"`
	VlanBasedIpLocalLearningDisable  types.Bool   `tfsdk:"vlan_based_ip_local_learning_disable"`
	VlanBasedIpLocalLearningEnable   types.Bool   `tfsdk:"vlan_based_ip_local_learning_enable"`
	VlanBasedDefaultGatewayAdvertise types.String `tfsdk:"vlan_based_default_gateway_advertise"`
	VlanBasedReOriginateRouteType5   types.Bool   `tfsdk:"vlan_based_re_originate_route_type5"`
}

func (data EVPNInstance) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/l2vpn/Cisco-IOS-XE-l2vpn:evpn_cont/evpn-instance/evpn/instance/instance=%v", url.QueryEscape(fmt.Sprintf("%v", data.EvpnInstanceNum.ValueInt64())))
}

func (data EVPNInstanceData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/l2vpn/Cisco-IOS-XE-l2vpn:evpn_cont/evpn-instance/evpn/instance/instance=%v", url.QueryEscape(fmt.Sprintf("%v", data.EvpnInstanceNum.ValueInt64())))
}

// if last path element has a key -> remove it
func (data EVPNInstance) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data EVPNInstance) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.EvpnInstanceNum.IsNull() && !data.EvpnInstanceNum.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"evpn-instance-num", strconv.FormatInt(data.EvpnInstanceNum.ValueInt64(), 10))
	}
	if !data.VlanBasedReplicationTypeIngress.IsNull() && !data.VlanBasedReplicationTypeIngress.IsUnknown() {
		if data.VlanBasedReplicationTypeIngress.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vlan-based.replication-type.ingress", map[string]string{})
		}
	}
	if !data.VlanBasedReplicationTypeStatic.IsNull() && !data.VlanBasedReplicationTypeStatic.IsUnknown() {
		if data.VlanBasedReplicationTypeStatic.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vlan-based.replication-type.static", map[string]string{})
		}
	}
	if !data.VlanBasedReplicationTypeP2mp.IsNull() && !data.VlanBasedReplicationTypeP2mp.IsUnknown() {
		if data.VlanBasedReplicationTypeP2mp.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vlan-based.replication-type.p2mp", map[string]string{})
		}
	}
	if !data.VlanBasedReplicationTypeMp2mp.IsNull() && !data.VlanBasedReplicationTypeMp2mp.IsUnknown() {
		if data.VlanBasedReplicationTypeMp2mp.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vlan-based.replication-type.mp2mp", map[string]string{})
		}
	}
	if !data.VlanBasedEncapsulation.IsNull() && !data.VlanBasedEncapsulation.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vlan-based.encapsulation", data.VlanBasedEncapsulation.ValueString())
	}
	if !data.VlanBasedAutoRouteTarget.IsNull() && !data.VlanBasedAutoRouteTarget.IsUnknown() {
		if data.VlanBasedAutoRouteTarget.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vlan-based.auto-route-target_cont.auto-route-target", map[string]string{})
		}
	}
	if !data.VlanBasedRd.IsNull() && !data.VlanBasedRd.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vlan-based.rd.rd-value", data.VlanBasedRd.ValueString())
	}
	if !data.VlanBasedRouteTarget.IsNull() && !data.VlanBasedRouteTarget.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vlan-based.route-target.rt-value", data.VlanBasedRouteTarget.ValueString())
	}
	if !data.VlanBasedRouteTargetBoth.IsNull() && !data.VlanBasedRouteTargetBoth.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vlan-based.route-target.both.rt-value", data.VlanBasedRouteTargetBoth.ValueString())
	}
	if !data.VlanBasedRouteTargetImport.IsNull() && !data.VlanBasedRouteTargetImport.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vlan-based.route-target.import.rt-value", data.VlanBasedRouteTargetImport.ValueString())
	}
	if !data.VlanBasedRouteTargetExport.IsNull() && !data.VlanBasedRouteTargetExport.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vlan-based.route-target.export.rt-value", data.VlanBasedRouteTargetExport.ValueString())
	}
	if !data.VlanBasedIpLocalLearningDisable.IsNull() && !data.VlanBasedIpLocalLearningDisable.IsUnknown() {
		if data.VlanBasedIpLocalLearningDisable.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vlan-based.ip.local-learning.disable", map[string]string{})
		}
	}
	if !data.VlanBasedIpLocalLearningEnable.IsNull() && !data.VlanBasedIpLocalLearningEnable.IsUnknown() {
		if data.VlanBasedIpLocalLearningEnable.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vlan-based.ip.local-learning.enable", map[string]string{})
		}
	}
	if !data.VlanBasedDefaultGatewayAdvertise.IsNull() && !data.VlanBasedDefaultGatewayAdvertise.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vlan-based.default-gateway.advertise", data.VlanBasedDefaultGatewayAdvertise.ValueString())
	}
	if !data.VlanBasedReOriginateRouteType5.IsNull() && !data.VlanBasedReOriginateRouteType5.IsUnknown() {
		if data.VlanBasedReOriginateRouteType5.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vlan-based.re-originate.route-type5", map[string]string{})
		}
	}
	return body
}

func (data *EVPNInstance) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "evpn-instance-num"); value.Exists() && !data.EvpnInstanceNum.IsNull() {
		data.EvpnInstanceNum = types.Int64Value(value.Int())
	} else {
		data.EvpnInstanceNum = types.Int64Null()
	}
	if value := res.Get(prefix + "vlan-based.replication-type.ingress"); !data.VlanBasedReplicationTypeIngress.IsNull() {
		if value.Exists() {
			data.VlanBasedReplicationTypeIngress = types.BoolValue(true)
		} else {
			data.VlanBasedReplicationTypeIngress = types.BoolValue(false)
		}
	} else {
		data.VlanBasedReplicationTypeIngress = types.BoolNull()
	}
	if value := res.Get(prefix + "vlan-based.replication-type.static"); !data.VlanBasedReplicationTypeStatic.IsNull() {
		if value.Exists() {
			data.VlanBasedReplicationTypeStatic = types.BoolValue(true)
		} else {
			data.VlanBasedReplicationTypeStatic = types.BoolValue(false)
		}
	} else {
		data.VlanBasedReplicationTypeStatic = types.BoolNull()
	}
	if value := res.Get(prefix + "vlan-based.replication-type.p2mp"); !data.VlanBasedReplicationTypeP2mp.IsNull() {
		if value.Exists() {
			data.VlanBasedReplicationTypeP2mp = types.BoolValue(true)
		} else {
			data.VlanBasedReplicationTypeP2mp = types.BoolValue(false)
		}
	} else {
		data.VlanBasedReplicationTypeP2mp = types.BoolNull()
	}
	if value := res.Get(prefix + "vlan-based.replication-type.mp2mp"); !data.VlanBasedReplicationTypeMp2mp.IsNull() {
		if value.Exists() {
			data.VlanBasedReplicationTypeMp2mp = types.BoolValue(true)
		} else {
			data.VlanBasedReplicationTypeMp2mp = types.BoolValue(false)
		}
	} else {
		data.VlanBasedReplicationTypeMp2mp = types.BoolNull()
	}
	if value := res.Get(prefix + "vlan-based.encapsulation"); value.Exists() && !data.VlanBasedEncapsulation.IsNull() {
		data.VlanBasedEncapsulation = types.StringValue(value.String())
	} else {
		data.VlanBasedEncapsulation = types.StringNull()
	}
	if value := res.Get(prefix + "vlan-based.auto-route-target_cont.auto-route-target"); !data.VlanBasedAutoRouteTarget.IsNull() {
		if value.Exists() {
			data.VlanBasedAutoRouteTarget = types.BoolValue(true)
		} else {
			data.VlanBasedAutoRouteTarget = types.BoolValue(false)
		}
	} else {
		data.VlanBasedAutoRouteTarget = types.BoolNull()
	}
	if value := res.Get(prefix + "vlan-based.rd.rd-value"); value.Exists() && !data.VlanBasedRd.IsNull() {
		data.VlanBasedRd = types.StringValue(value.String())
	} else {
		data.VlanBasedRd = types.StringNull()
	}
	if value := res.Get(prefix + "vlan-based.route-target.rt-value"); value.Exists() && !data.VlanBasedRouteTarget.IsNull() {
		data.VlanBasedRouteTarget = types.StringValue(value.String())
	} else {
		data.VlanBasedRouteTarget = types.StringNull()
	}
	if value := res.Get(prefix + "vlan-based.route-target.both.rt-value"); value.Exists() && !data.VlanBasedRouteTargetBoth.IsNull() {
		data.VlanBasedRouteTargetBoth = types.StringValue(value.String())
	} else {
		data.VlanBasedRouteTargetBoth = types.StringNull()
	}
	if value := res.Get(prefix + "vlan-based.route-target.import.rt-value"); value.Exists() && !data.VlanBasedRouteTargetImport.IsNull() {
		data.VlanBasedRouteTargetImport = types.StringValue(value.String())
	} else {
		data.VlanBasedRouteTargetImport = types.StringNull()
	}
	if value := res.Get(prefix + "vlan-based.route-target.export.rt-value"); value.Exists() && !data.VlanBasedRouteTargetExport.IsNull() {
		data.VlanBasedRouteTargetExport = types.StringValue(value.String())
	} else {
		data.VlanBasedRouteTargetExport = types.StringNull()
	}
	if value := res.Get(prefix + "vlan-based.ip.local-learning.disable"); !data.VlanBasedIpLocalLearningDisable.IsNull() {
		if value.Exists() {
			data.VlanBasedIpLocalLearningDisable = types.BoolValue(true)
		} else {
			data.VlanBasedIpLocalLearningDisable = types.BoolValue(false)
		}
	} else {
		data.VlanBasedIpLocalLearningDisable = types.BoolNull()
	}
	if value := res.Get(prefix + "vlan-based.ip.local-learning.enable"); !data.VlanBasedIpLocalLearningEnable.IsNull() {
		if value.Exists() {
			data.VlanBasedIpLocalLearningEnable = types.BoolValue(true)
		} else {
			data.VlanBasedIpLocalLearningEnable = types.BoolValue(false)
		}
	} else {
		data.VlanBasedIpLocalLearningEnable = types.BoolNull()
	}
	if value := res.Get(prefix + "vlan-based.default-gateway.advertise"); value.Exists() && !data.VlanBasedDefaultGatewayAdvertise.IsNull() {
		data.VlanBasedDefaultGatewayAdvertise = types.StringValue(value.String())
	} else {
		data.VlanBasedDefaultGatewayAdvertise = types.StringNull()
	}
	if value := res.Get(prefix + "vlan-based.re-originate.route-type5"); !data.VlanBasedReOriginateRouteType5.IsNull() {
		if value.Exists() {
			data.VlanBasedReOriginateRouteType5 = types.BoolValue(true)
		} else {
			data.VlanBasedReOriginateRouteType5 = types.BoolValue(false)
		}
	} else {
		data.VlanBasedReOriginateRouteType5 = types.BoolNull()
	}
}

func (data *EVPNInstance) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "vlan-based.replication-type.ingress"); value.Exists() {
		data.VlanBasedReplicationTypeIngress = types.BoolValue(true)
	} else {
		data.VlanBasedReplicationTypeIngress = types.BoolValue(false)
	}
	if value := res.Get(prefix + "vlan-based.replication-type.static"); value.Exists() {
		data.VlanBasedReplicationTypeStatic = types.BoolValue(true)
	} else {
		data.VlanBasedReplicationTypeStatic = types.BoolValue(false)
	}
	if value := res.Get(prefix + "vlan-based.replication-type.p2mp"); value.Exists() {
		data.VlanBasedReplicationTypeP2mp = types.BoolValue(true)
	} else {
		data.VlanBasedReplicationTypeP2mp = types.BoolValue(false)
	}
	if value := res.Get(prefix + "vlan-based.replication-type.mp2mp"); value.Exists() {
		data.VlanBasedReplicationTypeMp2mp = types.BoolValue(true)
	} else {
		data.VlanBasedReplicationTypeMp2mp = types.BoolValue(false)
	}
	if value := res.Get(prefix + "vlan-based.encapsulation"); value.Exists() {
		data.VlanBasedEncapsulation = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "vlan-based.auto-route-target_cont.auto-route-target"); value.Exists() {
		data.VlanBasedAutoRouteTarget = types.BoolValue(true)
	} else {
		data.VlanBasedAutoRouteTarget = types.BoolValue(false)
	}
	if value := res.Get(prefix + "vlan-based.rd.rd-value"); value.Exists() {
		data.VlanBasedRd = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "vlan-based.route-target.rt-value"); value.Exists() {
		data.VlanBasedRouteTarget = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "vlan-based.route-target.both.rt-value"); value.Exists() {
		data.VlanBasedRouteTargetBoth = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "vlan-based.route-target.import.rt-value"); value.Exists() {
		data.VlanBasedRouteTargetImport = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "vlan-based.route-target.export.rt-value"); value.Exists() {
		data.VlanBasedRouteTargetExport = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "vlan-based.ip.local-learning.disable"); value.Exists() {
		data.VlanBasedIpLocalLearningDisable = types.BoolValue(true)
	} else {
		data.VlanBasedIpLocalLearningDisable = types.BoolValue(false)
	}
	if value := res.Get(prefix + "vlan-based.ip.local-learning.enable"); value.Exists() {
		data.VlanBasedIpLocalLearningEnable = types.BoolValue(true)
	} else {
		data.VlanBasedIpLocalLearningEnable = types.BoolValue(false)
	}
	if value := res.Get(prefix + "vlan-based.default-gateway.advertise"); value.Exists() {
		data.VlanBasedDefaultGatewayAdvertise = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "vlan-based.re-originate.route-type5"); value.Exists() {
		data.VlanBasedReOriginateRouteType5 = types.BoolValue(true)
	} else {
		data.VlanBasedReOriginateRouteType5 = types.BoolValue(false)
	}
}

func (data *EVPNInstanceData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "vlan-based.replication-type.ingress"); value.Exists() {
		data.VlanBasedReplicationTypeIngress = types.BoolValue(true)
	} else {
		data.VlanBasedReplicationTypeIngress = types.BoolValue(false)
	}
	if value := res.Get(prefix + "vlan-based.replication-type.static"); value.Exists() {
		data.VlanBasedReplicationTypeStatic = types.BoolValue(true)
	} else {
		data.VlanBasedReplicationTypeStatic = types.BoolValue(false)
	}
	if value := res.Get(prefix + "vlan-based.replication-type.p2mp"); value.Exists() {
		data.VlanBasedReplicationTypeP2mp = types.BoolValue(true)
	} else {
		data.VlanBasedReplicationTypeP2mp = types.BoolValue(false)
	}
	if value := res.Get(prefix + "vlan-based.replication-type.mp2mp"); value.Exists() {
		data.VlanBasedReplicationTypeMp2mp = types.BoolValue(true)
	} else {
		data.VlanBasedReplicationTypeMp2mp = types.BoolValue(false)
	}
	if value := res.Get(prefix + "vlan-based.encapsulation"); value.Exists() {
		data.VlanBasedEncapsulation = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "vlan-based.auto-route-target_cont.auto-route-target"); value.Exists() {
		data.VlanBasedAutoRouteTarget = types.BoolValue(true)
	} else {
		data.VlanBasedAutoRouteTarget = types.BoolValue(false)
	}
	if value := res.Get(prefix + "vlan-based.rd.rd-value"); value.Exists() {
		data.VlanBasedRd = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "vlan-based.route-target.rt-value"); value.Exists() {
		data.VlanBasedRouteTarget = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "vlan-based.route-target.both.rt-value"); value.Exists() {
		data.VlanBasedRouteTargetBoth = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "vlan-based.route-target.import.rt-value"); value.Exists() {
		data.VlanBasedRouteTargetImport = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "vlan-based.route-target.export.rt-value"); value.Exists() {
		data.VlanBasedRouteTargetExport = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "vlan-based.ip.local-learning.disable"); value.Exists() {
		data.VlanBasedIpLocalLearningDisable = types.BoolValue(true)
	} else {
		data.VlanBasedIpLocalLearningDisable = types.BoolValue(false)
	}
	if value := res.Get(prefix + "vlan-based.ip.local-learning.enable"); value.Exists() {
		data.VlanBasedIpLocalLearningEnable = types.BoolValue(true)
	} else {
		data.VlanBasedIpLocalLearningEnable = types.BoolValue(false)
	}
	if value := res.Get(prefix + "vlan-based.default-gateway.advertise"); value.Exists() {
		data.VlanBasedDefaultGatewayAdvertise = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "vlan-based.re-originate.route-type5"); value.Exists() {
		data.VlanBasedReOriginateRouteType5 = types.BoolValue(true)
	} else {
		data.VlanBasedReOriginateRouteType5 = types.BoolValue(false)
	}
}

func (data *EVPNInstance) getDeletedItems(ctx context.Context, state EVPNInstance) []string {
	deletedItems := make([]string, 0)
	if !state.VlanBasedReplicationTypeIngress.IsNull() && data.VlanBasedReplicationTypeIngress.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/vlan-based/replication-type/ingress", state.getPath()))
	}
	if !state.VlanBasedReplicationTypeStatic.IsNull() && data.VlanBasedReplicationTypeStatic.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/vlan-based/replication-type/static", state.getPath()))
	}
	if !state.VlanBasedReplicationTypeP2mp.IsNull() && data.VlanBasedReplicationTypeP2mp.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/vlan-based/replication-type/p2mp", state.getPath()))
	}
	if !state.VlanBasedReplicationTypeMp2mp.IsNull() && data.VlanBasedReplicationTypeMp2mp.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/vlan-based/replication-type/mp2mp", state.getPath()))
	}
	if !state.VlanBasedEncapsulation.IsNull() && data.VlanBasedEncapsulation.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/vlan-based/encapsulation", state.getPath()))
	}
	if !state.VlanBasedAutoRouteTarget.IsNull() && data.VlanBasedAutoRouteTarget.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/vlan-based/auto-route-target_cont/auto-route-target", state.getPath()))
	}
	if !state.VlanBasedRd.IsNull() && data.VlanBasedRd.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/vlan-based/rd/rd-value", state.getPath()))
	}
	if !state.VlanBasedRouteTarget.IsNull() && data.VlanBasedRouteTarget.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/vlan-based/route-target/rt-value", state.getPath()))
	}
	if !state.VlanBasedRouteTargetBoth.IsNull() && data.VlanBasedRouteTargetBoth.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/vlan-based/route-target/both/rt-value", state.getPath()))
	}
	if !state.VlanBasedRouteTargetImport.IsNull() && data.VlanBasedRouteTargetImport.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/vlan-based/route-target/import/rt-value", state.getPath()))
	}
	if !state.VlanBasedRouteTargetExport.IsNull() && data.VlanBasedRouteTargetExport.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/vlan-based/route-target/export/rt-value", state.getPath()))
	}
	if !state.VlanBasedIpLocalLearningDisable.IsNull() && data.VlanBasedIpLocalLearningDisable.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/vlan-based/ip/local-learning/disable", state.getPath()))
	}
	if !state.VlanBasedIpLocalLearningEnable.IsNull() && data.VlanBasedIpLocalLearningEnable.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/vlan-based/ip/local-learning/enable", state.getPath()))
	}
	if !state.VlanBasedDefaultGatewayAdvertise.IsNull() && data.VlanBasedDefaultGatewayAdvertise.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/vlan-based/default-gateway/advertise", state.getPath()))
	}
	if !state.VlanBasedReOriginateRouteType5.IsNull() && data.VlanBasedReOriginateRouteType5.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/vlan-based/re-originate/route-type5", state.getPath()))
	}
	return deletedItems
}

func (data *EVPNInstance) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.VlanBasedReplicationTypeIngress.IsNull() && !data.VlanBasedReplicationTypeIngress.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/vlan-based/replication-type/ingress", data.getPath()))
	}
	if !data.VlanBasedReplicationTypeStatic.IsNull() && !data.VlanBasedReplicationTypeStatic.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/vlan-based/replication-type/static", data.getPath()))
	}
	if !data.VlanBasedReplicationTypeP2mp.IsNull() && !data.VlanBasedReplicationTypeP2mp.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/vlan-based/replication-type/p2mp", data.getPath()))
	}
	if !data.VlanBasedReplicationTypeMp2mp.IsNull() && !data.VlanBasedReplicationTypeMp2mp.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/vlan-based/replication-type/mp2mp", data.getPath()))
	}
	if !data.VlanBasedAutoRouteTarget.IsNull() && !data.VlanBasedAutoRouteTarget.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/vlan-based/auto-route-target_cont/auto-route-target", data.getPath()))
	}
	if !data.VlanBasedIpLocalLearningDisable.IsNull() && !data.VlanBasedIpLocalLearningDisable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/vlan-based/ip/local-learning/disable", data.getPath()))
	}
	if !data.VlanBasedIpLocalLearningEnable.IsNull() && !data.VlanBasedIpLocalLearningEnable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/vlan-based/ip/local-learning/enable", data.getPath()))
	}
	if !data.VlanBasedReOriginateRouteType5.IsNull() && !data.VlanBasedReOriginateRouteType5.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/vlan-based/re-originate/route-type5", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *EVPNInstance) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.VlanBasedReplicationTypeIngress.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/vlan-based/replication-type/ingress", data.getPath()))
	}
	if !data.VlanBasedReplicationTypeStatic.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/vlan-based/replication-type/static", data.getPath()))
	}
	if !data.VlanBasedReplicationTypeP2mp.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/vlan-based/replication-type/p2mp", data.getPath()))
	}
	if !data.VlanBasedReplicationTypeMp2mp.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/vlan-based/replication-type/mp2mp", data.getPath()))
	}
	if !data.VlanBasedEncapsulation.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/vlan-based/encapsulation", data.getPath()))
	}
	if !data.VlanBasedAutoRouteTarget.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/vlan-based/auto-route-target_cont/auto-route-target", data.getPath()))
	}
	if !data.VlanBasedRd.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/vlan-based/rd/rd-value", data.getPath()))
	}
	if !data.VlanBasedRouteTarget.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/vlan-based/route-target/rt-value", data.getPath()))
	}
	if !data.VlanBasedRouteTargetBoth.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/vlan-based/route-target/both/rt-value", data.getPath()))
	}
	if !data.VlanBasedRouteTargetImport.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/vlan-based/route-target/import/rt-value", data.getPath()))
	}
	if !data.VlanBasedRouteTargetExport.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/vlan-based/route-target/export/rt-value", data.getPath()))
	}
	if !data.VlanBasedIpLocalLearningDisable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/vlan-based/ip/local-learning/disable", data.getPath()))
	}
	if !data.VlanBasedIpLocalLearningEnable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/vlan-based/ip/local-learning/enable", data.getPath()))
	}
	if !data.VlanBasedDefaultGatewayAdvertise.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/vlan-based/default-gateway/advertise", data.getPath()))
	}
	if !data.VlanBasedReOriginateRouteType5.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/vlan-based/re-originate/route-type5", data.getPath()))
	}
	return deletePaths
}

func (data *EVPNInstance) getIdsFromPath() {
	reString := strings.ReplaceAll("Cisco-IOS-XE-native:native/l2vpn/Cisco-IOS-XE-l2vpn:evpn_cont/evpn-instance/evpn/instance/instance=%v", "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Id.ValueString())
	data.EvpnInstanceNum = types.Int64Value(helpers.Must(strconv.ParseInt(matches[1], 10, 0)))
}
