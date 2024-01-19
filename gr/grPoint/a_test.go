package grPoint

import (
	"context"
	"fmt"
	"github.com/d3v-friends/go-tools/fn/fnEnv"
	"github.com/d3v-friends/go-tools/fn/fnPanic"
	"github.com/d3v-friends/go-tools/fn/fnParams"
	"github.com/d3v-friends/go-tools/wr/wrGorm"
	"gorm.io/gorm"
	"sync"
	"testing"
)

func TestAll(test *testing.T) {
	NewTestTool(true)
	TestAccount(test)
	TestCoupon(test)
	TestCouponUseRequest(test)
	TestWallet(test)
	TestWalletUseRequest(test)

}

type TestTool struct {
	DB *gorm.DB
}

func NewTestTool(truncate ...bool) (res *TestTool) {
	fnPanic.On(fnEnv.ReadFromFile("../.env"))
	var c = fnPanic.Get(wrGorm.NewConnect(&wrGorm.ConnectArgs{
		Host:     fnEnv.Read("DB_HOST"),
		Username: fnEnv.Read("DB_USERNAME"),
		Password: fnEnv.Read("DB_PASSWORD"),
		Schema:   fnEnv.Read("DB_SCHEMA"),
	}))

	res = &TestTool{
		DB: c.DB,
	}

	var models = []wrGorm.MigrateModel{
		&Account{},
		&Coupon{},
		&CouponBalance{},
		&CouponUseRequest{},
		&CouponUseReceipt{},
		&Wallet{},
		&WalletBalance{},
		&WalletUseRequest{},
		&WalletUseReceipt{},
	}

	if fnParams.Get(truncate) {
		res.TruncateAll(models)
	}

	fnPanic.On(c.Migrate(
		models...,
	))

	return
}

func (x *TestTool) TruncateAll(models []wrGorm.MigrateModel) {
	x.DB.Exec("set FOREIGN_KEY_CHECKS = 0")
	for _, model := range models {
		x.DB.Exec(fmt.Sprintf("TRUNCATE TABLE %s", wrGorm.GetTableNm(model)))
	}
	x.DB.Exec("set FOREIGN_KEY_CHECKS = 1")
}

func (x *TestTool) Context() context.Context {
	var ctx = context.TODO()
	return wrGorm.SetDB(ctx, x.DB.WithContext(ctx))
}

func (x *TestTool) NewAccount() (res *Account) {
	var wg = &sync.WaitGroup{}
	wg.Add(1)
	fnPanic.On(CreateAccount(x.Context(), &CreateAccountArgs{
		Fn: func(ctx context.Context, i *Account) (err error) {
			res = i
			wg.Done()
			return
		},
	}))
	wg.Wait()
	return
}
