package prismacloudcompute

import (
	"log"
	"time"

	pc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/collection"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy/policyRuntimeContainer"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePoliciesRuntimeContainer() *schema.Resource {
	return &schema.Resource{
		Create: createPolicy,
		Read:   readPolicy,
		Update: updatePolicy,
		Delete: deletePolicy,

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
				Description: "ID",
			},
			"learningdisabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Learning disabled",
				Default:     true,
			},
			"rules": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Rules for the policies",
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name",
						},
						"collections": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Collections",
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"accountids": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Account IDs",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"appids": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "App IDs",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"clusters": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Clusters",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"coderepos": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Code repositories",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"color": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Color",
									},
									"containers": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Containers",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"description": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Description",
									},
									"functions": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Serverless functions",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"hosts": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Hosts",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"images": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Images",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"labels": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Labels",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"modified": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Last modified date",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Name",
									},
									"namespaces": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Namespaces",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"owner": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Owner",
									},
									"prisma": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Prisma",
									},
									"system": {
										Type:        schema.TypeBool,
										Optional:    true,
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
										Optional:    true,
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
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "DNS",
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
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "List of filesystems.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"backdoorfiles": {
										Type:        schema.TypeBool,
										Optional:    true,
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
										Optional:    true,
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
										Optional:    true,
										Description: "skipEncryptedBinaries",
									},
									"suspiciouselfheaders": {
										Type:        schema.TypeBool,
										Optional:    true,
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
							Description: "kubernetesenforcement",
						},
						"modified": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "modified",
						},
						"network": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "List of networks.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"backdoorfiles": {
										Type:        schema.TypeBool,
										Optional:    true,
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
													Optional:    true,
													Description: "deny",
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
									"blacklistoutboundports": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of blacklistOutboundPorts.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"deny": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "deny",
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
									"detectportscan": {
										Type:        schema.TypeBool,
										Optional:    true,
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
										Optional:    true,
										Description: "skipModifiedProc",
									},
									"skiprawsockets": {
										Type:        schema.TypeBool,
										Optional:    true,
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
													Optional:    true,
													Description: "deny",
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
									"whitelistoutboundports": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of whitelistOutboundPorts.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"deny": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "deny",
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
									"notes": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "notes",
									},
									"owner": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "owner",
									},
									"previousname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "previousName",
									},
									"processes": {
										Type:        schema.TypeMap,
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
																Optional:    true,
																Description: "blacklist",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"blockallbinaries": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "blockAllBinaries",
															},
															"checkcryptominers": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "checkCryptoMiners",
															},
															"checklateralmovement": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "checkLateralMovement",
															},
															"checknewbinaries": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "checkNewBinaries",
															},
															"checkparentchild": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "checkParentChild",
															},
															"checksuidbinaries": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "checkSuidBinaries",
															},
															"effect": {
																Type:        schema.TypeList,
																Optional:    true,
																Description: "effect",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"skipmodified": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "skipModified",
															},
															"skipreverseshell": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "skipReverseShell",
															},
															"whitelist": {
																Type:        schema.TypeList,
																Optional:    true,
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
										Optional:    true,
										Description: "wildFireAnalysis",
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
		},
	}
}

