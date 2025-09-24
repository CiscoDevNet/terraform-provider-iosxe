---
subcategory: "Guides"
page_title: "Selective Deploy"
description: |-
    Howto selectively deploy to devices.
---

# Selective Deploy

## Overview

Selective deployment allows you to deploy Terraform configurations to a subset of devices from your `devices` list, while keeping the remaining devices in a "frozen" state where no changes are applied. This capability enables staged rollouts, maintenance workflows, and risk mitigation strategies for large-scale network deployments.

### Key Benefits

- **Risk Reduction**: Deploy to a small subset first to validate changes
- **Staged Rollouts**: Progressively deploy across device groups
- **Maintenance Windows**: Exclude devices undergoing maintenance
- **Emergency Procedures**: Quickly isolate or target specific devices

### How It Works

When `selected_devices` is configured, only the specified devices from your `devices` list will be actively managed by Terraform. Non-selected devices maintain their current configuration state and are skipped during plan and apply operations. This setting overrides individual device `managed` attributes.

### Understanding State Behavior
When using selective deployment, Terraform maintains state for ALL devices but only modifies selected devices:

- **Selected Devices**: Configuration changes applied, state updated
- **Non-Selected Devices**: State preserved, no modifications made
- **State File**: Contains all devices regardless of selection

## Configuration Reference

### HCL Configuration

Configure selective deployment using the `selected_devices` attribute in your provider block:

```hcl
provider "iosxe" {
  selected_devices = ["switch-01", "switch-02"]
  devices = [
    { name = "switch-01", url = "https://10.1.1.10" },  # Managed
    { name = "switch-02", url = "https://10.1.1.20" },  # Managed
    { name = "switch-03", url = "https://10.1.1.30" },  # Skipped
    { name = "switch-04", url = "https://10.1.1.40" }   # Skipped
  ]
}

```

### Environment Variable Configuration

Alternatively, use the `IOSXE_SELECTED_DEVICES` environment variable with comma-separated device names:

```bash
export IOSXE_SELECTED_DEVICES="switch-01,switch-02"
```

```hcl
provider "iosxe" {
  devices = [
    { name = "switch-01", url = "https://10.1.1.10" },
    { name = "switch-02", url = "https://10.1.1.20" },
    { name = "switch-03", url = "https://10.1.1.30" },
    { name = "switch-04", url = "https://10.1.1.40" }
  ]
}

```

### Default Behavior

When `selected_devices` is not specified, all devices in the `devices` list are managed by default.
