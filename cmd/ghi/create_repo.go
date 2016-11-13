package main

import (
	"github.com/spf13/cobra"
	"io"
    "fmt"
)

type CreateRepoOptions struct {
    Name string
    IsPrivate string
}

func NewCmdCreateRepo(out io.Writer) *cobra.Command {
    options := &CreateRepoOptions{}

	cmd := &cobra.Command{
		Use:   "repo",
		Short: "Create repo",
		Long:  `Creates a Github repo.`,
		Run: func(cmd *cobra.Command, args []string) {
			err := RunCreateRepo(cmd, args, out, options)
			if err != nil {
				exitWithError(err)
			}
		},
	}

    return cmd
}

func RunCreateRepo(cmd *cobra.Command, args []string, out io.Writer, c *CreateRepoOptions) error {
	client := rootCommand.gclient.GetClient()
	repos, _, _ := client.Repositories.List("", nil)
	fmt.Println(repos)
	return nil
}
