package policy

import (
	"fmt"
	"net/http"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/collection"
)

const (
	ComplianceCiImagesEndpoint     = "api/v1/policies/compliance/ci/images"
	ComplianceCiServerlessEndpoint = "api/v1/policies/compliance/ci/serverless"
	ComplianceContainerEndpoint    = "api/v1/policies/compliance/container"
	ComplianceHostEndpoint         = "api/v1/policies/compliance/host"
	ComplianceServerlessEndpoint   = "api/v1/policies/compliance/serverless"
)

type CompliancePolicy struct {
	Rules []ComplianceRule `json:"rules,omitempty"`
	Type  string           `json:"policyType,omitempty"`
}

type ComplianceRule struct {
	BlockMessage     string                  `json:"blockMsg,omitempty"`
	Collections      []collection.Collection `json:"collections,omitempty"`
	Conditions       ComplianceConditions    `json:"condition,omitempty"`
	Disabled         bool                    `json:"disabled"`
	Effect           string                  `json:"effect,omitempty"`
	Name             string                  `json:"name,omitempty"`
	Notes            string                  `json:"notes,omitempty"`
	ShowPassedChecks bool                    `json:"allCompliance"`
	Verbose          bool                    `json:"verbose"`
}

type ComplianceConditions struct {
	Checks []ComplianceCheck `json:"vulnerabilities,omitempty"`
}

type ComplianceCheck struct {
	Block bool `json:"block"`
	Id    int  `json:"id,omitempty"`
}

// Get the current CI image compliance policy.
func GetComplianceCiImage(c api.Client) (CompliancePolicy, error) {
	var ans CompliancePolicy
	if err := c.Request(http.MethodGet, ComplianceCiImagesEndpoint, nil, nil, &ans); err != nil {
		return ans, fmt.Errorf("error getting CI image compliance policy: %s", err)
	}
	return ans, nil
}

// Get the current CI serverless compliance policy.
func GetComplianceCiServerless(c api.Client) (CompliancePolicy, error) {
	var ans CompliancePolicy
	if err := c.Request(http.MethodGet, ComplianceCiServerlessEndpoint, nil, nil, &ans); err != nil {
		return ans, fmt.Errorf("error getting CI serverless compliance policy: %s", err)
	}
	return ans, nil
}

// Get the current container compliance policy.
func GetComplianceContainer(c api.Client) (CompliancePolicy, error) {
	var ans CompliancePolicy
	if err := c.Request(http.MethodGet, ComplianceContainerEndpoint, nil, nil, &ans); err != nil {
		return ans, fmt.Errorf("error getting container compliance policy: %s", err)
	}
	return ans, nil
}

// Get the current host compliance policy.
func GetComplianceHost(c api.Client) (CompliancePolicy, error) {
	var ans CompliancePolicy
	if err := c.Request(http.MethodGet, ComplianceHostEndpoint, nil, nil, &ans); err != nil {
		return ans, fmt.Errorf("error getting host compliance policy: %s", err)
	}
	return ans, nil
}

// Get the current serverless compliance policy.
func GetComplianceServerless(c api.Client) (CompliancePolicy, error) {
	var ans CompliancePolicy
	if err := c.Request(http.MethodGet, ComplianceServerlessEndpoint, nil, nil, &ans); err != nil {
		return ans, fmt.Errorf("error getting serverless compliance policy: %s", err)
	}
	return ans, nil
}

// Update the current CI image compliance policy.
func UpdateComplianceCiImage(c api.Client, policy CompliancePolicy) error {
	return c.Request(http.MethodPut, ComplianceCiImagesEndpoint, nil, policy, nil)
}

// Update the current CI serverless compliance policy.
func UpdateComplianceCiServerless(c api.Client, policy CompliancePolicy) error {
	return c.Request(http.MethodPut, ComplianceCiServerlessEndpoint, nil, policy, nil)
}

// Update the current container compliance policy.
func UpdateComplianceContainer(c api.Client, policy CompliancePolicy) error {
	return c.Request(http.MethodPut, ComplianceContainerEndpoint, nil, policy, nil)
}

// Update the current host compliance policy.
func UpdateComplianceHost(c api.Client, policy CompliancePolicy) error {
	return c.Request(http.MethodPut, ComplianceHostEndpoint, nil, policy, nil)
}

// Update the current serverless compliance policy.
func UpdateComplianceServerless(c api.Client, policy CompliancePolicy) error {
	return c.Request(http.MethodPut, ComplianceServerlessEndpoint, nil, policy, nil)
}
