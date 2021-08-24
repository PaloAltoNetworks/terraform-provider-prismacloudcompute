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
  name        = "example collection 1"
  description = "My first collection with Terraform"
  color       = "#00FF00"
  images      = ["prefix2*", "prefix3*"]
  labels      = ["env:development", "env:staging"]
  namespaces  = ["hamilton"]
}

resource "prismacloudcompute_collection" "example2" {
  name        = "example collection 2"
  description = "My second collection with Terraform"
  color       = "#0000FF"
  namespaces  = ["iverson"]
}

resource "prismacloudcompute_policiescomplianceciimages" "ruleset" {
  policytype = "ciImagesCompliance"
  rule {
    name   = "example ci image compliance rule"
    effect = "alert"
    collections {
      name = "All"
    }
    condition {
      compliance_check {
        id    = 41
        block = true
      }
    }
  }
}

resource "prismacloudcompute_policiescompliancecontainer" "ruleset" {
  policytype = "containerCompliance"
  rule {
    name   = "example container compliance rule"
    effect = "alert"
    collections {
      name = "All"
    }
    condition {
      compliance_check {
        id    = 531
        block = false
      }
      compliance_check {
        id    = 41
        block = true
      }
    }
  }
}

resource "prismacloudcompute_policiesruntimecontainer" "ruleset" {
  learningdisabled = false
  rule {
    name = "example container runtime rule 1"
    collections {
      name = "example collection 1"
    }
    advancedprotection       = true
    cloudmetadataenforcement = true
    kubernetesenforcement    = true
    wildfireanalysis         = "alert"
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
  rule {
    name = "example container runtime rule 2"
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

resource "prismacloudcompute_policiesvulnerabilityciimages" "ruleset" {
  policytype = "ciImagesVulnerability"
  rule {
    name   = "example ci image vulnerability rule 1"
    effect = "alert"
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
  rule {
    name   = "example ci image vulnerability rule 2"
    effect = "alert"
    collections {
      name = "All"
    }
    alertthreshold = {
      disabled = false
      value    = 9
    }
    blockthreshold = {
      enabled = false
      value   = 0
    }
  }
}

resource "prismacloudcompute_policiesvulnerabilityimages" "ruleset" {
  policytype = "containerVulnerability"
  rule {
    name   = "example image vulnerability rule 1"
    effect = "alert"
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
  rule {
    name   = "example image vulnerability rule 2"
    effect = "alert"
    collections {
      name = "All"
    }
    alertthreshold = {
      disabled = false
      value    = 9
    }
    blockthreshold = {
      enabled = false
      value   = 0
    }
  }
}
