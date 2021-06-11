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

resource "prismacloudcompute_collection" "example" {
    name = "New Collection"
    color = "#FF0000"
}

/*resource "prismacloudcompute_policiesruntimecontainer" "example" {
    name = "My Policy"
    policy_type = "network"
    rule {
        name = "my rule"
        criteria = "savedSearchId"
        parameters = {
            "savedSearch": "true",
            "withIac": "false",
        }
        rule_type = "Network"
    }
}*/
