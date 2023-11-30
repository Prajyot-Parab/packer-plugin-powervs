package main

import (
	"fmt"
	"os"

	"github.com/ppc64le-cloud/packer-plugin-powervs/builder/powervs"
	powervsData "github.com/ppc64le-cloud/packer-plugin-powervs/datasource/powervs"
	powervsPP "github.com/ppc64le-cloud/packer-plugin-powervs/post-processor/powervs"
	powervsProv "github.com/ppc64le-cloud/packer-plugin-powervs/provisioner/powervs"
	powervsVersion "github.com/ppc64le-cloud/packer-plugin-powervs/version"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder(plugin.DEFAULT_NAME, new(powervs.Builder))
	pps.RegisterProvisioner(plugin.DEFAULT_NAME, new(powervsProv.Provisioner))
	pps.RegisterPostProcessor(plugin.DEFAULT_NAME, new(powervsPP.PostProcessor))
	pps.RegisterDatasource(plugin.DEFAULT_NAME, new(powervsData.Datasource))
	pps.SetVersion(powervsVersion.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
