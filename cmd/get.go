package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get ",
	Short: "Get resource(s) from github",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getReposCmd)
}
