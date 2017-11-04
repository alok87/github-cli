package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/alok87/github-cli/pkg/ghub"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout github-cli's current github session.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := deleteConfig(); err != nil {
			exitWithError(err)
		}
		fmt.Println("Logged out")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}

// deleteFile is an abstraction over `os.Remove` for deleting a file.
// Abstraction was created for testability of `deleteConfig`.
var deleteFile = func(filepath string) error {
	return os.Remove(filepath)
}

// deleteConfig deletes github-cli config file, which is present at the home
// directory of users.
func deleteConfig() error {
	homepath, err := homedir.Dir()
	if err != nil {
		return err
	}
	configpath := path.Join(homepath, ghub.ConfigName) + "." + ghub.ConfigType
	fmt.Printf("Deleting config file %s\n", configpath)
	return deleteFile(configpath)
}
