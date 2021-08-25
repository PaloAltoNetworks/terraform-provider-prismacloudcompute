package prismacloudcompute

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDsPoliciesComplianceHost(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsPoliciesComplianceHostConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloudcompute_policies_compliance_host.test", "total"),
				),
			},
		},
	})
}

func testAccDsPoliciesComplianceHostConfig() string {
	return `
data "prismacloudcompute_policies" "test" {
    learningDisabled = true
    rule {
        name = "my rule"
    }
}
`
}
