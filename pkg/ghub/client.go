package ghub

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
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

func (c *Gclient) setClient(ctx context.Context) error {
	var gitOauth string

	// TODO: Find a better approach than intializing config in this way
	viper.SetConfigName(ConfigName)
	viper.SetConfigType(ConfigType)
	viper.AddConfigPath("$HOME")

	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "unable to reading config")
	}

	if viper.IsSet("git_oauth") {
		gitOauth = viper.GetString("git_oauth")
	}

	if len(gitOauth) == 0 {
		fmt.Println("Login required!")

		// TODO: Replace syscalls with function calls.
		binary, lookErr := exec.LookPath("github-cli")
		if lookErr != nil {
			return errors.New("binary github-cli not found")
		}
		args := []string{"github-cli", "login"}
		env := os.Environ()
		execErr := syscall.Exec(binary, args, env)
		if execErr != nil {
			return errors.New("error running github-cli login")
		}
		return errors.New("login required")
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gitOauth},
	)
	tc := oauth2.NewClient(ctx, ts)
	c.client = github.NewClient(tc)

	return nil
}

// GetClient returns ghub client.
func (c *Gclient) GetClient(ctx context.Context) (*github.Client, error) {
	if err := c.setClient(ctx); err != nil {
		return nil, err
	}

	if err := c.setUser(ctx); err != nil {
		return nil, err
	}

	return c.client, nil
}

func (c *Gclient) setUser(ctx context.Context) error {
	currentUser, _, err := c.client.Users.Get(ctx, "")
	if err != nil {
		return errors.Wrap(err, "unable to get user")
	}
	c.User = *currentUser.Login

	return nil
}

// CheckConnection checks the connection to github with an OAuth token.
func (c *Gclient) CheckConnection(ctx context.Context, gitOauth string) error {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gitOauth},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	_, _, err := client.Users.Get(ctx, "")
	return err
}
