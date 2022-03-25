package usecase

import (
	"context"

	"authen-go/app/domain/dto"
	"authen-go/app/infrastructure/casbin"
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

	return casbin.Enforce(authorizeDto.CasbinUser, authorizeDto.RequestURI, authorizeDto.Method)
}
