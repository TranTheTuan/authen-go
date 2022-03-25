package migration

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `json:"username"`
	Password   string `json:"password"`
	CasbinUser string `json:"casbin_user"`
}

func migrationVersion20220325100400(tx *gorm.DB) error {
	if !tx.Migrator().HasTable(&User{}) {
		if err := tx.AutoMigrate(&User{}); err != nil {
			return err
		}
	}
	return nil
}
