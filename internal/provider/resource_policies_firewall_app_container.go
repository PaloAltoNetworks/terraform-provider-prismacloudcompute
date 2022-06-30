package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Validation slices of string, int
var (
	waasEffectStrings = []string{
		"ban",
		"prevent",
		"alert",
		"allow",
		"disable",
		"reCAPTCHA",
	}

	waasExceptionLocationStrings = []string{
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

	waasReCAPTCHATypeStrings = []string{
		"checkbox",
		"invisible",
	}

	waasRequestAnomalyThresholdInts = []int{
		3,
		6,
		9,
	}

	customrulesActionStrings = []string{
		"audit",
		"incident",
	}

	waasFileType = []string{
		"pdf",
		"officeLegacy",
		"officeOoxml",
		"odf",
		"jpeg",
		"png",
		"gif",
		"bmp",
		"ico",
		"avi",
		"mp4",
		"aac",
		"mp3",
		"wav",
		"zip",
		"gzip",
		"rar",
		"7zip",
	}

	waasSameSiteStrings = []string{
		"Lax",
		"Strict",
		"None",
	}

	waasMinTLSVersiontrings = []string{
		"1.0",
		"1.1",
		"1.2",
		"1.3",
	}

	waasParamLocationStrings = []string{
		"path",
		"query",
		"cookie",
		"header",
		"body",
		"json",
		"xml",
		"formData",
		"multipart",
	}

	waasParamStyle = []string{
		"simple",
		"spaceDelimited",
		"tabDelimited",
		"pipeDelimited",
		"form",
		"matrix",
		"label",
	}

	waasParamType = []string{
		"integer",
		"number",
		"string",
		"boolean",
		"array",
		"object",
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
												"effect": {
													Type:         schema.TypeString,
													Optional:     true,
													Default:      "alert",
													Description:  "Effect that will be used in the rule.",
													ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
												},
												"endpoints": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "The app's endpoints.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"base_path": {
																Type:        schema.TypeString,
																Required:    true,
																Default:     "/",
																Description: "Base path for the endpoint.",
															},
															"exposed_port": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Exposed port that the proxy is listening on.",
															},
															"grpc": {
																Type:        schema.TypeBool,
																Optional:    true,
																Default:     false,
																Description: "Indicates if the proxy supports gRPC (true) or not (false).",
															},
															"host": {
																Type:        schema.TypeString,
																Optional:    true,
																Default:     "*",
																Description: "URL address (name or IP) of the endpoint's API specification (e.g., petstore.swagger.io). The address can be prefixed with a wildcard (e.g., *.swagger.io).",
															},
															"http2": {
																Type:        schema.TypeBool,
																Optional:    true,
																Default:     false,
																Description: "Indicates if the proxy supports HTTP/2 (true) or not (false).",
															},
															"internal_port": {
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Internal port that the application is listening on.",
															},
															"tls": {
																Type:        schema.TypeBool,
																Optional:    true,
																Default:     false,
																Description: "Indicates if the connection is secured (true) or not (false).",
															},
														},
													},
												},
												"fallback_effect": {
													Type:         schema.TypeString,
													Optional:     true,
													Default:      "alert",
													Description:  "Fallback effect that will be used in the rule.",
													ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
												},
												"paths": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Paths of the API's endpoints.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"methods": {
																Type:        schema.TypeList,
																Required:    true,
																Description: "Supported operations for the path (e.g., PUT, GET, etc.).",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"method": {
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "Type of HTTP request (e.g., PUT, GET, etc.).",
																		},
																		"parameters": {
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "Parameters that are part of the HTTP request.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"allow_empty_value": {
																						Type:        schema.TypeBool,
																						Optional:    true,
																						Default:     false,
																						Description: "Indicates if the proxy supports HTTP/2 (true) or not (false).",
																					},
																					"array": {
																						Type:        schema.TypeBool,
																						Optional:    true,
																						Default:     false,
																						Description: "Indicates if multiple values of the specified type are allowed (true) or not (false).",
																					},
																					"explode": {
																						Type:        schema.TypeBool,
																						Optional:    true,
																						Default:     false,
																						Description: "Indicates if arrays should generate separate parameters for each array item or object property.",
																					},
																					"location": {
																						Type:         schema.TypeString,
																						Required:     true,
																						Description:  "The location of a parameter.",
																						ValidateFunc: validation.StringInSlice(waasParamLocationStrings, false),
																					},
																					"max": {
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Description: "Maximum allowable value for a numeric parameter.",
																					},
																					"min": {
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Description: "Maximum allowable value for a numeric parameter.",
																					},
																					"name": {
																						Type:        schema.TypeString,
																						Required:    true,
																						Description: "Name of the parameter.",
																					},
																					"required": {
																						Type:        schema.TypeBool,
																						Optional:    true,
																						Default:     false,
																						Description: "Indicates if the parameter is required (true) or not (false).",
																					},
																					"style": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						Description:  "Param format style, defined by OpenAPI specification It describes how the parameter value will be serialized depending on the type of the parameter.",
																						ValidateFunc: validation.StringInSlice(waasParamStyle, false),
																					},
																					"type": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						Description:  "Type of a parameter, defined by OpenAPI specification.",
																						ValidateFunc: validation.StringInSlice(waasParamType, false),
																					},
																				},
																			},
																		},
																	},
																},
															},
															"path": {
																Type:        schema.TypeString,
																Required:    true,
																Description: "Relative path to an endpoint such as '/pet/{petId}'.",
															},
														},
													},
												},

												"query_param_effect": {
													Type:         schema.TypeString,
													Optional:     true,
													Default:      "alert",
													Description:  "Query param effect that will be used in the rule.",
													ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
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
																ValidateFunc: validation.StringInSlice(waasExceptionLocationStrings, false),
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
																ValidateFunc: validation.StringInSlice(waasReCAPTCHATypeStrings, false),
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
																ValidateFunc: validation.IntInSlice(waasRequestAnomalyThresholdInts),
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
									"command_injection": {
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
																ValidateFunc: validation.StringInSlice(waasExceptionLocationStrings, false),
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
																ValidateFunc: validation.StringInSlice(waasExceptionLocationStrings, false),
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
													ValidateFunc: validation.StringInSlice(customrulesActionStrings, false),
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
									"disabled_event_id_header": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     true,
										Description: "Indicates if event ID header should be attached to the response or not.",
									},
									"dos_config": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Protect against DOS attacks.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"alert_avg": {
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Average request rate (requests / second) threshold before alerting.",
												},
												"alert_burst": {
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Burst request rate (requests / second) threshold before alerting.",
												},
												"ban_avg": {
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Average request rate (requests / second) threshold before baning.",
												},
												"ban_burst": {
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Burst request rate (requests / second) threshold before baning.",
												},
												"enabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Default:     false,
													Description: "Indicates if the rule is disabled or not.",
												},
												"excluded_network_lists": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Network IPs to exclude from DoS tracking.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"match_conditions": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"file_types": {
																Type:        schema.TypeList,
																Optional:    true,
																Description: "File types for request matching.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"methods": {
																Type:        schema.TypeList,
																Optional:    true,
																Description: "HTTP methods for request matching.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"response_code_ranges": {
																Type:        schema.TypeList,
																Optional:    true,
																Description: "Response codes for the request's response matching.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"end": {
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "End of the range. Can be omitted if using a single status code.",
																		},
																		"start": {
																			Type:        schema.TypeList,
																			Required:    true,
																			Description: "Start of the range. Can also be used for a single, non-range value.",
																		},
																	},
																},
															},
														},
													},
												},
												"track_session": {
													Type:        schema.TypeBool,
													Optional:    true,
													Default:     false,
													Description: "Indicates if the custom session ID generated during bot protection flow is tracked (true) or not (false).",
												},
											},
										},
									},
									"headers": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Configuration for inspecting HTTP headers.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"allow": {
													Type:        schema.TypeBool,
													Optional:    true,
													Default:     false,
													Description: "Indicates if the flow is to be allowed (true) or blocked (false).",
												},
												"effect": {
													Type:         schema.TypeString,
													Required:     true,
													Default:      "alert",
													Description:  "Effect if a header is found.",
													ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
												},
												"name": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "Header name.",
												},
												"required": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Indicates if the header must be present (true) or not (false).",
												},
												"values": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Wildcard expressions that represent the header value.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"intel_gathering": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Configuration for intelligence gathering protections.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"info_leakage_effect": {
													Type:         schema.TypeString,
													Required:     true,
													Default:      "alert",
													Description:  "Effect if a leakage is found.",
													ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
												},
												"remove_fingerprint_enabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Indicates if server fingerprints should be removed (true) or not (false).",
												},
											},
										},
									},
									"local_file_inclusion": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Protection against local file inclusion.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"effect": {
													Type:         schema.TypeString,
													Optional:     true,
													Default:      "alert",
													Description:  "Effect if a local file inclusion is detected.",
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
																ValidateFunc: validation.StringInSlice(waasExceptionLocationStrings, false),
															},
														},
													},
												},
											},
										},
									},
									"malformed_requests": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Protection against malformed requests.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"effect": {
													Type:         schema.TypeString,
													Optional:     true,
													Default:      "alert",
													Description:  "Effect if a malformed request is detected.",
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
																ValidateFunc: validation.StringInSlice(waasExceptionLocationStrings, false),
															},
														},
													},
												},
											},
										},
									},
									"malicious_uploads": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Protection against malicious uploads.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"allowed_extensions": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Allowed file extensions.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"allowed_file_types": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Allowed file types.",
													Elem: &schema.Schema{
														Type:         schema.TypeString,
														ValidateFunc: validation.StringInSlice(waasFileType, false),
													},
												},
												"effect": {
													Type:         schema.TypeString,
													Optional:     true,
													Default:      "alert",
													Description:  "Effect if a malicious upload is detected.",
													ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
												},
											},
										},
									},
									"network_controls": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Protection using access controls for IPs and countries.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"advanced_protection_effect": {
													Type:         schema.TypeString,
													Optional:     true,
													Default:      "alert",
													Description:  "Effect if controlled IPs / contries are detected.",
													ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
												},
												"countries": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Countries managed by access control.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"alert": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Alert are the denied countries for which we alert.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"allow": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Allow are the allowed countries for which we don't alert or prevent.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"allow_mode": {
																Type:        schema.TypeBool,
																Optional:    true,
																Default:     false,
																Description: "Indicates allowlist (true) or denylist (false) mode.",
															},
															"enabled": {
																Type:        schema.TypeBool,
																Optional:    true,
																Default:     false,
																Description: "Enabled indicates if access controls protection is enabled.",
															},
															"fallback_effect": {
																Type:         schema.TypeString,
																Optional:     true,
																Default:      "alert",
																Description:  "Effect if an access control is detected.",
																ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
															},
															"prevent": {
																Type:        schema.TypeList,
																Optional:    true,
																Description: "The denied countries.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
												"exception_subnets": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Network lists for which requests completely bypass WAAS checks and protections.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"subnets": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Subnets managed by access control.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"alert": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Alert are the denied subnets for which we alert.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"allow": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Allow are the allowed subnets for which we don't alert or prevent.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"allow_mode": {
																Type:        schema.TypeBool,
																Optional:    true,
																Default:     false,
																Description: "Indicates allowlist (true) or denylist (false) mode.",
															},
															"enabled": {
																Type:        schema.TypeBool,
																Optional:    true,
																Default:     false,
																Description: "Enabled indicates if access controls protection is enabled.",
															},
															"fallback_effect": {
																Type:         schema.TypeString,
																Optional:     true,
																Default:      "alert",
																Description:  "Effect if an access control is detected.",
																ValidateFunc: validation.StringInSlice(waasEffectStrings, false),
															},
															"prevent": {
																Type:        schema.TypeList,
																Optional:    true,
																Description: "The denied subnets.",
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
									"remote_host_forwarding": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Remote host to forward requests to.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Default:     false,
													Description: "Indicates if remote host forwarding is enabled (true) or not (false).",
												},
												"target": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "Remote host to forward requests to.",
												},
											},
										},
									},
									"response_headers": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Configuration for modifying HTTP response headers.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "Header name (will be canonicalized when possible).",
												},
												"override": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Indicates whether to override existing values (true) or add to them (false).",
												},
												"values": {
													Type:        schema.TypeList,
													Required:    true,
													Description: "New header values.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"session_cookie_ban": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "Indicates if bans in this app are made by session cookie ID (true) or false (not).",
									},
									"session_cookie_enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "Indicates if session cookies are enabled (true) or not (false).",
									},
									"session_cookie_samesite": {
										Type:         schema.TypeString,
										Optional:     true,
										Description:  "SameSite allows a server to define a cookie attribute making it impossible for the browser to send this cookie along with cross-site requests. The main goal is to mitigate the risk of cross-origin information leakage, and provide some protection against cross-site request forgery attacks.",
										ValidateFunc: validation.StringInSlice(waasSameSiteStrings, false),
									},
									"session_cookie_secure": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "Indicates the Secure attribute of the session cookie.",
									},
									"shellshock": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Protection against shellshock requests.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"effect": {
													Type:         schema.TypeString,
													Optional:     true,
													Default:      "alert",
													Description:  "Effect if a shellshock request is detected.",
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
																ValidateFunc: validation.StringInSlice(waasExceptionLocationStrings, false),
															},
														},
													},
												},
											},
										},
									},
									"sql_injection": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Protection against SQL injection requests.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"effect": {
													Type:         schema.TypeString,
													Optional:     true,
													Default:      "alert",
													Description:  "Effect if a SQL injection request is detected.",
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
																ValidateFunc: validation.StringInSlice(waasExceptionLocationStrings, false),
															},
														},
													},
												},
											},
										},
									},
									"tls_config": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "The user TLS configuration and the certificate data.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"hsts_config": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "The HTTP Strict Transport Security configuration in order to enforce HSTS header.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"enabled": {
																Type:        schema.TypeBool,
																Optional:    true,
																Default:     false,
																Description: "Enabled indicates if HSTS enforcement is enabled.",
															},
															"include_subdomains": {
																Type:        schema.TypeString,
																Optional:    true,
																Default:     false,
																Description: "Indicates if this rule applies to all of the site's subdomains as well.",
															},
															"max_age": {
																Type:        schema.TypeInt,
																Optional:    true,
																Default:     31536000,
																Description: "The time (in seconds) that the browser should remember that a site is only be accessed using HTTPS.",
															},
															"preload": {
																Type:        schema.TypeBool,
																Optional:    true,
																Default:     false,
																Description: "Indicates if it should support preload.",
															},
														},
													},
												},
												"metadata": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Computed:    true,
													Description: "The user TLS configuration and the certificate data.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"issuer_name": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The certificate issuer common name.",
															},
															"not_after": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The time the certificate is not valid (expiry time).",
															},
															"subject_name": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The certificate subject common name.",
															},
														},
													},
												},
												"min_tls_version": {
													Type:         schema.TypeString,
													Optional:     true,
													Default:      "1.2",
													Description:  "MinTLSVersion is the list of acceptable TLS versions.",
													ValidateFunc: validation.StringInSlice(waasMinTLSVersiontrings, false),
												},
											},
										},
									},
									"cross_site_scripting": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Protection against Cross Site Scripting (XSS) requests.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"effect": {
													Type:         schema.TypeString,
													Optional:     true,
													Default:      "alert",
													Description:  "Effect if a Cross Site Scripting (XSS) request is detected.",
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
																ValidateFunc: validation.StringInSlice(waasExceptionLocationStrings, false),
															},
														},
													},
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
							Description: "List of collections. Used to scope the rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"account_ids": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of account IDs.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"app_ids": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of application IDs.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"clusters": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of Kubernetes cluster names.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"color": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Hexadecimal representation of color code value.",
									},
									"containers": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of containers.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"description": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Free-form text..",
									},
									"functions": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of functions.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"hosts": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of hosts.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"images": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of images.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"labels": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of labels.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"modified": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Datetime when the collection was last modified.",
									},
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Collection name. Must be unique.",
									},
									"namespaces": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of Kubernetes namespaces.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"owner": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "User who created or last modified the collection.",
									},
									"prisma": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Indicates whether this collection originates from Prisma Cloud.",
									},
									"system": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Indicates whether this collection was created by the system (i.e., a non user) (true) or a real user (false).",
									},
								},
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
