package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var (
	waasEffectStrings = []string{
		"ban",
		"prevent",
		"alert",
		"allow",
		"disable",
		"reCAPTCHA",
	}

	waasExceptionLocation = []string{
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

	waasReCAPTCHAType = []string{
		"checkbox",
		"invisible",
	}

	waasRequestAnomalyThreshold = []int{
		3,
		6,
		9,
	}

	customrulesAction = []string{
		"audit",
		"incident",
	}
)

func resourcePoliciesFirewallAppContainer() *schema.Resource {
	return &schema.Resource{
		Create: createPolicyFirewallAppContainer,
		Read:   readPolicyFirewallAppContainer,
		Update: updatePolicyFirewallAppContainer,
		Delete: deletePolicyFirewallAppContainer,

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
													Type:         schema.TypeString,
													Optional:     true,
													Default:      "alert",
													Description:  "Effect that will be used in the rule.",
													ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
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
																Type:         schema.TypeString,
																Required:     true,
																Description:  "Exception HTTP field location.",
																ValidateFunc: validation.StringInSlice(waasExceptionLocation, false),
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
													Type:         schema.TypeString,
													Optional:     true,
													Default:      "alert",
													Description:  "Effect that will be used if body is more than inspection size.",
													ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
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
																Type:         schema.TypeString,
																Optional:     true,
																Default:      "alert",
																Description:  "Effect that will be used if the timeout of JS inspection is raised.",
																ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
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
																Type:         schema.TypeString,
																Optional:     true,
																Default:      "disable",
																Description:  "Effect that will be used if an archiving bot is crawling the app.",
																ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
															},
															"business_analytics": {
																Type:         schema.TypeString,
																Optional:     true,
																Default:      "disable",
																Description:  "Effect that will be used if a business analytics bot is crawling the app.",
																ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
															},
															"career_search": {
																Type:         schema.TypeString,
																Optional:     true,
																Default:      "disable",
																Description:  "Effect that will be used if a career search bot is crawling the app.",
																ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
															},
															"content_feed_clients": {
																Type:         schema.TypeString,
																Optional:     true,
																Default:      "disable",
																Description:  "Effect that will be used if a content feed clients bot is crawling the app.",
																ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
															},
															"educational": {
																Type:         schema.TypeString,
																Optional:     true,
																Default:      "disable",
																Description:  "Effect that will be used if an educational bot is crawling the app.",
																ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
															},
															"financial": {
																Type:         schema.TypeString,
																Optional:     true,
																Default:      "disable",
																Description:  "Effect that will be used if a financial bot is crawling the app.",
																ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
															},
															"media_search": {
																Type:         schema.TypeString,
																Optional:     true,
																Default:      "disable",
																Description:  "Effect that will be used if a media search bot is crawling the app.",
																ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
															},
															"news": {
																Type:         schema.TypeString,
																Optional:     true,
																Default:      "disable",
																Description:  "Effect that will be used if a news bot is crawling the app.",
																ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
															},
															"search_engine_crawlers": {
																Type:         schema.TypeString,
																Optional:     true,
																Default:      "disable",
																Description:  "Effect that will be used if a search engine crawler bot is crawling the app.",
																ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
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
																Type:         schema.TypeString,
																Optional:     true,
																Description:  "Encrypted secret key to use when invoking the reCAPTCHA service.",
																ValidateFunc: validation.StringInSlice(waasReCAPTCHAType, false),
															},
														},
													},
												},
												"session_validation": {
													Type:         schema.TypeString,
													Optional:     true,
													Default:      "disable",
													Description:  "Effect that will be used if the session is not validated.",
													ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
												},
												"unknown_bot_protection": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Unknown bot protection configuration.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"api_libraries": {
																Type:         schema.TypeString,
																Optional:     true,
																Default:      "disable",
																Description:  "Effect that will be used if an API library bot is crawling the app.",
																ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
															},
															"bot_impersonation": {
																Type:         schema.TypeString,
																Optional:     true,
																Default:      "disable",
																Description:  "Effect that will be used if a bot impersonation bot is crawling the app.",
																ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
															},
															"browser_impersonation": {
																Type:         schema.TypeString,
																Optional:     true,
																Default:      "disable",
																Description:  "Effect that will be used if a browser impersonation bot is crawling the app.",
																ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
															},
															"generic": {
																Type:         schema.TypeString,
																Optional:     true,
																Default:      "disable",
																Description:  "Effect that will be used if a generic bot is crawling the app.",
																ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
															},
															"http_libraries": {
																Type:         schema.TypeString,
																Optional:     true,
																Default:      "disable",
																Description:  "Effect that will be used if a HTTP library bot is crawling the app.",
																ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
															},
															"request_anomalies_effect": {
																Type:         schema.TypeString,
																Optional:     true,
																Default:      "disable",
																Description:  "Effect that will be used if anomalies are detected in a request done by a bot crawling the app.",
																ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
															},
															"request_anomalies_threshold": {
																Type:         schema.TypeInt,
																Optional:     true,
																Default:      9,
																Description:  "Score threshold for which request anomaly violation is triggered.",
																ValidateFunc: validation.IntInSlice(waasRequestAnomalyThreshold),
															},
															"web_automation_tools": {
																Type:         schema.TypeString,
																Optional:     true,
																Default:      "disable",
																Description:  "Effect that will be used if a web automation tool bot is crawling the app.",
																ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
															},
															"web_scrapers": {
																Type:         schema.TypeString,
																Optional:     true,
																Default:      "disable",
																Description:  "Effect that will be used if a web scraper bot is crawling the app.",
																ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
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
																Type:         schema.TypeString,
																Optional:     true,
																Default:      "alert",
																Description:  "Effect that will be used if the user-defined bot is detected.",
																ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
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
									"cmdi": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Protection against command injection.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"effect": {
													Type:         schema.TypeString,
													Optional:     true,
													Default:      "alert",
													Description:  "Effect if a command injection is detected.",
													ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
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
																Type:         schema.TypeString,
																Required:     true,
																Description:  "Exception HTTP field location.",
																ValidateFunc: validation.StringInSlice(waasExceptionLocation, false),
															},
														},
													},
												},
											},
										},
									},
									"code_injection": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Protection against code injection.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"effect": {
													Type:         schema.TypeString,
													Optional:     true,
													Default:      "alert",
													Description:  "Effect if a code injection is detected.",
													ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
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
																Type:         schema.TypeString,
																Required:     true,
																Description:  "Exception HTTP field location.",
																ValidateFunc: validation.StringInSlice(waasExceptionLocation, false),
															},
														},
													},
												},
											},
										},
									},
									"csrf_enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates whether Cross-Site Request Forgery (CSRF) protection is enabled (true) or not (false).",
									},
									"custom_block_response": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Custom block message config for a policy.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"body": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "Custom HTML for the block response.",
												},
												"code": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     403,
													Description: "Custom HTTP response code for the block response.",
												},
												"enabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Default:     false,
													Description: "Indicates if the custom block response is enabled (true) or not (false).",
												},
											},
										},
									},
									"custom_rules": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "List of custom runtime rules.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "Custom rule ID.",
												},
												"action": {
													Type:         schema.TypeString,
													Required:     true,
													Description:  "Action is the action to perform if the custom rule applies.",
													ValidateFunc: validation.StringInSlice(customrulesAction, false),
												},
												"effect": {
													Type:         schema.TypeString,
													Required:     true,
													Description:  "Effect if a code injection is detected.",
													ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
												},
											},
										},
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

func createPolicyFirewallAppContainer(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func readPolicyFirewallAppContainer(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func updatePolicyFirewallAppContainer(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deletePolicyFirewallAppContainer(d *schema.ResourceData, meta interface{}) error {
	return nil
}
