package prismacloudcompute

import (
	"log"

	pcc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policies"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Restrictions/suppression for suspected anti-malware.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allowedprocesses": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "A list of paths for files and processes to skip during anti-malware checks.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"cryptominer": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"customfeed": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"deniedprocesses": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "A rule containing paths of files and processes to alert/prevent and the required effect.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"effect": {
													Type:        schema.TypeString,
													Required:    false,
													Description: "Effect that will be used in the runtime rule.",
												},
												"paths": {
													Type:        schema.TypeList,
													Required:    false,
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
									"encryptedbinaries": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"executionflowhijack": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"intelligencefeed": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"reverseshell": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"serviceunknownoriginbinary": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"skipsshtracking": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"suspiciouselfheaders": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"tempfsproc": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"userunknownoriginbinary": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"webshell": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"wildfireanalysis": {
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
									"accountids": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of account IDs.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"appids": {
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
									"coderepos": {
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
						"customrules": {
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
									"blacklist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Deny-list of domain names (e.g., www.bad-url.com, *.bad-url.com).",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"denylisteffect": {
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
									"intelligencefeed": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"whitelist": {
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
						"fileintegrityrules": {
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
									"procwhitelist": {
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
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Indicates how to perform host forensic.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"activitiesdisabled": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Indicates if the host activity collection is enabled/disabled.",
									},
									"dockerenabled": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Indicates whether docker commands are collected.",
									},
									"readonlydockerenabled": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Indicates whether docker readonly commands are collected.",
									},
									"serviceactivitiesenabled": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Indicates whether activities from services are collected.",
									},
									"sshdenabled": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Indicates whether ssh commands are collected.",
									},
									"sudoenabled": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Indicates whether sudo commands are collected.",
									},
								},
							},
						},
						"loginspectionrules": {
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
									"blacklistips": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Deny-list of IP addresses.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"blacklistlisteningports": {
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
									"blacklistoutboundports": {
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
									"customfeed": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"detectportscan": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true' port scanning detection is enabled.",
									},
									"denylisteffect": {
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
									"intelligencefeed": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect that will be used in the runtime rule.",
									},
									"skipmodifiedproc": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', Prisma Cloud can detect malicious networking activity from modified processes.",
									},
									"skiprawsockets": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', raw socket detection will be skipped.",
									},
									"whitelistips": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Allow-listed IP addresses.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"whitelistlisteningports": {
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
									"whitelistoutboundports": {
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

	i, err := policies.Get(*client, policies.RuntimeHostEndpoint)
	if err != nil {
		return err
	}

	d.SetId(i.PolicyId)

	list := make([]interface{}, 0, 1)
	list = append(list, map[string]interface{}{
		"_id":   i.PolicyId,
		"rules": i.Rules,
	})

	if err := d.Set("listing", list); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}

	return nil
}
