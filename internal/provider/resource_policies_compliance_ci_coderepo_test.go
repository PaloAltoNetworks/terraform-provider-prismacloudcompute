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

func TestAccPolicyComplianceCiCoderepoConfig(t *testing.T) {
	fmt.Printf("\n\nStart TestAccPolicyComplianceCiCiCoderepoConfig")
	var o policy.ComplianceCoderepoPolicy
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPolicyComplianceCiCoderepoDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyComplianceCiCoderepoConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyComplianceCiCoderepoExists("prismacloudcompute_policies_compliance_coderepo.test", &o),
					testAccCheckPolicyComplianceCiCoderepoAttributes(&o, id, "network"),
				),
			},
			{
				Config: testAccPolicyComplianceCiCoderepoConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyComplianceCiCoderepoExists("prismacloudcompute_policies_compliance_coderepo.test", &o),
					testAccCheckPolicyComplianceCiCoderepoAttributes(&o, id, "network"),
				),
			},
		},
	})
}

func TestAccPolicyComplianceCiCoderepoNetwork(t *testing.T) {
	var o policy.ComplianceCoderepoPolicy
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPolicyComplianceCiCoderepoDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyComplianceCiCoderepoConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyComplianceCiCoderepoExists("prismacloudcompute_policies_compliance_coderepo.test", &o),
					testAccCheckPolicyComplianceCiCoderepoAttributes(&o, id, "network"),
				),
			},
			{
				Config: testAccPolicyComplianceCiCoderepoConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyComplianceCiCoderepoExists("prismacloudcompute_policies_compliance_coderepo.test", &o),
					testAccCheckPolicyComplianceCiCoderepoAttributes(&o, id, "network"),
				),
			},
		},
	})
}

func TestAccPolicyComplianceCiCoderepoAuditEvent(t *testing.T) {
	var o policy.ComplianceCoderepoPolicy
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPolicyComplianceCiCoderepoDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyComplianceCiCoderepoConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyComplianceCiCoderepoExists("prismacloudcompute_policies_compliance_coderepo.test", &o),
					testAccCheckPolicyComplianceCiCoderepoAttributes(&o, id, "network"),
				),
			},
			{
				Config: testAccPolicyComplianceCiCoderepoConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyComplianceCiCoderepoExists("prismacloudcompute_policies_compliance_coderepo.test", &o),
					testAccCheckPolicyComplianceCiCoderepoAttributes(&o, id, "network"),
				),
			},
		},
	})
}

func testAccCheckPolicyComplianceCiCoderepoExists(n string, o *policy.ComplianceCoderepoPolicy) resource.TestCheckFunc {
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
		lo, err := policy.GetComplianceCoderepo(*client)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}
		*o = lo

		return nil
	}
}

func testAccCheckPolicyComplianceCiCoderepoAttributes(o *policy.ComplianceCoderepoPolicy, id string, policyType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.Id != id {
			return fmt.Errorf("\n\nPolicyId is %s, expected %s", o.Id, id)
		} else {
			fmt.Printf("\n\nName is %s", o.Id)
		}

		if o.Type != policyType {
			return fmt.Errorf("PolicyType is %s, expected %s", o.Type, policyType)
		}

		return nil
	}
}

func testAccPolicyComplianceCiCoderepoDestroy(s *terraform.State) error {
	/*	client := testAccProvider.Meta().(*api.Client)

		for _, rs := range s.RootModule().Resources {

			if rs.Type != "prismacloudcompute_policycompliancecoderepo" {
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

func testAccPolicyComplianceCiCoderepoConfig(id string) string {
	var buf bytes.Buffer
	buf.Grow(500)

	buf.WriteString(fmt.Sprintf(`
resource "prismacloudcompute_policyComplianceCiCoderepo" "test" {
    name = %q
}`, id))

	return buf.String()
}
