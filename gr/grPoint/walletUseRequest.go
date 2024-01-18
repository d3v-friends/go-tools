package grPoint

import (
	"github.com/d3v-friends/go-tools/typ"
	"github.com/d3v-friends/go-tools/wr/wrGorm"
	"github.com/shopspring/decimal"
	"time"
)

type WalletUseRequest struct {
	Id        typ.UUID        `gorm:"primaryKey;type:char(36)"`
	UsedPoint decimal.Decimal `gorm:"type:decimal(20,2)"`
	Msg       string          `gorm:"type:text"`
	CreatedAt time.Time       `gorm:"index"`

	// ref
	WalletUseReceipt []*WalletUseReceipt `gorm:"foreignKey:WalletUseRequestId;references:Id"`
}

func (x *WalletUseRequest) Migrate() []wrGorm.Migrate {
	return make([]wrGorm.Migrate, 0)
}

type WalletUseReceipt struct {
	Id                 typ.UUID `gorm:"primaryKey"`
	WalletUseRequestId typ.UUID `gorm:"index"`
	WalletBalanceId    typ.UUID `gorm:"index"`

	// ref
	WalletUseRequest *WalletUseRequest `gorm:"foreignKey:WalletUseRequestId;references:Id"`
	WalletBalance    *WalletBalance    `gorm:"foreignKey:WalletBalanceId;references:Id"`
}

func (x *WalletUseReceipt) Migrate() []wrGorm.Migrate {
	return make([]wrGorm.Migrate, 0)
}
