package provider

import (
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePoliciesFirewallAppContainer() *schema.Resource {
	return &schema.Resource{
		CreateContext: createPolicyFirewallAppContainer,
		ReadContext:   readPolicyFirewallAppContainer,
		UpdateContext: updatePolicyFirewallAppContainer,
		DeleteContext: deletePolicyFirewallAppContainer,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique internal ID.",
			},
			"max_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     31000,
				Description: "Maximum port number to use in the application firewall.",
			},
			"min_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     30000,
				Description: "Minimum port number to use in the application firewall.",
			},
			"rules": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Rules in the policy.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_malformed_http_header_names": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates if validation of http request header names should allow non-compliant characters.",
						},
						"applications": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of application specifications in the rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"api": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "API specification.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"description": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Description of the app.",
												},
											},
										},
									},
									"app_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Datetime when the rule was last modified.",
									},
									"attack_tools": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Protection against attack tools and vulnerability scanners.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"effect": {
													Type:             schema.TypeString,
													Optional:         true,
													Default:          "alert",
													Description:      "Effect that will be used in the rule.",
													ValidateDiagFunc: validateWaasEffect,
												},
												"exceptions": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": {
																Type:        schema.TypeString,
																Required:    true,
																Description: "Field in HTTP request.",
															},
															"location": {
																Type:             schema.TypeString,
																Required:         true,
																Description:      "Exception HTTP field location.",
																ValidateDiagFunc: validateWaasExceptionLocation,
															},
														},
													},
												},
											},
										},
									},
									"ban_duration": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     5,
										Description: "Ban duration, in minutes.",
									},
									"body": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "App configuration related to HTTP body.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"inspection_limit_exceeded_effect": {
													Type:             schema.TypeString,
													Optional:         true,
													Default:          "alert",
													Description:      "Effect that will be used if body is more than inspection size.",
													ValidateDiagFunc: validateWaasEffect,
												},
												"inspection_size": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     131072,
													Description: "Max amount of data to inspect in request body.",
												},
												"skip": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Indicates that body inspection should be skipped.",
												},
											},
										},
									},
									"bot_protection": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Protection against bots.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"interstitial_page": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Indicates if an interstitial page is served (true) or not (false).",
												},
												"js_inspection": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "JS inspection configuration.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"enabled": {
																Type:        schema.TypeBool,
																Optional:    true,
																Default:     false,
																Description: "Indicates if JavaScript injection is enabled (true) or not (false).",
															},
															"timeout_effect": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "alert",
																Description:      "Effect that will be used if the timeout of JS inspection is raised.",
																ValidateDiagFunc: validateWaasEffect,
															},
														},
													},
												},
												"known_bot_protections": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Known bot protections configuration.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"archiving": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																Description:      "Effect that will be used if an archiving bot is crawling the app.",
																ValidateDiagFunc: validateWaasEffect,
															},
															"business_analytics": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																Description:      "Effect that will be used if a business analytics bot is crawling the app.",
																ValidateDiagFunc: validateWaasEffect,
															},
															"career_search": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																Description:      "Effect that will be used if a career search bot is crawling the app.",
																ValidateDiagFunc: validateWaasEffect,
															},
															"content_feed_clients": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																Description:      "Effect that will be used if a content feed clients bot is crawling the app.",
																ValidateDiagFunc: validateWaasEffect,
															},
															"educational": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																Description:      "Effect that will be used if an educational bot is crawling the app.",
																ValidateDiagFunc: validateWaasEffect,
															},
															"financial": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																Description:      "Effect that will be used if a financial bot is crawling the app.",
																ValidateDiagFunc: validateWaasEffect,
															},
															"media_search": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																Description:      "Effect that will be used if a media search bot is crawling the app.",
																ValidateDiagFunc: validateWaasEffect,
															},
															"news": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																Description:      "Effect that will be used if a news bot is crawling the app.",
																ValidateDiagFunc: validateWaasEffect,
															},
															"search_engine_crawlers": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																Description:      "Effect that will be used if a search engine crawler bot is crawling the app.",
																ValidateDiagFunc: validateWaasEffect,
															},
														},
													},
												},
												"recaptcha": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "reCaptcha protections configuration.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"all_sessions": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Indicates if the reCAPTCHA page is served at the start of every new session (true) or not (false).",
															},
															"enabled": {
																Type:        schema.TypeBool,
																Optional:    true,
																Default:     false,
																Description: "Indicates if reCAPTCHA integration is enabled (true) or not (false).",
															},
															"secret_key": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Secret key to use when invoking the reCAPTCHA service.",
															},
															"site_key": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Site key to use when invoking the reCAPTCHA service.",
															},
															"success_expiration": {
																Type:        schema.TypeInt,
																Optional:    true,
																Default:     24,
																Description: "Duration (hours) for which the indication of reCAPTCHA success is kept. Maximum value is 30 days * 24 = 720 hours.",
															},
															"type": {
																Type:             schema.TypeString,
																Optional:         true,
																Description:      "Encrypted secret key to use when invoking the reCAPTCHA service.",
																ValidateDiagFunc: validateWaasReCAPTCHAType,
															},
														},
													},
												},
												"session_validation": {
													Type:             schema.TypeString,
													Optional:         true,
													Default:          "disable",
													Description:      "Effect that will be used if the session is not validated.",
													ValidateDiagFunc: validateWaasEffect,
												},
												"unknown_bot_protection": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Unknown bot protection configuration.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"api_libraries": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																Description:      "Effect that will be used if an API library bot is crawling the app.",
																ValidateDiagFunc: validateWaasEffect,
															},
															"bot_impersonation": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																Description:      "Effect that will be used if a bot impersonation bot is crawling the app.",
																ValidateDiagFunc: validateWaasEffect,
															},
															"browser_impersonation": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																Description:      "Effect that will be used if a browser impersonation bot is crawling the app.",
																ValidateDiagFunc: validateWaasEffect,
															},
															"generic": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																Description:      "Effect that will be used if a generic bot is crawling the app.",
																ValidateDiagFunc: validateWaasEffect,
															},
															"http_libraries": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																Description:      "Effect that will be used if a HTTP library bot is crawling the app.",
																ValidateDiagFunc: validateWaasEffect,
															},
															"request_anomalies_effect": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																Description:      "Effect that will be used if anomalies are detected in a request done by a bot crawling the app.",
																ValidateDiagFunc: validateWaasEffect,
															},
															"request_anomalies_threshold": {
																Type:             schema.TypeInt,
																Optional:         true,
																Default:          9,
																Description:      "Score threshold for which request anomaly violation is triggered.",
																ValidateDiagFunc: validateWaasRequestAnomalyThreshold,
															},
															"web_automation_tools": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																Description:      "Effect that will be used if a web automation tool bot is crawling the app.",
																ValidateDiagFunc: validateWaasEffect,
															},
															"web_scrapers": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																Description:      "Effect that will be used if a web scraper bot is crawling the app.",
																ValidateDiagFunc: validateWaasEffect,
															},
														},
													},
												},
												"user_defined_bot_protection": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Effects to perform when user-defined bots are detected.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"effect": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "alert",
																Description:      "Effect that will be used if the user-defined bot is detected.",
																ValidateDiagFunc: validateWaasEffect,
															},
															"header_name": {
																Type:        schema.TypeString,
																Required:    true,
																Description: "HTTP header name to recognize the user-defined bot.",
															},
															"header_values": {
																Type:        schema.TypeList,
																Optional:    true,
																Description: "Header values.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"name": {
																Type:        schema.TypeString,
																Required:    true,
																Description: "Name of the user-defined bot.",
															},
															"subnets": {
																Type:        schema.TypeList,
																Optional:    true,
																Description: "Subnets where the bot originates. Specify using network lists.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
											},
										},
									},
									"certificate": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Certificate used to decrypt TLS.",
									},
									"clickjackingEnabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates whether clickjacking protection is enabled (true) or not (false).",
									},
								},
							},
						},
						"auto_protect_ports": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates if http ports should be automatically detected and protected.",
						},
						"collections": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "Collections used to scope the rule.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"disabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates if the rule is disabled or not.",
						},
						"modified": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Datetime when the rule was last modified.",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name of the rule.",
						},
						"notes": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Free-form text.",
						},
						"owmer": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User who created or last modified the rule.",
						},
						"previous_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Previous name of the rule. Required for rule renaming.",
						},
						"read_timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     5,
							Description: "Timeout of request reads in seconds, when no value is specified (0) the timeout is 5 seconds.",
						},
						"skip_api_learning": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates if API discovery is to be skipped (true) or not (false).",
						},
						"traffic_mirroring_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates if if traffic mirroring is enabled.",
						},
						"windows": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether the operating system of the app is windows, default is Linux.",
						},
					},
				},
			},
		},
	}
}

