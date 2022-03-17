package policy

import (
	"fmt"
	"net/http"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/collection"
)

const (
	ComplianceCodereposEndpoint   = "api/v1/policies/compliance/coderepos"
	ComplianceCiCodereposEndpoint = "api/v1/policies/compliance/ci/coderepos"
)

type ComplianceCoderepoPolicy struct {
	Rules []ComplianceCoderepoRule `json:"rules,omitempty"`
	Type  string                   `json:"policyType,omitempty"`
}

type ComplianceCoderepoRule struct {
	Collections     []collection.Collection           `json:"collections,omitempty"`
	Disabled        bool                              `json:"disabled"`
	Effect          string                            `json:"effect,omitempty"`
	GraceDays       int                               `json:"graceDays,omitempty"`
	GraceDaysPolicy ComplianceCoderepoGraceDaysPolicy `json:"graceDaysPolicy,omitempty"`
	Name            string                            `json:"name,omitempty"`
	Notes           string                            `json:"notes,omitempty"`
	License         ComplianceCoderepoLicense         `json:"license,omitempty"`
}

type ComplianceCoderepoLicense struct {
	AlertThreshold ComplianceCoderepoThreshold `json:"alertThreshold,omitempty"`
	BlockThreshold ComplianceCoderepoThreshold `json:"blockThreshold,omitempty"`
	Critical       []string                    `json:"critical,omitempty"`
	High           []string                    `json:"high,omitempty"`
	Medium         []string                    `json:"medium,omitempty"`
	Low            []string                    `json:"low,omitempty"`
}

type ComplianceCoderepoThreshold struct {
	Disabled bool `json:"disabled"`
	Enabled  bool `json:"enabled"`
	Value    int  `json:"value,omitempty"`
}

type ComplianceCoderepoGraceDaysPolicy struct {
	Enabled  bool `json:"enabled,omitempty"`
	Low      int  `json:"low,omitempty"`
	Medium   int  `json:"medium,omitempty"`
	High     int  `json:"high,omitempty"`
	Critical int  `json:"critical,omitempty"`
}

// Get the current CI coderepo compliance policy.
func GetComplianceCiCoderepo(c api.Client) (ComplianceCoderepoPolicy, error) {
	var ans ComplianceCoderepoPolicy
	if err := c.Request(http.MethodGet, ComplianceCiCodereposEndpoint, nil, nil, &ans); err != nil {
		return ans, fmt.Errorf("error getting CI coderepo compliance policy: %s", err)
	}
	return ans, nil
}

// Get the current coderepo compliance policy.
func GetComplianceCoderepo(c api.Client) (ComplianceCoderepoPolicy, error) {
	var ans ComplianceCoderepoPolicy
	if err := c.Request(http.MethodGet, ComplianceCodereposEndpoint, nil, nil, &ans); err != nil {
		return ans, fmt.Errorf("error getting coderepo compliance policy: %s", err)
	}
	return ans, nil
}

// Update the current CI coderepo compliance policy.
func UpdateComplianceCiCoderepo(c api.Client, policy ComplianceCoderepoPolicy) error {
	return c.Request(http.MethodPut, ComplianceCiCodereposEndpoint, nil, policy, nil)
}

// Update the current coderepo compliance policy.
func UpdateComplianceCoderepo(c api.Client, policy ComplianceCoderepoPolicy) error {
	return c.Request(http.MethodPut, ComplianceCodereposEndpoint, nil, policy, nil)
}
