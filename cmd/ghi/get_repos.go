package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/google/go-github/github"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type GetRepoOptions struct {
	Name      string
	IsPrivate string
}

func NewCmdGetRepos(out io.Writer) *cobra.Command {
	options := &GetRepoOptions{}

	cmd := &cobra.Command{
		Use:   "repos",
		Short: "Get repo",
		Long:  `Gets the list of Github repos for the logged user`,
		Run: func(cmd *cobra.Command, args []string) {
			err := RunGetRepos(cmd, args, out, options)
			if err != nil {
				exitWithError(err)
			}
		},
	}

	return cmd
}

func RunGetRepos(cmd *cobra.Command, args []string, out io.Writer, c *GetRepoOptions) error {
	var user string
	if len(args) > 0 {
		return cmd.Help()
	}
	client := rootCommand.gclient.GetClient()
	if len(args) == 0 {
		user = rootCommand.gclient.User
	}
	opt := &github.RepositoryListOptions{Type: "all", Sort: "updated"}
	repos, _, err := client.Repositories.List(user, opt)
	if err != nil {
		exitWithError(fmt.Errorf("Error Getting Repos"))
	}
	if len(repos) == 0 {
		fmt.Fprintf(os.Stderr, "No repos found\n")
		return nil
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"REPO", "LANGUAGE", "STARS", "FORKS"})

	for _, repo := range repos {
		table.Append([]string{*repo.HTMLURL, *repo.Language,
			strconv.Itoa(*repo.StargazersCount),
			strconv.Itoa(*repo.ForksCount)})
	}

	table.Render()

	return nil
}
