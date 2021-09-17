package provider

import (
	"context"
	"fmt"

	"github.com/CiscoDevNet/iosxe-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	host           = "host"
	deviceUsername = "device_username"
	devicePassword = "device_password"
	timeoutKey     = "request_timeout"
	insecureKey    = "insecure"
	proxyURLKey    = "proxy_url"
	proxyCredsKey  = "proxy_creds"
	caFileKey      = "ca_file"

	version = 1
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

// New instance of IOS XE Provider
func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				host: {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("HOST_IOSXE", nil),
					Description: `The host of the IOSXE switch with its port. E.g.: "https://10.1.1.5:443".`,
				},
				deviceUsername: {
					Type:        schema.TypeString,
					Required:    true,
					Sensitive:   true,
					DefaultFunc: schema.EnvDefaultFunc("DEVICE_USERNAME_IOSXE", nil),
					Description: `The Username of the IOSXE switch. E.g.: "admin".`,
				},
				devicePassword: {
					Type:        schema.TypeString,
					Required:    true,
					Sensitive:   true,
					DefaultFunc: schema.EnvDefaultFunc("DEVICE_PASSWORD_IOSXE", nil),
					Description: `The Password of the IOSXE switch. E.g.: "Cisco123".`,
				},
				insecureKey: {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     true,
					Description: "Allow insecure TLS. Default: true, means the API call is insecure",
				},
				timeoutKey: {
					Type:        schema.TypeInt,
					Optional:    true,
					Default:     30,
					Description: "Timeout for HTTP requests. Default value: 30",
				},
				caFileKey: {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("CA_FILE_IOSXE", nil),
					Description: "The path to CA certificate file (PEM). In case, certificate is based on legacy CN instead of ASN, set env. variable `GODEBUG=x509ignoreCN=0`. This can also be set by environment variable `CA_FILE_IOSXE`",
				},
				proxyURLKey: {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("PROXY_URL_IOSXE", nil),
					Description: "Proxy Server URL with port number. This can also be set by environment variable `PROXY_URL_IOSXE`",
				},
				proxyCredsKey: {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("PROXY_CREDENTIALS_IOSXE", nil),
					Description: "Proxy credential in format `username:password`. This can also be set by environment variable `PROXY_CREDENTIALS_IOSXE`",
				},
			},
			ResourcesMap: map[string]*schema.Resource{
				"iosxe_rest": resourceIOSXERest(),
			},
		}
		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type apiClient struct {
	CiscoIOSXEClient *client.V2
}

func configure(version string, p *schema.Provider) func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		var diagnostics diag.Diagnostics
		if diagnostics = validateInput(d); diagnostics.HasError() {
			return nil, diagnostics
		}
		iosxeV2Client, err := client.NewV2(
			d.Get(host).(string),
			d.Get(deviceUsername).(string),
			d.Get(devicePassword).(string),
			d.Get(timeoutKey).(int),
			d.Get(insecureKey).(bool),
			d.Get(caFileKey).(string),
			d.Get(proxyURLKey).(string),
			d.Get(proxyCredsKey).(string),
		)

		if err != nil {
			diagnostics = append(diagnostics, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create Cisco IOSXE client",
				Detail:   fmt.Sprintf("Unable to create Cisco IOSXE client. Error - %v", err),
			})
			return nil, diagnostics
		}

		return &apiClient{
			CiscoIOSXEClient: iosxeV2Client,
		}, nil
	}
}

func validateInput(d *schema.ResourceData) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	hostURL := d.Get(host).(string)
	if hostURL == "" {
		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Host URL is required",
			Detail:   "Host URL must be set for Cisco IOSXE provider",
		})
	}
	du := d.Get(deviceUsername).(string)
	if du == "" {
		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Device Username is required",
			Detail:   "Device Username must be set for Cisco IOSXE provider",
		})
	}
	dp := d.Get(devicePassword).(string)
	if dp == "" {
		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Device Password is required",
			Detail:   "Device Password must be set for Cisco IOSXE provider",
		})
	}
	return diagnostics
}
