package provider

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/policy"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
// )

// func TestPolicyRuntimeHostConfig(t *testing.T) {
// 	fmt.Printf("\n\nStart TestAccPolicyConfig")
// 	var o policy.Policy
// 	id := fmt.Sprintf("tf%s", acctest.RandString(6))

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccPolicyDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccPolicyConfig(id),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckPolicyExists("prismacloudcompute_policies_runtime_container.test", &o),
// 					testAccCheckPolicyAttributes(&o, id, true),
// 				),
// 			},
// 			{
// 				Config: testAccPolicyConfig(id),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckPolicyExists("prismacloudcompute_policies_runtime_container.test", &o),
// 					testAccCheckPolicyAttributes(&o, id, true),
// 				),
// 			},
// 		},
// 	})
// }

// func TestPolicyRuntimeHostNetwork(t *testing.T) {
// 	var o policy.Policy
// 	id := fmt.Sprintf("tf%s", acctest.RandString(6))

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccPolicyDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccPolicyConfig(id),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckPolicyExists("prismacloudcompute_policies_runtime_container.test", &o),
// 					testAccCheckPolicyAttributes(&o, id, true),
// 				),
// 			},
// 			{
// 				Config: testAccPolicyConfig(id),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckPolicyExists("prismacloudcompute_policies_runtime_container.test", &o),
// 					testAccCheckPolicyAttributes(&o, id, true),
// 				),
// 			},
// 		},
// 	})
// }

// func TestPolicyRuntimeHostAuditEvent(t *testing.T) {
// 	var o policy.Policy
// 	id := fmt.Sprintf("tf%s", acctest.RandString(6))

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccPolicyDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccPolicyConfig(id),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckPolicyExists("prismacloudcompute_policies_runtime_container.test", &o),
// 					testAccCheckPolicyAttributes(&o, id, true),
// 				),
// 			},
// 			{
// 				Config: testAccPolicyConfig(id),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckPolicyExists("prismacloudcompute_policies_runtime_container.test", &o),
// 					testAccCheckPolicyAttributes(&o, id, true),
// 				),
// 			},
// 		},
// 	})
// }

// // func testPolicyRuntimeHostExists(n string, o *policy.Policy) resource.TestCheckFunc {
// // 	return func(s *terraform.State) error {
// // 		// return fmt.Errorf("What is the name: %s", o.PolicyId)

// // 		rs, ok := s.RootModule().Resources[n]
// // 		if !ok {
// // 			return fmt.Errorf("Resource not found: %s", n)
// // 		}

// // 		if rs.Primary.ID == "" {
// // 			return fmt.Errorf("Object label Id is not set")
// // 		}

// // 		client := testAccProvider.Meta().(*api.Client)
// // 		lo, err := policy.Get(*client, policy.RuntimeHostEndpoint)
// // 		if err != nil {
// // 			return fmt.Errorf("Error in get: %s", err)
// // 		}
// // 		*o = lo

// // 		return nil
// // 	}
// // }

// // func testPolicyRuntimeHostAttributes(o *policy.Policy, id string, learningDisabled bool) resource.TestCheckFunc {
// // 	return func(s *terraform.State) error {
// // 		if o.PolicyId != id {
// // 			return fmt.Errorf("\n\nPolicyId is %s, expected %s", o.PolicyId, id)
// // 		} else {
// // 			fmt.Printf("\n\nName is %s", o.PolicyId)
// // 		}

// // 		if o.LearningDisabled != learningDisabled {
// // 			return fmt.Errorf("LearningDisabled is %t, expected %t", o.LearningDisabled, learningDisabled)
// // 		}

// // 		return nil
// // 	}
// // }

// // func testPolicyRuntimeHostDestroy(s *terraform.State) error {
// // 	/*	client := testAccProvider.Meta().(*api.Client)

// // 		for _, rs := range s.RootModule().Resources {

// // 			if rs.Type != "prismacloudcompute_policyruntimehost" {
// // 				continue
// // 			}

// // 			if rs.Primary.ID != "" {
// // 				name := rs.Primary.ID
// // 				if err := policy.Delete(client, name); err == nil {
// // 					return fmt.Errorf("Object %q still exists", name)
// // 				}
// // 			}
// // 			return nil
// // 		}
// // 	*/
// // 	return nil
// // }

// // func testPolicyRuntimeHostConfig(id string) string {
// // 	var buf bytes.Buffer
// // 	buf.Grow(500)

// // 	buf.WriteString(fmt.Sprintf(`
// // resource "prismacloudcompute_policyRuntimeHost" "test" {
// //     name = %q
// // }`, id))

// // 	return buf.String()
// // }
