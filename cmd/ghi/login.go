package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCmdLogin(out io.Writer) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "login [oauth-token]",
		Short: "Setup login for accessing Github",
		Run: func(cmd *cobra.Command, args []string) {
			err := RunLogin(cmd, args, out)
			if err != nil {
				exitWithError(err)
			}
		},
	}

	return cmd
}

func RunLogin(cmd *cobra.Command, args []string, out io.Writer) error {
	if len(args) != 1 {
		return cmd.Help()
	}
	gitOauth := args[0]
	configPath := "$HOME"
	configName := ".ghi"
	configType := "yaml"
	var isConfig = true
	configFile := path.Join(os.Getenv(configPath),
		(configName + "." + configType))

	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.SetConfigType(configType)

	if _, err := os.Stat(configFile); err != nil {
		isConfig = false
		var file, err = os.Create(configFile)
		if err != nil {
			exitWithError(fmt.Errorf("Could not create config: %s", configFile))
		}
		default_yaml := []byte("git_oauth: " + gitOauth + "\n" + "show_viper_config_file: " + "false")
		err = ioutil.WriteFile(configFile, default_yaml, 0644)
		if err != nil {
			exitWithError(fmt.Errorf("Could not write config to %s", configFile))
		}
		defer file.Close()
	} else {
		input, err := ioutil.ReadFile(configFile)
		if err != nil {
			exitWithError(err)
		}

		lines := strings.Split(string(input), "\n")

		for i, line := range lines {
			if strings.Contains(line, "git_oauth: ") {
				lines[i] = "git_oauth: " + gitOauth
			}
		}
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(configFile, []byte(output), 0644)
		if err != nil {
			exitWithError(err)
		}
	}
	rootCommand.gclient.SetClient()
	rootCommand.gclient.GetClient()
	if isConfig {
		fmt.Printf("Login Updated")
	} else {
		fmt.Printf("Login Succedded")
	}
	return nil
}
