package grCouponPoint

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/d3v-friends/go-tools/fn/fnPanic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlias(test *testing.T) {
	var tool = NewTestTool()
	var ctx = tool.Context()
	test.Run("create alias", func(t *testing.T) {

		var account = fnPanic.Get(CreateAccount(ctx))

		var alias = fnPanic.Get(CreateAlias(ctx, &CreateAliasArgs{
			AccountId: account.Id,
			Kind:      "username",
			Value:     gofakeit.Username(),
		}))

		assert.Equal(t, "username", alias.Kind)

	})
}
