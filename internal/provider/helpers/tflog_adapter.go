// Copyright © 2025 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package helpers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-netconf"
)

// TflogAdapter adapts go-netconf's Logger interface to Terraform's tflog package.
//
// This adapter bridges the gap between go-netconf's logging interface and Terraform's
// context-based logging. It leverages go-netconf's context-aware Logger interface,
// which passes context as the first parameter to all logging methods.
//
// The adapter automatically creates the "netconf" subsystem on first use, eliminating
// the need for manual subsystem creation in provider code.
//
// Thread-safety: This adapter is safe for concurrent use. Context is passed per
// log call via the Logger interface, eliminating the need for internal state.
//
// Usage in provider initialization:
//
//	logger := helpers.NewTflogAdapter()
//	client, err := netconf.NewClient(host,
//	    netconf.Username(username),
//	    netconf.Password(password),
//	    netconf.WithLogger(logger),
//	)
//
// Usage in resource operations:
//
//	// Context is automatically propagated through go-netconf's Logger interface
//	_, err := device.NetconfClient.GetConfig(ctx, "running", filter)
//	// Logs will automatically use the correct context and netconf subsystem
type TflogAdapter struct{}

var _ netconf.Logger = (*TflogAdapter)(nil)

// NewTflogAdapter creates a new Terraform logging adapter.
//
// The adapter automatically receives context from go-netconf's Logger interface
// on each logging call, ensuring proper context propagation without manual management.
func NewTflogAdapter() *TflogAdapter {
	return &TflogAdapter{}
}

// Debug logs a debug message with structured key-value pairs to tflog.SubsystemDebug.
//
// Debug logs are typically used for detailed operational information useful
// for troubleshooting and development.
//
// Context is provided by go-netconf's Logger interface, ensuring automatic
// propagation of trace IDs, request IDs, and deadlines.
//
// Logs are written to the "netconf" subsystem for proper organization and filtering.
// The subsystem is automatically created if it doesn't exist.
func (t *TflogAdapter) Debug(ctx context.Context, msg string, keysAndValues ...any) {
	if ctx == nil {
		return
	}
	// Ensure subsystem exists (idempotent operation)
	ctx = tflog.NewSubsystem(ctx, "netconf")

	fields := keysAndValuesToMap(keysAndValues)
	if fields != nil {
		tflog.SubsystemDebug(ctx, "netconf", msg, fields)
	} else {
		tflog.SubsystemDebug(ctx, "netconf", msg)
	}
}

// Info logs an informational message with structured key-value pairs to tflog.SubsystemInfo.
//
// Info logs represent normal operational messages that highlight progress
// or state changes.
//
// Context is provided by go-netconf's Logger interface, ensuring automatic
// propagation of trace IDs, request IDs, and deadlines.
//
// Logs are written to the "netconf" subsystem for proper organization and filtering.
// The subsystem is automatically created if it doesn't exist.
func (t *TflogAdapter) Info(ctx context.Context, msg string, keysAndValues ...any) {
	if ctx == nil {
		return
	}
	// Ensure subsystem exists (idempotent operation)
	ctx = tflog.NewSubsystem(ctx, "netconf")

	fields := keysAndValuesToMap(keysAndValues)
	if fields != nil {
		tflog.SubsystemInfo(ctx, "netconf", msg, fields)
	} else {
		tflog.SubsystemInfo(ctx, "netconf", msg)
	}
}

// Warn logs a warning message with structured key-value pairs to tflog.SubsystemWarn.
//
// Warnings indicate potentially harmful situations that don't prevent
// operation but should be addressed.
//
// Context is provided by go-netconf's Logger interface, ensuring automatic
// propagation of trace IDs, request IDs, and deadlines.
//
// Logs are written to the "netconf" subsystem for proper organization and filtering.
// The subsystem is automatically created if it doesn't exist.
func (t *TflogAdapter) Warn(ctx context.Context, msg string, keysAndValues ...any) {
	if ctx == nil {
		return
	}
	// Ensure subsystem exists (idempotent operation)
	ctx = tflog.NewSubsystem(ctx, "netconf")

	fields := keysAndValuesToMap(keysAndValues)
	if fields != nil {
		tflog.SubsystemWarn(ctx, "netconf", msg, fields)
	} else {
		tflog.SubsystemWarn(ctx, "netconf", msg)
	}
}

// Error logs an error message with structured key-value pairs to tflog.SubsystemError.
//
// Errors indicate serious problems that prevent successful operation.
//
// Context is provided by go-netconf's Logger interface, ensuring automatic
// propagation of trace IDs, request IDs, and deadlines.
//
// Logs are written to the "netconf" subsystem for proper organization and filtering.
// The subsystem is automatically created if it doesn't exist.
func (t *TflogAdapter) Error(ctx context.Context, msg string, keysAndValues ...any) {
	if ctx == nil {
		return
	}
	// Ensure subsystem exists (idempotent operation)
	ctx = tflog.NewSubsystem(ctx, "netconf")

	fields := keysAndValuesToMap(keysAndValues)
	if fields != nil {
		tflog.SubsystemError(ctx, "netconf", msg, fields)
	} else {
		tflog.SubsystemError(ctx, "netconf", msg)
	}
}

// keysAndValuesToMap converts variadic key-value pairs to a map for tflog.
//
// tflog expects structured fields as map[string]any, while go-netconf uses
// variadic key-value pairs (following Go's slog standard). This function performs
// the conversion.
//
// Example:
//
//	Input:  "operation", "GetConfig", "datastore", "running", "duration", 150
//	Output: map[string]any{"operation": "GetConfig", "datastore": "running", "duration": 150}
//
// If the number of elements is odd, the last key will have a nil value.
func keysAndValuesToMap(keysAndValues []any) map[string]any {
	if len(keysAndValues) == 0 {
		return nil
	}

	fields := make(map[string]any, len(keysAndValues)/2)
	for i := 0; i < len(keysAndValues); i += 2 {
		key, ok := keysAndValues[i].(string)
		if !ok {
			// Skip non-string keys
			continue
		}

		var value any
		if i+1 < len(keysAndValues) {
			value = keysAndValues[i+1]
		}

		fields[key] = value
	}

	return fields
}
