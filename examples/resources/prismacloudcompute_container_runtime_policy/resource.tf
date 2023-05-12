# This resource will work with the following provider stored in the following directory path
# mkdir -p ~/.terraform.d/plugins/paloaltonetworks.com/prismacloud/prismacloudcompute/0.7.1-release/darwin_amd64
# mv terraform-provider-prismacloudcompute ~/.terraform.d/plugins/paloaltonetworks.com/prismacloud/prismacloudcompute/0.7.1-release/darwin_amd64

terraform {
  required_providers {
    prismacloudcompute = {
      source  = "paloaltonetworks.com/prismacloud/prismacloudcompute"
      version = "0.7.1-release"
    }
  }
}

provider "prismacloudcompute" {
  console_url = ""
  username    = ""
  password    = ""
}

resource "prismacloudcompute_container_runtime_policy" "ruleset" {
  learning_disabled = false
  rule {
    name                              = "string"
    collections                       = ["string"]
    advanced_protection_effect        = true
    cloud_metadata_enforcement_effect = false
    previous_name                     = "string" # Required if Renaming the Rule
    skip_exec_sessions                = false    # true | false
    wildfire_analysis                 = "alert"  # "block" | "prevent" | "alert" | "disable"
    custom_rule {
      id     = 0
      action = "string"
      effect = "string" # "allow" | "ban" | "block" | "prevent" | "alert" | "disable"
    }
    custom_rule {
      id     = 1
      action = "string"
      effect = "string" # "allow" | "ban" | "block" | "prevent" | "alert" | "disable"
    }
    dns {
      default_effect = "alert" # "block" | "prevent" | "alert" | "disable"
      disabled       = true
      domain_list {
        allowed = ["0.0.0.0"]
        denied  = ["1.1.1.1"]
        effect  = "disable"
      }
    }
    filesystem {
      allowed_list          = ["string"]
      backdoor_files_effect = "disable" # "block" | "prevent" | "alert" | "disable"
      default_effect        = "alert"   # "block" | "prevent" | "alert" | "disable"
      denied_list {
        effect = "disable" # "block" | "prevent" | "alert" | "disable"
        paths  = ["string"]
      }
      disabled                      = true
      encrypted_binaries_effect     = "disable"
      new_files_effect              = "disable"
      suspicious_elf_headers_effect = "disable"
    }
    kubernetes_enforcement = false
    network {
      allowed_ips       = ["0.0.0.0"]
      default_effect    = "alert"
      denied_ips        = ["1.1.1.1"]
      denied_ips_effect = "disable"
      disabled          = true
      listening_ports {
        allowed {
          deny  = true
          end   = 333
          start = 222
        }
        denied {
          deny  = true
          end   = 5000
          start = 4000
        }
        denied {
          deny  = true
          end   = 222
          start = 111
        }
        effect = "disable" # "block" | "prevent" | "alert" | "disable"
      }
      modified_proc_effect = "disable" # "block" | "prevent" | "alert" | "disable"
      outbound_ports {
        allowed {
          deny  = true
          end   = 300
          start = 200
        }
        denied {
          deny  = true
          end   = 6000
          start = 5000
        }
        denied {
          deny  = true
          end   = 222
          start = 111
        }
        effect = "disable" # "block" | "prevent" | "alert" | "disable"
      }
      port_scan_effect   = "disable" # "block" | "prevent" | "alert" | "disable"
      raw_sockets_effect = "disable" # "block" | "prevent" | "alert" | "disable"
    }
    processes {
      modified_process_effect = "disable" # "block" | "prevent" | "alert" | "disable"
      crypto_miners_effect    = "disable" # "block" | "prevent" | "alert" | "disable"
      lateral_movement_effect = "disable" # "block" | "prevent" | "alert" | "disable"
      reverse_shell_effect    = "disable" # "block" | "prevent" | "alert" | "disable"
      suid_binaries_effect    = "disable" # "block" | "prevent" | "alert" | "disable"
      default_effect          = "alert"   # "block" | "prevent" | "alert" | "disable"
      check_parent_child      = false
      allowed_list            = []
      disabled                = false
      denied_list {
        effect = "disable" # "block" | "prevent" | "alert" | "disable"
        paths  = ["test"]
      }
    }
  }
}