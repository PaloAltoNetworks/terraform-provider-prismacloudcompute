package prismacloudcompute

import (
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/prismacloudcompute/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy"
)

func resourcePoliciesRuntimeContainer() *schema.Resource {
	return &schema.Resource{
		Create: createPolicyRuntimeContainer,
		Read:   readPolicyRuntimeContainer,
		Update: updatePolicyRuntimeContainer,
		Delete: deletePolicyRuntimeContainer,

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
						"advanced_protection": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether or not to enable advanced protection.",
						},
						"cloud_metadata_enforcement": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether or not to enable cloud metadata access monitoring.",
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
									"deny_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used. Can be set to 'block', 'prevent', 'alert', or 'disable'.",
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
									"allowed": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of allowed file system paths.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"backdoor_files": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to monitor files that can create or persist backdoors (SSH or admin account config files).",
									},
									"check_new_files": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to detect changes to binaries and certificates.",
									},
									"denied": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of denied file system paths.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"deny_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used. Can be set to 'block', 'prevent', 'alert', or 'disable'.",
									},
									"skip_encrypted_binaries": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to skip encrypted binaries.",
									},
									"suspicious_elf_headers": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to detect suspicious ELF headers.",
									},
								},
							},
						},
						"kubernetes_enforcement": {
							Type:        schema.TypeBool,
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
									"allowed_listening_port": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of allowed listening ports.",
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
									"allowed_outbound_ips": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of allowed outbound IP addresses.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"allowed_outbound_port": {
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
									"denied_listening_port": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of denied listening ports.",
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
									"denied_outbound_ips": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of denied outbound IP addresses.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"denied_outbound_port": {
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
									"deny_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used. Can be set to 'block', 'alert', or 'disable'.",
									},
									"detect_port_scan": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to detect port scans.",
									},
									"skip_modified_processes": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to skip network monitoring for modified processes.",
									},
									"skip_raw_sockets": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to skip raw socket detection.",
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
									"allowed": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of allowed processes.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"check_crypto_miners": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to detect crypto miners.",
									},
									"check_lateral_movement": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to detect processes that can be used for lateral movement exploits.",
									},
									"check_parent_child": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to check for parent-child relationship when comparing spawned processes in the model.",
									},
									"check_suid_binaries": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to check for process-elevating privileges (SUID bit).",
									},
									"denied": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of denied processes.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"deny_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used. Can be set to 'block', 'prevent', 'alert', or 'disable'.",
									},
									"skip_modified": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to skip detection of processes started from modified binaries",
									},
									"skip_reverse_shell": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not skip detection of reverse shells.",
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

func createPolicyRuntimeContainer(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedRules, err := convert.SchemaToRuntimeContainerRules(d)
	if err != nil {
		return fmt.Errorf("error creating %s policy: %s", policyTypeRuntimeContainer, err)
	}

	parsedPolicy := policy.RuntimeContainerPolicy{
		Rules: parsedRules,
	}

	if err := policy.UpdateRuntimeContainer(*client, parsedPolicy); err != nil {
		return fmt.Errorf("error creating %s policy: %s", policyTypeRuntimeContainer, err)
	}

	d.SetId(policyTypeRuntimeContainer)
	return readPolicyRuntimeContainer(d, meta)
}

func readPolicyRuntimeContainer(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	retrievedPolicy, err := policy.GetRuntimeContainer(*client)
	if err != nil {
		return fmt.Errorf("error reading %s policy: %s", policyTypeRuntimeContainer, err)
	}

	d.Set("learning_disabled", retrievedPolicy.LearningDisabled)
	if err := d.Set("rule", convert.RuntimeContainerRulesToSchema(retrievedPolicy.Rules)); err != nil {
		return fmt.Errorf("error reading %s policy: %s", policyTypeRuntimeContainer, err)
	}
	return nil
}

func updatePolicyRuntimeContainer(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedRules, err := convert.SchemaToRuntimeContainerRules(d)
	if err != nil {
		return fmt.Errorf("error updating %s policy: %s", policyTypeRuntimeContainer, err)
	}

	parsedPolicy := policy.RuntimeContainerPolicy{
		Rules: parsedRules,
	}

	if err := policy.UpdateRuntimeContainer(*client, parsedPolicy); err != nil {
		return fmt.Errorf("error updating %s policy: %s", policyTypeRuntimeContainer, err)
	}

	return readPolicyRuntimeContainer(d, meta)
}

func deletePolicyRuntimeContainer(d *schema.ResourceData, meta interface{}) error {
	// TODO: reset to default policy
	return nil
}
