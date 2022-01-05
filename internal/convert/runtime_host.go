package convert

import (
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/policy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaToRuntimeHostRules(d *schema.ResourceData) ([]policy.RuntimeHostRule, error) {
	parsedRules := make([]policy.RuntimeHostRule, 0)
	if rules, ok := d.GetOk("rule"); ok {
		presentRules := rules.([]interface{})
		for _, val := range presentRules {
			presentRule := val.(map[string]interface{})
			parsedRule := policy.RuntimeHostRule{}

			if presentRule["antimalware"].([]interface{})[0] != nil {
				presentAntiMalware := presentRule["antimalware"].([]interface{})[0].(map[string]interface{})
				parsedAntiMalware := policy.RuntimeHostAntiMalware{}

				parsedAntiMalware.AllowedProcesses = SchemaToStringSlice(presentAntiMalware["allowed_processes"].([]interface{}))
				parsedAntiMalware.CryptoMiner = presentAntiMalware["crypto_miners"].(string)
				parsedAntiMalware.CustomFeed = presentAntiMalware["custom_feed"].(string)

				if presentAntiMalware["denied_processes"].([]interface{})[0] != nil {
					presentDeniedProcesses := presentAntiMalware["denied_processes"].([]interface{})[0].(map[string]interface{})
					parsedAntiMalware.DeniedProcesses = policy.RuntimeHostDeniedProcesses{
						Effect: presentDeniedProcesses["effect"].(string),
						Paths:  SchemaToStringSlice(presentDeniedProcesses["paths"].([]interface{})),
					}
				} else {
					parsedAntiMalware.DeniedProcesses = policy.RuntimeHostDeniedProcesses{}
				}

				parsedAntiMalware.DetectCompilerGeneratedBinary = presentAntiMalware["detect_compiler_generated_binary"].(bool)
				parsedAntiMalware.EncryptedBinaries = presentAntiMalware["encrypted_binaries"].(string)
				parsedAntiMalware.ExecutionFlowHijack = presentAntiMalware["execution_flow_hijack"].(string)
				parsedAntiMalware.IntelligenceFeed = presentAntiMalware["intelligence_feed"].(string)
				parsedAntiMalware.ReverseShell = presentAntiMalware["reverse_shell"].(string)
				parsedAntiMalware.ServiceUnknownOriginBinary = presentAntiMalware["service_unknown_origin_binary"].(string)
				parsedAntiMalware.SkipSshTracking = presentAntiMalware["skip_ssh_tracking"].(bool)
				parsedAntiMalware.SuspiciousElfHeaders = presentAntiMalware["suspicious_elf_headers"].(string)
				parsedAntiMalware.TempFsProcesses = presentAntiMalware["temp_filesystem_processes"].(string)
				parsedAntiMalware.UserUnknownOriginBinary = presentAntiMalware["user_unknown_origin_binary"].(string)
				parsedAntiMalware.WebShell = presentAntiMalware["webshell"].(string)
				parsedAntiMalware.WildFireAnalysis = presentAntiMalware["wildfire_analysis"].(string)

				parsedRule.AntiMalware = parsedAntiMalware

			} else {
				parsedRule.AntiMalware = policy.RuntimeHostAntiMalware{}
			}

			parsedRule.Collections = PolicySchemaToCollections(presentRule["collections"].([]interface{}))

			presentCustomRules := presentRule["custom_rule"].([]interface{})
			parsedCustomRules := make([]policy.RuntimeHostCustomRule, 0, len(presentCustomRules))
			for _, val := range presentCustomRules {
				presentCustomRule := val.(map[string]interface{})
				parsedCustomRules = append(parsedCustomRules, policy.RuntimeHostCustomRule{
					Action: presentCustomRule["action"].(string),
					Effect: presentCustomRule["effect"].(string),
					Id:     presentCustomRule["id"].(int),
				})
			}
			parsedRule.CustomRules = parsedCustomRules

			parsedRule.Disabled = presentRule["disabled"].(bool)

			if presentRule["dns"].([]interface{})[0] != nil {
				presentDns := presentRule["dns"].([]interface{})[0].(map[string]interface{})
				parsedRule.Dns = policy.RuntimeHostDns{
					Allowed:          SchemaToStringSlice(presentDns["allowed"].([]interface{})),
					Denied:           SchemaToStringSlice(presentDns["denied"].([]interface{})),
					DenyEffect:       presentDns["deny_effect"].(string),
					IntelligenceFeed: presentDns["intelligence_feed"].(string),
				}
			} else {
				parsedRule.Dns = policy.RuntimeHostDns{}
			}

			presentFileIntegrityRules := presentRule["file_integrity_rule"].([]interface{})
			parsedFileIntegrityRules := make([]policy.RuntimeHostFileIntegrityRule, 0, len(presentFileIntegrityRules))
			for _, val := range presentFileIntegrityRules {
				presentFileIntegrityRule := val.(map[string]interface{})
				parsedFileIntegrityRules = append(parsedFileIntegrityRules, policy.RuntimeHostFileIntegrityRule{
					AllowedProcesses: SchemaToStringSlice(presentFileIntegrityRule["allowed_processes"].([]interface{})),
					ExcludedFiles:    SchemaToStringSlice(presentFileIntegrityRule["excluded_files"].([]interface{})),
					Metadata:         presentFileIntegrityRule["metadata"].(bool),
					Path:             presentFileIntegrityRule["path"].(string),
					Read:             presentFileIntegrityRule["read"].(bool),
					Recursive:        presentFileIntegrityRule["recursive"].(bool),
					Write:            presentFileIntegrityRule["write"].(bool),
				})
			}
			parsedRule.FileIntegrityRules = parsedFileIntegrityRules

			if presentRule["activities"].([]interface{})[0] != nil {
				presentActivities := presentRule["activities"].([]interface{})[0].(map[string]interface{})
				parsedRule.Forensic = policy.RuntimeHostForensic{
					ActivitiesDisabled:       presentActivities["disabled"].(bool),
					DockerEnabled:            presentActivities["docker_enabled"].(bool),
					ReadonlyDockerEnabled:    presentActivities["readonly_docker_enabled"].(bool),
					ServiceActivitiesEnabled: presentActivities["service_activities_enabled"].(bool),
					SshdEnabled:              presentActivities["sshd_enabled"].(bool),
					SudoEnabled:              presentActivities["sudo_enabled"].(bool),
				}
			} else {
				parsedRule.Forensic = policy.RuntimeHostForensic{}
			}

			presentLogInspectionRules := presentRule["log_inspection_rule"].([]interface{})
			parsedLogInspectionRules := make([]policy.RuntimeHostLogInspectionRule, 0, len(presentLogInspectionRules))
			for _, val := range presentLogInspectionRules {
				presentLogInspectionRule := val.(map[string]interface{})
				parsedLogInspectionRules = append(parsedLogInspectionRules, policy.RuntimeHostLogInspectionRule{
					Path:  presentLogInspectionRule["path"].(string),
					Regex: SchemaToStringSlice(presentLogInspectionRule["regex"].([]interface{})),
				})
			}
			parsedRule.LogInspectionRules = parsedLogInspectionRules

			parsedRule.Name = presentRule["name"].(string)

			if presentRule["network"].([]interface{})[0] != nil {
				presentNetwork := presentRule["network"].([]interface{})[0].(map[string]interface{})
				parsedRule.Network = policy.RuntimeHostNetwork{
					AllowedOutboundIps:   SchemaToStringSlice(presentNetwork["allowed_outbound_ips"].([]interface{})),
					CustomFeed:           presentNetwork["custom_feed"].(string),
					DeniedListeningPorts: schemaToRuntimeHostPorts(presentNetwork["denied_listening_port"].([]interface{})),
					DeniedOutboundIps:    SchemaToStringSlice(presentNetwork["denied_outbound_ips"].([]interface{})),
					DeniedOutboundPorts:  schemaToRuntimeHostPorts(presentNetwork["denied_outbound_port"].([]interface{})),
					DenyEffect:           presentNetwork["deny_effect"].(string),
					IntelligenceFeed:     presentNetwork["intelligence_feed"].(string),
				}
			} else {
				parsedRule.Network = policy.RuntimeHostNetwork{}
			}

			parsedRule.Notes = presentRule["notes"].(string)

			parsedRules = append(parsedRules, parsedRule)
		}
	}
	return parsedRules, nil
}

func schemaToRuntimeHostPorts(in []interface{}) []policy.RuntimeHostPort {
	parsedPorts := make([]policy.RuntimeHostPort, 0, len(in))
	for _, val := range in {
		presentPort := val.(map[string]interface{})
		parsedPorts = append(parsedPorts, policy.RuntimeHostPort{
			Deny:  presentPort["deny"].(bool),
			End:   presentPort["end"].(int),
			Start: presentPort["start"].(int),
		})
	}
	return parsedPorts
}

func RuntimeHostRulesToSchema(in []policy.RuntimeHostRule) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["activities"] = runtimeHostActivitiesToSchema(val.Forensic)
		m["antimalware"] = runtimeHostAntiMalwareToSchema(val.AntiMalware)
		m["collections"] = CollectionsToPolicySchema(val.Collections)
		m["custom_rule"] = runtimeHostCustomRulesToSchema(val.CustomRules)
		m["disabled"] = val.Disabled
		m["dns"] = runtimeHostDnsToSchema(val.Dns)
		m["file_integrity_rule"] = runtimeHostFileIntegrityRulesToSchema(val.FileIntegrityRules)
		m["log_inspection_rule"] = runtimeHostLogInspectionRulesToSchema(val.LogInspectionRules)
		m["name"] = val.Name
		m["network"] = runtimeHostNetworkToSchema(val.Network)
		m["notes"] = val.Notes
		ans = append(ans, m)
	}
	return ans
}

