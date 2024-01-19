package fnGit

import (
	"errors"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"os"
)

type CloneArgs struct {
	Repo      string
	Local     string
	Username  string
	AccessKey string
}

func Clone(i *CloneArgs) (_ *git.Repository, err error) {
	if _, err = os.Stat(i.Local); err == nil {
		if err = os.RemoveAll(i.Local); err != nil {
			return
		}
	}

	return git.PlainClone(i.Local, false, &git.CloneOptions{
		URL: i.Repo,
		Auth: &http.BasicAuth{
			Username: i.Username,
			Password: i.AccessKey,
		},
	})
}

type FetchArgs struct {
	Repo      string
	Local     string
	Username  string
	AccessKey string
}

func Fetch(i *FetchArgs) (err error) {
	var repo *git.Repository
	if repo, err = git.PlainOpen(i.Local); err != nil {
		return
	}

	err = repo.Fetch(&git.FetchOptions{
		RemoteURL: i.Repo,
		Auth: &http.BasicAuth{
			Username: i.Username,
			Password: i.AccessKey,
		},
	})

	if errors.Is(err, git.NoErrAlreadyUpToDate) {
		err = nil
	}

	return
}
