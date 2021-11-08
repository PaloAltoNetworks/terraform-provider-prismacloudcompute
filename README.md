# Terraform Provider for Prisma Cloud Compute by Palo Alto Networks
You can find the Prisma Cloud Compute provider in the [Terraform Registry](https://registry.terraform.io/providers/PaloAltoNetworks/prismacloudcompute/latest).

If you're interested in developing the provider, see below for a basic setup guide.

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
