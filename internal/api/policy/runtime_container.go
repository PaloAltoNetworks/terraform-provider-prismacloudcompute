package policy

import (
	"fmt"
	"net/http"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/collection"
)

const RuntimeContainerEndpoint = "api/v1/policies/runtime/container"

type RuntimeContainerPolicy struct {
	LearningDisabled bool                   `json:"learningDisabled,omitempty"`
	Rules            []RuntimeContainerRule `json:"rules,omitempty"`
}

type RuntimeContainerRule struct {
	AdvancedProtectionEffect       string                       `json:"advancedProtectionEffect"`
	CloudMetadataEnforcementEffect string                       `json:"cloudMetadataEnforcementEffect"`
	Collections                    []collection.Collection      `json:"collections,omitempty"`
	CustomRules                    []RuntimeContainerCustomRule `json:"customRules,omitempty"`
	Disabled                       bool                         `json:"disabled"`
	Dns                            RuntimeContainerDns          `json:"dns,omitempty"`
	Filesystem                     RuntimeContainerFilesystem   `json:"filesystem,omitempty"`
	KubernetesEnforcementEffect    string                       `json:"kubernetesEnforcementEffect"`
	Name                           string                       `json:"name,omitempty"`
	PreviousName                   string                       `json:"previousName,omitempty"`
	SkipExecSessions               bool                         `json:"skipExecSessions,omitempty"`
	Network                        RuntimeContainerNetwork      `json:"network,omitempty"`
	Notes                          string                       `json:"notes,omitempty"`
	Processes                      RuntimeContainerProcesses    `json:"processes,omitempty"`
	WildFireAnalysis               string                       `json:"wildFireAnalysis,omitempty"`
}

type RuntimeContainerCustomRule struct {
	Action string `json:"action,omitempty"`
	Effect string `json:"effect,omitempty"`
	Id     int    `json:"_id,omitempty"`
}

type RuntimeContainerDns struct {
	DefaultEffect string                        `json:"defaultEffect,omitempty"`
	Disabled      bool                          `json:"disabled,omitempty"`
	DomainList    RuntimeContainerDnsDomainList `json:"domainList,omitempty"`
}

type RuntimeContainerFilesystem struct {
	AllowedList                []string                   `json:"allowedList,omitempty"`
	BackdoorFilesEffect        string                     `json:"backdoorFilesEffect,omitempty"`
	DefaultEffect              string                     `json:"defaultEffect,omitempty"`
	DeniedList                 RuntimeContainerDeniedList `json:"deniedList,omitempty"`
	Disabled                   bool                       `json:"disabled,omitempty"`
	EncryptedBinariesEffect    string                     `json:"encryptedBinariesEffect,omitempty"`
	NewFilesEffect             string                     `json:"newFilesEffect,omitempty"`
	SuspiciousElfHeadersEffect string                     `json:"suspiciousElfHeadersEffect,omitempty"`
}

type RuntimeContainerNetwork struct {
	AllowedIps         []string                     `json:"allowedIPs,omitempty"`
	DefaultEffect      string                       `json:"defaultEffect,omitempty"`
	DeniedIps          []string                     `json:"deniedIPs,omitempty"`
	DeniedIpsEffect    string                       `json:"deniedIPsEffect,omitempty"`
	Disabled           bool                         `json:"disabled,omitempty"`
	ListeningPorts     RuntimeContainerNetworkPorts `json:"listeningPorts,omitempty"`
	ModifiedProcEffect string                       `json:"modifiedProcEffect,omitempty"`
	OutboundPorts      RuntimeContainerNetworkPorts `json:"outboundPorts,omitempty"`
	PortScanEffect     string                       `json:"portScanEffect,omitempty"`
	RawSocketsEffect   string                       `json:"rawSocketsEffect,omitempty"`
}

type RuntimeContainerNetworkPorts struct {
	Allowed []RuntimeContainerPort `json:"allowed,omitempty"`
	Denied  []RuntimeContainerPort `json:"denied,omitempty"`
	Effect  string                 `json:"effect,omitempty"`
}
type RuntimeContainerDnsDomainList struct {
	Allowed []string `json:"allowed,omitempty"`
	Denied  []string `json:"denied,omitempty"`
	Effect  string   `json:"effect,omitempty"`
}

type RuntimeContainerPort struct {
	Deny  bool `json:"deny"`
	End   int  `json:"end,omitempty"`
	Start int  `json:"start,omitempty"`
}

type RuntimeContainerProcesses struct {
	ModifiedProcessEffect string                     `json:"modifiedProcessEffect,omitempty"`
	CryptoMinersEffect    string                     `json:"cryptoMinersEffect,omitempty"`
	LateralMovementEffect string                     `json:"lateralMovementEffect,omitempty"`
	ReverseShellEffect    string                     `json:"reverseShellEffect,omitempty"`
	SuidBinariesEffect    string                     `json:"suidBinariesEffect,omitempty"`
	DefaultEffect         string                     `json:"defaultEffect,omitempty"`
	CheckParentChild      bool                       `json:"checkParentChild"`
	AllowedList           []string                   `json:"allowedList,omitempty"`
	Disabled              bool                       `json:"disabled"`
	DeniedList            RuntimeContainerDeniedList `json:"deniedList"`
}

type RuntimeContainerDeniedList struct {
	Effect string   `json:"effect,omitempty"`
	Paths  []string `json:"paths,omitempty"`
}

// Get the current container runtime policy.
func GetRuntimeContainer(c api.Client) (RuntimeContainerPolicy, error) {
	var ans RuntimeContainerPolicy
	if err := c.Request(http.MethodGet, RuntimeContainerEndpoint, nil, nil, &ans); err != nil {
		return ans, fmt.Errorf("error getting container runtime policy: %s", err)
	}
	return ans, nil
}

// Update the current container runtime policy.
func UpdateRuntimeContainer(c api.Client, policy RuntimeContainerPolicy) error {
	return c.Request(http.MethodPut, RuntimeContainerEndpoint, nil, policy, nil)
}

// Add new container runtime policy rule
func SetRuntimeContainerRule(c api.Client, policy RuntimeContainerPolicy) error {
	var err error
	for _, val := range policy.Rules {
		err = c.Request(http.MethodPost, RuntimeContainerEndpoint, nil, val, nil)

		if err != nil {
			return fmt.Errorf("error creating container runtime policy rule: %s", err)
		}
	}
	return nil
}
