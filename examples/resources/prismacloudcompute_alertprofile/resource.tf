resource "prismacloudcompute_alertprofile" "test" {
  name               = "webhook-example"
  enabled            = true
  alert_profile_type = "webhook"

  enable_immediate_vulnerabilities_alerts = false

  alert_profile_config {
    prisma_cloud_integration_id = prismacloud_integration.this.id
    webhook_url = "https://webhook.url"
    custom_json = <<-EOT
                {
                  "types": "#type",
                  "time": "#time",
                  "container": "#container",
                  "image": "#image",
                  "host": "#host",
                  "fqdn": "#fqdn",
                  "function": "#function",
                  "region": "#region",
                  "runtime": "#runtime",
                  "appID": "#appID",
                  "rule": "#rule",
                  "message": "#message",
                  "aggregated": "#aggregated",
                  "rest": "#rest",
                  "forensics": "#forensics",
                  "accountID": "#accountID",
                  "cluster": "#cluster",
                  "labels": #labels,
                  "collections": #collections,
                  "complianceIssues": #complianceIssues,
                  "vulnerabilities": #vulnerabilities
                }
            EOT
  }

  alert_triggers {
    access {
      enabled   = true
      all_rules = true
    }

    admission {
      enabled   = true
      all_rules = true
    }

    app_embedded_defender_runtime {
      enabled   = true
      all_rules = true
    }

    cloud_native_network_firewall {
      enabled = true
    }

    container_and_image_compliance {
      enabled   = true
      all_rules = true
    }

    container_runtime {
      enabled   = true
      all_rules = true
    }

    defender_health {
      enabled = true
    }

    host_compliance {
      enabled   = true
      all_rules = true
    }

    host_runtime {
      enabled   = true
      all_rules = true
    }

    host_vulnerabilities {
      enabled   = true
      all_rules = true
    }
    image_vulnerabilities {
      enabled   = true
      all_rules = false
      rules     = ["rule1", "rule2"]
    }

    incidents {
      enabled = true
    }

    kubernetes_audits {
      enabled   = true
      all_rules = true
    }

    serverless_runtime {
      enabled   = true
      all_rules = true
    }

    waas_firewall_app_embedded_defender {
      enabled   = true
      all_rules = true
    }

    waas_firewall_container {
      enabled   = true
      all_rules = true
    }

    waas_firewall_host {
      enabled   = true
      all_rules = true
    }

    waas_firewall_serverless {
      enabled   = true
      all_rules = true
    }

    waas_health {
      enabled = true
    }

  }
}
