package convert

import (
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/alertprofile"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Converts Alertprofile policy object to schema policy
func AlertProfilePoliciesToSchema(d *alertprofile.Policy) interface{} {

	alertTriggerPolicies := make(map[string]interface{})

	if d.Docker.Enabled {
		alertTriggerPolicies["access"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.Docker.Enabled,
				"all_rules": d.Docker.Allrules,
				"rules":     d.Docker.Rules,
			},
		}
	}

	if d.Admission.Enabled {
		alertTriggerPolicies["admission"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.Admission.Enabled,
				"all_rules": d.Admission.Allrules,
				"rules":     d.Admission.Rules,
			},
		}
	}

	if d.AppEmbeddedRuntime.Enabled {
		alertTriggerPolicies["app_embedded_defender_runtime"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.AppEmbeddedRuntime.Enabled,
				"all_rules": d.AppEmbeddedRuntime.Allrules,
				"rules":     d.AppEmbeddedRuntime.Rules,
			},
		}
	}

	if d.NetworkFirewall.Enabled {
		alertTriggerPolicies["cloud_native_network_firewall"] = []interface{}{
			map[string]interface{}{
				"enabled": d.NetworkFirewall.Enabled,
			},
		}
	}

	if d.ContainerComplianceScan.Enabled {
		alertTriggerPolicies["container_and_image_compliance"] = []interface{}{
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

	if d.Defender.Enabled {
		alertTriggerPolicies["defender_health"] = []interface{}{
			map[string]interface{}{
				"enabled": d.Defender.Enabled,
			},
		}
	}

	if d.HostComplianceScan.Enabled {
		alertTriggerPolicies["host_compliance"] = []interface{}{
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
		alertTriggerPolicies["host_vulnerabilities"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.HostVulnerability.Enabled,
				"all_rules": d.HostVulnerability.Allrules,
				"rules":     d.HostVulnerability.Rules,
			},
		}
	}

	if d.ContainerVulnerability.Enabled {
		alertTriggerPolicies["image_vulnerabilities"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.ContainerVulnerability.Enabled,
				"all_rules": d.ContainerVulnerability.Allrules,
				"rules":     d.ContainerVulnerability.Rules,
			},
		}
	}

	if d.Incident.Enabled {
		alertTriggerPolicies["incidents"] = []interface{}{
			map[string]interface{}{
				"enabled": d.Incident.Enabled,
			},
		}
	}

	if d.KubernetesAudit.Enabled {
		alertTriggerPolicies["kubernetes_audits"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.KubernetesAudit.Enabled,
				"all_rules": d.KubernetesAudit.Allrules,
				"rules":     d.KubernetesAudit.Rules,
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

	if d.AppEmbeddedAppFirewall.Enabled {
		alertTriggerPolicies["waas_firewall_app_embedded_defender"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.AppEmbeddedAppFirewall.Enabled,
				"all_rules": d.AppEmbeddedAppFirewall.Allrules,
				"rules":     d.AppEmbeddedAppFirewall.Rules,
			},
		}
	}

	if d.ContainerAppFirewall.Enabled {
		alertTriggerPolicies["waas_firewall_container"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.ContainerAppFirewall.Enabled,
				"all_rules": d.ContainerAppFirewall.Allrules,
				"rules":     d.ContainerAppFirewall.Rules,
			},
		}
	}

	if d.HostAppFirewall.Enabled {
		alertTriggerPolicies["waas_firewall_host"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.HostAppFirewall.Enabled,
				"all_rules": d.HostAppFirewall.Allrules,
				"rules":     d.HostAppFirewall.Rules,
			},
		}
	}

	if d.ServerlessAppFirewall.Enabled {
		alertTriggerPolicies["waas_firewall_serverless"] = []interface{}{
			map[string]interface{}{
				"enabled":   d.ServerlessAppFirewall.Enabled,
				"all_rules": d.ServerlessAppFirewall.Allrules,
				"rules":     d.ServerlessAppFirewall.Rules,
			},
		}
	}

	if d.WaasHealth.Enabled {
		alertTriggerPolicies["waas_health"] = []interface{}{
			map[string]interface{}{
				"enabled": d.WaasHealth.Enabled,
			},
		}
	}

	return alertTriggerPolicies
}

