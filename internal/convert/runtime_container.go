package convert

import (
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/policy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaToRuntimeContainerRules(d *schema.ResourceData) ([]policy.RuntimeContainerRule, error) {
	parsedRules := make([]policy.RuntimeContainerRule, 0)
	if rules, ok := d.GetOk("rule"); ok {
		presentRules := rules.([]interface{})
		for _, val := range presentRules {
			presentRule := val.(map[string]interface{})
			parsedRule := policy.RuntimeContainerRule{}

			parsedRule.AdvancedProtectionEffect = presentRule["advanced_protection_effect"].(string)
			parsedRule.CloudMetadataEnforcementEffect = presentRule["cloud_metadata_enforcement_effect"].(string)
			parsedRule.PreviousName = presentRule["previous_name"].(string)
			parsedRule.SkipExecSessions = presentRule["skip_exec_sessions"].(bool)

			parsedRule.Collections = PolicySchemaToCollections(presentRule["collections"].([]interface{}))

			presentCustomRules := presentRule["custom_rule"].([]interface{})
			parsedCustomRules := make([]policy.RuntimeContainerCustomRule, 0, len(presentCustomRules))
			for _, val := range presentCustomRules {
				presentCustomRule := val.(map[string]interface{})
				parsedCustomRules = append(parsedCustomRules, policy.RuntimeContainerCustomRule{
					Action: presentCustomRule["action"].(string),
					Effect: presentCustomRule["effect"].(string),
					Id:     presentCustomRule["id"].(int),
				})
			}
			parsedRule.CustomRules = parsedCustomRules

			parsedRule.Disabled = presentRule["disabled"].(bool)

			if presentRule["dns"].([]interface{})[0] != nil {
				presentDns := presentRule["dns"].([]interface{})[0].(map[string]interface{})
				parsedRule.Dns = policy.RuntimeContainerDns{
					DefaultEffect: presentDns["default_effect"].(string),
					Disabled:      presentDns["disabled"].(bool),
					DomainList:    schemaToRuntimeContainerDnsDomainList(presentDns["domain_list"].([]interface{})),
				}
			} else {
				parsedRule.Dns = policy.RuntimeContainerDns{}
			}

			if presentRule["filesystem"].([]interface{})[0] != nil {
				presentFilesystem := presentRule["filesystem"].([]interface{})[0].(map[string]interface{})
				parsedRule.Filesystem = policy.RuntimeContainerFilesystem{
					AllowedList:                SchemaToStringSlice(presentFilesystem["allowed_list"].([]interface{})),
					BackdoorFilesEffect:        presentFilesystem["backdoor_files_effect"].(string),
					DefaultEffect:              presentFilesystem["default_effect"].(string),
					DeniedList:                 schemaToRuntimeContainerDeniedList(presentFilesystem["denied_list"].([]interface{})),
					Disabled:                   presentFilesystem["disabled"].(bool),
					EncryptedBinariesEffect:    presentFilesystem["encrypted_binaries_effect"].(string),
					NewFilesEffect:             presentFilesystem["new_files_effect"].(string),
					SuspiciousElfHeadersEffect: presentFilesystem["suspicious_elf_headers_effect"].(string),
				}
			} else {
				parsedRule.Filesystem = policy.RuntimeContainerFilesystem{}
			}

			parsedRule.KubernetesEnforcementEffect = presentRule["kubernetes_enforcement_effect"].(string)
			parsedRule.Name = presentRule["name"].(string)

			if presentRule["network"].([]interface{})[0] != nil {
				presentNetwork := presentRule["network"].([]interface{})[0].(map[string]interface{})
				parsedRule.Network = policy.RuntimeContainerNetwork{
					AllowedIps:         SchemaToStringSlice(presentNetwork["allowed_ips"].([]interface{})),
					DefaultEffect:      presentNetwork["default_effect"].(string),
					DeniedIps:          SchemaToStringSlice(presentNetwork["denied_ips"].([]interface{})),
					DeniedIpsEffect:    presentNetwork["denied_ips_effect"].(string),
					Disabled:           presentNetwork["disabled"].(bool),
					ListeningPorts:     schemaToRuntimeContainerNetworkPorts(presentNetwork["listening_ports"].([]interface{})),
					ModifiedProcEffect: presentNetwork["modified_proc_effect"].(string),
					OutboundPorts:      schemaToRuntimeContainerNetworkPorts(presentNetwork["outbound_ports"].([]interface{})),
					PortScanEffect:     presentNetwork["port_scan_effect"].(string),
					RawSocketsEffect:   presentNetwork["raw_sockets_effect"].(string),
				}
			} else {
				parsedRule.Network = policy.RuntimeContainerNetwork{}
			}

			parsedRule.Notes = presentRule["notes"].(string)

			if presentRule["processes"].([]interface{})[0] != nil {
				presentProcesses := presentRule["processes"].([]interface{})[0].(map[string]interface{})
				parsedRule.Processes = policy.RuntimeContainerProcesses{
					ModifiedProcessEffect: presentProcesses["modified_process_effect"].(string),
					CryptoMinersEffect:    presentProcesses["crypto_miners_effect"].(string),
					LateralMovementEffect: presentProcesses["lateral_movement_effect"].(string),
					ReverseShellEffect:    presentProcesses["reverse_shell_effect"].(string),
					SuidBinariesEffect:    presentProcesses["suid_binaries_effect"].(string),
					DefaultEffect:         presentProcesses["default_effect"].(string),
					CheckParentChild:      presentProcesses["check_parent_child"].(bool),
					AllowedList:           SchemaToStringSlice(presentProcesses["allowed_list"].([]interface{})),
					Disabled:              presentProcesses["disabled"].(bool),
					DeniedList:            schemaToRuntimeContainerDeniedList(presentProcesses["denied_list"].([]interface{})),
				}
			} else {
				parsedRule.Processes = policy.RuntimeContainerProcesses{}
			}

			parsedRule.WildFireAnalysis = presentRule["wildfire_analysis"].(string)

			parsedRules = append(parsedRules, parsedRule)
		}
	}
	return parsedRules, nil
}

func schemaToRuntimeContainerNetworkPorts(in []interface{}) policy.RuntimeContainerNetworkPorts {
	parsedNetworkPorts := policy.RuntimeContainerNetworkPorts{}

	for _, val := range in {
		presentRule := val.(map[string]interface{})
		parsedNetworkPorts.Allowed = schemaToRuntimeContainerPorts(presentRule["allowed"].([]interface{}))
		parsedNetworkPorts.Denied = schemaToRuntimeContainerPorts(presentRule["denied"].([]interface{}))
		parsedNetworkPorts.Effect = presentRule["effect"].(string)
	}

	return parsedNetworkPorts
}

func schemaToRuntimeContainerDnsDomainList(in []interface{}) policy.RuntimeContainerDnsDomainList {
	parsedDomainList := policy.RuntimeContainerDnsDomainList{}

	for _, val := range in {
		presentRule := val.(map[string]interface{})
		parsedDomainList.Allowed = SchemaToStringSlice(presentRule["allowed"].([]interface{}))
		parsedDomainList.Denied = SchemaToStringSlice(presentRule["denied"].([]interface{}))
		parsedDomainList.Effect = presentRule["effect"].(string)
	}

	return parsedDomainList
}

func schemaToRuntimeContainerDeniedList(in []interface{}) policy.RuntimeContainerDeniedList {
	parsedDeniedList := policy.RuntimeContainerDeniedList{}

	for _, val := range in {
		presentRule := val.(map[string]interface{})
		parsedDeniedList.Effect = presentRule["effect"].(string)
		parsedDeniedList.Paths = SchemaToStringSlice(presentRule["paths"].([]interface{}))
	}

	return parsedDeniedList
}

func schemaToRuntimeContainerPorts(in []interface{}) []policy.RuntimeContainerPort {
	parsedPorts := make([]policy.RuntimeContainerPort, 0, len(in))
	for _, val := range in {
		presentPort := val.(map[string]interface{})
		parsedPorts = append(parsedPorts, policy.RuntimeContainerPort{
			Deny:  presentPort["deny"].(bool),
			End:   presentPort["end"].(int),
			Start: presentPort["start"].(int),
		})
	}
	return parsedPorts
}

func RuntimeContainerRulesToSchema(in []policy.RuntimeContainerRule) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["advanced_protection_effect"] = val.AdvancedProtectionEffect
		m["cloud_metadata_enforcement_effect"] = val.CloudMetadataEnforcementEffect
		m["previous_name"] = val.PreviousName
		m["skip_exec_sessions"] = val.SkipExecSessions
		m["collections"] = CollectionsToPolicySchema(val.Collections)
		m["custom_rule"] = runtimeContainerCustomRulesToSchema(val.CustomRules)
		m["disabled"] = val.Disabled
		m["dns"] = runtimeContainerDnsToSchema(val.Dns)
		m["filesystem"] = runtimeContainerFileystemToSchema(val.Filesystem)
		m["kubernetes_enforcement_effect"] = val.KubernetesEnforcementEffect
		m["name"] = val.Name
		m["network"] = runtimeContainerNetworkToSchema(val.Network)
		m["notes"] = val.Notes
		m["processes"] = runtimeContainerProcessesToSchema(val.Processes)
		m["wildfire_analysis"] = val.WildFireAnalysis
		ans = append(ans, m)
	}
	return ans
}

