package migration

import (
	"github.com/TranTheTuan/authen-go/app/domain/common"

	"gorm.io/gorm"
)

type IndexStore struct {
	gorm.Model
	DataJSON common.JSON            `gorm:"column:data_json;type:json"`
	Data     map[string]interface{} `gorm:"-"`
}

func migrationVersion2022032316040(tx *gorm.DB) error {
	if !tx.Migrator().HasTable(&IndexStore{}) {
		if err := tx.AutoMigrate(&IndexStore{}); err != nil {
			return err
		}
	}
	return nil
}
