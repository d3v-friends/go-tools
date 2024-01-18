package grCouponPoint

import (
	"context"
	"fmt"
	"github.com/d3v-friends/go-tools/fn/fnPanic"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sync"
	"testing"
)

func TestRequest(test *testing.T) {
	var tool = NewTestTool(true)
	var ctx = tool.Context()
	var db = tool.DB
	test.Run("request", func(t *testing.T) {
		var account = fnPanic.Get(CreateAccount(db))
		var _ = fnPanic.Get(CreateCoupon(db, &CreateCouponArgs{
			AccountId: account.Id,
			Currency:  decimal.NewFromInt(10000),
			Price:     decimal.NewFromInt(10),
			Point:     decimal.NewFromInt(1000),
		}))

		fnPanic.On(CreateRequest(ctx, &CreateRequestArgs{
			AccountId: account.Id,
			Point:     decimal.NewFromInt(20),
			Fn: func(ctx context.Context, bal *Balance) (err error) {
				return
			},
		}))

		var balance = fnPanic.Get(FindBalanceOne(ctx, &FindBalanceOneArgs{
			AccountId: &account.Id,
		}))

		assert.Equal(t, "980", balance.Point.String())
		assert.Equal(t, "9800", balance.Currency.String())
	})

	test.Run("shortage", func(t *testing.T) {

		var account = fnPanic.Get(CreateAccount(db))
		var err = CreateRequest(ctx, &CreateRequestArgs{
			AccountId: account.Id,
			Point:     decimal.NewFromInt(10),
			Fn: func(ctx context.Context, bal *Balance) (err error) {
				return
			},
		})

		if err == nil {
			err = fmt.Errorf("no chared")
			t.Fatal(err)
		}

		var balance = fnPanic.Get(FindBalanceOne(ctx, &FindBalanceOneArgs{
			AccountId: &account.Id,
		}))

		assert.Equal(t, "0", balance.Point.String())
		assert.Equal(t, "0", balance.Currency.String())
	})

	test.Run("use more coupons", func(t *testing.T) {
		var account = fnPanic.Get(CreateAccount(db))

		// expected
		var count = 10
		var currency = 1000
		var price = 10
		var point = 100

		for i := 0; i < count; i++ {
			fnPanic.Get(CreateCoupon(db, &CreateCouponArgs{
				AccountId: account.Id,
				Currency:  decimal.NewFromInt(int64(currency)),
				Price:     decimal.NewFromInt(int64(price)),
				Point:     decimal.NewFromInt(int64(point)),
			}))
		}

		var balance = fnPanic.Get(FindBalanceOne(ctx, &FindBalanceOneArgs{
			AccountId: &account.Id,
		}))

		assert.Equal(t, fmt.Sprintf("%d", count*currency), balance.Currency.String())
		assert.Equal(t, fmt.Sprintf("%d", count*point), balance.Point.String())

		var wg = &sync.WaitGroup{}
		wg.Add(count)
		var totalUsedPoint = 0
		for i := 0; i < count; i++ {
			var usePoint = rand.Int63n(int64(point))
			totalUsedPoint += int(usePoint)

			go func(p decimal.Decimal, w *sync.WaitGroup) {
				_ = CreateRequest(ctx, &CreateRequestArgs{
					AccountId: account.Id,
					Point:     p,
					Fn: func(ctx context.Context, bal *Balance) (err error) {
						return
					},
				})
				w.Done()
			}(decimal.NewFromInt(usePoint), wg)
		}
		wg.Wait()

		balance = fnPanic.Get(FindBalanceOne(ctx, &FindBalanceOneArgs{
			AccountId: &account.Id,
		}))

		assert.Equal(t, fmt.Sprintf("%d", count*currency-price*totalUsedPoint), balance.Currency.String())
		assert.Equal(t, fmt.Sprintf("%d", count*point-totalUsedPoint), balance.Point.String())
	})

	// todo cancel 테스트 하기

}
