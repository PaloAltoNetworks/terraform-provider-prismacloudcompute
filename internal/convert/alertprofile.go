package convert

import (
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/alertprofile"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Converts Alertprofile policy object to schema policy
func AlertProfilePoliciesToSchema(d *alertprofile.Policy) interface{} {

	alertTriggerPolicies := make(map[string]interface{})

	if d.Admission.Enabled {
		alertTriggerPolicies["admission"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.Admission.Enabled,
				"all_rules": d.Admission.Allrules,
				"rules":     d.Admission.Rules,
			},
		}
	}

	if d.AgentlessAppFirewall.Enabled {
		alertTriggerPolicies["agentless_app_firewall"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.AgentlessAppFirewall.Enabled,
				"all_rules": d.AgentlessAppFirewall.Allrules,
				"rules":     d.AgentlessAppFirewall.Rules,
			},
		}
	}

	if d.AppEmbeddedAppFirewall.Enabled {
		alertTriggerPolicies["app_embedded_app_firewall"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.AppEmbeddedAppFirewall.Enabled,
				"all_rules": d.AppEmbeddedAppFirewall.Allrules,
				"rules":     d.AppEmbeddedAppFirewall.Rules,
			},
		}
	}

	if d.AppEmbeddedRuntime.Enabled {
		alertTriggerPolicies["app_embedded_runtime"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.AppEmbeddedRuntime.Enabled,
				"all_rules": d.AppEmbeddedRuntime.Allrules,
				"rules":     d.AppEmbeddedRuntime.Rules,
			},
		}
	}

	if d.CloudDiscovery.Enabled {
		alertTriggerPolicies["cloud_discovery"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.CloudDiscovery.Enabled,
				"all_rules": d.CloudDiscovery.Allrules,
				"rules":     d.CloudDiscovery.Rules,
			},
		}
	}

	if d.CodeRepoVulnerability.Enabled {
		alertTriggerPolicies["code_repo_vulnerability"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.CodeRepoVulnerability.Enabled,
				"all_rules": d.CodeRepoVulnerability.Allrules,
				"rules":     d.CodeRepoVulnerability.Rules,
			},
		}
	}

	if d.ContainerAppFirewall.Enabled {
		alertTriggerPolicies["container_app_firewall"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.ContainerAppFirewall.Enabled,
				"all_rules": d.ContainerAppFirewall.Allrules,
				"rules":     d.ContainerAppFirewall.Rules,
			},
		}
	}

	if d.ContainerCompliance.Enabled {
		alertTriggerPolicies["container_compliance"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.ContainerCompliance.Enabled,
				"all_rules": d.ContainerCompliance.Allrules,
				"rules":     d.ContainerCompliance.Rules,
			},
		}
	}

	if d.ContainerComplianceScan.Enabled {
		alertTriggerPolicies["container_compliance_scan"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.ContainerComplianceScan.Enabled,
				"all_rules": d.ContainerComplianceScan.Allrules,
				"rules":     d.ContainerComplianceScan.Rules,
			},
		}
	}

	if d.ContainerRuntime.Enabled {
		alertTriggerPolicies["container_runtime"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.ContainerRuntime.Enabled,
				"all_rules": d.ContainerRuntime.Allrules,
				"rules":     d.ContainerRuntime.Rules,
			},
		}
	}

	if d.ContainerVulnerability.Enabled {
		alertTriggerPolicies["container_vulnerability"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.ContainerVulnerability.Enabled,
				"all_rules": d.ContainerVulnerability.Allrules,
				"rules":     d.ContainerVulnerability.Rules,
			},
		}
	}

	if d.Defender.Enabled {
		alertTriggerPolicies["defender"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.Defender.Enabled,
				"all_rules": d.Defender.Allrules,
				"rules":     d.Defender.Rules,
			},
		}
	}

	if d.Docker.Enabled {
		alertTriggerPolicies["docker"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.Docker.Enabled,
				"all_rules": d.Docker.Allrules,
				"rules":     d.Docker.Rules,
			},
		}
	}

	if d.HostAppFirewall.Enabled {
		alertTriggerPolicies["host_app_firewall"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.HostAppFirewall.Enabled,
				"all_rules": d.HostAppFirewall.Allrules,
				"rules":     d.HostAppFirewall.Rules,
			},
		}
	}

	if d.HostCompliance.Enabled {
		alertTriggerPolicies["host_compliance"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.HostCompliance.Enabled,
				"all_rules": d.HostCompliance.Allrules,
				"rules":     d.HostCompliance.Rules,
			},
		}
	}

	if d.HostComplianceScan.Enabled {
		alertTriggerPolicies["host_compliance_scan"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.HostComplianceScan.Enabled,
				"all_rules": d.HostComplianceScan.Allrules,
				"rules":     d.HostComplianceScan.Rules,
			},
		}
	}

	if d.HostRuntime.Enabled {
		alertTriggerPolicies["host_runtime"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.HostRuntime.Enabled,
				"all_rules": d.HostRuntime.Allrules,
				"rules":     d.HostRuntime.Rules,
			},
		}
	}

	if d.HostVulnerability.Enabled {
		alertTriggerPolicies["host_vulnerability"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.HostVulnerability.Enabled,
				"all_rules": d.HostVulnerability.Allrules,
				"rules":     d.HostVulnerability.Rules,
			},
		}
	}

	if d.Incident.Enabled {
		alertTriggerPolicies["incident"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.Incident.Enabled,
				"all_rules": d.Incident.Allrules,
				"rules":     d.Incident.Rules,
			},
		}
	}

	if d.KubernetesAudit.Enabled {
		alertTriggerPolicies["kubernetes_audit"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.KubernetesAudit.Enabled,
				"all_rules": d.KubernetesAudit.Allrules,
				"rules":     d.KubernetesAudit.Rules,
			},
		}
	}

	if d.NetworkFirewall.Enabled {
		alertTriggerPolicies["network_firewall"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.NetworkFirewall.Enabled,
				"all_rules": d.NetworkFirewall.Allrules,
				"rules":     d.NetworkFirewall.Rules,
			},
		}
	}

	if d.RegistryVulnerability.Enabled {
		alertTriggerPolicies["registry_vulnerability"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.RegistryVulnerability.Enabled,
				"all_rules": d.RegistryVulnerability.Allrules,
				"rules":     d.RegistryVulnerability.Rules,
			},
		}
	}

	if d.ServerlessAppFirewall.Enabled {
		alertTriggerPolicies["serverless_app_firewall"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.ServerlessAppFirewall.Enabled,
				"all_rules": d.ServerlessAppFirewall.Allrules,
				"rules":     d.ServerlessAppFirewall.Rules,
			},
		}
	}

	if d.ServerlessRuntime.Enabled {
		alertTriggerPolicies["serverless_runtime"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.ServerlessRuntime.Enabled,
				"all_rules": d.ServerlessRuntime.Allrules,
				"rules":     d.ServerlessRuntime.Rules,
			},
		}
	}

	if d.VmCompliance.Enabled {
		alertTriggerPolicies["vm_compliance"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.VmCompliance.Enabled,
				"all_rules": d.VmCompliance.Allrules,
				"rules":     d.VmCompliance.Rules,
			},
		}
	}

	if d.WaasHealth.Enabled {
		alertTriggerPolicies["waas_health"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.WaasHealth.Enabled,
				"all_rules": d.WaasHealth.Allrules,
				"rules":     d.WaasHealth.Rules,
			},
		}
	}

	return alertTriggerPolicies
}

