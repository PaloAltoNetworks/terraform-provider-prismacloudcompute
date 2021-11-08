resource "prismacloudcompute_container_runtime_policy" "ruleset" {
  learning_disabled = false
  rule {
    name                       = "Default - alert on suspicious runtime behavior"
    collections                = ["All"]
    advanced_protection        = true
    cloud_metadata_enforcement = false
    dns {
      allowed     = []
      denied      = []
      deny_effect = "disable"
    }
    filesystem {
      allowed                 = []
      backdoor_files          = true
      check_new_files         = true
      denied                  = []
      deny_effect             = "alert"
      skip_encrypted_binaries = false
      suspicious_elf_headers  = true
    }
    kubernetes_enforcement = false
    network {
      allowed_outbound_ips    = []
      denied_outbound_ips     = []
      deny_effect             = "alert"
      detect_port_scan        = true
      skip_modified_processes = false
      skip_raw_sockets        = false
    }
    processes {
      allowed                = []
      check_crypto_miners    = true
      check_lateral_movement = true
      denied                 = []
      deny_effect            = "alert"
    }
    wildfire_analysis = "alert"
  }
}