func validateWaasEffect(i interface{}, path cty.Path) diag.Diagnostics {
	v, ok := i.(string)
	if !ok {
		return diag.Errorf("expected type to be string")
	}
	validStrings := []string{
		"ban",
		"prevent",
		"alert",
		"allow",
		"disable",
		"reCAPTCHA",
	}
	if !stringInSlice(v, validStrings) {
		return diag.Errorf("%q must be one of %v", v, validStrings)
	}
	return nil
}

func validateWaasExceptionLocation(i interface{}, path cty.Path) diag.Diagnostics {
	v, ok := i.(string)
	if !ok {
		return diag.Errorf("expected type to be string")
	}
	validStrings := []string{
		"path",
		"query",
		"queryValues",
		"cookie",
		"UserAgentHeader",
		"header",
		"body",
		"rawBody",
		"XMLPath",
		"JSONPath",
	}
	if !stringInSlice(v, validStrings) {
		return diag.Errorf("%q must be one of %v", v, validStrings)
	}
	return nil
}

func validateWaasReCAPTCHAType(i interface{}, path cty.Path) diag.Diagnostics {
	v, ok := i.(string)
	if !ok {
		return diag.Errorf("expected type to be string")
	}
	validStrings := []string{
		"checkbox",
		"invisible",
	}
	if !stringInSlice(v, validStrings) {
		return diag.Errorf("%q must be one of %v", v, validStrings)
	}
	return nil
}

func validateWaasRequestAnomalyThreshold(i interface{}, path cty.Path) diag.Diagnostics {
	v, ok := i.(int)
	if !ok {
		return diag.Errorf("expected type to be int")
	}
	validIntegers := []int{
		3,
		6,
		9,
	}
	if !intInSlice(v, validIntegers) {
		return diag.Errorf("%q must be one of %v", v, validIntegers)
	}
	return nil
}