func runtimeContainerCustomRulesToSchema(in []policy.RuntimeContainerCustomRule) []interface{} {
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

func runtimeContainerDnsToSchema(in policy.RuntimeContainerDns) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["default_effect"] = in.DefaultEffect
	m["disabled"] = in.Disabled
	m["domain_list"] = runtimeContainerDnsDomainListToSchema(in.DomainList)
	ans = append(ans, m)
	return ans
}

func runtimeContainerFileystemToSchema(in policy.RuntimeContainerFilesystem) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["allowed_list"] = in.AllowedList
	m["backdoor_files_effect"] = in.BackdoorFilesEffect
	m["default_effect"] = in.DefaultEffect
	m["denied_list"] = runtimeContainerDeniedListToSchema(in.DeniedList)
	m["disabled"] = in.Disabled
	m["encrypted_binaries_effect"] = in.EncryptedBinariesEffect
	m["new_files_effect"] = in.NewFilesEffect
	m["suspicious_elf_headers_effect"] = in.SuspiciousElfHeadersEffect
	ans = append(ans, m)
	return ans
}

func runtimeContainerNetworkToSchema(in policy.RuntimeContainerNetwork) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["allowed_ips"] = in.AllowedIps
	m["default_effect"] = in.DefaultEffect
	m["denied_ips"] = in.DeniedIps
	m["denied_ips_effect"] = in.DeniedIpsEffect
	m["disabled"] = in.Disabled
	m["listening_ports"] = runtimeContainerNetworkPortsToSchema(in.ListeningPorts)
	m["modified_proc_effect"] = in.ModifiedProcEffect
	m["outbound_ports"] = runtimeContainerNetworkPortsToSchema(in.OutboundPorts)
	m["port_scan_effect"] = in.PortScanEffect
	m["raw_sockets_effect"] = in.RawSocketsEffect
	ans = append(ans, m)
	return ans
}

