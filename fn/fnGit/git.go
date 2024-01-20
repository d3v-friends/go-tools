package fnGit

import (
	"context"
	"errors"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"os"
)

type GitInfo interface {
	GetRepository() string
	GetLocalPath() string
	GetUsername() string
	GetAccessKey() string
}

func Clone(
	ctx context.Context,
	i GitInfo,
) (err error) {
	if _, err = os.Stat(i.GetLocalPath()); err == nil {
		if err = os.RemoveAll(i.GetLocalPath()); err != nil {
			return
		}
	}

	_, err = git.PlainCloneContext(
		ctx,
		i.GetLocalPath(),
		false,
		&git.CloneOptions{
			URL: i.GetRepository(),
			Auth: &http.BasicAuth{
				Username: i.GetUsername(),
				Password: i.GetAccessKey(),
			},
		})

	return
}

func Fetch(ctx context.Context, i GitInfo) (err error) {
	var repo *git.Repository
	if repo, err = git.PlainOpen(i.GetLocalPath()); err != nil {
		return
	}

	err = repo.FetchContext(ctx, &git.FetchOptions{
		RemoteURL: i.GetRepository(),
		Auth: &http.BasicAuth{
			Username: i.GetUsername(),
			Password: i.GetAccessKey(),
		},
	})

	if errors.Is(err, git.NoErrAlreadyUpToDate) {
		err = nil
	}

	return
}
