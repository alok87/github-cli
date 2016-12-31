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

var ConfigName = ".github-cli"
var ConfigType = "yaml"

func (c *Gclient) SetClient() {
	var gitOauth string

	// TODO: Find a better approach than intializing config in this way
	viper.SetConfigName(ConfigName)
	viper.AddConfigPath("$HOME")
	viper.SetConfigType(ConfigType)

	if err := viper.ReadInConfig(); err == nil {
		if viper.IsSet("git_oauth") {
			gitOauth = viper.GetString("git_oauth")
		}
	}

	if len(gitOauth) == 0 {
		fmt.Println("Login required !")
		binary, lookErr := exec.LookPath("github-cli")
		if lookErr != nil {
			exitWithError(fmt.Errorf("Binary github-cli not found"))
		}
		args := []string{"github-cli", "login"}
		env := os.Environ()
		execErr := syscall.Exec(binary, args, env)
		if execErr != nil {
			exitWithError(fmt.Errorf("Error running github-cli login"))
		}
		exitWithError(fmt.Errorf("Login required"))
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gitOauth},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	c.client = github.NewClient(tc)
}

func (c *Gclient) GetClient() *github.Client {
	c.SetClient()
	c.SetUser()
	return c.client
}

func (c *Gclient) SetUser() string {
	currentUser, _, _ := c.client.Users.Get("")
	c.User = *currentUser.Login
	return c.User
}

func (c *Gclient) CheckConnection(gitOauth string) error {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gitOauth},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)
	_, _, err := client.Users.Get("")
	return err
}

// exitWithError will terminate execution with an error result
// It prints the error to stderr and exits with a non-zero exit code
func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}
