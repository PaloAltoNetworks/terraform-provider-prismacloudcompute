package provider

const (
	policyTypeAdmission               = "admission"
	policyTypeComplianceCiImage       = "ciImagesCompliance"
	policyTypeComplianceCoderepo      = "codeRepoCompliance"
	policyTypeComplianceCiCoderepo    = "ciCodeRepoCompliance"
	policyTypeComplianceContainer     = "containerCompliance"
	policyTypeComplianceHost          = "hostCompliance"
	policyTypeRuntimeContainer        = "containerRuntime"
	policyTypeRuntimeHost             = "hostRuntime"
	policyTypeVulnerabilityCiCoderepo = "ciCodeRepoVulnerability"
	policyTypeVulnerabilityCiImage    = "ciImagesVulnerability"
	policyTypeVulnerabilityCoderepo   = "codeRepoVulnerability"
	policyTypeVulnerabilityHost       = "hostVulnerability"
	policyTypeVulnerabilityImage      = "containerVulnerability"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
