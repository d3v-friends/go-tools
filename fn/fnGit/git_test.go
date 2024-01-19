package fnGit

import (
	"github.com/d3v-friends/go-tools/fn/fnEnv"
	"github.com/d3v-friends/go-tools/fn/fnPanic"
	"testing"
)

func TestManager(test *testing.T) {
	fnPanic.On(fnEnv.ReadFromFile(".env"))
	test.Run("clone", func(t *testing.T) {
		var err error
		if _, err = Clone(&CloneArgs{
			Repo:      fnEnv.Read("GIT_REPO"),
			Local:     fnEnv.Read("GIT_LOCAL"),
			Username:  fnEnv.Read("GIT_USERNAME"),
			AccessKey: fnEnv.Read("GIT_ACCESS_KEY"),
		}); err != nil {
			t.Error(err)
		}

		if err = Fetch(&FetchArgs{
			Repo:      fnEnv.Read("GIT_REPO"),
			Local:     fnEnv.Read("GIT_LOCAL"),
			Username:  fnEnv.Read("GIT_USERNAME"),
			AccessKey: fnEnv.Read("GIT_ACCESS_KEY"),
		}); err != nil {
			t.Error(err)
		}
	})
}
