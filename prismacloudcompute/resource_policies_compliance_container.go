package prismacloudcompute

import (
	"log"
	"time"

	pc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/collection"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy/policyComplianceContainer"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePoliciesComplianceContainer() *schema.Resource {
	return &schema.Resource{
		Create: createPolicyComplianceContainer,
		Read:   readPolicyComplianceContainer,
		Update: updatePolicyComplianceContainer,
		Delete: deletePolicyComplianceContainer,

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
			"policytype": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Type of policy. For example: 'docker', 'containerVulnerability', 'containerCompliance', etc.",
				Default:     true,
			},
			"rule": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of policy rules.",
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Action to take.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"alertthreshold": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "The compliance container policy alert threshold. Threshold values typically vary between 0 and 10 (non-inclusive).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', enables alerts.",
									},
									"disabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', suppresses alerts for all compliance containers.",
									},
									"value": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Minimum severity to trigger alerts. Supported values range from 0 to 9, where 0=off, 1=low, and 9=critical.",
									},
								},
							},
						},
						"allcompliance": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If set to 'true', reports the results of all (passed and failed) compliance checks.",
						},
						"auditallowed": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If set to 'true', Prisma Cloud audits successful transactions.",
						},
						"blockmsg": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Represents the block message in a policy.",
						},
						"blockthreshold": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "The compliance container policy block threshold. Threshold values typically vary between 0 and 10 (non-inclusive).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', enables blocking.",
									},
									"disabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to 'true', suppresses blocking for all compliance containers.",
									},
									"value": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Minimum severity to trigger blocking. Supported values range from 0 to 9, where 0=off, 1=low, and 9=critical.",
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
										Description: "A hex color code for a collection.",
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
										Description: "If set to 'true', this collection was created by the system (i.e., a non-user). Otherwise (false) a real user.",
									},
								},
							},
						},
						"condition": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Rule conditions. Conditions only apply for their respective policy type.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"compliance_check": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Block and scan severity-based vulnerabilities conditions.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"block": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If set to 'true', the effect is blocked.",
												},
												"id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Compliance container ID.",
												},
											},
										},
									},
								},
							},
						},
						"cverules": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of Common Vulnerability and Exposure (CVE) IDs classified for special handling/exceptions.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"description": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Free-form text for documenting the exception.",
									},
									"effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the relevant action for a compliance container. Can be set to 'ignore', 'alert', or 'block'.",
									},
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "CVE ID",
									},
									"expiration": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The compliance container expiration date.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"date": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The date the compliance container expires.",
												},
												"enabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If set to 'true', the grace period is enabled.",
												},
											},
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
						"effect": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The effect of evaluating the given policy. Can be set to 'allow', 'deny', 'block', or 'alert'.",
						},
						"gracedays": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Number of days to suppress the rule's block effect. Measured from date the compliance container was fixed. If there's no fix, measured from the date the compliance container was published.",
						},
						"group": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Applicable groups.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"license": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "The configuration of the compliance policy license.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"alertThreshold": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The license severity threshold to indicate whether to perform an alert action. Threshold values typically vary between 0 and 10 (non-inclusive).",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If set to 'true', the alert action is enabled.",
												},
												"disabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If set to 'true', suppresses alerts for all compliance containers.",
												},
												"value": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "The minimum severity score for which the alert action is enabled.",
												},
											},
										},
									},
									"blockThreshold": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The license severity threshold to indicate whether to perform a block action. Threshold values typically vary between 0 and 10 (non-inclusive).",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If set to 'true', the block action is enabled.",
												},
												"disabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If set to 'true', suppresses blocking for all compliance containers.",
												},
												"value": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "The minimum severity score for which the block action is enabled.",
												},
											},
										},
									},
									"critical": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The list of licenses with critical severity.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"high": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The list of licenses with high severity.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"low": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The list of licenses with low severity.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"medium": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The list of licenses with medium severity.",
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
						"notes": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Free-form text notes.",
						},
						"onlyfixed": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If set to 'true', applies rule only when vendor fixes are available.",
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
						"principal": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Applicable users.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"tags": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of tags classified for special handling/exceptions.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"description": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Free-form text for documenting the exception.",
									},
									"effect": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the relevant action for a compliance container. Can be set to 'ignore', 'alert', or 'block'.",
									},
									"expiration": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The compliance container expiration date.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"date": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Date of the compliance container expiration.",
												},
												"enabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If set to 'true', the grace period is enabled.",
												},
											},
										},
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Tag name.",
									},
								},
							},
						},
						"verbose": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If set to 'true', displays a detailed message when an operation is blocked.",
						},
					},
				},
			},
		},
	}
}