func parsePolicy(d *schema.ResourceData, id string) policyRuntimeContainer.Policy {
	ans := policyRuntimeContainer.Policy{
		PolicyId:         id,
		LearningDisabled: d.Get("learningdisabled").(bool),
	}

	rules := d.Get("rules").([]interface{})
	ans.Rules = make([]policy.Rule, 0, len(rules))
	if len(rules) > 0 {

		item := rules[0].(map[string]interface{})

		rule := policy.Rule{}

		if item["advancedprotection"] != nil {
			rule.AdvancedProtection = item["advancedprotection"].(bool)
		}
		if item["cloudmetadataenforcement"] != nil {
			rule.CloudMetadataEnforcement = item["cloudmetadataenforcement"].(bool)
		}
		if item["collections"] != nil {
			colls := item["collections"].([]interface{})
			rule.Collections = make([]collection.Collection, 0, len(colls))
			if len(colls) > 0 {
				collItem := colls[0].(map[string]interface{})

				coll := collection.Collection{
					Name: collItem["name"].(string),
				}
				if collItem["accountIDs"] != nil && len(collItem["accountIDs"].([]interface{})) > 0 {
					coll.AccountIDs = collItem["accountIDs"].([]interface{})[0].([]string)
				}
				if collItem["appIDs"] != nil && len(collItem["appIDs"].([]interface{})) > 0 {
					coll.AppIDs = collItem["appIDs"].([]interface{})[0].([]string)
				}
				if collItem["clusters"] != nil && len(collItem["clusters"].([]interface{})) > 0 {
					coll.Clusters = collItem["clusters"].([]interface{})[0].([]string)
				}
				if collItem["codeRepos"] != nil && len(collItem["codeRepos"].([]interface{})) > 0 {
					coll.CodeRepos = collItem["codeRepos"].([]interface{})[0].([]string)
				}
				if collItem["color"] != nil {
					coll.Color = collItem["color"].(string)
				}
				if collItem["containers"] != nil && len(collItem["containers"].([]interface{})) > 0 {
					coll.Containers = collItem["containers"].([]interface{})[0].([]string)
				}
				if collItem["description"] != nil {
					coll.Description = collItem["description"].(string)
				}
				if collItem["functions"] != nil && len(collItem["functions"].([]interface{})) > 0 {
					coll.Functions = collItem["functions"].([]interface{})[0].([]string)
				}
				if collItem["hosts"] != nil && len(collItem["hosts"].([]interface{})) > 0 {
					coll.Hosts = collItem["hosts"].([]interface{})[0].([]string)
				}
				if collItem["images"] != nil && len(collItem["images"].([]interface{})) > 0 {
					coll.Images = collItem["images"].([]interface{})[0].([]string)
				}
				if collItem["labels"] != nil && len(collItem["labels"].([]interface{})) > 0 {
					coll.Labels = collItem["labels"].([]interface{})[0].([]string)
				}
				if collItem["modified"] != nil {
					coll.Modified = collItem["modified"].(string)
				}
				if collItem["namespaces"] != nil && len(collItem["namespaces"].([]interface{})) > 0 {
					coll.Namespaces = collItem["namespaces"].([]interface{})[0].([]string)
				}
				if collItem["owner"] != nil {
					coll.Owner = collItem["owner"].(string)
				}
				if collItem["prisma"] != nil {
					coll.Prisma = collItem["prisma"].(bool)
				}
				if collItem["system"] != nil {
					coll.System = collItem["system"].(bool)
				}

				rule.Collections = append(rule.Collections, coll)
			}

		}
		if item["customrules"] != nil {
			custRules := item["customrules"].([]interface{})
			rule.CustomRules = make([]policy.CustomRule, 0, len(custRules))
			if len(custRules) > 0 {
				custRuleItem := custRules[0].(map[string]interface{})

				custRule := policy.CustomRule{
					Id:     custRuleItem["_id"].(int),
					Action: custRuleItem["action"].([]string),
					Effect: custRuleItem["effect"].([]string),
				}
				rule.CustomRules = append(rule.CustomRules, custRule)
			}
		}
		if item["disabled"] != nil {
			rule.Disabled = item["disabled"].(bool)
		}
		if item["dns"] != nil {
			dnsSet := item["dns"].(interface{})
			dnsItem := dnsSet.(map[string]interface{})

			rule.Dns = policy.Dns{}
			if dnsItem["blacklist"] != nil {
				rule.Dns.Blacklist = dnsItem["blacklist"].([]string)
			}
			if dnsItem["effect"] != nil {
				rule.Dns.Effect = dnsItem["effect"].([]string)
			}
			if dnsItem["whitelist"] != nil {
				rule.Dns.Whitelist = dnsItem["whitelist"].([]string)
			}
		}
		if item["filesystem"] != nil {
			fileSysSet := item["filesystem"].(interface{})
			fileSysItem := fileSysSet.(map[string]interface{})

			rule.Filesystem = policy.Filesystem{}
			if fileSysItem["blacklist"] != nil {
				rule.Filesystem.BackdoorFiles = fileSysItem["backdoorFiles"].(bool)
			}
			if fileSysItem["blacklist"] != nil {
				rule.Filesystem.Blacklist = fileSysItem["blacklist"].([]string)
			}
			if fileSysItem["checkNewFiles"] != nil {
				rule.Filesystem.CheckNewFiles = fileSysItem["checkNewFiles"].(bool)
			}
			if fileSysItem["effect"] != nil {
				rule.Filesystem.Effect = fileSysItem["effect"].([]string)
			}
			if fileSysItem["skipEncryptedBinaries"] != nil {
				rule.Filesystem.SkipEncryptedBinaries = fileSysItem["skipEncryptedBinaries"].(bool)
			}
			if fileSysItem["suspiciousELFHeaders"] != nil {
				rule.Filesystem.SuspiciousELFHeaders = fileSysItem["suspiciousELFHeaders"].(bool)
			}
			if fileSysItem["whitelist"] != nil {
				rule.Filesystem.Whitelist = fileSysItem["whitelist"].(string)
			}
		}
		if item["kubernetesenforcement"] != nil {
			rule.KubernetesEnforcement = item["kubernetesenforcement"].(bool)
		}
		if item["modified"] != nil {
			rule.Modified = item["modified"].(string)
		}
		if item["name"] != nil {
			rule.Name = item["name"].(string)
		}
		if item["network"] != nil {
			networkSet := item["network"].(interface{})
			networkItem := networkSet.(map[string]interface{})
			if networkItem["blacklistIPs"] != nil {
				rule.Network.BlacklistIPs = networkItem["blacklistIPs"].([]string)
			}

			if networkItem["blacklistListeningPorts"] != nil {
				blacklistListenPorts := networkItem["blacklistListeningPorts"].([]interface{})
				rule.Network.BlacklistListeningPorts = make([]policy.ListPort, 0, len(blacklistListenPorts))
				if len(blacklistListenPorts) > 0 {
					rule.Network.BlacklistListeningPorts = append(rule.Network.BlacklistListeningPorts, getListPort(blacklistListenPorts[0]))
				}
			}

			if networkItem["blacklistOutboundPorts"] != nil {
				blacklistOutPorts := networkItem["blacklistOutboundPorts"].([]interface{})
				rule.Network.BlacklistOutboundPorts = make([]policy.ListPort, 0, len(blacklistOutPorts))
				if len(blacklistOutPorts) > 0 {
					rule.Network.BlacklistOutboundPorts = append(rule.Network.BlacklistOutboundPorts, getListPort(blacklistOutPorts[0]))
				}
			}
			if networkItem["blacklistOutboundPorts"] != nil {
				rule.Network.DetectPortScan = networkItem["detectPortScan"].(bool)
			}
			if networkItem["effect"] != nil {
				rule.Network.Effect = networkItem["effect"].([]string)
			}
			if networkItem["skipModifiedProc"] != nil {
				rule.Network.SkipModifiedProc = networkItem["skipModifiedProc"].(bool)
			}
			if networkItem["skipRawSockets"] != nil {
				rule.Network.SkipRawSockets = networkItem["skipRawSockets"].(bool)
			}
			if networkItem["whitelistIPs"] != nil {
				rule.Network.WhitelistIPs = networkItem["whitelistIPs"].(string)
			}

			if networkItem["whitelistListeningPorts"] != nil {
				whitelistListenPorts := networkItem["whitelistListeningPorts"].([]interface{})
				rule.Network.WhitelistListeningPorts = make([]policy.ListPort, 0, len(whitelistListenPorts))
				if len(whitelistListenPorts) > 0 {
					rule.Network.WhitelistListeningPorts = append(rule.Network.WhitelistListeningPorts, getListPort(whitelistListenPorts[0]))
				}
			}

			if networkItem["whitelistOutboundPorts"] != nil {
				whitelistOutPorts := networkItem["whitelistOutboundPorts"].([]interface{})
				rule.Network.WhitelistOutboundPorts = make([]policy.ListPort, 0, len(whitelistOutPorts))
				if len(whitelistOutPorts) > 0 {
					rule.Network.WhitelistOutboundPorts = append(rule.Network.WhitelistOutboundPorts, getListPort(whitelistOutPorts[0]))
				}
			}
		}
		if item["notes"] != nil {
			rule.Notes = item["notes"].(string)
		}
		if item["owner"] != nil {
			rule.Owner = item["owner"].(string)
		}
		if item["previousname"] != nil {
			rule.PreviousName = item["previousname"].(string)
		}
		if item["processes"] != nil {
			processSet := item["processes"].(interface{})
			processItem := processSet.(map[string]interface{})

			rule.Processes = policy.Processes{}

			if processItem["blacklist"] != nil {
				rule.Processes.Blacklist = processItem["blacklist"].([]string)
			}
			if processItem["blockAllBinaries"] != nil {
				rule.Processes.BlockAllBinaries = processItem["blockAllBinaries"].(bool)
			}
			if processItem["checkCryptoMiners"] != nil {
				rule.Processes.CheckCryptoMiners = processItem["checkCryptoMiners"].(bool)
			}
			if processItem["checkLateralMovement"] != nil {
				rule.Processes.CheckLateralMovement = processItem["checkLateralMovement"].(bool)
			}
			if processItem["checkNewBinaries"] != nil {
				rule.Processes.CheckNewBinaries = processItem["checkNewBinaries"].(bool)
			}
			if processItem["checkParentChild"] != nil {
				rule.Processes.CheckParentChild = processItem["checkParentChild"].(bool)
			}
			if processItem["checkSuidBinaries"] != nil {
				rule.Processes.CheckSuidBinaries = processItem["checkSuidBinaries"].(bool)
			}
			if processItem["effect"] != nil {
				rule.Processes.Effect = processItem["effect"].([]string)
			}
			if processItem["skipModified"] != nil {
				rule.Processes.SkipModified = processItem["skipModified"].(bool)
			}
			if processItem["skipReverseShell"] != nil {
				rule.Processes.SkipReverseShell = processItem["skipReverseShell"].(bool)
			}
			if processItem["whitelist"] != nil {
				rule.Processes.Whitelist = processItem["whitelist"].([]string)
			}
		}
		if item["wildfireanalysis"] != nil {
			rule.WildFireAnalysis = item["wildfireanalysis"].([]string)
		}

		ans.Rules = append(ans.Rules, rule)
	}

	return ans
}

