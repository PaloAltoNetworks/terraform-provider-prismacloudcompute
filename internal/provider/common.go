package provider

const (
	policyTypeAdmission            = "admission"
	policyTypeComplianceCiImage    = "ciImagesCompliance"
	policyTypeComplianceContainer  = "containerCompliance"
	policyTypeComplianceHost       = "hostCompliance"
	policyTypeRuntimeContainer     = "containerRuntime"
	policyTypeRuntimeHost          = "hostRuntime"
	policyTypeVulnerabilityCiImage = "ciImagesVulnerability"
	policyTypeVulnerabilityHost    = "hostVulnerability"
	policyTypeVulnerabilityImage   = "containerVulnerability"
	policyTypeWaasContainer        = "containerAppFirewall"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
