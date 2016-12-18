package cmd

import (
	goflag "flag"
	"fmt"
	"io"
	"os"

	"github.com/alok87/github-cli/pkg/ghub"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type RootCmd struct {
	configFile   string
	cobraCommand *cobra.Command
	gclient      *ghub.Gclient
}

var rootCommand = RootCmd{
	cobraCommand: &cobra.Command{
		Use:   "github-cli",
		Short: "Use github-cli to perform github tasks.",
		Long:  "Use github-cli to perform github tasks.",
	},
	gclient: &ghub.Gclient{
		Name: "Github Client",
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	NewCmdRoot(os.Stdout)
}

func NewCmdRoot(out io.Writer) *cobra.Command {
	//options := &RootOptions{}
	cmd := rootCommand.cobraCommand

	cmd.PersistentFlags().AddGoFlagSet(goflag.CommandLine)
	cmd.PersistentFlags().StringVar(&rootCommand.configFile, "config", "", "config file (default is $HOME/.github-cli.yaml)")

	// create subcommands
	cmd.AddCommand(NewCmdLogin(out))
	cmd.AddCommand(NewCmdGet(out))
	cmd.AddCommand(NewCmdCreate(out))
	cmd.AddCommand(NewCmdDelete(out))

	return cmd
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if rootCommand.configFile != "" {
		// enable ability to specify config file via flag
		viper.SetConfigFile(rootCommand.configFile)
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

func (c *RootCmd) AddCommand(cmd *cobra.Command) {
	c.cobraCommand.AddCommand(cmd)
}

func Execute() {
	goflag.Set("logtostderr", "true")
	goflag.CommandLine.Parse([]string{})
	if err := rootCommand.cobraCommand.Execute(); err != nil {
		exitWithError(err)
	}
}

// exitWithError will terminate execution with an error result
// It prints the error to stderr and exits with a non-zero exit code
func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}
