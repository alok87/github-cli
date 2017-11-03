package ghub

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/google/go-github/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

// Gclient is ghub client struct.
type Gclient struct {
	Name   string
	client *github.Client
	User   string
}

// ConfigName is the configuration file name.
var ConfigName = ".github-cli"

// ConfigType is the configuration file type.
var ConfigType = "yaml"

func (c *Gclient) setClient(ctx context.Context) {
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
	tc := oauth2.NewClient(ctx, ts)
	c.client = github.NewClient(tc)
}

// GetClient returns ghub client.
func (c *Gclient) GetClient(ctx context.Context) *github.Client {
	c.setClient(ctx)
	c.setUser(ctx)
	return c.client
}

func (c *Gclient) setUser(ctx context.Context) string {
	currentUser, _, _ := c.client.Users.Get(ctx, "")
	c.User = *currentUser.Login
	return c.User
}

// CheckConnection checks the connection to github.
func (c *Gclient) CheckConnection(ctx context.Context, gitOauth string) error {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gitOauth},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	_, _, err := client.Users.Get(ctx, "")
	return err
}

// exitWithError will terminate execution with an error result
// It prints the error to stderr and exits with a non-zero exit code
func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}
