terraform {
  required_providers {
    prismacloudcompute = {
      source  = "PaloAltoNetworks/prismacloudcompute"
      version = "0.1.0"
    }
  }
}

provider "prismacloudcompute" {
  config_file = "creds.json"
}
