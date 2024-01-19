package grPoint

import "github.com/d3v-friends/go-tools/wr/wrGorm"

var MigrateModels = []wrGorm.MigrateModel{
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
