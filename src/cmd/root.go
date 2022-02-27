package cmd

import (
	"github.com/spf13/cobra"

	"github.com/rna-vt/devicecommander/src/utilities"
)

func init() {
	cobra.OnInitialize(utilities.ConfigureEnvironment)
}

var RootCmd = &cobra.Command{
	Use:   "device-commander",
	Short: "A tool for running and managaing a device-commander network.",
	Long: `device-commander is the primary jumpoff point for running a network of devices.
This tool will provide several helpful tools for managing and running a 
network of compliant devices.`,
}

func Execute() error {
	return RootCmd.Execute()
}
