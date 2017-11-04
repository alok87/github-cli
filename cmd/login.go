package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/alok87/github-cli/pkg/ghub"
	"github.com/pkg/errors"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login [oauth-token]",
	Short: "Setup login for accessing Github",
	Run: func(cmd *cobra.Command, args []string) {
		err := runLogin(cmd, args)
		if err != nil {
			exitWithError(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func runLogin(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return cmd.Help()
	}

	var configYaml []byte

	// madeConfigFile tells if a new config file is created. A new config is
	// created when no previous config file is found.
	var madeConfigFile = false

	gitOauth := args[0]
	configFile := path.Join(os.Getenv("HOME"), (ghub.ConfigName + "." + ghub.ConfigType))

	if _, err := os.Stat(configFile); err != nil {
		// Create the config file if it doesn't exist.
		var file, err = os.Create(configFile)
		if err != nil {
			return errors.Wrapf(err, "could not create config: %s", configFile)
		}
		file.Close()

		// Create the config.
		// TODO: Do with a proper yaml parser/writer.
		configYaml = []byte("git_oauth: " + gitOauth + "\n" + "show_viper_config_file: " + "false")
		madeConfigFile = true
	} else {
		// Read the existing config file.
		input, err := ioutil.ReadFile(configFile)
		if err != nil {
			return errors.Wrapf(err, "could not read config: %s", configFile)
		}

		// Split the config file content into a slice of configs.
		lines := strings.Split(string(input), "\n")

		// An update in this case is addition of a new oauth key. If the key
		// already exists, this process is an update, else, we need to append
		// the config with the provided key.
		isUpdate := false

		// "git_oauth" property string.
		gitLine := "git_oauth: " + gitOauth

		// Check if the config already consists of "git_oauth" property, if so,
		// mark it as an update.
		// TODO: Handle this with a go-yaml parser.
		for i, line := range lines {
			if strings.Contains(line, "git_oauth: ") {
				isUpdate = true
				lines[i] = gitLine
			}
		}
		output := strings.Join(lines, "\n")

		// If it isn't an update to existing config file, add the "git_oauth"
		// property.
		if !isUpdate {
			output = output + "\n" + gitLine + "\n"
		}
		configYaml = []byte(output)
	}

	ctx := context.Background()
	err := gc.CheckConnection(ctx, gitOauth)
	if err != nil {
		if madeConfigFile {
			// Connection failed with the new config. Delete the newly created
			// config file.
			// TODO: This could be avoided by performing a connection check
			// initially before creating the config file.
			if err := os.Remove(configFile); err != nil {
				return errors.Wrapf(err, "could not delete config: %s", configFile)
			}
		}
		return errors.Wrap(err, "invalid oauth key")
	}

	// Write the config
	err = ioutil.WriteFile(configFile, configYaml, 0644)
	if err != nil {
		errors.Wrapf(err, "could not write config to %s", configFile)
	}

	if madeConfigFile {
		fmt.Println("Logged in successfully.")
	} else {
		fmt.Println("Logged in and config updated.")
	}
	return nil
}
