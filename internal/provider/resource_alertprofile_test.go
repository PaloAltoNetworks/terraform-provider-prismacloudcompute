package provider

import (
	"fmt"
	"testing"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/alertprofile"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccAlertProfileWebhook(t *testing.T) {
	var o alertprofile.AlertProfile
	nameWebhook := fmt.Sprintf("test-%s", acctest.RandString(6))
	nameSlack := fmt.Sprintf("test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAlertProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAlertProfileWebhookConfig(nameWebhook),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAlertProfileExists("prismacloudcompute_alertprofile.test_webhook", &o),
					testAccCheckAlertProfileWebhookAttributes(&o, nameWebhook),
				),
			},
			{
				Config: testAccAlertProfileSlackConfig(nameSlack),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAlertProfileExists("prismacloudcompute_alertprofile.test_slack", &o),
					testAccCheckAlertProfileWebhookAttributes(&o, nameSlack),
				),
			},
		},
	})
}

func testAccCheckAlertProfileExists(n string, o *alertprofile.AlertProfile) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object Name is not set")
		}

		client := testAccProvider.Meta().(*api.Client)
		id := rs.Primary.ID
		lo, err := alertprofile.GetAlertprofile(*client, id)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}

		*o = *lo

		return nil
	}
}

func testAccCheckAlertProfileWebhookAttributes(o *alertprofile.AlertProfile, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.Name != name {
			return fmt.Errorf("\nName is %s, expected %s", o.Name, name)
		} else {
			fmt.Printf("\nName is %s", o.Name)
		}

		return nil
	}
}

func testAccCheckAlertProfileDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*api.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type != "prismacloudcompute_alertprofile" {
			continue
		}

		if rs.Primary.ID != "" {
			id := rs.Primary.ID
			if err := alertprofile.DeleteAlertprofile(*client, id); err == nil {
				return fmt.Errorf("Object %q still exists", id)
			}
		}
		return nil
	}

	return nil
}

func testAccAlertProfileSlackConfig(name string) string {
	return fmt.Sprintf(`	
	resource "prismacloudcompute_alertprofile" "test_slack" {
		name               = "%s"
		enable_immediate_vulnerabilities_alerts = false
	
		slack {
			webhook_url = "https://webhook.url"
		}
	
		policy {
		  admission {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  agentless_app_firewall {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  app_embedded_app_firewall {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  app_embedded_runtime {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  container_app_firewall {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  container_compliance_scan {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  container_vulnerability {
			enabled   = true
			all_rules = true
			rules = []
		  }
	
		  defender {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  host_app_firewall {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  host_compliance_scan {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  host_runtime {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  incident {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  kubernetes_audit {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  network_firewall {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  registry_vulnerability {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  serverless_app_firewall {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  serverless_runtime {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  waas_health {
			enabled = true
			all_rules = true
			rules = []
		  }
		}
	}
	`, name)
}

func testAccAlertProfileWebhookConfig(name string) string {
	return fmt.Sprintf(`	
	resource "prismacloudcompute_alertprofile" "test_webhook" {
		name               = "%s"
		enable_immediate_vulnerabilities_alerts = false
	
		webhook {
			url = "https://webhook.url"
			credential_id = ""
			custom_json = <<-EOT
						{
						"types": "#type",
						"time": "#time",
						"container": "#container",
						"image": "#image",
						"host": "#host",
						"fqdn": "#fqdn",
						"function": "#function",
						"region": "#region",
						"runtime": "#runtime",
						"appID": "#appID",
						"rule": "#rule",
						"message": "#message",
						"aggregated": "#aggregated",
						"rest": "#rest",
						"forensics": "#forensics",
						"accountID": "#accountID",
						"cluster": "#cluster",
						"labels": #labels,
						"collections": #collections,
						"complianceIssues": #complianceIssues,
						"vulnerabilities": #vulnerabilities
						}
					EOT
		}
	
		policy {
		  admission {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  agentless_app_firewall {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  app_embedded_app_firewall {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  app_embedded_runtime {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  container_app_firewall {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  container_compliance_scan {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  container_vulnerability {
			enabled   = true
			all_rules = true
			rules = []
		  }
	
		  defender {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  host_app_firewall {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  host_compliance_scan {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  host_runtime {
			enabled = true
			all_rules = true
			rules = []
		  }
		  
		  host_vulnerability {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  incident {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  kubernetes_audit {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  network_firewall {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  registry_vulnerability {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  serverless_app_firewall {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  serverless_runtime {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  vm_compliance {
			enabled = true
			all_rules = true
			rules = []
		  }
	
		  waas_health {
			enabled = true
			all_rules = true
			rules = []
		  }
		}
	}
	`, name)
}
