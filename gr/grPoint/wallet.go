package grPoint

import (
	"context"
	"errors"
	"github.com/d3v-friends/go-tools/typ"
	"github.com/d3v-friends/go-tools/wr/wrGorm"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

type Wallets []*Wallet

/*------------------------------------------------------------------------------------------------*/

type FindWalletArgs struct {
	Id        []typ.UUID
	AccountId []typ.UUID
	Lock      *bool
}

func (x FindWalletArgs) Query(db *gorm.DB) *gorm.DB {
	var query = db.Model(&Wallet{})

	if len(x.Id) != 0 {
		query = query.Where("`wallets`.`id` IN (?)", x.Id)
	}

	if len(x.AccountId) != 0 {
		query = query.Where("`wallets`.`account_id` IN (?)", x.AccountId)
	}

	if x.Lock != nil && *x.Lock {
		query = query.Clauses(clause.Locking{
			Strength: "UPDATE",
		})
	}

	query = query.
		Joins("WalletBalance").
		Order("`wallets`.`created_at` DESC")

	return query
}

func FindOneWalletCtx(ctx context.Context, i *FindWalletArgs) (*Wallet, error) {
	return FindOneWallet(wrGorm.GetDBP(ctx), i)
}

func FindOneWallet(tx *gorm.DB, i *FindWalletArgs) (wallet *Wallet, err error) {
	wallet = new(Wallet)

	var rows *gorm.DB
	if rows = i.Query(tx).Take(wallet); rows.Error != nil {
		err = rows.Error
		return
	}

	return
}

func FindAllWallets(tx *gorm.DB, i *FindWalletArgs) (ls Wallets, err error) {
	var rows *gorm.DB
	ls = make(Wallets, 0)

	if rows = i.Query(tx).Find(&ls); rows.Error != nil {
		if errors.Is(rows.Error, gorm.ErrEmptySlice) {
			return
		}
		err = rows.Error
		return
	}

	return
}
