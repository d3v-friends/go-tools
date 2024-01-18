package grCouponPoint

import (
	"github.com/d3v-friends/go-tools/typ"
	"github.com/d3v-friends/go-tools/wr/wrGorm"
	"gorm.io/gorm"
	"time"
)

type Account struct {
	Id        typ.UUID  `gorm:"primaryKey;type:char(36)"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// refer
	CouponList []*Coupon `gorm:"foreignKey:AccountId;references:Id"`
	AliasList  []*Alias  `gorm:"foreignKey:AccountId;references:Id"`
}

func (x *Account) Migrate() []wrGorm.Migrate {
	return []wrGorm.Migrate{}
}

/*------------------------------------------------------------------------------------------------*/

func CreateAccount(tx *gorm.DB) (res *Account, err error) {
	var accountId = typ.NewUUID()
	var now = time.Now()
	res = &Account{
		Id:         accountId,
		CreatedAt:  now,
		UpdatedAt:  now,
		CouponList: make([]*Coupon, 0),
		AliasList:  make([]*Alias, 0),
	}

	var rows *gorm.DB
	if rows = tx.Create(res); rows.Error != nil {
		err = rows.Error
		return
	}
	return
}
