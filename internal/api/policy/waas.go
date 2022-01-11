package policy

import (
	"fmt"
	"net/http"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/collection"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/common"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/rule"
)

const WaasContainerEndpoint = "/api/v1/policies/firewall/app/container"

type WaasContainerPolicy struct {
	Id      string              `json:"_id"`
	MaxPort int                 `json:"maxPort"`
	MinPort int                 `json:"minPort"`
	Rules   []WaasContainerRule `json:"rules"`
}

type WaasContainerRule struct {
	AllowMalformedHttpHeaderNames bool                           `json:"allowMalformedHttpHeaderNames,omitempty"`
	ApplicationsSpec              []WaasContainerApplicationSpec `json:"applicationsSpec,omitempty"`
	Collections                   []collection.Collection        `json:"collections"`
	Name                          string                         `json:"name"`
	Notes                         string                         `json:"notes,omitempty"`
	ReadTimeoutSeconds            int                            `json:"readTimeoutSeconds,omitempty"`
}

type WaasContainerApplicationSpec struct {
	ApiSpec               WaasContainerApiSpec             `json:"apiSpec"`
	AppId                 string                           `json:"appID"`
	AttackTools           WaasContainerProtectionConfig    `json:"attackTools"`
	BanDurationMinutes    int                              `json:"banDurationMinutes,omitempty"`
	Body                  WaasContainerBody                `json:"body"`
	BotProtectionSpec     WaasContainerBotProtectionSpec   `json:"botProtectionSpec"`
	Certificate           common.Secret                    `json:"certificate"`
	ClickjackingEnabled   bool                             `json:"clickjackingEnabled"`
	Cmdi                  WaasContainerProtectionConfig    `json:"cmdi"`
	CodeInjection         WaasContainerProtectionConfig    `json:"codeInjection"`
	CsrfEnabled           bool                             `json:"csrfEnabled"`
	CustomBlockResponse   WaasContainerCustomBlockResponse `json:"customBlockResponse"`
	CustomRules           []rule.CustomRule                `json:"customRules,omitempty"`
	DisableEventIdHeader  bool                             `json:"disableEventIDHeader,omitempty"`
	DosConfig             WaasContainerDosConfig           `json:"dosConfig"`
	HeaderSpecs           []WaasContainerHeaderSpec        `json:"headerSpecs,omitempty"`
	IntelGathering        WaasContainerIntelGathering      `json:"intelGathering"`
	Lfi                   WaasContainerProtectionConfig    `json:"lfi"`
	MalformedReq          WaasContainerProtectionConfig    `json:"malformedReq"`
	MaliciousUpload       WaasContainerMaliciousUpload     `json:"maliciousUpload"`
	NetworkControls       WaasContainerNetworkControls     `json:"networkControls"`
	SessionCookieBan      bool                             `json:"sessionCookieBan,omitempty"`
	SessionCookieEnabled  bool                             `json:"sessionCookieEnabled,omitempty"`
	SessionCookieSameSite string                           `json:"sessionCookieSameSite,omitempty"`
	SessionCookieSecure   bool                             `json:"sessionCookieSecure,omitempty"`
	Shellshock            WaasContainerProtectionConfig    `json:"shellshock"`
	Sqli                  WaasContainerProtectionConfig    `json:"sqli"`
	TlsConfig             WaasContainerTlsConfig           `json:"tlsConfig"`
	Xss                   WaasContainerProtectionConfig    `json:"xss"`
}

type WaasContainerApiSpec struct {
	Description              string                         `json:"description,omitempty"`
	Effect                   string                         `json:"effect"` // "API protections - Parameter violations"
	Endpoints                []WaasContainerApiSpecEndpoint `json:"endpoints,omitempty"`
	FallbackEffect           string                         `json:"fallbackEffect"` // "API protections - Unspecified paths/methods"
	Paths                    []WaasContainerApiSpecPath     `json:"paths,omitempty"`
	QueryParamFallbackEffect string                         `json:"queryParamFallbackEffect"` // "Unspecified query params"
	SkipLearning             bool                           `json:"skipLearning"`
}

type WaasContainerApiSpecEndpoint struct {
	BasePath     string `json:"basePath"`
	ExposedPort  int    `json:"exposedPort"`
	Grpc         bool   `json:"grpc"`
	Host         string `json:"host"`
	Http2        bool   `json:"http2"`
	InternalPort int    `json:"internalPort"`
	Tls          bool   `json:"tls"`
}

