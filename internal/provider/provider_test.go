package provider

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
var providerFactories = map[string]func() (*schema.Provider, error){
	"iosxe": func() (*schema.Provider, error) {
		return New("dev")(), nil
	},
}

func TestProvider(t *testing.T) {
	if err := New("dev")().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	_ = New("dev")()
}

// testAccPreCheck To validate required environment values are available
func testAccPreCheck(t *testing.T) {
	missing := []string{}
	for _, envKey := range []string{"HOST_IOSXE", "DEVICE_USERNAME_IOSXE", "DEVICE_PASSWORD_IOSXE"} {
		if v := os.Getenv(envKey); v == "" {
			missing = append(missing, envKey)
		}
	}
	if len(missing) > 0 {
		t.Fatalf("Required environment variables missing: %v", missing)
	}
}

// testAccPreCheck To validate the deletion of features in acceptance test
func testAccCheckIOSXEGeneralizeDestroy() func(s *terraform.State) error {
	return func(s *terraform.State) error {
		r, _ := regexp.Compile("iosxe_rest.*")
		for _, rs := range s.RootModule().Resources {
			if r.MatchString(rs.Type) {
				return fmt.Errorf("%s still exists, ResourceID: %v", "iosxe_rest", rs)
			}
		}
		return nil
	}
}
