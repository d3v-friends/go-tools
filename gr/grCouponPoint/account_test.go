package grCouponPoint

import (
	"github.com/d3v-friends/go-tools/fn/fnPanic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccount(test *testing.T) {
	var tool = NewTestTool()

	test.Run("create account", func(t *testing.T) {
		var ctx = tool.Context()

		var account = fnPanic.Get(CreateAccount(ctx))

		assert.Equal(t, 0, len(account.CouponList))

	})
}