func parsePolicyComplianceContainer(d *schema.ResourceData, id string) policyComplianceContainer.Policy {
	ans := policyComplianceContainer.Policy{
		PolicyId: id,
	}
	if d.Get("policytype") != nil {
		ans.PolicyType = d.Get("policytype").(string)
	}

	rules := d.Get("rule").([]interface{})
	ans.Rules = make([]policy.Rule, 0, len(rules))
	if len(rules) > 0 {
		for i := 0; i < len(rules); i++ {
			item := rules[i].(map[string]interface{})

			rule := policy.Rule{}

			if item["alertthreshold"] != nil {
				thresholdInterface := item["alertthreshold"].(interface{})
				rule.AlertThreshold = getThreshold(thresholdInterface)
			}
			if item["blockthreshold"] != nil {
				thresholdInterface := item["blockthreshold"].(interface{})
				rule.BlockThreshold = getThreshold(thresholdInterface)
			}
			if item["collections"] != nil {
				colls := item["collections"].([]interface{})

				rule.Collections = make([]collection.Collection, 0, len(colls))
				if len(colls) > 0 {
					collItem := colls[0].(map[string]interface{})

					rule.Collections = append(rule.Collections, getCollection(collItem))
				}
			}
			if item["condition"] != nil {
				condition := policy.Condition{}
				// rule.condition is a list with guaranteed length 1, so grab first element and cast it
				cond := item["condition"].([]interface{})[0].(map[string]interface{})
				if cond["compliance_check"] != nil {
					compliance_checks := cond["compliance_check"].(*schema.Set).List()
					condition.Vulnerabilities = make([]policy.Vulnerability, 0, len(compliance_checks))

					for _, v := range compliance_checks {
						check := v.(map[string]interface{})
						vulnerability := policy.Vulnerability{}
						if check["block"] != nil {
							vulnerability.Block = check["block"].(bool)
						}
						if check["id"] != nil {
							vulnerability.Id = check["id"].(int)
						}
						condition.Vulnerabilities = append(condition.Vulnerabilities, vulnerability)
					}
				}
				rule.Condition = condition
			}
			if item["customrules"] != nil {
				custRules := item["customrules"].([]interface{})
				rule.CustomRules = make([]policy.CustomRule, 0, len(custRules))
				if len(custRules) > 0 {
					custRuleItem := custRules[0].(map[string]interface{})

					custRule := policy.CustomRule{
						Id:     custRuleItem["_id"].(int),
						Action: custRuleItem["action"].([]string),
						Effect: custRuleItem["effect"].(string),
					}
					rule.CustomRules = append(rule.CustomRules, custRule)
				}
			}
			if item["disabled"] != nil {
				rule.Disabled = item["disabled"].(bool)
			}
			if item["effect"] != nil {
				rule.Effect = item["effect"].(string)
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
	}

	return ans
}

func savePolicyComplianceContainer(d *schema.ResourceData, obj policyComplianceContainer.Policy) {
	d.Set("_id", obj.PolicyId)
	d.Set("policytype", obj.PolicyType)
	d.Set("rule", obj.Rules)

	// Rule.
	if len(obj.Rules) > 0 {
		rv := map[string]interface{}{
			"advancedprotection":       obj.Rules[0].AdvancedProtection,
			"cloudmetadataenforcement": obj.Rules[0].CloudMetadataEnforcement,
			"collections":              obj.Rules[0].Collections,
			"condition":                obj.Rules[0].Condition,
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

		if err := d.Set("rule", []interface{}{rv}); err != nil {
			log.Printf("[WARN] Error setting 'rule' for %q: %s", d.Id(), err)
		}
	}

}

func createPolicyComplianceContainer(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	obj := parsePolicyComplianceContainer(d, "")

	if err := policyComplianceContainer.Create(client, obj); err != nil {
		return err
	}

	PollApiUntilSuccess(func() error {
		_, err := policyComplianceContainer.Get(client)
		return err
	})

	pol, err := policyComplianceContainer.Get(client)
	if err != nil {
		return err
	}

	d.SetId(pol.PolicyId)
	return readPolicyComplianceContainer(d, meta)
}

func readPolicyComplianceContainer(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)

	obj, err := policyComplianceContainer.Get(client)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	savePolicyComplianceContainer(d, obj)

	return nil
}

func updatePolicyComplianceContainer(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()
	obj := parsePolicyComplianceContainer(d, id)

	if err := policyComplianceContainer.Update(client, obj); err != nil {
		return err
	}

	return readPolicyComplianceContainer(d, meta)
}

func deletePolicyComplianceContainer(d *schema.ResourceData, meta interface{}) error {
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
