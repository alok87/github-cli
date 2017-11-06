package cmd

import (
	"bytes"
	"testing"

	"github.com/google/go-github/github"
)

func TestNewRepo(t *testing.T) {
	testCases := []struct {
		name         string
		repo         *github.Repository
		wantLanguage string
	}{
		{
			name: "normal repo",
			repo: &github.Repository{
				HTMLURL:         stringPtr("repo"),
				Language:        stringPtr("Java"),
				StargazersCount: intPtr(1),
				ForksCount:      intPtr(1),
			},
			wantLanguage: "Java",
		},
		{
			name: "repo with unknown language",
			repo: &github.Repository{
				HTMLURL:         stringPtr("repo"),
				Language:        nil,
				StargazersCount: intPtr(0),
				ForksCount:      intPtr(0),
			},
			wantLanguage: "-",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := newRepo(tc.repo)

			if r.Language != tc.wantLanguage {
				t.Errorf("unexpected repo language: \n\t(GOT) %v\n\t(WNT) %v", r.Language, tc.wantLanguage)
			}
		})
	}
}

func TestRenderTable(t *testing.T) {
	testCases := []struct {
		name      string
		repos     []*github.Repository
		wantTable string
	}{
		{
			name: "empty table",
			wantTable: `
+------+----------+-------+-------+
| REPO | LANGUAGE | STARS | FORKS |
+------+----------+-------+-------+
+------+----------+-------+-------+
`,
		},
		{
			name: "single repo table",
			repos: []*github.Repository{
				{
					HTMLURL:         stringPtr("https://github.com/foo/bar"),
					Language:        stringPtr("Go"),
					StargazersCount: intPtr(5),
					ForksCount:      intPtr(3),
				},
			},
			wantTable: `
+----------------------------+----------+-------+-------+
|            REPO            | LANGUAGE | STARS | FORKS |
+----------------------------+----------+-------+-------+
| https://github.com/foo/bar | Go       |     5 |     3 |
+----------------------------+----------+-------+-------+
`,
		},
		{
			name: "multiple repo table",
			repos: []*github.Repository{
				{
					HTMLURL:         stringPtr("reponame"),
					Language:        stringPtr("Ruby"),
					StargazersCount: intPtr(0),
					ForksCount:      intPtr(0),
				},
				{
					HTMLURL:         stringPtr("reponame2"),
					Language:        nil,
					StargazersCount: intPtr(100),
					ForksCount:      intPtr(70),
				},
				{
					HTMLURL:         stringPtr("reponame3"),
					Language:        stringPtr("Javascript"),
					StargazersCount: intPtr(3),
					ForksCount:      intPtr(1),
				},
			},
			wantTable: `
+-----------+------------+-------+-------+
|   REPO    |  LANGUAGE  | STARS | FORKS |
+-----------+------------+-------+-------+
| reponame  | Ruby       |     0 |     0 |
| reponame2 |          - |   100 |    70 |
| reponame3 | Javascript |     3 |     1 |
+-----------+------------+-------+-------+
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Start with a newline to keep the wantTable pretty.
			buf := bytes.NewBufferString("\n")

			renderTable(tc.repos, buf)

			if buf.String() != tc.wantTable {
				t.Errorf("unexpected table output: \n\t(GOT) %v\n\t(WNT) %v", buf.String(), tc.wantTable)
			}
		})
	}
}

func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}
