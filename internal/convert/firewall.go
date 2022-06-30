package convert

import (
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/policy"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/waas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaToFirewallAppPolicy(d *schema.ResourceData) (policy.FirewallAppPolicy, error) {
	policy := policy.FirewallAppPolicy{}

	if val, ok := d.GetOk("max_port"); ok {
		policy.MaxPort = val.(int)
	}

	if val, ok := d.GetOk("min_port"); ok {
		policy.MinPort = val.(int)
	}

	if schemaRules, ok := d.GetOk("rules"); ok {
		iterRules := schemaRules.([]interface{})

		for _, schemaRule := range iterRules {
			mapRule := schemaRule.(map[string]interface{})
			rule := waas.WaasRule{}

			if val, ok := mapRule["allow_malformed_http_header_names"]; ok {
				rule.AllowMalformedHttpHeaderNames = val.(bool)
			}

			if schemaApplications, ok := mapRule["applications"]; ok {
				iterApplications := schemaApplications.([]interface{})

				for _, schemaApplication := range iterApplications {
					mapApplication := schemaApplication.(map[string]interface{})
					application := waas.WaasApplicationSpec{}

					if schemaAPIs, ok := mapApplication["api"]; ok {
						iterAPIs := schemaAPIs.([]interface{})

						if len(iterAPIs) == 1 {
							mapAPI := iterAPIs[0].(map[string]interface{})
							api := waas.WaasAPISpec{}

							if val, ok := mapAPI["description"]; ok {
								api.Description = val.(string)
							}

							if val, ok := mapAPI["effect"]; ok {
								api.Effect = val.(string)
							}

							if schemaEndpoints, ok := mapAPI["endpoints"]; ok {
								iterEndpoints := schemaEndpoints.([]interface{})
								for _, schemaEndpoint := range iterEndpoints {
									mapEndpoint := schemaEndpoint.(map[string]interface{})
									endpoint := waas.WaasEndpoint{}

									if val, ok := mapEndpoint["base_path"]; ok {
										endpoint.BasePath = val.(string)
									}

									if val, ok := mapEndpoint["exposed_port"]; ok {
										endpoint.ExposedPort = val.(int)
									}

									if val, ok := mapEndpoint["grpc"]; ok {
										endpoint.GRPC = val.(bool)
									}

									if val, ok := mapEndpoint["host"]; ok {
										endpoint.Host = val.(string)
									}

									if val, ok := mapEndpoint["http2"]; ok {
										endpoint.HTTP2 = val.(bool)
									}

									if val, ok := mapEndpoint["internal_port"]; ok {
										endpoint.InternalPort = val.(int)
									}

									if val, ok := mapEndpoint["tls"]; ok {
										endpoint.TLS = val.(bool)
									}

									api.Endpoints = append(api.Endpoints, endpoint)
								}
							}

							if val, ok := mapAPI["fallback_effect"]; ok {
								api.FallbackEffect = val.(string)
							}

							if schemaPaths, ok := mapAPI["paths"]; ok {
								iterPaths := schemaPaths.([]interface{})
								for _, schemaPath := range iterPaths {
									mapPath := schemaPath.(map[string]interface{})
									path := waas.WaasPath{}

									if schemaMethods, ok := mapPath["methods"]; ok {
										iterMethods := schemaMethods.([]interface{})
										for _, schemaMethod := range iterMethods {
											mapMethod := schemaMethod.(map[string]interface{})
											method := waas.WaasMethod{}

											if val, ok := mapMethod["method"]; ok {
												method.Method = val.(string)
											}

											if schemaParameters, ok := mapMethod["parameters"]; ok {
												iterParameters := schemaParameters.([]interface{})
												for _, schemaParameter := range iterParameters {
													mapParameter := schemaParameter.(map[string]interface{})
													parameter := waas.WaasParam{}

													if val, ok := mapParameter["allow_empty_value"]; ok {
														parameter.AllowEmptyValue = val.(bool)
													}

													if val, ok := mapParameter["array"]; ok {
														parameter.Array = val.(bool)
													}

													if val, ok := mapParameter["explode"]; ok {
														parameter.Explode = val.(bool)
													}

													if val, ok := mapParameter["location"]; ok {
														parameter.Location = val.(string)
													}

													if val, ok := mapParameter["max"]; ok {
														parameter.Max = val.(int)
													}

													if val, ok := mapParameter["min"]; ok {
														parameter.Min = val.(int)
													}

													if val, ok := mapParameter["name"]; ok {
														parameter.Name = val.(string)
													}

													if val, ok := mapParameter["required"]; ok {
														parameter.Required = val.(bool)
													}

													if val, ok := mapParameter["style"]; ok {
														parameter.Style = val.(string)
													}

													if val, ok := mapParameter["type"]; ok {
														parameter.Type = val.(string)
													}

													method.Parameters = append(method.Parameters, parameter)
												}
											}

											path.Methods = append(path.Methods, method)
										}
									}

									if val, ok := mapPath["path"]; ok {
										path.Path = val.(string)
									}

									api.Paths = append(api.Paths, path)
								}
							}

							if val, ok := mapAPI["query_param_effect"]; ok {
								api.QueryParamFallbackEffect = val.(string)
							}

							application.APISpec = api
						}
					}

					if val, ok := mapApplication["app_id"]; ok {
						application.AppID = val.(string)
					}

					if schemaAttackTools, ok := mapApplication["attack_tools"]; ok {
						iterAttackTools := schemaAttackTools.([]interface{})

						if len(iterAttackTools) == 1 {
							mapAttackTools := iterAttackTools[0].(map[string]interface{})
							attackTools := waas.WaasProtectionConfig{}

							if val, ok := mapAttackTools["effect"]; ok {
								attackTools.Effect = val.(string)
							}

							if schemaAttackToolsExceptions, ok := mapAttackTools["exceptions"]; ok {
								iterAttackToolsExceptions := schemaAttackToolsExceptions.([]interface{})

								for _, schemaAttackException := range iterAttackToolsExceptions {
									mapAttackToolsException := schemaAttackException.(map[string]interface{})
									attackToolsException := waas.WaasExceptionField{}

									if val, ok := mapAttackToolsException["key"]; ok {
										attackToolsException.Key = val.(string)
									}

									if val, ok := mapAttackToolsException["location"]; ok {
										attackToolsException.Location = val.(string)
									}

									attackTools.ExceptionFields = append(attackTools.ExceptionFields, attackToolsException)
								}
							}

							application.AttackTools = attackTools
						}
					}

					if val, ok := mapApplication["ban_duration"]; ok {
						application.BanDurationMinutes = val.(int)
					}

				}
			}
		}
	}

	return policy, nil

}
