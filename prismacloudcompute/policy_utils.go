package prismacloudcompute

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/collection"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy"
)

const (
	policyTypeComplianceCiImages    = "ciImagesCompliance"
	policyTypeComplianceContainer   = "containerCompliance"
	policyTypeComplianceHost        = "hostCompliance"
	policyTypeRuntimeContainer      = "containerRuntime"
	policyTypeRuntimeHost           = "hostRuntime"
	policyTypeVulnerabilityCiImages = "ciImagesVulnerability"
	policyTypeVulnerabilityHost     = "hostVulnerability"
	policyTypeVulnerabilityImages   = "containerVulnerability"
)

func parsePolicy(d *schema.ResourceData, policyId, policyType string) (*policy.Policy, error) {
	parsedPolicy := policy.Policy{
		PolicyId:   policyId,
		PolicyType: policyType,
	}

	if d.Get("learning_disabled") != nil {
		parsedPolicy.LearningDisabled = d.Get("learning_disabled").(bool)
	}

	if d.Get("rule") != nil {
		parsedRules, err := parseRules(d.Get("rule").([]interface{}))
		if err != nil {
			return nil, fmt.Errorf("error parsing rules: %s", err)
		}
		parsedPolicy.Rules = parsedRules
	} else {
		parsedPolicy.Rules = make([]policy.Rule, 0)
	}

	return &parsedPolicy, nil
}

