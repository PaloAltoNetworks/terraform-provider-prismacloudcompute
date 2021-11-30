package prismacloudcompute

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPolicyAdmissionConfig(t *testing.T) {
	fmt.Printf("\n\nStart TestAccPolicyAdmissionConfig")
	var o policy.Policy
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPolicyAdmissionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyAdmissionConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyAdmissionExists("prismacloudcompute_policies_admission.test", &o),
					testAccCheckPolicyAdmissionAttributes(&o, id, true),
				),
			},
			{
				Config: testAccPolicyAdmissionConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyAdmissionExists("prismacloudcompute_policies_admission.test", &o),
					testAccCheckPolicyAdmissionAttributes(&o, id, true),
				),
			},
		},
	})
}

func TestAccPolicyAdmissionNetwork(t *testing.T) {
	var o policy.Policy
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPolicyAdmissionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyAdmissionConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyAdmissionExists("prismacloudcompute_policies_admission.test", &o),
					testAccCheckPolicyAdmissionAttributes(&o, id, true),
				),
			},
			{
				Config: testAccPolicyAdmissionConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyAdmissionExists("prismacloudcompute_policies_admission.test", &o),
					testAccCheckPolicyAdmissionAttributes(&o, id, true),
				),
			},
		},
	})
}

func TestAccPolicyAdmissionAuditEvent(t *testing.T) {
	var o policy.Policy
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPolicyAdmissionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyAdmissionConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyAdmissionExists("prismacloudcompute_policies_admission.test", &o),
					testAccCheckPolicyAdmissionAttributes(&o, id, true),
				),
			},
			{
				Config: testAccPolicyAdmissionConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyAdmissionExists("prismacloudcompute_policies_admission.test", &o),
					testAccCheckPolicyAdmissionAttributes(&o, id, true),
				),
			},
		},
	})
}

func testAccCheckPolicyAdmissionExists(n string, o *policy.Policy) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// return fmt.Errorf("What is the name: %s", o.PolicyId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object label Id is not set")
		}

		client := testAccProvider.Meta().(*pcc.Client)
		lo, err := policy.Get(*client, policy.ComplianceHostEndpoint)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}
		*o = lo

		return nil
	}
}

func testAccCheckPolicyAdmissionAttributes(o *policy.Policy, id string, policyType string) resource.TestCheckFunc {
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

func testAccPolicyAdmissionDestroy(s *terraform.State) error {
	/*	client := testAccProvider.Meta().(*pcc.Client)

		for _, rs := range s.RootModule().Resources {

			if rs.Type != "prismacloudcompute_policyadmission" {
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

func testAccPolicyAdmissionConfig(id string) string {
	var buf bytes.Buffer
	buf.Grow(500)

	buf.WriteString(fmt.Sprintf(`
resource "prismacloudcompute_policyAdmission" "test" {
    name = %q
}`, id))

	return buf.String()
}
