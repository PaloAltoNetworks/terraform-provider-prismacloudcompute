package prismacloudcompute

import (
	"github.com/paloaltonetworks/prisma-cloud-compute-go/collection"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy"
)

const (
	policyTypeComplianceCiImage    = "ciImagesCompliance"
	policyTypeComplianceContainer  = "containerCompliance"
	policyTypeComplianceHost       = "hostCompliance"
	policyTypeRuntimeContainer     = "containerRuntime"
	policyTypeRuntimeHost          = "hostRuntime"
	policyTypeVulnerabilityCiImage = "ciImagesVulnerability"
	policyTypeVulnerabilityHost    = "hostVulnerability"
	policyTypeVulnerabilityImage   = "containerVulnerability"
)

func parseCollections(in []interface{}) []collection.Collection {
	parsedCollections := make([]collection.Collection, 0, len(in))
	for _, val := range in {
		parsedCollections = append(parsedCollections, collection.Collection{
			Name: val.(string),
		})
	}
	return parsedCollections
}

func flattenVulnerabilityHostAlertThreshold(in policy.VulnerabilityHostThreshold) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["disabled"] = in.Disabled
	m["value"] = in.Value
	ans = append(ans, m)
	return ans
}

func flattenVulnerabilityImageAlertThreshold(in policy.VulnerabilityImageThreshold) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["disabled"] = in.Disabled
	m["value"] = in.Value
	ans = append(ans, m)
	return ans
}

func flattenRuntimeHostAntiMalware(in policy.RuntimeHostAntiMalware) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["allowed_processes"] = in.AllowedProcesses
	m["crypto_miners"] = in.CryptoMiner
	m["custom_feed"] = in.CustomFeed
	m["denied_processes"] = flattenRuntimeHostDeniedProcesses(in.DeniedProcesses)
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

func flattenVulnerabilityImageBlockThreshold(in policy.VulnerabilityImageThreshold) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["enabled"] = in.Enabled
	m["value"] = in.Value
	ans = append(ans, m)
	return ans
}

func flattenCollections(in []collection.Collection) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		ans = append(ans, val.Name)
	}
	return ans
}

func flattenComplianceConditions(in policy.ComplianceConditions) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["compliance_check"] = flattenChecks(in.Checks)
	ans = append(ans, m)
	return ans
}

func flattenRuntimeContainerCustomRules(in []policy.RuntimeContainerCustomRule) []interface{} {
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

func flattenRuntimeHostCustomRules(in []policy.RuntimeHostCustomRule) []interface{} {
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

func flattenVulnerabilityHostCveRules(in []policy.VulnerabilityHostCveRule) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["description"] = val.Description
		m["effect"] = val.Effect
		m["expiration"] = flattenVulnerabilityHostExpiration(val.Expiration)
		m["id"] = val.Id
		ans = append(ans, m)
	}
	return ans
}

func flattenVulnerabilityImageCveRules(in []policy.VulnerabilityImageCveRule) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["description"] = val.Description
		m["effect"] = val.Effect
		m["expiration"] = flattenVulnerabilityImageExpiration(val.Expiration)
		m["id"] = val.Id
		ans = append(ans, m)
	}
	return ans
}

func flattenRuntimeHostDeniedProcesses(in policy.RuntimeHostDeniedProcesses) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["effect"] = in.Effect
	m["paths"] = in.Paths
	ans = append(ans, m)
	return ans
}

func flattenRuntimeContainerDns(in policy.RuntimeContainerDns) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["allowed"] = in.Allowed
	m["denied"] = in.Denied
	m["deny_effect"] = in.DenyEffect
	ans = append(ans, m)
	return ans
}

func flattenRuntimeHostDns(in policy.RuntimeHostDns) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["allowed"] = in.Allowed
	m["denied"] = in.Denied
	m["deny_effect"] = in.DenyEffect
	m["intelligence_feed"] = in.IntelligenceFeed
	ans = append(ans, m)
	return ans
}

func flattenVulnerabilityHostExpiration(in policy.VulnerabilityHostExpiration) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["date"] = in.Date
	m["enabled"] = in.Enabled
	ans = append(ans, m)
	return ans
}

