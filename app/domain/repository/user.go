package repository

import (
	"github.com/TranTheTuan/authen-go/app/domain/model"
)

type UserRepositoryInterface interface {
	FindUserByUsername(username string) (*model.User, error)
	// IsUsernameExist(username string) (bool, error)
	CreateUser(user *model.User) (*model.User, error)
}
