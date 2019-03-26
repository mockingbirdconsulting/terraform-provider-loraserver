package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/mockingbirdconsulting/terraform-provider-loraserver/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
