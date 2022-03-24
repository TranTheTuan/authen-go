package usecase

import (
	"authen-go/app/domain/model"
	"authen-go/app/domain/service"
	"authen-go/app/infrastructure/util"
	"errors"
)

type UserUsecaseInterface interface {
	Login(user *model.User) (*util.TokenInfo, error)
	Register(user *model.User) (*model.User, error)
}

type UserUsecase struct {
	userService service.UserServiceInterface
}

func NewUserUsecase(userService service.UserServiceInterface) *UserUsecase {
	return &UserUsecase{
		userService: userService,
	}
}

func (u *UserUsecase) Login(user *model.User) (*util.TokenInfo, error) {
	loggedInUser, err := u.userService.Login(user)
	if err != nil {
		return nil, err
	}

	j := util.NewJWT()
	tokenInfo, err := j.GenerateToken(loggedInUser.ID, loggedInUser.Username)
	if err != nil {
		return nil, err
	}

	return &tokenInfo, nil
}

func (u *UserUsecase) Register(user *model.User) (*model.User, error) {
	hashedPw, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPw
	isExist, err := u.userService.IsUsernameExist(user.Username)
	if err != nil {
		return nil, err
	}
	if !isExist {
		return u.userService.Register(user)
	}
	err = errors.New("username is not available")
	return nil, err
}
