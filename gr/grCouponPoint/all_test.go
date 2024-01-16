package grCouponPoint

import (
	"context"
	"fmt"
	"github.com/d3v-friends/go-tools/fn/fnEnv"
	"github.com/d3v-friends/go-tools/fn/fnPanic"
	"github.com/d3v-friends/go-tools/fn/fnParams"
	"github.com/d3v-friends/go-tools/wr/wrGorm"
	"gorm.io/gorm"
	"testing"
)

type TestTool struct {
	DB *gorm.DB
}

func TestAll(test *testing.T) {
	TestAccount(test)
	TestCoupon(test)
	TestRequest(test)
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
		&Alias{},
		&Coupon{},
		&Request{},
		&Usage{},
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
	return wrGorm.SetDB(context.TODO(), x.DB)
}