func runtimeContainerNetworkPortsToSchema(in policy.RuntimeContainerNetworkPorts) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["allowed"] = runtimeContainerPortsToSchema(in.Allowed)
	m["denied"] = runtimeContainerPortsToSchema(in.Denied)
	m["effect"] = in.Effect
	ans = append(ans, m)
	return ans
}

func runtimeContainerPortsToSchema(in []policy.RuntimeContainerPort) []interface{} {
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

func runtimeContainerProcessesToSchema(in policy.RuntimeContainerProcesses) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["modified_process_effect"] = in.ModifiedProcessEffect
	m["crypto_miners_effect"] = in.CryptoMinersEffect
	m["lateral_movement_effect"] = in.LateralMovementEffect
	m["reverse_shell_effect"] = in.ReverseShellEffect
	m["suid_binaries_effect"] = in.SuidBinariesEffect
	m["default_effect"] = in.DefaultEffect
	m["check_parent_child"] = in.CheckParentChild
	m["allowed_list"] = in.AllowedList
	m["disabled"] = in.Disabled
	m["denied_list"] = runtimeContainerDeniedListToSchema(in.DeniedList)
	ans = append(ans, m)
	return ans
}

func runtimeContainerDnsDomainListToSchema(in policy.RuntimeContainerDnsDomainList) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["allowed"] = in.Allowed
	m["denied"] = in.Denied
	m["effect"] = in.Effect
	ans = append(ans, m)
	return ans
}

func runtimeContainerDeniedListToSchema(in policy.RuntimeContainerDeniedList) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["effect"] = in.Effect
	m["paths"] = in.Paths
	ans = append(ans, m)
	return ans
}
