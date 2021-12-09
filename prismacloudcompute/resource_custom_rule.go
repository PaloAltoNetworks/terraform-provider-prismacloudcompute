package prismacloudcompute

import (
	"fmt"
	"strconv"

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
				Description: "The ID of the custom rule.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A free-form text description of the custom rule.",
			},
			"message": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Message to display for a custom rule event.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A unique custom rule name.",
			},
			"script": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An custom rule expression.",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The custom rule type. Can be set to 'processes', 'filesystem', 'network-outgoing', 'kubernetes-audit', 'waas-request', 'waas-response'",
			},
		},
	}
}

func createCustomRule(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedCustomRule := convert.SchemaToCustomRule(d)
	if err := rule.CreateCustomRule(*client, parsedCustomRule); err != nil {
		return fmt.Errorf("error creating custom rule '%+v': %s", parsedCustomRule, err)
	}

	d.SetId(strconv.Itoa(parsedCustomRule.Id))
	return readCustomRule(d, meta)
}

func readCustomRule(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("error reading custom rule: %s", err)
	}
	retrievedCustomRule, err := rule.GetCustomRule(*client, id)
	if err != nil {
		return fmt.Errorf("error reading custom rule: %s", err)
	}

	if err := d.Set("description", retrievedCustomRule.Description); err != nil {
		return fmt.Errorf("error reading custom rule: %s", err)
	}
	if err := d.Set("id", retrievedCustomRule.Id); err != nil {
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
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("error reading custom rule: %s", err)
	}
	if err := rule.DeleteCustomRule(*client, id); err != nil {
		return fmt.Errorf("error updating custom rule '%s': %s", d.Id(), err)
	}

	d.SetId("")
	return nil
}
