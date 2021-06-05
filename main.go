package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/terraform-providers/terraform-provider-prismacloudcompute/prismacloudcompute"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: prismacloudcompute.Provider,
	})
}
