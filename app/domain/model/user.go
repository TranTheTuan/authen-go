package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `json:"username"`
	Password   string `json:"password"`
	CasbinUser string `json:"casbin_user"`
}

func (User) TableName() string {
	return "users"
}
