package migration

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func migrationVersion2021122816360(tx *gorm.DB) error {
	if !tx.Migrator().HasTable(&User{}) {
		if err := tx.AutoMigrate(&User{}); err != nil {
			return err
		}
	}
	return nil
}
