package cmd

import (
	"fmt"
	"strings"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"github.com/kubernetes/editor"
)

type CreateIssueOptions struct {
	Title     *string   `json:"title,omitempty"`
    Body      *string   `json:"body,omitempty"`
    Labels    *[]string `json:"labels,omitempty"`
    Assignee  *string   `json:"assignee,omitempty"`
    State     *string   `json:"state,omitempty"`
    Milestone *int      `json:"milestone,omitempty"`
    Assignees *[]string `json:"assignees,omitempty"`
}

var createIssueOptions = &CreateIssueOptions{}
var createissueCmd = &cobra.Command{
	Use:   "issue [repo_name]",
	Short: "Create issue",
	Long:  `Creates a issue.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := RunCreateIssue(cmd, args, createIssueOptions)
		if err != nil {
			exitWithError(err)
		}
	},
}

func RunCreateIssue(cmd *cobra.Command, args []string, c *CreateIssueOptions) error {
	if len(args) != 1 {
		return cmd.Help()
	}
	repoInput := strings.Split(args[0], "/")
	if len(repoInput) == 1 {
		repoName := repoInput
		repoOrg := "alok87"
	}
	repoName := args[0]


	// client := rootCommand.gclient.GetClient()
	client := gc.GetClient()
	repo := &github.Repository{
		Name:    github.String(repoName),
		Private: github.Bool(false),
	}
	_, _, err := client.Repositories.Create("", repo)
	if err != nil {
		if strings.Fields(err.Error())[2] == "422" {
			exitWithError(fmt.Errorf("Repo %s already exists", repoName))
		}
		exitWithError(err)
	}
	fmt.Printf("Repo %s created in github", repoName)
	return nil
}
