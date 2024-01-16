package grCouponPoint

import (
	"context"
	"fmt"
	"github.com/d3v-friends/go-tools/typ"
	"github.com/d3v-friends/go-tools/wr/wrGorm"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type Coupon struct {
	Id         typ.UUID        `gorm:"primaryKey;type:char(36)"`
	HasBalance bool            `gorm:"index"`                     // 잔고여부
	AccountId  typ.UUID        `gorm:"index;type:char(36)"`       // 소유권
	Currency   decimal.Decimal `gorm:"type:decimal(20,2)"`        // 구매한 쿠폰 가격
	Price      decimal.Decimal `gorm:"type:decimal(20,2)"`        // 1건당 가격 예 15원
	TotalPoint decimal.Decimal `gorm:"type:decimal(20,2)"`        // ( currency / price ) 갯수, 잔여 남으면 + 1건 해준다
	UsageId    typ.UUID        `gorm:"uniqueIndex;type:char(36)"` // 잔고 정보
	CreatedAt  time.Time       `gorm:"autoCreateTime"`
	UpdatedAt  time.Time       `gorm:"autoUpdateTime"`

	// refer
	Usage *Usage `gorm:"foreignKey:UsageId;references:Id"`
}

func (x *Coupon) Migrate() []wrGorm.Migrate {
	return []wrGorm.Migrate{}
}

/*------------------------------------------------------------------------------------------------*/

type CreateCouponArgs struct {
	AccountId typ.UUID        // accountID
	Currency  decimal.Decimal // 원화
	Price     decimal.Decimal // 1개당 가격
	Point     decimal.Decimal // 전체 전송 가능 갯수
}

func CreateCoupon(ctx context.Context, i *CreateCouponArgs) (res *Coupon, err error) {
	var usageId = typ.NewUUID()
	var couponId = typ.NewUUID()
	var now = time.Now()
	res = &Coupon{
		Id:         couponId,
		HasBalance: true,
		AccountId:  i.AccountId,
		Currency:   i.Currency,
		Price:      i.Price,
		TotalPoint: i.Point,
		UsageId:    usageId,
		CreatedAt:  now,
		UpdatedAt:  now,
		Usage: &Usage{
			Id:           usageId,
			CouponId:     couponId,
			AccountId:    i.AccountId,
			Point:        i.Point,
			ChangedPoint: i.Point,
			PrevPoint:    decimal.Zero,
			CreatedAt:    now,
		},
	}

	var db = wrGorm.GetDBP(ctx)
	var rows *gorm.DB
	if rows = db.Model(&Coupon{}).Create(res); rows.Error != nil {
		err = rows.Error
		return
	}

	return
}

/*------------------------------------------------------------------------------------------------*/

type FindBalanceOneArgs struct {
	AccountId  *typ.UUID
	Alias      *FindAliasArgs
	HasBalance *bool
}

func (x *FindBalanceOneArgs) GetAccountId(db *gorm.DB) (res typ.UUID, err error) {
	if x.AccountId != nil {
		res = *x.AccountId
		return
	}

	var alias = new(Alias)
	var rows = db.Model(&Alias{}).
		Select("`aliases`.`account_id`").
		Where("`aliases`.`kind` = ?", x.Alias.Kind).
		Where("`aliases`.`value` = ?", x.Alias.Value).
		Take(alias)

	if rows.RowsAffected == 0 {
		err = fmt.Errorf("not found account_id")
		return
	}

	res = alias.AccountId
	return
}

type FindAliasArgs struct {
	Kind  string
	Value string
}

type Balance struct {
	AccountId  typ.UUID
	Point      decimal.Decimal // 문자 보낼수 있는 갯수
	Currency   decimal.Decimal // 잔여 금액
	CouponList []*Coupon       // 잔여 금액이 있는 쿠폰 리스트
	QueryAt    time.Time       // 조회시각
}

// FindBalanceOne 잔고 조회
func FindBalanceOne(ctx context.Context, i *FindBalanceOneArgs) (res *Balance, err error) {
	return findBalanceOne(wrGorm.GetDBP(ctx), i)
}

func findBalanceOne(db *gorm.DB, i *FindBalanceOneArgs) (res *Balance, err error) {
	var accountId typ.UUID
	if accountId, err = i.GetAccountId(db); err != nil {
		return
	}

	var query = db.
		Model(&Coupon{}).
		Where("`coupons`.`account_id` = ?", accountId)

	if i.HasBalance != nil {
		query = query.Where("`coupons`.`has_balance` = ?", *i.HasBalance)
	}

	var rows *gorm.DB
	var couponList = make([]*Coupon, 0)
	if rows = query.
		Order("`coupons`.`created_at` DESC").
		Joins("Usage").
		Find(&couponList); rows.Error != nil {
		err = rows.Error
		return
	}

	res = &Balance{
		AccountId:  accountId,
		Point:      decimal.Zero,
		Currency:   decimal.Zero,
		CouponList: couponList,
		QueryAt:    time.Now(),
	}

	// 쿠폰 잔고 더하기
	for _, coupon := range couponList {
		var couponCurrency = coupon.Usage.Point.Mul(coupon.Price)
		res.Currency = res.Currency.Add(couponCurrency)
		res.Point = res.Point.Add(coupon.Usage.Point)
	}

	return
}

func FindBalanceAll(ctx context.Context, i *FindBalanceOneArgs) (res []*Balance, err error) {
	panic("not implement")
}
