package provider

import (
	"context"
	"log"
	"strings"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/policy"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/convert"
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     31000,
				Description: "The maximum port number available to use by the firewall.",
			},
			"min_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     30000,
				Description: "The minimum port number available to use by the firewall.",
			},
			"rule": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Rules that make up the policy.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"notes": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"collections": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "Collections used to scope the rule.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"read_timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     5,
							Description: "The read timeout in seconds.",
						},
						"application": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"app_definition": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"app_id": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "Unique identifier for the application.",
												},
												"endpoint_setup": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"description": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"api_endpoint_discovery_enabled": { // SkipLearning
																Type:        schema.TypeBool,
																Optional:    true,
																Default:     true,
																Description: "Whether or not to automatically discover API endpoints.",
															},
															"tls_config": {
																Type:        schema.TypeList,
																Optional:    true,
																Description: "",
																MaxItems:    1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"certificate": {
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Concatenated certificate and key.",
																		},
																		"min_tls_version": {
																			Type:             schema.TypeString,
																			Optional:         true,
																			Default:          "1.2",
																			ValidateDiagFunc: validateTlsVersion,
																		},
																		"hsts_config": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"enabled": {
																						Type:     schema.TypeBool,
																						Optional: true,
																						Default:  false,
																					},
																					"max_age": {
																						Type:        schema.TypeInt,
																						Optional:    true,
																						Default:     31536000,
																						Description: "Number of seconds that user agents regard the web server as a known HSTS host.",
																					},
																					"include_subdomains": {
																						Type:     schema.TypeBool,
																						Optional: true,
																						Default:  false,
																					},
																					"preload": {
																						Type:     schema.TypeBool,
																						Optional: true,
																						Default:  false,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"endpoint": {
																Type:     schema.TypeList,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"host": {
																			Type:        schema.TypeString,
																			Required:    true,
																			Description: "DNS name or IP address.",
																		},
																		"external_port": {
																			Type:             schema.TypeInt,
																			Required:         true,
																			Description:      "The port number that is exposed to the internet.",
																			ValidateDiagFunc: validatePort,
																		},
																		"base_path": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																		"internal_port": {
																			Type:             schema.TypeInt,
																			Required:         true,
																			Description:      "The internal port that your app listens on. WAAS sends traffic to this port.",
																			ValidateDiagFunc: validatePort,
																		},
																		"tls": {
																			Type:     schema.TypeBool,
																			Optional: true,
																			Default:  false,
																		},
																		"grpc": {
																			Type:     schema.TypeBool,
																			Optional: true,
																			Default:  false,
																		},
																		"http2": {
																			Type:     schema.TypeBool,
																			Optional: true,
																			Default:  false,
																		},
																	},
																},
															},
														},
													},
												},
												"api_protection": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"violation_effect": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																Description:      "The effect to take if a violation is encountered.",
																ValidateDiagFunc: validateStandardEffect,
															},
															"unspecified_api_effect": { // FallbackEffect
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																Description:      "The effect to take if an unspecified API path or method is encountered.",
																ValidateDiagFunc: validateStandardEffect,
															},
															"unspecified_query_params_effect": { // QueryParamFallbackEffect
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																Description:      "The effect to take if an unspecified query parameter is encountered. Only evaluated if there are no API violations.",
																ValidateDiagFunc: validateStandardEffect,
															},
															"path": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"path": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																		"method": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"method": {
																						Type:             schema.TypeString,
																						Required:         true,
																						ValidateDiagFunc: validateHttpMethod,
																					},
																					"parameter": {
																						Type:     schema.TypeList,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"name": {
																									Type:     schema.TypeString,
																									Required: true,
																								},
																								"type": {
																									Type:             schema.TypeString,
																									Required:         true,
																									ValidateDiagFunc: validateParameterType,
																								},
																								"location": {
																									Type:             schema.TypeString,
																									Required:         true,
																									ValidateDiagFunc: validateParameterLocation,
																								},
																								"style": {
																									Type:             schema.TypeString,
																									Required:         true,
																									ValidateDiagFunc: validateParameterStyle,
																								},
																								"min": {
																									Type:        schema.TypeInt,
																									Optional:    true,
																									Description: "The minimum value allowed. Only applicable when type is integer or number.",
																								},
																								"max": {
																									Type:        schema.TypeInt,
																									Optional:    true,
																									Description: "The maximum value allowed. Only applicable when type is integer or number.",
																								},
																								"explode": {
																									Type:     schema.TypeBool,
																									Optional: true,
																								},
																								"array": {
																									Type:     schema.TypeBool,
																									Optional: true,
																								},
																								"required": {
																									Type:     schema.TypeBool,
																									Optional: true,
																								},
																								"allow_empty_value": {
																									Type:     schema.TypeBool,
																									Optional: true,
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
														},
													},
												},
											},
										},
									},
									"app_firewall": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"sql_injection": {
													Type:     schema.TypeList,
													Required: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"effect": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "alert",
																ValidateDiagFunc: validateStandardEffect,
															},
															"exception": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"location": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																		"value": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},
														},
													},
												},
												"xss": {
													Type:     schema.TypeList,
													Required: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"effect": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "alert",
																ValidateDiagFunc: validateStandardEffect,
															},
															"exception": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"location": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																		"value": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},
														},
													},
												},
												"command_injection": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"effect": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "alert",
																ValidateDiagFunc: validateStandardEffect,
															},
															"exception": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"location": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																		"value": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},
														},
													},
												},
												"code_injection": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"effect": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "alert",
																ValidateDiagFunc: validateStandardEffect,
															},
															"exception": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"location": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																		"value": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},
														},
													},
												},
												"local_file_inclusion": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"effect": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "alert",
																ValidateDiagFunc: validateStandardEffect,
															},
															"exception": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"location": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																		"value": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},
														},
													},
												},
												"attack_tools": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"effect": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "alert",
																ValidateDiagFunc: validateStandardEffect,
															},
															"exception": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"location": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																		"value": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},
														},
													},
												},
												"shellshock": {
													Type:     schema.TypeList,
													Required: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"effect": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "alert",
																ValidateDiagFunc: validateStandardEffect,
															},
														},
													},
												},
												"malformed_request": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"effect": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "alert",
																ValidateDiagFunc: validateStandardEffect,
															},
														},
													},
												},
												"advanced_threat_protection": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"effect": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "alert",
																ValidateDiagFunc: validateStandardEffect,
															},
														},
													},
												},
												"information_leakage": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"effect": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																ValidateDiagFunc: validateStandardEffect,
															},
														},
													},
												},
												"csrf_protection_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  true,
												},
												"clickjacking_prevention_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  true,
												},
												"remove_fingerprints": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  true,
												},
											},
										},
									},
									"dos_protection": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  false,
												},
												"alert": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"average": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"burst": {
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
												"ban": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"average": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"burst": {
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
												"match_condition": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"methods": {
																Type:     schema.TypeList,
																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"file_types": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"response_code_range": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"start": {
																			Type:     schema.TypeInt,
																			Required: true,
																		},
																		"end": {
																			Type:     schema.TypeInt,
																			Required: true,
																		},
																	},
																},
															},
														},
													},
												},
												"excluded_networks": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"access_control": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"network_controls": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ip_access_control": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"enabled": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																		"allowed": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Whether or not to use an allow list or alert/prevent lists.",
																		},
																		"allowed_network_lists": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"alerted_network_lists": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"prevented_network_lists": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"fallback_effect": {
																			Type:             schema.TypeString,
																			Optional:         true,
																			ValidateDiagFunc: validateAlertPreventEffect,
																		},
																	},
																},
															},
															"geo_access_control": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"enabled": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																		"allowed": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Whether or not to use an allow list or alert/prevent lists.",
																		},
																		"allowed_countries": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"alerted_countries": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"prevented_countries": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"fallback_effect": {
																			Type:             schema.TypeString,
																			Optional:         true,
																			ValidateDiagFunc: validateAlertPreventEffect,
																		},
																	},
																},
															},
															"network_list_exceptions": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
												"http_header": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
															"allowed": {
																Type:        schema.TypeBool,
																Required:    true,
																Description: "Whether the specified values are allowed or blocked.",
															},
															"values": {
																Type:     schema.TypeList,
																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"effect": {
																Type:             schema.TypeString,
																Required:         true,
																ValidateDiagFunc: validateAllowPreventEffect,
															},
															"required": {
																Type:     schema.TypeBool,
																Optional: true,
															},
														},
													},
												},
												"file_uploads": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"allowed_extensions": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"allowed_file_types": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"effect": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "alert",
																ValidateDiagFunc: validateAlertPreventEffect,
															},
														},
													},
												},
											},
										},
									},
									"bot_protection": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"known_bots": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"search_engine_crawlers": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																ValidateDiagFunc: validateStandardEffect,
															},
															"business_analytics": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																ValidateDiagFunc: validateStandardEffect,
															},
															"educational": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																ValidateDiagFunc: validateStandardEffect,
															},
															"news": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																ValidateDiagFunc: validateStandardEffect,
															},
															"financial": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																ValidateDiagFunc: validateStandardEffect,
															},
															"content_feed_clients": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																ValidateDiagFunc: validateStandardEffect,
															},
															"archiving": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																ValidateDiagFunc: validateStandardEffect,
															},
															"career_search": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																ValidateDiagFunc: validateStandardEffect,
															},
															"media_search": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																ValidateDiagFunc: validateStandardEffect,
															},
														},
													},
												},
												"unknown_bots": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"generic": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																ValidateDiagFunc: validateStandardRecaptchaEffect,
															},
															"web_automation_tools": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																ValidateDiagFunc: validateStandardRecaptchaEffect,
															},
															"web_scrapers": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																ValidateDiagFunc: validateStandardRecaptchaEffect,
															},
															"api_libraries": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																ValidateDiagFunc: validateStandardRecaptchaEffect,
															},
															"http_libraries": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																ValidateDiagFunc: validateStandardRecaptchaEffect,
															},
															"request_anomalies": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"effect": {
																			Type:             schema.TypeString,
																			Optional:         true,
																			Default:          "disable",
																			ValidateDiagFunc: validateStandardRecaptchaEffect,
																		},
																		"threshold": {
																			Type:             schema.TypeString,
																			Optional:         true,
																			Default:          "lax",
																			ValidateDiagFunc: validateRequestAnomaliesEnforcementThreshold,
																		},
																	},
																},
															},
															"bot_impersonation": {
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "disable",
																ValidateDiagFunc: validateStandardRecaptchaEffect,
															},
														},
													},
												},
												"user_defined_bot": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
															"header_name": {
																Type:     schema.TypeString,
																Required: true,
															},
															"header_values": {
																Type:     schema.TypeList,
																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"source_subnets": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"effect": {
																Type:             schema.TypeString,
																Required:         true,
																ValidateDiagFunc: validateStandardRecaptchaEffect,
															},
														},
													},
												},
												"active_bot_detection": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"session_validation_failure_effect": {
																Type:             schema.TypeString,
																Optional:         true,
																ValidateDiagFunc: validateStandardEffect,
															},
															"javascript_based_detection": {
																Type:        schema.TypeList,
																Optional:    true,
																MaxItems:    1,
																Description: "Detect browser impersonation by injecting JavaScript to collect browser attributes and flag anomalies typical to various bot frameworks.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"enabled": {
																			Type:     schema.TypeBool,
																			Optional: true,
																			Default:  false,
																		},
																		"browser_impersonation": { // ApplicationsSpec.BotProtectionSpec.UnknownBotProtectionSpec.BrowserImpersonation
																			Type:             schema.TypeString,
																			Optional:         true,
																			Default:          "disable",
																			Description:      "Automated tools or services that impersonate common web browser software.",
																			ValidateDiagFunc: validateStandardRecaptchaEffect,
																		},
																		"injection_timeout_effect": {
																			Type:             schema.TypeString,
																			Optional:         true,
																			Default:          "disable",
																			Description:      "The web client failed to pass the bot-detection JavaScript injection check in reasonable time.",
																			ValidateDiagFunc: validateStandardEffect,
																		},
																	},
																},
															},
															"recaptcha": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"enabled": {
																			Type:     schema.TypeBool,
																			Optional: true,
																			Default:  false,
																		},
																		"site_key": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"secret_key": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"type": {
																			Type:             schema.TypeString,
																			Optional:         true,
																			Description:      "Type MUST match the challenge type selected during site registration.",
																			ValidateDiagFunc: validateRecaptchaType,
																		},
																		"every_new_session": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Use reCAPTCHA for every new session or according to policy. WAAS can respond with reCAPTCHA when it detects unknown bots.",
																		},
																		"success_expiration": {
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Default:     24,
																			Description: "Success expiration in hours.",
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
									"custom_rule": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"rule_id": { // WaasContainerApplicationSpec.CustomRules.Id
													Type:     schema.TypeInt,
													Required: true,
												},
												"effect": { // WaasContainerApplicationSpec.CustomRules.Effect
													Type:             schema.TypeString,
													Optional:         true,
													Default:          "alert",
													ValidateDiagFunc: validateStandardEffect,
												},
											},
										},
									},
									"advanced_settings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"session_cookies": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"enabled": { // WaasContainerApplicationSpec.SessionCookieEnabled
																Type:        schema.TypeBool,
																Optional:    true,
																Default:     false,
																Description: "Prisma session cookies are required for active bot detection and app DoS protection based on session. By enabling this feature, WAAS will set a Prisma session cookie and mandate its use for any communication with the application.",
															},
															"same_site": { // WaasContainerApplicationSpec.SessionCookieSameSite
																Type:             schema.TypeString,
																Optional:         true,
																Default:          "lax",
																ValidateDiagFunc: validateSameSite,
															},
															"secure": { // WaasContainerApplicationSpec.SessionCookieSecure
																Type:        schema.TypeBool,
																Optional:    true,
																Default:     false,
																Description: "Whether or not to set the `secure` attribute.",
															},
															"apply_ban_on_session": { // WaasContainerApplicationSpec.SessionCookieBan
																Type:        schema.TypeBool,
																Optional:    true,
																Default:     true,
																Description: "Whether or not to apply firewall, bot protection, and custom rules ban on session or client IP.",
															},
															"rate_limit_on_session": { // WaasContainerApplicationSpec.WaasContainerDosConfig.TrackSession
																Type:        schema.TypeBool,
																Optional:    true,
																Default:     true,
																Description: "Whether or not to limit rate and apply DoS protection ban on session or client IP.",
															},
														},
													},
												},
												"ban_duration": { // WaasContainerApplicationSpec.BanDurationMinutes
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     5,
													Description: "Ban duration in minutes.",
												},
												"http_body_inspection": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"enabled": { // WaasContainerApplicationSpec.Body.Skip
																Type:        schema.TypeBool,
																Optional:    true,
																Default:     true,
																Description: "Whether or not to inspect request bodies.",
															},
															"inspection_size": { // WaasContainerApplicationSpec.Body.InspectionSizeBytes
																Type:        schema.TypeInt,
																Optional:    true,
																Default:     131072, // 128 KB
																Description: "Number of bytes of request body to inspect.",
															},
														},
													},
												},
												"custom_waas_response": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"enabled": { // WaasContainerApplicationSpec.CustomBlockResponse.Enabled
																Type:     schema.TypeBool,
																Optional: true,
															},
															"status_code": { // WaasContainerApplicationSpec.CustomBlockResponse.Code
																Type:             schema.TypeInt,
																Required:         true,
																Description:      "HTTP response code.",
																ValidateDiagFunc: validateHttpResponseCode,
															},
															"body": { // WaasContainerApplicationSpec.CustomBlockResponse.Body
																Type:             schema.TypeString,
																Required:         true,
																Description:      "HTML to use as response body. Include Prisma Event IDs as part of customized responses by adding the following placeholder in user-provided HTML: '#eventID'.",
																ValidateDiagFunc: validateHtmlBody,
															},
														},
													},
												},
												"enable_event_id_header": { // WaasContainerApplicationSpec.DisableEventIdHeader
													Type:        schema.TypeBool,
													Optional:    true,
													Default:     true,
													Description: "Whether or not to enable event ID generation in response headers.",
												},
											},
										},
									},
								},
							},
						},
						// "remote_host_forwarding": { // TODO: only relevant for host WAAS
						// 	Type:     schema.TypeList,
						// 	Optional: true,
						// 	MaxItems: 1,
						// 	Elem: &schema.Resource{
						// 		Schema: map[string]*schema.Schema{
						// 			"enabled": {
						// 				Type:     schema.TypeBool,
						// 				Optional: true,
						// 			},
						// 			"target": {
						// 				Type:     schema.TypeString,
						// 				Optional: true,
						// 			},
						// 		},
						// 	},
						// },
						// "windows": { // TODO: only relevant for host WAAS
						// 	Type:        schema.TypeBool,
						// 	Optional:    true,
						// 	Default:     false,
						// 	Description: "Whether or not the host OS is Windows.",
						// },
					},
				},
			},
		},
	}
}

