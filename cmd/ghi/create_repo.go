package main

import (
	"github.com/spf13/cobra"
	"github.com/google/go-github/github"
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
		Use:   "repo [name]",
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
	if len(args) != 1 {
		return cmd.Help()
	}
	repoName := args[0]

	client := rootCommand.gclient.GetClient()
	repo := &github.Repository{
		Name:    github.String(repoName),
		Private: github.Bool(false),
	}
	_, _, err:= client.Repositories.Create("", repo)
	if err != nil {
		return err
	}
	fmt.Printf("Repo %s created in github", repoName)
	return nil
}
