package auth

import (
	"fmt"
	"net/http"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
)

const UsersEndpoint = "api/v1/users"

type User struct {
	AuthType    string           `json:"authType,omitempty"`
	Password    string           `json:"password,omitempty"`
	Permissions []UserPermission `json:"permissions,omitempty"`
	Role        string           `json:"role,omitempty"`
	Username    string           `json:"username,omitempty"`
}

type UserPermission struct {
	Collections []string `json:"collections,omitempty"`
	Project     string   `json:"project,omitempty"`
}

// Get all users.
func ListUsers(c api.Client) ([]User, error) {
	var ans []User
	if err := c.Request(http.MethodGet, UsersEndpoint, nil, nil, &ans); err != nil {
		return nil, fmt.Errorf("error listing users: %s", err)
	}
	return ans, nil
}

// Get a specific user.
func GetUser(c api.Client, name string) (*User, error) {
	users, err := ListUsers(c)
	if err != nil {
		return nil, fmt.Errorf("error getting user '%s': %s", name, err)
	}
	for _, val := range users {
		if val.Username == name {
			return &val, nil
		}
	}
	return nil, fmt.Errorf("user '%s' not found", name)
}

// Create a new user.
func CreateUser(c api.Client, user User) error {
	return c.Request(http.MethodPost, UsersEndpoint, nil, user, nil)
}

// Update an existing user.
func UpdateUser(c api.Client, user User) error {
	return c.Request(http.MethodPut, UsersEndpoint, nil, user, nil)
}

// Delete an existing user.
func DeleteUser(c api.Client, name string) error {
	return c.Request(http.MethodDelete, fmt.Sprintf("%s/%s", UsersEndpoint, name), nil, nil, nil)
}
