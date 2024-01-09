package fnGorm

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

const ctxDB = "CTX_GORM_DB"

func GetDB(ctx context.Context) (db *gorm.DB, err error) {
	var isOk bool
	if db, isOk = ctx.Value(ctxDB).(*gorm.DB); !isOk {
		err = fmt.Errorf("not found db")
		return
	}
	return
}

func GetDBP(ctx context.Context) (db *gorm.DB) {
	var err error
	if db, err = GetDB(ctx); err != nil {
		panic(err)
	}
	return
}

func SetDB(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, ctxDB, db)
}
