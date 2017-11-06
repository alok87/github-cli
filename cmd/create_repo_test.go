package cmd

import (
	"io/ioutil"
	"os"
	"testing"

	git "gopkg.in/src-d/go-git.v4"
)

func TestAddRemote(t *testing.T) {
	testCases := []struct {
		name         string
		isGitRepo    bool
		remoteExists bool
		wantErr      error
	}{
		{
			name:      "not a git repo",
			isGitRepo: false,
			wantErr:   errAddRemoteNoRepo,
		},
		{
			name:      "a git repo",
			isGitRepo: true,
			wantErr:   nil,
		},
		{
			name:         "remote already exists",
			isGitRepo:    true,
			remoteExists: true,
			wantErr:      errAddRemoteExists,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repoName := "some-repo-name"
			username := "userxyz"

			// Create a tempdir
			tempdir, err := ioutil.TempDir("", "github-cli-test")
			if err != nil {
				t.Fatal("failed to create temporary directory")
			}

			if tc.isGitRepo {
				// Initialize a git repo.
				_, err := git.PlainInit(tempdir, false)
				if err != nil {
					t.Fatal("failed to initialize git repo")
				}

				// Add a remote.
				if tc.remoteExists {
					addRemote(username, tempdir, repoName)
				}
			}

			err = addRemote(username, tempdir, repoName)
			if err != tc.wantErr {
				t.Fatalf("unexpected error: \n\t(GOT): %v\n\t(WNT): %v", err, tc.wantErr)
			}

			// Cleanup the test assets.
			os.RemoveAll(tempdir)
		})
	}
}
