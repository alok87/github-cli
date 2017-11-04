package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/alok87/github-cli/pkg/utils"
	"github.com/spf13/cobra"
)

// DeleteRepoOptions holds options for deleting a repo.
type DeleteRepoOptions struct {
	Name      string
	IsPrivate string
}

var deleteRepoOptions = &DeleteRepoOptions{}
var deleteRepoCmd = &cobra.Command{
	Use:   "repo [name]",
	Short: "Delete repo",
	Long:  `Deletes a Github repo.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := runDeleteRepo(cmd, args, deleteRepoOptions)
		if err != nil {
			exitWithError(err)
		}
	},
}

func runDeleteRepo(cmd *cobra.Command, args []string, o *DeleteRepoOptions) error {
	if len(args) != 1 {
		return cmd.Help()
	}
	repoName := args[0]

	ctx := context.Background()
	client, err := gc.GetClient(ctx)
	if err != nil {
		return err
	}
	user := gc.User
	repoURL := user + "/" + repoName
	c := utils.AskForConfirmation("Are you sure you want to delete " + repoURL + " ?")
	if c {
		_, err := client.Repositories.Delete(ctx, user, repoName)
		if err != nil {
			if strings.Fields(err.Error())[2] == "404" {
				exitWithError(fmt.Errorf("Repo %s does not exist", repoName))
			}
		}
		fmt.Printf("Repo %s deleted in github.\n", repoName)
	}
	return nil
}
