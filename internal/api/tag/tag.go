package tag

import (
	"fmt"
	"net/http"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
)

const TagsEndpoint = "api/v1/tags"

type Vuln struct {
	CheckBaseLayer bool     `json:"checkBaseLayer,omitempty"`
	Comment        string   `json:"comment,omitempty"`
	Id             string   `json:"id,omitempty"`
	PackageName    string   `json:"packageName,omitempty"`
	ResourceType   string   `json:"resourceType,omitempty"`
	Resources      []string `json:"resources,omitempty"`
}

type Tag struct {
	Color       string `json:"color,omitempty"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Vulns       []Vuln `json:"vulns,omitempty"`
}

// Get all tags.
func ListTags(c api.Client) ([]Tag, error) {
	var ans []Tag
	if err := c.Request(http.MethodGet, TagsEndpoint, nil, nil, &ans); err != nil {
		return nil, fmt.Errorf("error listing tags: %s", err)
	}
	return ans, nil
}

// Get a specific tag.
func GetTag(c api.Client, name string) (*Tag, error) {
	tags, err := ListTags(c)
	if err != nil {
		return nil, err
	}
	for _, val := range tags {
		if val.Name == name {
			if val.Vulns == nil {
				val.Vulns = []Vuln{}
			}
			return &val, nil
		}
	}
	return nil, fmt.Errorf("tag '%s' not found", name)
}

// Create a new tag.
func CreateTag(c api.Client, tag Tag) error {
	// there's a bug in Prisma that causes tags to not work correctly if created
	// with vulns, so we create it without vulns and then add them individually
	tagWithoutVulns := tag
	tagWithoutVulns.Vulns = []Vuln{}

	err := c.Request(http.MethodPost, TagsEndpoint, nil, tagWithoutVulns, nil)
	if err != nil {
		return err
	}

	return createVuln(c, tag, tag.Vulns)
}

func createVuln(c api.Client, tag Tag, vuln []Vuln) error {
	for _, val := range vuln {
		err := c.Request(http.MethodPost, fmt.Sprintf("%s/%s/vuln", TagsEndpoint, tag.Name), nil, val, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

// Update an existing tag.
func UpdateTag(c api.Client, tag Tag) error {
	return c.Request(http.MethodPut, fmt.Sprintf("%s/%s", TagsEndpoint, tag.Name), nil, tag, nil)
}

// Delete an existing tag.
func DeleteTag(c api.Client, name string) error {
	return c.Request(http.MethodDelete, fmt.Sprintf("%s/%s", TagsEndpoint, name), nil, nil, nil)
}
