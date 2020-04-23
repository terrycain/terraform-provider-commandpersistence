package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/terrycain/terraform-provider-commandpersistence/commandpersistence"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: commandpersistence.Provider})
}
