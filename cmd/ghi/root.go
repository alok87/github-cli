package main

import (
	goflag "flag"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"os"
)

type RootCmd struct {
	configFile   string
	cobraCommand *cobra.Command
}

var rootCommand = RootCmd{
	cobraCommand: &cobra.Command{
		Use:   "ghi",
		Short: "ghi is github-cli to perform github tasks.",
		Long:  `ghi is github-cli to perform github tasks.`,
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
	cmd.PersistentFlags().StringVar(&rootCommand.configFile, "config", "", "config file (default is $HOME/.ghi.yaml)")

	// defaultStateStore := os.Getenv("KOPS_STATE_STORE")

	// create subcommands
	// cmd.AddCommand(NewCmdCreate(f, out))
	// cmd.AddCommand(NewCmdEdit(f, out))
	// cmd.AddCommand(NewCmdUpdate(f, out))

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
		fmt.Println("Using config file:", viper.ConfigFileUsed())
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
