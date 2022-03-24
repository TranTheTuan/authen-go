package repository

import (
	"authen-go/app/domain/model"
)

type UserRepositoryInterface interface {
	Login(username string) (*model.User, error)
	IsUsernameExist(username string) (bool, error)
	Register(user *model.User) (*model.User, error)
}
