package fnGit

import (
	"github.com/d3v-friends/go-tools/fn/fnEnv"
	"github.com/d3v-friends/go-tools/fn/fnPanic"
	"testing"
)

func TestManager(test *testing.T) {
	fnPanic.On(fnEnv.ReadFromFile(".env"))
	var manager = new(Manager)
	test.Run("reset credential", func(t *testing.T) {
		var err = manager.ResetCredential(&ResetCredentialArgs{
			Username:  fnEnv.Read("GIT_USERNAME"),
			Email:     fnEnv.Read("GIT_EMAIL"),
			AccessKey: fnEnv.Read("GIT_ACCESS_KEY"),
		})

		if err != nil {
			t.Fatal(err)
		}
	})

	test.Run("clone", func(t *testing.T) {
		var err = manager.Clone(&CloneArgs{
			Repo:      fnEnv.Read("GIT_REPO"),
			Local:     fnEnv.Read("GIT_LOCAL"),
			Username:  fnEnv.Read("GIT_USERNAME"),
			AccessKey: fnEnv.Read("GIT_ACCESS_KEY"),
		})

		if err != nil {
			t.Error(err)
		}

	})
}
