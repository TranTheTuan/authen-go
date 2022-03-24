package service

import (
	"authen-go/app/domain/model"
	"authen-go/app/domain/repository"
	"authen-go/app/infrastructure/util"
	"errors"
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
	loggedInUser, err := u.userRepo.Login(user.Username)
	if err != nil {
		return nil, err
	}
	if !util.CheckPasswordHash(user.Password, loggedInUser.Password) {
		return nil, errors.New("Login failed")
	}
	return loggedInUser, nil
}

func (u *UserService) IsUsernameExist(username string) (bool, error) {
	return u.userRepo.IsUsernameExist(username)
}

func (u *UserService) Register(user *model.User) (*model.User, error) {
	return u.userRepo.Register(user)
}
