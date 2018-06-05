package main

import (
	"github.com/andrewgithub.com/logicmonitor/terraform-provider-helm/terraform-provider-helm/pkg/provider"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return provider.Provider()
		},
	})
}
