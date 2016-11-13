package main

import (
	"io"

	"github.com/spf13/cobra"
)

func NewCmdDelete(out io.Writer) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "delete ",
		Short: "Delete a resource in github",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	// create subcommands
	cmd.AddCommand(NewCmdDeleteRepo(out))
	return cmd
}
