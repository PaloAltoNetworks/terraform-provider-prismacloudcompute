package policy

import (
	"fmt"
	"net/http"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/collection"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/common"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/customrules"
)

const (
	FirewallAppAppEmbeddedEndpoint = "api/v1/policies/firewall/app/app-embedded"
	FirewallAppAPISpecEndpoint     = "api/v1/policies/firewall/app/apispec"
)

type WaasEndpoint struct {
	BasePath     string `json:"basePath,omitempty"`
	ExposedPort  int    `json:"exposedPort,omitempty"`
	GRPC         bool   `json:"grpc,omitempty"`
	Host         string `json:"host,omitempty"`
	HTTP2        bool   `json:"http2,omitempty"`
	InternalPort int    `json:"internalPort,omitempty"`
	TLS          bool   `json:"tls,omitempty"`
}

type WaasParam struct {
	AllowEmptyValue bool   `json:"allowEmptyValue,omitempty"`
	Array           bool   `json:"array,omitempty"`
	Explode         bool   `json:"explode,omitempty"`
	Location        string `json:"location,omitempty"`
	Max             int    `json:"max,omitempty"`
	Min             int    `json:"min,omitempty"`
	Name            string `json:"name,omitempty"`
	Required        bool   `json:"required,omitempty"`
	Style           string `json:"style,omitempty"`
	Type            string `json:"type,omitempty"`
}

type WaasMethod struct {
	Method     string      `json:"method,omitempty"`
	Parameters []WaasParam `json:"parameters,omitempty"`
}

type WaasPath struct {
	Methods []WaasMethod `json:"methods,omitempty"`
	Path    string       `json:"path,omitempty"`
}

type WaasExceptionField struct {
	Key      string `json:"key,omitempty"`
	Location string `json:"location,omitempty"`
}

type WaasProtectionConfig struct {
	Effect          string               `json:"effect,omitempty"`
	ExceptionFields []WaasExceptionField `json:"exceptionFields,omitempty"`
}

type WaasBodyConfig struct {
	InspectionLimitExceededEffect string `json:"inspectionLimitExceededEffect,omitempty"`
	InspectionSizeBytes           int    `json:"inspectionSizeBytes,omitempty"`
	Skip                          bool   `json:"skip,omitempty"`
}

type WaasJSInjectionSpec struct {
	Enabled       bool   `json:"enabled,omitempty"`
	TimeoutEffect string `json:"timeoutEffect,omitempty"`
}

type WaasKnownBotProtectionsSpec struct {
	Archiving            string `json:"archiving,omitempty"`
	BusinessAnalytics    string `json:"businessAnalytics,omitempty"`
	CareerSearch         string `json:"careerSearch,omitempty"`
	ContentFeedClients   string `json:"contentFeedClients,omitempty"`
	Educational          string `json:"educational,omitempty"`
	Financial            string `json:"financial,omitempty"`
	MediaSearch          string `json:"mediaSearch,omitempty"`
	News                 string `json:"news,omitempty"`
	SearchEngineCrawlers string `json:"searchEngineCrawlers,omitempty"`
}

type WaasReCAPTCHASpec struct {
	AllSessions            bool          `json:"allSessions,omitempty"`
	Enabled                bool          `json:"enabled,omitempty"`
	SecretKey              common.Secret `json:"secretKey,omitempty"`
	SiteKey                string        `json:"siteKey,omitempty"`
	SuccessExpirationHours int           `json:"successExpirationHours,omitempty"`
	Type                   string        `json:"type,omitempty"`
}

type WaasRequestAnomalies struct {
	Effect    string `json:"effect,omitempty"`
	Threshold int    `json:"threshold,omitempty"`
}

type WaasUnknownBotProtectionSpec struct {
	APILibraries         string               `json:"apiLibraries,omitempty"`
	BotImpersonation     string               `json:"botImpersonation,omitempty"`
	BrowserImpersonation string               `json:"browserImpersonation,omitempty"`
	Generic              string               `json:"generic,omitempty"`
	RequestAnomalies     WaasRequestAnomalies `json:"requestAnomalies,omitempty"`
	WebAutomationTools   string               `json:"webAutomationTools,omitempty"`
	WebScrapers          string               `json:"webScrapers,omitempty"`
}

type WaasUserDefinedBot struct {
	Effect       string   `json:"effect,omitempty"`
	HeaderName   string   `json:"headerName,omitempty"`
	HeaderValues []string `json:"headerValues,omitempty"`
	Name         string   `json:"name,omitempty"`
	Subnets      []string `json:"subnets,omitempty"`
}

type WaasBotProtectionSpec struct {
	InterstitialPage         bool                         `json:"interstitialPage,omitempty"`
	JSInjectionSpec          WaasJSInjectionSpec          `json:"jsInjectionSpec,omitempty"`
	KnownBotProtectionsSpec  WaasKnownBotProtectionsSpec  `json:"knownBotProtectionsSpec,omitempty"`
	ReCAPTCHASpec            WaasReCAPTCHASpec            `json:"reCAPTCHASpec,omitempty"`
	SessionValidation        string                       `json:"sessionValidation,omitempty"`
	UnknownBotProtectionSpec WaasUnknownBotProtectionSpec `json:"unknownBotProtectionSpec,omitempty"`
	UserDefinedBots          []WaasUserDefinedBot         `json:"userDefinedBots,omitempty"`
}

