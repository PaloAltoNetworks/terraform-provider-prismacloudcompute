package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	PrismacloudcomputeJsonConfigFileEnvVar = "PRISMACLOUDCOMPUTE_CONFIG_FILE"
)

var (
	testAccProviders                   map[string]*schema.Provider
	testAccProvider                    *schema.Provider
	sessionTimeoutOrig, sessionTimeout int
)

func init() {
	var err error

	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"prismacloudcompute": testAccProvider,
	}

	client := &api.Client{}
	if err = client.Initialize(os.Getenv(PrismacloudcomputeJsonConfigFileEnvVar)); err == nil {
		if err != nil {
			fmt.Printf("Error initializing client")
		}
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if os.Getenv(PrismacloudcomputeJsonConfigFileEnvVar) == "" {
		t.Fatalf("%s must be set for acceptance tests", PrismacloudcomputeJsonConfigFileEnvVar)
	}
}
