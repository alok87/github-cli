package cmd

import (
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

// openRepoCmd represents the openRepo command
var openRepoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Open Github repo under the logged in user in a web browser",
	Long: `Opens Github repo page in a web browser. The repo url is derived
from the logged in username. Hence, a url would be of the form

	github.com/username/reponame
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get the project name from the directory name.
		cwd, err := os.Getwd()
		if err != nil {
			exitWithError(err)
		}
		projectName := filepath.Base(cwd)

		// Initialize the gc client to have Github account details.
		ctx := context.Background()
		_, err = gc.GetClient(ctx)
		if err != nil {
			exitWithError(err)
		}

		url := "https://" + path.Join("github.com", gc.User, projectName)
		fmt.Printf("Opening %s in web browser...\n", url)
		if err = browser.OpenURL(url); err != nil {
			exitWithError(err)
		}
	},
}
