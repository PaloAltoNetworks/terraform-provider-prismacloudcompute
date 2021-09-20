package prismacloudcompute

import (
	"log"

	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePoliciesRuntimeHost() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePoliciesRuntimeHostRead,

		Schema: map[string]*schema.Schema{
			// Input.
			"filters": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Filter policy results",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Output.
			"_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the policy set.",
			},
			"owner": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The policy owner.",
			},
			"rules": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of rules in the policy.",
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
									"detectvompilergeneratedbinary": {
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
							Description: "List of collections. Used to scope the rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									// Output.
									"account_ids": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of account IDs.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"application_ids": {
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
									"code_repositories": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of code repositories.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"color": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Hex color code for a collection.",
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
										Type:        schema.TypeString,
										Optional:    true,
										Description: "A free-form text description of the collection.",
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
										Optional:    true,
										Description: "Date/time when the collection was last modified.",
									},
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Unique collection name.",
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
										Optional:    true,
										Description: "User who created or last modified the collection.",
									},
									"prisma": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "If set to 'true' this collection originates from Prisma Cloud.",
									},
									"system": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "If set to 'true', this collection was created by the system (i.e., a non-user). Otherwise it was created by a real user.",
									},
								},
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
										Required:    true,
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
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The DNS runtime rule",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"denylist": {
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
									"effect": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The effect used in the runtime rule. Can be set to 'block', 'prevent', 'alert', 'disable'.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"intelligence_feed": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"allowlist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of allowed domain names (e.g., *.gmail.com, *.s3.amazon.com).",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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
									"dir": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Indicates that the path is a directory.",
									},
									"exclusions": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Filenames that should be ignored while generating audits These filenames may contain a wildcard regex pattern, e.g. foo*.log, *.cache.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"metadata": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Indicates that metadata changes should be monitored (e.g. chmod, chown).",
									},
									"path": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Path to monitor.",
									},
									"allowed_processes": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "The processes to ignore Filesystem events caused by these processes DO NOT generate file integrity events.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"read": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Indicates that eads operations should be monitored.",
									},
									"recursive": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Indicates that monitoring should be recursive.",
									},
									"write": {
										Type:        schema.TypeBool,
										Required:    true,
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
										Required:    true,
										Description: "Indicates if the host activity collection is enabled/disabled.",
									},
									"docker_enabled": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Indicates whether docker commands are collected.",
									},
									"readonly_docker_enabled": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Indicates whether docker readonly commands are collected.",
									},
									"service_activities_enabled": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Indicates whether activities from services are collected.",
									},
									"sshd_enabled": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Indicates whether ssh commands are collected.",
									},
									"sudo_enabled": {
										Type:        schema.TypeBool,
										Required:    true,
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
										Required:    true,
										Description: "the log path.",
									},
									"regex": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Regular expressions associated with the rule if it is a custom one.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"modified": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Date/time when the rule was last modified.",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the rule.",
						},
						"network": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Represents the restrictions or suppression for networking.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"denied_outbound_ips": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Deny-list of IP addresses.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"denied_listening_port": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Deny-list of listening ports.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"deny": {
													Type:        schema.TypeBool,
													Required:    true,
													Description: "If set to 'true' the connection is denied.",
												},
												"end": {
													Type:        schema.TypeInt,
													Required:    true,
													Description: "end",
												},
												"start": {
													Type:        schema.TypeInt,
													Required:    true,
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
													Required:    true,
													Description: "If set to 'true' the connection is denied.",
												},
												"end": {
													Type:        schema.TypeInt,
													Required:    true,
													Description: "end",
												},
												"start": {
													Type:        schema.TypeInt,
													Required:    true,
													Description: "start",
												},
											},
										},
									},
									"custom_feed": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"detect_port_scan": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true' port scanning detection is enabled.",
									},
									"deny_effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"effect": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The effect used in the runtime rule. Can be set to 'block', 'prevent', 'alert', 'disable'.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"intelligence_feed": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
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
													Required:    true,
													Description: "If set to 'true', the connection is denied.",
												},
												"end": {
													Type:        schema.TypeInt,
													Required:    true,
													Description: "end",
												},
												"start": {
													Type:        schema.TypeInt,
													Required:    true,
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
													Required:    true,
													Description: "If set to 'true', the connection is denied.",
												},
												"end": {
													Type:        schema.TypeInt,
													Required:    true,
													Description: "end",
												},
												"start": {
													Type:        schema.TypeInt,
													Required:    true,
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
						"owner": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "User who created or last modified the rule.",
						},
						"previousname": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Previous name of the rule. Required for rule renaming.",
						},
					},
				},
			},
		},
	}
}

func dataSourcePoliciesRuntimeHostRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)

	i, err := policy.GetRuntimeHost(*client)
	if err != nil {
		return err
	}

	d.SetId(policyTypeRuntimeHost)

	list := make([]interface{}, 0, 1)
	list = append(list, map[string]interface{}{
		"rules": i.Rules,
	})

	if err := d.Set("listing", list); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}

	return nil
}
