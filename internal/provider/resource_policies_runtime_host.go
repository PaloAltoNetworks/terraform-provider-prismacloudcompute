package provider

import (
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/policy"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePoliciesRuntimeHost() *schema.Resource {
	return &schema.Resource{
		Create: createPolicyRuntimeHost,
		Read:   readPolicyRuntimeHost,
		Update: updatePolicyRuntimeHost,
		Delete: deletePolicyRuntimeHost,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the policy.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"rule": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Rules that make up the policy.",
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"activities": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Activities configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"disabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to disable host activity collection.",
									},
									"docker_enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to collect docker commands.",
									},
									"readonly_docker_enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to collect read-only docker commands.",
									},
									"service_activities_enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to collect activity from services.",
									},
									"sshd_enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to collect new SSH sessions.",
									},
									"sudo_enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to collect commands ran with sudo or su.",
									},
								},
							},
						},
						"antimalware": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Anti-malware configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allowed_processes": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of processes and files to allow during anti-malware checks.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"crypto_miners": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used when crypto miners are detected. Can be set to 'prevent', 'alert', or 'disable'.",
									},
									"custom_feed": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used when malware from custom feeds is detected. Can be set to 'alert' or 'disable'.",
									},
									"denied_processes": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Denied processes configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"effect": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The effect to be used. Can be set to 'prevent' or 'alert'.",
												},
												"paths": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "List of processes and files to deny during anti-malware checks.",
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
										Description: "Whether or not to detect compiler-generated binaries.",
									},
									"encrypted_binaries": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used when encrypted or packed binaries are detected. Can be set to 'alert' or 'disable'.",
									},
									"execution_flow_hijack": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used when execution flow hijacking is detected. Can be set to 'alert' or 'disable'.",
									},
									"intelligence_feed": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used when malware according to Prisma Cloud Compute is detected. Can be set to 'alert' or 'disable'.",
									},
									"reverse_shell": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used when reverse shell attacks are detected. Can be set to 'alert' or 'disable'.",
									},
									"service_unknown_origin_binary": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used when non-packaged binaries are created or ran by a service. Can be set to 'prevent', 'alert', or 'disable'.",
									},
									"skip_ssh_tracking": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to skip tracking of SSH events.",
									},
									"suspicious_elf_headers": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used when binaries with suspicious ELF headers are detected. Can be set to 'alert' or 'disable'.",
									},
									"temp_filesystem_processes": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used when processes are ran from a temporary file system. Can be set to 'prevent', 'alert', or 'disable'.",
									},
									"user_unknown_origin_binary": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used when non-packaged binaries are created or ran by a user. Can be set to 'prevent', 'alert', or 'disable'.",
									},
									"webshell": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used when webshell attacks are detected. Can be set to 'prevent', 'alert', or 'disable'.",
									},
									"wildfire_analysis": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used when WildFire analysis is enabled. Can be set to 'alert' or 'disable'.",
									},
								},
							},
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
										Description: "The effect to be used. Can be set to 'prevent', 'alert', or 'allow'.",
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
										Description: "The effect to be used. Can be set to 'prevent', 'alert', or 'disable'.",
									},
									"intelligence_feed": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used when resolving suspicious domains according to Prisma Cloud Compute. Can be set to 'prevent', 'alert', or 'disable'.",
									},
								},
							},
						},
						"file_integrity_rule": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of file integrity rules.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allowed_processes": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of processes allowed to generate file system events on monitored files.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"excluded_files": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of file names to ignore. Pattern matching is supported.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"metadata": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to monitor file metadata changes.",
									},
									"path": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Path to monitor.",
									},
									"read": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to monitor file reads.",
									},
									"recursive": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to recursively monitor files starting at `path`.",
									},
									"write": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether or not to monitor file writes.",
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
										Description: "Path to the log file.",
									},
									"regex": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of regular expressions to use when inspecting the log file.",
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
							Description: "Unique name of the rule.",
						},
						"network": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Network configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allowed_outbound_ips": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of allowed outbound IP addresses.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"custom_feed": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used when connecting to suspicious IPs according to custom feeds. Can be set to 'alert' or 'disable'.",
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
										Description: "The effect to be used. Can be set to 'alert' or 'disable'.",
									},
									"intelligence_feed": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used when connecting to suspicious IPs according to Prisma Cloud Compute. Can be set to 'alert' or 'disable'.",
									},
								},
							},
						},
						"notes": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Free-form text field.",
						},
					},
				},
			},
		},
	}
}

func createPolicyRuntimeHost(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	parsedRules, err := convert.SchemaToRuntimeHostRules(d)
	if err != nil {
		return fmt.Errorf("error creating %s policy: %s", policyTypeRuntimeHost, err)
	}

	parsedPolicy := policy.RuntimeHostPolicy{
		Rules: parsedRules,
	}

	if err := policy.UpdateRuntimeHost(*client, parsedPolicy); err != nil {
		return fmt.Errorf("error creating %s policy: %s", policyTypeRuntimeHost, err)
	}

	d.SetId(policyTypeRuntimeHost)
	return readPolicyRuntimeHost(d, meta)
}

func readPolicyRuntimeHost(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	retrievedPolicy, err := policy.GetRuntimeHost(*client)
	if err != nil {
		return fmt.Errorf("error reading %s policy: %s", policyTypeRuntimeHost, err)
	}

	if err := d.Set("rule", convert.RuntimeHostRulesToSchema(retrievedPolicy.Rules)); err != nil {
		return fmt.Errorf("error reading %s policy: %s", policyTypeRuntimeHost, err)
	}

	return nil
}

func updatePolicyRuntimeHost(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	parsedRules, err := convert.SchemaToRuntimeHostRules(d)
	if err != nil {
		return fmt.Errorf("error updating %s policy: %s", policyTypeRuntimeHost, err)
	}

	parsedPolicy := policy.RuntimeHostPolicy{
		Rules: parsedRules,
	}

	if err := policy.UpdateRuntimeHost(*client, parsedPolicy); err != nil {
		return fmt.Errorf("error updating %s policy: %s", policyTypeRuntimeHost, err)
	}

	return readPolicyRuntimeHost(d, meta)
}

func deletePolicyRuntimeHost(d *schema.ResourceData, meta interface{}) error {
	// TODO: reset to default policy
	return nil
}