func savePolicy(d *schema.ResourceData, obj policyRuntimeContainer.Policy) {
	d.Set("_id", obj.PolicyId)
	d.Set("learningdisabled", obj.LearningDisabled)
	d.Set("rules", obj.Rules)

	// Rule.
	if len(obj.Rules) > 0 {
		rv := map[string]interface{}{
			"advancedprotection":       obj.Rules[0].AdvancedProtection,
			"cloudmetadataenforcement": obj.Rules[0].CloudMetadataEnforcement,
			"collections":              obj.Rules[0].Collections,
			"customrules":              obj.Rules[0].CustomRules,
			"disabled":                 obj.Rules[0].Disabled,
			"dns":                      obj.Rules[0].Dns,
			"filesystem":               obj.Rules[0].Filesystem,
			"kubernetesenforcement":    obj.Rules[0].KubernetesEnforcement,
			"modified":                 obj.Rules[0].Modified,
			"name":                     obj.Rules[0].Name,
			"network":                  obj.Rules[0].Network,
			"notes":                    obj.Rules[0].Notes,
			"owner":                    obj.Rules[0].Owner,
			"previousname":             obj.Rules[0].PreviousName,
			"processes":                obj.Rules[0].Processes,
			"wildfireanalysis":         obj.Rules[0].WildFireAnalysis,
		}

		if err := d.Set("rules", []interface{}{rv}); err != nil {
			log.Printf("[WARN] Error setting 'rules' for %q: %s", d.Id(), err)
		}
	}

}

