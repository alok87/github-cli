package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

// CreateRepoOptions holds options for creating a repo.
type CreateRepoOptions struct {
	Name      string
	IsPrivate string
}

var createRepoOptions = &CreateRepoOptions{}
var createRepoCmd = &cobra.Command{
	Use:   "repo [name]",
	Short: "Create repo",
	Long:  `Creates a Github repo.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := runCreateRepo(cmd, args, createRepoOptions)
		if err != nil {
			exitWithError(err)
		}
	},
}

func runCreateRepo(cmd *cobra.Command, args []string, c *CreateRepoOptions) error {
	if len(args) != 1 {
		return cmd.Help()
	}
	repoName := args[0]

	ctx := context.Background()
	client := gc.GetClient(ctx)
	repo := &github.Repository{
		Name:    github.String(repoName),
		Private: github.Bool(false),
	}
	_, _, err := client.Repositories.Create(ctx, "", repo)
	if err != nil {
		if strings.Fields(err.Error())[2] == "422" {
			exitWithError(fmt.Errorf("Repo %s already exists", repoName))
		}
		exitWithError(err)
	}
	fmt.Printf("Repo %s created in github", repoName)
	return nil
}
