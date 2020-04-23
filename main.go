package main

import (
    "github.com/hashicorp/terraform/plugin"
    "github.com/terrycain/terraform-provider-external2/external"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: external.Provider})
}
