package settings

import (
	"fmt"
	"net/http"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
)

const SettingsRegistryEndpoint = "api/v1/settings/registry"

type RegistrySettings struct {
	Specifications []RegistrySpecification `json:"specifications,omitempty"`
}

type RegistrySpecification struct {
	Cap                      int      `json:"cap,omitempty"`
	Collections              []string `json:"collections,omitempty"`
	Credential               string   `json:"credentialID,omitempty"`
	ExcludedRepositories     []string `json:"excludedRepositories,omitempty"`
	ExcludedTags             []string `json:"excludedTags,omitempty"`
	HarborDeploymentSecurity bool     `json:"harborDeploymentSecurity,omitempty"`
	JfrogRepoTypes           []string `json:"jfrogRepoTypes,omitempty"`
	Namespace                string   `json:"namespace,omitempty"`
	Os                       string   `json:"os,omitempty"`
	Tag                      string   `json:"tag,omitempty"`
	Registry                 string   `json:"registry,omitempty"`
	Repository               string   `json:"repository,omitempty"`
	Scanners                 int      `json:"scanners,omitempty"`
	Version                  string   `json:"version,omitempty"`
	VersionPattern           string   `json:"versionPattern,omitempty"`
}

// Get the current registry scan settings.
func GetRegistrySettings(c api.Client) (RegistrySettings, error) {
	var ans RegistrySettings
	if err := c.Request(http.MethodGet, SettingsRegistryEndpoint, nil, nil, &ans); err != nil {
		return ans, fmt.Errorf("error getting registry settings: %s", err)
	}
	return ans, nil
}

// Update the current registry scan settings.
func UpdateRegistrySettings(c api.Client, registry RegistrySettings) error {
	return c.Request(http.MethodPut, SettingsRegistryEndpoint, nil, registry, nil)
}
