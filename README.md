# Terraform Provider for Prisma Cloud Compute by Palo Alto Networks

---
**WORK IN PROGRESS. NOT READY FOR USE.**

---
<a href="https://www.terraform.io"><img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px"></a>

## Requirements
- [Go](https://golang.org/doc/install) (only if building the provider)
- [Terraform](https://www.terraform.io/downloads.html)

## Building the provider
0. Set `$GOPATH` if not already set.
    ```bash
    export GOPATH=$(go env GOPATH)
    ```
1. Clone this repository and navigate to its directory.
    ```bash
    git clone git@github.com:PaloAltoNetworks/terraform-provider-prismacloudcompute.git $GOPATH/src/github.com/terraform-providers/terraform-provider-prismacloudcompute && cd $_
    ```
2. Build the provider.
    ```bash
    make build
    ```
3. For local testing, symlink the resultant binary to the appropriate location.
    ```bash
    # macOS-specific path; adjust as necessary
    mkdir -p ~/.terraform.d/plugins/paloaltonetworks.com/prismacloud/compute/0.0.1/darwin_amd64/ && ln -fs ~/go/bin/terraform-provider-prismacloudcompute ~/.terraform.d/plugins/paloaltonetworks.com/prismacloud/compute/0.0.1/darwin_amd64/terraform-provider-compute_v0.0.1
    ```
4. Point your terraform file to this local plugin.
    ```terraform
    terraform {
      required_providers {
        prismacloudcompute = {
          source  = "paloaltonetworks.com/prismacloud/compute"
          version = "~> 0.0.1"
        }
      }
    }
    ```

## Developing the provider
See Makefile for available `make` targets.
