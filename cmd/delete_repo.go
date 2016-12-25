package cmd

import (
	"fmt"
	"io"
	"strings"

	"github.com/alok87/github-cli/pkg/utils"
	"github.com/spf13/cobra"
)

type DeleteRepoOptions struct {
	Name      string
	IsPrivate string
}

func NewCmdDeleteRepo(out io.Writer) *cobra.Command {
	options := &DeleteRepoOptions{}

	cmd := &cobra.Command{
		Use:   "repo [name]",
		Short: "Delete repo",
		Long:  `Deletes a Github repo.`,
		Run: func(cmd *cobra.Command, args []string) {
			err := RunDeleteRepo(cmd, args, out, options)
			if err != nil {
				exitWithError(err)
			}
		},
	}

	return cmd
}

func RunDeleteRepo(cmd *cobra.Command, args []string, out io.Writer, o *DeleteRepoOptions) error {
	if len(args) != 1 {
		return cmd.Help()
	}
	repoName := args[0]
	// client := rootCommand.gclient.GetClient()
	client := gc.GetClient()
	// user := rootCommand.gclient.User
	user := gc.User
	repoUrl := user + "/" + repoName
	c := utils.AskForConfirmation("Are you sure you want to delete " + repoUrl + " ?")
	if c {
		_, err := client.Repositories.Delete(user, repoName)
		if err != nil {
			if strings.Fields(err.Error())[2] == "404" {
				exitWithError(fmt.Errorf("Repo %s does not exist", repoName))
			}
		}
		fmt.Printf("Repo %s deleted in github", repoName)
	}
	return nil
}
