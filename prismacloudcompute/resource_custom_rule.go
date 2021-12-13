package prismacloudcompute

import (
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/prismacloudcompute/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/rule"
)

func resourceCustomRule() *schema.Resource {
	return &schema.Resource{
		Create: createCustomRule,
		Read:   readCustomRule,
		Update: updateCustomRule,
		Delete: deleteCustomRule,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

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
				Optional:    true,
				Description: "Message to display for a custom rule event.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique custom rule name.",
			},
			"script": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Custom rule expression.",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Custom rule type. Can be set to 'processes', 'filesystem', 'network-outgoing', 'kubernetes-audit', 'waas-request', or 'waas-response'.",
			},
		},
	}
}

func createCustomRule(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedCustomRule := convert.SchemaToCustomRule(d)
	id, err := rule.CreateCustomRule(*client, parsedCustomRule)

	if err != nil {
		return fmt.Errorf("error creating custom rule '%+v': %s", parsedCustomRule, err)
	}
	if err := d.Set("prisma_id", id); err != nil {
		return fmt.Errorf("error creating custom rule '%+v': %s", parsedCustomRule, err)
	}
	d.SetId(parsedCustomRule.Name)
	return readCustomRule(d, meta)
}

func readCustomRule(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	retrievedCustomRule, err := rule.GetCustomRule(*client, d.Get("prisma_id").(int))
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

	return nil
}

func updateCustomRule(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedCustomRule := convert.SchemaToCustomRule(d)

	if err := rule.UpdateCustomRule(*client, parsedCustomRule); err != nil {
		return fmt.Errorf("error updating custom rule: %s", err)
	}

	return readCustomRule(d, meta)
}

func deleteCustomRule(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	if err := rule.DeleteCustomRule(*client, d.Get("prisma_id").(int)); err != nil {
		return fmt.Errorf("error updating custom rule '%s': %s", d.Id(), err)
	}

	d.SetId("")
	return nil
}
