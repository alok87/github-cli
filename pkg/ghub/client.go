package ghub

import (
    "github.com/google/go-github/github"
    "golang.org/x/oauth2"
    "github.com/spf13/viper"
    "fmt"
    "os"
)

type Gclient struct {
    Name string
    client *github.Client
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
    if len(gitOauth) == 0  {
        exitWithError(fmt.Errorf("Missing git_oauth in configfile %s.%s",
                                 configName, configType))
    }
    ts := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: gitOauth},
    )
    tc := oauth2.NewClient(oauth2.NoContext, ts)
    c.client = github.NewClient(tc)
}

func (c *Gclient) GetClient() *github.Client {
    return c.client
}

// exitWithError will terminate execution with an error result
// It prints the error to stderr and exits with a non-zero exit code
func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "\n%v\n", err)
	os.Exit(1)
}
