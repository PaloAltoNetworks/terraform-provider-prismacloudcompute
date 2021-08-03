terraform {
  required_providers {
    prismacloudcompute = {
      source  = "hashicorp/prismacloudcompute"
      version = "~>1.0.0"
    }
  }
}

provider "prismacloudcompute" {
  json_config_file = "creds.json"
}

#resource "prismacloudcompute_collection" "example1" {
#  name   = "example collection 1"
#  color  = "#FF0000"
#  appids = [ "app1" ]
#  coderepos = [ "coderepo1", "prefix1*" ]
#  images = [ "prefix2*", "prefix3*" ]
#  labels = [ "env:development", "env:staging" ]
#  namespaces = [ "hamilton" ]
#}

#resource "prismacloudcompute_policiesruntimecontainer" "example2" {
#    learningdisabled = true
#    rules {
#        name = "my-rule"
#	collections {
#		name = "All"
#	}
#        processes = {
#            effect = "alert"
#        }
#        network = {
#            effect = "alert"
#        }
#        dns = {
#            effect = "alert"
#        }
#        filesystem = {
#            effect = "alert"
#        }
#        wildfireanalysis = "alert"
#    }
#}

resource "prismacloudcompute_policiesvulnerabilityimages" "example3" {
    policytype = "containerVulnerability"
    rules {
        name = "my-rule"
	collections {
		name = "All"
	}
        alertthreshold = {
                disabled = false
                value = 4
        }
        blockthreshold = {
                enabled = false
                value = 0
        }
    }
}

resource "prismacloudcompute_policiescompliancecontainer" "example4" {
    policytype = "containerCompliance"
    rules {
        name = "my-rule"
        effect = "alert"
	collections {
		name = "All"
	}
        alertthreshold = {
                disabled = false
                value = 4
        }
        blockthreshold = {
                enabled = false
                value = 0
        }
   }
}
