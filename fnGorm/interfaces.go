package fnGorm

import "gorm.io/gorm"

type (
	MigrateModel interface {
		Migrate() []FnMigrate
	}
	FnMigrate func(tx *gorm.DB) (err error)
)
