package cmd

import (
	goflag "flag"
	"fmt"
	"os"

	"github.com/alok87/github-cli/pkg/ghub"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var gitSha string
var version string
var buildDate string

var gc = &ghub.Gclient{Name: "Github Client"}
var configFile string

// RootCmd is the github-cli root command.
var RootCmd = &cobra.Command{
	Use:   "github-cli",
	Short: "Use github-cli to perform github tasks.",
	Long:  "Use github-cli to perform github tasks.",
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().AddGoFlagSet(goflag.CommandLine)
	RootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.github-cli.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if configFile != "" {
		// enable ability to specify config file via flag
		viper.SetConfigFile(configFile)
	}

	viper.SetConfigName(".github-cli") // name of config file (without extension)
	viper.AddConfigPath("$HOME")       // adding home directory as first search path
	viper.AutomaticEnv()               // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// Print config file name based on `show_vuper_config_file`.
		if viper.IsSet("show_viper_config_file") {
			if viper.GetBool("show_viper_config_file") {
				fmt.Println("Using config file:", viper.ConfigFileUsed())
			}
		} else {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
}

// Execute runs the root command.
func Execute() {
	goflag.Set("logtostderr", "true")
	goflag.CommandLine.Parse([]string{})
	if err := RootCmd.Execute(); err != nil {
		exitWithError(err)
	}
}

// exitWithError will terminate execution with an error result
// It prints the error to stderr and exits with a non-zero exit code
func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}