func validateHtmlBody(i interface{}, path cty.Path) diag.Diagnostics {
	v, ok := i.(string)
	if !ok {
		return diag.Errorf("expected type to be string")
	}

	if !strings.HasPrefix(v, "<html") || !strings.HasSuffix(v, "</html>") {
		return diag.Errorf("value must be wrapped in html tags (<html></html>)")
	}
	return nil
}

func validateSameSite(i interface{}, path cty.Path) diag.Diagnostics {
	v, ok := i.(string)
	if !ok {
		return diag.Errorf("expected type to be string")
	}
	validStrings := []string{
		"lax",
		"strict",
		"none",
	}
	if !stringInSlice(v, validStrings) {
		return diag.Errorf("%q must be one of %v", v, validStrings)
	}
	return nil
}

func validateHttpMethod(i interface{}, path cty.Path) diag.Diagnostics {
	v, ok := i.(string)
	if !ok {
		return diag.Errorf("expected type to be string")
	}
	validStrings := []string{
		"GET",
		"PUT",
		"POST",
		"DELETE",
		"OPTIONS",
		"HEAD",
		"PATCH",
	}
	if !stringInSlice(v, validStrings) {
		return diag.Errorf("%q must be one of %v", v, validStrings)
	}
	return nil
}

func validateParameterLocation(i interface{}, path cty.Path) diag.Diagnostics {
	v, ok := i.(string)
	if !ok {
		return diag.Errorf("expected type to be string")
	}
	validStrings := []string{
		"body",
		"cookie",
		"formData",
		"header",
		"json",
		"multipart",
		"path",
		"query",
		"xml",
	}
	if !stringInSlice(v, validStrings) {
		return diag.Errorf("%q must be one of %v", v, validStrings)
	}
	return nil
}

