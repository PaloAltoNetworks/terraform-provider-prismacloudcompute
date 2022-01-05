package rule

import (
	"fmt"
	"net/http"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
)

const CustomRulesEndpoint = "api/v1/custom-rules"

type CustomRule struct {
	Description string `json:"description,omitempty"`
	Id          int    `json:"_id,omitempty"`
	Message     string `json:"message,omitempty"`
	Name        string `json:"name,omitempty"`
	Script      string `json:"script,omitempty"`
	Type        string `json:"type,omitempty"`
}

// Get all custom rules.
func ListCustomRules(c api.Client) ([]CustomRule, error) {
	var ans []CustomRule
	if err := c.Request(http.MethodGet, CustomRulesEndpoint, nil, nil, &ans); err != nil {
		return nil, fmt.Errorf("error listing custom rules: %s", err)
	}
	return ans, nil
}

// Get a specific custom rule by ID.
func GetCustomRule(c api.Client, id int) (*CustomRule, error) {
	rules, err := ListCustomRules(c)
	if err != nil {
		return nil, fmt.Errorf("error getting custom rules '%d': %s", id, err)
	}
	for _, val := range rules {
		if val.Id == id {
			return &val, nil
		}
	}
	return nil, fmt.Errorf("custom rule '%d' not found", id)
}

// Get a specific custom rule by Name.
func GetCustomRuleByName(c api.Client, name string) (*CustomRule, error) {
	rules, err := ListCustomRules(c)
	if err != nil {
		return nil, fmt.Errorf("error getting custom rules '%s': %s", name, err)
	}
	for _, val := range rules {
		if val.Name == name {
			return &val, nil
		}
	}
	return nil, fmt.Errorf("custom rule '%s' not found", name)
}

// Create a new custom rule.
func CreateCustomRule(c api.Client, rule CustomRule) (int, error) {
	id, err := GenerateCustomRuleId(c)
	if err != nil {
		return -1, fmt.Errorf("error getting custom rules '%d': %s", id, err)
	}
	rule.Id = id
	return id, UpdateCustomRule(c, rule)
}

// Helper method to generate an ID for new custom rule.
// Finds the maximum custom rule ID and increments it by 1.
func GenerateCustomRuleId(c api.Client) (int, error) {
	rules, err := ListCustomRules(c)
	if err != nil {
		return -1, fmt.Errorf("error getting custom rules: %s", err)
	}

	// Assuming rules may not be sorted by ID.
	maxId := 0
	for _, val := range rules {
		if val.Id > maxId {
			maxId = val.Id
		}
	}
	return maxId + 1, nil
}

// Update an existing collection.
func UpdateCustomRule(c api.Client, rule CustomRule) error {
	return c.Request(http.MethodPut, fmt.Sprintf("%s/%d", CustomRulesEndpoint, rule.Id), nil, rule, nil)
}

// Delete an existing collection.
func DeleteCustomRule(c api.Client, id int) error {
	return c.Request(http.MethodDelete, fmt.Sprintf("%s/%d", CustomRulesEndpoint, id), nil, nil, nil)
}