func parseRules(rules []interface{}) ([]policy.Rule, error) {
	ruleSlice := make([]policy.Rule, 0)

	if len(rules) > 0 {
		for _, v := range rules {
			item := v.(map[string]interface{})
			rule := policy.Rule{}

			if item["advanced_protection"] != nil {
				rule.AdvancedProtection = item["advanced_protection"].(bool)
			}

			if item["alert_threshold"] != nil {
				parsedThreshold, err := parseThreshold(item["alert_threshold"].(map[string]interface{}))
				if err != nil {
					return nil, fmt.Errorf("error parsing threshold: %s", err)
				}
				rule.AlertThreshold = *parsedThreshold
			}

			if item["antimalware"] != nil {
				anitMalware := item["antimalware"].(map[string]interface{})
				if anitMalware["allowed_processes"] != nil {
					rule.AntiMalware.AllowedProcesses = anitMalware["allowed_processes"].([]string)
				}
				if anitMalware["crypto_miner"] != nil {
					rule.AntiMalware.CryptoMiner = anitMalware["crypto_miner"].(string)
				}
				if anitMalware["custom_feed"] != nil {
					rule.AntiMalware.CustomFeed = anitMalware["custom_feed"].(string)
				}
				if anitMalware["denied_processes"] != nil {
					deniedProcessesString := anitMalware["denied_processes"].(string)
					if deniedProcessesString != "" {
						var deniedProcesses policy.DeniedProcesses
						if err := json.Unmarshal([]byte(deniedProcessesString), &deniedProcesses); err != nil {
							panic(err)
						}
						rule.AntiMalware.DeniedProcesses = deniedProcesses
					}
				}
				if anitMalware["detect_compiler_generated_binary"] != nil {
					rule.AntiMalware.DetectCompilerGeneratedBinary = anitMalware["detect_compiler_generated_binary"].(bool)
				}
				if anitMalware["encrypted_binaries"] != nil {
					rule.AntiMalware.EncryptedBinaries = anitMalware["encrypted_binaries"].(string)
				}
				if anitMalware["execution_flow_hijack"] != nil {
					rule.AntiMalware.ExecutionFlowHijack = anitMalware["execution_flow_hijack"].(string)
				}
				if anitMalware["intelligence_feed"] != nil {
					rule.AntiMalware.IntelligenceFeed = anitMalware["intelligence_feed"].(string)
				}
				if anitMalware["reverse_shell"] != nil {
					rule.AntiMalware.ReverseShell = anitMalware["reverse_shell"].(string)
				}
				if anitMalware["service_unknown_origin_binary"] != nil {
					rule.AntiMalware.ServiceUnknownOriginBinary = anitMalware["service_unknown_origin_binary"].(string)
				}
				if anitMalware["skip_ssh_tracking"] != nil {
					rule.AntiMalware.SkipSshTracking = anitMalware["skip_ssh_tracking"].(bool)
				}
				if anitMalware["suspicious_elf_headers"] != nil {
					rule.AntiMalware.SuspiciousElfHeaders = anitMalware["suspicious_elf_headers"].(string)
				}
				if anitMalware["temp_filesystem_processes"] != nil {
					rule.AntiMalware.TempFsProc = anitMalware["temp_filesystem_processes"].(string)
				}
				if anitMalware["user_unknown_origin_binary"] != nil {
					rule.AntiMalware.UserUnknownOriginBinary = anitMalware["user_unknown_origin_binary"].(string)
				}
				if anitMalware["webshell"] != nil {
					rule.AntiMalware.WebShell = anitMalware["webshell"].(string)
				}
				if anitMalware["wildfire_analysis"] != nil {
					rule.AntiMalware.WildFireAnalysis = anitMalware["wildfire_analysis"].(string)
				}
			}

			if item["block_message"] != nil {
				rule.BlockMsg = item["block_message"].(string)
			}

			if item["block_threshold"] != nil {
				parsedThreshold, err := parseThreshold(item["block_threshold"].(map[string]interface{}))
				if err != nil {
					return nil, fmt.Errorf("error parsing threshold: %s", err)
				}
				rule.BlockThreshold = *parsedThreshold
			}

			if item["cloud_metadata_enforcement"] != nil {
				rule.CloudMetadataEnforcement = item["cloud_metadata_enforcement"].(bool)
			}

			if item["collections"] != nil {
				colls := parseStringArray(item["collections"].([]interface{}))
				for _, v := range colls {
					coll := collection.Collection{Name: v}
					rule.Collections = append(rule.Collections, coll)
				}
			}

			if item["conditions"] != nil &&
				len(item["conditions"].([]interface{})) > 0 &&
				item["conditions"].([]interface{})[0] != nil {
				presentCondition := item["conditions"].([]interface{})[0].(map[string]interface{})
				condition := policy.Condition{}

				if presentCondition["compliance_check"] != nil {
					complianceChecks := presentCondition["compliance_check"].(*schema.Set).List()
					for _, v := range complianceChecks {
						condition.Vulnerabilities = append(condition.Vulnerabilities, policy.Vulnerability{
							Block: v.(map[string]interface{})["block"].(bool),
							Id:    v.(map[string]interface{})["id"].(int),
						})
					}
					rule.Condition = condition
				}
			}

			if item["custom_rule"] != nil {
				custRules := item["custom_rule"].([]interface{})
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

			if item["cve_rule"] != nil && len(item["cve_rule"].([]interface{})) > 0 {
				cveRules := item["cve_rule"].([]interface{})
				for _, v := range cveRules {
					cveRule := policy.CveRule{}
					if v.(map[string]interface{})["description"] != nil {
						cveRule.Description = v.(map[string]interface{})["description"].(string)
					}
					if v.(map[string]interface{})["effect"] != nil {
						cveRule.Effect = v.(map[string]interface{})["effect"].(string)
					}
					cveRuleExpiration := policy.Expiration{}
					if v.(map[string]interface{})["expiration"] != nil {
						if v.(map[string]interface{})["expiration"].(map[string]interface{})["date"] != nil {
							cveRuleExpiration.Date = v.(map[string]interface{})["expiration"].(map[string]interface{})["date"].(string)
						}
						if v.(map[string]interface{})["expiration"].(map[string]interface{})["enabled"] != nil {
							cveRuleExpiration.Enabled, _ = strconv.ParseBool(v.(map[string]interface{})["expiration"].(map[string]interface{})["enabled"].(string))
						}
					}
					cveRule.Expiration = cveRuleExpiration
					if v.(map[string]interface{})["id"] != nil {
						cveRule.Id = v.(map[string]interface{})["id"].(string)
					}

					rule.CveRules = append(rule.CveRules, cveRule)
				}
			}

			if item["disabled"] != nil {
				rule.Disabled = item["disabled"].(bool)
			}

			if item["dns"] != nil {
				dns := item["dns"].(map[string]interface{})
				if dns["denylist"] != nil {
					rule.Dns.Blacklist = dns["denylist"].([]string)
				}
				if dns["effect"] != nil {
					rule.Dns.Effect = dns["effect"].(string)
				}
				if dns["allowlist"] != nil {
					rule.Dns.Whitelist = dns["allowlist"].([]string)
				}
				if dns["allow"] != nil {
					rule.Dns.Allow = dns["allow"].([]string)
				}
				if dns["deny"] != nil {
					rule.Dns.Deny = dns["deny"].([]string)
				}
				if dns["deny_effect"] != nil {
					rule.Dns.DenyListEffect = dns["deny_effect"].(string)
				}
				if dns["intelligence_feed"] != nil {
					rule.Dns.IntelligenceFeed = dns["intelligence_feed"].(string)
				}
			}

			if item["effect"] != nil {
				rule.Effect = item["effect"].(string)
			}

			if item["file_integrity_rule"] != nil {
				fileIntegrityRules := item["file_integrity_rule"].([]interface{})
				rule.FileIntegrityRules = make([]policy.FileIntegrityRule, 0, len(fileIntegrityRules))
				for _, v := range fileIntegrityRules {
					presentFileIntegrityRule := v.(map[string]interface{})
					fileIntegrityRule := policy.FileIntegrityRule{}
					if presentFileIntegrityRule["allowed_processes"] != nil {
						fileIntegrityRule.ProcWhitelist = presentFileIntegrityRule["allowed_processes"].([]string)
					}
					if presentFileIntegrityRule["dir"] != nil {
						fileIntegrityRule.Dir = presentFileIntegrityRule["dir"].(bool)
					}
					if presentFileIntegrityRule["exclusions"] != nil {
						fileIntegrityRule.Exclusions = presentFileIntegrityRule["exclusions"].([]string)
					}
					if presentFileIntegrityRule["metadata"] != nil {
						fileIntegrityRule.Metadata = presentFileIntegrityRule["metadata"].(bool)
					}
					if presentFileIntegrityRule["path"] != nil {
						fileIntegrityRule.Path = presentFileIntegrityRule["path"].(string)
					}
					if presentFileIntegrityRule["read"] != nil {
						fileIntegrityRule.Read = presentFileIntegrityRule["read"].(bool)
					}
					if presentFileIntegrityRule["recursive"] != nil {
						fileIntegrityRule.Recursive = presentFileIntegrityRule["recursive"].(bool)
					}
					if presentFileIntegrityRule["write"] != nil {
						fileIntegrityRule.Write = presentFileIntegrityRule["write"].(bool)
					}
					rule.FileIntegrityRules = append(rule.FileIntegrityRules, fileIntegrityRule)
				}
			}

			if item["filesystem"] != nil {
				filesystem := item["filesystem"].(map[string]interface{})
				if filesystem["allowlist"] != nil {
					rule.Filesystem.Whitelist = filesystem["allowlist"].([]string)
				}
				if filesystem["backdoor_files"] != nil {
					rule.Filesystem.BackdoorFiles = filesystem["backdoor_files"].(bool)
				}
				if filesystem["check_new_files"] != nil {
					rule.Filesystem.CheckNewFiles = filesystem["check_new_files"].(bool)
				}
				if filesystem["denylist"] != nil {
					rule.Filesystem.Blacklist = filesystem["denylist"].([]string)
				}
				if filesystem["effect"] != nil {
					rule.Filesystem.Effect = filesystem["effect"].(string)
				}
				if filesystem["skip_encrypted_binaries"] != nil {
					rule.Filesystem.SkipEncryptedBinaries = filesystem["skip_encrypted_binaries"].(bool)
				}
				if filesystem["suspicious_elf_headers"] != nil {
					rule.Filesystem.SuspiciousElfHeaders = filesystem["suspicious_elf_headers"].(bool)
				}
			}

			if item["forensic"] != nil {
				forensic := item["forensic"].(map[string]interface{})
				if forensic["activities_disabled"] != nil {
					activitiesDisabled, err := strconv.ParseBool(forensic["activities_disabled"].(string))
					if err == nil {
						rule.Forensic.ActivitiesDisabled = activitiesDisabled
					}
				}
				if forensic["docker_enabled"] != nil {
					dockerEnabled, err := strconv.ParseBool(forensic["docker_enabled"].(string))
					if err == nil {
						rule.Forensic.DockerEnabled = dockerEnabled
					}
				}
				if forensic["readOnlydockerenabled"] != nil {
					readonlyDockerEnabled, err := strconv.ParseBool(forensic["readOnlydockerenabled"].(string))
					if err == nil {
						rule.Forensic.ReadonlyDockerEnabled = readonlyDockerEnabled
					}
				}
				if forensic["service_activities_enabled"] != nil {
					serviceActivitiesEnabled, err := strconv.ParseBool(forensic["service_activities_enabled"].(string))
					if err == nil {
						rule.Forensic.ServiceActivitiesEnabled = serviceActivitiesEnabled
					}
				}
				if forensic["sshd_enabled"] != nil {
					sshdEnabled, err := strconv.ParseBool(forensic["sshd_enabled"].(string))
					if err == nil {
						rule.Forensic.SshdEnabled = sshdEnabled
					}
				}
				if forensic["sudo_enabled"] != nil {
					sudoEnabled, err := strconv.ParseBool(forensic["sudo_enabled"].(string))
					if err == nil {
						rule.Forensic.SudoEnabled = sudoEnabled
					}
				}
			}

			if item["grace_days"] != nil {
				rule.GraceDays = item["grace_days"].(int)
			}

			if item["kubernetes_enforcement"] != nil {
				rule.KubernetesEnforcement = item["kubernetes_enforcement"].(bool)
			}

			if item["log_inspection_rule"] != nil {
				logInspectionRules := item["log_inspection_rule"].([]interface{})
				rule.LogInspectionRules = make([]policy.LogInspectionRule, 0, len(logInspectionRules))
				for _, v := range logInspectionRules {
					presentLogInspectionRules := v.(map[string]interface{})
					logInspectionRule := policy.LogInspectionRule{}
					if presentLogInspectionRules["path"] != nil {
						logInspectionRule.Path = presentLogInspectionRules["path"].(string)
					}
					if presentLogInspectionRules["regex"] != nil {
						logInspectionRule.Regex = presentLogInspectionRules["regex"].([]string)
					}
					rule.LogInspectionRules = append(rule.LogInspectionRules, logInspectionRule)
				}
			}

			if item["name"] != nil {
				rule.Name = item["name"].(string)
			}

			if item["network"] != nil {
				network := item["network"].(map[string]interface{})
				if network["allowed_outbound_ips"] != nil {
					rule.Network.WhitelistIps = network["allowed_outbound_ips"].([]string)
					rule.Network.AllowedOutboundIps = network["allowed_outbound_ips"].([]string)
				}
				if network["allowed_listening_port"] != nil {
					allowedListeningPorts := network["allowed_listening_port"].([]interface{})
					rule.Network.WhitelistListeningPorts = make([]policy.ListPort, 0, len(allowedListeningPorts))
					for _, v := range allowedListeningPorts {
						allowedListeningPort := parseListPort(v.(map[string]interface{}))
						rule.Network.WhitelistListeningPorts = append(rule.Network.WhitelistListeningPorts, allowedListeningPort)
					}
				}
				if network["allowed_outbound_port"] != nil {
					allowedOutboundPorts := network["allowed_outbound_port"].([]interface{})
					rule.Network.WhitelistOutboundPorts = make([]policy.ListPort, 0, len(allowedOutboundPorts))
					for _, v := range allowedOutboundPorts {
						allowedOutboundPort := parseListPort(v.(map[string]interface{}))
						rule.Network.WhitelistOutboundPorts = append(rule.Network.WhitelistOutboundPorts, allowedOutboundPort)
					}
				}
				if network["custom_feed"] != nil {
					rule.Network.CustomFeed = network["custom_feed"].(string)
				}
				if network["deny_effect"] != nil {
					rule.Network.DenyListEffect = network["deny_effect"].(string)
				}
				if network["denied_listening_port"] != nil {
					deniedListeningPorts := network["denied_listening_port"].([]interface{})
					rule.Network.BlacklistListeningPorts = make([]policy.ListPort, 0, len(deniedListeningPorts))
					rule.Network.DeniedListeningPorts = make([]policy.ListPort, 0, len(deniedListeningPorts))
					for _, v := range deniedListeningPorts {
						deniedListeningPort := parseListPort(v.(map[string]interface{}))
						rule.Network.BlacklistListeningPorts = append(rule.Network.BlacklistListeningPorts, deniedListeningPort)
						rule.Network.DeniedListeningPorts = append(rule.Network.DeniedListeningPorts, deniedListeningPort)
					}
				}
				if network["denied_outbound_ips"] != nil {
					rule.Network.BlacklistIps = network["denied_outbound_ips"].([]string)
					rule.Network.DeniedOutboundIps = network["denied_outbound_ips"].([]string)
				}
				if network["denied_outbound_port"] != nil {
					deniedOutboundPorts := network["denied_outbound_port"].([]interface{})
					rule.Network.BlacklistOutboundPorts = make([]policy.ListPort, 0, len(deniedOutboundPorts))
					rule.Network.DeniedOutboundPorts = make([]policy.ListPort, 0, len(deniedOutboundPorts))
					for _, v := range deniedOutboundPorts {
						deniedOutboundPort := parseListPort(v.(map[string]interface{}))
						rule.Network.BlacklistOutboundPorts = append(rule.Network.BlacklistOutboundPorts, deniedOutboundPort)
						rule.Network.DeniedOutboundPorts = append(rule.Network.DeniedOutboundPorts, deniedOutboundPort)
					}
				}
				if network["detect_port_scan"] != nil {
					rule.Network.DetectPortScan = network["detect_port_scan"].(bool)
				}
				if network["effect"] != nil {
					rule.Network.Effect = network["effect"].(string)
				}
				if network["intelligence_feed"] != nil {
					rule.Network.IntelligenceFeed = network["intelligence_feed"].(string)
				}
				if network["skip_modified_processes"] != nil {
					rule.Network.SkipModifiedProc = network["skip_modified_processes"].(bool)
				}
				if network["skip_raw_sockets"] != nil {
					rule.Network.SkipRawSockets = network["skip_raw_sockets"].(bool)
				}
			}

			if item["notes"] != nil {
				rule.Notes = item["notes"].(string)
			}

			if item["only_fixed"] != nil {
				rule.OnlyFixed = item["only_fixed"].(bool)
			}

			if item["processes"] != nil {
				processes := item["processes"].(map[string]interface{})
				if processes["allowlist"] != nil {
					rule.Processes.Whitelist = processes["allowlist"].([]string)
				}
				if processes["block_all_processes"] != nil {
					rule.Processes.BlockAllBinaries = processes["block_all_processes"].(bool)
				}
				if processes["crypto_miners"] != nil {
					rule.Processes.CheckCryptoMiners = processes["crypto_miners"].(bool)
				}
				if processes["check_lateral_movement"] != nil {
					rule.Processes.CheckLateralMovement = processes["check_lateral_movement"].(bool)
				}
				if processes["check_new_binaries"] != nil {
					rule.Processes.CheckNewBinaries = processes["check_new_binaries"].(bool)
				}
				if processes["check_parent_child"] != nil {
					rule.Processes.CheckParentChild = processes["check_parent_child"].(bool)
				}
				if processes["check_suid_binaries"] != nil {
					rule.Processes.CheckSuidBinaries = processes["check_suid_binaries"].(bool)
				}
				if processes["denylist"] != nil {
					rule.Processes.Blacklist = processes["denylist"].([]string)
				}
				if processes["effect"] != nil {
					rule.Processes.Effect = processes["effect"].(string)
				}
				if processes["skip_modified"] != nil {
					rule.Processes.SkipModified = processes["skip_modified"].(bool)
				}
				if processes["skip_reverse_shell"] != nil {
					rule.Processes.SkipReverseShell = processes["skip_reverse_shell"].(bool)
				}
			}

			if item["tags"] != nil && len(item["tags"].([]interface{})) > 0 {
				tags := item["tags"].([]interface{})
				rule.Tags = make([]policy.Tag, 0, len(tags))
				for _, v := range tags {
					presentTag := v.(map[string]interface{})
					tag := policy.Tag{}
					if presentTag["description"] != nil {
						tag.Description = presentTag["description"].(string)
					}
					if presentTag["effect"] != nil {
						tag.Effect = presentTag["effect"].(string)
					}
					if presentTag["expiration"] != nil {
						presentTagExpiration := presentTag["expiration"].(map[string]interface{})
						tagExpiration := policy.Expiration{}
						if presentTagExpiration["date"] != nil {
							tagExpiration.Date = presentTagExpiration["date"].(string)
						}
						if presentTagExpiration["enabled"] != nil {
							tagExpiration.Enabled, _ = strconv.ParseBool(presentTagExpiration["enabled"].(string))
						}
						tag.Expiration = tagExpiration
					}
					if presentTag["name"] != nil {
						tag.Name = presentTag["name"].(string)
					}
					rule.Tags = append(rule.Tags, tag)
				}
			}

			if item["wildfire_analysis"] != nil {
				rule.WildFireAnalysis = item["wildfire_analysis"].(string)
			}

			ruleSlice = append(ruleSlice, rule)
		}
	}

	return ruleSlice, nil
}

func flattenAlertThreshold(t policy.Threshold) map[string]interface{} {
	ans := make(map[string]interface{})
	ans["disabled"] = strconv.FormatBool(t.Disabled)
	ans["value"] = strconv.Itoa(t.Value)
	return ans
}

func flattenAntiMalware(a policy.AntiMalware) map[string]interface{} {
	ans := make(map[string]interface{})
	ans["allowed_processes"] = a.AllowedProcesses
	ans["crypto_miner"] = a.CryptoMiner
	ans["custom_feed"] = a.CustomFeed
	ans["denied_processes"] = flattenDeniedProcesses(a.DeniedProcesses)
	ans["detect_compiler_generated_binary"] = a.DetectCompilerGeneratedBinary
	ans["encrypted_binaries"] = a.EncryptedBinaries
	ans["execution_flow_hijack"] = a.ExecutionFlowHijack
	ans["intelligence_feed"] = a.IntelligenceFeed
	ans["reverse_shell"] = a.ReverseShell
	ans["service_unknown_origin_binary"] = a.ServiceUnknownOriginBinary
	ans["skip_ssh_tracking"] = a.SkipSshTracking
	ans["suspicious_elf_headers"] = a.SuspiciousElfHeaders
	ans["temp_filesystem_processes"] = a.TempFsProc
	ans["user_unknown_origin_binary"] = a.UserUnknownOriginBinary
	ans["webshell"] = a.WebShell
	ans["wildfire_analysis"] = a.WildFireAnalysis
	return ans
}

func flattenBlockThreshold(t policy.Threshold) map[string]interface{} {
	ans := make(map[string]interface{})
	ans["enabled"] = strconv.FormatBool(t.Enabled)
	ans["value"] = strconv.Itoa(t.Value)
	return ans
}

func flattenCollections(c []collection.Collection) []interface{} {
	ans := make([]interface{}, 0, len(c))
	for _, val := range c {
		ans = append(ans, val.Name)
	}
	return ans
}

func flattenConditions(c policy.Condition) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["compliance_check"] = flattenVulnerabilities(c.Vulnerabilities)
	ans = append(ans, m)
	return ans
}

