package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID:      "20220325100400",
			Migrate: migrationVersion20220325100400,
		},
		{
			ID:      "2022032316040",
			Migrate: migrationVersion2022032316040,
		},
	})

	return m.Migrate()
}
