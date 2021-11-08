resource "prismacloudcompute_host_runtime_policy" "ruleset" {
  rule {
    name        = "Default - alert on suspicious runtime behavior"
    collections = ["All"]
    activities {
      disabled                   = false
      docker_enabled             = false
      readonly_docker_enabled    = false
      service_activities_enabled = false
      sshd_enabled               = false
      sudo_enabled               = false
    }
    antimalware {
      allowed_processes = []
      crypto_miners     = "alert"
      custom_feed       = "alert"
      denied_processes {
        effect = "alert"
        paths  = []
      }
      encrypted_binaries            = "alert"
      execution_flow_hijack         = "alert"
      intelligence_feed             = "alert"
      reverse_shell                 = "alert"
      service_unknown_origin_binary = "alert"
      suspicious_elf_headers        = "alert"
      temp_filesystem_processes     = "alert"
      user_unknown_origin_binary    = "alert"
      webshell                      = "alert"
      wildfire_analysis             = "alert"
    }
    dns {
      allowed           = []
      denied            = []
      deny_effect       = "disable"
      intelligence_feed = "disable"
    }
    network {
      allowed_outbound_ips = []
      custom_feed          = "alert"
      denied_outbound_ips  = []
      deny_effect          = "alert"
      intelligence_feed    = "alert"
    }
  }
}
