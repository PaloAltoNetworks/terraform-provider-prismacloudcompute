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

func TestAccPolicyRuntimeContainerConfig(t *testing.T) {
	fmt.Printf("\n\nStart TestAccPolicyConfig")
	var o policy.RuntimeContainerPolicy
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPolicyRuntimeContainerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyRuntimeContainerConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyRuntimeContainerExists("prismacloudcompute_policies_runtime_Container.test", &o),
					testAccCheckPolicyRuntimeContainerAttributes(&o),
				),
			},
			{
				Config: testAccPolicyRuntimeContainerConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyRuntimeContainerExists("prismacloudcompute_policies_runtime_Container.test", &o),
					testAccCheckPolicyRuntimeContainerAttributes(&o),
				),
			},
		},
	})
}

func TestAccPolicyRuntimeContainerNetwork(t *testing.T) {
	var o policy.RuntimeContainerPolicy
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPolicyRuntimeContainerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyRuntimeContainerConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyRuntimeContainerExists("prismacloudcompute_policies_runtime_Container.test", &o),
					testAccCheckPolicyRuntimeContainerAttributes(&o),
				),
			},
			{
				Config: testAccPolicyRuntimeContainerConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyRuntimeContainerExists("prismacloudcompute_policies_runtime_Container.test", &o),
					testAccCheckPolicyRuntimeContainerAttributes(&o),
				),
			},
		},
	})
}

func TestAccPolicyRuntimeContainerAuditEvent(t *testing.T) {
	var o policy.RuntimeContainerPolicy
	id := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPolicyRuntimeContainerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyRuntimeContainerConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyRuntimeContainerExists("prismacloudcompute_policies_runtime_Container.test", &o),
					testAccCheckPolicyRuntimeContainerAttributes(&o),
				),
			},
			{
				Config: testAccPolicyRuntimeContainerConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyRuntimeContainerExists("prismacloudcompute_policies_runtime_Container.test", &o),
					testAccCheckPolicyRuntimeContainerAttributes(&o),
				),
			},
		},
	})
}

func testAccCheckPolicyRuntimeContainerExists(n string, o *policy.RuntimeContainerPolicy) resource.TestCheckFunc {
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
		lo, err := policy.GetRuntimeContainer(*client)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}
		*o = lo

		return nil
	}
}

func testAccCheckPolicyRuntimeContainerAttributes(o *policy.RuntimeContainerPolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		return nil
	}
}

func testAccPolicyRuntimeContainerDestroy(s *terraform.State) error {
	/*	client := testAccProvider.Meta().(*api.Client)

		for _, rs := range s.RootModule().Resources {

			if rs.Type != "prismacloudcompute_policyruntimeContainer" {
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

func testAccPolicyRuntimeContainerConfig(id string) string {
	var buf bytes.Buffer
	buf.Grow(500)

	buf.WriteString(fmt.Sprintf(`
resource "prismacloudcompute_policyRuntimeContainer" "test" {
    name = %q
}`, id))

	return buf.String()
}
