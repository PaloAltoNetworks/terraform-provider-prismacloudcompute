# This points to the local provider built according to the README.md.
terraform {
  required_providers {
    prismacloudcompute = {
      source  = "paloaltonetworks.com/prismacloud/compute"
      version = "~> 0.0.1"
    }
  }
}

provider "prismacloudcompute" {
  json_config_file = "creds.json"
}

resource "prismacloudcompute_collection" "example1" {
  name       = "example collection 1"
  color      = "#FF0000"
  appids     = ["app1"]
  coderepos  = ["coderepo1", "prefix1*"]
  images     = ["prefix2*", "prefix3*"]
  labels     = ["env:development", "env:staging"]
  namespaces = ["hamilton"]
}

resource "prismacloudcompute_policiesruntimecontainer" "example2" {
  learningdisabled = false
  rules {
    name = "example container runtime rule"
    collections {
      name = "All"
    }
    wildfireanalysis = "alert"
    processes = {
      "effect" : "alert"
    }
    network = {
      "effect" : "alert"
    }
    dns = {
      "effect" : "disable"
    }
    filesystem = {
      effect = "alert"
    }
  }
}

resource "prismacloudcompute_policiesvulnerabilityimages" "example3" {
  policytype = "containerVulnerability"
  rules {
    name = "example image vulnerability rule"
    collections {
      name = "All"
    }
    alertthreshold = {
      disabled = false
      value    = 4
    }
    blockthreshold = {
      enabled = false
      value   = 0
    }
  }
}


# resource "prismacloudcompute_policiescompliancecontainer" "example4" {
#   policytype = "containerCompliance"
#   rules {
#     name   = "example container compliance rule"
#     effect = "alert"
#     collections {
#       name = "All"
#     }
#     condition = {
#       vulnerabilities = {
#         id    = 531
#         block = false
#       }
#     }
#   }
# }
