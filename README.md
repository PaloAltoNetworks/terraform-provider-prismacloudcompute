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
1. Fetch the repository and navigate to its directory.
    ```bash
    go get github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute
    cd ~/go/src/github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute
    ```
2. Install the provider.
This also moves the compiled binary to the appropriate location.
    ```bash
    # macOS-specific OS_ARCH; adjust as necessary
    make install OS_ARCH=darwin_amd64 VERSION=0.0.1
    ```
4. Point your terraform file to this local plugin.
    ```terraform
    terraform {
      required_providers {
        prismacloudcompute = {
          source  = "paloaltonetworks.com/prismacloud/prismacloudcompute"
          version = "0.0.1"
        }
      }
    }
    ```

## Developing the provider
See Makefile for available `make` targets.
