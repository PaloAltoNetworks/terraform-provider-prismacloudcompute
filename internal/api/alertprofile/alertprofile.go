package alertprofile

import (
	"fmt"
	"net/http"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
)

const AlertprofilesEndpoint = "api/v1/alert-profiles"

// type Email struct {
// 	Enabled      bool   `json:"enabled"`
// 	SmtpAddress  string `json:"smtpAddress"`
// 	Port         int    `json:"port"`
// 	CredentialId string `json:"credentialId"`
// 	From         string `json:"from"`
// 	Ssl          string `json:"ssl"`
// }

// type Jira struct {
// 	Enabled      bool   `json:"enabled"`
// 	BaseUrl      string `json:"baseUrl"`
// 	CredentialId int    `json:"credentialId"`
// 	CaCert       string `json:"caCert"`
// 	ProjectKey   string `json:"projectKey"`
// 	IssueType    string `json:"issueType"`
// 	Priority     string `json:"priority"`
// 	Labels       string `json:"labels"`
// 	Assignee     string `json:"assignee"`
// }

// type SecurityCenter struct {
// 	Enabled      bool   `json:"enabled"`
// 	CredentialId int    `json:"credentialId"`
// 	SourceID     string `json:"sourceID"`
// }

// type GcpPubsub struct {
// 	Enabled      bool   `json:"enabled"`
// 	CredentialId int    `json:"credentialId"`
// 	Topic        string `json:"topic"`
// }

// type SecurityHub struct {
// 	Enabled      bool   `json:"enabled"`
// 	CredentialId string `json:"credentialId"`
// 	Region       string `json:"region"`
// 	AccountID    string `json:"accountID"`
// }

// Alert profiles types
type Slack struct {
	Enabled    bool   `json:"enabled"`
	WebhookUrl string `json:"webhookUrl"`
}

type Webhook struct {
	Enabled      bool   `json:"enabled"`
	CredentialId string `json:"credentialId,omitempty"`
	Url          string `json:"url,omitempty"`
	CaCert       string `json:"caCert,omitempty"`
	Json         string `json:"json,omitempty"`
}

// Policy types

// Admission audits
type Admission struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// WAAS Firewall (serverless)
type AgentlessAppFirewall struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// WAAS Firewall (App-Embedded Defender)
type AppEmbeddedAppFirewall struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// App-Embedded Defender runtime
type AppEmbeddedRuntime struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

type CloudDiscovery struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

type CodeRepoVulnerability struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// WAAS Firewall (container)
type ContainerAppFirewall struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

type ContainerCompliance struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// Container and image compliance
type ContainerComplianceScan struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// Container runtime
type ContainerRuntime struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// Deployed image vulnerabilities
type ContainerVulnerability struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// Defender health
type Defender struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// Access
type Docker struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// WAAS Firewall (host)
type HostAppFirewall struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

