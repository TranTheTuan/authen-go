package service

import (
	"errors"

	"github.com/jinzhu/gorm"

	"authen-go/app/domain/model"
	"authen-go/app/domain/repository"
	"authen-go/app/infrastructure/casbin"
	"authen-go/app/infrastructure/util"
)

type UserServiceInterface interface {
	Login(user *model.User) (*model.User, error)
	IsUsernameExist(username string) (bool, error)
	Register(user *model.User) (*model.User, error)
}

type UserService struct {
	userRepo repository.UserRepositoryInterface
}

func NewUserService(userRepo repository.UserRepositoryInterface) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) Login(user *model.User) (*model.User, error) {
	loggedInUser, err := u.userRepo.FindUserByUsername(user.Username)
	if err != nil {
		return nil, err
	}
	if !util.CheckPasswordHash(user.Password, loggedInUser.Password) {
		return nil, errors.New("Login failed")
	}
	return loggedInUser, nil
}

func (u *UserService) IsUsernameExist(username string) (bool, error) {
	_, err := u.userRepo.FindUserByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (u *UserService) Register(user *model.User) (*model.User, error) {
	user, err := u.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	_, err = casbin.AddRole(user.CasbinUser, "MEMBER_PERSONA_USER")
	if err != nil {
		return nil, err
	}
	return user, nil
}
