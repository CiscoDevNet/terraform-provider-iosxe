package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/CiscoDevNet/iosxe-go-client/client"
	"github.com/CiscoDevNet/iosxe-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIOSXERest() *schema.Resource {
	return &schema.Resource{
		Description: "Manages Cisco IOS XE Generic Rest Resource",

		CreateContext: resourceIOSXERestCreate,
		ReadContext:   resourceIOSXERestRead,
		UpdateContext: resourceIOSXERestCreate,
		DeleteContext: resourceIOSXERestDelete,

		Schema: map[string]*schema.Schema{
			"path": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The endopoint for the feature. E.g.: "/data/Cisco-IOS-XE-native:native"`,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringIsNotWhiteSpace,
				),
			},
			"method": {
				Type:     schema.TypeString,
				Required: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice([]string{
						"GET",
						"PUT",
						"POST",
						"DELETE",
						"PATCH",
					}, false),
				),
				Description: `The HTTP method. Allowed values for method are "GET", "POST", "PUT", "PATCH" and "DELETE".`,
			},
			"payload": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The payload for the HTTP "POST", "PATCH", and "PUT". The provider will set it to null-value at the time of HTTP "GET" and "DELETE"`,
			},
			"response": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The HTTP response from the HTTP "GET". The provider will set it to null-value at the time of HTTP "POST", "PATCH", "PUT", and "DELETE"`,
			},
		},
	}
}

func resourceIOSXERestCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Println("[DEBUG] Begining Create method ")

	iosxeAC, _ := meta.(*apiClient)
	iosxeClient := iosxeAC.CiscoIOSXEClient

	path := d.Get("path").(string)
	payload := d.Get("payload").(string)
	oper := d.Get("method").(string)

	iosxeGM := &models.GenericModel{
		JSONPayload: payload,
	}

	switch oper {
	case "POST":
		diags := resourceIOSXERestPost(iosxeClient, path, iosxeGM, d)
		if diags != nil {
			return diags
		}
	case "PATCH":
		diags := resourceIOSXERestPatch(iosxeClient, path, iosxeGM, d)
		if diags != nil {
			return diags
		}
	case "GET":
		err := d.Set("payload", nil)
		if err != nil {
			return diag.FromErr(err)
		}

		diags := resourceIOSXERestGet(iosxeClient, path, d)
		if diags.HasError() {
			return diags
		}
	case "PUT":
		diags := resourceIOSXERestPut(iosxeClient, path, iosxeGM, d)
		if diags != nil {
			return diags
		}
	case "DELETE":
		diags := resourceIOSXERestRemove(iosxeClient, path, d)
		if diags != nil {
			return diags
		}
	default:
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Detail:   "[iosxe_rest] Can only use HTTP GET, POST, PATCH, PUT, or DELETE method",
				Summary:  fmt.Sprintf("Cannot use HTTP method: %v", oper),
			},
		}
	}

	d.SetId(fmt.Sprintf("%v", hashcode(oper+path+payload)))

	return resourceIOSXERestRead(ctx, d, meta)
}

func resourceIOSXERestRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}

// func resourceIOSXERestUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

// 	return nil
// }

func resourceIOSXERestDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}

func resourceIOSXERestGet(iosxeClient *client.V2, path string, d *schema.ResourceData) diag.Diagnostics {
	_, data, err := iosxeClient.Get(path, nil)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Println("[DEBUG] GET Response: ", data)
	err = d.Set("response", data.String())
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceIOSXERestPost(iosxeClient *client.V2, path string, iosxeGM *models.GenericModel, d *schema.ResourceData) diag.Diagnostics {
	if iosxeGM.JSONPayload == "" {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Detail:   "[iosxe_rest] Payload is required at the time of POST call",
				Summary:  "POST call failed. Payload empty.",
			},
		}
	}
	_, _, err := iosxeClient.Create(path, iosxeGM)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("response", nil)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceIOSXERestPatch(iosxeClient *client.V2, path string, iosxeGM *models.GenericModel, d *schema.ResourceData) diag.Diagnostics {
	if iosxeGM.JSONPayload == "" {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Detail:   "[iosxe_rest] Payload is required at the time of PATCH call",
				Summary:  "PATCh call failed. Payload empty.",
			},
		}
	}
	_, _, err := iosxeClient.Patch(path, iosxeGM)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("response", nil)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceIOSXERestPut(iosxeClient *client.V2, path string, iosxeGM *models.GenericModel, d *schema.ResourceData) diag.Diagnostics {
	if iosxeGM.JSONPayload == "" {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Detail:   "[iosxe_rest] Payload is required at the time of PUT call",
				Summary:  "PUT call failed. Payload empty.",
			},
		}
	}
	_, err := iosxeClient.Update(path, iosxeGM)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("response", nil)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceIOSXERestRemove(iosxeClient *client.V2, path string, d *schema.ResourceData) diag.Diagnostics {
	_, err := iosxeClient.Delete(path)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("response", nil)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("payload", nil)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
