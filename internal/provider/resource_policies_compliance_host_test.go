package provider

// import (
// 	"bytes"
// 	"fmt"
// 	"testing"

// 	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
// 	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/policy"

// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
// )

// func TestAccPolicyComplianceHostConfig(t *testing.T) {
// 	fmt.Printf("\n\nStart TestAccPolicyComplianceHostConfig")
// 	var o policy.Policy
// 	id := fmt.Sprintf("tf%s", acctest.RandString(6))

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccPolicyComplianceHostDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccPolicyComplianceHostConfig(id),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckPolicyComplianceHostExists("prismacloudcompute_policies_compliance_container.test", &o),
// 					testAccCheckPolicyComplianceHostAttributes(&o, id, "network"),
// 				),
// 			},
// 			{
// 				Config: testAccPolicyComplianceHostConfig(id),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckPolicyComplianceHostExists("prismacloudcompute_policies_compliance_container.test", &o),
// 					testAccCheckPolicyComplianceHostAttributes(&o, id, "network"),
// 				),
// 			},
// 		},
// 	})
// }

// func TestAccPolicyComplianceHostNetwork(t *testing.T) {
// 	var o policy.Policy
// 	id := fmt.Sprintf("tf%s", acctest.RandString(6))

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccPolicyComplianceHostDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccPolicyComplianceHostConfig(id),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckPolicyComplianceHostExists("prismacloudcompute_policies_compliance_container.test", &o),
// 					testAccCheckPolicyComplianceHostAttributes(&o, id, "network"),
// 				),
// 			},
// 			{
// 				Config: testAccPolicyComplianceHostConfig(id),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckPolicyComplianceHostExists("prismacloudcompute_policies_compliance_container.test", &o),
// 					testAccCheckPolicyComplianceHostAttributes(&o, id, "network"),
// 				),
// 			},
// 		},
// 	})
// }

// func TestAccPolicyComplianceHostAuditEvent(t *testing.T) {
// 	var o policy.Policy
// 	id := fmt.Sprintf("tf%s", acctest.RandString(6))

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccPolicyComplianceHostDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccPolicyComplianceHostConfig(id),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckPolicyComplianceHostExists("prismacloudcompute_policies_compliance_container.test", &o),
// 					testAccCheckPolicyComplianceHostAttributes(&o, id, "network"),
// 				),
// 			},
// 			{
// 				Config: testAccPolicyComplianceHostConfig(id),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckPolicyComplianceHostExists("prismacloudcompute_policies_compliance_container.test", &o),
// 					testAccCheckPolicyComplianceHostAttributes(&o, id, "network"),
// 				),
// 			},
// 		},
// 	})
// }

// func testAccCheckPolicyComplianceHostExists(n string, o *policy.Policy) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		// return fmt.Errorf("What is the name: %s", o.PolicyId)

// 		rs, ok := s.RootModule().Resources[n]
// 		if !ok {
// 			return fmt.Errorf("Resource not found: %s", n)
// 		}

// 		if rs.Primary.ID == "" {
// 			return fmt.Errorf("Object label Id is not set")
// 		}

// 		client := testAccProvider.Meta().(*api.Client)
// 		lo, err := policy.Get(*client, policy.ComplianceHostEndpoint)
// 		if err != nil {
// 			return fmt.Errorf("Error in get: %s", err)
// 		}
// 		*o = lo

// 		return nil
// 	}
// }

// func testAccCheckPolicyComplianceHostAttributes(o *policy.Policy, id string, policyType string) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		if o.PolicyId != id {
// 			return fmt.Errorf("\n\nPolicyId is %s, expected %s", o.PolicyId, id)
// 		} else {
// 			fmt.Printf("\n\nName is %s", o.PolicyId)
// 		}

// 		if o.PolicyType != policyType {
// 			return fmt.Errorf("PolicyType is %s, expected %s", o.PolicyType, policyType)
// 		}

// 		return nil
// 	}
// }

// func testAccPolicyComplianceHostDestroy(s *terraform.State) error {
// 	/*	client := testAccProvider.Meta().(*api.Client)

// 		for _, rs := range s.RootModule().Resources {

// 			if rs.Type != "prismacloudcompute_policycompliancehost" {
// 				continue
// 			}

// 			if rs.Primary.ID != "" {
// 				name := rs.Primary.ID
// 				if err := policy.Delete(client, name); err == nil {
// 					return fmt.Errorf("Object %q still exists", name)
// 				}
// 			}
// 			return nil
// 		}
// 	*/
// 	return nil
// }

// func testAccPolicyComplianceHostConfig(id string) string {
// 	var buf bytes.Buffer
// 	buf.Grow(500)

// 	buf.WriteString(fmt.Sprintf(`
// resource "prismacloudcompute_policyComplianceHost" "test" {
//     name = %q
// }`, id))

// 	return buf.String()
// }
