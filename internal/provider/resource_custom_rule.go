package provider

import (
	"fmt"
        "strings"
	"strconv"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/rule"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCustomRule() *schema.Resource {
	return &schema.Resource{
		Create: createCustomRule,
		Read:   readCustomRule,
		Update: updateCustomRule,
		Delete: deleteCustomRule,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				
				name, id, err := CustomRuleParseId(d.Id())

				intVar, err := strconv.Atoi(id)

				if err != nil {
        				return []*schema.ResourceData{d}, nil
        				      }

				var pid int = intVar
        			d.Set("prisma_id", pid)
				d.SetId(name)
        		return []*schema.ResourceData{d}, nil
      		},
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


func CustomRuleParseId(id string) (string, string, error) {
  parts := strings.SplitN(id, ":", 2)

  if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
    return "", "", fmt.Errorf("unexpected format of ID (%s), expected attribute1:attribute2", id)
  }

  return parts[0], parts[1], nil
}

func createCustomRule(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
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
	client := meta.(*api.Client)
	retrievedCustomRule, err := rule.GetCustomRuleById(*client, d.Get("prisma_id").(int))
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
	client := meta.(*api.Client)
	parsedCustomRule := convert.SchemaToCustomRule(d)

	if err := rule.UpdateCustomRule(*client, parsedCustomRule); err != nil {
		return fmt.Errorf("error updating custom rule: %s", err)
	}

	return readCustomRule(d, meta)
}

func deleteCustomRule(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	if err := rule.DeleteCustomRule(*client, d.Get("prisma_id").(int)); err != nil {
		return fmt.Errorf("error updating custom rule '%s': %s", d.Id(), err)
	}

	d.SetId("")
	return nil
}
