package grCouponPoint

import (
	"github.com/d3v-friends/go-tools/typ"
	"github.com/d3v-friends/go-tools/wr/wrGorm"
	"github.com/shopspring/decimal"
	"time"
)

// Usage 쿠폰 사용 히스토리
type Usage struct {
	Id           typ.UUID        `gorm:"primaryKey;type:char(36)"`
	CouponId     typ.UUID        `gorm:"index;type:char(36)"`
	AccountId    typ.UUID        `gorm:"index;type:char(36)"`
	RequestId    *typ.UUID       `gorm:"index;type:char(36)"` // 포인트 사용 아이디
	Point        decimal.Decimal `gorm:"type:decimal(20,2)"`  // 잔고
	ChangedPoint decimal.Decimal `gorm:"type:decimal(20,2)"`  // 변경된 값
	PrevPoint    decimal.Decimal `gorm:"type:decimal(20,2)"`  // 이전 잔고
	CreatedAt    time.Time       `gorm:"autoCreateTime"`
}

func (x *Usage) Migrate() []wrGorm.Migrate {
	return []wrGorm.Migrate{}
}