func flattenCustomRules(c []policy.CustomRule) []interface{} {
	ans := make([]interface{}, 0, len(c))
	for _, val := range c {
		m := make(map[string]interface{})
		m["action"] = val.Action
		m["effect"] = val.Effect
		m["id"] = val.Id
		ans = append(ans, m)
	}
	return ans
}

func flattenCveRules(c []policy.CveRule) []interface{} {
	ans := make([]interface{}, 0, len(c))
	for _, val := range c {
		m := make(map[string]interface{})
		m["description"] = val.Description
		m["effect"] = val.Effect
		m["expiration"] = flattenExpiration(val.Expiration)
		m["id"] = val.Id
		ans = append(ans, m)
	}
	return ans
}

func flattenDeniedProcesses(d policy.DeniedProcesses) map[string]interface{} {
	ans := make(map[string]interface{})
	ans["effect"] = d.Effect
	ans["paths"] = d.Paths
	return ans
}

func flattenDns(d policy.Dns) map[string]interface{} {
	ans := make(map[string]interface{})
	ans["allow"] = d.Allow
	ans["denylist"] = d.Blacklist
	ans["deny"] = d.Deny
	ans["deny_effect"] = d.DenyListEffect
	ans["effect"] = d.Effect
	ans["intelligence_feed"] = d.IntelligenceFeed
	ans["allowlist"] = d.Whitelist
	return ans
}

