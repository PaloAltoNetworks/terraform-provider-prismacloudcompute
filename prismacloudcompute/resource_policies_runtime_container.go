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
				Description: "ID of the policy set.",
			},
			"learningdisabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If set to 'true', automatic behavioural learning is enabled.",
				Default:     true,
			},
			"rules": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Rules in the policies.",
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name of the rule.",
						},
						"collections": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of collections. Used to scope the rule.",
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
										Description: "List of labels",
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
										Optional:    true,
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
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', this collection originates from Prisma Cloud.",
									},
									"system": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true' this collection was created by the system (i.e., a non-user). Otherwise it was created by a real user.",
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
									"blacklist": {
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
									"whitelist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Allow-listed domain names (e.g., *.gmail.com, *.s3.amazon.com).",
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
							Description: "Represents restrictions or suppression for filesystem changes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"backdoorfiles": {
										Type:        schema.TypeBool,
										Optional:    true,
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
										Optional:    true,
										Description: "If set to 'true', Detects changes to binaries and certificates.",
									},
									"effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect that will be used in the runtime rule. Can be set to 'block', 'prevent', 'alert', or 'disable'.",
									},
									"skipencryptedbinaries": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', the encrypted binaries check will be skipped.",
									},
									"suspiciouselfheaders": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', enables malware detection based on suspicious ELF headers.",
									},
									"whitelist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of allowed filesystem paths.",
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
						"network": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Represents the restrictions and suppression for networking.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"blacklistips": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Deny-listed IP addresses.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"blacklistlisteningports": {
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
									"blacklistoutboundports": {
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
									"detectportscan": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', port scanning detection is enabled.",
									},
									"effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Effect used in the runtime rule. Can be set to: 'block', 'prevent', 'alert', or 'disable'.",
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
									"whitelistoutboundports": {
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
						"processes": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Represents restrictions or suppression for running processes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"blacklist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of processes to deny.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"blockallbinaries": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', blocks all processes except for the main process.",
									},
									"checkcryptominers": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', detect crypto miners.",
									},
									"checklateralmovement": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', enables detection of processes that can be used for lateral movement exploits.",
									},
									"checknewbinaries": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', binaries which don't belong to the original image are allowed to run.",
									},
									"checkparentchild": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', enables check for parent-child relationship when comparing spawned processes in the model.",
									},
									"checksuidbinaries": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', enables check for process elevanting privileges (SUID bit).",
									},
									"effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The effect to be used in the runtime rule. Can be set to 'block', 'prevent', 'alert', 'disable'.",
									},
									"skipmodified": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', trigger audits/incidents when a modified proc is spawned.",
									},
									"skipreverseshell": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', reverse shell detection is disabled.",
									},
									"whitelist": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Allow-list of processes.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"wildfireanalysis": {
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

				rule.Collections = append(rule.Collections, getCollection(collItem))
			}

		}
		if item["customrules"] != nil {
			custRules := item["customrules"].([]interface{})
			rule.CustomRules = make([]policy.CustomRule, 0, len(custRules))
			if len(custRules) > 0 {
				for i := 0; i < len(custRules); i++ {
					custRuleItem := custRules[i].(map[string]interface{})

					custRule := policy.CustomRule{
						Id:     custRuleItem["_id"].(int),
						Action: custRuleItem["action"].([]string),
						Effect: custRuleItem["effect"].(string),
					}
					rule.CustomRules = append(rule.CustomRules, custRule)
				}
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
				rule.Dns.Effect = dnsItem["effect"].(string)
			}
			if dnsItem["whitelist"] != nil {
				rule.Dns.Whitelist = dnsItem["whitelist"].([]string)
			}
		}
		if item["filesystem"] != nil {
			fileSysSet := item["filesystem"].(interface{})
			fileSysItem := fileSysSet.(map[string]interface{})

			rule.Filesystem = policy.Filesystem{}
			if fileSysItem["backdoorFiles"] != nil {
				rule.Filesystem.BackdoorFiles = fileSysItem["backdoorFiles"].(bool)
			}
			if fileSysItem["blacklist"] != nil {
				rule.Filesystem.Blacklist = fileSysItem["blacklist"].([]string)
			}
			if fileSysItem["checkNewFiles"] != nil {
				rule.Filesystem.CheckNewFiles = fileSysItem["checkNewFiles"].(bool)
			}
			if fileSysItem["effect"] != nil {
				rule.Filesystem.Effect = fileSysItem["effect"].(string)
			}
			if fileSysItem["skipEncryptedBinaries"] != nil {
				rule.Filesystem.SkipEncryptedBinaries = fileSysItem["skipEncryptedBinaries"].(bool)
			}
			if fileSysItem["suspiciousELFHeaders"] != nil {
				rule.Filesystem.SuspiciousELFHeaders = fileSysItem["suspiciousELFHeaders"].(bool)
			}
			if fileSysItem["whitelist"] != nil {
				rule.Filesystem.Whitelist = fileSysItem["whitelist"].([]string)
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
					for i := 0; i < len(blacklistListenPorts); i++ {
						rule.Network.BlacklistListeningPorts = append(rule.Network.BlacklistListeningPorts, getListPort(blacklistListenPorts[i]))
					}
				}
			}

			if networkItem["blacklistOutboundPorts"] != nil {
				blacklistOutPorts := networkItem["blacklistOutboundPorts"].([]interface{})
				rule.Network.BlacklistOutboundPorts = make([]policy.ListPort, 0, len(blacklistOutPorts))
				if len(blacklistOutPorts) > 0 {
					for i := 0; i < len(blacklistOutPorts); i++ {
						rule.Network.BlacklistOutboundPorts = append(rule.Network.BlacklistOutboundPorts, getListPort(blacklistOutPorts[i]))
					}
				}
			}
			if networkItem["blacklistOutboundPorts"] != nil {
				rule.Network.DetectPortScan = networkItem["detectPortScan"].(bool)
			}
			if networkItem["effect"] != nil {
				rule.Network.Effect = networkItem["effect"].(string)
			}
			if networkItem["skipModifiedProc"] != nil {
				rule.Network.SkipModifiedProc = networkItem["skipModifiedProc"].(bool)
			}
			if networkItem["skipRawSockets"] != nil {
				rule.Network.SkipRawSockets = networkItem["skipRawSockets"].(bool)
			}
			if networkItem["whitelistIPs"] != nil {
				rule.Network.WhitelistIPs = networkItem["whitelistIPs"].([]string)
			}

			if networkItem["whitelistListeningPorts"] != nil {
				whitelistListenPorts := networkItem["whitelistListeningPorts"].([]interface{})
				rule.Network.WhitelistListeningPorts = make([]policy.ListPort, 0, len(whitelistListenPorts))
				if len(whitelistListenPorts) > 0 {
					for i := 0; i < len(whitelistListenPorts); i++ {
						rule.Network.WhitelistListeningPorts = append(rule.Network.WhitelistListeningPorts, getListPort(whitelistListenPorts[i]))
					}
				}
			}

			if networkItem["whitelistOutboundPorts"] != nil {
				whitelistOutPorts := networkItem["whitelistOutboundPorts"].([]interface{})
				rule.Network.WhitelistOutboundPorts = make([]policy.ListPort, 0, len(whitelistOutPorts))
				if len(whitelistOutPorts) > 0 {
					for i := 0; i < len(whitelistOutPorts); i++ {
						rule.Network.WhitelistOutboundPorts = append(rule.Network.WhitelistOutboundPorts, getListPort(whitelistOutPorts[i]))
					}
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
				rule.Processes.Effect = processItem["effect"].(string)
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
			rule.WildFireAnalysis = item["wildfireanalysis"].(string)
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

func getCollection(collItem map[string]interface{}) collection.Collection {
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
	return coll
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
