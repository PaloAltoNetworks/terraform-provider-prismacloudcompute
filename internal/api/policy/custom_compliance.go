package policy

import (
	"fmt"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"net/http"
)

const CustomCompliancesEndpoint = "api/v1/custom-compliance"

type CustomCompliance struct {
	Id       int    `json:"_id,omitempty"`
	Name     string `json:"name,omitempty"`
	Title    string `json:"title,omitempty"`
	Severity string `json:"severity,omitempty"`
	Script   string `json:"script,omitempty"`
}

// Get all custom Compliances.
func ListCustomCompliance(c api.Client) ([]CustomCompliance, error) {
	var ans []CustomCompliance
	if err := c.Request(http.MethodGet, CustomCompliancesEndpoint, nil, nil, &ans); err != nil {
		return nil, fmt.Errorf("error listing custom Compliances: %s", err)
	}
	return ans, nil
}

// Get a specific custom Compliance by ID.
func GetCustomComplianceById(c api.Client, id int) (*CustomCompliance, error) {
	compliances, err := ListCustomCompliance(c)
	if err != nil {
		return nil, err
	}
	for _, val := range compliances {
		if val.Id == id {
			return &val, nil
		}
	}
	return nil, fmt.Errorf("custom Compliance '%d' not found", id)
}

// Get a specific custom Compliance by name.
func GetCustomComplianceByName(c api.Client, name string) (*CustomCompliance, error) {
	compliances, err := ListCustomCompliance(c)
	if err != nil {
		return nil, err
	}
	for _, val := range compliances {
		if val.Name == name {
			return &val, nil
		}
	}
	return nil, fmt.Errorf("custom Compliance '%s' not found", name)
}

// Create a new custom compliance.
func CreateCustomCompliance(c api.Client, compliance CustomCompliance) (int, error) {
	id, err := GenerateCustomComplianceId(c)
	if err != nil {
		return -1, err
	}
	compliance.Id = id
	return id, UpdateCustomCompliance(c, compliance)
}

// Helper method to generate an ID for new custom Compliance.
// Finds the maximum custom Compliance ID and increments it by 1.
func GenerateCustomComplianceId(c api.Client) (int, error) {
	compliances, err := ListCustomCompliance(c)
	if err != nil {
		return -1, err
	}

	// Assuming Compliances may not be sorted by ID.
	maxId := 9000
	for _, val := range compliances {
		if val.Id > maxId {
			maxId = val.Id
		}
	}
	return maxId + 1, nil
}

// Update an existing custom Compliance.
func UpdateCustomCompliance(c api.Client, compliance CustomCompliance) error {
	var ans CustomCompliance

	return c.Request(http.MethodPut, CustomCompliancesEndpoint, nil, compliance, &ans)
}

// Delete an existing custom Compliance.
func DeleteCustomCompliance(c api.Client, id int) error {
	return c.Request(http.MethodDelete, fmt.Sprintf("%s/%d", CustomCompliancesEndpoint, id), nil, nil, nil)
}