func flattenExpiration(e policy.Expiration) map[string]interface{} {
	ans := make(map[string]interface{})
	ans["date"] = e.Date
	ans["enabled"] = strconv.FormatBool(e.Enabled)
	return ans
}

func flattenFileIntegrityRules(f []policy.FileIntegrityRule) []interface{} {
	ans := make([]interface{}, 0, len(f))
	for _, val := range f {
		m := make(map[string]interface{})
		m["dir"] = val.Dir
		m["exclusions"] = val.Exclusions
		m["metadata"] = val.Metadata
		m["path"] = val.Path
		m["allowed_processes"] = val.ProcWhitelist
		m["read"] = val.Read
		m["recursive"] = val.Recursive
		m["write"] = val.Write
		ans = append(ans, m)
	}
	return ans
}

func flattenFilesystem(f policy.Filesystem) map[string]interface{} {
	ans := make(map[string]interface{})
	ans["backdoor_files"] = f.BackdoorFiles
	ans["denylist"] = f.Blacklist
	ans["check_new_files"] = f.CheckNewFiles
	ans["effect"] = f.Effect
	ans["skip_encrypted_binaries"] = f.SkipEncryptedBinaries
	ans["suspicious_elf_headers"] = f.SuspiciousElfHeaders
	ans["allowlist"] = f.Whitelist
	return ans
}