// Converts a alertprofile schema to a alertprofile object for SDK compatibility.
func SchemaToAlertprofile(d *schema.ResourceData) (alertprofile.AlertProfile, error) {

	parsedAlertProfile := alertprofile.AlertProfile{}

	if val, ok := d.GetOk("name"); ok {
		parsedAlertProfile.Name = val.(string)
	}

	if val, ok := d.GetOk("enable_immediate_vulnerabilities_alerts"); ok {
		parsedAlertProfile.VulnerabilityImmediateAlertsEnabled = val.(bool)
	}

	if val, ok := d.GetOk("owner"); ok {
		parsedAlertProfile.Owner = val.(string)
	}

	if wc, ok := d.GetOk("webhook"); ok {
		for _, val := range wc.([]interface{}) {
			parsedAlertProfile.Webhook.Url = val.(map[string]interface{})["url"].(string)
			parsedAlertProfile.Webhook.CredentialId = val.(map[string]interface{})["credential_id"].(string)
			parsedAlertProfile.Webhook.CaCert = val.(map[string]interface{})["custom_ca"].(string)
			parsedAlertProfile.Webhook.Json = val.(map[string]interface{})["custom_json"].(string)
		}
	}

	parsedAlertProfile.Webhook.Enabled = true // Currently only supporting webhook alert profiles

	if alertTriggers, ok := d.GetOk("policy"); ok {
		for _, alertTrigger := range alertTriggers.([]interface{}) {
			for _, cv := range alertTrigger.(map[string]interface{})["admission"].([]interface{}) {
				parsedAlertProfile.Policy.Admission.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.Admission.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.Admission.Rules = append(parsedAlertProfile.Policy.Admission.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["agentless_app_firewall"].([]interface{}) {
				parsedAlertProfile.Policy.AgentlessAppFirewall.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.AgentlessAppFirewall.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.AgentlessAppFirewall.Rules = append(parsedAlertProfile.Policy.AgentlessAppFirewall.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["app_embedded_app_firewall"].([]interface{}) {
				parsedAlertProfile.Policy.AppEmbeddedAppFirewall.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.AppEmbeddedAppFirewall.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.AppEmbeddedAppFirewall.Rules = append(parsedAlertProfile.Policy.AppEmbeddedAppFirewall.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["app_embedded_runtime"].([]interface{}) {
				parsedAlertProfile.Policy.AppEmbeddedRuntime.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.AppEmbeddedRuntime.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.AppEmbeddedRuntime.Rules = append(parsedAlertProfile.Policy.AppEmbeddedRuntime.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["cloud_discovery"].([]interface{}) {
				parsedAlertProfile.Policy.CloudDiscovery.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.CloudDiscovery.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.CloudDiscovery.Rules = append(parsedAlertProfile.Policy.CloudDiscovery.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["code_repo_vulnerability"].([]interface{}) {
				parsedAlertProfile.Policy.CodeRepoVulnerability.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.CodeRepoVulnerability.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.CodeRepoVulnerability.Rules = append(parsedAlertProfile.Policy.CodeRepoVulnerability.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["container_app_firewall"].([]interface{}) {
				parsedAlertProfile.Policy.ContainerAppFirewall.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.ContainerAppFirewall.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.ContainerAppFirewall.Rules = append(parsedAlertProfile.Policy.ContainerAppFirewall.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["container_compliance"].([]interface{}) {
				parsedAlertProfile.Policy.ContainerCompliance.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.ContainerCompliance.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.ContainerCompliance.Rules = append(parsedAlertProfile.Policy.ContainerCompliance.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["container_compliance_scan"].([]interface{}) {
				parsedAlertProfile.Policy.ContainerComplianceScan.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.ContainerComplianceScan.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.ContainerComplianceScan.Rules = append(parsedAlertProfile.Policy.ContainerComplianceScan.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["container_vulnerability"].([]interface{}) {
				parsedAlertProfile.Policy.ContainerVulnerability.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.ContainerVulnerability.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.ContainerVulnerability.Rules = append(parsedAlertProfile.Policy.ContainerVulnerability.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["defender"].([]interface{}) {
				parsedAlertProfile.Policy.Defender.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.Defender.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.Defender.Rules = append(parsedAlertProfile.Policy.Defender.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["docker"].([]interface{}) {
				parsedAlertProfile.Policy.Docker.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.Docker.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.Docker.Rules = append(parsedAlertProfile.Policy.Docker.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["host_app_firewall"].([]interface{}) {
				parsedAlertProfile.Policy.HostAppFirewall.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.HostAppFirewall.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.HostAppFirewall.Rules = append(parsedAlertProfile.Policy.HostAppFirewall.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["host_compliance"].([]interface{}) {
				parsedAlertProfile.Policy.HostCompliance.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.HostCompliance.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.HostCompliance.Rules = append(parsedAlertProfile.Policy.HostCompliance.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["host_compliance_scan"].([]interface{}) {
				parsedAlertProfile.Policy.HostComplianceScan.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.HostComplianceScan.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.HostComplianceScan.Rules = append(parsedAlertProfile.Policy.HostComplianceScan.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["host_runtime"].([]interface{}) {
				parsedAlertProfile.Policy.HostRuntime.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.HostRuntime.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.HostRuntime.Rules = append(parsedAlertProfile.Policy.HostRuntime.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["host_vulnerability"].([]interface{}) {
				parsedAlertProfile.Policy.HostVulnerability.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.HostVulnerability.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.HostVulnerability.Rules = append(parsedAlertProfile.Policy.HostVulnerability.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["incident"].([]interface{}) {
				parsedAlertProfile.Policy.Incident.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.Incident.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.Incident.Rules = append(parsedAlertProfile.Policy.Incident.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["kubernetes_audit"].([]interface{}) {
				parsedAlertProfile.Policy.KubernetesAudit.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.KubernetesAudit.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.KubernetesAudit.Rules = append(parsedAlertProfile.Policy.KubernetesAudit.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["network_firewall"].([]interface{}) {
				parsedAlertProfile.Policy.NetworkFirewall.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.NetworkFirewall.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.NetworkFirewall.Rules = append(parsedAlertProfile.Policy.NetworkFirewall.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["registry_vulnerability"].([]interface{}) {
				parsedAlertProfile.Policy.RegistryVulnerability.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.RegistryVulnerability.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.RegistryVulnerability.Rules = append(parsedAlertProfile.Policy.RegistryVulnerability.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["serverless_app_firewall"].([]interface{}) {
				parsedAlertProfile.Policy.ServerlessAppFirewall.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.ServerlessAppFirewall.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.ServerlessAppFirewall.Rules = append(parsedAlertProfile.Policy.ServerlessAppFirewall.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["serverless_runtime"].([]interface{}) {
				parsedAlertProfile.Policy.ServerlessRuntime.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.ServerlessRuntime.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.ServerlessRuntime.Rules = append(parsedAlertProfile.Policy.ServerlessRuntime.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["vm_compliance"].([]interface{}) {
				parsedAlertProfile.Policy.VmCompliance.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.VmCompliance.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.VmCompliance.Rules = append(parsedAlertProfile.Policy.VmCompliance.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["vm_vulnerability"].([]interface{}) {
				parsedAlertProfile.Policy.VmVulnerability.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.VmVulnerability.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.VmVulnerability.Rules = append(parsedAlertProfile.Policy.VmVulnerability.Rules, rule.(string))
				}
			}

			for _, cv := range alertTrigger.(map[string]interface{})["waas_health"].([]interface{}) {
				parsedAlertProfile.Policy.WaasHealth.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertProfile.Policy.WaasHealth.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertProfile.Policy.WaasHealth.Rules = append(parsedAlertProfile.Policy.WaasHealth.Rules, rule.(string))
				}
			}
		}
	}

	return parsedAlertProfile, nil
}
