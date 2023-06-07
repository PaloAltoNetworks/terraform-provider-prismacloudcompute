package provider

import (
	"context"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/policy"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePoliciesRuntimeContainer() *schema.Resource {
	return &schema.Resource{
		CreateContext: createPolicyRuntimeContainer,
		ReadContext:   readPolicyRuntimeContainer,
		UpdateContext: updatePolicyRuntimeContainer,
		DeleteContext: deletePolicyRuntimeContainer,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the policy.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"learning_disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether or not to disable automatic behavioral learning.",
				Default:     false,
			},
			"rule": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Rules that make up the policy.",
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advanced_protection_effect": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Whether or not to enable advanced protection.",
						},
						"cloud_metadata_enforcement_effect": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Whether or not to enable cloud metadata access monitoring.",
						},
						"previous_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"skip_exec_sessions": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "",
						},
						"collections": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Collections used to scope the rule.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"custom_rule": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of custom rules.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The action to perform if the custom rule applies. Can be set to 'audit' or 'incident'.",
									},
									"effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used. Can be set to 'block', 'prevent', 'alert', or 'allow'.",
									},
									"id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Custom rule number.",
									},
								},
							},
						},
						"disabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether or not to disable the rule.",
						},
						"dns": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "DNS configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"default_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"disabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "",
									},
									"domain_list": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"allowed": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Allowed domains. Wildcard prefixes are supported.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"denied": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Denied domains. Wildcard prefixes are supported.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"effect": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
								},
							},
						},
						"filesystem": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "File system configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allowed_list": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"backdoor_files_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"default_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"denied_list": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"effect": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"paths": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"disabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "",
									},
									"encrypted_binaries_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"new_files_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"suspicious_elf_headers_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"kubernetes_enforcement_effect": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Whether or not to detect attacks against the cluster.",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique name of the rule.",
						},
						"network": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Network configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allowed_ips": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"default_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"denied_ips": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"denied_ips_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"disabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "",
									},
									"listening_ports": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"effect": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"allowed": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "List of denied outbound ports.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"deny": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Whether or not to deny the connection.",
															},
															"end": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "End of the port range.",
															},
															"start": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Start of the port range.",
															},
														},
													},
												},
												"denied": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "List of denied outbound ports.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"deny": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Whether or not to deny the connection.",
															},
															"end": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "End of the port range.",
															},
															"start": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Start of the port range.",
															},
														},
													},
												},
											},
										},
									},
									"modified_proc_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"outbound_ports": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"effect": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"allowed": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "List of allowed outbound ports.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"deny": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Whether or not to deny the connection.",
															},
															"end": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "End of the port range.",
															},
															"start": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Start of the port range.",
															},
														},
													},
												},
												"denied": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "List of denied outbound ports.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"deny": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Whether or not to deny the connection.",
															},
															"end": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "End of the port range.",
															},
															"start": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Start of the port range.",
															},
														},
													},
												},
											},
										},
									},
									"port_scan_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"raw_sockets_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"notes": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Free-form text field.",
						},
						"processes": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Processes configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"check_parent_child": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to check for parent-child relationship when comparing spawned processes in the model.",
									},
									"allowed_list": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of allowed processes.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"modified_process_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"crypto_miners_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"lateral_movement_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"reverse_shell_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"suid_binaries_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"default_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"disabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not skip detection of reverse shells.",
									},
									"denied_list": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"effect": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"paths": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "",
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
						"wildfire_analysis": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The effect to be used when WildFire analysis is enabled. Can be set to 'block', 'alert', or 'disable'.",
						},
					},
				},
			},
		},
	}
}

func createPolicyRuntimeContainer(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)
	parsedRules, err := convert.SchemaToRuntimeContainerRules(d)
	if err != nil {
		return diag.Errorf("error creating %s policy: %s", policyTypeRuntimeContainer, err)
	}

	var learningDisabled bool
	if val, ok := d.GetOk("learning_disabled"); ok {
		d.Set("learning_disabled", val)
		learningDisabled = val.(bool)
	}

	parsedPolicy := policy.RuntimeContainerPolicy{
		LearningDisabled: learningDisabled,
		Rules:            parsedRules,
	}

	if err := policy.UpdateRuntimeContainer(*client, parsedPolicy); err != nil {
		return diag.Errorf("error creating %s policy: %s", policyTypeRuntimeContainer, err)
	}

	d.SetId(policyTypeRuntimeContainer)

	return readPolicyRuntimeContainer(ctx, d, meta)
}

func readPolicyRuntimeContainer(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)

	var diags diag.Diagnostics

	retrievedPolicy, err := policy.GetRuntimeContainer(*client)
	if err != nil {
		return diag.Errorf("error reading %s policy: %s", policyTypeRuntimeContainer, err)
	}

	d.Set("learning_disabled", retrievedPolicy.LearningDisabled)
	if err := d.Set("rule", convert.RuntimeContainerRulesToSchema(retrievedPolicy.Rules)); err != nil {
		return diag.Errorf("error reading %s policy: %s", policyTypeRuntimeContainer, err)
	}
	return diags
}

func updatePolicyRuntimeContainer(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.Client)

	var learningDisabled bool
	if val, ok := d.GetOk("learning_disabled"); ok {
		d.Set("learning_disabled", val)
		learningDisabled = val.(bool)
	}

	parsedRules, err := convert.SchemaToRuntimeContainerRules(d)

	if err != nil {
		return diag.Errorf("error updating %s policy: %s", policyTypeRuntimeContainer, err)
	}

	parsedPolicy := policy.RuntimeContainerPolicy{
		LearningDisabled: learningDisabled,
		Rules:            parsedRules,
	}

	if err := policy.UpdateRuntimeContainer(*client, parsedPolicy); err != nil {
		return diag.Errorf("error updating %s policy: %s", policyTypeRuntimeContainer, err)
	}

	return readPolicyRuntimeContainer(ctx, d, meta)
}

func deletePolicyRuntimeContainer(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// TODO: reset to default policy
	var diags diag.Diagnostics
	return diags
}
