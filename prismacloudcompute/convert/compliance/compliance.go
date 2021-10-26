package compliance

import (
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/prismacloudcompute/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/collection"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy"
)

func SchemaToComplianceCiRules(d *schema.ResourceData) ([]policy.ComplianceRule, error) {
	parsedRules := make([]policy.ComplianceRule, 0)
	if rules, ok := d.GetOk("rule"); ok {
		presentRules := rules.([]interface{})
		for _, val := range presentRules {
			presentRule := val.(map[string]interface{})
			parsedRule := policy.ComplianceRule{}

			parsedRule.Collections = convert.PolicySchemaToCollections(presentRule["collections"].([]interface{}))

			presentChecks := presentRule["compliance_check"].([]interface{})
			parsedConditions := policy.ComplianceConditions{
				Checks: make([]policy.ComplianceCheck, 0, len(presentChecks)),
			}
			for _, val := range presentChecks {
				presentCheck := val.(map[string]interface{})
				parsedConditions.Checks = append(parsedConditions.Checks, policy.ComplianceCheck{
					Block: presentCheck["block"].(bool),
					Id:    presentCheck["id"].(int),
				})
			}
			parsedRule.Conditions = parsedConditions

			parsedRule.Disabled = presentRule["disabled"].(bool)
			parsedRule.Effect = presentRule["effect"].(string)
			parsedRule.Name = presentRule["name"].(string)
			parsedRule.Notes = presentRule["notes"].(string)
			parsedRule.Verbose = presentRule["verbose"].(bool)

			parsedRules = append(parsedRules, parsedRule)
		}
	}
	return parsedRules, nil
}

func SchemaToComplianceDeployedRules(d *schema.ResourceData) ([]policy.ComplianceRule, error) {
	parsedRules := make([]policy.ComplianceRule, 0)
	if rules, ok := d.GetOk("rule"); ok {
		presentRules := rules.([]interface{})
		for _, val := range presentRules {
			presentRule := val.(map[string]interface{})
			parsedRule := policy.ComplianceRule{}

			parsedRule.BlockMessage = presentRule["block_message"].(string)

			presentCollections := presentRule["collections"].([]interface{})
			parsedCollections := make([]collection.Collection, 0, len(presentCollections))
			for _, val := range presentCollections {
				parsedCollection := collection.Collection{
					Name: val.(string),
				}
				parsedCollections = append(parsedCollections, parsedCollection)
			}
			parsedRule.Collections = parsedCollections

			presentChecks := presentRule["compliance_check"].([]interface{})
			parsedConditions := policy.ComplianceConditions{
				Checks: make([]policy.ComplianceCheck, 0, len(presentChecks)),
			}
			for _, val := range presentChecks {
				presentCheck := val.(map[string]interface{})
				parsedConditions.Checks = append(parsedConditions.Checks, policy.ComplianceCheck{
					Block: presentCheck["block"].(bool),
					Id:    presentCheck["id"].(int),
				})
			}
			parsedRule.Conditions = parsedConditions

			parsedRule.Disabled = presentRule["disabled"].(bool)
			parsedRule.Effect = presentRule["effect"].(string)
			parsedRule.Name = presentRule["name"].(string)
			parsedRule.Notes = presentRule["notes"].(string)
			parsedRule.ShowPassedChecks = presentRule["show_passed_checks"].(bool)
			parsedRule.Verbose = presentRule["verbose"].(bool)

			parsedRules = append(parsedRules, parsedRule)
		}
	}
	return parsedRules, nil
}

func ComplianceCiRulesToSchema(in []policy.ComplianceRule) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["collections"] = convert.CollectionsToPolicySchema(val.Collections)
		m["compliance_check"] = complianceConditionsToSchema(val.Conditions)
		m["disabled"] = val.Disabled
		m["effect"] = val.Effect
		m["name"] = val.Name
		m["notes"] = val.Notes
		m["verbose"] = val.Verbose
		ans = append(ans, m)
	}
	return ans
}

func ComplianceDeployedRulesToSchema(in []policy.ComplianceRule) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["block_message"] = val.BlockMessage
		m["collections"] = convert.CollectionsToPolicySchema(val.Collections)
		m["compliance_check"] = complianceConditionsToSchema(val.Conditions)
		m["disabled"] = val.Disabled
		m["effect"] = val.Effect
		m["name"] = val.Name
		m["notes"] = val.Notes
		m["show_passed_checks"] = val.ShowPassedChecks
		m["verbose"] = val.Verbose
		ans = append(ans, m)
	}
	return ans
}

func complianceConditionsToSchema(in policy.ComplianceConditions) []interface{} {
	ans := make([]interface{}, 0, len(in.Checks))
	for _, val := range in.Checks {
		m := make(map[string]interface{})
		m["block"] = val.Block
		m["id"] = val.Id
		ans = append(ans, m)
	}
	return ans
}
