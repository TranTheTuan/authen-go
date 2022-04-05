package grpc

import (
	"context"

	"github.com/TranTheTuan/authen-go/app/domain/dto"
	"github.com/TranTheTuan/authen-go/app/domain/usecase"
	"github.com/TranTheTuan/authen-go/app/infrastructure/util"
	pbAuth "github.com/TranTheTuan/pbtypes/build/go/auth"
)

type AuthorizeServiceServer struct {
	authorUsecase usecase.AuthorUsecaseInterface

	pbAuth.UnimplementedAuthAuthorizeServiceServer
	pbAuth.UnimplementedAuthVerifyServiceServer
}

func NewAuthorizeServiceServer(authorUsecase usecase.AuthorUsecaseInterface) *AuthorizeServiceServer {
	return &AuthorizeServiceServer{
		authorUsecase: authorUsecase,
	}
}

func (a *AuthorizeServiceServer) Authorize(ctx context.Context, in *pbAuth.AuthorizeRequest) (*pbAuth.AuthorizeResponse, error) {
	isAuthorized, err := a.authorUsecase.Authorize(ctx, &dto.AuthorizeDTO{
		CasbinUser: in.CasbinUser,
		RequestURI: in.RequestUri,
		Method:     in.Method,
	})
	if err != nil {
		return nil, err
	}
	return &pbAuth.AuthorizeResponse{
		Pass: isAuthorized,
	}, nil
}

func (a *AuthorizeServiceServer) VerifyToken(ctx context.Context, in *pbAuth.VerifyTokenRequest) (*pbAuth.VerifyTokenResponse, error) {
	j := util.NewJWT()
	claim, err := j.ParseToken(in.Token)
	if err != nil {
		return nil, err
	}

	return &pbAuth.VerifyTokenResponse{
		Id: uint32(claim.ID),
	}, nil
}
