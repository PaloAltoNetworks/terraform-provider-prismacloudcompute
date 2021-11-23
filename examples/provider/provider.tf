terraform {
  required_providers {
    prismacloudcompute = {
      source  = "PaloAltoNetworks/prismacloudcompute"
      version = "0.1.0"
    }
  }
}

provider "prismacloudcompute" {
  # Configure provider with file
  #
  config_file = "creds.json"

  # Alternatively, you can use variables
  #
  # console_url = "https://foo.bar.com"
  # username = "myUsername"
  # password = "myPassword"
}
