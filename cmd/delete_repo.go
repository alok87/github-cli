package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/alok87/github-cli/pkg/utils"
	"github.com/spf13/cobra"
)

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
		err := RunDeleteRepo(cmd, args, deleteRepoOptions)
		if err != nil {
			exitWithError(err)
		}
	},
}

func RunDeleteRepo(cmd *cobra.Command, args []string, o *DeleteRepoOptions) error {
	if len(args) != 1 {
		return cmd.Help()
	}
	repoName := args[0]

	ctx := context.Background()
	client := gc.GetClient(ctx)
	user := gc.User
	repoUrl := user + "/" + repoName
	c := utils.AskForConfirmation("Are you sure you want to delete " + repoUrl + " ?")
	if c {
		_, err := client.Repositories.Delete(ctx, user, repoName)
		if err != nil {
			if strings.Fields(err.Error())[2] == "404" {
				exitWithError(fmt.Errorf("Repo %s does not exist", repoName))
			}
		}
		fmt.Printf("Repo %s deleted in github", repoName)
	}
	return nil
}
