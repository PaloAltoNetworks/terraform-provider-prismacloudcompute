terraform {
  required_providers {
    prismacloudcompute = {
      source  = "prismacloudcompute"
      version = "~> 1.0"
    }
  }
}

provider "prismacloudcompute" {
  json_config_file = "creds.json"
}

/*resource "prismacloudcompute_collection" "example1" {
    name = "New Collection"
    color = "#FF0000"
}*/

/*resource "prismacloudcompute_policiesruntimecontainer" "example2" {
    learningdisabled = true
    rules {
        name = "my-rule"
	collections {
		name = "All"
	}
        processes = {
            effect = "alert"
        }
        network = {
            effect = "alert"
        }
        dns = {
            effect = "alert"
        }
        filesystem = {
            effect = "alert"
        }
        wildfireanalysis = "alert"
    }
}*/

resource "prismacloudcompute_policiesvulnerabilityimages" "example3" {
    policytype = "containerVulnerability"
    rules {
        name = "my-rule"
	collections {
		name = "All"
	}
    }
}
