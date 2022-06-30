package provider

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/policy"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCustomComplianceConfig(t *testing.T) {
	fmt.Printf("\n\nStart TestAccCustomComplianceConfig")
	var o policy.CustomCompliance
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCustomComplianceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomComplianceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomComplianceExists("prismacloudcompute_custom_Compliance.test", &o),
					testAccCheckCustomComplianceAttributes(&o, name, "description", "#000000"),
				),
			},
			{
				Config: testAccCustomComplianceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomComplianceExists("prismacloudcompute_custom_Compliance.test", &o),
					testAccCheckCustomComplianceAttributes(&o, name, "description", "#000000"),
				),
			},
		},
	})
}

func TestAccCustomComplianceNetwork(t *testing.T) {
	var o policy.CustomCompliance
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCustomComplianceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomComplianceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomComplianceExists("prismacloudcompute_custom_Compliance.test", &o),
					testAccCheckCustomComplianceAttributes(&o, name, "description", "#000000"),
				),
			},
			{
				Config: testAccCustomComplianceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomComplianceExists("prismacloudcompute_custom_Compliance.test", &o),
					testAccCheckCustomComplianceAttributes(&o, name, "description", "#000000"),
				),
			},
		},
	})
}

func TestAccCustomComplianceAuditEvent(t *testing.T) {
	var o policy.CustomCompliance
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCustomComplianceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomComplianceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomComplianceExists("prismacloudcompute_custom_Compliance.test", &o),
					testAccCheckCustomComplianceAttributes(&o, name, "description", "#000000"),
				),
			},
			{
				Config: testAccCustomComplianceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomComplianceExists("prismacloudcompute_custom_Compliance.test", &o),
					testAccCheckCustomComplianceAttributes(&o, name, "description", "#000000"),
				),
			},
		},
	})
}

func testAccCheckCustomComplianceExists(n string, o *policy.CustomCompliance) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// return fmt.Errorf("What is the name: %s", o.Name)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object label Name is not set")
		}

		client := testAccProvider.Meta().(*api.Client)
		name := rs.Primary.ID
		lo, err := policy.GetCustomComplianceByName(*client, name)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}
		o = lo

		return nil
	}
}

func testAccCheckCustomComplianceAttributes(o *policy.CustomCompliance, name string, description string, color string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.Name != name {
			return fmt.Errorf("\n\nName is %s, expected %s", o.Name, name)
		} else {
			fmt.Printf("\n\nName is %s", o.Name)
		}

		return nil
	}
}

func testAccCustomComplianceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*api.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type != "prismacloudcompute_custom_Compliance" {
			continue
		}

		if rs.Primary.ID != "" {
			id, _ := strconv.Atoi(rs.Primary.ID)
			if err := policy.DeleteCustomCompliance(*client, id); err == nil {
				return fmt.Errorf("Object %q still exists", id)
			}
		}
		return nil
	}

	return nil
}

func testAccCustomComplianceConfig(name string) string {
	var buf bytes.Buffer
	buf.Grow(500)

	buf.WriteString(fmt.Sprintf(`
resource "prismacloudcompute_custom_Compliances" "test" {
    name = %q
}`, name))

	return buf.String()
}
