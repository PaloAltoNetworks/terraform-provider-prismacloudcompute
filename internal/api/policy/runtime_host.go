package policy

import (
	"fmt"
	"net/http"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/collection"
)

const RuntimeHostEndpoint = "api/v1/policies/runtime/host"

type RuntimeHostPolicy struct {
	Rules []RuntimeHostRule `json:"rules,omitempty"`
}

type RuntimeHostRule struct {
	AntiMalware        RuntimeHostAntiMalware         `json:"antiMalware,omitempty"`
	Collections        []collection.Collection        `json:"collections,omitempty"`
	CustomRules        []RuntimeHostCustomRule        `json:"customRules,omitempty"`
	Disabled           bool                           `json:"disabled"`
	Dns                RuntimeHostDns                 `json:"dns,omitempty"`
	FileIntegrityRules []RuntimeHostFileIntegrityRule `json:"fileIntegrityRules,omitempty"`
	Forensic           RuntimeHostForensic            `json:"forensic,omitempty"`
	LogInspectionRules []RuntimeHostLogInspectionRule `json:"logInspectionRules,omitempty"`
	Name               string                         `json:"name,omitempty"`
	Network            RuntimeHostNetwork             `json:"network,omitempty"`
	Notes              string                         `json:"notes,omitempty"`
}

type RuntimeHostAntiMalware struct {
	AllowedProcesses              []string                   `json:"allowedProcesses,omitempty"`
	CryptoMiner                   string                     `json:"cryptoMiner,omitempty"`
	CustomFeed                    string                     `json:"customFeed,omitempty"`
	DeniedProcesses               RuntimeHostDeniedProcesses `json:"deniedProcesses,omitempty"`
	DetectCompilerGeneratedBinary bool                       `json:"detectCompilerGeneratedBinary"`
	EncryptedBinaries             string                     `json:"encryptedBinaries,omitempty"`
	ExecutionFlowHijack           string                     `json:"executionFlowHijack,omitempty"`
	IntelligenceFeed              string                     `json:"intelligenceFeed,omitempty"`
	ReverseShell                  string                     `json:"reverseShell,omitempty"`
	ServiceUnknownOriginBinary    string                     `json:"serviceUnknownOriginBinary,omitempty"`
	SkipSshTracking               bool                       `json:"skipSSHTracking,omitempty"`
	SuspiciousElfHeaders          string                     `json:"suspiciousELFHeaders,omitempty"`
	TempFsProcesses               string                     `json:"tempFSProc,omitempty"`
	UserUnknownOriginBinary       string                     `json:"userUnknownOriginBinary,omitempty"`
	WebShell                      string                     `json:"webShell,omitempty"`
	WildFireAnalysis              string                     `json:"wildFireAnalysis,omitempty"`
}

type RuntimeHostCustomRule struct {
	Action string `json:"action,omitempty"`
	Effect string `json:"effect,omitempty"`
	Id     int    `json:"_id,omitempty"`
}

type RuntimeHostDeniedProcesses struct {
	Effect string   `json:"effect,omitempty"`
	Paths  []string `json:"paths,omitempty"`
}

type RuntimeHostDns struct {
	Allowed          []string `json:"allow,omitempty"`
	Denied           []string `json:"deny,omitempty"`
	DenyEffect       string   `json:"denyListEffect,omitempty"`
	IntelligenceFeed string   `json:"intelligenceFeed,omitempty"`
}

type RuntimeHostFileIntegrityRule struct {
	AllowedProcesses []string `json:"procWhitelist,omitempty"`
	ExcludedFiles    []string `json:"exclusions,omitempty"`
	Metadata         bool     `json:"metadata"`
	Path             string   `json:"path,omitempty"`
	Read             bool     `json:"read"`
	Recursive        bool     `json:"recursive"`
	Write            bool     `json:"write"`
}

type RuntimeHostForensic struct {
	ActivitiesDisabled       bool `json:"activitiesDisabled"`
	DockerEnabled            bool `json:"dockerEnabled"`
	ReadonlyDockerEnabled    bool `json:"readonlyDockerEnabled"`
	ServiceActivitiesEnabled bool `json:"serviceActivitiesEnabled"`
	SshdEnabled              bool `json:"sshdEnabled"`
	SudoEnabled              bool `json:"sudoEnabled"`
}

type RuntimeHostLogInspectionRule struct {
	Path  string   `json:"path,omitempty"`
	Regex []string `json:"regex,omitempty"`
}

type RuntimeHostNetwork struct {
	AllowedOutboundIps   []string          `json:"allowedOutboundIPs,omitempty"`
	CustomFeed           string            `json:"customFeed,omitempty"`
	DeniedListeningPorts []RuntimeHostPort `json:"deniedListeningPorts,omitempty"`
	DeniedOutboundIps    []string          `json:"deniedOutboundIPs,omitempty"`
	DeniedOutboundPorts  []RuntimeHostPort `json:"deniedOutboundPorts,omitempty"`
	DenyEffect           string            `json:"denyListEffect,omitempty"`
	IntelligenceFeed     string            `json:"intelligenceFeed,omitempty"`
}

type RuntimeHostPort struct {
	Deny  bool `json:"deny"`
	End   int  `json:"end,omitempty"`
	Start int  `json:"start,omitempty"`
}

// Get the current host runtime policy.
func GetRuntimeHost(c api.Client) (RuntimeHostPolicy, error) {
	var ans RuntimeHostPolicy
	if err := c.Request(http.MethodGet, RuntimeHostEndpoint, nil, nil, &ans); err != nil {
		return ans, fmt.Errorf("error getting host runtime policy: %s", err)
	}
	return ans, nil
}

// Update the current host runtime policy.
func UpdateRuntimeHost(c api.Client, policy RuntimeHostPolicy) error {
	return c.Request(http.MethodPut, RuntimeHostEndpoint, nil, policy, nil)
}