type HostCompliance struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// Host compliance
type HostComplianceScan struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// Host runtime
type HostRuntime struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// Host vulnerabilities
type HostVulnerability struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// Incidents
type Incident struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// Kubernetes audits
type KubernetesAudit struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// Cloud Native Network Segmentation (CNNS)
type NetworkFirewall struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// Registry image vulnerabilities
type RegistryVulnerability struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// WAAS Firewall (serverless)
type ServerlessAppFirewall struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// Serverless runtime
type ServerlessRuntime struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// VM images compliance
type VmCompliance struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// VM images vulnerabilities
type VmVulnerability struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// WaasHealth - WAAS health
type WaasHealth struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// Policy struct
type Policy struct {
	Admission               Admission               `json:"admission,omitempty"`
	AgentlessAppFirewall    AgentlessAppFirewall    `json:"agentlessAppFirewall,omitempty"`
	AppEmbeddedAppFirewall  AppEmbeddedAppFirewall  `json:"appEmbeddedAppFirewall,omitempty"`
	AppEmbeddedRuntime      AppEmbeddedRuntime      `json:"appEmbeddedRuntime,omitempty"`
	CloudDiscovery          CloudDiscovery          `json:"cloudDiscovery,omitempty"`
	CodeRepoVulnerability   CodeRepoVulnerability   `json:"codeRepoVulnerability"`
	ContainerAppFirewall    ContainerAppFirewall    `json:"containerAppFirewall,omitempty"`
	ContainerCompliance     ContainerCompliance     `json:"containerCompliance,omitempty"`
	ContainerComplianceScan ContainerComplianceScan `json:"containerComplianceScan,omitempty"`
	ContainerRuntime        ContainerRuntime        `json:"containerRuntime,omitempty"`
	ContainerVulnerability  ContainerVulnerability  `json:"containerVulnerability,omitempty"`
	Defender                Defender                `json:"defender,omitempty"`
	Docker                  Docker                  `json:"docker,omitempty"`
	HostAppFirewall         HostAppFirewall         `json:"hostAppFirewall,omitempty"`
	HostCompliance          HostCompliance          `json:"hostCompliance,omitempty"`
	HostComplianceScan      HostComplianceScan      `json:"hostComplianceScan,omitempty"`
	HostRuntime             HostRuntime             `json:"hostRuntime,omitempty"`
	HostVulnerability       HostVulnerability       `json:"hostVulnerability,omitempty"`
	Incident                Incident                `json:"incident,omitempty"`
	KubernetesAudit         KubernetesAudit         `json:"kubernetesAudit,omitempty"`
	NetworkFirewall         NetworkFirewall         `json:"networkFirewall,omitempty"`
	RegistryVulnerability   RegistryVulnerability   `json:"registryVulnerability,omitempty"`
	ServerlessAppFirewall   ServerlessAppFirewall   `json:"serverlessAppFirewall,omitempty"`
	ServerlessRuntime       ServerlessRuntime       `json:"serverlessRuntime,omitempty"`
	VmCompliance            VmCompliance            `json:"vmCompliance,omitempty"`
	VmVulnerability         VmVulnerability         `json:"vmVulnerability,omitempty"`
	WaasHealth              WaasHealth              `json:"waasHealth,omitempty"`
}

// AlertProfile struct
type AlertProfile struct {
	Id                                  string  `json:"_id"`
	Name                                string  `json:"name"`
	VulnerabilityImmediateAlertsEnabled bool    `json:"vulnerabilityImmediateAlertsEnabled,omitempty"`
	Owner                               string  `json:"owner,omitempty"`
	Slack                               Slack   `json:"slack,omitempty"`
	Webhook                             Webhook `json:"webhook,omitempty"`
	Policy                              Policy  `json:"policy,omitempty"`
}

// Get all Alertprofiles.
func ListAlertprofiles(c api.Client) ([]AlertProfile, error) {
	var ans []AlertProfile
	if err := c.Request(http.MethodGet, AlertprofilesEndpoint, nil, nil, &ans); err != nil {
		return nil, fmt.Errorf("error listing Alert Profiles: %s", err)
	}
	return ans, nil
}

// Get a specific Alertprofile.
func GetAlertprofile(c api.Client, name string) (*AlertProfile, error) {
	Alertprofiles, err := ListAlertprofiles(c)
	if err != nil {
		return nil, err
	}
	for _, val := range Alertprofiles {
		if val.Name == name {
			return &val, nil
		}
	}
	return nil, fmt.Errorf("error Alert Profile '%s' not found", name)
}

// Create a new Alertprofile.
func CreateAlertprofile(c api.Client, Alertprofile AlertProfile) error {
	return c.Request(http.MethodPost, AlertprofilesEndpoint, nil, Alertprofile, nil)
}

// Update an existing Alertprofile.
func UpdateAlertprofile(c api.Client, Alertprofile AlertProfile) error {
	return c.Request(http.MethodPost, AlertprofilesEndpoint, nil, Alertprofile, nil)
}

// Delete an existing Alertprofile.
func DeleteAlertprofile(c api.Client, name string) error {
	return c.Request(http.MethodDelete, fmt.Sprintf("%s/%s", AlertprofilesEndpoint, name), nil, nil, nil)
}
