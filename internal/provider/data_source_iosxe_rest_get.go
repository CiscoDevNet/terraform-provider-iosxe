package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func datasourceIOSXERest() *schema.Resource {
	return &schema.Resource{
		Description: "Represents Cisco IOS XE Generic Rest Resource",

		ReadContext: datasourceIOSXERestRead,

		Schema: map[string]*schema.Schema{
			"path": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The endopoint for the feature. E.g.: "/data/Cisco-IOS-XE-native:native"`,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringIsNotWhiteSpace,
				),
			},
			"response": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The HTTP response from the HTTP "GET". The provider will set it to null-value at the time of HTTP "POST", "PATCH", "PUT", and "DELETE."`,
			},
		},
	}
}

func datasourceIOSXERestRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Println("[DEBUG] Begining Read method ")

	iosxeAC, _ := meta.(*apiClient)
	iosxeClient := iosxeAC.CiscoIOSXEClient

	path := d.Get("path").(string)

	diags := resourceIOSXERestGet(iosxeClient, path, d)
	if diags.HasError() {
		return diags
	}

	d.SetId(fmt.Sprintf("%v", hashcode("GET"+path)))

	return nil
}
