package auth

import (
	"fmt"
	"net/http"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
)

const GroupsEndpoint = "api/v1/groups"

type Group struct {
	Id          string            `json:"groupId,omitempty"` // Group identifier in the Azure SAML identification process.
	LdapGroup   bool              `json:"ldapGroup,omitempty"`
	Name        string            `json:"groupName,omitempty"`
	OauthGroup  bool              `json:"oauthGroup,omitempty"`
	OidcGroup   bool              `json:"oidcGroup,omitempty"`
	Permissions []GroupPermission `json:"permissions,omitempty"`
	Role        string            `json:"role,omitempty"`
	SamlGroup   bool              `json:"samlGroup,omitempty"`
	Users       []GroupUser       `json:"user,omitempty"`
}

type GroupPermission struct {
	Collections []string `json:"collections,omitempty"`
	Project     string   `json:"project,omitempty"`
}

type GroupUser struct {
	Username string `json:"username,omitempty"`
}

// Get all groups.
func ListGroups(c api.Client) ([]Group, error) {
	var ans []Group
	if err := c.Request(http.MethodGet, GroupsEndpoint, nil, nil, &ans); err != nil {
		return nil, fmt.Errorf("error listing groups: %s", err)
	}
	return ans, nil
}

// Get a specific group.
func GetGroup(c api.Client, name string) (*Group, error) {
	groups, err := ListGroups(c)
	if err != nil {
		return nil, fmt.Errorf("error getting group '%s': %s", name, err)
	}
	for _, val := range groups {
		if val.Name == name {
			return &val, nil
		}
	}
	return nil, fmt.Errorf("group '%s' not found", name)
}

// Create a new group.
func CreateGroup(c api.Client, group Group) error {
	return c.Request(http.MethodPost, GroupsEndpoint, nil, group, nil)
}

// Update an existing group.
func UpdateGroup(c api.Client, group Group) error {
	return c.Request(http.MethodPut, fmt.Sprintf("%s/%s", GroupsEndpoint, group.Name), nil, group, nil)
}

// Delete an existing group.
func DeleteGroup(c api.Client, name string) error {
	return c.Request(http.MethodDelete, fmt.Sprintf("%s/%s", GroupsEndpoint, name), nil, nil, nil)
}
