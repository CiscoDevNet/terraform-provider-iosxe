---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_route_map Data Source - terraform-provider-iosxe"
subcategory: "System"
description: |-
  This data source can read the Route Map configuration.
---

# iosxe_route_map (Data Source)

This data source can read the Route Map configuration.

## Example Usage

```terraform
data "iosxe_route_map" "example" {
  name = "RM1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) WORD;;Route map tag

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `entries` (Attributes List) Sequence to insert to/delete from existing route-map entry (see [below for nested schema](#nestedatt--entries))
- `id` (String) The path of the retrieved object.

<a id="nestedatt--entries"></a>
### Nested Schema for `entries`

Read-Only:

- `continue` (Boolean) Continue on a different entry within the route-map
- `continue_sequence_number` (Number) Route-map entry sequence number
- `description` (String) Route-map comment
- `match_as_paths` (List of Number) AS path access-list
- `match_as_paths_legacy` (List of Number) AS path access-list (OBSOLETE - please use route-map configuration in Cisco-IOS-XE-bgp.yang)
- `match_community_list_exact_match` (Boolean) Do exact matching of communities
- `match_community_lists` (List of String) Named Access List
- `match_community_lists_legacy` (List of String) Named Access List (OBSOLETE- please use community-list in Cisco-IOS-XE-bgp.yang)
- `match_extcommunity_lists` (List of String) Named Access List
- `match_extcommunity_lists_legacy` (List of String) Named Access List (OBSOLETE- please use extcommunity-list in Cisco-IOS-XE-bgp.yang)
- `match_interfaces` (List of String)
- `match_ip_address_access_lists` (List of String)
- `match_ip_address_prefix_lists` (List of String) Match entries of prefix-lists
- `match_ip_next_hop_access_lists` (List of String)
- `match_ip_next_hop_prefix_lists` (List of String) Match entries of prefix-lists
- `match_ipv6_address_access_lists` (String) IPv6 access-list name
- `match_ipv6_address_prefix_lists` (String) IPv6 prefix-list name
- `match_ipv6_next_hop_access_lists` (String) IPv6 access-list name
- `match_ipv6_next_hop_prefix_lists` (String) IPv6 prefix-list name
- `match_local_preferences` (List of Number)
- `match_local_preferences_legacy` (List of Number)
- `match_route_type_external` (Boolean) external route (BGP, EIGRP and OSPF type 1/2)
- `match_route_type_external_type_1` (Boolean) OSPF external type 1 route
- `match_route_type_external_type_2` (Boolean) OSPF external type 2 route
- `match_route_type_internal` (Boolean) internal route (including OSPF intra/inter area)
- `match_route_type_level_1` (Boolean) IS-IS level-1 route
- `match_route_type_level_2` (Boolean) IS-IS level-2 route
- `match_route_type_local` (Boolean) locally generated route (OBSOLETE - please use route-map configuration in Cisco-IOS-XE-bgp.yang)
- `match_source_protocol_bgp` (List of String) Border Gateway Protocol (BGP)
- `match_source_protocol_connected` (Boolean) Connected
- `match_source_protocol_eigrp` (List of String) Border Gateway Protocol (BGP)
- `match_source_protocol_isis` (Boolean) ISO IS-IS
- `match_source_protocol_lisp` (Boolean) Locator ID Separation Protocol (LISP)
- `match_source_protocol_ospf` (List of String) Open Shortest Path First (OSPF)
- `match_source_protocol_ospfv3` (List of String) OSPFv3
- `match_source_protocol_rip` (Boolean) Routing Information Protocol (RIP)
- `match_source_protocol_static` (Boolean) Static routes
- `match_tags` (List of Number) Tag value (DEPRECATED - please use tag-val)
- `match_track` (Number) tracking object
- `operation` (String) Route map permit/deny set operations
- `seq` (Number)
- `set_as_path_prepend_as` (String) BGP AS number
- `set_as_path_prepend_as_legacy` (String) <1-65535>;;AS number (OBSOLETE - please use route-map configuration in Cisco-IOS-XE-bgp.yang)
- `set_as_path_prepend_last_as` (Number)
- `set_as_path_prepend_last_as_legacy` (Number)
- `set_as_path_replace_any` (Boolean) Replace each AS number in the AS-path with the local AS
- `set_as_path_replace_as` (Attributes List) (see [below for nested schema](#nestedatt--entries--set_as_path_replace_as))
- `set_as_path_tag` (Boolean) Set the tag as an AS-path attribute
- `set_as_path_tag_legacy` (Boolean) Set the tag as an AS-path attribute (OBSOLETE - please use route-map configuration in Cisco-IOS-XE-bgp.yang)
- `set_communities` (List of String)
- `set_communities_additive` (Boolean)
- `set_communities_additive_legacy` (Boolean)
- `set_communities_legacy` (List of String)
- `set_community_list_delete` (Boolean) Delete matching communities
- `set_community_list_delete_legacy` (Boolean) Delete matching communities (OBSOLETE - please use route-map configuration in Cisco-IOS-XE-bgp.yang)
- `set_community_list_expanded` (Number)
- `set_community_list_expanded_legacy` (Number)
- `set_community_list_name` (String)
- `set_community_list_name_legacy` (String)
- `set_community_list_standard` (Number)
- `set_community_list_standard_legacy` (Number)
- `set_community_none` (Boolean) No community attribute
- `set_community_none_legacy` (Boolean) No community attribute (OBSOLETE - please use route-map configuration in Cisco-IOS-XE-bgp.yang)
- `set_default_interfaces` (List of String) SPAN source interface
- `set_extcomunity_rt` (List of String)
- `set_extcomunity_rt_legacy` (List of String)
- `set_extcomunity_soo` (String)
- `set_extcomunity_soo_legacy` (String)
- `set_extcomunity_vpn_distinguisher` (String) VPN Distinguisher Extended Community
- `set_extcomunity_vpn_distinguisher_additive` (Boolean)
- `set_extcomunity_vpn_distinguisher_legacy` (String)
- `set_global` (Boolean)
- `set_interfaces` (List of String) Interface specific information
- `set_ip_address` (String) Specify prefix-list
- `set_ip_default_global_next_hop_address` (List of String)
- `set_ip_default_next_hop_address` (List of String)
- `set_ip_global_next_hop_address` (List of String)
- `set_ip_next_hop_address` (List of String) IP address of next hop
- `set_ip_next_hop_self` (Boolean) Use self address (for BGP only)
- `set_ip_qos_group` (Number)
- `set_ipv6_address` (List of String) IPv6 prefix-list
- `set_ipv6_default_global_next_hop` (String) Next hop along path
- `set_ipv6_default_next_hop` (List of String) Default next hop IPv6 address
- `set_ipv6_next_hop` (List of String) Next hop IPv6 address
- `set_level_1` (Boolean) Import into a level-1 area
- `set_level_1_2` (Boolean) Import into level-1 and level-2
- `set_level_2` (Boolean) Import into level-2 sub-domain
- `set_local_preference` (Number) Preference value
- `set_local_preference_legacy` (Number) Preference value (OBSOLETE - please use route-map configuration in Cisco-IOS-XE-bgp.yang)
- `set_metric_change` (String) +/-<metric>;;Add or subtract metric
- `set_metric_delay` (String) EIGRP delay metric, in 10 microsecond units
- `set_metric_loading` (Number) EIGRP Effective bandwidth metric (Loading) where 255 is 100% loaded
- `set_metric_mtu` (Number) EIGRP MTU of the path
- `set_metric_reliability` (Number) EIGRP reliability metric where 255 is 100% reliable
- `set_metric_type` (String) Type of metric for destination routing protocol
- `set_metric_value` (Number) Metric value or Bandwidth in Kbits per second
- `set_tag` (Number) Tag value (DEPRECATED - please use tag-val)
- `set_vrf` (String) VPN Routing/Forwarding instance name
- `set_weight` (Number) BGP weight for routing table
- `set_weight_legacy` (Number) BGP weight for routing table (OBSOLETE - please use route-map configuration in Cisco-IOS-XE-bgp.yang)

<a id="nestedatt--entries--set_as_path_replace_as"></a>
### Nested Schema for `entries.set_as_path_replace_as`

Read-Only:

- `as_number` (String) <1-65535>;;AS number
