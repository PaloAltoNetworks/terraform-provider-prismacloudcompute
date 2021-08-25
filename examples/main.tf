# This points to the local provider built according to the README.md.
terraform {
  required_providers {
    prismacloudcompute = {
      source  = "paloaltonetworks.com/prismacloud/prismacloudcompute"
      version = "~> 0.0.1"
    }
  }
}

provider "prismacloudcompute" {
  json_config_file = "creds.json"
}
/*
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
    name = "example container runtime rule 1"
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
  rules {
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

resource "prismacloudcompute_policiescompliancecontainer" "example4" {
  policytype = "containerCompliance"
  rules {
    name   = "example container compliance rule"
    effect = "alert"
    collections {
      name = "All"
    }
    condition = {
      vulnerabilities = "[{\"id\": 531, \"block\": false}]"
    }
  }
}

resource "prismacloudcompute_policiesvulnerabilityciimages" "example5" {
  policytype = "ciImagesVulnerability"
  rules {
    effect = "alert"
    name = "example ci image vulnerability rule"
    collections {
      name = "All"
    }
    alertthreshold = {
      value    = 1
      disabled = false
    }
    blockthreshold = {
      value   = 0
      enabled = false
    }
  }
}

resource "prismacloudcompute_policiescomplianceciimages" "example6" {
  policytype = "ciImagesCompliance"
  rules {
    effect = "alert"
    condition = {
      vulnerabilities = "[{\"id\": 41, \"block\": false}, {\"id\": 422, \"block\": false},{\"id\": 424, \"block\": false}, {\"id\": 425, \"block\": false}, {\"id\": 426, \"block\": false}, {\"id\": 448,\"block\": false}, {\"id\": 5041, \"block\": false}]"
    }
    name = "example ci image compliance rule"
    collections {
      name = "All"
    }
  }
}
*/

resource "prismacloudcompute_policiesruntimehost" "example7" {
  rules {
    name = "example host runtime rule 1"
    collections {
      name = "All"
    }
    forensic = {
       activitiesdisabled = false
       sshdenabled = false
       sudoenabled = false
       serviceactivitiesenabled = false
       dockerenabled = false
       readonlydockerenabled = false
    }
    network {
       denylisteffect = "alert"
       customfeed = "alert"
       intelligencefeed = "alert"
    }
    dns {
       denylisteffect = "disable"
       intelligencefeed = "disable"
    }
    antimalware = {
       deniedprocesses = "{\"effect\": \"alert\"}"
       cryptominer = "alert"
       serviceunknownoriginbinary = "alert"
       userunknownoriginbinary = "alert"
       encryptedbinaries = "alert"
       suspiciouselfheaders = "alert"
       tempfsproc = "alert"
       reverseshell = "alert"
       webshell = "alert"
       executionflowhijack = "alert"
       customfeed = "alert"
       intelligencefeed = "alert"
       wildfireanalysis = "alert"
    }
  }  
}

/*
resource "prismacloudcompute_policiesvulnerabilityhost" "example8" {
  policytype = "hostVulnerability"
  rules {
    name = "example host vulnerability rule 1"
    effect = "alert"
    action = tolist(["*"])
    condition = {
       vulnerabilities = ""
    }
    blockmsg = ""
    principal = tolist([])
    group = tolist(["*"])
    verbose = false
    allcompliance = false
    onlyfixed = false
    cverules {}
    tags {}
    collections {
      name = "All"
    }
    alertthreshold = {
      value    = 1
      disabled = false
    }
    blockthreshold = {
      value   = 0
      enabled = false
    }
  }
}

resource "prismacloudcompute_policiescompliancehost" "example9" {
  policytype = "hostCompliance"
  rules {
    name = "example host compliance rule 1"
    effect = "alert"
    action = tolist(["*"])
    condition = {
       vulnerabilities = "[{\"id\": 11, \"block\": false}, {\"id\": 111, \"block\": true}]"
    }
    blockmsg = ""
    verbose = false
    collections {
      name = "All"
    }
  }
}
*/