func flattenForensic(f policy.Forensic) map[string]interface{} {
	ans := make(map[string]interface{})
	ans["activities_disabled"] = f.ActivitiesDisabled
	ans["docker_enabled"] = f.DockerEnabled
	ans["readOnlydockerenabled"] = f.ReadonlyDockerEnabled
	ans["service_activities_enabled"] = f.ServiceActivitiesEnabled
	ans["sshd_enabled"] = f.SshdEnabled
	ans["sudo_enabled"] = f.SudoEnabled
	return ans
}

func flattenLogInspectionRules(l []policy.LogInspectionRule) []interface{} {
	ans := make([]interface{}, 0, len(l))
	for _, val := range l {
		m := make(map[string]interface{})
		m["path"] = val.Path
		m["regex"] = val.Regex
		ans = append(ans, m)
	}
	return ans
}

func flattenNetwork(n policy.Network) map[string]interface{} {
	ans := make(map[string]interface{})
	ans["date"] = n.AllowedOutboundIps
	ans["date"] = n.BlacklistIps
	ans["date"] = flattenPorts(n.BlacklistListeningPorts)
	ans["date"] = flattenPorts(n.BlacklistOutboundPorts)
	ans["date"] = n.CustomFeed
	ans["date"] = flattenPorts(n.DeniedListeningPorts)
	ans["date"] = n.DeniedOutboundIps
	ans["date"] = flattenPorts(n.DeniedOutboundPorts)
	ans["date"] = n.DenyListEffect
	ans["date"] = n.DetectPortScan
	ans["date"] = n.Effect
	ans["date"] = n.IntelligenceFeed
	ans["date"] = n.SkipModifiedProc
	ans["date"] = n.SkipRawSockets
	ans["date"] = n.WhitelistIps
	ans["date"] = flattenPorts(n.WhitelistListeningPorts)
	ans["date"] = flattenPorts(n.WhitelistOutboundPorts)
	return ans
}

