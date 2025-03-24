package cmd

import "github.com/spf13/cobra"

var rootCMD = &cobra.Command{
	Use:   "dpf",
	Short: "dpf is a tool that allows you to find duplicates in some env-files",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func NewCommand() *cobra.Command {
	return rootCMD
}
