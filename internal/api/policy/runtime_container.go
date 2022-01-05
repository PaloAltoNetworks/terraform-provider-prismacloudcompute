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
	AdvancedProtection       bool                         `json:"advancedProtection"`
	CloudMetadataEnforcement bool                         `json:"cloudMetadataEnforcement"`
	Collections              []collection.Collection      `json:"collections,omitempty"`
	CustomRules              []RuntimeContainerCustomRule `json:"customRules,omitempty"`
	Disabled                 bool                         `json:"disabled"`
	Dns                      RuntimeContainerDns          `json:"dns,omitempty"`
	Filesystem               RuntimeContainerFilesystem   `json:"filesystem,omitempty"`
	KubernetesEnforcement    bool                         `json:"kubernetesEnforcement"`
	Name                     string                       `json:"name,omitempty"`
	Network                  RuntimeContainerNetwork      `json:"network,omitempty"`
	Notes                    string                       `json:"notes,omitempty"`
	Processes                RuntimeContainerProcesses    `json:"processes,omitempty"`
	WildFireAnalysis         string                       `json:"wildFireAnalysis,omitempty"`
}

type RuntimeContainerCustomRule struct {
	Action string `json:"action,omitempty"`
	Effect string `json:"effect,omitempty"`
	Id     int    `json:"_id,omitempty"`
}

type RuntimeContainerDns struct {
	Allowed    []string `json:"whitelist,omitempty"`
	Denied     []string `json:"blacklist,omitempty"`
	DenyEffect string   `json:"effect,omitempty"`
}

type RuntimeContainerFilesystem struct {
	Allowed               []string `json:"whitelist,omitempty"`
	BackdoorFiles         bool     `json:"backdoorFiles"`
	CheckNewFiles         bool     `json:"checkNewFiles"`
	Denied                []string `json:"blacklist,omitempty"`
	DenyEffect            string   `json:"effect,omitempty"`
	SkipEncryptedBinaries bool     `json:"skipEncryptedBinaries"`
	SuspiciousElfHeaders  bool     `json:"suspiciousELFHeaders"`
}

type RuntimeContainerNetwork struct {
	AllowedListeningPorts []RuntimeContainerPort `json:"whitelistListeningPorts,omitempty"`
	AllowedOutboundIps    []string               `json:"whitelistIPs,omitempty"`
	AllowedOutboundPorts  []RuntimeContainerPort `json:"whitelistOutboundPorts,omitempty"`
	DeniedListeningPorts  []RuntimeContainerPort `json:"blacklistListeningPorts,omitempty"`
	DeniedOutboundIps     []string               `json:"blacklistIPs,omitempty"`
	DeniedOutboundPorts   []RuntimeContainerPort `json:"blacklistOutboundPorts,omitempty"`
	DenyEffect            string                 `json:"effect,omitempty"`
	DetectPortScan        bool                   `json:"detectPortScan"`
	SkipModifiedProcesses bool                   `json:"skipModifiedProc"`
	SkipRawSockets        bool                   `json:"skipRawSockets"`
}

type RuntimeContainerPort struct {
	Deny  bool `json:"deny"`
	End   int  `json:"end,omitempty"`
	Start int  `json:"start,omitempty"`
}

type RuntimeContainerProcesses struct {
	Allowed              []string `json:"whitelist,omitempty"`
	CheckCryptoMiners    bool     `json:"checkCryptoMiners"`
	CheckLateralMovement bool     `json:"checkLateralMovement"`
	CheckParentChild     bool     `json:"checkParentChild"`
	CheckSuidBinaries    bool     `json:"checkSuidBinaries"`
	Denied               []string `json:"blacklist,omitempty"`
	DenyEffect           string   `json:"effect,omitempty"`
	SkipModified         bool     `json:"skipModified"`
	SkipReverseShell     bool     `json:"skipReverseShell"`
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
