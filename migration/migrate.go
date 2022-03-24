package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID:      "2021122816360",
			Migrate: migrationVersion2021122816360,
		},
		{
			ID:      "2022032316040",
			Migrate: migrationVersion2022032316040,
		},
	})

	return m.Migrate()
}
