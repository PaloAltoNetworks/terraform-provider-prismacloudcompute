package prismacloudcompute

import (
	"github.com/paloaltonetworks/prisma-cloud-compute-go/collection"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy"
	"strconv"
	"encoding/json"
)

func parseRules(rules []interface{}) []policy.Rule {
	rulesList := make([]policy.Rule, 0, len(rules))
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
			if item["condition"] != nil {
				cond := item["condition"].(map[string]interface{})

				condition := policy.Condition{}

				if cond["vulnerabilities"] != nil {
					vulnString := cond["vulnerabilities"].(string)
					if vulnString != "" {
						var vulnArray []policy.Vulnerability
						if err := json.Unmarshal([]byte(vulnString), &vulnArray); err != nil {
        	panic(err)
        }
						for i := 0; i < len(vulnArray); i++ {
						vuln := vulnArray[i]
						condition.Vulnerabilities = append(condition.Vulnerabilities, vuln)
						}
					}
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
				if networkItem["detectPortScan"] != nil {
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

			rulesList = append(rulesList, rule)
		}
	}

	return rulesList
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
		coll.Prisma = collItem["prisma"].(interface{}).(bool)
	}
	if collItem["system"] != nil {
		coll.System = collItem["system"].(interface{}).(bool)
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

func getThreshold(thresholdInterface interface{}) policy.Threshold {
	thresholdItem := thresholdInterface.(map[string]interface{})

	threshold := policy.Threshold{}
	if thresholdItem["enabled"] != nil {
		enbl, err := strconv.ParseBool(thresholdItem["enabled"].(string))
		if err == nil {
			threshold.Enabled = enbl
		}
	}
	if thresholdItem["disabled"] != nil {
		disbl, err := strconv.ParseBool(thresholdItem["disabled"].(string))
		if err == nil {
			threshold.Disabled = disbl
		}
	}
	if thresholdItem["value"] != nil {
		val, err := strconv.Atoi(thresholdItem["value"].(string))
		if err == nil {
			threshold.Value = val
		}
	}
	return threshold
}
