package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/terraform-providers/terraform-provider-prismacloudcompute/prismacloudcompute"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: prismacloudcompute.Provider,
	})
}
