package usecase

import (
	"context"

	"github.com/TranTheTuan/authen-go/app/domain/dto"
	"github.com/TranTheTuan/authen-go/app/infrastructure/casbin"
)

type AuthorUsecaseInterface interface {
	Authorize(ctx context.Context, authorizeDto *dto.AuthorizeDTO) (bool, error)
}

type AuthorUsecase struct{}

func NewAuthorUsecase() *AuthorUsecase {
	return &AuthorUsecase{}
}

func (a *AuthorUsecase) Authorize(ctx context.Context, authorizeDto *dto.AuthorizeDTO) (bool, error) {
	casbin := casbin.GetCasbin()

	isAuthorized, _, err := casbin.EnforceEx("1", authorizeDto.RequestURI, authorizeDto.Method)
	// fmt.Println(casbin.GetPolicy())
	return isAuthorized, err
}
