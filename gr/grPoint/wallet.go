package grPoint

import (
	"github.com/d3v-friends/go-tools/typ"
	"github.com/d3v-friends/go-tools/wr/wrGorm"
	"github.com/shopspring/decimal"
	"time"
)

type Wallet struct {
	Id              typ.UUID  `gorm:"primaryKey;type:char(36)"`
	AccountId       typ.UUID  `gorm:"index;type:char(36)"`
	WalletBalanceId typ.UUID  `gorm:"uniqueIndex;type:char(36)"`
	CreatedAt       time.Time `bson:"index"`
	UpdatedAt       time.Time `bson:"index"`

	// ref
	WalletBalance *WalletBalance `gorm:"foreignKey:Id;references:WalletId"`
}

func (x *Wallet) Migrate() []wrGorm.Migrate {
	return make([]wrGorm.Migrate, 0)
}

type WalletBalance struct {
	Id           typ.UUID        `gorm:"primaryKey;type:char(36)"`
	WalletId     typ.UUID        `gorm:"index;type:char(36)"`
	Point        decimal.Decimal `gorm:"type:decimal(20,2)"`
	PrevPoint    decimal.Decimal `gorm:"type:decimal(20,2)"`
	ChangedPoint decimal.Decimal `gorm:"type:decimal(20,2)"`
	Memo         string          `gorm:"type:text"`
	CreatedAt    time.Time       `gorm:"index"`

	// ref
	Wallet *Wallet `gorm:"foreignKey:Id;references:WalletId;"`
}

func (x *WalletBalance) Migrate() []wrGorm.Migrate {
	return make([]wrGorm.Migrate, 0)
}
