package policy

import (
	"fmt"
	"net/http"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/cnnf"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/waas"
)

const (
	FirewallAppAppEmbeddedEndpoint = "api/v1/policies/firewall/app/app-embedded"
	FirewallAppContainerEndpoint   = "api/v1/policies/firewall/app/container"
	FirewallAppHostEndpoint        = "api/v1/policies/firewall/app/host"
	FirewallAppAPISpecEndpoint     = "api/v1/policies/firewall/app/apispec"
	FirewallAppNetworkListEndpoint = "api/v1/policies/firewall/app/network-list"
	FirewallAppOutOfBandEndpoint   = "api/v1/policies/firewall/app/out-of-band"
	FirewallAppServerlessEndpoint  = "api/v1/policies/firewall/app/serverless"

	FirewallNetworkEndpoint = "api/v1/policies/firewall/network"
)

type FirewallAppPolicy struct {
	Id      int             `json:"_id,omitempty"`
	MaxPort int             `json:"maxPort,omitempty"`
	MinPort int             `json:"minPort,omitempty"`
	Rules   []waas.WaasRule `json:"rules,omitempty"`
}

type FirewallAppNetworkList struct {
	Id           int      `json:"_id,omitempty"`
	Description  string   `json:"description,omitempty"`
	Disabled     bool     `json:"disabled,omitempty"`
	Modified     string   `json:"modified,omitempty"`
	Name         string   `json:"name,omitempty"`
	Notes        string   `json:"notes,omitempty"`
	Owner        string   `json:"owner,omitempty"`
	PreviousName string   `json:"previousName,omitempty"`
	Subnets      []string `json:"subnets,omitempty"`
}

type FirewallNetworkPolicy struct {
	Id               int             `json:"_id,omitempty"`
	ContainerEnabled bool            `json:"containerEnabled,omitempty"`
	ContainerRules   []cnnf.Rule     `json:"containerRules,omitempty"`
	HostEnabled      bool            `json:"hostEnabled,omitempty"`
	HostRules        []cnnf.Rule     `json:"hostRules,omitempty"`
	Modified         string          `json:"modified,omitempty"`
	NetworkEntities  []cnnf.Entities `json:"networkEntities,omitempty"`
	Owner            string          `json:"owner,omitempty"`
}

// Resolve endpoints defined in an OpenAPI/Swagger specification
func ResolveFirewallAppAPISpec(c api.Client, content string) (waas.WaasAPISpec, error) {
	var ans waas.WaasAPISpec
	if err := c.Request(http.MethodPost, FirewallAppAPISpecEndpoint, nil, content, &ans); err != nil {
		return ans, fmt.Errorf("error resolving app firewall api spec: %s", err)
	}
	return ans, nil
}

// Create the current firewall app network list.
func CreateFirewallAppNetworkList(c api.Client, networkList FirewallAppNetworkList) error {
	return c.Request(http.MethodPost, FirewallAppNetworkListEndpoint, nil, networkList, nil)
}

// Get the current app-embedded app firewall policy.
func GetFirewallAppAppEmbedded(c api.Client) (FirewallAppPolicy, error) {
	var ans FirewallAppPolicy
	if err := c.Request(http.MethodGet, FirewallAppAppEmbeddedEndpoint, nil, nil, &ans); err != nil {
		return ans, fmt.Errorf("error getting app-embedded app firewall policy: %s", err)
	}
	return ans, nil
}

// Get the current container app firewall policy.
func GetFirewallAppContainer(c api.Client) (FirewallAppPolicy, error) {
	var ans FirewallAppPolicy
	if err := c.Request(http.MethodGet, FirewallAppContainerEndpoint, nil, nil, &ans); err != nil {
		return ans, fmt.Errorf("error getting container app firewall policy: %s", err)
	}
	return ans, nil
}

