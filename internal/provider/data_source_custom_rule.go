package provider

import (
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/rule"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCustomRule() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve ID of a custom rule.",
		Read:        dataSourceCustomRuleRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "ID of the custom rule.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"prisma_id": {
				Description: "Prisma Cloud Compute ID of the custom rule.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Free-form text description of the custom rule.",
			},
			"message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Message to display for a custom rule event.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique custom rule name.",
			},
			"script": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Custom rule expression.",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Custom rule type. Can be set to 'processes', 'filesystem', 'network-outgoing', 'kubernetes-audit', 'waas-request', or 'waas-response'.",
			},
		},
	}
}

func dataSourceCustomRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)

	if name := d.Get("name").(string); name != "" {
		retrievedCustomRule, err := rule.GetCustomRuleByName(*client, name)
		if err != nil {
			return fmt.Errorf("error reading custom rule: %s", err)
		}
		if err := d.Set("description", retrievedCustomRule.Description); err != nil {
			return fmt.Errorf("error reading custom rule: %s", err)
		}
		if err := d.Set("prisma_id", retrievedCustomRule.Id); err != nil {
			return fmt.Errorf("error reading custom rule: %s", err)
		}
		if err := d.Set("message", retrievedCustomRule.Message); err != nil {
			return fmt.Errorf("error reading custom rule: %s", err)
		}
		if err := d.Set("name", retrievedCustomRule.Name); err != nil {
			return fmt.Errorf("error reading custom rule: %s", err)
		}
		if err := d.Set("script", retrievedCustomRule.Script); err != nil {
			return fmt.Errorf("error reading custom rule: %s", err)
		}
		if err := d.Set("type", retrievedCustomRule.Type); err != nil {
			return fmt.Errorf("error reading custom rule: %s", err)
		}
		d.SetId(retrievedCustomRule.Name)

		return nil
	}

	return fmt.Errorf("missing name parameter")
}
