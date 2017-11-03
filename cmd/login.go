package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var loginCmd = &cobra.Command{
	Use:   "login [oauth-token]",
	Short: "Setup login for accessing Github",
	Run: func(cmd *cobra.Command, args []string) {
		err := RunLogin(cmd, args)
		if err != nil {
			exitWithError(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)
}

func RunLogin(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return cmd.Help()
	}
	gitOauth := args[0]
	configPath := os.Getenv("HOME")
	configName := ".github-cli"
	configType := "yaml"
	var configYaml []byte
	var madeConfigFile = false
	configFile := path.Join(configPath,
		(configName + "." + configType))
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.SetConfigType(configType)

	if _, err := os.Stat(configFile); err != nil {
		var file, err = os.Create(configFile)
		defer file.Close()
		if err != nil {
			exitWithError(fmt.Errorf("Could not create config: %s", configFile))
		}
		configYaml = []byte("git_oauth: " + gitOauth + "\n" + "show_viper_config_file: " + "false")
		defer file.Close()
		madeConfigFile = true
	} else {
		input, err := ioutil.ReadFile(configFile)
		if err != nil {
			exitWithError(err)
		}

		lines := strings.Split(string(input), "\n")

		isUpdate := false
		gitLine := "git_oauth: " + gitOauth
		for i, line := range lines {
			if strings.Contains(line, "git_oauth: ") {
				isUpdate = true
				lines[i] = gitLine
			}
		}
		output := strings.Join(lines, "\n")
		if !isUpdate {
			output = output + "\n" + gitLine + "\n"
		}
		configYaml = []byte(output)
	}

	ctx := context.Background()
	err := gc.CheckConnection(ctx, gitOauth)
	if err != nil {
		if madeConfigFile {
			err_del := os.Remove(configFile)
			if err_del != nil {
				exitWithError(fmt.Errorf("Could not delete config: %s", configFile))
			}
		}
		exitWithError(fmt.Errorf("Invalid oauth key"))
	}
	err = ioutil.WriteFile(configFile, configYaml, 0644)
	if err != nil {
		exitWithError(fmt.Errorf("Could not write config to %s", configFile))
	}

	if madeConfigFile {
		fmt.Printf("Login Succedded")
	} else {
		fmt.Printf("Login Succedded and Updated")
	}
	return nil
}
