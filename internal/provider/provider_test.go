package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	PrismacloudcomputeJsonConfigFileEnvVar = "../creds.json"
)

var (
	testAccProviders map[string]*schema.Provider
	testAccProvider  *schema.Provider
	// sessionTimeoutOrig, sessionTimeout int
)

// func init() {
// 	fmt.Printf("\n\nStart Provider init()\n")
// 	var err error

// 	testAccProvider = Provider().(*schema.Provider)
// 	testAccProviders = map[string]terraform.ResourceProvider{
// 		"prismacloudcompute": testAccProvider,
// 	}

// 	client := &api.Client{}
// 	if err = client.Initialize(os.Getenv(PrismacloudcomputeJsonConfigFileEnvVar)); err == nil {
// 		if err != nil {
// 			fmt.Sprintf("Error initializing client")
// 		}
// 	}
// }

func TestProvider(t *testing.T) {
	fmt.Printf("\n\nStart Provider TestProvider()\n")
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	fmt.Printf("\n\nStart Provider TestProvider_impl()\n")
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	fmt.Printf("\n\nStart Provider testAccPreCheck()\n")
	if os.Getenv(PrismacloudcomputeJsonConfigFileEnvVar) == "" {
		t.Fatalf("%s must be set for acceptance tests", PrismacloudcomputeJsonConfigFileEnvVar)
	}
}
