package main

import (
	"io"

	"github.com/spf13/cobra"
)

func NewCmdGet(out io.Writer) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "get ",
		Short: "Get resource(s) from github",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	// create subcommands
	cmd.AddCommand(NewCmdGetRepos(out))
	return cmd
}
