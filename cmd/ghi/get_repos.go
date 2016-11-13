package main

import (
	"fmt"
	"io"
	"os"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"k8s.io/kops/util/pkg/tables"
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
		exitWithError(fmt.Errorf("err"))
	}
	if len(repos) == 0 {
		fmt.Fprintf(os.Stderr, "No repos found\n")
		return nil
	}
	t := &tables.Table{}
	t.AddColumn("REPO", func(r *github.Repository) string {
		return *r.HTMLURL
	})
	t.AddColumn("LANGUAGE", func(r *github.Repository) string {
		if r.Language != nil {
			return *r.Language
		} else {
			return "-"
		}
	})
	t.AddColumn("STARS", func(r *github.Repository) int {
		return *r.StargazersCount
	})
	t.AddColumn("FORKS", func(r *github.Repository) int {
		return *r.ForksCount
	})
	return t.Render(repos, os.Stdout, "REPO", "LANGUAGE", "STARS", "FORKS")
}
