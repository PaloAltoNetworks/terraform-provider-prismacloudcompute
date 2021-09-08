package prismacloudcompute

import (
	"log"

	pcc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policies"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourcePoliciesRuntimeContainer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePoliciesRuntimeContainerRead,

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
			"learningdisabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If set to 'true', automatic behavioural learning is enabled.",
			},
			"rules": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of rules in the policy.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advancedprotection": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If set to 'true', enables advanced protection (e.g., custom or premium feeds for container, added whitelist rules for serverless).",
						},
						"cloudmetadataenforcement": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Catches containers that access the cloud provider metadata API.",
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
										Required:    true,
										Description: "List of account IDs.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"appids": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "List of application IDs.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"clusters": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "List of Kubernetes cluster names.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"coderepos": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "List of code repositories.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"color": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Hex color code for a collection.",
									},
									"containers": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "List of containers.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"description": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "A free-form text description of the collection.",
									},
									"functions": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "List of functions.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"hosts": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "List of hosts.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"images": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "List of images.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"labels": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "List of labels.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"modified": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Date/time when the collection was last modified.",
									},
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Unique collection name.",
									},
									"namespaces": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "List of Kubernetes namespaces.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"owner": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "User who created or last modified the collection.",
									},
									"prisma": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "If set to 'true' this collection originates from Prisma Cloud.",
									},
									"system": {
										Type:        schema.TypeString,
										Required:    true,
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
									"effect": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The effect used in the runtime rule. Can be set to 'block', 'prevent', 'alert', 'disable'.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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
						"filesystem": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Represents restrictions or suppression for filesystem changes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"backdoorfiles": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "If set to 'true', monitors files that can create or persist backdoors (SSH or admin account config files).",
									},
									"blacklist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of denied file system paths.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"checknewfiles": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "If set to 'true', detects changes to binaries and certificates.",
									},
									"effect": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The effect that will be used in the runtime rule. Can be set to 'block', 'prevent', 'alert', or 'disable'.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"skipencryptedbinaries": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "If set to 'true', the encrypted binaries check should be skipped.",
									},
									"suspiciouselfheaders": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "If set to 'true', enables malware detection based on suspicious ELF headers.",
									},
									"whitelist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of allowed file system paths.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"kubernetesenforcement": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Detects containers that attempt to compromise the orchestrator.",
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
									"detectportscan": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "If set to 'true' port scanning detection is enabled.",
									},
									"effect": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The effect that will be used in the runtime rule. Vaules: ['block', 'prevent', 'alert', 'disable'].",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"skipmodifiedproc": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "If set to 'true', Prisma Cloud can detect malicious networking activity from modified processes.",
									},
									"skiprawsockets": {
										Type:        schema.TypeBool,
										Required:    true,
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
							Required:    true,
							Description: "A free-form text description of the collection.",
						},
						"owner": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "User who created or last modified the rule.",
						},
						"previousname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Previous name of the rule. Required for rule renaming.",
						},
						"processes": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Represents restrictions or suppression for running processes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"blacklist": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "List of processes to deny.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"blockallbinaries": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "If set to 'true', all processes are blocked except the main process.",
									},
									"checkcryptominers": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Detect crypto miners.",
									},
									"checklateralmovement": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "If set to 'true', enables dectection of processes that can be used for lateral movement exploits.",
									},
									"checknewbinaries": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "If set to 'true', binaries which don't belong to the original image are allowed to run.",
									},
									"checkparentchild": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "If set to 'true', enables checking for parent child relationship when comparing spawned processes in the model.",
									},
									"checksuidbinaries": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "If set to 'true', enables check for process elevating privileges (SUID bit).",
									},
									"effect": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "The effect that will be used in the runtime rule. Can be set to 'block', 'prevent', 'alert', or 'disable'.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"skipmodified": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Indicates whether to trigger audits/incidents when a modified proc is spawned.",
									},
									"skipreverseshell": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Indicates whether reverse shell detection is disabled.",
									},
									"whitelist": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "List of processes to allow.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"wildfireanalysis": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "The effect that will be used in the runtime rule. Can be set to 'block', 'prevent', 'alert', or 'disable'.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourcePoliciesRuntimeContainerRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pcc.Client)

	i, err := policies.Get(*client, policies.RuntimeContainerEndpoint)
	if err != nil {
		return err
	}

	d.SetId(i.PolicyId)

	list := make([]interface{}, 0, 1)
	list = append(list, map[string]interface{}{
		"_id":              i.PolicyId,
		"learningdisabled": i.LearningDisabled,
		"rules":            i.Rules,
	})

	if err := d.Set("listing", list); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}

	return nil
}
