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

// type Slack struct {
// 	Enabled    bool   `json:"enabled"`
// 	WebhookUrl string `json:"webhookUrl"`
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

//Alert profiles types
type Webhook struct {
	Enabled      bool   `json:"enabled"`
	CredentialId string `json:"credentialId,omitempty"`
	Url          string `json:"url,omitempty"`
	CaCert       string `json:"caCert,omitempty"`
	Json         string `json:"json,omitempty"`
}

// Policy types

//Docker - Access
type Docker struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// Admisson - Admission audits
type Admission struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// AppEmbeddedRuntime - App-Embedded Defender runtime
type AppEmbeddedRuntime struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// NetworkFirewall - Cloud Native Network Firewall (CNNF)
type NetworkFirewall struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// ContainerComplianceScan - Container and image compliance
type ContainerComplianceScan struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// ContainerRuntime - Container runtime
type ContainerRuntime struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// Defender - Defender health
type Defender struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// HostComplianceScan - Host compliance
type HostComplianceScan struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// HostRuntime  - Host runtime
type HostRuntime struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// HostVulnerability - Host vulnerabilities
type HostVulnerability struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// ContainerVulnerability - Image vulnerabilities (registry and deployed)
type ContainerVulnerability struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// Incident - Incidents
type Incident struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// KubernetesAudit - Kubernetes audits
type KubernetesAudit struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// ServerlessRuntime - Serverless runtime
type ServerlessRuntime struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// AppEmbeddedAppFirewall - WAAS Firewall (App-Embedded Defender)
type AppEmbeddedAppFirewall struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// ContainerAppFirewall - WAAS Firewall (container)
type ContainerAppFirewall struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// HostAppFirewall - WAAS Firewall (host)
type HostAppFirewall struct {
	Enabled  bool     `json:"enabled"`
	Allrules bool     `json:"allRules"`
	Rules    []string `json:"rules,omitempty"`
}

// ServerlessAppFirewall - WAAS Firewall (serverless)
type ServerlessAppFirewall struct {
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
	Docker                  Docker                  `json:"docker,omitempty"`
	Admission               Admission               `json:"admission,omitempty"`
	AppEmbeddedRuntime      AppEmbeddedRuntime      `json:"appEmbeddedRuntime,omitempty"`
	NetworkFirewall         NetworkFirewall         `json:"networkFirewall,omitempty"`
	ContainerComplianceScan ContainerComplianceScan `json:"containerComplianceScan,omitempty"`
	ContainerRuntime        ContainerRuntime        `json:"containerRuntime,omitempty"`
	Defender                Defender                `json:"defender,omitempty"`
	HostComplianceScan      HostComplianceScan      `json:"hostComplianceScan,omitempty"`
	HostRuntime             HostRuntime             `json:"hostRuntime,omitempty"`
	HostVulnerability       HostVulnerability       `json:"hostVulnerability,omitempty"`
	ContainerVulnerability  ContainerVulnerability  `json:"containerVulnerability,omitempty"`
	Incident                Incident                `json:"incident,omitempty"`
	KubernetesAudit         KubernetesAudit         `json:"kubernetesAudit,omitempty"`
	ServerlessRuntime       ServerlessRuntime       `json:"serverlessRuntime,omitempty"`
	AppEmbeddedAppFirewall  AppEmbeddedAppFirewall  `json:"appEmbeddedAppFirewall,omitempty"`
	ContainerAppFirewall    ContainerAppFirewall    `json:"containerAppFirewall,omitempty"`
	HostAppFirewall         HostAppFirewall         `json:"hostAppFirewall,omitempty"`
	ServerlessAppFirewall   ServerlessAppFirewall   `json:"serverlessAppFirewall,omitempty"`
	WaasHealth              WaasHealth              `json:"waasHealth,omitempty"`
}

// Alertprofile struct
type Alertprofile struct {
	Id                                  string  `json:"_id"`
	Name                                string  `json:"name"`
	External                            bool    `json:"external,omitempty"`
	IntegrationID                       string  `json:"integrationID,omitempty"`
	Webhook                             Webhook `json:"webhook,omitempty"`
	Policy                              Policy  `json:"policy,omitempty"`
	VulnerabilityImmediateAlertsEnabled bool    `json:"vulnerabilityImmediateAlertsEnabled,omitempty"`
	Owner                               string  `json:"owner,omitempty"`
}

// Get all Alertprofiles.
func ListAlertprofiles(c api.Client) ([]Alertprofile, error) {
	var ans []Alertprofile
	if err := c.Request(http.MethodGet, AlertprofilesEndpoint, nil, nil, &ans); err != nil {
		return nil, fmt.Errorf("error listing Alert Profiles: %s", err)
	}
	return ans, nil
}

// Get a specific Alertprofile.
func GetAlertprofile(c api.Client, name string) (*Alertprofile, error) {
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
func CreateAlertprofile(c api.Client, Alertprofile Alertprofile) error {
	return c.Request(http.MethodPost, AlertprofilesEndpoint, nil, Alertprofile, nil)
}

// Update an existing Alertprofile.
func UpdateAlertprofile(c api.Client, Alertprofile Alertprofile) error {
	return c.Request(http.MethodPost, AlertprofilesEndpoint, nil, Alertprofile, nil)
}

// Delete an existing Alertprofile.
func DeleteAlertprofile(c api.Client, name string) error {
	return c.Request(http.MethodDelete, fmt.Sprintf("%s/%s", AlertprofilesEndpoint, name), nil, nil, nil)
}
