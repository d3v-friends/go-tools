package fnGit

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"os"
	"path/filepath"
)

type Manager struct {
}

type ResetCredentialArgs struct {
	Username  string
	Email     string
	AccessKey string
}

// ResetCredential
// 주의
// 1. 기존 credential 모두 초기화 된다.
func (x *Manager) ResetCredential(i *ResetCredentialArgs) (err error) {
	var home string
	if home, err = os.UserHomeDir(); err != nil {
		return
	}

	var configFp = filepath.Join(home, ".gitconfig")
	if _, err = os.Stat(configFp); err == nil {
		_ = os.Remove(configFp)
	}

	var configFile *os.File
	if configFile, err = os.Create(configFp); err != nil {
		return
	}

	defer func() {
		_ = configFile.Close()
	}()

	var credFp = filepath.Join(home, ".git-credentials")
	if _, err = configFile.WriteString(fmt.Sprintf(
		fmtConfig,
		i.Username,
		i.Email,
		credFp,
	)); err != nil {
		return
	}

	if _, err = os.Stat(credFp); err == nil {
		if err = os.Remove(credFp); err != nil {
			return
		}
	}

	var credFile *os.File
	if credFile, err = os.Create(credFp); err != nil {
		return
	}

	defer func() {
		_ = credFile.Close()
	}()

	if _, err = credFile.WriteString(fmt.Sprintf(
		fmtCredentials,
		i.Username,
		i.AccessKey,
	)); err != nil {
		return
	}

	return
}

// CloneArgs
// repo: https://github.com/[account]/[repo]
// local: c:\\project\\abc

type CloneArgs struct {
	Repo      string
	Local     string
	Username  string
	AccessKey string
}

func (x *Manager) Clone(i *CloneArgs) (err error) {
	_, err = git.PlainClone(i.Local, false, &git.CloneOptions{
		URL: i.Repo,
		Auth: &http.BasicAuth{
			Username: i.Username,
			Password: i.AccessKey,
		},
	})
	return
}

var fmtConfig = `
[user]
	name = %s
	email = %s
[core]
	autocrlf = true
[credential]
	helper = store --file %s
`

var fmtCredentials = `
https://%s:%s@github.com
`
