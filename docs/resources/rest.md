---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_rest Resource - terraform-provider-iosxe"
subcategory: "Generic Rest Resource"
description: |-
  Manages Cisco IOS XE Generic Rest Resource
---

# iosxe_rest (Resource)

Manages Cisco IOS XE Generic Rest Resource

## Example Usage

```terraform
resource "iosxe_rest" "example" {
  method = "GET"
  path = "/data/Cisco-IOS-XE-native:native"
  payload = ""
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **method** (String) The HTTP method. Allowed values for method are "GET", "POST", "PUT", "PATCH" and "DELETE".
- **path** (String) The endopoint for the feature. E.g.: "/data/Cisco-IOS-XE-native:native"

### Optional

- **payload** (String) The payload for the HTTP "POST", "PATCH", and "PUT". The provider will set it to null-value at the time of HTTP "GET" and "DELETE".

### Read-Only

- **response** (String) The HTTP response from the HTTP "GET". The provider will set it to null-value at the time of HTTP "POST", "PATCH", "PUT", and "DELETE".

---

> Note: The API requires sequential input for the payload (for some features like NAT and BGP) whereas Terraform sorts/reorders the specified payload which can cause the call to fail.