func validateParameterStyle(i interface{}, path cty.Path) diag.Diagnostics {
	v, ok := i.(string)
	if !ok {
		return diag.Errorf("expected type to be string")
	}
	validStrings := []string{
		"form",
		"label",
		"matrix",
		"pipeDelimited",
		"simple",
		"spaceDelimited",
		"tabDelimited",
	}
	if !stringInSlice(v, validStrings) {
		return diag.Errorf("%q must be one of %v", v, validStrings)
	}
	return nil
}

func validateParameterType(i interface{}, path cty.Path) diag.Diagnostics {
	v, ok := i.(string)
	if !ok {
		return diag.Errorf("expected type to be string")
	}
	validStrings := []string{
		"array",
		"boolean",
		"integer",
		"number",
		"object",
		"string",
	}
	if !stringInSlice(v, validStrings) {
		return diag.Errorf("%q must be one of %v", v, validStrings)
	}
	return nil
}

func validateRequestAnomaliesEnforcementThreshold(i interface{}, path cty.Path) diag.Diagnostics {
	v, ok := i.(string)
	if !ok {
		return diag.Errorf("expected type to be string")
	}
	validStrings := []string{
		"lax",
		"moderate",
		"strict",
	}
	if !stringInSlice(v, validStrings) {
		return diag.Errorf("%q must be one of %v", v, validStrings)
	}
	return nil
}

