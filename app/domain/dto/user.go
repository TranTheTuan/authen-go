package dto

import "github.com/TranTheTuan/authen-go/app/domain/model"

type CreateUserDto struct {
	Username string `json:"user_name`
	Password string `json:"password`
}

func (CreateUserDto) TableName() string {
	return model.User{}.TableName()
}
