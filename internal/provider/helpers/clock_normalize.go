// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package helpers

import (
	"fmt"
	"strings"
)

// NormalizeTimeString normalizes a time string like "2:30" to "02:30"
// by ensuring the hour component always has a leading zero when it is
// a single digit. IOS-XE strips leading zeros from time values in
// NETCONF get-config responses, but users typically configure them
// with leading zeros (e.g., "02:00").
func NormalizeTimeString(value string) string {
	parts := strings.SplitN(value, ":", 2)
	if len(parts) != 2 {
		return value
	}
	hour := parts[0]
	minute := parts[1]
	if len(hour) == 1 {
		hour = "0" + hour
	}
	if len(minute) == 1 {
		minute = "0" + minute
	}
	return fmt.Sprintf("%s:%s", hour, minute)
}
