package cmd

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/google/go-github/github"
)

func getFakeRepo(url string, lang string, stars int, forks int) github.Repository {
	// NEED HELP in cleaning this dirt.
	urlPtr := new(string)
	*urlPtr = url
	langPtr := new(string)
	*langPtr = lang
	starsPtr := new(int)
	*starsPtr = stars
	forksPtr := new(int)
	*forksPtr = forks
	return github.Repository{HTMLURL: urlPtr, Language: langPtr,
		StargazersCount: starsPtr, ForksCount: forksPtr}
}

func TestGetRepos(t *testing.T) {
	cases := []struct {
		url   string
		lang  string
		stars int
		forks int
	}{
		{
			url:   "example.com",
			lang:  "Java",
			stars: 5,
			forks: 20,
		},
		{
			url:   "",
			lang:  "chinese",
			stars: 0,
			forks: 0,
		},
		{
			url:   "foo.com",
			lang:  "",
			stars: 0,
			forks: 0,
		},
	}

	for _, c := range cases {
		fakeRepo := getFakeRepo(c.url, c.lang, c.stars, c.forks)
		repo := newRepo(&fakeRepo)
		if repo.URL != c.url {
			t.Fatalf("Expected URL to be %s but got %s", c.url, repo.URL)
		}
		if repo.Language != c.lang {
			t.Fatalf("Expected Language to be %s but got %s", c.lang, repo.Language)
		}
		if !reflect.DeepEqual(repo.Stars, c.stars) {
			t.Fatalf("Expected Stars to be %d but got %v", c.stars, repo.Stars)
		}
		if !reflect.DeepEqual(repo.Forks, c.forks) {
			t.Fatalf("Expected Forks to be %d but got %v", c.forks, repo.Forks)
		}
		str := repo.RepoString()
		expected := []string{c.url, c.lang, strconv.Itoa(c.stars),
			strconv.Itoa(c.forks)}
		if !reflect.DeepEqual(str, expected) {
			t.Fatalf("Expected %v but got %v", expected, str)
		}
	}

	// Test newRepo for replacing nil language with "-"
	fURL := new(string)
	*fURL = "foo"
	fStars := new(int)
	*fStars = 3
	fForks := new(int)
	*fForks = 9
	nilLangRepo := github.Repository{HTMLURL: fURL, Language: nil,
		StargazersCount: fStars, ForksCount: fForks}
	repo := newRepo(&nilLangRepo)
	if repo.Language != "-" {
		t.Fatalf("Expected Language to be - but for %s", repo.Language)
	}
}
