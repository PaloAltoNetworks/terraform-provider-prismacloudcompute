package auth

import (
	"fmt"
	"net/http"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
)

const RolesEndpoint = "api/v1/rbac/roles"

type Role struct {
	Description string           `json:"description,omitempty"`
	Name        string           `json:"name,omitempty"`
	Permissions []RolePermission `json:"perms,omitempty"`
}

type RolePermission struct {
	Name      string `json:"name,omitempty"`
	ReadWrite bool   `json:"readWrite,omitempty"`
}

// Get all roles.
func ListRoles(c api.Client) ([]Role, error) {
	var ans []Role
	if err := c.Request(http.MethodGet, RolesEndpoint, nil, nil, &ans); err != nil {
		return nil, fmt.Errorf("error listing roles: %s", err)
	}
	return ans, nil
}

// Get a specific role.
func GetRole(c api.Client, name string) (*Role, error) {
	roles, err := ListRoles(c)
	if err != nil {
		return nil, fmt.Errorf("error getting role '%s': %s", name, err)
	}
	for _, val := range roles {
		if val.Name == name {
			return &val, nil
		}
	}
	return nil, fmt.Errorf("role '%s' not found", name)
}

// Create a new role.
func CreateRole(c api.Client, role Role) error {
	return c.Request(http.MethodPost, RolesEndpoint, nil, role, nil)
}

// Update an existing role.
func UpdateRole(c api.Client, role Role) error {
	return c.Request(http.MethodPut, RolesEndpoint, nil, role, nil)
}

// Delete an existing role.
func DeleteRole(c api.Client, name string) error {
	return c.Request(http.MethodDelete, fmt.Sprintf("%s/%s", RolesEndpoint, name), nil, nil, nil)
}
