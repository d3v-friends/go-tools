package grCouponPoint

import (
	"github.com/d3v-friends/go-tools/fn/fnPanic"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoupon(test *testing.T) {
	var tool = NewTestTool()
	var ctx = tool.Context()
	test.Run("create coupon", func(t *testing.T) {
		var account = fnPanic.Get(CreateAccount(ctx))

		var coupon = fnPanic.Get(CreateCoupon(ctx, &CreateCouponArgs{
			AccountId: account.Id,
			Currency:  decimal.NewFromInt(10000),
			Price:     decimal.NewFromInt(10),
			Point:     decimal.NewFromInt(1000),
		}))

		assert.Equal(t, "10000", coupon.Currency.String())
		assert.Equal(t, "10", coupon.Price.String())
		assert.Equal(t, "1000", coupon.TotalPoint.String())

	})
}