func runtimeHostActivitiesToSchema(in policy.RuntimeHostForensic) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["disabled"] = in.ActivitiesDisabled
	m["docker_enabled"] = in.DockerEnabled
	m["readonly_docker_enabled"] = in.ReadonlyDockerEnabled
	m["service_activities_enabled"] = in.ServiceActivitiesEnabled
	m["sshd_enabled"] = in.SshdEnabled
	m["sudo_enabled"] = in.SudoEnabled
	ans = append(ans, m)
	return ans
}

func runtimeHostAntiMalwareToSchema(in policy.RuntimeHostAntiMalware) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["allowed_processes"] = in.AllowedProcesses
	m["crypto_miners"] = in.CryptoMiner
	m["custom_feed"] = in.CustomFeed
	m["denied_processes"] = runtimeHostDeniedProcessesToSchema(in.DeniedProcesses)
	m["detect_compiler_generated_binary"] = in.DetectCompilerGeneratedBinary
	m["encrypted_binaries"] = in.EncryptedBinaries
	m["execution_flow_hijack"] = in.ExecutionFlowHijack
	m["intelligence_feed"] = in.IntelligenceFeed
	m["reverse_shell"] = in.ReverseShell
	m["service_unknown_origin_binary"] = in.ServiceUnknownOriginBinary
	m["skip_ssh_tracking"] = in.SkipSshTracking
	m["suspicious_elf_headers"] = in.SuspiciousElfHeaders
	m["temp_filesystem_processes"] = in.TempFsProcesses
	m["user_unknown_origin_binary"] = in.UserUnknownOriginBinary
	m["webshell"] = in.WebShell
	m["wildfire_analysis"] = in.WildFireAnalysis
	ans = append(ans, m)
	return ans
}