func createPolicy(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	obj := parsePolicy(d, "")

	if err := policyRuntimeContainer.Create(client, obj); err != nil {
		return err
	}

	PollApiUntilSuccess(func() error {
		_, err := policyRuntimeContainer.Get(client)
		return err
	})

	pol, err := policyRuntimeContainer.Get(client)
	if err != nil {
		return err
	}

	d.SetId(pol.PolicyId)
	return readPolicy(d, meta)
}

func readPolicy(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)

	obj, err := policyRuntimeContainer.Get(client)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	savePolicy(d, obj)

	return nil
}

func updatePolicy(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()
	obj := parsePolicy(d, id)

	if err := policyRuntimeContainer.Update(client, obj); err != nil {
		return err
	}

	return readPolicy(d, meta)
}

func deletePolicy(d *schema.ResourceData, meta interface{}) error {
	/*	client := meta.(*pc.Client)
		id := d.Id()

		err := policy.Delete(client, id)
		if err != nil {
			if err != pc.ObjectNotFoundError {
				return err
			}
		}*/

	d.SetId("")
	return nil
}

func getListPort(listPortInterface interface{}) policy.ListPort {
	listPortItem := listPortInterface.(map[string]interface{})

	listPort := policy.ListPort{
		Deny:  listPortItem["deny"].(bool),
		End:   listPortItem["end"].(int),
		Start: listPortItem["start"].(int),
	}
	return listPort
}
