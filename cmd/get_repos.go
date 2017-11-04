package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/google/go-github/github"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

const defaultReposPerPage = 10

var reposPerPage int
var headers = []string{"REPO", "LANGUAGE", "STARS", "FORKS"}

// GetRepoOptions holds options for fetching repo.
type GetRepoOptions struct {
	Name      string
	IsPrivate string
}

// Repo is struct of a repo that is used in get_repos table view.
type Repo struct {
	URL      string
	Language string
	Stars    int
	Forks    int
}

// RepoString converts all the attributes of Repo into an array of strings.
// Used in appending table in tablewriter.
func (r Repo) RepoString() []string {
	return []string{r.URL, r.Language,
		strconv.Itoa(r.Stars), strconv.Itoa(r.Forks)}
}

var getRepoOptions = &GetRepoOptions{}
var getReposCmd = &cobra.Command{
	Use:   "repos",
	Short: "Get repo",
	Long:  `Gets the list of Github repos for the logged user`,
	Run: func(cmd *cobra.Command, args []string) {
		err := runGetRepos(cmd, args, getRepoOptions)
		if err != nil {
			exitWithError(err)
		}
	},
}

func init() {
	getReposCmd.Flags().IntVarP(&reposPerPage, "limit", "l",
		defaultReposPerPage, "number of repos to fetch")
}

// newRepo creates a new Repo object, given a github repository.
var newRepo = func(r *github.Repository) Repo {
	var lang string
	if r.Language != nil {
		lang = *r.Language
	} else {
		lang = "-"
	}
	return Repo{*r.HTMLURL, lang, *r.StargazersCount, *r.ForksCount}
}

// getRepos fetches and returns all the repos of the logged in user.
var getRepos = func() ([]*github.Repository, error) {
	ctx := context.Background()
	client, err := gc.GetClient(ctx)
	if err != nil {
		return nil, err
	}
	user := gc.User
	opt := &github.RepositoryListOptions{
		Type: "all", Sort: "updated",
		ListOptions: github.ListOptions{PerPage: reposPerPage}}
	repos, _, err := client.Repositories.List(ctx, user, opt)
	if err != nil {
		return nil, err
	}

	if len(repos) == 0 {
		fmt.Fprintf(os.Stderr, "No repos found\n")
	}

	return repos, nil
}

// renderTable extracts the given github repos and renders a table using the
// extracted data.
var renderTable = func(repos []*github.Repository) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)

	for _, repo := range repos {
		r := newRepo(repo)
		table.Append(r.RepoString())
	}

	table.Render()
}

func runGetRepos(cmd *cobra.Command, args []string, c *GetRepoOptions) error {
	if len(args) > 0 {
		return cmd.Help()
	}
	repos, err := getRepos()
	if err != nil {
		return err
	}

	renderTable(repos)

	return nil
}
