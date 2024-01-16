package grCouponPoint

import (
	"github.com/d3v-friends/go-tools/wr/wrGorm"
	"gorm.io/gorm"
)

func ConnectAndMigrate(i *wrGorm.ConnectArgs) (db *gorm.DB, err error) {
	var c *wrGorm.FnGorm
	if c, err = wrGorm.NewConnect(i); err != nil {
		return
	}

	if err = c.Migrate(
		&Account{},
		&Alias{},
		&Coupon{},
		&Request{},
		&Usage{},
	); err != nil {
		return
	}

	db = c.DB
	return
}