type WaasContainerApiSpecPath struct {
	Methods []WaasContainerApiSpecPathMethod `json:"methods,omitempty"`
	Path    string                           `json:"path"`
}

type WaasContainerApiSpecPathMethod struct {
	Method     string                                    `json:"method"`
	Parameters []WaasContainerApiSpecPathMethodParameter `json:"parameters,omitempty"`
}

type WaasContainerApiSpecPathMethodParameter struct {
	AllowEmptyValue bool   `json:"allowEmptyValue,omitempty"`
	Array           bool   `json:"array,omitempty"`
	Explode         bool   `json:"explode,omitempty"`
	Location        string `json:"location"`
	Max             int    `json:"max,omitempty"`
	Min             int    `json:"min,omitempty"`
	Name            string `json:"name"`
	Required        bool   `json:"required,omitempty"`
	Style           string `json:"style,omitempty"`
	Type            string `json:"type"`
}

type WaasContainerProtectionConfig struct {
	Effect          string                        `json:"effect"`
	ExceptionFields []WaasContainerExceptionField `json:"exceptionFields"`
}

type WaasContainerExceptionField struct {
	Key      string `json:"key"`
	Location string `json:"location"`
}

type WaasContainerBody struct {
	InspectionSizeBytes int  `json:"inspectionSizeBytes"`
	Skip                bool `json:"skip,omitempty"`
}

type WaasContainerBotProtectionSpec struct {
	InterstitialPage         bool                                                   `json:"interstitialPage"`
	JsInjectionSpec          WaasContainerBotProtectionSpecJsInjectionSpec          `json:"jsInjectionSpec"`
	KnownBotProtectionsSpec  WaasContainerBotProtectionSpecKnownBotProtectionsSpec  `json:"knownBotProtectionsSpec"`
	RecaptchaSpec            WaasContainerBotProtectionSpecRecaptchaSpec            `json:"reCAPTCHASpec"`
	SessionValidation        string                                                 `json:"sessionValidation"`
	UnknownBotProtectionSpec WaasContainerBotProtectionSpecUnknownBotProtectionSpec `json:"unknownBotProtectionSpec"`
	UserDefinedBots          []WaasContainerBotProtectionSpecUserDefinedBot         `json:"userDefinedBots"`
}

type WaasContainerBotProtectionSpecJsInjectionSpec struct {
	Enabled       bool   `json:"enabled"`
	TimeoutEffect string `json:"timeoutEffect"`
}

type WaasContainerBotProtectionSpecKnownBotProtectionsSpec struct {
	Archiving            string `json:"archiving"`
	BusinessAnalytics    string `json:"businessAnalytics"`
	CareerSearch         string `json:"careerSearch"`
	ContentFeedClients   string `json:"contentFeedClients"`
	Educational          string `json:"educational"`
	Financial            string `json:"financial"`
	MediaSearch          string `json:"mediaSearch"`
	News                 string `json:"news"`
	SearchEngineCrawlers string `json:"searchEngineCrawlers"`
}

type WaasContainerBotProtectionSpecRecaptchaSpec struct {
	AllSessions            bool          `json:"allSessions"` // "Friction" in the UI
	Enabled                bool          `json:"enabled"`
	SecretKey              common.Secret `json:"secretKey"`
	SiteKey                string        `json:"siteKey"`
	SuccessExpirationHours int           `json:"successExpirationHours"`
	Type                   string        `json:"type"`
}

type WaasContainerBotProtectionSpecUnknownBotProtectionSpec struct {
	ApiLibraries         string                                                                 `json:"apiLibraries"`
	BotImpersonation     string                                                                 `json:"botImpersonation"`
	BrowserImpersonation string                                                                 `json:"browserImpersonation"`
	Generic              string                                                                 `json:"generic"`
	HttpLibraries        string                                                                 `json:"httpLibraries"`
	RequestAnomalies     WaasContainerBotProtectionSpecUnknownBotProtectionSpecRequestAnomalies `json:"requestAnomalies"`
	WebAutomationTools   string                                                                 `json:"webAutomationTools"`
	WebScrapers          string                                                                 `json:"webScrapers"`
}

type WaasContainerBotProtectionSpecUnknownBotProtectionSpecRequestAnomalies struct {
	Effect    string `json:"effect"`
	Threshold int    `json:"threshold"` // "Lax" == 9, "Moderate" == 6, "Strict" == 3
}

