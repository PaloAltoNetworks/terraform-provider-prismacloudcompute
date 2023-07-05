package account

import (
	"fmt"
	"net/http"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/auth"
)

const CloudScanRulesEndpoint = "api/v1/cloud-scan-rules"

// Serverless scan specs struct
type ServerLessScanSpec struct {
	Enabled         bool `json:"enabled,omitempty"`
	Cap             int  `json:"cap,omitempty"`
	ScanAllVersions bool `json:"scanAllVersions,omitempty"`
	ScanLayers      bool `json:"scanLayers,omitempty"`
}

type AgentlessScanSpec struct {
	Enabled              bool     `json:"enabled,omitempty"`
	HubAccount           bool     `json:"hubAccount,omitempty"`
	ConsoleAddr          string   `json:"consoleAddr,omitempty"`
	ScanNonRunning       bool     `json:"scanNonRunning,omitempty"`
	ProxyAddress         string   `json:"proxyAddress,omitempty"`
	ProxyCA              string   `json:"proxyCA,omitempty"`
	SkipPermissionsCheck bool     `json:"skipPermissionsCheck,omitempty"`
	AutoScale            bool     `json:"autoScale,omitempty"`
	Scanners             int      `json:"scanners,omitempty"`
	SecurityGroup        string   `json:"securityGroup,omitempty"`
	SubNet               string   `json:"subnet,omitempty"`
	Regions              []string `json:"regions,omitempty"`
	CustomTags           []Tag    `json:"customTags,omitempty"`
	IncludedTags         []Tag    `json:"includedTags,omitempty"`
}

type Tag struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type CloudScanRule struct {
	CredentialId                string             `json:"credentialId"`
	Credential                  auth.Credential    `json:"credential,omitempty"`
	DiscoveryEnabled            bool               `json:"discoveryEnabled,omitempty"`
	ServerlessRadarEnabled      bool               `json:"serverlessRadarEnabled,omitempty"`
	VmTagsEnabled               bool               `json:"vmTagsEnabled,omitempty"`
	DiscoverAllFunctionVersions bool               `json:"discoverAllFunctionVersions,omitempty"`
	ServerlessRadarCap          int                `json:"serverlessRadarCap,omitempty"`
	AgentlessScanSpec           AgentlessScanSpec  `json:"agentlessScanSpec,omitempty"`
	ServerlessScanSpec          ServerLessScanSpec `json:"serverlessScanSpec,omitempty"`
	AwsRegionType               string             `json:"awsRegionType,omitempty"`
}

// Get all cloud can rules
func ListCloudScanRules(c api.Client) ([]CloudScanRule, error) {
	var ans []CloudScanRule
	if err := c.Request(http.MethodGet, CloudScanRulesEndpoint, nil, nil, &ans); err != nil {
		return nil, fmt.Errorf("error listing Cloud Scan Rules: %s", err)
	}
	return ans, nil
}

// Get a specific cloud scan rule
func GetCloudScanRule(c api.Client, name string) (*CloudScanRule, error) {
	var ans []CloudScanRule

	if err := c.Request(http.MethodGet, CloudScanRulesEndpoint, map[string]string{"search": name}, nil, &ans); err != nil {
		return nil, fmt.Errorf("error searching Cloud Scan Rules: %s", err)
	}
	for _, val := range ans {
		if val.CredentialId == name {
			return &val, nil
		}
	}
	return nil, fmt.Errorf("Cloud Scan Rule '%s' not found", name)
}

// Create/Update cloud scan rules
func UpdateCloudScanRule(c api.Client, rule []CloudScanRule) error {
	return c.Request(http.MethodPut, CloudScanRulesEndpoint, nil, rule, nil)
}

// Delete an existing cloud scan rule
func DeleteCloudScanRule(c api.Client, name string) error {
	return c.Request(http.MethodDelete, fmt.Sprintf("%s/%s", CloudScanRulesEndpoint, name), nil, nil, nil)
}
