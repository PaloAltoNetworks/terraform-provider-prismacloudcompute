package prismacloudcompute

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDsPoliciesComplianceContainer(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsPoliciesComplianceContainerConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloudcompute_policies_compliance_container.test", "total"),
				),
			},
		},
	})
}

func testAccDsPoliciesComplianceContainerConfig() string {
	return `
data "prismacloudcompute_policies" "test" {
    learningDisabled = true
    rule {
        name = "my rule"
    }
}
`
}
