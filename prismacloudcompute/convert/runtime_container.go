package convert

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy"
)

func SchemaToRuntimeContainerRules(d *schema.ResourceData) ([]policy.RuntimeContainerRule, error) {
	parsedRules := make([]policy.RuntimeContainerRule, 0)
	if rules, ok := d.GetOk("rule"); ok {
		presentRules := rules.([]interface{})
		for _, val := range presentRules {
			presentRule := val.(map[string]interface{})
			parsedRule := policy.RuntimeContainerRule{}

			parsedRule.AdvancedProtection = presentRule["advanced_protection"].(bool)
			parsedRule.CloudMetadataEnforcement = presentRule["cloud_metadata_enforcement"].(bool)

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
					Allowed:    SchemaToStringSlice(presentDns["allowed"].([]interface{})),
					Denied:     SchemaToStringSlice(presentDns["denied"].([]interface{})),
					DenyEffect: presentDns["deny_effect"].(string),
				}
			} else {
				parsedRule.Dns = policy.RuntimeContainerDns{}
			}

			if presentRule["filesystem"].([]interface{})[0] != nil {
				presentFilesystem := presentRule["filesystem"].([]interface{})[0].(map[string]interface{})
				parsedRule.Filesystem = policy.RuntimeContainerFilesystem{
					Allowed:               SchemaToStringSlice(presentFilesystem["allowed"].([]interface{})),
					BackdoorFiles:         presentFilesystem["backdoor_files"].(bool),
					CheckNewFiles:         presentFilesystem["check_new_files"].(bool),
					Denied:                SchemaToStringSlice(presentFilesystem["denied"].([]interface{})),
					DenyEffect:            presentFilesystem["deny_effect"].(string),
					SkipEncryptedBinaries: presentFilesystem["skip_encrypted_binaries"].(bool),
					SuspiciousElfHeaders:  presentFilesystem["suspicious_elf_headers"].(bool),
				}
			} else {
				parsedRule.Filesystem = policy.RuntimeContainerFilesystem{}
			}

			parsedRule.KubernetesEnforcement = presentRule["kubernetes_enforcement"].(bool)
			parsedRule.Name = presentRule["name"].(string)

			if presentRule["network"].([]interface{})[0] != nil {
				presentNetwork := presentRule["network"].([]interface{})[0].(map[string]interface{})
				parsedRule.Network = policy.RuntimeContainerNetwork{
					AllowedListeningPorts: schemaToRuntimeContainerPorts(presentNetwork["allowed_listening_port"].([]interface{})),
					AllowedOutboundIps:    SchemaToStringSlice(presentNetwork["allowed_outbound_ips"].([]interface{})),
					AllowedOutboundPorts:  schemaToRuntimeContainerPorts(presentNetwork["allowed_outbound_port"].([]interface{})),
					DeniedListeningPorts:  schemaToRuntimeContainerPorts(presentNetwork["denied_listening_port"].([]interface{})),
					DeniedOutboundIps:     SchemaToStringSlice(presentNetwork["denied_outbound_ips"].([]interface{})),
					DeniedOutboundPorts:   schemaToRuntimeContainerPorts(presentNetwork["denied_outbound_port"].([]interface{})),
					DenyEffect:            presentNetwork["deny_effect"].(string),
					DetectPortScan:        presentNetwork["detect_port_scan"].(bool),
					SkipModifiedProcesses: presentNetwork["skip_modified_processes"].(bool),
					SkipRawSockets:        presentNetwork["skip_raw_sockets"].(bool),
				}
			} else {
				parsedRule.Network = policy.RuntimeContainerNetwork{}
			}

			parsedRule.Notes = presentRule["notes"].(string)

			if presentRule["processes"].([]interface{})[0] != nil {
				presentProcesses := presentRule["processes"].([]interface{})[0].(map[string]interface{})
				parsedRule.Processes = policy.RuntimeContainerProcesses{
					Allowed:              SchemaToStringSlice(presentProcesses["allowed"].([]interface{})),
					CheckCryptoMiners:    presentProcesses["check_crypto_miners"].(bool),
					CheckLateralMovement: presentProcesses["check_lateral_movement"].(bool),
					CheckParentChild:     presentProcesses["check_parent_child"].(bool),
					CheckSuidBinaries:    presentProcesses["check_suid_binaries"].(bool),
					Denied:               SchemaToStringSlice(presentProcesses["denied"].([]interface{})),
					DenyEffect:           presentProcesses["deny_effect"].(string),
					SkipModified:         presentProcesses["skip_modified"].(bool),
					SkipReverseShell:     presentProcesses["skip_reverse_shell"].(bool),
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
		m["advanced_protection"] = val.AdvancedProtection
		m["cloud_metadata_enforcement"] = val.CloudMetadataEnforcement
		m["collections"] = CollectionsToPolicySchema(val.Collections)
		m["custom_rule"] = runtimeContainerCustomRulesToSchema(val.CustomRules)
		m["disabled"] = val.Disabled
		m["dns"] = runtimeContainerDnsToSchema(val.Dns)
		m["filesystem"] = runtimeContainerFileystemToSchema(val.Filesystem)
		m["kubernetes_enforcement"] = val.KubernetesEnforcement
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
	m["allowed"] = in.Allowed
	m["denied"] = in.Denied
	m["deny_effect"] = in.DenyEffect
	ans = append(ans, m)
	return ans
}

func runtimeContainerFileystemToSchema(in policy.RuntimeContainerFilesystem) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["allowed"] = in.Allowed
	m["backdoor_files"] = in.BackdoorFiles
	m["check_new_files"] = in.CheckNewFiles
	m["denied"] = in.Denied
	m["deny_effect"] = in.DenyEffect
	m["skip_encrypted_binaries"] = in.SkipEncryptedBinaries
	m["suspicious_elf_headers"] = in.SuspiciousElfHeaders
	ans = append(ans, m)
	return ans
}

func runtimeContainerNetworkToSchema(in policy.RuntimeContainerNetwork) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["allowed_listening_port"] = runtimeContainerPortsToSchema(in.AllowedListeningPorts)
	m["allowed_outbound_ips"] = in.AllowedOutboundIps
	m["allowed_outbound_port"] = runtimeContainerPortsToSchema(in.AllowedOutboundPorts)
	m["denied_listening_port"] = runtimeContainerPortsToSchema(in.DeniedListeningPorts)
	m["denied_outbound_ips"] = in.DeniedOutboundIps
	m["denied_outbound_port"] = runtimeContainerPortsToSchema(in.DeniedOutboundPorts)
	m["deny_effect"] = in.DenyEffect
	m["detect_port_scan"] = in.DetectPortScan
	m["skip_modified_processes"] = in.SkipModifiedProcesses
	m["skip_raw_sockets"] = in.SkipRawSockets
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
	m["allowed"] = in.Allowed
	m["check_crypto_miners"] = in.CheckCryptoMiners
	m["check_lateral_movement"] = in.CheckLateralMovement
	m["check_parent_child"] = in.CheckParentChild
	m["check_suid_binaries"] = in.CheckSuidBinaries
	m["denied"] = in.Denied
	m["deny_effect"] = in.DenyEffect
	m["skip_modified"] = in.SkipModified
	m["skip_reverse_shell"] = in.SkipReverseShell
	ans = append(ans, m)
	return ans
}