func runtimeHostCustomRulesToSchema(in []policy.RuntimeHostCustomRule) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["action"] = val.Action
		m["effect"] = val.Effect
		m["id"] = val.Id
		ans = append(ans, m)
	}
	return ans
}

func runtimeHostDeniedProcessesToSchema(in policy.RuntimeHostDeniedProcesses) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["effect"] = in.Effect
	m["paths"] = in.Paths
	ans = append(ans, m)
	return ans
}

func runtimeHostDnsToSchema(in policy.RuntimeHostDns) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["allowed"] = in.Allowed
	m["denied"] = in.Denied
	m["deny_effect"] = in.DenyEffect
	m["intelligence_feed"] = in.IntelligenceFeed
	ans = append(ans, m)
	return ans
}

func runtimeHostFileIntegrityRulesToSchema(in []policy.RuntimeHostFileIntegrityRule) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["allowed_processes"] = val.AllowedProcesses
		m["excluded_files"] = val.ExcludedFiles
		m["metadata"] = val.Metadata
		m["path"] = val.Path
		m["read"] = val.Read
		m["recursive"] = val.Recursive
		m["write"] = val.Write
		ans = append(ans, m)
	}
	return ans
}

func runtimeHostLogInspectionRulesToSchema(in []policy.RuntimeHostLogInspectionRule) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["path"] = val.Path
		m["regex"] = val.Regex
		ans = append(ans, m)
	}
	return ans
}

func runtimeHostNetworkToSchema(in policy.RuntimeHostNetwork) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["allowed_outbound_ips"] = in.AllowedOutboundIps
	m["custom_feed"] = in.CustomFeed
	m["denied_listening_port"] = runtimeHostPortsToSchema(in.DeniedListeningPorts)
	m["denied_outbound_ips"] = in.DeniedOutboundIps
	m["denied_outbound_port"] = runtimeHostPortsToSchema(in.DeniedOutboundPorts)
	m["deny_effect"] = in.DenyEffect
	m["intelligence_feed"] = in.IntelligenceFeed
	ans = append(ans, m)
	return ans
}

func runtimeHostPortsToSchema(in []policy.RuntimeHostPort) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["deny"] = val.Deny
		m["end"] = val.End
		m["start"] = val.Start
		ans = append(ans, m)
	}
	return ans
}
