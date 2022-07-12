package provider

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/policy"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPolicyComplianceContainerConfig(t *testing.T) {
	fmt.Printf("\n\nStart TestAccPolicyComplianceContainerConfig")
	var o policy.CompliancePolicy
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPolicyComplianceContainerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyComplianceContainerConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyComplianceContainerExists("prismacloudcompute_policies_compliance_container.test", &o),
					testAccCheckPolicyComplianceContainerAttributes(&o, id, "network"),
				),
			},
			{
				Config: testAccPolicyComplianceContainerConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyComplianceContainerExists("prismacloudcompute_policies_compliance_container.test", &o),
					testAccCheckPolicyComplianceContainerAttributes(&o, id, "network"),
				),
			},
		},
	})
}

func TestAccPolicyComplianceContainerNetwork(t *testing.T) {
	var o policy.CompliancePolicy
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPolicyComplianceContainerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyComplianceContainerConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyComplianceContainerExists("prismacloudcompute_policies_compliance_container.test", &o),
					testAccCheckPolicyComplianceContainerAttributes(&o, id, "network"),
				),
			},
			{
				Config: testAccPolicyComplianceContainerConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyComplianceContainerExists("prismacloudcompute_policies_compliance_container.test", &o),
					testAccCheckPolicyComplianceContainerAttributes(&o, id, "network"),
				),
			},
		},
	})
}

func TestAccPolicyComplianceContainerAuditEvent(t *testing.T) {
	var o policy.CompliancePolicy
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPolicyComplianceContainerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyComplianceContainerConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyComplianceContainerExists("prismacloudcompute_policies_compliance_container.test", &o),
					testAccCheckPolicyComplianceContainerAttributes(&o, id, "network"),
				),
			},
			{
				Config: testAccPolicyComplianceContainerConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyComplianceContainerExists("prismacloudcompute_policies_compliance_container.test", &o),
					testAccCheckPolicyComplianceContainerAttributes(&o, id, "network"),
				),
			},
		},
	})
}

func testAccCheckPolicyComplianceContainerExists(n string, o *policy.CompliancePolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// return fmt.Errorf("What is the name: %s", o.PolicyId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object label Id is not set")
		}

		client := testAccProvider.Meta().(*api.Client)
		lo, err := policy.GetComplianceContainer(*client)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}
		*o = lo

		return nil
	}
}

func testAccCheckPolicyComplianceContainerAttributes(o *policy.CompliancePolicy, id string, policyType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if o.Type != policyType {
			return fmt.Errorf("PolicyType is %s, expected %s", o.Type, policyType)
		}

		return nil
	}
}

func testAccPolicyComplianceContainerDestroy(s *terraform.State) error {
	/*	client := testAccProvider.Meta().(*api.Client)

		for _, rs := range s.RootModule().Resources {

			if rs.Type != "prismacloudcompute_policyComplianceContainer" {
				continue
			}

			if rs.Primary.ID != "" {
				name := rs.Primary.ID
				if err := policy.Delete(client, name); err == nil {
					return fmt.Errorf("Object %q still exists", name)
				}
			}
			return nil
		}
	*/
	return nil
}

func testAccPolicyComplianceContainerConfig(id string) string {
	var buf bytes.Buffer
	buf.Grow(500)

	buf.WriteString(fmt.Sprintf(`
resource "prismacloudcompute_policyComplianceContainer" "test" {
    name = %q
}`, id))

	return buf.String()
}
