package convert

import (
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/rule"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Converts a custom rule schema to a custom rule object for SDK compatibility.
func SchemaToCustomRule(d *schema.ResourceData) rule.CustomRule {
	parsedRule := rule.CustomRule{}

	if val, ok := d.GetOk("prisma_id"); ok {
		parsedRule.Id = val.(int)
	}

	if val, ok := d.GetOk("attack_techniques"); ok {
		parsedRule.AttackTechniques = SchemaToStringSlice(val.([]interface{}))
	}

	if val, ok := d.GetOk("description"); ok {
		parsedRule.Description = val.(string)
	}

	if val, ok := d.GetOk("message"); ok {
		parsedRule.Message = val.(string)
	}

	if val, ok := d.GetOk("min_version"); ok {
		parsedRule.MinVersion = val.(string)
	}

	if val, ok := d.GetOk("name"); ok {
		parsedRule.Name = val.(string)
	}

	if val, ok := d.GetOk("script"); ok {
		parsedRule.Script = val.(string)
	}

	if val, ok := d.GetOk("type"); ok {
		parsedRule.Type = val.(string)
	}

	if val, ok := d.GetOk("vuln_ids"); ok {
		parsedRule.VulnIDs = SchemaToStringSlice(val.([]interface{}))
	}

	return parsedRule
}
