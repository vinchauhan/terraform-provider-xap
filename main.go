package main

import(
	"github.com/vinchauhan/terraform-provider-xap/xap"
	"github.com/hashicorp/terraform/plugin"
)
func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: xap.Provider})
}