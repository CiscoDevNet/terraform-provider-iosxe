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

// Hand-written support for iosxe_wireless_ap_tag_source_priority (not generated).
//
// The C9800 only lets priorities 2 and 3 of the AP tag source table be changed; priorities 0
// (static), 1 (location) and 4 (default) are fixed. It also validates every single edit-config
// on its own and rejects one that, even transiently, assigns the same source to two priorities
// ("Tag Source is already assigned to Priority N"), whether the swap is sent as a merge, a
// container replace or a delete plus create in one payload. Swapping the sources of priorities
// 2 and 3 therefore takes two edit-config operations: first remove the priorities whose source
// changes, then write the desired priorities.

import (
	"context"
	"fmt"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-netconf"
	"github.com/netascode/xmldot"
)

// wirelessAPTagSourceFactoryPriorities is the factory order of the configurable priorities. It is
// restored when the resource is destroyed.
var wirelessAPTagSourceFactoryPriorities = map[int64]string{
	2: "tag-source-filter",
	3: "tag-source-ap-pnp",
}

// wirelessAPTagSourceConfigurablePriorities lists the configurable priorities in write order.
var wirelessAPTagSourceConfigurablePriorities = []int64{2, 3}

// onlyConfigurableTagSourcePriorities drops the fixed priorities (0, 1 and 4) read from the device.
func onlyConfigurableTagSourcePriorities(items []WirelessAPTagSourcePriorityPriorities) []WirelessAPTagSourcePriorityPriorities {
	if items == nil {
		return nil
	}
	result := make([]WirelessAPTagSourcePriorityPriorities, 0, len(items))
	for _, item := range items {
		if _, ok := wirelessAPTagSourceFactoryPriorities[item.Priority.ValueInt64()]; ok {
			result = append(result, item)
		}
	}
	return result
}

// tagSourcePriorityTarget returns the complete table of the configurable priorities to write. The
// device requires both priorities to hold a source and every source to be used once, so a
// priority that is not configured gets the source left over by the configured ones, which is
// its factory source when the configured priorities keep the factory order.
func tagSourcePriorityTarget(items []WirelessAPTagSourcePriorityPriorities) map[int64]string {
	target := make(map[int64]string, len(wirelessAPTagSourceConfigurablePriorities))
	used := make(map[string]bool, len(wirelessAPTagSourceConfigurablePriorities))
	for _, item := range items {
		if item.Priority.IsNull() || item.Priority.IsUnknown() || item.Source.IsNull() || item.Source.IsUnknown() {
			continue
		}
		target[item.Priority.ValueInt64()] = item.Source.ValueString()
		used[item.Source.ValueString()] = true
	}
	for _, priority := range wirelessAPTagSourceConfigurablePriorities {
		if _, ok := target[priority]; ok {
			continue
		}
		candidates := []string{wirelessAPTagSourceFactoryPriorities[priority]}
		for _, p := range wirelessAPTagSourceConfigurablePriorities {
			candidates = append(candidates, wirelessAPTagSourceFactoryPriorities[p])
		}
		for _, source := range candidates {
			if !used[source] {
				target[priority] = source
				used[source] = true
				break
			}
		}
	}
	return target
}

// tagSourcePrioritiesToRemove returns the configurable priorities whose current source differs
// from the target. They are removed in a separate edit-config before the target is written.
func tagSourcePrioritiesToRemove(current, target map[int64]string) []int64 {
	var result []int64
	for _, priority := range wirelessAPTagSourceConfigurablePriorities {
		if source, ok := current[priority]; ok && source != target[priority] {
			result = append(result, priority)
		}
	}
	return result
}

// readTagSourcePriorities reads the current priority -> source table from the device.
func (data WirelessAPTagSourcePriority) readTagSourcePriorities(ctx context.Context, device *IosxeProviderDataDevice) (map[int64]string, error) {
	res, err := device.NetconfClient.GetConfig(ctx, "running", helpers.GetXpathFilter(data.getXPath()))
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve object (%s), got error: %s", data.getPath(), err)
	}
	current := make(map[int64]string)
	helpers.GetFromXPath(res.Res, "data"+data.getXPath()+"/tag-source-priority-config").ForEach(
		func(_ int, v xmldot.Result) bool {
			current[v.Get("priority").Int()] = v.Get("tag-src").String()
			return true
		},
	)
	return current, nil
}

// applyTagSourcePriorities writes the target table to the device. Priorities whose source changes
// are removed in a first edit-config, so that no single edit assigns a source to two priorities.
func (data WirelessAPTagSourcePriority) applyTagSourcePriorities(ctx context.Context, device *IosxeProviderDataDevice, target map[int64]string) error {
	current, err := data.readTagSourcePriorities(ctx, device)
	if err != nil {
		return err
	}

	remove := tagSourcePrioritiesToRemove(current, target)
	if len(remove) > 0 {
		b := netconf.Body{}
		for _, priority := range remove {
			b = helpers.RemoveFromXPath(b, fmt.Sprintf(data.getXPath()+"/tag-source-priority-config[priority='%d']", priority))
		}
		tflog.Debug(ctx, fmt.Sprintf("%s: Removing tag source priorities %v before writing %v", data.getPath(), remove, target))
		if err := helpers.EditConfig(ctx, device.NetconfClient, b.Res(), device.AutoCommit); err != nil {
			return err
		}
	}

	items := make([]WirelessAPTagSourcePriorityPriorities, 0, len(target))
	for _, priority := range wirelessAPTagSourceConfigurablePriorities {
		if source, ok := target[priority]; ok && (current[priority] != source || len(remove) > 0) {
			items = append(items, WirelessAPTagSourcePriorityPriorities{
				Priority: types.Int64Value(priority),
				Source:   types.StringValue(source),
			})
		}
	}
	if len(items) == 0 {
		return nil
	}
	write := WirelessAPTagSourcePriority{Device: data.Device, Priorities: items}
	return helpers.EditConfig(ctx, device.NetconfClient, write.toBodyXML(ctx, write), device.AutoCommit)
}
