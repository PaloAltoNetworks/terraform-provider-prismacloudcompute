package provider

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/rule"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCustomRuleConfig(t *testing.T) {
	fmt.Printf("\n\nStart TestAccCustomRuleConfig")
	var o rule.CustomRule
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCustomRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomRuleConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomRuleExists("prismacloudcompute_custom_rule.test", &o),
					testAccCheckCustomRuleAttributes(&o, name, "description", "#000000"),
				),
			},
			{
				Config: testAccCustomRuleConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomRuleExists("prismacloudcompute_custom_rule.test", &o),
					testAccCheckCustomRuleAttributes(&o, name, "description", "#000000"),
				),
			},
		},
	})
}

func TestAccCustomRuleNetwork(t *testing.T) {
	var o rule.CustomRule
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCustomRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomRuleConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomRuleExists("prismacloudcompute_custom_rule.test", &o),
					testAccCheckCustomRuleAttributes(&o, name, "description", "#000000"),
				),
			},
			{
				Config: testAccCustomRuleConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomRuleExists("prismacloudcompute_custom_rule.test", &o),
					testAccCheckCustomRuleAttributes(&o, name, "description", "#000000"),
				),
			},
		},
	})
}

func TestAccCustomRuleAuditEvent(t *testing.T) {
	var o rule.CustomRule
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCustomRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomRuleConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomRuleExists("prismacloudcompute_custom_rule.test", &o),
					testAccCheckCustomRuleAttributes(&o, name, "description", "#000000"),
				),
			},
			{
				Config: testAccCustomRuleConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomRuleExists("prismacloudcompute_custom_rule.test", &o),
					testAccCheckCustomRuleAttributes(&o, name, "description", "#000000"),
				),
			},
		},
	})
}

func testAccCheckCustomRuleExists(n string, o *rule.CustomRule) resource.TestCheckFunc {
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
		lo, err := rule.GetCustomRuleByName(*client, name)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}
		o = lo

		return nil
	}
}

func testAccCheckCustomRuleAttributes(o *rule.CustomRule, name string, description string, typeOfCustomRuleAttribute string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.Name != name {
			return fmt.Errorf("\n\nName is %s, expected %s", o.Name, name)
		} else {
			fmt.Printf("\n\nName is %s", o.Name)
		}

		if o.Description != description {
			return fmt.Errorf("Description is %s, expected %s", o.Description, description)
		}

		if o.Type != typeOfCustomRuleAttribute {
			return fmt.Errorf("Type is %q, expected %q", o.Type, typeOfCustomRuleAttribute)
		}

		return nil
	}
}

func testAccCustomRuleDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*api.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type != "prismacloudcompute_custom_rule" {
			continue
		}

		if rs.Primary.ID != "" {
			id, _ := strconv.Atoi(rs.Primary.ID)
			if err := rule.DeleteCustomRule(*client, id); err == nil {
				return fmt.Errorf("Object %q still exists", id)
			}
		}
		return nil
	}

	return nil
}

func testAccCustomRuleConfig(name string) string {
	var buf bytes.Buffer
	buf.Grow(500)

	buf.WriteString(fmt.Sprintf(`
resource "prismacloudcompute_custom_rules" "test" {
    name = %q
}`, name))

	return buf.String()
}
