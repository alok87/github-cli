package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

func NewCmdCreate(out io.Writer) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "create ",
		Short: "Create a resource in github",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	// create subcommands
	cmd.AddCommand(NewCmdCreateRepo(out))
	return cmd
}
