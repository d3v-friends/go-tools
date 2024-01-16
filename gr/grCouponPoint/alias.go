package grCouponPoint

import (
	"context"
	"fmt"
	"github.com/d3v-friends/go-tools/typ"
	"github.com/d3v-friends/go-tools/wr/wrGorm"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Alias struct {
	Id        typ.UUID  `gorm:"primaryKey;type:char(36)"`
	Kind      string    `gorm:"varchar(50);index:kind_value_1,unique"`
	Value     string    `gorm:"varchar(50);index:kind_value_1,unique"`
	AccountId typ.UUID  `gorm:"index;type:char(36);"`
	CreatedAt time.Time `gorm:"autoUpdateTime"`

	// refer
	Account *Account `gorm:"foreignKey:AccountId;references:Id"`
}

func (x *Alias) Migrate() []wrGorm.Migrate {
	return []wrGorm.Migrate{}
}

/*------------------------------------------------------------------------------------------------*/

type CreateAliasArgs struct {
	AccountId typ.UUID
	Kind      string
	Value     string
}

func CreateAlias(ctx context.Context, i *CreateAliasArgs) (res *Alias, err error) {
	res = &Alias{
		Id:        typ.NewUUID(),
		Kind:      strings.ToLower(i.Kind),
		Value:     strings.ToLower(i.Value),
		AccountId: i.AccountId,
		CreatedAt: time.Now(),
	}

	var rows *gorm.DB
	var db = wrGorm.GetDBP(ctx)
	if rows = db.Create(res); rows.Error != nil {
		err = rows.Error
		return
	}

	return
}

/*------------------------------------------------------------------------------------------------*/

type FindAccountByAliasArgs struct {
	Alias string
}

// FindAccountByAlias 사용가능한 쿠폰 리스트까지 로딩하여 보여준다.
func FindAccountByAlias(ctx context.Context, i *FindAccountByAliasArgs) (res *Account, err error) {
	var db = wrGorm.GetDBP(ctx)
	var rows *gorm.DB
	var alias = new(Alias)
	if rows = db.
		Model(&Alias{}).
		Joins("Account").
		Where("`aliases`.`id` = ?", i.Alias).
		Take(alias); rows.RowsAffected == 0 {
		err = fmt.Errorf("not found alias: alias=%s", i.Alias)
		return
	}

	if alias.Account != nil {
		err = fmt.Errorf("not found account: alias=%s", alias.Id)
		return
	}

	alias.Account.CouponList = make([]*Coupon, 0)
	if rows = db.
		Model(&Coupon{}).
		Joins("Usage").
		Where("`coupons`.`account_id` = ?", alias.AccountId).
		Where("`coupons`.`has_balance` = ?", true).
		Order("`coupons`.`created_at` DESC").
		Find(&alias.Account.CouponList); rows.Error != nil {
		err = rows.Error
		return
	}

	res = alias.Account
	return
}