// Converts a alertprofile schema to a alertprofile object for SDK compatibility.
func SchemaToAlertprofile(d *schema.ResourceData) (alertprofile.Alertprofile, error) {

	parsedAlertprofile := alertprofile.Alertprofile{}

	if val, ok := d.GetOk("name"); ok {
		parsedAlertprofile.Name = val.(string)
	}

	var isEnabled bool
	if val, ok := d.GetOk("enabled"); ok {
		isEnabled = val.(bool)
	}

	if val, ok := d.GetOk("enable_immediate_vulnerabilities_alerts"); ok {
		parsedAlertprofile.VulnerabilityImmediateAlertsEnabled = val.(bool)
	}

	if ap, ok := d.GetOk("alert_profile_config"); ok {
		for _, val := range ap.([]interface{}) {
			if val.(map[string]interface{})["prisma_cloud_integration_id"] != nil {
				parsedAlertprofile.IntegrationID = val.(map[string]interface{})["prisma_cloud_integration_id"].(string)
				parsedAlertprofile.External = true
			}

			parsedAlertprofile.Webhook.Url = val.(map[string]interface{})["webhook_url"].(string)
			parsedAlertprofile.Webhook.CredentialId = val.(map[string]interface{})["credential_id"].(string)
			parsedAlertprofile.Webhook.CaCert = val.(map[string]interface{})["custom_ca"].(string)
			parsedAlertprofile.Webhook.Json = val.(map[string]interface{})["custom_json"].(string)
		}
	}

	if val, ok := d.GetOk("alert_profile_type"); ok {
		if val == "webhook" {
			parsedAlertprofile.Webhook.Enabled = isEnabled
		}
	}

	if alertTriggers, ok := d.GetOk("alert_triggers"); ok {

		for _, alertTriger := range alertTriggers.([]interface{}) {
			for _, cv := range alertTriger.(map[string]interface{})["access"].([]interface{}) {
				parsedAlertprofile.Policy.Docker.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertprofile.Policy.Docker.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertprofile.Policy.Docker.Rules = append(parsedAlertprofile.Policy.Docker.Rules, rule.(string))
				}
			}
		}

		for _, alertTriger := range alertTriggers.([]interface{}) {
			for _, cv := range alertTriger.(map[string]interface{})["admission"].([]interface{}) {
				parsedAlertprofile.Policy.Admission.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertprofile.Policy.Admission.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertprofile.Policy.Admission.Rules = append(parsedAlertprofile.Policy.Docker.Rules, rule.(string))
				}
			}
		}

		for _, alertTriger := range alertTriggers.([]interface{}) {
			for _, cv := range alertTriger.(map[string]interface{})["app_embedded_defender_runtime"].([]interface{}) {
				parsedAlertprofile.Policy.AppEmbeddedRuntime.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertprofile.Policy.AppEmbeddedRuntime.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertprofile.Policy.AppEmbeddedRuntime.Rules = append(parsedAlertprofile.Policy.Docker.Rules, rule.(string))
				}
			}
		}

		for _, alertTriger := range alertTriggers.([]interface{}) {
			for _, cv := range alertTriger.(map[string]interface{})["cloud_native_network_firewall"].([]interface{}) {
				parsedAlertprofile.Policy.NetworkFirewall.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertprofile.Policy.NetworkFirewall.Allrules = true
			}
		}

		for _, alertTriger := range alertTriggers.([]interface{}) {
			for _, cv := range alertTriger.(map[string]interface{})["container_and_image_compliance"].([]interface{}) {
				parsedAlertprofile.Policy.ContainerComplianceScan.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertprofile.Policy.ContainerComplianceScan.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertprofile.Policy.ContainerComplianceScan.Rules = append(parsedAlertprofile.Policy.Docker.Rules, rule.(string))
				}
			}
		}

		for _, alertTriger := range alertTriggers.([]interface{}) {
			for _, cv := range alertTriger.(map[string]interface{})["container_runtime"].([]interface{}) {
				parsedAlertprofile.Policy.ContainerRuntime.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertprofile.Policy.ContainerRuntime.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertprofile.Policy.ContainerRuntime.Rules = append(parsedAlertprofile.Policy.Docker.Rules, rule.(string))
				}
			}
		}

		for _, alertTriger := range alertTriggers.([]interface{}) {
			for _, cv := range alertTriger.(map[string]interface{})["defender_health"].([]interface{}) {
				parsedAlertprofile.Policy.Defender.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertprofile.Policy.Defender.Allrules = true
			}
		}

		for _, alertTriger := range alertTriggers.([]interface{}) {
			for _, cv := range alertTriger.(map[string]interface{})["host_compliance"].([]interface{}) {
				parsedAlertprofile.Policy.HostComplianceScan.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertprofile.Policy.HostComplianceScan.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertprofile.Policy.HostComplianceScan.Rules = append(parsedAlertprofile.Policy.Docker.Rules, rule.(string))
				}
			}
		}

		for _, alertTriger := range alertTriggers.([]interface{}) {
			for _, cv := range alertTriger.(map[string]interface{})["host_runtime"].([]interface{}) {
				parsedAlertprofile.Policy.HostRuntime.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertprofile.Policy.HostRuntime.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertprofile.Policy.HostRuntime.Rules = append(parsedAlertprofile.Policy.Docker.Rules, rule.(string))
				}
			}
		}

		for _, alertTriger := range alertTriggers.([]interface{}) {
			for _, cv := range alertTriger.(map[string]interface{})["host_vulnerabilities"].([]interface{}) {
				parsedAlertprofile.Policy.HostVulnerability.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertprofile.Policy.HostVulnerability.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertprofile.Policy.HostVulnerability.Rules = append(parsedAlertprofile.Policy.Docker.Rules, rule.(string))
				}
			}
		}

		for _, alertTriger := range alertTriggers.([]interface{}) {
			for _, cv := range alertTriger.(map[string]interface{})["image_vulnerabilities"].([]interface{}) {
				parsedAlertprofile.Policy.ContainerVulnerability.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertprofile.Policy.ContainerVulnerability.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertprofile.Policy.ContainerVulnerability.Rules = append(parsedAlertprofile.Policy.ContainerVulnerability.Rules, rule.(string))
				}
			}
		}

		for _, alertTriger := range alertTriggers.([]interface{}) {
			for _, cv := range alertTriger.(map[string]interface{})["incidents"].([]interface{}) {
				parsedAlertprofile.Policy.Incident.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertprofile.Policy.Incident.Allrules = true
			}
		}

		for _, alertTriger := range alertTriggers.([]interface{}) {
			for _, cv := range alertTriger.(map[string]interface{})["kubernetes_audits"].([]interface{}) {
				parsedAlertprofile.Policy.KubernetesAudit.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertprofile.Policy.KubernetesAudit.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertprofile.Policy.KubernetesAudit.Rules = append(parsedAlertprofile.Policy.Docker.Rules, rule.(string))
				}
			}
		}

		for _, alertTriger := range alertTriggers.([]interface{}) {
			for _, cv := range alertTriger.(map[string]interface{})["serverless_runtime"].([]interface{}) {
				parsedAlertprofile.Policy.ServerlessRuntime.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertprofile.Policy.ServerlessRuntime.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertprofile.Policy.ServerlessRuntime.Rules = append(parsedAlertprofile.Policy.Docker.Rules, rule.(string))
				}
			}
		}

		for _, alertTriger := range alertTriggers.([]interface{}) {
			for _, cv := range alertTriger.(map[string]interface{})["waas_firewall_app_embedded_defender"].([]interface{}) {
				parsedAlertprofile.Policy.AppEmbeddedAppFirewall.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertprofile.Policy.AppEmbeddedAppFirewall.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertprofile.Policy.AppEmbeddedAppFirewall.Rules = append(parsedAlertprofile.Policy.Docker.Rules, rule.(string))
				}
			}
		}

		for _, alertTriger := range alertTriggers.([]interface{}) {
			for _, cv := range alertTriger.(map[string]interface{})["waas_firewall_container"].([]interface{}) {
				parsedAlertprofile.Policy.ContainerAppFirewall.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertprofile.Policy.ContainerAppFirewall.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertprofile.Policy.ContainerAppFirewall.Rules = append(parsedAlertprofile.Policy.Docker.Rules, rule.(string))
				}
			}
		}

		for _, alertTriger := range alertTriggers.([]interface{}) {
			for _, cv := range alertTriger.(map[string]interface{})["waas_firewall_host"].([]interface{}) {
				parsedAlertprofile.Policy.HostAppFirewall.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertprofile.Policy.HostAppFirewall.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertprofile.Policy.HostAppFirewall.Rules = append(parsedAlertprofile.Policy.Docker.Rules, rule.(string))
				}
			}
		}

		for _, alertTriger := range alertTriggers.([]interface{}) {
			for _, cv := range alertTriger.(map[string]interface{})["waas_firewall_serverless"].([]interface{}) {
				parsedAlertprofile.Policy.ServerlessAppFirewall.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertprofile.Policy.ServerlessAppFirewall.Allrules = cv.(map[string]interface{})["all_rules"].(bool)

				for _, rule := range cv.(map[string]interface{})["rules"].([]interface{}) {
					parsedAlertprofile.Policy.ServerlessAppFirewall.Rules = append(parsedAlertprofile.Policy.Docker.Rules, rule.(string))
				}
			}
		}

		for _, alertTriger := range alertTriggers.([]interface{}) {
			for _, cv := range alertTriger.(map[string]interface{})["waas_health"].([]interface{}) {
				parsedAlertprofile.Policy.WaasHealth.Enabled = cv.(map[string]interface{})["enabled"].(bool)
				parsedAlertprofile.Policy.WaasHealth.Allrules = true
			}
		}
	}

	return parsedAlertprofile, nil
}
