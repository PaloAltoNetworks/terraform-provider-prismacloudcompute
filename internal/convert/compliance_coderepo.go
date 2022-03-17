package convert

import (
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/policy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaToComplianceCoderepoRules(d *schema.ResourceData) ([]policy.ComplianceCoderepoRule, error) {
	parsedRules := make([]policy.ComplianceCoderepoRule, 0)
	if rules, ok := d.GetOk("rule"); ok {
		presentRules := rules.([]interface{})
		for _, val := range presentRules {
			presentRule := val.(map[string]interface{})
			parsedRule := policy.ComplianceCoderepoRule{}

			if len(presentRule["license"].([]interface{})) > 0 && presentRule["license"].([]interface{})[0] != nil {
				presentLicense := presentRule["license"].([]interface{})[0].(map[string]interface{})
				if len(presentLicense["critical"].([]interface{})) > 0 && presentLicense["critical"].([]interface{})[0] != nil {
					parsedRule.License.Critical = SchemaToStringSlice(presentLicense["critical"].([]interface{}))
				}
				if len(presentLicense["high"].([]interface{})) > 0 && presentLicense["high"].([]interface{})[0] != nil {
					parsedRule.License.High = SchemaToStringSlice(presentLicense["high"].([]interface{}))
				}
				if len(presentLicense["medium"].([]interface{})) > 0 && presentLicense["medium"].([]interface{})[0] != nil {
					parsedRule.License.Medium = SchemaToStringSlice(presentLicense["medium"].([]interface{}))
				}
				if len(presentLicense["low"].([]interface{})) > 0 && presentLicense["low"].([]interface{})[0] != nil {
					parsedRule.License.Low = SchemaToStringSlice(presentLicense["low"].([]interface{}))
				}
				if presentLicense["alert_threshold"].([]interface{})[0] != nil {
					presentAlertThreshold := presentLicense["alert_threshold"].([]interface{})[0].(map[string]interface{})
					parsedRule.License.AlertThreshold = policy.ComplianceCoderepoThreshold{
						Enabled: presentAlertThreshold["enabled"].(bool),
						Value:   presentAlertThreshold["value"].(int),
					}
				}
			} else {
				parsedRule.License = policy.ComplianceCoderepoLicense{}
			}

			parsedRule.Collections = PolicySchemaToCollections(presentRule["collections"].([]interface{}))

			parsedRule.Disabled = presentRule["disabled"].(bool)
			parsedRule.Effect = presentRule["effect"].(string)
			parsedRule.Name = presentRule["name"].(string)
			parsedRule.Notes = presentRule["notes"].(string)

			parsedRules = append(parsedRules, parsedRule)
		}
	}
	return parsedRules, nil
}

func ComplianceCoderepoRulesToSchema(in []policy.ComplianceCoderepoRule) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["collections"] = CollectionsToPolicySchema(val.Collections)
		m["license"] = complianceCoderepoLicenseToSchema(val.License)
		m["disabled"] = val.Disabled
		m["effect"] = val.Effect
		m["name"] = val.Name
		m["notes"] = val.Notes
		ans = append(ans, m)
	}
	return ans
}

func complianceCoderepoThresholdToSchema(in policy.ComplianceCoderepoThreshold) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["enabled"] = in.Enabled
	m["value"] = in.Value
	ans = append(ans, m)
	return ans
}

func complianceCoderepoLicenseToSchema(in policy.ComplianceCoderepoLicense) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["alert_threshold"] = complianceCoderepoThresholdToSchema(in.AlertThreshold)
	m["critical"] = in.Critical
	m["high"] = in.High
	m["medium"] = in.Medium
	m["low"] = in.Low

	ans = append(ans, m)
	return ans
}
