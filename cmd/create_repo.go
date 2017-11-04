package cmd

import (
	"context"
	"fmt"

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
	client, err := gc.GetClient(ctx)
	if err != nil {
		return err
	}

	repo := &github.Repository{
		Name:    github.String(repoName),
		Private: github.Bool(false),
	}
	_, _, err = client.Repositories.Create(ctx, "", repo)
	if err != nil {
		// Convert error to github.ErrorResponse. Since "repo already exists" is
		// a custom response, e.Code is "custom", which doesn't tell the exact
		// reason. Hence, compare the error message to ensure it's an already
		// existing repo error.
		// https://developer.github.com/v3/#client-errors
		e := err.(*github.ErrorResponse).Errors
		if len(e) > 0 {
			if e[0].Message == "name already exists on this account" {
				return fmt.Errorf("repo %s already exists", repoName)
			}
		}
		return err
	}
	fmt.Printf("Repo %s created in github.\n", repoName)
	return nil
}