func flattenVulnerabilityImageExpiration(in policy.VulnerabilityImageExpiration) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["date"] = in.Date
	m["enabled"] = in.Enabled
	ans = append(ans, m)
	return ans
}

func flattenRuntimeHostFileIntegrityRules(in []policy.RuntimeHostFileIntegrityRule) []interface{} {
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

func flattenRuntimeContainerFilesystem(in policy.RuntimeContainerFilesystem) []interface{} {
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

func flattenRuntimeHostForensic(in policy.RuntimeHostForensic) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["activities_disabled"] = in.ActivitiesDisabled
	m["docker_enabled"] = in.DockerEnabled
	m["readonly_docker_enabled"] = in.ReadonlyDockerEnabled
	m["service_activities_enabled"] = in.ServiceActivitiesEnabled
	m["sshd_enabled"] = in.SshdEnabled
	m["sudo_enabled"] = in.SudoEnabled
	ans = append(ans, m)
	return ans
}

func flattenRuntimeHostLogInspectionRules(in []policy.RuntimeHostLogInspectionRule) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["path"] = val.Path
		m["regex"] = val.Regex
		ans = append(ans, m)
	}
	return ans
}

func flattenRuntimeContainerNetwork(in policy.RuntimeContainerNetwork) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["allowed_listening_port"] = flattenRuntimeContainerPorts(in.AllowedListeningPorts)
	m["allowed_outbound_ips"] = in.AllowedOutboundIps
	m["allowed_outbound_port"] = flattenRuntimeContainerPorts(in.AllowedOutboundPorts)
	m["denied_listening_port"] = flattenRuntimeContainerPorts(in.DeniedListeningPorts)
	m["denied_outbound_ips"] = in.DeniedOutboundIps
	m["denied_outbound_port"] = flattenRuntimeContainerPorts(in.DeniedOutboundPorts)
	m["deny_effect"] = in.DenyEffect
	m["detect_port_scan"] = in.DetectPortScan
	m["skip_modified_processes"] = in.SkipModifiedProcesses
	m["skip_raw_sockets"] = in.SkipRawSockets
	ans = append(ans, m)
	return ans
}

func flattenRuntimeContainerPorts(in []policy.RuntimeContainerPort) []interface{} {
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

func flattenRuntimeHostNetwork(in policy.RuntimeHostNetwork) []interface{} {
	ans := make([]interface{}, 0, 1)
	m := make(map[string]interface{})
	m["allowed_outbound_ips"] = in.AllowedOutboundIps
	m["denied_listening_port"] = flattenRuntimeHostPorts(in.DeniedListeningPorts)
	m["denied_outbound_ips"] = in.DeniedOutboundIps
	m["denied_outbound_port"] = flattenRuntimeHostPorts(in.DeniedOutboundPorts)
	m["deny_effect"] = in.DenyEffect
	m["intelligence_feed"] = in.IntelligenceFeed
	ans = append(ans, m)
	return ans
}

func flattenRuntimeHostPorts(in []policy.RuntimeHostPort) []interface{} {
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

func flattenRuntimeContainerProcesses(in policy.RuntimeContainerProcesses) []interface{} {
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

func flattenVulnerabilityHostTagRules(in []policy.VulnerabilityHostTagRule) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["description"] = val.Description
		m["effect"] = val.Effect
		m["expiration"] = flattenVulnerabilityHostExpiration(val.Expiration)
		m["name"] = val.Name
		ans = append(ans, m)
	}
	return ans
}

func flattenVulnerabilityImageTagRules(in []policy.VulnerabilityImageTagRule) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["description"] = val.Description
		m["effect"] = val.Effect
		m["expiration"] = flattenVulnerabilityImageExpiration(val.Expiration)
		m["name"] = val.Name
		ans = append(ans, m)
	}
	return ans
}

func flattenChecks(in []policy.ComplianceCheck) []interface{} {
	ans := make([]interface{}, 0, len(in))
	for _, val := range in {
		m := make(map[string]interface{})
		m["block"] = val.Block
		m["id"] = val.Id
		ans = append(ans, m)
	}
	return ans
}
