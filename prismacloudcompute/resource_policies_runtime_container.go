package prismacloudcompute

import (
	"fmt"
	"time"

	pcc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePoliciesRuntimeContainer() *schema.Resource {
	return &schema.Resource{
		Create: createPolicyRuntimeContainer,
		Read:   readPolicyRuntimeContainer,
		Update: updatePolicyRuntimeContainer,
		Delete: deletePolicyRuntimeContainer,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the policy set.",
				Default:     policyTypeRuntimeContainer,
			},
			"learning_disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If set to 'true', automatic behavioural learning is enabled.",
				Default:     false,
			},
			"rule": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Rules in the policy.",
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advanced_protection": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Prisma Cloud advanced threat protection",
						},
						"cloud_metadata_enforcement": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Suspicious queries to cloud provider APIs",
						},
						"collections": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of collections used to scope the rule.",
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
									"_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Custom rule ID.",
									},
									"action": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The action to perform if the custom rule applies. Can be set to 'audit', 'incident'.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used for the custom rule. Can be set to 'block', 'prevent', 'alert', 'allow', 'ban', or 'disable'.",
									},
								},
							},
						},
						"disabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If set to 'true', the rule is currently disabled.",
						},
						"dns": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "The DNS runtime rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allowlist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Allow-listed domain names (e.g., *.gmail.com, *.s3.amazon.com).",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"denylist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Deny-listed domain names (e.g., www.bad-url.com, *.bad-url.com).",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used in the runtime rule. Can be set to 'block', 'prevent', 'alert', 'disable'.",
									},
								},
							},
						},
						"filesystem": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Represents restrictions or suppression for filesystem changes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allowlist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of allowed filesystem paths.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"backdoor_files": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', monitors files that can create or persist backdoors (SSH or admin account config files).",
									},
									"check_new_files": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', Detects changes to binaries and certificates.",
									},
									"denylist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of denied file system paths.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect that will be used in the runtime rule. Can be set to 'block', 'prevent', 'alert', or 'disable'.",
									},
									"skip_encrypted_binaries": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', the encrypted binaries check will be skipped.",
									},
									"suspicious_elf_headers": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', enables malware detection based on suspicious ELF headers.",
									},
								},
							},
						},
						"kubernetes_enforcement": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Detects containers that attempt to compromise the orchestrator.",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name of the rule.",
						},
						"network": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Represents the restrictions and suppression for networking.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"denied_outbound_ips": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Deny-listed IP addresses.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"denied_listening_port": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Deny-listed listening ports.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"deny": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If set to 'true', the connection is denied.",
												},
												"end": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "end",
												},
												"start": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "start",
												},
											},
										},
									},
									"denied_outbound_port": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Deny-listed outbound ports.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"deny": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If set to 'true', the connection is denied.",
												},
												"end": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "end",
												},
												"start": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "start",
												},
											},
										},
									},
									"detect_port_scan": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', port scanning detection is enabled.",
									},
									"effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect used in the runtime rule. Can be set to: 'block', 'prevent', 'alert', or 'disable'.",
									},
									"skip_modified_processes": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', Prisma Cloud can detect malicious networking activity from modified processes.",
									},
									"skip_raw_sockets": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', raw socket detection will be skipped.",
									},
									"allowed_outbound_ips": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Allow-listed IP addresses.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"allowed_listening_port": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Allow-listed listening ports.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"deny": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If set to 'true', the connection is denied.",
												},
												"end": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "end",
												},
												"start": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "start",
												},
											},
										},
									},
									"allowed_outbound_port": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Allow-listed outbound ports.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"deny": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If set to 'true', the connection is denied.",
												},
												"end": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "end",
												},
												"start": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "start",
												},
											},
										},
									},
								},
							},
						},
						"notes": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A free-form text description of the collection.",
						},
						"processes": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Represents restrictions or suppression for running processes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allowlist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Allow-list of processes.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"check_lateral_movement": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', enables detection of processes that can be used for lateral movement exploits.",
									},
									"check_parent_child": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', enables check for parent-child relationship when comparing spawned processes in the model.",
									},
									"check_suid_binaries": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', enables check for process elevanting privileges (SUID bit).",
									},
									"crypto_miners": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', detect crypto miners.",
									},
									"denylist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of processes to deny.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used in the runtime rule. Can be set to 'block', 'prevent', 'alert', 'disable'.",
									},
									"skip_modified": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Processes started from modified binaries",
									},
									"skip_reverse_shell": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', reverse shell detection is disabled.",
									},
								},
							},
						},
						"wildfire_analysis": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The effect that will be used in the runtime rule. Can be set to 'block', 'prevent', 'alert', or 'disable'.",
						},
					},
				},
			},
		},
	}
}

func parsePolicyRuntimeContainer(d *schema.ResourceData, policyId string) (*policy.Policy, error) {
	parsedPolicy, err := parsePolicy(d, policyId, "")
	if err != nil {
		return nil, fmt.Errorf("error parsing %s policy: %s", policyId, err)
	}
	return parsedPolicy, nil
}

func createPolicyRuntimeContainer(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedPolicy, err := parsePolicyRuntimeContainer(d, "")
	if err != nil {
		return fmt.Errorf("error creating %s policy: %s", policyTypeRuntimeContainer, err)
	}

	if err := policy.Update(*client, policy.RuntimeContainerEndpoint, *parsedPolicy); err != nil {
		return err
	}

	pol, err := policy.Get(*client, policy.RuntimeContainerEndpoint)
	if err != nil {
		return err
	}

	d.SetId(pol.PolicyId)
	return readPolicyRuntimeContainer(d, meta)
}

func readPolicyRuntimeContainer(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)

	retrievedPolicy, err := policy.Get(*client, policy.RuntimeContainerEndpoint)
	if err != nil {
		return err
	}

	d.Set("_id", policyTypeRuntimeContainer)
	d.Set("learning_disabled", retrievedPolicy.LearningDisabled)
	if err := d.Set("rule", flattenPolicyRuntimeContainerRules(retrievedPolicy.Rules)); err != nil {
		return fmt.Errorf("error setting rule for resource %s: %s", d.Id(), err)
	}

	return nil
}

func flattenPolicyRuntimeContainerRules(in []policy.Rule) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["advanced_protection"] = val.AdvancedProtection
		m["cloud_metadata_enforcement"] = val.CloudMetadataEnforcement
		m["collections"] = flattenCollections(val.Collections)
		m["custom_rule"] = flattenCustomRules(val.CustomRules)
		m["disabled"] = val.Disabled
		m["dns"] = flattenDns(val.Dns)
		m["filesystem"] = flattenFilesystem(val.Filesystem)
		m["kubernetes_enforcement"] = val.KubernetesEnforcement
		m["name"] = val.Name
		m["network"] = flattenNetwork(val.Network)
		m["notes"] = val.Notes
		m["processes"] = flattenProcesses(val.Processes)
		m["wildfire_analysis"] = val.WildFireAnalysis
		ans = append(ans, m)
	}
	return ans
}

func updatePolicyRuntimeContainer(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	id := d.Id()
	parsedPolicy, err := parsePolicyRuntimeContainer(d, id)
	if err != nil {
		return fmt.Errorf("error updating %s policy: %s", policyTypeRuntimeContainer, err)
	}

	if err := policy.Update(*client, policy.RuntimeContainerEndpoint, *parsedPolicy); err != nil {
		return err
	}

	return readPolicyRuntimeContainer(d, meta)
}

func deletePolicyRuntimeContainer(d *schema.ResourceData, meta interface{}) error {
	return nil
}
