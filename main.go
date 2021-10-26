package main

import (
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/prismacloudcompute"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: prismacloudcompute.Provider,
	})
}