func flattenPorts(p []policy.ListPort) []interface{} {
	ans := make([]interface{}, 0, len(p))
	for _, val := range p {
		m := make(map[string]interface{})
		m["deny"] = val.Deny
		m["end"] = val.End
		m["start"] = val.Start
		ans = append(ans, m)
	}
	return ans
}

func flattenProcesses(p policy.Processes) map[string]interface{} {
	ans := make(map[string]interface{})
	ans["denylist"] = p.Blacklist
	ans["block_all_processes"] = p.BlockAllBinaries
	ans["crypto_miners"] = p.CheckCryptoMiners
	ans["check_lateral_movement"] = p.CheckLateralMovement
	ans["check_new_binaries"] = p.CheckNewBinaries
	ans["check_parent_child"] = p.CheckParentChild
	ans["check_suid_binaries"] = p.CheckSuidBinaries
	ans["effect"] = p.Effect
	ans["skip_modified"] = p.SkipModified
	ans["date"] = p.SkipReverseShell
	ans["skip_reverse_shell"] = p.Whitelist
	return ans
}

func flattenTags(t []policy.Tag) []interface{} {
	ans := make([]interface{}, 0, len(t))
	for _, val := range t {
		m := make(map[string]interface{})
		m["description"] = val.Description
		m["effect"] = val.Effect
		m["expiration"] = flattenExpiration(val.Expiration)
		m["name"] = val.Name
		ans = append(ans, m)
	}
	return ans
}

