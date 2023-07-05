package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/rule"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCustomRule() *schema.Resource {
	return &schema.Resource{
		CreateContext: createCustomRule,
		ReadContext:   readCustomRule,
		UpdateContext: updateCustomRule,
		DeleteContext: deleteCustomRule,

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
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the custom rule.",
			},
			"prisma_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Prisma Cloud Compute ID of the custom rule.",
			},
			"attack_techniques": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Attack techniques from the MITRE ATT&CK Framework that the rule is concerned with.",
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
			"min_version": {
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
			"vuln_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "List of vulnerability IDs.",
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

func createCustomRule(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)
	parsedCustomRule := convert.SchemaToCustomRule(d)
	id, err := rule.CreateCustomRule(*client, parsedCustomRule)

	if err != nil {
		return diag.Errorf("error creating custom rule '%+v': %s", parsedCustomRule, err)
	}
	if err := d.Set("prisma_id", id); err != nil {
		return diag.Errorf("error creating custom rule '%+v': %s", parsedCustomRule, err)
	}
	d.SetId(parsedCustomRule.Name)
	return readCustomRule(ctx, d, meta)
}

func readCustomRule(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)

	var diags diag.Diagnostics

	retrievedCustomRule, err := rule.GetCustomRuleByName(*client, d.Id())

	if err != nil {
		return diag.Errorf("error reading custom rule: %s", err)
	}

	if err := d.Set("description", retrievedCustomRule.Description); err != nil {
		return diag.Errorf("error reading custom rule: %s", err)
	}
	if err := d.Set("prisma_id", retrievedCustomRule.Id); err != nil {
		return diag.Errorf("error reading custom rule: %s", err)
	}
	if err := d.Set("message", retrievedCustomRule.Message); err != nil {
		return diag.Errorf("error reading custom rule: %s", err)
	}
	if err := d.Set("name", retrievedCustomRule.Name); err != nil {
		return diag.Errorf("error reading custom rule: %s", err)
	}
	if err := d.Set("script", retrievedCustomRule.Script); err != nil {
		return diag.Errorf("error reading custom rule: %s", err)
	}
	if err := d.Set("type", retrievedCustomRule.Type); err != nil {
		return diag.Errorf("error reading custom rule: %s", err)
	}

	return diags
}

func updateCustomRule(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)
	parsedCustomRule := convert.SchemaToCustomRule(d)

	if err := rule.UpdateCustomRule(*client, parsedCustomRule); err != nil {
		return diag.Errorf("error updating custom rule: %s", err)
	}

	return readCustomRule(ctx, d, meta)
}

func deleteCustomRule(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)

	var diags diag.Diagnostics

	if err := rule.DeleteCustomRule(*client, d.Get("prisma_id").(int)); err != nil {
		return diag.Errorf("error updating custom rule '%s': %s", d.Id(), err)
	}

	d.SetId("")

	return diags
}