func validateRecaptchaType(i interface{}, path cty.Path) diag.Diagnostics {
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

func validateAlertPreventEffect(i interface{}, path cty.Path) diag.Diagnostics {
	v, ok := i.(string)
	if !ok {
		return diag.Errorf("expected type to be string")
	}
	validStrings := []string{
		"alert",
		"prevent",
	}
	if !stringInSlice(v, validStrings) {
		return diag.Errorf("%q must be one of %v", v, validStrings)
	}
	return nil
}

func validateAllowPreventEffect(i interface{}, path cty.Path) diag.Diagnostics {
	v, ok := i.(string)
	if !ok {
		return diag.Errorf("expected type to be string")
	}
	validStrings := []string{
		"allow",
		"prevent",
	}
	if !stringInSlice(v, validStrings) {
		return diag.Errorf("%q must be one of %v", v, validStrings)
	}
	return nil
}

func validateStandardEffect(i interface{}, path cty.Path) diag.Diagnostics {
	v, ok := i.(string)
	if !ok {
		return diag.Errorf("expected type to be string")
	}
	validStrings := []string{
		"disable",
		"alert",
		"prevent",
		"ban",
	}
	if !stringInSlice(v, validStrings) {
		return diag.Errorf("%q must be one of %v", v, validStrings)
	}
	return nil
}

func validateStandardRecaptchaEffect(i interface{}, path cty.Path) diag.Diagnostics {
	v, ok := i.(string)
	if !ok {
		return diag.Errorf("expected type to be string")
	}
	validStrings := []string{
		"disable",
		"alert",
		"prevent",
		"ban",
		"recaptcha",
	}
	if !stringInSlice(v, validStrings) {
		return diag.Errorf("%q must be one of %v", v, validStrings)
	}
	return nil
}

func validateTlsVersion(i interface{}, path cty.Path) diag.Diagnostics {
	v, ok := i.(string)
	if !ok {
		return diag.Errorf("expected type to be string")
	}
	validStrings := []string{
		"1.0",
		"1.1",
		"1.2",
		"1.3",
	}
	if !stringInSlice(v, validStrings) {
		return diag.Errorf("%q must be one of %v", v, validStrings)
	}
	return nil
}

func validatePort(i interface{}, path cty.Path) diag.Diagnostics {
	v, ok := i.(int)
	if !ok {
		return diag.Errorf("expected type to be int")
	}
	if v < 1 || v > 65535 {
		return diag.Errorf("expected value to be between 1 and 65535. got %d", v)
	}
	return nil
}

func validateHttpResponseCode(i interface{}, path cty.Path) diag.Diagnostics {
	v, ok := i.(int)
	if !ok {
		return diag.Errorf("expected type to be int")
	}
	if v < 100 || v > 599 {
		return diag.Errorf("expected value to be between 100 and 599. got %d", v)
	}
	return nil
}

func createPolicyFirewallAppContainer(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api.Client)
	parsedRules, err := convert.SchemaToWaasContainerRules(d)
	if err != nil {
		return diag.FromErr(err)
	}

	parsedPolicy := policy.WaasContainerPolicy{
		Id:    policyTypeWaasContainer,
		Rules: parsedRules,
	}
	if val, ok := d.GetOk("min_port"); ok {
		parsedPolicy.MinPort = val.(int)
	}
	if val, ok := d.GetOk("max_port"); ok {
		parsedPolicy.MaxPort = val.(int)
	}

	log.Printf("[DEBUG] parsedPolicy: %+v", parsedPolicy)

	if err := policy.UpdateWaasContainer(*client, parsedPolicy); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(parsedPolicy.Id)
	return readPolicyFirewallAppContainer(ctx, d, m)
}