func flattenVulnerabilities(v []policy.Vulnerability) *schema.Set {
	ans := schema.Set{
		F: schema.HashResource(&schema.Resource{
			Schema: map[string]*schema.Schema{
				"block": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Whether or not to block if this check is failed. Setting to 'false' will only alert on failure.",
				},
				"id": {
					Type:        schema.TypeInt,
					Optional:    true,
					Description: "Compliance check ID.",
				},
			},
		}),
	}
	for _, val := range v {
		m := make(map[string]interface{})
		m["block"] = val.Block
		m["id"] = val.Id
		ans.Add(m)
	}
	return &ans
}

func parseListPort(listPort map[string]interface{}) policy.ListPort {
	parsedListPort := policy.ListPort{}
	if listPort["deny"] != nil {
		parsedListPort.Deny = listPort["deny"].(bool)
	}
	if listPort["end"] != nil {
		parsedListPort.End = listPort["end"].(int)
	}
	if listPort["start"] != nil {
		parsedListPort.Start = listPort["start"].(int)
	}
	return parsedListPort
}

func parseThreshold(threshold map[string]interface{}) (*policy.Threshold, error) {
	parsedThreshold := policy.Threshold{}
	if threshold["enabled"] != nil {
		enabled, err := strconv.ParseBool(threshold["enabled"].(string))
		if err != nil {
			return nil, fmt.Errorf("error parsing threshold.enabled: %s", err)
		}
		parsedThreshold.Enabled = enabled
	}
	if threshold["disabled"] != nil {
		disabled, err := strconv.ParseBool(threshold["disabled"].(string))
		if err != nil {
			return nil, fmt.Errorf("error parsing threshold.disabled: %s", err)
		}
		parsedThreshold.Disabled = disabled
	}
	if threshold["value"] != nil {
		value, err := strconv.Atoi(threshold["value"].(string))
		if err != nil {
			return nil, fmt.Errorf("error parsing threshold.value: %s", err)
		}
		parsedThreshold.Value = value
	}
	return &parsedThreshold, nil
}
