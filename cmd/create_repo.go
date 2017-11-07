package cmd

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
)

var remote bool

var (
	errAddRemoteNoRepo = errors.New("failed to add remote, current directory is not a git repo")
	errAddRemoteExists = errors.New("failed to add remote \"origin\", remote already exists")
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

func init() {
	createRepoCmd.Flags().BoolVarP(&remote, "remote", "r", false, "add a remote \"origin\" in current git repo associated with the created remote repo")
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

	if remote {
		// Get the current repo path.
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		if err := addRemote(gc.User, cwd, repoName); err != nil {
			return err
		}
		fmt.Println("Added remote \"origin\".")
	}

	return nil
}

// addRemote adds a remote to a local repo.
// username is the github username, repoPath is the path to local repo,
// repoName is the remote repo name.
func addRemote(username, repoPath, repoName string) error {
	// Load the git repo at repoPath.
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		if err == git.ErrRepositoryNotExists {
			return errAddRemoteNoRepo
		}
		return err
	}

	// Construct a remoteURL path.
	remoteURL := "https://" + path.Join("github.com", username, repoName)
	// Default refspec.
	refSpec := config.RefSpec("+refs/heads/*:refs/remotes/origin/*")

	config := config.RemoteConfig{
		Name:  "origin",
		URLs:  []string{remoteURL},
		Fetch: []config.RefSpec{refSpec},
	}
	_, err = r.CreateRemote(&config)
	if err == git.ErrRemoteExists {
		return errAddRemoteExists
	}
	return err
}
