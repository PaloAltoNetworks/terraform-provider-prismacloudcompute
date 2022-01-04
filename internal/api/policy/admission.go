package policy

import (
	"fmt"
	"net/http"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
)

const AdmissionEndpoint = "api/v1/policies/admission"

type AdmissionPolicy struct {
	Id    string          `json:"_id,omitempty"`
	Rules []AdmissionRule `json:"rules,omitempty"`
}

type AdmissionRule struct {
	Description string `json:"description,omitempty"`
	Disabled    bool   `json:"disabled"`
	Effect      string `json:"effect,omitempty"`
	Name        string `json:"name,omitempty"`
	Script      string `json:"script,omitempty"`
}

// Get the current host admission policy.
func GetAdmission(c api.Client) (AdmissionPolicy, error) {
	var ans AdmissionPolicy
	if err := c.Request(http.MethodGet, AdmissionEndpoint, nil, nil, &ans); err != nil {
		return ans, fmt.Errorf("error getting admission policy: %s", err)
	}
	return ans, nil
}

// Update the current host admission policy.
func UpdateAdmission(c api.Client, policy AdmissionPolicy) error {
	return c.Request(http.MethodPut, AdmissionEndpoint, nil, policy, nil)
}
