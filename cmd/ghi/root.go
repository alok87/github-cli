package main

import (
	goflag "flag"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/alok87/github-cli/pkg/ghub"
	"io"
	"os"
)

type RootCmd struct {
	configFile   string
	cobraCommand *cobra.Command
	gclient *ghub.Gclient
}

var rootCommand = RootCmd{
	cobraCommand: &cobra.Command{
		Use:   "ghi",
		Short: "ghi is github-cli to perform github tasks.",
		Long:  `ghi is github-cli to perform github tasks.`,
	},
	gclient: &ghub.Gclient{
		Name: "Github Client",
	},
}

func init() {
	rootCommand.gclient.SetClient()
	cobra.OnInitialize(initConfig)
	NewCmdRoot(os.Stdout)
}

func NewCmdRoot(out io.Writer) *cobra.Command {
	//options := &RootOptions{}

	cmd := rootCommand.cobraCommand

	cmd.PersistentFlags().AddGoFlagSet(goflag.CommandLine)
	cmd.PersistentFlags().StringVar(&rootCommand.configFile, "config", "", "config file (default is $HOME/.ghi.yaml)")

	// create subcommands
	cmd.AddCommand(NewCmdCreate(out))
	// cmd.AddCommand(NewLoginCommand(out))

	return cmd
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if rootCommand.configFile != "" {
		// enable ability to specify config file via flag
		viper.SetConfigFile(rootCommand.configFile)
	}

	viper.SetConfigName(".ghi")  // name of config file (without extension)
	viper.AddConfigPath("$HOME") // adding home directory as first search path
	viper.AutomaticEnv()         // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
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