type WaasContainerBotProtectionSpecUserDefinedBot struct {
	Effect       string   `json:"effect"`
	HeaderName   string   `json:"headerName"`
	HeaderValues []string `json:"headerValues"`
	Name         string   `json:"name"`
	Subnets      []string `json:"subnets"`
}

type WaasContainerCustomBlockResponse struct {
	Body    string `json:"body,omitempty"`
	Code    int    `json:"code,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
}

type WaasContainerDosConfig struct {
	Alert                WaasContainerDosConfigThreshold        `json:"alert"`
	Ban                  WaasContainerDosConfigThreshold        `json:"ban"`
	Enabled              bool                                   `json:"enabled"`
	ExcludedNetworkLists []string                               `json:"excludedNetworkLists,omitempty"`
	MatchConditions      []WaasContainerDosConfigMatchCondition `json:"matchConditions,omitempty"`
	TrackSession         bool                                   `json:"trackSession,omitempty"`
}

type WaasContainerDosConfigThreshold struct {
	Average int `json:"average,omitempty"`
	Burst   int `json:"burst,omitempty"`
}

type WaasContainerDosConfigMatchCondition struct {
	FileTypes          []string                                                `json:"fileTypes,omitempty"`
	Methods            []string                                                `json:"methods,omitempty"` // HTTP methods
	ResponseCodeRanges []WaasContainerDosConfigMatchConditionResponseCodeRange `json:"responseCodeRanges,omitempty"`
}

type WaasContainerDosConfigMatchConditionResponseCodeRange struct {
	End   int `json:"end,omitempty"`
	Start int `json:"start"`
}

type WaasContainerHeaderSpec struct {
	Allow    bool     `json:"allow"`
	Effect   string   `json:"effect"`
	Name     string   `json:"name"`
	Required bool     `json:"required,omitempty"`
	Values   []string `json:"values"`
}

type WaasContainerIntelGathering struct {
	InfoLeakageEffect         string `json:"infoLeakageEffect"`
	RemoveFingerprintsEnabled bool   `json:"removeFingerprintsEnabled"`
}

type WaasContainerMaliciousUpload struct {
	AllowedExtensions []string `json:"allowedExtensions"`
	AllowedFileTypes  []string `json:"allowedFileTypes"`
	Effect            string   `json:"effect,omitempty"`
}

type WaasContainerNetworkControls struct {
	AdvancedProtectionEffect string                             `json:"advancedProtectionEffect"`
	Countries                WaasContainerNetworkControlsAccess `json:"countries"`
	ExceptionSubnets         []string                           `json:"exceptionSubnets,omitempty"`
	Subnets                  WaasContainerNetworkControlsAccess `json:"subnets"`
}

type WaasContainerNetworkControlsAccess struct {
	Alert          []string `json:"alert,omitempty"`
	Allow          []string `json:"allow,omitempty"`
	AllowMode      bool     `json:"allowMode,omitempty"`
	Enabled        bool     `json:"enabled"`
	FallbackEffect string   `json:"fallbackEffect"`
	Prevent        []string `json:"prevent,omitempty"`
}

type WaasContainerRemoteHostForwarding struct {
	Enabled bool   `json:"enabled,omitempty"`
	Target  string `json:"target,omitempty"`
}

type WaasContainerTlsConfig struct {
	HstsConfig    WaasContainerTlsConfigHstsConfig `json:"HSTSConfig"`
	MinTlsVersion string                           `json:"minTLSVersion"`
}

type WaasContainerTlsConfigHstsConfig struct {
	Enabled           bool `json:"enabled"`
	IncludeSubdomains bool `json:"includeSubdomains"`
	MaxAgeSeconds     int  `json:"maxAgeSeconds"`
	Preload           bool `json:"preload"`
}

// Get the current WAAS container policy.
func GetWaasContainer(c api.Client) (WaasContainerPolicy, error) {
	var ans WaasContainerPolicy
	if err := c.Request(http.MethodGet, VulnerabilityCiImagesEndpoint, nil, nil, &ans); err != nil {
		return ans, fmt.Errorf("error getting WAAS container policy: %s", err)
	}
	return ans, nil
}

// Update the current WAAS container policy.
func UpdateWaasContainer(c api.Client, policy WaasContainerPolicy) error {
	return c.Request(http.MethodPut, WaasContainerEndpoint, nil, policy, nil)
}
