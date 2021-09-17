package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/CiscoDevNet/iosxe-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIOSXERest() *schema.Resource {
	return &schema.Resource{
		Description: "",

		CreateContext: resourceIOSXERestCreate,
		ReadContext:   resourceIOSXERestRead,
		UpdateContext: resourceIOSXERestCreate,
		DeleteContext: resourceIOSXERestDelete,

		Schema: map[string]*schema.Schema{
			"path": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The endopoint for the feature. E.g.: "/data/Cisco-IOS-XE-native:native"`,
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
		if payload == "" {
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
	case "PATCH":
		if payload == "" {
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
	case "GET":
		_, data, err := iosxeClient.Get(path, nil)
		if err != nil {
			return diag.FromErr(err)
		}
		log.Println("[DEBUG] GET Response: ", data)
		err = d.Set("response", data.String())
		if err != nil {
			return diag.FromErr(err)
		}
		err = d.Set("payload", nil)
		if err != nil {
			return diag.FromErr(err)
		}
	case "PUT":
		if payload == "" {
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
	case "DELETE":
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
