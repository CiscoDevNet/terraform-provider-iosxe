package provider

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

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

func testAccProviderFactoriesInit(provider **schema.Provider, providerName string) map[string]func() (*schema.Provider, error) {
	var factories = make(map[string]func() (*schema.Provider, error))

	p := New("dev")()

	factories[providerName] = func() (*schema.Provider, error) {
		return p, nil
	}

	if provider != nil {
		*provider = p
	}

	return factories
}

func testAccProviderFactoriesInternal(provider **schema.Provider) map[string]func() (*schema.Provider, error) {
	return testAccProviderFactoriesInit(provider, "iosxe")
}

// testAccPreCheck To validate the deletion of features in acceptance test
func testAccCheckIOSXEGeneralizeDestroy(providerInstance *schema.Provider) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		pConfig := (*providerInstance).Meta().(*apiClient)
		iosxeClient := pConfig.CiscoIOSXEClient
		r, _ := regexp.Compile("iosxe_rest.*")

		for _, rs := range s.RootModule().Resources {
			if r.MatchString(rs.Type) {
				resp, _, err := iosxeClient.Get(rs.Primary.Attributes["path"], nil)
				log.Printf("LOL CHECK %#v\n\n\n%#v\n", resp, err)
				if err != nil {
					if !(strings.Contains(err.Error(), "not-found") || resp.Status == "404 Not Found" || resp.Status == "204 No Content") {
						return fmt.Errorf("%s still exists, ResourceID: %v", rs.Type, rs.Primary.Attributes["path"])
					}
				} else if !(resp.ContentLength == 0 || resp.Status == "204 No Content") {
					return fmt.Errorf("%s still exists, ResourceID: %v", rs.Type, rs.Primary.Attributes["path"])
				}
			}
		}
		return nil
	}
}
