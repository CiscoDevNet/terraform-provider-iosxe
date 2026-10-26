// Copyright © 2025 Cisco Systems, Inc. and its affiliates.
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

package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/xmldot"
)

// Devices omit leaves equal to their YANG default from get-config. On a normal (non-import) read, a String or
// Int64 attribute whose state value equals the YANG default must survive the leaf being absent, while any other
// value must still be cleared so drift is detected.

func getConfigReply(inner string) xmldot.Result {
	return xmldot.Get("<rpc-reply><data>"+inner+"</data></rpc-reply>", "rpc-reply")
}

func TestUpdateFromBodyXMLKeepsYangDefaults(t *testing.T) {
	ctx := context.Background()
	reply := getConfigReply(`<wlan-cfg-data xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-wireless-wlan-cfg"><wlan-policies><wlan-policy>` +
		`<policy-profile-name>P1</policy-profile-name>` +
		`<guest-mm-db-export-entries><guest-mm-db-export-entry><ip>192.0.2.1</ip></guest-mm-db-export-entry></guest-mm-db-export-entries>` +
		`</wlan-policy></wlan-policies></wlan-cfg-data>`)

	tests := []struct {
		name               string
		vlan               types.String
		sessionTimeout     types.Int64
		anchorPriority     types.String
		wantVlan           types.String
		wantSessionTimeout types.Int64
		wantAnchorPriority types.String
	}{
		{
			name:               "state at YANG default is kept when leaf is absent",
			vlan:               types.StringValue("1"),
			sessionTimeout:     types.Int64Value(28800),
			anchorPriority:     types.StringValue("export-anchor-tertiary"),
			wantVlan:           types.StringValue("1"),
			wantSessionTimeout: types.Int64Value(28800),
			wantAnchorPriority: types.StringValue("export-anchor-tertiary"),
		},
		{
			name:               "state not at YANG default is cleared when leaf is absent",
			vlan:               types.StringValue("20"),
			sessionTimeout:     types.Int64Value(600),
			anchorPriority:     types.StringValue("export-anchor-primary"),
			wantVlan:           types.StringNull(),
			wantSessionTimeout: types.Int64Null(),
			wantAnchorPriority: types.StringNull(),
		},
		{
			name:               "null state stays null",
			vlan:               types.StringNull(),
			sessionTimeout:     types.Int64Null(),
			anchorPriority:     types.StringNull(),
			wantVlan:           types.StringNull(),
			wantSessionTimeout: types.Int64Null(),
			wantAnchorPriority: types.StringNull(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := WirelessPolicyProfile{
				PolicyProfileName: types.StringValue("P1"),
				Vlan:              tt.vlan,
				SessionTimeout:    tt.sessionTimeout,
				MobilityAnchors: []WirelessPolicyProfileMobilityAnchors{
					{Ip: types.StringValue("192.0.2.1"), Priority: tt.anchorPriority},
				},
			}
			data.updateFromBodyXML(ctx, reply)
			if !data.Vlan.Equal(tt.wantVlan) {
				t.Errorf("Vlan = %v, want %v", data.Vlan, tt.wantVlan)
			}
			if !data.SessionTimeout.Equal(tt.wantSessionTimeout) {
				t.Errorf("SessionTimeout = %v, want %v", data.SessionTimeout, tt.wantSessionTimeout)
			}
			if !data.MobilityAnchors[0].Priority.Equal(tt.wantAnchorPriority) {
				t.Errorf("MobilityAnchors[0].Priority = %v, want %v", data.MobilityAnchors[0].Priority, tt.wantAnchorPriority)
			}
		})
	}
}

func TestUpdateFromBodyXMLReadsPresentLeaves(t *testing.T) {
	ctx := context.Background()
	reply := getConfigReply(`<wlan-cfg-data xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-wireless-wlan-cfg"><wlan-policies><wlan-policy>` +
		`<policy-profile-name>P1</policy-profile-name><interface-name>30</interface-name>` +
		`<wlan-timeout><session-timeout>1800</session-timeout></wlan-timeout>` +
		`</wlan-policy></wlan-policies></wlan-cfg-data>`)

	// Changed out of band: the device value wins even though state held the YANG default.
	data := WirelessPolicyProfile{
		PolicyProfileName: types.StringValue("P1"),
		Vlan:              types.StringValue("1"),
		SessionTimeout:    types.Int64Value(28800),
	}
	data.updateFromBodyXML(ctx, reply)
	if got, want := data.Vlan, types.StringValue("30"); !got.Equal(want) {
		t.Errorf("Vlan = %v, want %v", got, want)
	}
	if got, want := data.SessionTimeout, types.Int64Value(1800); !got.Equal(want) {
		t.Errorf("SessionTimeout = %v, want %v", got, want)
	}
}

func TestUpdateFromBodyXMLClearsMissingListItem(t *testing.T) {
	ctx := context.Background()
	// The anchor list entry no longer exists on the device: its attributes must be cleared even when at default.
	reply := getConfigReply(`<wlan-cfg-data xmlns="http://cisco.com/ns/yang/Cisco-IOS-XE-wireless-wlan-cfg"><wlan-policies><wlan-policy>` +
		`<policy-profile-name>P1</policy-profile-name>` +
		`</wlan-policy></wlan-policies></wlan-cfg-data>`)
	data := WirelessPolicyProfile{
		PolicyProfileName: types.StringValue("P1"),
		MobilityAnchors: []WirelessPolicyProfileMobilityAnchors{
			{Ip: types.StringValue("192.0.2.1"), Priority: types.StringValue("export-anchor-tertiary")},
		},
	}
	data.updateFromBodyXML(ctx, reply)
	if !data.MobilityAnchors[0].Priority.IsNull() {
		t.Errorf("MobilityAnchors[0].Priority = %v, want null", data.MobilityAnchors[0].Priority)
	}
}