type WaasCustomBlockResponseConfig struct {
	Body    string `json:"body,omitempty"`
	Code    int    `json:"code,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
}

type WaasDoSRates struct {
	Average int `json:"average,omitempty"`
	Burst   int `json:"burst,omitempty"`
}

type WaasStatusCodeRange struct {
	End   int `json:"end,omitempty"`
	Start int `json:"start,omitempty"`
}

type WaasDoSMatchCondition struct {
	FileTypes          []string              `json:"fileTypes,omitempty"`
	Methods            []string              `json:"methods,omitempty"`
	ResponseCodeRanges []WaasStatusCodeRange `json:"responseCodeRanges,omitempty"`
}

type WaasDoSConfig struct {
	Alert                WaasDoSRates            `json:"alert,omitempty"`
	Ban                  WaasDoSRates            `json:"ban,omitempty"`
	Enabled              bool                    `json:"enabled,omitempty"`
	ExcludedNetworkLists []string                `json:"excludedNetworkLists,omitempty"`
	MatchConditions      []WaasDoSMatchCondition `json:"matchConditions,omitempty"`
	TrackSession         bool                    `json:"trackSession,omitempty"`
}

type WaasHeaderSpec struct {
	Allow    bool     `json:"allow,omitempty"`
	Effect   string   `json:"effect,omitempty"`
	Name     string   `json:"name,omitempty"`
	Required bool     `json:"required,omitempty"`
	Values   []string `json:"values,omitempty"`
}

type WaasIntelGatheringConfig struct {
	InfoLeakageEffect         string `json:"infoLeakageEffect,omitempty"`
	RemoveFingerprintsEnabled bool   `json:"removeFingerprintsEnabled,omitempty"`
}

type WaasMaliciousUploadConfig struct {
	AllowedExtensions []string `json:"allowedExtensions,omitempty"`
	AllowedFileTypes  []string `json:"allowedFileTypes,omitempty"`
	Effect            string   `json:"effect,omitempty"`
}

type WaasAccessControls struct {
	Alert          []string `json:"alert,omitempty"`
	Allow          []string `json:"allow,omitempty"`
	AllowMode      bool     `json:"allowMode,omitempty"`
	Enabled        bool     `json:"enabled,omitempty"`
	FallbackEffect string   `json:"fallbackEffect,omitempty"`
	Prevent        []string `json:"prevent,omitempty"`
}

type WaasNetworkControls struct {
	AdvancedProtectionEffect string             `json:"advancedProtectionEffect,omitempty"`
	Countries                WaasAccessControls `json:"countries,omitempty"`
	ExceptionSubnets         []string           `json:"exceptionSubnets,omitempty"`
	Subnets                  WaasAccessControls `json:"subnets,omitempty"`
}

type WaasRemoteHostForwardingConfig struct {
	Enabled bool   `json:"enabled,omitempty"`
	Target  string `json:"target,omitempty"`
}

type WaasResponseHeaderSpec struct {
	Name     string   `json:"name,omitempty"`
	Override bool     `json:"override,omitempty"`
	Values   []string `json:"values,omitempty"`
}

type WaasHSTSConfig struct {
	Enabled           bool `json:"enabled,omitempty"`
	IncludeSubdomains bool `json:"includeSubdomains,omitempty"`
	MaxAgeSeconds     int  `json:"maxAgeSeconds,omitempty"`
	Preload           bool `json:"preload,omitempty"`
}

type WaasCertificateMeta struct {
	IssuerName  string `json:"issuerName,omitempty"`
	NotAfter    string `json:"notAfter,omitempty"`
	SubjectName string `json:"subjectName,omitempty"`
}

type WaasTLSConfig struct {
	HSTSConfig    WaasHSTSConfig      `json:"HSTSConfig,omitempty"`
	Metadata      WaasCertificateMeta `json:"metadata,omitempty"`
	MinTLSVersion string              `json:"minTLSVersion,omitempty"`
}

type WaasApplicationSpec struct {
	APISpec               WaasAPISpec                    `json:"apiSpec,omitempty"`
	AppID                 string                         `json:"appID,omitempty"`
	AttackTools           WaasProtectionConfig           `json:"attackTools,omitempty"`
	BanDurationMinutes    int                            `json:"banDurationMinutes,omitempty"`
	Body                  WaasBodyConfig                 `json:"body,omitempty"`
	BotProtectionSpec     WaasBotProtectionSpec          `json:"botProtectionSpec,omitempty"`
	Certificate           common.Secret                  `json:"certificate,omitempty"`
	ClickjackingEnabled   bool                           `json:"clickjackingEnabled,omitempty"`
	CMDI                  WaasProtectionConfig           `json:"cmdi,omitempty"`
	CodeInjection         WaasProtectionConfig           `json:"codeInjection,omitempty"`
	CSRFEnabled           bool                           `json:"csrfEnabled,omitempty"`
	CustomBlockResponse   WaasCustomBlockResponseConfig  `json:"customBlockResponse,omitempty"`
	CustomRules           []customrules.Ref              `json:"customRules,omitempty"`
	DisableEventIDHeader  bool                           `json:"disableEventIDHeader,omitempty"`
	DoSConfig             WaasDoSConfig                  `json:"dosConfig,omitempty"`
	HeaderSpecs           []WaasHeaderSpec               `json:"headerSpecs,omitempty"`
	IntelGathering        WaasIntelGatheringConfig       `json:"intelGathering,omitempty"`
	LFI                   WaasProtectionConfig           `json:"lfi,omitempty"`
	MalformedReq          WaasProtectionConfig           `json:"malformedReq,omitempty"`
	MaliciousUpload       WaasMaliciousUploadConfig      `json:"maliciousUpload,omitempty"`
	NetworkControls       WaasNetworkControls            `json:"networkControls,omitempty"`
	RemoteHostForwarding  WaasRemoteHostForwardingConfig `json:"remoteHostForwarding,omitempty"`
	ResponseHeaderSpecs   WaasResponseHeaderSpec         `json:"responseHeaderSpecs,omitempty"`
	SessionCookieBan      bool                           `json:"sessionCookieBan,omitempty"`
	SessionCookieEnabled  bool                           `json:"sessionCookieEnabled,omitempty"`
	SessionCookieSameSite string                         `json:"sessionCookieSameSite,omitempty"`
	SessionCookieSecure   bool                           `json:"sessionCookieSecure,omitempty"`
	Shellshock            WaasProtectionConfig           `json:"shellshock,omitempty"`
	SQLI                  WaasProtectionConfig           `json:"sqli,omitempty"`
	TLSConfig             WaasTLSConfig                  `json:"tlsConfig,omitempty"`
	XSS                   WaasProtectionConfig           `json:"xss,omitempty"`
}

type WaasTrafficMirroringConfig struct {
	Enabled bool `json:"enabled,omitempty"`
}

type WaasRule struct {
	AllowMalformedHttpHeaderNames bool                       `json:"allowMalformedHttpHeaderNames,omitempty"`
	ApplicationsSpec              []WaasApplicationSpec      `json:"applicationsSpec,omitempty"`
	AutoProtectPorts              bool                       `json:"autoProtectPorts,omitempty"`
	Collections                   []collection.Collection    `json:"collections,omitempty"`
	Disabled                      bool                       `json:"disabled,omitempty"`
	Modified                      string                     `json:"modified,omitempty"`
	Name                          string                     `json:"name,omitempty"`
	Notes                         string                     `json:"notes,omitempty"`
	Owner                         string                     `json:"owner,omitempty"`
	PreviousName                  string                     `json:"previousName,omitempty"`
	ReadTimeoutSeconds            int                        `json:"readTimeoutSeconds,omitempty"`
	SkipAPILearning               bool                       `json:"skipAPILearning,omitempty"`
	TrafficMirroring              WaasTrafficMirroringConfig `json:"trafficMirroring,omitempty"`
	Windows                       bool                       `json:"windows,omitempty"`
}

type WaasAPISpec struct {
	Description              string         `json:"description,omitempty"`
	Effect                   string         `json:"effect,omitempty"`
	Endpoints                []WaasEndpoint `json:"endpoints,omitempty"`
	FallbackEffect           string         `json:"fallbackEffect,omitempty"`
	Paths                    []WaasPath     `json:"paths,omitempty"`
	QueryParamFallbackEffect string         `json:"queryParamFallbackEffect,omitempty"`
}

type FirewallAppAppEmbeddedPolicy struct {
	Id      int        `json:"_id,omitempty"`
	MaxPort int        `json:"maxPort,omitempty"`
	MinPort int        `json:"minPort,omitempty"`
	Rules   []WaasRule `json:"rules,omitempty"`
}

// Resolve endpoints defined in an OpenAPI/Swagger specification
func ResolveFirewallAppAPISpec(c api.Client, content string) (WaasAPISpec, error) {
	var ans WaasAPISpec
	if err := c.Request(http.MethodPost, FirewallAppAPISpecEndpoint, nil, content, &ans); err != nil {
		return ans, fmt.Errorf("error resolving app firewall api spec: %s", err)
	}
	return ans, nil
}

// Get the current app-embedded app firewall policy.
func GetFirewallAppAppEmbedded(c api.Client) (FirewallAppAppEmbeddedPolicy, error) {
	var ans FirewallAppAppEmbeddedPolicy
	if err := c.Request(http.MethodGet, FirewallAppAppEmbeddedEndpoint, nil, nil, &ans); err != nil {
		return ans, fmt.Errorf("error getting app-embedded app firewall policy: %s", err)
	}
	return ans, nil
}

// Update the current app-embedded app firewall policy.
func UpdateFirewallAppAppEmbedded(c api.Client, policy FirewallAppAppEmbeddedPolicy) error {
	return c.Request(http.MethodPut, FirewallAppAppEmbeddedEndpoint, nil, policy, nil)
}
