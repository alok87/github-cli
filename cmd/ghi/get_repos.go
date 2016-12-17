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

type Repo struct {
	URL      string
	Language string
	Stars    int
	Forks    int
}

func (r Repo) RepoString() []string {
	return []string{r.URL, r.Language,
		strconv.Itoa(r.Stars), strconv.Itoa(r.Forks)}
}

var newRepo = func(r *github.Repository) Repo {
	var lang string
	if r.Language != nil {
		lang = *r.Language
	} else {
		lang = "-"
	}
	return Repo{*r.HTMLURL, lang, *r.StargazersCount, *r.ForksCount}
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

var getRepos = func() ([]*github.Repository, error) {
	client := rootCommand.gclient.GetClient()
	// User should be fetched only after the above client init, else user remains
	// empty.
	user := rootCommand.gclient.User
	opt := &github.RepositoryListOptions{Type: "all", Sort: "updated"}
	repos, _, err := client.Repositories.List(user, opt)
	if err != nil {
		return nil, err
	}

	if len(repos) == 0 {
		fmt.Fprintf(os.Stderr, "No repos found\n")
	}

	return repos, nil
}

func RunGetRepos(cmd *cobra.Command, args []string, out io.Writer, c *GetRepoOptions) error {
	if len(args) > 0 {
		return cmd.Help()
	}
	repos, err := getRepos()
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"REPO", "LANGUAGE", "STARS", "FORKS"})

	for _, repo := range repos {
		r := newRepo(repo)
		table.Append(r.RepoString())
	}

	table.Render()

	return nil
}
