package collection

import (
	"fmt"
	"net/http"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
)

const CollectionsEndpoint = "api/v1/collections"

type Collection struct {
	AccountIds  []string `json:"accountIDs,omitempty"`
	AppIds      []string `json:"appIDs,omitempty"`
	Clusters    []string `json:"clusters,omitempty"`
	CodeRepos   []string `json:"codeRepos,omitempty"`
	Color       string   `json:"color,omitempty"`
	Containers  []string `json:"containers,omitempty"`
	Description string   `json:"description,omitempty"`
	Functions   []string `json:"functions,omitempty"`
	Hosts       []string `json:"hosts,omitempty"`
	Images      []string `json:"images,omitempty"`
	Labels      []string `json:"labels,omitempty"`
	Name        string   `json:"name,omitempty"`
	Namespaces  []string `json:"namespaces,omitempty"`
}

// Get all collections.
func ListCollections(c api.Client) ([]Collection, error) {
	var ans []Collection
	if err := c.Request(http.MethodGet, CollectionsEndpoint, nil, nil, &ans); err != nil {
		return nil, fmt.Errorf("error listing collections: %s", err)
	}
	return ans, nil
}

// Get a specific collection.
func GetCollection(c api.Client, name string) (*Collection, error) {
	collections, err := ListCollections(c)
	if err != nil {
		return nil, fmt.Errorf("error getting collection '%s': %s", name, err)
	}
	for _, val := range collections {
		if val.Name == name {
			return &val, nil
		}
	}
	return nil, fmt.Errorf("collection '%s' not found", name)
}

// Create a new collection.
func CreateCollection(c api.Client, collection Collection) error {
	return c.Request(http.MethodPost, CollectionsEndpoint, nil, collection, nil)
}

// Update an existing collection.
func UpdateCollection(c api.Client, collection Collection) error {
	return c.Request(http.MethodPut, fmt.Sprintf("%s/%s", CollectionsEndpoint, collection.Name), nil, collection, nil)
}

// Delete an existing collection.
func DeleteCollection(c api.Client, name string) error {
	return c.Request(http.MethodDelete, fmt.Sprintf("%s/%s", CollectionsEndpoint, name), nil, nil, nil)
}
