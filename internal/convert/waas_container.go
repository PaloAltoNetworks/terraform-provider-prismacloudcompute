package convert

import (
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/policy"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/rule"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaToWaasContainerRules(d *schema.ResourceData) ([]policy.WaasContainerRule, error) {
	parsedRules := make([]policy.WaasContainerRule, 0)
	if rules, ok := d.GetOk("rule"); ok {
		presentRules := rules.([]interface{})
		for _, val := range presentRules {
			presentRule := val.(map[string]interface{})
			parsedRule := policy.WaasContainerRule{}

			parsedRule.Name = presentRule["name"].(string)
			parsedRule.Notes = presentRule["notes"].(string)
			parsedRule.Collections = PolicySchemaToCollections(presentRule["collections"].([]interface{}))
			parsedRule.ReadTimeoutSeconds = presentRule["read_timeout"].(int)

			presentApplications := presentRule["application"].([]interface{})
			parsedApplications := make([]policy.WaasContainerApplicationSpec, 0)
			for _, val := range presentApplications {
				presentApplication := val.(map[string]interface{})
				parsedApplication := policy.WaasContainerApplicationSpec{}
				if isPresent(presentApplication["app_definition"].([]interface{})) {
					presentAppDefinition := presentApplication["app_definition"].([]interface{})[0].(map[string]interface{})
					parsedApplication.AppId = presentAppDefinition["app_id"].(string)

					if isPresent(presentAppDefinition["endpoint_setup"].([]interface{})) {
						presentEndpointSetup := presentAppDefinition["endpoint_setup"].([]interface{})[0].(map[string]interface{})
						parsedApplication.ApiSpec.Description = presentEndpointSetup["description"].(string)
						parsedApplication.ApiSpec.SkipLearning = !presentEndpointSetup["api_endpoint_discovery_enabled"].(bool)

						if isPresent(presentEndpointSetup["tls_config"].([]interface{})) {
							presentTlsConfig := presentEndpointSetup["tls_config"].([]interface{})[0].(map[string]interface{})
							parsedApplication.Certificate.Plain = presentTlsConfig["certificate"].(string)
							parsedApplication.TlsConfig.MinTlsVersion = presentTlsConfig["min_tls_version"].(string)

							if isPresent(presentTlsConfig["hsts_config"].([]interface{})) {
								presentHstsConfig := presentTlsConfig["hsts_config"].([]interface{})[0].(map[string]interface{})
								parsedApplication.TlsConfig.HstsConfig.Enabled = presentHstsConfig["enabled"].(bool)
								parsedApplication.TlsConfig.HstsConfig.MaxAgeSeconds = presentHstsConfig["max_age"].(int)
								parsedApplication.TlsConfig.HstsConfig.IncludeSubdomains = presentHstsConfig["include_subdomains"].(bool)
								parsedApplication.TlsConfig.HstsConfig.Preload = presentHstsConfig["preload"].(bool)
							}
						}
					}

					if isPresent(presentAppDefinition["api_protection"].([]interface{})) {
						presentApiProtection := presentAppDefinition["api_protection"].([]interface{})[0].(map[string]interface{})
						parsedApplication.ApiSpec.Effect = presentApiProtection["violation_effect"].(string)
						parsedApplication.ApiSpec.FallbackEffect = presentApiProtection["unspecified_api_effect"].(string)
						parsedApplication.ApiSpec.QueryParamFallbackEffect = presentApiProtection["unspecified_query_params_effect"].(string)

						parsedPaths := make([]policy.WaasContainerApiSpecPath, 0)
						for _, val := range presentApiProtection["paths"].([]interface{}) {
							presentPath := val.(map[string]interface{})
							parsedPath := policy.WaasContainerApiSpecPath{}
							parsedPath.Path = presentPath["path"].(string)

							parsedMethods := make([]policy.WaasContainerApiSpecPathMethod, 0)
							for _, val := range presentPath["methods"].([]interface{}) {
								presentMethod := val.(map[string]interface{})
								parsedMethod := policy.WaasContainerApiSpecPathMethod{}
								parsedMethod.Method = presentMethod["method"].(string)

								parsedParameters := make([]policy.WaasContainerApiSpecPathMethodParameter, 0)
								for _, val := range presentMethod["parameters"].([]interface{}) {
									presentParameter := val.(map[string]interface{})
									parsedParameter := policy.WaasContainerApiSpecPathMethodParameter{}
									parsedParameter.Name = presentParameter["name"].(string)
									parsedParameter.Type = presentParameter["type"].(string)
									parsedParameter.Location = presentParameter["location"].(string)
									parsedParameter.Style = presentParameter["style"].(string)
									parsedParameter.Min = presentParameter["min"].(int)
									parsedParameter.Max = presentParameter["max"].(int)
									parsedParameter.Explode = presentParameter["explode"].(bool)
									parsedParameter.Array = presentParameter["array"].(bool)
									parsedParameter.Required = presentParameter["required"].(bool)
									parsedParameter.AllowEmptyValue = presentParameter["allow_empty_value"].(bool)
									parsedParameters = append(parsedParameters, parsedParameter)
								}
								parsedMethod.Parameters = parsedParameters
								parsedMethods = append(parsedMethods, parsedMethod)
							}
							parsedPath.Methods = parsedMethods
							parsedPaths = append(parsedPaths, parsedPath)
						}
						parsedApplication.ApiSpec.Paths = parsedPaths
					}
				}

				if isPresent(presentApplication["app_firewall"].([]interface{})) {
					presentAppFirewall := presentApplication["app_firewall"].([]interface{})[0].(map[string]interface{})

					if isPresent(presentAppFirewall["sql_injection"].([]interface{})) {
						presentSqlInjection := presentAppFirewall["sql_injection"].([]interface{})[0].(map[string]interface{})
						parsedApplication.Sqli.Effect = presentSqlInjection["effect"].(string)
						parsedExceptions := make([]policy.WaasContainerExceptionField, 0)
						for _, val := range presentSqlInjection["exceptions"].([]interface{}) {
							presentException := val.(map[string]interface{})
							parsedException := policy.WaasContainerExceptionField{}
							parsedException.Location = presentException["location"].(string)
							parsedException.Key = presentException["value"].(string)
							parsedExceptions = append(parsedExceptions, parsedException)
						}
						parsedApplication.Sqli.ExceptionFields = parsedExceptions
					}

					if isPresent(presentAppFirewall["xss"].([]interface{})) {
						presentXss := presentAppFirewall["xss"].([]interface{})[0].(map[string]interface{})
						parsedApplication.Xss.Effect = presentXss["effect"].(string)
						parsedExceptions := make([]policy.WaasContainerExceptionField, 0)
						for _, val := range presentXss["exceptions"].([]interface{}) {
							presentException := val.(map[string]interface{})
							parsedException := policy.WaasContainerExceptionField{}
							parsedException.Location = presentException["location"].(string)
							parsedException.Key = presentException["value"].(string)
							parsedExceptions = append(parsedExceptions, parsedException)
						}
						parsedApplication.Xss.ExceptionFields = parsedExceptions
					}

					if isPresent(presentAppFirewall["command_injection"].([]interface{})) {
						presentCommandInjection := presentAppFirewall["command_injection"].([]interface{})[0].(map[string]interface{})
						parsedApplication.Cmdi.Effect = presentCommandInjection["effect"].(string)
						parsedExceptions := make([]policy.WaasContainerExceptionField, 0)
						for _, val := range presentCommandInjection["exceptions"].([]interface{}) {
							presentException := val.(map[string]interface{})
							parsedException := policy.WaasContainerExceptionField{}
							parsedException.Location = presentException["location"].(string)
							parsedException.Key = presentException["value"].(string)
							parsedExceptions = append(parsedExceptions, parsedException)
						}
						parsedApplication.Cmdi.ExceptionFields = parsedExceptions
					}

					if isPresent(presentAppFirewall["code_injection"].([]interface{})) {
						presentCodeInjection := presentAppFirewall["code_injection"].([]interface{})[0].(map[string]interface{})
						parsedApplication.CodeInjection.Effect = presentCodeInjection["effect"].(string)
						parsedExceptions := make([]policy.WaasContainerExceptionField, 0)
						for _, val := range presentCodeInjection["exceptions"].([]interface{}) {
							presentException := val.(map[string]interface{})
							parsedException := policy.WaasContainerExceptionField{}
							parsedException.Location = presentException["location"].(string)
							parsedException.Key = presentException["value"].(string)
							parsedExceptions = append(parsedExceptions, parsedException)
						}
						parsedApplication.CodeInjection.ExceptionFields = parsedExceptions
					}

					if isPresent(presentAppFirewall["local_file_inclusion"].([]interface{})) {
						presentLocalFileInclusion := presentAppFirewall["local_file_inclusion"].([]interface{})[0].(map[string]interface{})
						parsedApplication.Lfi.Effect = presentLocalFileInclusion["effect"].(string)
						parsedExceptions := make([]policy.WaasContainerExceptionField, 0)
						for _, val := range presentLocalFileInclusion["exceptions"].([]interface{}) {
							presentException := val.(map[string]interface{})
							parsedException := policy.WaasContainerExceptionField{}
							parsedException.Location = presentException["location"].(string)
							parsedException.Key = presentException["value"].(string)
							parsedExceptions = append(parsedExceptions, parsedException)
						}
						parsedApplication.Lfi.ExceptionFields = parsedExceptions
					}

					if isPresent(presentAppFirewall["attack_tools"].([]interface{})) {
						presentAttackTools := presentAppFirewall["attack_tools"].([]interface{})[0].(map[string]interface{})
						parsedApplication.AttackTools.Effect = presentAttackTools["effect"].(string)
						parsedExceptions := make([]policy.WaasContainerExceptionField, 0)
						for _, val := range presentAttackTools["exceptions"].([]interface{}) {
							presentException := val.(map[string]interface{})
							parsedException := policy.WaasContainerExceptionField{}
							parsedException.Location = presentException["location"].(string)
							parsedException.Key = presentException["value"].(string)
							parsedExceptions = append(parsedExceptions, parsedException)
						}
						parsedApplication.AttackTools.ExceptionFields = parsedExceptions
					}

					if isPresent(presentAppFirewall["shellshock"].([]interface{})) {
						presentShellshock := presentAppFirewall["shellshock"].([]interface{})[0].(map[string]interface{})
						parsedApplication.Shellshock.Effect = presentShellshock["effect"].(string)
					}

					if isPresent(presentAppFirewall["malformed_request"].([]interface{})) {
						presentMalformedRequest := presentAppFirewall["malformed_request"].([]interface{})[0].(map[string]interface{})
						parsedApplication.MalformedReq.Effect = presentMalformedRequest["effect"].(string)
					}

					if isPresent(presentAppFirewall["advanced_threat_protection"].([]interface{})) {
						presentAdvancedThreatProtection := presentAppFirewall["advanced_threat_protection"].([]interface{})[0].(map[string]interface{})
						parsedApplication.NetworkControls.AdvancedProtectionEffect = presentAdvancedThreatProtection["effect"].(string)
					}

					if isPresent(presentAppFirewall["information_leakage"].([]interface{})) {
						presentInformationLeakage := presentAppFirewall["information_leakage"].([]interface{})[0].(map[string]interface{})
						parsedApplication.IntelGathering.InfoLeakageEffect = presentInformationLeakage["effect"].(string)
					}

					parsedApplication.CsrfEnabled = presentAppFirewall["csrf_protection_enabled"].(bool)
					parsedApplication.ClickjackingEnabled = presentAppFirewall["clickjacking_prevention_enabled"].(bool)
					parsedApplication.IntelGathering.RemoveFingerprintsEnabled = presentAppFirewall["remove_fingerprints"].(bool)
				}

				if isPresent(presentApplication["dos_protection"].([]interface{})) {
					presentDosProtection := presentApplication["dos_protection"].([]interface{})[0].(map[string]interface{})
					parsedApplication.DosConfig.Enabled = presentDosProtection["enabled"].(bool)

					if isPresent(presentDosProtection["alert"].([]interface{})) {
						presentAlert := presentDosProtection["alert"].([]interface{})[0].(map[string]interface{})
						parsedApplication.DosConfig.Alert.Average = presentAlert["average"].(int)
						parsedApplication.DosConfig.Alert.Burst = presentAlert["burst"].(int)
					}

					if isPresent(presentDosProtection["ban"].([]interface{})) {
						presentBan := presentDosProtection["ban"].([]interface{})[0].(map[string]interface{})
						parsedApplication.DosConfig.Ban.Average = presentBan["average"].(int)
						parsedApplication.DosConfig.Ban.Burst = presentBan["burst"].(int)
					}

					parsedMatchConditions := make([]policy.WaasContainerDosConfigMatchCondition, 0)
					for _, val := range presentDosProtection["match_condition"].([]interface{}) {
						presentMatchCondition := val.(map[string]interface{})
						parsedMatchCondition := policy.WaasContainerDosConfigMatchCondition{}
						parsedMatchCondition.Methods = presentMatchCondition["methods"].([]string)
						parsedMatchCondition.FileTypes = presentMatchCondition["file_types"].([]string)

						parsedResponseCodeRanges := make([]policy.WaasContainerDosConfigMatchConditionResponseCodeRange, 0)
						for _, val := range presentMatchCondition["response_code_range"].([]interface{}) {
							presentResponseCodeRange := val.(map[string]interface{})
							parsedResponseCodeRange := policy.WaasContainerDosConfigMatchConditionResponseCodeRange{}
							parsedResponseCodeRange.Start = presentResponseCodeRange["start"].(int)
							parsedResponseCodeRange.End = presentResponseCodeRange["end"].(int)
							parsedResponseCodeRanges = append(parsedResponseCodeRanges, parsedResponseCodeRange)
						}
						parsedMatchCondition.ResponseCodeRanges = parsedResponseCodeRanges
						parsedMatchConditions = append(parsedMatchConditions, parsedMatchCondition)
					}
					parsedApplication.DosConfig.MatchConditions = parsedMatchConditions
					parsedApplication.DosConfig.ExcludedNetworkLists = presentDosProtection["excluded_networks"].([]string)
				}

				if isPresent(presentApplication["access_control"].([]interface{})) {
					presentAccessControl := presentApplication["access_control"].([]interface{})[0].(map[string]interface{})

					if isPresent(presentAccessControl["network_controls"].([]interface{})) {
						presentNetworkControls := presentAccessControl["network_controls"].([]interface{})[0].(map[string]interface{})

						if isPresent(presentNetworkControls["ip_access_control"].([]interface{})) {
							presentIpAccessControl := presentNetworkControls["ip_access_control"].([]interface{})[0].(map[string]interface{})
							parsedApplication.NetworkControls.Subnets.Enabled = presentIpAccessControl["enabled"].(bool)
							parsedApplication.NetworkControls.Subnets.AllowMode = presentIpAccessControl["allowed"].(bool)
							parsedApplication.NetworkControls.Subnets.Allow = presentIpAccessControl["allowed_network_lists"].([]string)
							parsedApplication.NetworkControls.Subnets.Alert = presentIpAccessControl["alerted_network_lists"].([]string)
							parsedApplication.NetworkControls.Subnets.Prevent = presentIpAccessControl["prevented_network_lists"].([]string)
							parsedApplication.NetworkControls.Subnets.FallbackEffect = presentIpAccessControl["fallback_effect"].(string)
						}

						if isPresent(presentNetworkControls["geo_access_control"].([]interface{})) {
							presentGeoAccessControl := presentNetworkControls["geo_access_control"].([]interface{})[0].(map[string]interface{})
							parsedApplication.NetworkControls.Countries.Enabled = presentGeoAccessControl["enabled"].(bool)
							parsedApplication.NetworkControls.Countries.AllowMode = presentGeoAccessControl["allowed"].(bool)
							parsedApplication.NetworkControls.Countries.Allow = presentGeoAccessControl["allowed_countries"].([]string)
							parsedApplication.NetworkControls.Countries.Alert = presentGeoAccessControl["alerted_countries"].([]string)
							parsedApplication.NetworkControls.Countries.Prevent = presentGeoAccessControl["prevented_countries"].([]string)
							parsedApplication.NetworkControls.Countries.FallbackEffect = presentGeoAccessControl["fallback_effect"].(string)
						}

						parsedApplication.NetworkControls.ExceptionSubnets = presentNetworkControls["network_list_exceptions"].([]string)
					}

					parsedHttpHeaders := make([]policy.WaasContainerHeaderSpec, 0)
					for _, val := range presentAccessControl["http_header"].([]interface{}) {
						presentHttpHeader := val.(map[string]interface{})
						parsedHttpHeader := policy.WaasContainerHeaderSpec{}
						parsedHttpHeader.Name = presentHttpHeader["name"].(string)
						parsedHttpHeader.Allow = presentHttpHeader["allowed"].(bool)
						parsedHttpHeader.Values = presentHttpHeader["values"].([]string)
						parsedHttpHeader.Effect = presentHttpHeader["effect"].(string)
						parsedHttpHeader.Required = presentHttpHeader["required"].(bool)
						parsedHttpHeaders = append(parsedHttpHeaders, parsedHttpHeader)
					}
					parsedApplication.HeaderSpecs = parsedHttpHeaders

					if isPresent(presentAccessControl["file_uploads"].([]interface{})) {
						presentFileUploads := presentAccessControl["file_uploads"].([]interface{})[0].(map[string]interface{})
						parsedApplication.MaliciousUpload.AllowedExtensions = presentFileUploads["allowed_extensions"].([]string)
						parsedApplication.MaliciousUpload.AllowedFileTypes = presentFileUploads["allowed_file_types"].([]string)
						parsedApplication.MaliciousUpload.Effect = presentFileUploads["effect"].(string)
					}
				}

				if isPresent(presentApplication["bot_protection"].([]interface{})) {
					presentBotProtection := presentApplication["bot_protection"].([]interface{})[0].(map[string]interface{})

					if isPresent(presentBotProtection["known_bots"].([]interface{})) {
						presentKnownBots := presentBotProtection["known_bots"].([]interface{})[0].(map[string]interface{})
						parsedApplication.BotProtectionSpec.KnownBotProtectionsSpec.SearchEngineCrawlers = presentKnownBots["search_engine_crawlers"].(string)
						parsedApplication.BotProtectionSpec.KnownBotProtectionsSpec.BusinessAnalytics = presentKnownBots["business_analytics"].(string)
						parsedApplication.BotProtectionSpec.KnownBotProtectionsSpec.Educational = presentKnownBots["educational"].(string)
						parsedApplication.BotProtectionSpec.KnownBotProtectionsSpec.News = presentKnownBots["news"].(string)
						parsedApplication.BotProtectionSpec.KnownBotProtectionsSpec.Financial = presentKnownBots["financial"].(string)
						parsedApplication.BotProtectionSpec.KnownBotProtectionsSpec.ContentFeedClients = presentKnownBots["content_feed_clients"].(string)
						parsedApplication.BotProtectionSpec.KnownBotProtectionsSpec.Archiving = presentKnownBots["archiving"].(string)
						parsedApplication.BotProtectionSpec.KnownBotProtectionsSpec.CareerSearch = presentKnownBots["career_search"].(string)
						parsedApplication.BotProtectionSpec.KnownBotProtectionsSpec.MediaSearch = presentKnownBots["media_search"].(string)
					}

					if isPresent(presentBotProtection["unknown_bots"].([]interface{})) {
						presentUnknownBots := presentBotProtection["unknown_bots"].([]interface{})[0].(map[string]interface{})
						parsedApplication.BotProtectionSpec.UnknownBotProtectionSpec.Generic = presentUnknownBots["generic"].(string)
						parsedApplication.BotProtectionSpec.UnknownBotProtectionSpec.WebAutomationTools = presentUnknownBots["web_automation_tools"].(string)
						parsedApplication.BotProtectionSpec.UnknownBotProtectionSpec.WebScrapers = presentUnknownBots["web_scrapers"].(string)
						parsedApplication.BotProtectionSpec.UnknownBotProtectionSpec.ApiLibraries = presentUnknownBots["api_libraries"].(string)
						parsedApplication.BotProtectionSpec.UnknownBotProtectionSpec.HttpLibraries = presentUnknownBots["http_libraries"].(string)

						if isPresent(presentUnknownBots["request_anomalies"].([]interface{})) {
							presentRequestAnomalies := presentUnknownBots["request_anomalies"].([]interface{})[0].(map[string]interface{})
							parsedApplication.BotProtectionSpec.UnknownBotProtectionSpec.RequestAnomalies.Effect = presentRequestAnomalies["effect"].(string)
							parsedApplication.BotProtectionSpec.UnknownBotProtectionSpec.RequestAnomalies.Threshold = convertRequestAnomaliesThreshold(presentRequestAnomalies["threshold"].(string))
						}
						parsedApplication.BotProtectionSpec.UnknownBotProtectionSpec.BotImpersonation = presentUnknownBots["bot_impersonation"].(string)
					}

					parsedUserDefinedBots := make([]policy.WaasContainerBotProtectionSpecUserDefinedBot, 0)
					for _, val := range presentBotProtection["user_defined_bot"].([]interface{}) {
						presentUserDefinedBot := val.(map[string]interface{})
						parsedUserDefinedBot := policy.WaasContainerBotProtectionSpecUserDefinedBot{}
						parsedUserDefinedBot.Name = presentUserDefinedBot["name"].(string)
						parsedUserDefinedBot.HeaderName = presentUserDefinedBot["header_name"].(string)
						parsedUserDefinedBot.HeaderValues = presentUserDefinedBot["header_values"].([]string)
						parsedUserDefinedBot.Subnets = presentUserDefinedBot["source_subnets"].([]string)
						parsedUserDefinedBot.Effect = presentUserDefinedBot["effect"].(string)
						parsedUserDefinedBots = append(parsedUserDefinedBots, parsedUserDefinedBot)
					}
					parsedApplication.BotProtectionSpec.UserDefinedBots = parsedUserDefinedBots

					if isPresent(presentBotProtection["active_bot_detection"].([]interface{})) {
						presentActiveBotDetection := presentBotProtection["active_bot_detection"].([]interface{})[0].(map[string]interface{})
						parsedApplication.BotProtectionSpec.SessionValidation = presentActiveBotDetection["session_validation_failure_effect"].(string)

						if isPresent(presentActiveBotDetection["javascript_based_detection"].([]interface{})) {
							presentJavascriptBasedDetection := presentActiveBotDetection["javascript_based_detection"].([]interface{})[0].(map[string]interface{})
							parsedApplication.BotProtectionSpec.JsInjectionSpec.Enabled = presentJavascriptBasedDetection["enabled"].(bool)
							parsedApplication.BotProtectionSpec.UnknownBotProtectionSpec.BrowserImpersonation = presentJavascriptBasedDetection["browser_impersonation"].(string)
							parsedApplication.BotProtectionSpec.JsInjectionSpec.TimeoutEffect = presentJavascriptBasedDetection["injection_timeout_effect"].(string)
						}

						if isPresent(presentActiveBotDetection["recaptcha"].([]interface{})) {
							presentRecaptcha := presentActiveBotDetection["recaptcha"].([]interface{})[0].(map[string]interface{})
							parsedApplication.BotProtectionSpec.RecaptchaSpec.Enabled = presentRecaptcha["enabled"].(bool)
							parsedApplication.BotProtectionSpec.RecaptchaSpec.SiteKey = presentRecaptcha["site_key"].(string)
							parsedApplication.BotProtectionSpec.RecaptchaSpec.SecretKey.Plain = presentRecaptcha["secret_key"].(string)
							parsedApplication.BotProtectionSpec.RecaptchaSpec.Type = presentRecaptcha["type"].(string)
							parsedApplication.BotProtectionSpec.RecaptchaSpec.AllSessions = presentRecaptcha["every_new_session"].(bool)
							parsedApplication.BotProtectionSpec.RecaptchaSpec.SuccessExpirationHours = presentRecaptcha["success_expiration"].(int)
						}
					}
				}

				parsedCustomRules := make([]rule.CustomRule, 0)
				for _, val := range presentApplication["custom_rule"].([]interface{}) {
					presentCustomRule := val.(map[string]interface{})
					parsedCustomRule := rule.CustomRule{}
					parsedCustomRule.Id = presentCustomRule["rule_id"].(int)
					parsedCustomRule.Effect = presentCustomRule["effect"].(string)
					parsedCustomRules = append(parsedCustomRules, parsedCustomRule)
				}
				parsedApplication.CustomRules = parsedCustomRules

				if isPresent(presentApplication["advanced_settings"].([]interface{})) {
					presentAdvancedSettings := presentApplication["advanced_settings"].([]interface{})[0].(map[string]interface{})

					if isPresent(presentAdvancedSettings["session_cookies"].([]interface{})) {
						presentSessionCookies := presentAdvancedSettings["session_cookies"].([]interface{})[0].(map[string]interface{})
						parsedApplication.SessionCookieEnabled = presentSessionCookies["enabled"].(bool)
						parsedApplication.SessionCookieSameSite = presentSessionCookies["same_site"].(string)
						parsedApplication.SessionCookieSecure = presentSessionCookies["secure"].(bool)
						parsedApplication.SessionCookieBan = presentSessionCookies["apply_ban_on_session"].(bool)
						parsedApplication.DosConfig.TrackSession = presentSessionCookies["rate_limit_on_session"].(bool)
					}

					parsedApplication.BanDurationMinutes = presentAdvancedSettings["ban_duration"].(int)

					if isPresent(presentAdvancedSettings["http_body_inspection"].([]interface{})) {
						presentHttpBodyInspection := presentAdvancedSettings["http_body_inspection"].([]interface{})[0].(map[string]interface{})
						parsedApplication.Body.Skip = !presentHttpBodyInspection["enabled"].(bool)
						parsedApplication.Body.InspectionSizeBytes = presentHttpBodyInspection["inspection_size"].(int)
					}

					if isPresent(presentAdvancedSettings["custom_waas_response"].([]interface{})) {
						presentCustomWaasResponse := presentAdvancedSettings["custom_waas_response"].([]interface{})[0].(map[string]interface{})
						parsedApplication.CustomBlockResponse.Enabled = presentCustomWaasResponse["enabled"].(bool)
						parsedApplication.CustomBlockResponse.Code = presentCustomWaasResponse["status_code"].(int)
						parsedApplication.CustomBlockResponse.Body = presentCustomWaasResponse["body"].(string)
					}

					parsedApplication.DisableEventIdHeader = !presentAdvancedSettings["enable_event_id_header"].(bool)
				}

				parsedApplications = append(parsedApplications, parsedApplication)
			}

			parsedRule.ApplicationsSpec = parsedApplications
			parsedRules = append(parsedRules, parsedRule)
		}
	}
	return parsedRules, nil
}

func WaasContainerRulesToSchema(in []policy.WaasContainerRule) []interface{} {
	rules := make([]interface{}, 0, len(in))
	for _, i := range in {
		rule := make(map[string]interface{})
		rule["name"] = i.Name
		rule["notes"] = i.Notes
		rule["collections"] = CollectionsToPolicySchema(i.Collections)
		rule["read_timeout"] = i.ReadTimeoutSeconds

		// application
		applicationSlice := make([]interface{}, 0, len(i.ApplicationsSpec))
		for _, j := range i.ApplicationsSpec {
			application := make(map[string]interface{})

			// app_definition
			appDefinitionSlice := make([]interface{}, 0, 1)
			appDefinition := make(map[string]interface{})
			appDefinition["app_id"] = j.AppId

			// endpoint_setup
			endpointSetupSlice := make([]interface{}, 0, 1)
			endpointSetup := make(map[string]interface{})
			endpointSetup["description"] = j.ApiSpec.Description
			endpointSetup["api_endpoint_discovery_enabled"] = !j.ApiSpec.SkipLearning

			// tls_config
			tlsConfigSlice := make([]interface{}, 0, 1)
			tlsConfig := make(map[string]interface{})
			tlsConfig["certificate"] = j.Certificate.Plain
			tlsConfig["min_tls_version"] = j.TlsConfig.MinTlsVersion

			// hsts_config
			hstsConfigSlice := make([]interface{}, 0, 1)
			hstsConfig := make(map[string]interface{})
			hstsConfig["enabled"] = j.TlsConfig.HstsConfig.Enabled
			hstsConfig["max_age"] = j.TlsConfig.HstsConfig.MaxAgeSeconds
			hstsConfig["include_subdomains"] = j.TlsConfig.HstsConfig.IncludeSubdomains
			hstsConfig["preload"] = j.TlsConfig.HstsConfig.Preload
			tlsConfig["hsts_config"] = hstsConfigSlice
			endpointSetup["tls_config"] = tlsConfigSlice

			// endpoint
			endpointSlice := make([]interface{}, 0, len(j.ApiSpec.Endpoints))
			for _, k := range j.ApiSpec.Endpoints {
				endpoint := make(map[string]interface{})
				endpoint["host"] = k.Host
				endpoint["external_port"] = k.ExposedPort
				endpoint["base_path"] = k.BasePath
				endpoint["internal_port"] = k.InternalPort
				endpoint["tls"] = k.Tls
				endpoint["grpc"] = k.Grpc
				endpoint["http2"] = k.Http2
				endpointSlice = append(endpointSlice, endpoint)
			}
			endpointSetup["endpoints"] = endpointSlice
			endpointSetupSlice = append(endpointSetupSlice, endpointSetup)
			appDefinition["endpoint_setup"] = endpointSetupSlice

			// api_protection
			apiProtectionSlice := make([]interface{}, 0, 1)
			apiProtection := make(map[string]interface{})
			apiProtection["violation_effect"] = j.ApiSpec.Effect
			apiProtection["unspecified_api_effect"] = j.ApiSpec.FallbackEffect
			apiProtection["unspecified_query_params_effect"] = j.ApiSpec.QueryParamFallbackEffect

			// path
			pathSlice := make([]interface{}, 0, len(j.ApiSpec.Paths))
			for _, l := range j.ApiSpec.Paths {
				path := make(map[string]interface{})
				path["path"] = l.Path

				// method
				methodSlice := make([]interface{}, 0, len(l.Methods))
				for _, m := range l.Methods {
					method := make(map[string]interface{})
					method["method"] = m.Method

					// parameter
					parameterSlice := make([]interface{}, 0, len(m.Parameters))
					for _, n := range m.Parameters {
						parameter := make(map[string]interface{})
						parameter["name"] = n.Name
						parameter["type"] = n.Type
						parameter["location"] = n.Location
						parameter["style"] = n.Style
						parameter["min"] = n.Min
						parameter["max"] = n.Max
						parameter["explode"] = n.Explode
						parameter["array"] = n.Array
						parameter["required"] = n.Required
						parameter["allow_empty_value"] = n.AllowEmptyValue
						parameterSlice = append(parameterSlice, parameter)
					}
					method["parameter"] = parameterSlice
					methodSlice = append(methodSlice, method)
				}
				path["method"] = methodSlice
				pathSlice = append(pathSlice, path)
			}
			apiProtection["path"] = pathSlice
			apiProtectionSlice = append(apiProtectionSlice, apiProtection)
			appDefinition["api_protection"] = apiProtectionSlice
			appDefinitionSlice = append(appDefinitionSlice, appDefinition)
			application["app_definition"] = appDefinitionSlice

			// app_firewall
			appFirewallSlice := make([]interface{}, 0, 1)
			appFirewall := make(map[string]interface{})

			// sql_injection
			sqlInjectionSlice := make([]interface{}, 0, 1)
			sqlInjection := make(map[string]interface{})
			sqlInjection["effect"] = j.Sqli.Effect

			// exception
			sqlInjectionExceptionSlice := make([]interface{}, 0, len(j.Sqli.ExceptionFields))
			for _, o := range j.Sqli.ExceptionFields {
				exception := make(map[string]interface{})
				exception["location"] = o.Location
				exception["value"] = o.Key
				sqlInjectionExceptionSlice = append(sqlInjectionExceptionSlice, exception)
			}
			sqlInjection["exception"] = sqlInjectionExceptionSlice
			sqlInjectionSlice = append(sqlInjectionSlice, sqlInjection)
			appFirewall["sql_injection"] = sqlInjectionSlice

			// xss
			xssSlice := make([]interface{}, 0, 1)
			xss := make(map[string]interface{})
			xss["effect"] = j.Xss.Effect

			// exception
			xssExceptionSlice := make([]interface{}, 0, len(j.Xss.ExceptionFields))
			for _, o := range j.Xss.ExceptionFields {
				exception := make(map[string]interface{})
				exception["location"] = o.Location
				exception["value"] = o.Key
				xssExceptionSlice = append(xssExceptionSlice, exception)
			}
			xss["exception"] = xssExceptionSlice
			xssSlice = append(xssSlice, xss)
			appFirewall["xss"] = xssSlice

			// command_injection
			commandInjectionSlice := make([]interface{}, 0, 1)
			commandInjection := make(map[string]interface{})
			commandInjection["effect"] = j.Cmdi.Effect

			// exception
			commandInjectionExceptionSlice := make([]interface{}, 0, len(j.Cmdi.ExceptionFields))
			for _, o := range j.Cmdi.ExceptionFields {
				exception := make(map[string]interface{})
				exception["location"] = o.Location
				exception["value"] = o.Key
				commandInjectionExceptionSlice = append(commandInjectionExceptionSlice, exception)
			}
			commandInjection["exception"] = commandInjectionExceptionSlice
			commandInjectionSlice = append(commandInjectionSlice, commandInjection)
			appFirewall["command_injection"] = commandInjectionSlice

			// code_injection
			codeInjectionSlice := make([]interface{}, 0, 1)
			codeInjection := make(map[string]interface{})
			codeInjection["effect"] = j.CodeInjection.Effect

			// exception
			codeInjectionExceptionSlice := make([]interface{}, 0, len(j.CodeInjection.ExceptionFields))
			for _, o := range j.CodeInjection.ExceptionFields {
				exception := make(map[string]interface{})
				exception["location"] = o.Location
				exception["value"] = o.Key
				codeInjectionExceptionSlice = append(codeInjectionExceptionSlice, exception)
			}
			codeInjection["exception"] = codeInjectionExceptionSlice
			codeInjectionSlice = append(codeInjectionSlice, codeInjection)
			appFirewall["code_injection"] = codeInjectionSlice

			// local_file_inclusion
			localFileInclusionSlice := make([]interface{}, 0, 1)
			localFileInclusion := make(map[string]interface{})
			localFileInclusion["effect"] = j.Lfi.Effect

			// exception
			localFileInclusionExceptionSlice := make([]interface{}, 0, len(j.Lfi.ExceptionFields))
			for _, o := range j.Lfi.ExceptionFields {
				exception := make(map[string]interface{})
				exception["location"] = o.Location
				exception["value"] = o.Key
				localFileInclusionExceptionSlice = append(localFileInclusionExceptionSlice, exception)
			}
			localFileInclusion["exception"] = localFileInclusionExceptionSlice
			localFileInclusionSlice = append(localFileInclusionSlice, localFileInclusion)
			appFirewall["local_file_inclusion"] = localFileInclusionSlice

			// attack_tools
			attackToolsSlice := make([]interface{}, 0, 1)
			attackTools := make(map[string]interface{})
			attackTools["effect"] = j.AttackTools.Effect

			// exception
			attackToolsExceptionSlice := make([]interface{}, 0, len(j.AttackTools.ExceptionFields))
			for _, o := range j.AttackTools.ExceptionFields {
				exception := make(map[string]interface{})
				exception["location"] = o.Location
				exception["value"] = o.Key
				attackToolsExceptionSlice = append(attackToolsExceptionSlice, exception)
			}
			attackTools["exception"] = attackToolsExceptionSlice
			attackToolsSlice = append(attackToolsSlice, attackTools)
			appFirewall["attack_tools"] = attackToolsSlice

			// shellshock
			shellshockSlice := make([]interface{}, 0, 1)
			shellshock := make(map[string]interface{})
			shellshock["effect"] = j.Shellshock.Effect
			appFirewall["shellshock"] = shellshockSlice

			// malformed_request
			malformedRequestSlice := make([]interface{}, 0, 1)
			malformedRequest := make(map[string]interface{})
			malformedRequest["effect"] = j.MalformedReq.Effect
			appFirewall["malformed_request"] = malformedRequestSlice

			// advanced_threat_protection
			advancedThreatProtectionSlice := make([]interface{}, 0, 1)
			advancedThreatProtection := make(map[string]interface{})
			advancedThreatProtection["effect"] = j.NetworkControls.AdvancedProtectionEffect
			appFirewall["advanced_threat_protection"] = advancedThreatProtectionSlice

			// information_leakage
			informationLeakageSlice := make([]interface{}, 0, 1)
			informationLeakage := make(map[string]interface{})
			informationLeakage["effect"] = j.IntelGathering.InfoLeakageEffect
			appFirewall["information_leakage"] = informationLeakageSlice

			appFirewall["csrf_protection_enabled"] = j.CsrfEnabled
			appFirewall["clickjacking_prevention_enabled"] = j.ClickjackingEnabled
			appFirewall["remove_fingerprints"] = j.IntelGathering.RemoveFingerprintsEnabled
			appFirewallSlice = append(appFirewallSlice, appFirewall)
			application["app_firewall"] = appFirewallSlice

			// dos_protection
			dosProtectionSlice := make([]interface{}, 0, 1)
			dosProtection := make(map[string]interface{})
			dosProtection["enabled"] = j.DosConfig.Enabled

			// alert
			dosProtectionAlertSlice := make([]interface{}, 0, 1)
			dosProtectionAlert := make(map[string]interface{})
			dosProtectionAlert["average"] = j.DosConfig.Alert.Average
			dosProtectionAlert["burst"] = j.DosConfig.Alert.Burst
			dosProtectionAlertSlice = append(dosProtectionAlertSlice, dosProtectionAlert)
			dosProtection["alert"] = dosProtectionAlertSlice

			// ban
			dosProtectionBanSlice := make([]interface{}, 0, 1)
			dosProtectionBan := make(map[string]interface{})
			dosProtectionBan["average"] = j.DosConfig.Ban.Average
			dosProtectionBan["burst"] = j.DosConfig.Ban.Burst
			dosProtectionBanSlice = append(dosProtectionBanSlice, dosProtectionBan)
			dosProtection["ban"] = dosProtectionBanSlice

			// match_condition
			dosProtectionMatchConditionSlice := make([]interface{}, 0, len(j.DosConfig.MatchConditions))
			for _, o := range j.DosConfig.MatchConditions {
				matchCondition := make(map[string]interface{})
				matchCondition["methods"] = o.Methods
				matchCondition["file_types"] = o.FileTypes

				// response_code_range
				responseCodeRangeSlice := make([]interface{}, 0, len(o.ResponseCodeRanges))
				for _, p := range o.ResponseCodeRanges {
					responseCodeRange := make(map[string]interface{})
					responseCodeRange["start"] = p.Start
					responseCodeRange["end"] = p.End
					responseCodeRangeSlice = append(responseCodeRangeSlice, responseCodeRange)
				}
				matchCondition["response_code_range"] = responseCodeRangeSlice
				dosProtectionMatchConditionSlice = append(dosProtectionMatchConditionSlice, matchCondition)
			}
			dosProtection["match_condition"] = dosProtectionMatchConditionSlice
			dosProtection["excluded_networks"] = j.DosConfig.ExcludedNetworkLists
			dosProtectionSlice = append(dosProtectionSlice, dosProtection)
			application["dos_protection"] = dosProtectionSlice

			// access_control
			accessControlSlice := make([]interface{}, 0, 1)
			accessControl := make(map[string]interface{})

			// network_controls
			networkControlsSlice := make([]interface{}, 0, 1)
			networkControls := make(map[string]interface{})

			// ip_access_control
			ipAccessControlSlice := make([]interface{}, 0, 1)
			ipAccessControl := make(map[string]interface{})
			ipAccessControl["enabled"] = j.NetworkControls.Subnets.Enabled
			ipAccessControl["allow"] = j.NetworkControls.Subnets.AllowMode
			ipAccessControl["allowed_network_lists"] = j.NetworkControls.Subnets.Allow
			ipAccessControl["alerted_network_lists"] = j.NetworkControls.Subnets.Alert
			ipAccessControl["prevented_network_lists"] = j.NetworkControls.Subnets.Prevent
			ipAccessControl["fallback_effect"] = j.NetworkControls.Subnets.FallbackEffect
			ipAccessControlSlice = append(ipAccessControlSlice, ipAccessControl)
			networkControls["ip_access_control"] = ipAccessControlSlice

			// geo_access_control
			geoAccessControlSlice := make([]interface{}, 0, 1)
			geoAccessControl := make(map[string]interface{})
			geoAccessControl["enabled"] = j.NetworkControls.Countries.Enabled
			geoAccessControl["allow"] = j.NetworkControls.Countries.AllowMode
			geoAccessControl["allowed_countries"] = j.NetworkControls.Countries.Allow
			geoAccessControl["alerted_countries"] = j.NetworkControls.Countries.Alert
			geoAccessControl["prevented_countries"] = j.NetworkControls.Countries.Prevent
			geoAccessControl["fallback_effect"] = j.NetworkControls.Countries.FallbackEffect
			geoAccessControlSlice = append(geoAccessControlSlice, geoAccessControl)
			networkControls["geo_access_control"] = geoAccessControlSlice

			networkControls["network_list_exceptions"] = j.NetworkControls.ExceptionSubnets
			accessControl["network_controls"] = networkControlsSlice

			// http_header
			httpHeaderSlice := make([]interface{}, 0, len(j.HeaderSpecs))
			for _, k := range j.HeaderSpecs {
				httpHeader := make(map[string]interface{})
				httpHeader["name"] = k.Name
				httpHeader["allowed"] = k.Allow
				httpHeader["values"] = k.Values
				httpHeader["effect"] = k.Effect
				httpHeader["required"] = k.Required
				httpHeaderSlice = append(httpHeaderSlice, httpHeader)
			}
			accessControl["http_header"] = httpHeaderSlice

			// file_uploads
			fileUploadsSlice := make([]interface{}, 0, 1)
			fileUploads := make(map[string]interface{})
			fileUploads["allowed_file_types"] = j.MaliciousUpload.AllowedFileTypes
			fileUploads["allowed_extensions"] = j.MaliciousUpload.AllowedExtensions
			fileUploads["effect"] = j.MaliciousUpload.Effect
			fileUploadsSlice = append(fileUploadsSlice, fileUploads)
			accessControl["file_uploads"] = fileUploadsSlice
			application["access_control"] = accessControlSlice

			// bot_protection
			botProtectionSlice := make([]interface{}, 0, 1)
			botProtection := make(map[string]interface{})

			// known_bots
			knownBotsSlice := make([]interface{}, 0, 1)
			knownBots := make(map[string]interface{})
			knownBots["search_engine_crawlers"] = j.BotProtectionSpec.KnownBotProtectionsSpec.SearchEngineCrawlers
			knownBots["business_analytics"] = j.BotProtectionSpec.KnownBotProtectionsSpec.BusinessAnalytics
			knownBots["educational"] = j.BotProtectionSpec.KnownBotProtectionsSpec.Educational
			knownBots["news"] = j.BotProtectionSpec.KnownBotProtectionsSpec.News
			knownBots["financial"] = j.BotProtectionSpec.KnownBotProtectionsSpec.Financial
			knownBots["content_feed_clients"] = j.BotProtectionSpec.KnownBotProtectionsSpec.ContentFeedClients
			knownBots["archiving"] = j.BotProtectionSpec.KnownBotProtectionsSpec.Archiving
			knownBots["career_search"] = j.BotProtectionSpec.KnownBotProtectionsSpec.CareerSearch
			knownBots["media_search"] = j.BotProtectionSpec.KnownBotProtectionsSpec.MediaSearch
			knownBotsSlice = append(knownBotsSlice, knownBots)
			botProtection["known_bots"] = knownBotsSlice

			// unknown_bots
			unknownBotsSlice := make([]interface{}, 0, 1)
			unknownBots := make(map[string]interface{})
			unknownBots["generic"] = j.BotProtectionSpec.UnknownBotProtectionSpec.Generic
			unknownBots["web_automation_tools"] = j.BotProtectionSpec.UnknownBotProtectionSpec.WebAutomationTools
			unknownBots["web_scrapers"] = j.BotProtectionSpec.UnknownBotProtectionSpec.WebScrapers
			unknownBots["api_libraries"] = j.BotProtectionSpec.UnknownBotProtectionSpec.ApiLibraries
			unknownBots["http_libraries"] = j.BotProtectionSpec.UnknownBotProtectionSpec.HttpLibraries

			// unknown_bots.request_anomalies
			requestAnomaliesSlice := make([]interface{}, 0, 1)
			requestAnomalies := make(map[string]interface{})
			requestAnomalies["effect"] = j.BotProtectionSpec.UnknownBotProtectionSpec.RequestAnomalies.Effect
			requestAnomalies["threshold"] = j.BotProtectionSpec.UnknownBotProtectionSpec.RequestAnomalies.Threshold // TODO: convert to int
			requestAnomaliesSlice = append(requestAnomaliesSlice, requestAnomalies)
			unknownBots["request_anomalies"] = requestAnomaliesSlice
			unknownBots["bot_impersonation"] = j.BotProtectionSpec.UnknownBotProtectionSpec.BotImpersonation
			unknownBotsSlice = append(unknownBotsSlice, unknownBots)
			botProtection["unknown_bots"] = unknownBotsSlice

			// bot_protection.user_defined_bots
			userDefinedBotsSlice := make([]interface{}, 0, len(j.BotProtectionSpec.UserDefinedBots))
			for _, k := range j.BotProtectionSpec.UserDefinedBots {
				userDefinedBot := make(map[string]interface{})
				userDefinedBot["name"] = k.Name
				userDefinedBot["header_name"] = k.HeaderName
				userDefinedBot["header_values"] = k.HeaderValues
				userDefinedBot["source_subnets"] = k.Subnets
				userDefinedBot["effect"] = k.Effect
				userDefinedBotsSlice = append(userDefinedBotsSlice, userDefinedBot)
			}
			botProtection["user_defined_bots"] = userDefinedBotsSlice

			// bot_protection.active_bot_detection
			activeBotDetectionSlice := make([]interface{}, 0, 1)
			activeBotDetection := make(map[string]interface{})
			activeBotDetection["session_validation_failure_effect"] = j.BotProtectionSpec.SessionValidation

			// bot_protection.active_bot_detection.javascript_based_detection
			javascriptBasedDetectionSlice := make([]interface{}, 0, 1)
			javascriptBasedDetection := make(map[string]interface{})
			javascriptBasedDetection["enabled"] = j.BotProtectionSpec.JsInjectionSpec.Enabled
			javascriptBasedDetection["browser_impersonation"] = j.BotProtectionSpec.UnknownBotProtectionSpec.BrowserImpersonation
			javascriptBasedDetection["injection_timeout_effect"] = j.BotProtectionSpec.JsInjectionSpec.TimeoutEffect
			javascriptBasedDetectionSlice = append(javascriptBasedDetectionSlice, javascriptBasedDetection)
			activeBotDetection["javascript_based_detection"] = javascriptBasedDetectionSlice

			// bot_protection.active_bot_detection.recaptcha
			recaptchaSlice := make([]interface{}, 0, 1)
			recaptcha := make(map[string]interface{})
			recaptcha["enabled"] = j.BotProtectionSpec.RecaptchaSpec.Enabled
			recaptcha["site_key"] = j.BotProtectionSpec.RecaptchaSpec.SiteKey
			recaptcha["secret_key"] = j.BotProtectionSpec.RecaptchaSpec.SecretKey
			recaptcha["type"] = j.BotProtectionSpec.RecaptchaSpec.Type
			recaptcha["every_new_session"] = j.BotProtectionSpec.RecaptchaSpec.AllSessions
			recaptcha["success_expiration"] = j.BotProtectionSpec.RecaptchaSpec.SuccessExpirationHours
			recaptchaSlice = append(recaptchaSlice, recaptcha)
			activeBotDetection["recaptcha"] = recaptchaSlice
			botProtection["active_bot_detection"] = activeBotDetectionSlice
			botProtectionSlice = append(botProtectionSlice, botProtection)
			application["bot_protection"] = botProtectionSlice

			// custom_rule
			customRuleSlice := make([]interface{}, 0, len(j.CustomRules))
			for _, k := range j.CustomRules {
				customRule := make(map[string]interface{})
				customRule["rule_id"] = k.Id
				customRule["effect"] = k.Effect
				customRuleSlice = append(customRuleSlice, customRule)
			}
			application["custom_rule"] = customRuleSlice

			// advanced_settings
			advancedSettingsSlice := make([]interface{}, 0, 1)
			advancedSettings := make(map[string]interface{})

			// session_cookies
			sessionCookiesSlice := make([]interface{}, 0, 1)
			sessionCookies := make(map[string]interface{})
			sessionCookies["enabled"] = j.SessionCookieEnabled
			sessionCookies["same_site"] = j.SessionCookieSameSite
			sessionCookies["secure"] = j.SessionCookieSecure
			sessionCookies["apply_ban_on_session"] = j.SessionCookieBan
			sessionCookies["rate_limit_on_session"] = j.DosConfig.TrackSession
			sessionCookiesSlice = append(sessionCookiesSlice, sessionCookies)
			advancedSettings["session_cookies"] = sessionCookiesSlice
			advancedSettings["ban_duration"] = j.BanDurationMinutes

			// http_body_inspection
			httpBodyInspectionSlice := make([]interface{}, 0, 1)
			httpBodyInspection := make(map[string]interface{})
			httpBodyInspection["enabled"] = !j.Body.Skip
			httpBodyInspection["inspection_size"] = j.Body.InspectionSizeBytes
			httpBodyInspectionSlice = append(httpBodyInspectionSlice, httpBodyInspection)
			advancedSettings["http_body_inspection"] = httpBodyInspectionSlice

			//custom_waas_response
			customWaasResponseSlice := make([]interface{}, 0, 1)
			customWaasResponse := make(map[string]interface{})
			customWaasResponse["enabled"] = j.CustomBlockResponse.Enabled
			customWaasResponse["status_code"] = j.CustomBlockResponse.Code
			customWaasResponse["body"] = j.CustomBlockResponse.Body
			customWaasResponseSlice = append(customWaasResponseSlice, customWaasResponse)
			advancedSettings["custom_waas_response"] = customWaasResponseSlice
			advancedSettings["enable_event_id_header"] = !j.DisableEventIdHeader
			advancedSettingsSlice = append(advancedSettingsSlice, advancedSettings)
			application["advanced_settings"] = advancedSettingsSlice
			applicationSlice = append(applicationSlice, application)
		}
		rule["application"] = applicationSlice
		rules = append(rules, rule)
	}
	return rules
}

func isPresent(item []interface{}) bool {
	if len(item) == 0 {
		return false
	}
	if item[0] == nil {
		return false
	}
	return true
}

func convertRequestAnomaliesThreshold(threshold string) int {
	switch threshold {
	case "lax":
		return 9
	case "moderate":
		return 6
	case "strict":
		return 3
	default:
		return -1
	}
}
