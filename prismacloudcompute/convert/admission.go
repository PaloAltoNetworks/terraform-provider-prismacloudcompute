package convert

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy"
)

func SchemaToAdmissionRules(d *schema.ResourceData) ([]policy.AdmissionRule, error) {
	parsedRules := make([]policy.AdmissionRule, 0)
	if rules, ok := d.GetOk("rule"); ok {
		presentRules := rules.([]interface{})
		for _, val := range presentRules {
			presentRule := val.(map[string]interface{})
			parsedRule := policy.AdmissionRule{}

			parsedRule.Description = presentRule["description"].(string)
			parsedRule.Disabled = presentRule["disabled"].(bool)
			parsedRule.Effect = presentRule["effect"].(string)
			parsedRule.Name = presentRule["name"].(string)
			parsedRule.Script = presentRule["script"].(string)

			parsedRules = append(parsedRules, parsedRule)
		}
	}
	return parsedRules, nil
}

func AdmissionRulesToSchema(in []policy.AdmissionRule) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["description"] = val.Description
		m["disabled"] = val.Disabled
		m["effect"] = val.Effect
		m["name"] = val.Name
		m["script"] = val.Script
		ans = append(ans, m)
	}
	return ans
}
