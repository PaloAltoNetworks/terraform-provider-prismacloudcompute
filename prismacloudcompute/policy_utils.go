package prismacloudcompute

import (
	"encoding/json"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/collections"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policies"
)

const (
	policyTypeComplianceCIImages    = "ciImagesCompliance"
	policyTypeComplianceContainer   = "containerCompliance"
	policyTypeComplianceHost        = "hostCompliance"
	policyTypeRuntimeContainer      = "containerRuntime"
	policyTypeRuntimeHost           = "hostRuntime"
	policyTypeVulnerabilityCIImages = "ciImagesVulnerability"
	policyTypeVulnerabilityHost     = "hostVulnerability"
	policyTypeVulnerabilityImages   = "containerVulnerability"
)

func parsePolicy(rd *schema.ResourceData, policyID, policyType string) policies.Policy {
	policy := policies.Policy{
		PolicyId:   policyID,
		PolicyType: policyType,
		Rules:      parseRules(rd.Get("rule").([]interface{})),
	}

	if rd.Get("learningdisabled") != nil {
		policy.LearningDisabled = rd.Get("learningdisabled").(bool)
	}

	return policy
}

func parseRules(rules []interface{}) []policies.Rule {
	rulesList := make([]policies.Rule, 0, len(rules))
	if len(rules) > 0 {
		for i := 0; i < len(rules); i++ {
			item := rules[i].(map[string]interface{})

			rule := policies.Rule{}

			if item["alertthreshold"] != nil {
				thresholdInterface := item["alertthreshold"]
				rule.AlertThreshold = getThreshold(thresholdInterface)
			}
			if item["antimalware"] != nil {
				antiMalwareSet := item["antimalware"]
				antiMalwareItem := antiMalwareSet.(map[string]interface{})

				rule.AntiMalware = policies.AntiMalware{}
				if antiMalwareItem["allowedprocesses"] != nil {
					rule.AntiMalware.AllowedProcesses = antiMalwareItem["allowedprocesses"].([]string)
				}
				if antiMalwareItem["cryptominer"] != nil {
					rule.AntiMalware.CryptoMiner = antiMalwareItem["cryptominer"].(string)
				}
				if antiMalwareItem["customfeed"] != nil {
					rule.AntiMalware.CustomFeed = antiMalwareItem["customfeed"].(string)
				}
				if antiMalwareItem["deniedprocesses"] != nil {
					deniedProcessesString := antiMalwareItem["deniedprocesses"].(string)
					if deniedProcessesString != "" {
						var deniedProcesses policies.DeniedProcesses
						if err := json.Unmarshal([]byte(deniedProcessesString), &deniedProcesses); err != nil {
							panic(err)
						}
						rule.AntiMalware.DeniedProcesses = deniedProcesses
					}
				}
				if antiMalwareItem["detectcompilergeneratedbinary"] != nil {
					rule.AntiMalware.DetectCompilerGeneratedBinary = antiMalwareItem["detectcompilergeneratedbinary"].(bool)
				}
				if antiMalwareItem["encryptedbinaries"] != nil {
					rule.AntiMalware.EncryptedBinaries = antiMalwareItem["encryptedbinaries"].(string)
				}
				if antiMalwareItem["executionflowhijack"] != nil {
					rule.AntiMalware.ExecutionFlowHijack = antiMalwareItem["executionflowhijack"].(string)
				}
				if antiMalwareItem["intelligencefeed"] != nil {
					rule.AntiMalware.IntelligenceFeed = antiMalwareItem["intelligencefeed"].(string)
				}
				if antiMalwareItem["reverseshell"] != nil {
					rule.AntiMalware.ReverseShell = antiMalwareItem["reverseshell"].(string)
				}
				if antiMalwareItem["serviceunknownoriginbinary"] != nil {
					rule.AntiMalware.ServiceUnknownOriginBinary = antiMalwareItem["serviceunknownoriginbinary"].(string)
				}
				if antiMalwareItem["skipsshtracking"] != nil {
					rule.AntiMalware.SkipSSHTracking = antiMalwareItem["skipsshtracking"].(bool)
				}
				if antiMalwareItem["suspiciouselfheaders"] != nil {
					rule.AntiMalware.SuspiciousELFHeaders = antiMalwareItem["suspiciouselfheaders"].(string)
				}
				if antiMalwareItem["tempfsproc"] != nil {
					rule.AntiMalware.TempFSProc = antiMalwareItem["tempfsproc"].(string)
				}
				if antiMalwareItem["userunknownoriginbinary"] != nil {
					rule.AntiMalware.UserUnknownOriginBinary = antiMalwareItem["userunknownoriginbinary"].(string)
				}
				if antiMalwareItem["webshell"] != nil {
					rule.AntiMalware.WebShell = antiMalwareItem["webshell"].(string)
				}
				if antiMalwareItem["wildfireanalysis"] != nil {
					rule.AntiMalware.WildFireAnalysis = antiMalwareItem["wildfireanalysis"].(string)
				}
			}
			if item["blockthreshold"] != nil {
				thresholdInterface := item["blockthreshold"]
				rule.BlockThreshold = getThreshold(thresholdInterface)
			}

			if item["advancedprotection"] != nil {
				rule.AdvancedProtection = item["advancedprotection"].(bool)
			}
			if item["cloudmetadataenforcement"] != nil {
				rule.CloudMetadataEnforcement = item["cloudmetadataenforcement"].(bool)
			}

			if item["collections"] != nil {
				colls := parseStringArray(item["collections"].([]interface{}))
				for _, v := range colls {
					coll := collections.Collection{Name: v}
					rule.Collections = append(rule.Collections, coll)
				}
			}

			if item["conditions"] != nil &&
				len(item["conditions"].([]interface{})) > 0 &&
				item["conditions"].([]interface{})[0] != nil {
				presentCondition := item["conditions"].([]interface{})[0].(map[string]interface{})
				condition := policies.Condition{}

				if presentCondition["compliance_check"] != nil {
					complianceChecks := presentCondition["compliance_check"].(*schema.Set).List()
					for _, v := range complianceChecks {
						condition.Vulnerabilities = append(condition.Vulnerabilities, policies.Vulnerability{
							Block: v.(map[string]interface{})["block"].(bool),
							Id:    v.(map[string]interface{})["id"].(int),
						})
					}
					rule.Condition = condition
				}
			}
			if item["customrules"] != nil {
				custRules := item["customrules"].([]interface{})
				rule.CustomRules = make([]policies.CustomRule, 0, len(custRules))
				if len(custRules) > 0 {
					for i := 0; i < len(custRules); i++ {
						custRuleItem := custRules[i].(map[string]interface{})

						custRule := policies.CustomRule{
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
				dnsSet := item["dns"]
				dnsItem := dnsSet.(map[string]interface{})

				rule.Dns = policies.Dns{}
				if dnsItem["blacklist"] != nil {
					rule.Dns.Blacklist = dnsItem["blacklist"].([]string)
				}
				if dnsItem["effect"] != nil {
					rule.Dns.Effect = dnsItem["effect"].(string)
				}
				if dnsItem["whitelist"] != nil {
					rule.Dns.Whitelist = dnsItem["whitelist"].([]string)
				}
				if dnsItem["allow"] != nil {
					rule.Dns.Allow = dnsItem["allow"].([]string)
				}
				if dnsItem["deny"] != nil {
					rule.Dns.Deny = dnsItem["deny"].([]string)
				}
				if dnsItem["denylisteffect"] != nil {
					rule.Dns.DenyListEffect = dnsItem["denylisteffect"].(string)
				}
				if dnsItem["intelligencefeed"] != nil {
					rule.Dns.IntelligenceFeed = dnsItem["intelligencefeed"].(string)
				}
			}

			if item["effect"] != nil {
				rule.Effect = item["effect"].(string)
			}

			if item["filesystem"] != nil {
				fileSysSet := item["filesystem"]
				fileSysItem := fileSysSet.(map[string]interface{})

				rule.Filesystem = policies.Filesystem{}
				if fileSysItem["backdoorFiles"] != nil {
					rule.Filesystem.BackdoorFiles = fileSysItem["backdoorFiles"].(bool)
				}
				if fileSysItem["blacklist"] != nil {
					rule.Filesystem.Blacklist = fileSysItem["blacklist"].([]string)
				}
				if fileSysItem["checknewfiles"] != nil {
					rule.Filesystem.CheckNewFiles = fileSysItem["checknewfiles"].(bool)
				}
				if fileSysItem["effect"] != nil {
					rule.Filesystem.Effect = fileSysItem["effect"].(string)
				}
				if fileSysItem["skipencryptedbinaries"] != nil {
					rule.Filesystem.SkipEncryptedBinaries = fileSysItem["skipencryptedbinaries"].(bool)
				}
				if fileSysItem["suspiciouselfheaders"] != nil {
					rule.Filesystem.SuspiciousELFHeaders = fileSysItem["suspiciouselfheaders"].(bool)
				}
				if fileSysItem["whitelist"] != nil {
					rule.Filesystem.Whitelist = fileSysItem["whitelist"].([]string)
				}
			}
			if item["kubernetesenforcement"] != nil {
				rule.KubernetesEnforcement = item["kubernetesenforcement"].(bool)
			}
			if item["fileintegrityrules"] != nil {
				fileIntegrityRulesSet := item["fileintegrityrules"].([]interface{})
				if len(fileIntegrityRulesSet) > 0 {
					fileIntegrityRulesItem := fileIntegrityRulesSet[0].(map[string]interface{})

					rule.FileIntegrityRules = []policies.FileIntegrityRules{}
					if fileIntegrityRulesItem["dir"] != nil {
						rule.FileIntegrityRules[0].Dir = fileIntegrityRulesItem["dir"].(bool)
					}
					if fileIntegrityRulesItem["exclusions"] != nil {
						rule.FileIntegrityRules[0].Exclusions = fileIntegrityRulesItem["exclusions"].([]string)
					}
					if fileIntegrityRulesItem["metadata"] != nil {
						rule.FileIntegrityRules[0].Metadata = fileIntegrityRulesItem["metadata"].(bool)
					}
					if fileIntegrityRulesItem["path"] != nil {
						rule.FileIntegrityRules[0].Path = fileIntegrityRulesItem["path"].(string)
					}
					if fileIntegrityRulesItem["procwhitelist"] != nil {
						rule.FileIntegrityRules[0].ProcWhitelist = fileIntegrityRulesItem["procwhitelist"].([]string)
					}
					if fileIntegrityRulesItem["read"] != nil {
						rule.FileIntegrityRules[0].Read = fileIntegrityRulesItem["read"].(bool)
					}
					if fileIntegrityRulesItem["recursive"] != nil {
						rule.FileIntegrityRules[0].Recursive = fileIntegrityRulesItem["recursive"].(bool)
					}
					if fileIntegrityRulesItem["write"] != nil {
						rule.FileIntegrityRules[0].Write = fileIntegrityRulesItem["write"].(bool)
					}
				}
			}
			if item["forensic"] != nil {
				forensicSet := item["forensic"]
				forensicItem := forensicSet.(map[string]interface{})
				rule.Forensic = policies.Forensic{}
				if forensicItem["activitiesdisabled"] != nil {
					activitiesDisabled, err := strconv.ParseBool(forensicItem["activitiesdisabled"].(string))
					if err == nil {
						rule.Forensic.ActivitiesDisabled = activitiesDisabled
					}
				}
				if forensicItem["dockerenabled"] != nil {
					dockerEnabled, err := strconv.ParseBool(forensicItem["dockerenabled"].(string))
					if err == nil {
						rule.Forensic.DockerEnabled = dockerEnabled
					}
				}
				if forensicItem["readOnlydockerenabled"] != nil {
					readonlyDockerEnabled, err := strconv.ParseBool(forensicItem["readOnlydockerenabled"].(string))
					if err == nil {
						rule.Forensic.ReadonlyDockerEnabled = readonlyDockerEnabled
					}
				}
				if forensicItem["serviceactivitiesenabled"] != nil {
					serviceActivitiesEnabled, err := strconv.ParseBool(forensicItem["serviceactivitiesenabled"].(string))
					if err == nil {
						rule.Forensic.ServiceActivitiesEnabled = serviceActivitiesEnabled
					}
				}
				if forensicItem["sshdenabled"] != nil {
					sshdEnabled, err := strconv.ParseBool(forensicItem["sshdenabled"].(string))
					if err == nil {
						rule.Forensic.SshdEnabled = sshdEnabled
					}
				}
				if forensicItem["sudoenabled"] != nil {
					sudoEnabled, err := strconv.ParseBool(forensicItem["sudoenabled"].(string))
					if err == nil {
						rule.Forensic.SudoEnabled = sudoEnabled
					}
				}
			}
			if item["loginspectionrules"] != nil {
				logInspectionRulesSet := item["loginspectionrules"].([]interface{})
				if len(logInspectionRulesSet) > 0 {
					logInspectionRulesItem := logInspectionRulesSet[0].(map[string]interface{})
					if logInspectionRulesItem["path"] != nil {
						rule.LogInspectionRules[0].Path = logInspectionRulesItem["path"].(string)
					}
					if logInspectionRulesItem["regex"] != nil {
						rule.LogInspectionRules[0].Regex = logInspectionRulesItem["regex"].([]string)
					}
				}
			}
			if item["modified"] != nil {
				rule.Modified = item["modified"].(string)
			}
			if item["name"] != nil {
				rule.Name = item["name"].(string)
			}
			if item["network"] != nil {
				networkSet := item["network"]
				networkItem := networkSet.(map[string]interface{})
				rule.Network = policies.Network{}
				if networkItem["blacklistips"] != nil {
					rule.Network.BlacklistIPs = networkItem["blacklistips"].([]string)
				}

				if networkItem["blacklistlisteningports"] != nil {
					blacklistListenPorts := networkItem["blacklistlisteningports"].([]interface{})
					rule.Network.BlacklistListeningPorts = make([]policies.ListPort, 0, len(blacklistListenPorts))
					if len(blacklistListenPorts) > 0 {
						for i := 0; i < len(blacklistListenPorts); i++ {
							rule.Network.BlacklistListeningPorts = append(rule.Network.BlacklistListeningPorts, getListPort(blacklistListenPorts[i]))
						}
					}
				}

				if networkItem["blacklistoutboundports"] != nil {
					blacklistOutPorts := networkItem["blacklistoutboundports"].([]interface{})
					rule.Network.BlacklistOutboundPorts = make([]policies.ListPort, 0, len(blacklistOutPorts))
					if len(blacklistOutPorts) > 0 {
						for i := 0; i < len(blacklistOutPorts); i++ {
							rule.Network.BlacklistOutboundPorts = append(rule.Network.BlacklistOutboundPorts, getListPort(blacklistOutPorts[i]))
						}
					}
				}
				if networkItem["detectportscan"] != nil {
					rule.Network.DetectPortScan = networkItem["detectportscan"].(bool)
				}
				if networkItem["effect"] != nil {
					rule.Network.Effect = networkItem["effect"].(string)
				}
				if networkItem["skipmodifiedproc"] != nil {
					rule.Network.SkipModifiedProc = networkItem["skipmodifiedproc"].(bool)
				}
				if networkItem["skiprawsockets"] != nil {
					rule.Network.SkipRawSockets = networkItem["skiprawsockets"].(bool)
				}
				if networkItem["whitelistips"] != nil {
					rule.Network.WhitelistIPs = networkItem["whitelistips"].([]string)
				}

				if networkItem["whitelistlisteningports"] != nil {
					whitelistListenPorts := networkItem["whitelistlisteningports"].([]interface{})
					rule.Network.WhitelistListeningPorts = make([]policies.ListPort, 0, len(whitelistListenPorts))
					if len(whitelistListenPorts) > 0 {
						for i := 0; i < len(whitelistListenPorts); i++ {
							rule.Network.WhitelistListeningPorts = append(rule.Network.WhitelistListeningPorts, getListPort(whitelistListenPorts[i]))
						}
					}
				}

				if networkItem["whitelistoutboundports"] != nil {
					whitelistOutPorts := networkItem["whitelistoutboundports"].([]interface{})
					rule.Network.WhitelistOutboundPorts = make([]policies.ListPort, 0, len(whitelistOutPorts))
					if len(whitelistOutPorts) > 0 {
						for i := 0; i < len(whitelistOutPorts); i++ {
							rule.Network.WhitelistOutboundPorts = append(rule.Network.WhitelistOutboundPorts, getListPort(whitelistOutPorts[i]))
						}
					}
				}
				if networkItem["allowedoutboundips"] != nil {
					rule.Network.AllowedOutboundIPs = networkItem["allowedoutboundips"].([]string)
				}
				if networkItem["customfeed"] != nil {
					rule.Network.CustomFeed = networkItem["customfeed"].(string)
				}
				if networkItem["deniedlisteningports"] != nil {
					deniedListeningPorts := networkItem["deniedlisteningports"].([]interface{})
					rule.Network.DeniedListeningPorts = make([]policies.ListPort, 0, len(deniedListeningPorts))
					if len(deniedListeningPorts) > 0 {
						for i := 0; i < len(deniedListeningPorts); i++ {
							rule.Network.DeniedListeningPorts = append(rule.Network.DeniedListeningPorts, getListPort(deniedListeningPorts[i]))
						}
					}
				}
				if networkItem["deniedoutboundips"] != nil {
					rule.Network.DeniedOutboundIPs = networkItem["deniedoutboundips"].([]string)
				}
				if networkItem["deniedoutboundports"] != nil {
					deniedOutboundPorts := networkItem["deniedoutboundports"].([]interface{})
					rule.Network.DeniedOutboundPorts = make([]policies.ListPort, 0, len(deniedOutboundPorts))
					if len(deniedOutboundPorts) > 0 {
						for i := 0; i < len(deniedOutboundPorts); i++ {
							rule.Network.DeniedOutboundPorts = append(rule.Network.DeniedOutboundPorts, getListPort(deniedOutboundPorts[i]))
						}
					}
				}
				if networkItem["denylisteffect"] != nil {
					rule.Network.DenyListEffect = networkItem["denylisteffect"].(string)
				}
				if networkItem["intelligencefeed"] != nil {
					rule.Network.IntelligenceFeed = networkItem["intelligencefeed"].(string)
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
				processSet := item["processes"]
				processItem := processSet.(map[string]interface{})

				rule.Processes = policies.Processes{}

				if processItem["blacklist"] != nil {
					rule.Processes.Blacklist = processItem["blacklist"].([]string)
				}
				if processItem["blockallbinaries"] != nil {
					rule.Processes.BlockAllBinaries = processItem["blockallbinaries"].(bool)
				}
				if processItem["checkcryptominers"] != nil {
					rule.Processes.CheckCryptoMiners = processItem["checkcryptominers"].(bool)
				}
				if processItem["checklateralmovement"] != nil {
					rule.Processes.CheckLateralMovement = processItem["checklateralmovement"].(bool)
				}
				if processItem["checknewbinaries"] != nil {
					rule.Processes.CheckNewBinaries = processItem["checknewbinaries"].(bool)
				}
				if processItem["checkparentchild"] != nil {
					rule.Processes.CheckParentChild = processItem["checkparentchild"].(bool)
				}
				if processItem["checksuidbinaries"] != nil {
					rule.Processes.CheckSuidBinaries = processItem["checksuidbinaries"].(bool)
				}
				if processItem["effect"] != nil {
					rule.Processes.Effect = processItem["effect"].(string)
				}
				if processItem["skipmodified"] != nil {
					rule.Processes.SkipModified = processItem["skipmodified"].(bool)
				}
				if processItem["skipreverseshell"] != nil {
					rule.Processes.SkipReverseShell = processItem["skipreverseshell"].(bool)
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

func getListPort(listPortInterface interface{}) policies.ListPort {
	listPortItem := listPortInterface.(map[string]interface{})

	listPort := policies.ListPort{
		Deny:  listPortItem["deny"].(bool),
		End:   listPortItem["end"].(int),
		Start: listPortItem["start"].(int),
	}
	return listPort
}

func getThreshold(thresholdInterface interface{}) policies.Threshold {
	thresholdItem := thresholdInterface.(map[string]interface{})

	threshold := policies.Threshold{}
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
