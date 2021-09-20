package prismacloudcompute

import (
	"fmt"
	"time"

	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"rule": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Rules in the policy.",
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"antimalware": {
							Type:        schema.TypeList,
							MaxItems:    1,
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
									"crypto_miners": {
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
										Type:        schema.TypeList,
										MaxItems:    1,
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
										Type:        schema.TypeBool,
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
									"action": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The action to perform if the custom rule applies. Can be set to 'audit' or 'incident'.",
									},
									"effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect that will be used for a custom rule. Can be set to 'block', 'prevent', 'alert', 'allow', 'ban', or 'disable'.",
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
							Description: "The DNS runtime rule",
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
									"excluded_files": {
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
							Type:        schema.TypeList,
							MaxItems:    1,
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
							Type:        schema.TypeList,
							MaxItems:    1,
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

func createPolicyRuntimeHost(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedPolicy, err := parsePolicyRuntimeHost(d)
	if err != nil {
		return fmt.Errorf("error creating %s policy: %s", policyTypeRuntimeHost, err)
	}

	if err := policy.UpdateRuntimeHost(*client, *parsedPolicy); err != nil {
		return fmt.Errorf("error creating %s policy: %s", policyTypeRuntimeHost, err)
	}

	d.SetId(policyTypeRuntimeHost)
	return readPolicyRuntimeHost(d, meta)
}

func readPolicyRuntimeHost(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	retrievedPolicy, err := policy.GetRuntimeHost(*client)
	if err != nil {
		return fmt.Errorf("error reading %s policy: %s", policyTypeRuntimeHost, err)
	}

	if err := d.Set("rule", flattenPolicyRuntimeHostRules(retrievedPolicy.Rules)); err != nil {
		return fmt.Errorf("error reading %s policy: %s", policyTypeRuntimeHost, err)
	}

	return nil
}

func updatePolicyRuntimeHost(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)
	parsedPolicy, err := parsePolicyRuntimeHost(d)
	if err != nil {
		return fmt.Errorf("error updating %s policy: %s", policyTypeRuntimeHost, err)
	}

	if err := policy.UpdateRuntimeHost(*client, *parsedPolicy); err != nil {
		return fmt.Errorf("error updating %s policy: %s", policyTypeRuntimeHost, err)
	}

	return readPolicyRuntimeHost(d, meta)
}

func deletePolicyRuntimeHost(d *schema.ResourceData, meta interface{}) error {
	// TODO: reset to default policy
	return nil
}

func parsePolicyRuntimeHost(d *schema.ResourceData) (*policy.RuntimeHostPolicy, error) {
	parsedPolicy := policy.RuntimeHostPolicy{
		Rules: make([]policy.RuntimeHostRule, 0),
	}
	if rules, ok := d.GetOk("rule"); ok {
		rulesList := rules.([]interface{})
		parsedRules := make([]policy.RuntimeHostRule, 0, len(rulesList))
		for _, val := range rulesList {
			rule := val.(map[string]interface{})
			parsedRule := policy.RuntimeHostRule{}

			parsedRule.AntiMalware = parseRuntimeHostAntiMalware(rule["antimalware"].([]interface{}))
			parsedRule.Collections = parseCollections(rule["collections"].([]interface{}))
			parsedRule.CustomRules = parseRuntimeHostCustomRules(rule["custom_rule"].([]interface{}))
			parsedRule.Disabled = rule["disabled"].(bool)
			parsedRule.Dns = parseRuntimeHostDns(rule["dns"].([]interface{}))
			parsedRule.FileIntegrityRules = parseRuntimeHostFileIntegrityRules(rule["file_integrity_rule"].([]interface{}))
			parsedRule.Forensic = parseRuntimeHostForensic(rule["forensic"].([]interface{}))
			parsedRule.LogInspectionRules = parseRuntimeHostLogInspectionRules(rule["log_inspection_rule"].([]interface{}))
			parsedRule.Name = rule["name"].(string)
			parsedRule.Network = parseRuntimeHostNetwork(rule["network"].([]interface{}))
			parsedRule.Notes = rule["notes"].(string)

			parsedRules = append(parsedRules, parsedRule)
		}
		parsedPolicy.Rules = parsedRules
	}
	return &parsedPolicy, nil
}

func parseRuntimeHostAntiMalware(in []interface{}) policy.RuntimeHostAntiMalware {
	parsedAntiMalware := policy.RuntimeHostAntiMalware{}
	if in[0] == nil {
		return parsedAntiMalware
	}
	presentAntiMalware := in[0].(map[string]interface{})
	if presentAntiMalware["allowed_processes"] != nil {
		parsedAntiMalware.AllowedProcesses = parseStringArray(presentAntiMalware["allowed_processes"].([]interface{}))
	}
	if presentAntiMalware["crypto_miners"] != nil {
		parsedAntiMalware.CryptoMiner = presentAntiMalware["crypto_miners"].(string)
	}
	if presentAntiMalware["custom_feed"] != nil {
		parsedAntiMalware.CustomFeed = presentAntiMalware["custom_feed"].(string)
	}
	if presentAntiMalware["denied_processes"] != nil {
		parsedAntiMalware.DeniedProcesses = parseRuntimeHostDeniedProcesses(presentAntiMalware["denied_processes"].([]interface{}))
	}
	if presentAntiMalware["detect_compiler_generated_binary"] != nil {
		parsedAntiMalware.DetectCompilerGeneratedBinary = presentAntiMalware["detect_compiler_generated_binary"].(bool)
	}
	if presentAntiMalware["encrypted_binaries"] != nil {
		parsedAntiMalware.EncryptedBinaries = presentAntiMalware["encrypted_binaries"].(string)
	}
	if presentAntiMalware["execution_flow_hijack"] != nil {
		parsedAntiMalware.ExecutionFlowHijack = presentAntiMalware["execution_flow_hijack"].(string)
	}
	if presentAntiMalware["intelligence_feed"] != nil {
		parsedAntiMalware.IntelligenceFeed = presentAntiMalware["intelligence_feed"].(string)
	}
	if presentAntiMalware["reverse_shell"] != nil {
		parsedAntiMalware.ReverseShell = presentAntiMalware["reverse_shell"].(string)
	}
	if presentAntiMalware["service_unknown_origin_binary"] != nil {
		parsedAntiMalware.ServiceUnknownOriginBinary = presentAntiMalware["service_unknown_origin_binary"].(string)
	}
	if presentAntiMalware["skip_ssh_tracking"] != nil {
		parsedAntiMalware.SkipSshTracking = presentAntiMalware["skip_ssh_tracking"].(bool)
	}
	if presentAntiMalware["suspicious_elf_headers"] != nil {
		parsedAntiMalware.SuspiciousElfHeaders = presentAntiMalware["suspicious_elf_headers"].(string)
	}
	if presentAntiMalware["temp_filesystem_processes"] != nil {
		parsedAntiMalware.TempFsProcesses = presentAntiMalware["temp_filesystem_processes"].(string)
	}
	if presentAntiMalware["user_unknown_origin_binary"] != nil {
		parsedAntiMalware.UserUnknownOriginBinary = presentAntiMalware["user_unknown_origin_binary"].(string)
	}
	if presentAntiMalware["webshell"] != nil {
		parsedAntiMalware.WebShell = presentAntiMalware["webshell"].(string)
	}
	if presentAntiMalware["wildfire_analysis"] != nil {
		parsedAntiMalware.WildFireAnalysis = presentAntiMalware["wildfire_analysis"].(string)
	}
	return parsedAntiMalware
}

func parseRuntimeHostCustomRules(in []interface{}) []policy.RuntimeHostCustomRule {
	parsedCustomRules := make([]policy.RuntimeHostCustomRule, 0, len(in))
	for _, val := range in {
		presentCustomRule := val.(map[string]interface{})
		parsedCustomRule := policy.RuntimeHostCustomRule{}
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

func parseRuntimeHostDns(in []interface{}) policy.RuntimeHostDns {
	parsedDns := policy.RuntimeHostDns{}
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
	if presentDns["intelligence_feed"] != nil {
		parsedDns.IntelligenceFeed = presentDns["intelligence_feed"].(string)
	}
	return parsedDns
}

func parseRuntimeHostFileIntegrityRules(in []interface{}) []policy.RuntimeHostFileIntegrityRule {
	parsedFileIntegrityRules := make([]policy.RuntimeHostFileIntegrityRule, 0, len(in))
	for _, val := range in {
		presentFileIntegrityRule := val.(map[string]interface{})
		parsedFileIntegrityRule := policy.RuntimeHostFileIntegrityRule{}
		if presentFileIntegrityRule["allowed_processes"] != nil {
			parsedFileIntegrityRule.AllowedProcesses = parseStringArray(presentFileIntegrityRule["allowed_processes"].([]interface{}))
		}
		if presentFileIntegrityRule["excluded_files"] != nil {
			parsedFileIntegrityRule.ExcludedFiles = parseStringArray(presentFileIntegrityRule["excluded_files"].([]interface{}))
		}
		if presentFileIntegrityRule["metadata"] != nil {
			parsedFileIntegrityRule.Metadata = presentFileIntegrityRule["metadata"].(bool)
		}
		if presentFileIntegrityRule["path"] != nil {
			parsedFileIntegrityRule.Path = presentFileIntegrityRule["path"].(string)
		}
		if presentFileIntegrityRule["read"] != nil {
			parsedFileIntegrityRule.Read = presentFileIntegrityRule["read"].(bool)
		}
		if presentFileIntegrityRule["recursive"] != nil {
			parsedFileIntegrityRule.Recursive = presentFileIntegrityRule["recursive"].(bool)
		}
		if presentFileIntegrityRule["write"] != nil {
			parsedFileIntegrityRule.Write = presentFileIntegrityRule["write"].(bool)
		}
		parsedFileIntegrityRules = append(parsedFileIntegrityRules, parsedFileIntegrityRule)
	}
	return parsedFileIntegrityRules
}

func parseRuntimeHostForensic(in []interface{}) policy.RuntimeHostForensic {
	parsedForensic := policy.RuntimeHostForensic{}
	if in[0] == nil {
		return parsedForensic
	}
	presentForensic := in[0].(map[string]interface{})
	if presentForensic["activities_disabled"] != nil {
		parsedForensic.ActivitiesDisabled = presentForensic["activities_disabled"].(bool)
	}
	if presentForensic["docker_enabled"] != nil {
		parsedForensic.DockerEnabled = presentForensic["docker_enabled"].(bool)
	}
	if presentForensic["readonly_docker_enabled"] != nil {
		parsedForensic.ReadonlyDockerEnabled = presentForensic["readonly_docker_enabled"].(bool)
	}
	if presentForensic["service_activities_enabled"] != nil {
		parsedForensic.ServiceActivitiesEnabled = presentForensic["service_activities_enabled"].(bool)
	}
	if presentForensic["sshd_enabled"] != nil {
		parsedForensic.SshdEnabled = presentForensic["sshd_enabled"].(bool)
	}
	if presentForensic["sudo_enabled"] != nil {
		parsedForensic.SudoEnabled = presentForensic["sudo_enabled"].(bool)
	}
	return parsedForensic
}

func parseRuntimeHostLogInspectionRules(in []interface{}) []policy.RuntimeHostLogInspectionRule {
	parsedLogInspectionRules := make([]policy.RuntimeHostLogInspectionRule, 0, len(in))
	for _, val := range in {
		presentLogInspectionRule := val.(map[string]interface{})
		parsedLogInspectionRule := policy.RuntimeHostLogInspectionRule{}
		if presentLogInspectionRule["path"] != nil {
			parsedLogInspectionRule.Path = presentLogInspectionRule["path"].(string)
		}
		if presentLogInspectionRule["regex"] != nil {
			parsedLogInspectionRule.Regex = parseStringArray(presentLogInspectionRule["regex"].([]interface{}))
		}
		parsedLogInspectionRules = append(parsedLogInspectionRules, parsedLogInspectionRule)
	}
	return parsedLogInspectionRules
}

func parseRuntimeHostNetwork(in []interface{}) policy.RuntimeHostNetwork {
	parsedNetwork := policy.RuntimeHostNetwork{}
	if in[0] == nil {
		return parsedNetwork
	}
	presentNetwork := in[0].(map[string]interface{})
	if presentNetwork["allowed_outbound_ips"] != nil {
		parsedNetwork.AllowedOutboundIps = parseStringArray(presentNetwork["allowed_outbound_ips"].([]interface{}))
	}
	if presentNetwork["custom_feed"] != nil {
		parsedNetwork.CustomFeed = presentNetwork["custom_feed"].(string)
	}
	if presentNetwork["denied_listening_port"] != nil {
		parsedNetwork.DeniedListeningPorts = parseRuntimeHostPorts(presentNetwork["denied_listening_port"].([]interface{}))

	}
	if presentNetwork["denied_outbound_ips"] != nil {
		parsedNetwork.DeniedOutboundIps = parseStringArray(presentNetwork["denied_outbound_ips"].([]interface{}))
	}
	if presentNetwork["denied_outbound_port"] != nil {
		parsedNetwork.DeniedOutboundPorts = parseRuntimeHostPorts(presentNetwork["denied_outbound_port"].([]interface{}))
	}
	if presentNetwork["deny_effect"] != nil {
		parsedNetwork.DenyEffect = presentNetwork["deny_effect"].(string)
	}
	if presentNetwork["intelligence_feed"] != nil {
		parsedNetwork.IntelligenceFeed = presentNetwork["intelligence_feed"].(string)
	}
	return parsedNetwork
}

func parseRuntimeHostDeniedProcesses(in []interface{}) policy.RuntimeHostDeniedProcesses {
	parsedDeniedProcesses := policy.RuntimeHostDeniedProcesses{}
	if in[0] == nil {
		return parsedDeniedProcesses
	}
	presentDeniedProcesses := in[0].(map[string]interface{})
	if presentDeniedProcesses["effect"] != nil {
		parsedDeniedProcesses.Effect = presentDeniedProcesses["effect"].(string)
	}
	if presentDeniedProcesses["paths"] != nil {
		parsedDeniedProcesses.Paths = parseStringArray(presentDeniedProcesses["paths"].([]interface{}))
	}
	return parsedDeniedProcesses
}

func parseRuntimeHostPorts(in []interface{}) []policy.RuntimeHostPort {
	parsedPorts := make([]policy.RuntimeHostPort, 0, len(in))
	for _, val := range in {
		presentPort := val.(map[string]interface{})
		parsedPort := policy.RuntimeHostPort{}
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

func flattenPolicyRuntimeHostRules(in []policy.RuntimeHostRule) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["antimalware"] = flattenRuntimeHostAntiMalware(val.AntiMalware)
		m["collections"] = flattenCollections(val.Collections)
		m["custom_rule"] = flattenRuntimeHostCustomRules(val.CustomRules)
		m["disabled"] = val.Disabled
		m["dns"] = flattenRuntimeHostDns(val.Dns)
		m["file_integrity_rule"] = flattenRuntimeHostFileIntegrityRules(val.FileIntegrityRules)
		m["forensic"] = flattenRuntimeHostForensic(val.Forensic)
		m["log_inspection_rule"] = flattenRuntimeHostLogInspectionRules(val.LogInspectionRules)
		m["name"] = val.Name
		m["network"] = flattenRuntimeHostNetwork(val.Network)
		m["notes"] = val.Notes
		ans = append(ans, m)
	}
	return ans
}