func readPolicyFirewallAppContainer(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api.Client)
	retrievedPolicy, err := policy.GetWaasContainer(*client)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("min_port", retrievedPolicy.MinPort)
	d.Set("max_port", retrievedPolicy.MaxPort)
	if err := d.Set("rule", convert.WaasContainerRulesToSchema(retrievedPolicy.Rules)); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updatePolicyFirewallAppContainer(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api.Client)
	parsedRules, err := convert.SchemaToWaasContainerRules(d)
	if err != nil {
		return diag.FromErr(err)
	}

	parsedPolicy := policy.WaasContainerPolicy{
		Id:    policyTypeWaasContainer,
		Rules: parsedRules,
	}
	if val, ok := d.GetOk("min_port"); ok {
		parsedPolicy.MinPort = val.(int)
	}
	if val, ok := d.GetOk("max_port"); ok {
		parsedPolicy.MaxPort = val.(int)
	}

	if err := policy.UpdateWaasContainer(*client, parsedPolicy); err != nil {
		return diag.FromErr(err)
	}

	return readPolicyFirewallAppContainer(ctx, d, m)
}

func deletePolicyFirewallAppContainer(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api.Client)
	blankPolicy := policy.WaasContainerPolicy{
		Id:      policyTypeWaasContainer,
		MinPort: 30000,
		MaxPort: 31000,
	}

	if err := policy.UpdateWaasContainer(*client, blankPolicy); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
