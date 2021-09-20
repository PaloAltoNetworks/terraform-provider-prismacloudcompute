package prismacloudcompute

import (
	"fmt"
	"time"

	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
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
									"id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Custom rule ID.",
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
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The DNS runtime rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allowed": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Allowed domains (e.g. gmail.com, *.s3.amazon.com).",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"denied": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Denied domains (e.g. www.bad-url.com, *.bad-url.com).",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"deny_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used in the runtime rule. Can be set to 'block', 'prevent', 'alert', 'disable'.",
									},
								},
							},
						},
						"filesystem": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Represents restrictions or suppression for filesystem changes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allowed": {
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
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Represents the restrictions and suppression for networking.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
									"allowed_outbound_ips": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Allow-listed IP addresses.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
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
									"denied_outbound_ips": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Deny-listed IP addresses.",
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
									"deny_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect used in the runtime rule. Can be set to: 'block', 'prevent', 'alert', or 'disable'.",
									},
									"detect_port_scan": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', port scanning detection is enabled.",
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
								},
							},
						},
						"notes": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A free-form text description of the collection.",
						},
						"processes": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Represents restrictions or suppression for running processes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allowed": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Allow-list of processes.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"check_crypto_miners": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', detect crypto miners.",
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
									"denied": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of processes to deny.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"deny_effect": {
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

func createPolicyRuntimeContainer(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedPolicy, err := parsePolicyRuntimeContainer(d)
	if err != nil {
		return fmt.Errorf("error creating %s policy: %s", policyTypeRuntimeContainer, err)
	}

	if err := policy.UpdateRuntimeContainer(*client, *parsedPolicy); err != nil {
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
	if err := d.Set("rule", flattenPolicyRuntimeContainerRules(retrievedPolicy.Rules)); err != nil {
		return fmt.Errorf("error reading %s policy: %s", policyTypeRuntimeContainer, err)
	}
	return nil
}

func updatePolicyRuntimeContainer(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedPolicy, err := parsePolicyRuntimeContainer(d)
	if err != nil {
		return fmt.Errorf("error updating %s policy: %s", policyTypeRuntimeContainer, err)
	}

	if err := policy.UpdateRuntimeContainer(*client, *parsedPolicy); err != nil {
		return fmt.Errorf("error updating %s policy: %s", policyTypeRuntimeContainer, err)
	}

	return readPolicyRuntimeContainer(d, meta)
}

func deletePolicyRuntimeContainer(d *schema.ResourceData, meta interface{}) error {
	// TODO: reset to default policy
	return nil
}

func parsePolicyRuntimeContainer(d *schema.ResourceData) (*policy.RuntimeContainerPolicy, error) {
	parsedPolicy := policy.RuntimeContainerPolicy{
		Rules: make([]policy.RuntimeContainerRule, 0),
	}
	if rules, ok := d.GetOk("rule"); ok {
		rulesList := rules.([]interface{})
		parsedRules := make([]policy.RuntimeContainerRule, 0, len(rulesList))
		for _, val := range rulesList {
			rule := val.(map[string]interface{})
			parsedRule := policy.RuntimeContainerRule{}

			parsedRule.AdvancedProtection = rule["advanced_protection"].(bool)
			parsedRule.CloudMetadataEnforcement = rule["cloud_metadata_enforcement"].(bool)
			parsedRule.Collections = parseCollections(rule["collections"].([]interface{}))
			parsedRule.CustomRules = parseRuntimeContainerCustomRules(rule["custom_rule"].([]interface{}))
			parsedRule.Disabled = rule["disabled"].(bool)
			parsedRule.Dns = parseRuntimeContainerDns(rule["dns"].([]interface{}))
			parsedRule.Filesystem = parseRuntimeContainerFilesystem(rule["filesystem"].([]interface{}))
			parsedRule.KubernetesEnforcement = rule["kubernetes_enforcement"].(bool)
			parsedRule.Name = rule["name"].(string)
			parsedRule.Network = parseRuntimeContainerNetwork(rule["network"].([]interface{}))
			parsedRule.Notes = rule["notes"].(string)
			parsedRule.Processes = parseRuntimeContainerProcesses(rule["processes"].([]interface{}))
			parsedRule.WildFireAnalysis = rule["wildfire_analysis"].(string)

			parsedRules = append(parsedRules, parsedRule)
		}
		parsedPolicy.Rules = parsedRules
	}
	return &parsedPolicy, nil
}

func parseRuntimeContainerCustomRules(in []interface{}) []policy.RuntimeContainerCustomRule {
	parsedCustomRules := make([]policy.RuntimeContainerCustomRule, 0, len(in))
	for _, val := range in {
		presentCustomRule := val.(map[string]interface{})
		parsedCustomRule := policy.RuntimeContainerCustomRule{}
		if presentCustomRule["action"] != nil {
			parsedCustomRule.Action = presentCustomRule["action"].(string)
		}
		if presentCustomRule["effect"] != nil {
			parsedCustomRule.Effect = presentCustomRule["effect"].(string)
		}
		if presentCustomRule["id"] != nil {
			parsedCustomRule.Id = presentCustomRule["id"].(int)
		}
		parsedCustomRules = append(parsedCustomRules, parsedCustomRule)
	}
	return parsedCustomRules
}

func parseRuntimeContainerDns(in []interface{}) policy.RuntimeContainerDns {
	parsedDns := policy.RuntimeContainerDns{}
	if in[0] == nil {
		return parsedDns
	}
	presentDns := in[0].(map[string]interface{})
	if presentDns["allowed"] != nil {
		parsedDns.Allowed = parseStringArray(presentDns["allowed"].([]interface{}))
	}
	if presentDns["denied"] != nil {
		parsedDns.Denied = parseStringArray(presentDns["denied"].([]interface{}))
	}
	if presentDns["deny_effect"] != nil {
		parsedDns.DenyEffect = presentDns["deny_effect"].(string)
	}
	return parsedDns
}

func parseRuntimeContainerFilesystem(in []interface{}) policy.RuntimeContainerFilesystem {
	parsedFilesystem := policy.RuntimeContainerFilesystem{}
	if in[0] == nil {
		return parsedFilesystem
	}
	presentFilesystem := in[0].(map[string]interface{})
	if presentFilesystem["allowed"] != nil {
		parsedFilesystem.Allowed = parseStringArray(presentFilesystem["allowed"].([]interface{}))
	}
	if presentFilesystem["backdoor_files"] != nil {
		parsedFilesystem.BackdoorFiles = presentFilesystem["backdoor_files"].(bool)
	}
	if presentFilesystem["check_new_files"] != nil {
		parsedFilesystem.CheckNewFiles = presentFilesystem["check_new_files"].(bool)
	}
	if presentFilesystem["denied"] != nil {
		parsedFilesystem.Denied = parseStringArray(presentFilesystem["denied"].([]interface{}))
	}
	if presentFilesystem["deny_effect"] != nil {
		parsedFilesystem.DenyEffect = presentFilesystem["deny_effect"].(string)
	}
	if presentFilesystem["skip_encrypted_binaries"] != nil {
		parsedFilesystem.SkipEncryptedBinaries = presentFilesystem["skip_encrypted_binaries"].(bool)
	}
	if presentFilesystem["suspicious_elf_headers"] != nil {
		parsedFilesystem.SuspiciousElfHeaders = presentFilesystem["suspicious_elf_headers"].(bool)
	}
	return parsedFilesystem
}

func parseRuntimeContainerNetwork(in []interface{}) policy.RuntimeContainerNetwork {
	parsedNetwork := policy.RuntimeContainerNetwork{}
	if in[0] == nil {
		return parsedNetwork
	}
	presentNetwork := in[0].(map[string]interface{})
	if presentNetwork["allowed_listening_port"] != nil {
		parsedNetwork.AllowedListeningPorts = parseRuntimeContainerPorts(presentNetwork["allowed_listening_port"].([]interface{}))
	}
	if presentNetwork["allowed_outbound_ips"] != nil {
		parsedNetwork.AllowedOutboundIps = parseStringArray(presentNetwork["allowed_outbound_ips"].([]interface{}))
	}
	if presentNetwork["allowed_outbound_port"] != nil {
		parsedNetwork.AllowedOutboundPorts = parseRuntimeContainerPorts(presentNetwork["allowed_outbound_port"].([]interface{}))

	}
	if presentNetwork["denied_listening_port"] != nil {
		parsedNetwork.DeniedListeningPorts = parseRuntimeContainerPorts(presentNetwork["denied_listening_port"].([]interface{}))

	}
	if presentNetwork["denied_outbound_ips"] != nil {
		parsedNetwork.DeniedOutboundIps = parseStringArray(presentNetwork["denied_outbound_ips"].([]interface{}))
	}
	if presentNetwork["denied_outbound_port"] != nil {
		parsedNetwork.DeniedOutboundPorts = parseRuntimeContainerPorts(presentNetwork["denied_outbound_port"].([]interface{}))
	}
	if presentNetwork["deny_effect"] != nil {
		parsedNetwork.DenyEffect = presentNetwork["deny_effect"].(string)
	}
	if presentNetwork["detect_port_scan"] != nil {
		parsedNetwork.DetectPortScan = presentNetwork["detect_port_scan"].(bool)
	}
	if presentNetwork["skip_modified_processes"] != nil {
		parsedNetwork.SkipModifiedProcesses = presentNetwork["skip_modified_processes"].(bool)
	}
	if presentNetwork["skip_raw_sockets"] != nil {
		parsedNetwork.SkipRawSockets = presentNetwork["skip_raw_sockets"].(bool)
	}
	return parsedNetwork
}

func parseRuntimeContainerProcesses(in []interface{}) policy.RuntimeContainerProcesses {
	parsedProcesses := policy.RuntimeContainerProcesses{}
	if in[0] == nil {
		return parsedProcesses
	}
	presentProcesses := in[0].(map[string]interface{})
	if presentProcesses["allowed"] != nil {
		parsedProcesses.Allowed = parseStringArray(presentProcesses["allowed"].([]interface{}))
	}
	if presentProcesses["check_crypto_miners"] != nil {
		parsedProcesses.CheckCryptoMiners = presentProcesses["check_crypto_miners"].(bool)
	}
	if presentProcesses["check_lateral_movement"] != nil {
		parsedProcesses.CheckLateralMovement = presentProcesses["check_lateral_movement"].(bool)
	}
	if presentProcesses["check_parent_child"] != nil {
		parsedProcesses.CheckParentChild = presentProcesses["check_parent_child"].(bool)
	}
	if presentProcesses["check_suid_binaries"] != nil {
		parsedProcesses.CheckSuidBinaries = presentProcesses["check_suid_binaries"].(bool)
	}
	if presentProcesses["denied"] != nil {
		parsedProcesses.Denied = parseStringArray(presentProcesses["denied"].([]interface{}))
	}
	if presentProcesses["deny_effect"] != nil {
		parsedProcesses.DenyEffect = presentProcesses["deny_effect"].(string)
	}
	if presentProcesses["skip_modified"] != nil {
		parsedProcesses.SkipModified = presentProcesses["skip_modified"].(bool)
	}
	if presentProcesses["skip_reverse_shell"] != nil {
		parsedProcesses.SkipReverseShell = presentProcesses["skip_reverse_shell"].(bool)
	}
	return parsedProcesses
}

func parseRuntimeContainerPorts(in []interface{}) []policy.RuntimeContainerPort {
	parsedPorts := make([]policy.RuntimeContainerPort, 0, len(in))
	for _, val := range in {
		presentPort := val.(map[string]interface{})
		parsedPort := policy.RuntimeContainerPort{}
		if presentPort["deny"] != nil {
			parsedPort.Deny = presentPort["deny"].(bool)
		}
		if presentPort["end"] != nil {
			parsedPort.End = presentPort["end"].(int)
		}
		if presentPort["start"] != nil {
			parsedPort.Start = presentPort["start"].(int)
		}
		parsedPorts = append(parsedPorts, parsedPort)
	}
	return parsedPorts
}

func flattenPolicyRuntimeContainerRules(in []policy.RuntimeContainerRule) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["advanced_protection"] = val.AdvancedProtection
		m["cloud_metadata_enforcement"] = val.CloudMetadataEnforcement
		m["collections"] = flattenCollections(val.Collections)
		m["custom_rule"] = flattenRuntimeContainerCustomRules(val.CustomRules)
		m["disabled"] = val.Disabled
		m["dns"] = flattenRuntimeContainerDns(val.Dns)
		m["filesystem"] = flattenRuntimeContainerFilesystem(val.Filesystem)
		m["kubernetes_enforcement"] = val.KubernetesEnforcement
		m["name"] = val.Name
		m["network"] = flattenRuntimeContainerNetwork(val.Network)
		m["notes"] = val.Notes
		m["processes"] = flattenRuntimeContainerProcesses(val.Processes)
		m["wildfire_analysis"] = val.WildFireAnalysis
		ans = append(ans, m)
	}
	return ans
}
