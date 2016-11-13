package ghub

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/google/go-github/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

type Gclient struct {
	Name   string
	client *github.Client
	User   string
}

func (c *Gclient) SetClient() {
	var gitOauth string
	configName := ".ghi"
	configType := "yaml"

	// TODO: Find a better approach than intializing config in this way
	viper.SetConfigName(configName)
	viper.AddConfigPath("$HOME")
	viper.SetConfigType(configType)

	if err := viper.ReadInConfig(); err == nil {
		if viper.IsSet("git_oauth") {
			gitOauth = viper.GetString("git_oauth")
		}
	}
	if len(gitOauth) == 0 {
		fmt.Println("Login required !")
		binary, lookErr := exec.LookPath("ghi")
		if lookErr != nil {
			exitWithError(fmt.Errorf("Binary ghi not found"))
		}
		args := []string{"ghi", "login"}
		env := os.Environ()
		execErr := syscall.Exec(binary, args, env)
		if execErr != nil {
			exitWithError(fmt.Errorf("Error running ghi login"))
		}
		exitWithError(fmt.Errorf("Login required"))
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gitOauth},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	c.client = github.NewClient(tc)
	currentUser, _, _ := c.client.Users.Get("")
	c.User = *currentUser.Login
}

func (c *Gclient) GetClient() *github.Client {
	c.SetClient()
	return c.client
}

// exitWithError will terminate execution with an error result
// It prints the error to stderr and exits with a non-zero exit code
func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "\n%v\n", err)
	os.Exit(1)
}
