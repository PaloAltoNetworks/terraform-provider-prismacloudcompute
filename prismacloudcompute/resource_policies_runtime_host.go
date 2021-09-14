package prismacloudcompute

import (
	"fmt"
	"time"

	pcc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePoliciesRuntimeHost() *schema.Resource {
	return &schema.Resource{
		Create: createPolicyRuntimeHost,
		Read:   readPolicyRuntimeHost,
		Update: updatePolicyRuntimeHost,
		Delete: deletePolicyRuntimeHost,

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
				Type:     schema.TypeString,
				Optional: true,
				Default:  policyTypeRuntimeHost,
			},
			"rule": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Rules in the policy.",
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"antimalware": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Restrictions/suppression for suspected anti-malware.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allowed_processes": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "A list of paths for files and processes to skip during anti-malware checks.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"crypto_miner": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"custom_feed": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"denied_processes": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "A rule containing paths of files and processes to alert/prevent and the required effect.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"effect": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Effect that will be used in the runtime rule.",
												},
												"paths": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Paths to alert/prevent when an event with one of the paths is triggered.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"detect_compiler_generated_binary": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Represents what happens when a compiler service writes a binary.",
									},
									"encrypted_binaries": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"execution_flow_hijack": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"intelligence_feed": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"reverse_shell": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"service_unknown_origin_binary": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"skip_ssh_tracking": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"suspicious_elf_headers": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"temp_filesystem_processes": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"user_unknown_origin_binary": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"webshell": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"wildfire_analysis": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
								},
							},
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
							Description: "List of custom runtime rules.",
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
										Description: "The action to perform if the custom rule applies. Can be set to 'audit' or 'incident'.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"effect": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The effect that will be used for a custom rule. Can be set to 'block', 'prevent', 'alert', 'allow', 'ban', or 'disable'.",
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
							Description: "If set to 'true', the rule is currently disabled.",
						},
						"dns": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "The DNS runtime rule",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allow": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of allowed domain names (e.g., *.gmail.com, *.s3.amazon.com).",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"deny": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Deny-list of domain names (e.g., www.bad-url.com, *.bad-url.com).",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"deny_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"intelligence_feed": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
								},
							},
						},
						"file_integrity_rule": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "file integrity monitoring rules..",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allowed_processes": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The processes to ignore Filesystem events caused by these processes DO NOT generate file integrity events.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"exclusions": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Filenames that should be ignored while generating audits These filenames may contain a wildcad regex pattern, e.g. foo*.log, *.cache.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"metadata": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates that metadata changes should be monitored (e.g. chmod, chown).",
									},
									"path": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Path to monitor.",
									},
									"read": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates that eads operations should be monitored.",
									},
									"recursive": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates that monitoring should be recursive.",
									},
									"write": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates that write operations should be monitored.",
									},
								},
							},
						},
						"forensic": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Indicates how to perform host forensic.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"activities_disabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates if the host activity collection is enabled/disabled.",
									},
									"docker_enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates whether docker commands are collected.",
									},
									"readonly_docker_enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates whether docker readonly commands are collected.",
									},
									"service_activities_enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates whether activities from services are collected.",
									},
									"sshd_enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates whether ssh commands are collected.",
									},
									"sudo_enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates whether sudo commands are collected.",
									},
								},
							},
						},
						"log_inspection_rule": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of log inspection rules.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"path": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "the log path.",
									},
									"regex": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Regular expressions associated with the rule if it is a custom one.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the rule.",
						},
						"network": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Represents the restrictions or suppression for networking.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allowed_outbound_ips": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Allow-listed IP addresses.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"custom_feed": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"denied_listening_port": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Deny-list of listening ports.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"deny": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If set to 'true' the connection is denied.",
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
									"denied_outbound_ips": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Deny-list of IP addresses.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
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
													Description: "If set to 'true' the connection is denied.",
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
									"deny_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"intelligence_feed": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
								},
							},
						},
						"notes": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A free-form text description of the collection.",
						},
					},
				},
			},
		},
	}
}

func parsePolicyRuntimeHost(d *schema.ResourceData, policyId string) (*policy.Policy, error) {
	parsedPolicy, err := parsePolicy(d, policyId, "")
	if err != nil {
		return nil, fmt.Errorf("error parsing %s policy: %s", policyId, err)
	}
	return parsedPolicy, nil
}

func flattenPolicyRuntimeHostRules(in []policy.Rule) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["antimalware"] = flattenAntiMalware(val.AntiMalware)
		m["collections"] = flattenCollections(val.Collections)
		m["custom_rule"] = flattenCustomRules(val.CustomRules)
		m["disabled"] = val.Disabled
		m["dns"] = flattenDns(val.Dns)
		m["file_integrity_rule"] = flattenFileIntegrityRules(val.FileIntegrityRules)
		m["forensic"] = flattenForensic(val.Forensic)
		m["log_inspection_rule"] = flattenLogInspectionRules(val.LogInspectionRules)
		m["name"] = val.Name
		m["network"] = flattenNetwork(val.Network)
		m["notes"] = val.Notes
		ans = append(ans, m)
	}
	return ans
}

func createPolicyRuntimeHost(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedPolicy, err := parsePolicyRuntimeHost(d, "")
	if err != nil {
		return fmt.Errorf("error creating %s policy: %s", policyTypeRuntimeHost, err)
	}

	if err := policy.Update(*client, policy.RuntimeHostEndpoint, *parsedPolicy); err != nil {
		return err
	}

	pol, err := policy.Get(*client, policy.RuntimeHostEndpoint)
	if err != nil {
		return err
	}

	d.SetId(pol.PolicyId)
	return readPolicyRuntimeHost(d, meta)
}

func readPolicyRuntimeHost(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)

	retrievedPolicy, err := policy.Get(*client, policy.RuntimeHostEndpoint)
	if err != nil {
		return err
	}

	d.Set("_id", policyTypeRuntimeHost)
	if err := d.Set("rule", flattenPolicyRuntimeHostRules(retrievedPolicy.Rules)); err != nil {
		return fmt.Errorf("error setting rule for resource %s: %s", d.Id(), err)
	}

	return nil
}

func updatePolicyRuntimeHost(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	id := d.Id()
	parsedPolicy, err := parsePolicyRuntimeHost(d, id)
	if err != nil {
		return fmt.Errorf("error updating %s policy: %s", policyTypeRuntimeHost, err)
	}

	if err := policy.Update(*client, policy.RuntimeHostEndpoint, *parsedPolicy); err != nil {
		return err
	}

	return readPolicyRuntimeHost(d, meta)
}

func deletePolicyRuntimeHost(d *schema.ResourceData, meta interface{}) error {
	return nil
}
