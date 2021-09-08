package prismacloudcompute

import (
	"bytes"
	"fmt"
	"testing"

	pcc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policies"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccPolicyComplianceCiImagesConfig(t *testing.T) {
	fmt.Printf("\n\nStart TestAccPolicyComplianceCiImagesConfig")
	var o policies.Policy
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPolicyComplianceCiImagesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyComplianceCiImagesConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyComplianceCiImagesExists("prismacloudcompute_policies_compliance_container.test", &o),
					testAccCheckPolicyComplianceCiImagesAttributes(&o, id, "network"),
				),
			},
			{
				Config: testAccPolicyComplianceCiImagesConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyComplianceCiImagesExists("prismacloudcompute_policies_compliance_container.test", &o),
					testAccCheckPolicyComplianceCiImagesAttributes(&o, id, "network"),
				),
			},
		},
	})
}

func TestAccPolicyComplianceCiImagesNetwork(t *testing.T) {
	var o policies.Policy
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPolicyComplianceCiImagesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyComplianceCiImagesConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyComplianceCiImagesExists("prismacloudcompute_policies_compliance_container.test", &o),
					testAccCheckPolicyComplianceCiImagesAttributes(&o, id, "network"),
				),
			},
			{
				Config: testAccPolicyComplianceCiImagesConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyComplianceCiImagesExists("prismacloudcompute_policies_compliance_container.test", &o),
					testAccCheckPolicyComplianceCiImagesAttributes(&o, id, "network"),
				),
			},
		},
	})
}

func TestAccPolicyComplianceCiImagesAuditEvent(t *testing.T) {
	var o policies.Policy
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPolicyComplianceCiImagesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyComplianceCiImagesConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyComplianceCiImagesExists("prismacloudcompute_policies_compliance_container.test", &o),
					testAccCheckPolicyComplianceCiImagesAttributes(&o, id, "network"),
				),
			},
			{
				Config: testAccPolicyComplianceCiImagesConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyComplianceCiImagesExists("prismacloudcompute_policies_compliance_container.test", &o),
					testAccCheckPolicyComplianceCiImagesAttributes(&o, id, "network"),
				),
			},
		},
	})
}

func testAccCheckPolicyComplianceCiImagesExists(n string, o *policies.Policy) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		return fmt.Errorf("What is the name: %s", o.PolicyId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object label Id is not set")
		}

		client := testAccProvider.Meta().(*pcc.Client)
		lo, err := policies.Get(*client, policies.ComplianceCiImagesEndpoint)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}
		*o = lo

		return nil
	}
}

func testAccCheckPolicyComplianceCiImagesAttributes(o *policies.Policy, id string, policyType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.PolicyId != id {
			return fmt.Errorf("\n\nPolicyId is %s, expected %s", o.PolicyId, id)
		} else {
			fmt.Printf("\n\nName is %s", o.PolicyId)
		}

		if o.PolicyType != policyType {
			return fmt.Errorf("PolicyType is %s, expected %s", o.PolicyType, policyType)
		}

		return nil
	}
}

func testAccPolicyComplianceCiImagesDestroy(s *terraform.State) error {
	/*	client := testAccProvider.Meta().(*pcc.Client)

		for _, rs := range s.RootModule().Resources {

			if rs.Type != "prismacloudcompute_policycomplianceciimages" {
				continue
			}

			if rs.Primary.ID != "" {
				name := rs.Primary.ID
				if err := policies.Delete(client, name); err == nil {
					return fmt.Errorf("Object %q still exists", name)
				}
			}
			return nil
		}
	*/
	return nil
}

func testAccPolicyComplianceCiImagesConfig(id string) string {
	var buf bytes.Buffer
	buf.Grow(500)

	buf.WriteString(fmt.Sprintf(`
resource "prismacloudcompute_policyComplianceCiImages" "test" {
    name = %q
}`, id))

	return buf.String()
}
