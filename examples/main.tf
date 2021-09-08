# This points to the local provider built according to the README.md.
terraform {
  required_providers {
    prismacloudcompute = {
      source  = "paloaltonetworks.com/prismacloud/prismacloudcompute"
      version = "0.0.1"
    }
  }
}

provider "prismacloudcompute" {
  config_file = "creds.json"
}

# resource "prismacloudcompute_collection" "example1" {
#   name        = "example collection 1"
#   description = "My first collection with Terraform"
#   color       = "#00FF00"
#   images      = ["prefix2*", "prefix3*"]
#   labels      = ["env:development", "env:staging"]
#   namespaces  = ["hamilton"]
# }

# resource "prismacloudcompute_collection" "example2" {
#   name        = "example collection 2"
#   description = "My second collection with Terraform"
#   color       = "#0000FF"
#   namespaces  = ["iverson"]
# }

resource "prismacloudcompute_policiescomplianceciimages" "ruleset" {
  rule {
    name        = "example ci image compliance rule 1"
    effect      = "alert"
    collections = ["All"]
    conditions {
      compliance_check {
        block = true
        id    = 41
      }
      compliance_check {
        block = false
        id    = 422
      }
    }
  }
  rule {
    name        = "example ci image compliance rule 2"
    effect      = "alert"
    collections = ["All"]
    conditions {}
  }
}

resource "prismacloudcompute_policiescompliancecontainer" "ruleset" {
  rule {
    name        = "example container compliance rule 1"
    effect      = "alert"
    collections = ["All"]
    conditions {
      compliance_check {
        block = true
        id    = 41
      }
      compliance_check {
        block = false
        id    = 422
      }
    }
  }
  rule {
    name        = "example container compliance rule 2"
    effect      = "alert"
    collections = ["All"]
    conditions {}
  }
}

resource "prismacloudcompute_policiescompliancehost" "ruleset" {
  rule {
    name        = "example host compliance rule 1"
    effect      = "alert, block"
    collections = ["All"]
    conditions {
      compliance_check {
        block = true
        id    = 41
      }
      compliance_check {
        block = false
        id    = 422
      }
    }
  }
  rule {
    name        = "example host compliance rule 2"
    effect      = "alert"
    collections = ["All"]
    conditions {}
  }
}

resource "prismacloudcompute_policiesruntimecontainer" "ruleset" {
  learningdisabled = false
  rule {
    name                     = "example container runtime rule 1"
    collections              = ["All"]
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
    name             = "example container runtime rule 2"
    collections      = ["All"]
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

resource "prismacloudcompute_policiesruntimehost" "ruleset" {
  rule {
    name        = "example host runtime rule 1"
    collections = ["All"]
    antimalware = {
      cryptominer                = "alert"
      customfeed                 = "alert"
      deniedprocesses            = "{\"effect\": \"alert\"}"
      encryptedbinaries          = "alert"
      executionflowhijack        = "alert"
      intelligencefeed           = "alert"
      reverseshell               = "alert"
      serviceunknownoriginbinary = "alert"
      suspiciouselfheaders       = "alert"
      tempfsproc                 = "alert"
      userunknownoriginbinary    = "alert"
      webshell                   = "alert"
      wildfireanalysis           = "alert"
    }
    dns = {
      denylisteffect   = "disable"
      intelligencefeed = "disable"
    }
    forensic = {
      activitiesdisabled       = false
      dockerenabled            = false
      readonlydockerenabled    = false
      serviceactivitiesenabled = false
      sshdenabled              = false
      sudoenabled              = false
    }
    network = {
      customfeed       = "alert"
      denylisteffect   = "alert"
      intelligencefeed = "alert"
    }
  }
  rule {
    name        = "example host runtime rule 2"
    collections = ["All"]
    antimalware = {
      cryptominer                = "disable"
      customfeed                 = "alert"
      deniedprocesses            = "{\"effect\": \"alert\"}"
      encryptedbinaries          = "alert"
      executionflowhijack        = "alert"
      intelligencefeed           = "alert"
      reverseshell               = "alert"
      serviceunknownoriginbinary = "alert"
      suspiciouselfheaders       = "alert"
      tempfsproc                 = "alert"
      userunknownoriginbinary    = "alert"
      webshell                   = "alert"
      wildfireanalysis           = "alert"
    }
    dns = {
      denylisteffect   = "disable"
      intelligencefeed = "disable"
    }
    forensic = {
      activitiesdisabled       = false
      dockerenabled            = false
      readonlydockerenabled    = false
      serviceactivitiesenabled = false
      sshdenabled              = false
      sudoenabled              = false
    }
    network = {
      customfeed       = "alert"
      denylisteffect   = "alert"
      intelligencefeed = "alert"
    }
  }
}

resource "prismacloudcompute_policiesvulnerabilityciimages" "ruleset" {
  rule {
    name        = "example ci image vulnerability rule 1"
    effect      = "alert"
    collections = ["All"]
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
    name        = "example ci image vulnerability rule 2"
    effect      = "alert"
    collections = ["All"]
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

resource "prismacloudcompute_policiesvulnerabilityhost" "ruleset" {
  rule {
    name        = "example host vulnerability rule 1"
    effect      = "alert"
    collections = ["All"]
    alertthreshold = {
      disabled = false
      value    = 4
    }
  }
  rule {
    name        = "example host vulnerability rule 2"
    effect      = "alert"
    collections = ["All"]
    alertthreshold = {
      disabled = false
      value    = 9
    }
  }
}

resource "prismacloudcompute_policiesvulnerabilityimages" "ruleset" {
  rule {
    name        = "example image vulnerability rule 1"
    effect      = "alert"
    collections = ["All"]
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
    name        = "example image vulnerability rule 2"
    effect      = "alert"
    collections = ["All"]
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
