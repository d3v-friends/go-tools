package grCouponPoint

import (
	"context"
	"fmt"
	"github.com/d3v-friends/go-tools/typ"
	"github.com/d3v-friends/go-tools/wr/wrGorm"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

// Request 포인트 사용
type Request struct {
	Id        typ.UUID        `gorm:"primaryKey;type:char(36)"`
	Status    RequestStatus   `gorm:"index"`
	AccountId typ.UUID        `gorm:"index"`
	Point     decimal.Decimal `gorm:"type:decimal(20,2)"`
	Msg       string          `gorm:"type:text"`
	CreatedAt time.Time       `gorm:"autoCreateTime"`
}

func (x *Request) Migrate() []wrGorm.Migrate {
	return []wrGorm.Migrate{}
}

/*------------------------------------------------------------------------------------------------*/

type CreateRequestArgs struct {
	AccountId typ.UUID
	Point     decimal.Decimal
	Fn        func(ctx context.Context, balance *Balance) error
}

func CreateRequest(ctx context.Context, i *CreateRequestArgs) (resErr error) {
	var request = &Request{
		Id:        typ.NewUUID(),
		Status:    RequestStatusRequested,
		AccountId: i.AccountId,
		Point:     i.Point,
		CreatedAt: time.Now(),
	}

	var db = wrGorm.GetDBP(ctx)
	var rows *gorm.DB
	if rows = db.Create(request); rows.Error != nil {
		resErr = rows.Error
		return
	}

	defer func() {
		var updates = map[string]any{
			"status": RequestStatusSucceed,
		}

		if resErr != nil {
			updates["status"] = RequestStatusFail
			updates["msg"] = resErr.Error()
		}

		db.
			Model(&Request{}).
			Where("`requests`.`id` = ?", request.Id).
			Updates(updates)
	}()

	resErr = db.Transaction(func(tx *gorm.DB) (txErr error) {
		var couponList = make([]*Coupon, 0)
		var rows *gorm.DB

		if rows = tx.
			Model(&Coupon{}).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("`coupons`.`account_id` = ?", i.AccountId).
			Where("`coupons`.`has_balance` = ?", true).
			Order("`coupons`.`created_at` ASC").
			Joins("Usage").
			Find(&couponList); rows.Error != nil {
			txErr = rows.Error
			return
		}

		var totalPoint = decimal.Zero
		for _, coupon := range couponList {
			totalPoint = totalPoint.Add(coupon.Usage.Point)
		}

		if totalPoint.LessThan(i.Point) {
			txErr = fmt.Errorf("shoartage point")
			return
		}

		var now = time.Now()
		var leftPoint = i.Point.Copy()

		for _, coupon := range couponList {
			var nextUsage *Usage
			if coupon.Usage.Point.GreaterThanOrEqual(leftPoint) {
				nextUsage = &Usage{
					Id:           typ.NewUUID(),
					CouponId:     coupon.Id,
					AccountId:    coupon.AccountId,
					RequestId:    &request.Id,
					Point:        coupon.Usage.Point.Sub(leftPoint),
					ChangedPoint: leftPoint.Neg(),
					PrevPoint:    coupon.Usage.Point,
					CreatedAt:    now,
				}
				leftPoint = decimal.Zero
			} else {
				nextUsage = &Usage{
					Id:           typ.NewUUID(),
					CouponId:     coupon.Id,
					AccountId:    coupon.AccountId,
					RequestId:    &request.Id,
					Point:        decimal.Zero,
					ChangedPoint: coupon.Usage.Point.Neg(),
					PrevPoint:    coupon.Usage.Point,
					CreatedAt:    now,
				}
				leftPoint = leftPoint.Sub(coupon.Usage.Point)
			}

			if rows = tx.Model(&Usage{}).Create(nextUsage); rows.Error != nil {
				txErr = rows.Error
				return
			}

			var updates = map[string]any{
				"has_balance": !nextUsage.Point.Equal(decimal.Zero),
				"usage_id":    nextUsage.Id,
				"updated_at":  now,
			}

			if rows = tx.
				Model(&Coupon{}).
				Where("`coupons`.`id` = ?", coupon.Id).
				Updates(updates); rows.RowsAffected == 0 {
				txErr = fmt.Errorf("not found coupon: coupon=%s", coupon.Id)
				return
			}

			if leftPoint.Equal(decimal.Zero) {
				break
			}
		}

		var balance *Balance
		if balance, txErr = findBalanceOne(tx, &FindBalanceOneArgs{
			AccountId: &i.AccountId,
		}); txErr != nil {
			return
		}

		return i.Fn(ctx, balance)
	})

	return
}