// Get the current host app firewall policy.
func GetFirewallAppHost(c api.Client) (FirewallAppPolicy, error) {
	var ans FirewallAppPolicy
	if err := c.Request(http.MethodGet, FirewallAppHostEndpoint, nil, nil, &ans); err != nil {
		return ans, fmt.Errorf("error getting host app firewall policy: %s", err)
	}
	return ans, nil
}

// Get the current out of band app firewall policy.
func GetFirewallAppOutOfBand(c api.Client) (FirewallAppPolicy, error) {
	var ans FirewallAppPolicy
	if err := c.Request(http.MethodGet, FirewallAppOutOfBandEndpoint, nil, nil, &ans); err != nil {
		return ans, fmt.Errorf("error getting out of band app firewall policy: %s", err)
	}
	return ans, nil
}

// Get the current serverless app firewall policy.
func GetFirewallAppServerless(c api.Client) (FirewallAppPolicy, error) {
	var ans FirewallAppPolicy
	if err := c.Request(http.MethodGet, FirewallAppServerlessEndpoint, nil, nil, &ans); err != nil {
		return ans, fmt.Errorf("error getting serverless app firewall policy: %s", err)
	}
	return ans, nil
}

// Get the current firewall app network list.
func GetFirewallAppNetworkList(c api.Client) (FirewallAppNetworkList, error) {
	var ans FirewallAppNetworkList
	if err := c.Request(http.MethodGet, FirewallAppNetworkListEndpoint, nil, nil, &ans); err != nil {
		return ans, fmt.Errorf("error getting firewall app network list: %s", err)
	}
	return ans, nil
}

// Get the current firewall network policy.
func GetFirewallNetworkPolicy(c api.Client) (FirewallNetworkPolicy, error) {
	var ans FirewallNetworkPolicy
	if err := c.Request(http.MethodGet, FirewallNetworkEndpoint, nil, nil, &ans); err != nil {
		return ans, fmt.Errorf("error getting firewall network policy: %s", err)
	}
	return ans, nil
}

// Update the current app-embedded app firewall policy.
func UpdateFirewallAppAppEmbedded(c api.Client, policy FirewallAppPolicy) error {
	return c.Request(http.MethodPut, FirewallAppAppEmbeddedEndpoint, nil, policy, nil)
}

// Update the current container app firewall policy.
func UpdateFirewallAppContainer(c api.Client, policy FirewallAppPolicy) error {
	return c.Request(http.MethodPut, FirewallAppContainerEndpoint, nil, policy, nil)
}

// Update the current host app firewall policy.
func UpdateFirewallAppHost(c api.Client, policy FirewallAppPolicy) error {
	return c.Request(http.MethodPut, FirewallAppHostEndpoint, nil, policy, nil)
}

// Update the current out of band app firewall policy.
func UpdateFirewallAppOutOfBand(c api.Client, policy FirewallAppPolicy) error {
	return c.Request(http.MethodPut, FirewallAppOutOfBandEndpoint, nil, policy, nil)
}

// Update the current serverless app firewall policy.
func UpdateFirewallAppServerless(c api.Client, policy FirewallAppPolicy) error {
	return c.Request(http.MethodPut, FirewallAppServerlessEndpoint, nil, policy, nil)
}

// Update the current firewall app network list.
func UpdateFirewallAppNetworkList(c api.Client, networkList FirewallAppNetworkList) error {
	return c.Request(http.MethodPut, FirewallAppNetworkListEndpoint, nil, networkList, nil)
}

// Update the current firewall network policy.
func UpdateFirewallNetworkPolicy(c api.Client, policy FirewallNetworkPolicy) error {
	return c.Request(http.MethodPut, FirewallNetworkEndpoint, nil, policy, nil)
}

// Delete a firewall app network list.
func DeleteFirewallAppNetworkList(c api.Client, name string) error {
	return c.Request(http.MethodDelete, fmt.Sprintf("%s/%s", FirewallAppNetworkListEndpoint, name), nil, nil, nil)
}
