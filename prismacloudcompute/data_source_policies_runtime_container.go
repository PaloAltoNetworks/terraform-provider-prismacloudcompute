package prismacloudcompute

import (
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy/policyRuntimeContainer"

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
				Description: "ID of the policy set",
			},
			"learningdisabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determines if learning is enabled or not",
			},
			"rules": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of policy rules",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advancedprotection": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether advanced protection (e.g., custom or premium feeds for container, added whitelist rules for serverless) is enabled (true) or not (false).",
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
										Description: "Account IDs",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"appids": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "App IDs",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"clusters": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Clusters",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"coderepos": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Code repositories",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"color": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Color",
									},
									"containers": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Containers",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"description": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Description",
									},
									"functions": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Serverless functions",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"hosts": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Hosts",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"images": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Images",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"labels": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Labels",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"modified": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Last modified date",
									},
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Name",
									},
									"namespaces": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Namespaces",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"owner": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Owner",
									},
									"prisma": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Prisma",
									},
									"system": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "System",
									},
								},
							},
						},
						"customrules": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of custom rules.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"_id": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Rule ID",
									},
									"action": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of actions.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"effect": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of effects.",
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
							Description: "disabled",
						},
						"dns": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of dns.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"blacklist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of blacklist items.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"effect": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of blacklist items.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"whitelist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of blacklist items.",
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
							Description: "List of filesystems.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"backdoorfiles": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "backdoorFiles",
									},
									"blacklist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of blacklist items.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"checknewfiles": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "checkNewFiles",
									},
									"effect": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of effects.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"skipencryptedbinaries": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "skipEncryptedBinaries",
									},
									"suspiciouselfheaders": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "suspiciousELFHeaders",
									},
									"whitelist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of whitelist.",
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
							Description: "kubernetesEnforcement",
						},
						"modified": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "modified",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "name",
						},
						"network": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of networks.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"backdoorfiles": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "backdoorFiles",
									},
									"blacklistips": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of blacklistIPs.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"blacklistlisteningports": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of blacklistListeningPorts.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"deny": {
													Type:        schema.TypeBool,
													Required:    true,
													Description: "deny",
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
										Description: "List of blacklistOutboundPorts.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"deny": {
													Type:        schema.TypeBool,
													Required:    true,
													Description: "deny",
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
										Description: "detectPortScan",
									},
									"effect": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of effects.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"skipmodifiedproc": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "skipModifiedProc",
									},
									"skiprawsockets": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "skipRawSockets",
									},
									"whitelistips": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of whitelistIPs.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"whitelistlisteningports": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of whitelistListeningPorts.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"deny": {
													Type:        schema.TypeBool,
													Required:    true,
													Description: "deny",
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
										Description: "List of whitelistOutboundPorts.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"deny": {
													Type:        schema.TypeBool,
													Required:    true,
													Description: "deny",
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
							Description: "notes",
						},
						"owner": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "owner",
						},
						"previousname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "previousName",
						},
						"processes": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of processes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"processes": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of processes.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"blacklist": {
													Type:        schema.TypeList,
													Required:    true,
													Description: "blacklist",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"blockallbinaries": {
													Type:        schema.TypeBool,
													Required:    true,
													Description: "blockAllBinaries",
												},
												"checkcryptominers": {
													Type:        schema.TypeBool,
													Required:    true,
													Description: "checkCryptoMiners",
												},
												"checklateralmovement": {
													Type:        schema.TypeBool,
													Required:    true,
													Description: "checkLateralMovement",
												},
												"checknewbinaries": {
													Type:        schema.TypeBool,
													Required:    true,
													Description: "checkNewBinaries",
												},
												"checkparentchild": {
													Type:        schema.TypeBool,
													Required:    true,
													Description: "checkParentChild",
												},
												"checksuidbinaries": {
													Type:        schema.TypeBool,
													Required:    true,
													Description: "checkSuidBinaries",
												},
												"effect": {
													Type:        schema.TypeList,
													Required:    true,
													Description: "effect",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"skipmodified": {
													Type:        schema.TypeBool,
													Required:    true,
													Description: "skipModified",
												},
												"skipreverseshell": {
													Type:        schema.TypeBool,
													Required:    true,
													Description: "skipReverseShell",
												},
												"whitelist": {
													Type:        schema.TypeList,
													Required:    true,
													Description: "blacklist",
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
						"wildfireanalysis": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "wildFireAnalysis",
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
	client := meta.(*pc.Client)

	i, err := policyRuntimeContainer.Get(client)
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
