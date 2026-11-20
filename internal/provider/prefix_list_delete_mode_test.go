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

package provider

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestPrefixListDeleteModeSchema(t *testing.T) {
	response := resource.SchemaResponse{}
	NewPrefixListResource().Schema(context.Background(), resource.SchemaRequest{}, &response)
	attribute, ok := response.Schema.Attributes["delete_mode"]
	if !ok || !attribute.IsOptional() {
		t.Fatal("shared prefix-list container requires an optional delete_mode")
	}
}

func TestPrefixListDeletePathsAreKeyed(t *testing.T) {
	model := PrefixList{Prefixes: []PrefixListPrefixes{
		{Name: types.StringValue("EVPN-Campus-TO-FUSION"), Seq: types.Int64Value(10)},
		{Name: types.StringValue("EVPN-IoT-TO-FUSION"), Seq: types.Int64Value(10)},
	}, PrefixListDescription: []PrefixListPrefixListDescription{
		{Name: types.StringValue("EVPN-Campus-TO-FUSION"), Description: types.StringValue("owned description")},
	}}
	body := model.addDeletePathsXML(context.Background(), "")
	if strings.Count(body, `operation="remove"`) != 3 ||
		!strings.Contains(body, "EVPN-Campus-TO-FUSION") ||
		!strings.Contains(body, "EVPN-IoT-TO-FUSION") ||
		!strings.Contains(body, "<no>10</no>") ||
		!strings.Contains(body, "prefix-list-description") ||
		strings.Contains(body, "owned description") ||
		strings.Contains(body, `<prefix-lists nc:operation="remove"`) {
		t.Fatalf("delete must target only keyed managed entries: %s", body)
	}
}
