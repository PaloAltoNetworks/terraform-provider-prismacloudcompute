# Terraform Provider for Prisma Cloud Compute
You can find the Prisma Cloud Compute provider in the [Terraform Registry](https://registry.terraform.io/providers/PaloAltoNetworks/prismacloudcompute/latest).

## Basic setup
```terraform
terraform {
  required_providers {
    prismacloudcompute = {
      source  = "PaloAltoNetworks/prismacloudcompute"
      version = "0.4.0"
    }
  }
}

provider "prismacloudcompute" {
  # Configure provider with file
  #
  config_file = "creds.json"

  # Alternatively, you can use variables
  #
  # console_url = "https://console.example.com"
  # username = "myUsername"
  # password = "myPassword"
}
```
Complete documentation can be found in the [marketplace listing](https://registry.terraform.io/providers/PaloAltoNetworks/prismacloudcompute/latest/docs).

## Contributing
Contributions are welcome!
Please read the [contributing guide](CONTRIBUTING.md) for more information.

## Support
Please read our [support document](SUPPORT.md) for details on how to get support for this project.
