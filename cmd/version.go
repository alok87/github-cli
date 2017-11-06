package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show github-cli version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:", version)
		fmt.Println("Rev:", gitSha)
		fmt.Println("Built:", buildDate)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
