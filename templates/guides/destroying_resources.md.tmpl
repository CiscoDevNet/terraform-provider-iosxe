---
subcategory: "Guides"
page_title: "Destroying Resources"
description: |-
    Destroying Resources
---

# Destroying Resources

When it comes to destroying resources, one can use the `delete_mode` attribute to control the behavior of the resource when it is destroyed. The `delete_mode` attribute can be set to `all` or `attributes`. The default value is documented under the `delete_mode` attribute in the resource documentation.

- `all`: Delete the entire object (YANG container) being managed.
- `attributes`: Delete only the individual resource attributes configured explicitly and leave everything else as-is.

Depending on the use case, one or the other may be more appropriate. For example, if one wants to use Terraform to just configure a specific BGP setting using the `iosxe_bgp` resource but manage the remaining BGP configuration manually, one can use the `attributes` delete mode.

```terraform
resource "iosxe_bgp" "bgp" { 
  asn                  = "65000"
  log_neighbor_changes = true
  delete_mode          = "attributes"
}
```

This way, only the `log_neighbor_changes` attribute will be removed when the resource is destroyed, while the rest of the BGP configuration will remain unchanged. In other words, the `delete_mode` attribute controls the behavior of the `destroy` operation.

If the `delete_mode` attribute is set to `all`, the entire BGP configuration would have been removed when the resource was destroyed.

For some resources only one mode is supported, for example the `iosxe_access_list_standard` resource only supports the `all` mode and does not expose a `delete_mode` attribute, because it is not possible to manage individual access list entries outside of Terraform.
