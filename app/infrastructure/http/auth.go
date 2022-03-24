package http

import (
	"encoding/json"
	"net/http"

	"authen-go/app/domain/model"
	"authen-go/app/domain/usecase"
)

type AuthHandler struct {
	userUsecase usecase.UserUsecaseInterface
}

func NewAuthHandler(userUsecase usecase.UserUsecaseInterface) *AuthHandler {
	return &AuthHandler{
		userUsecase: userUsecase,
	}
}

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user *model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, false, "body request is invalid", nil, err)
		return
	}

	tokenInfo, err := a.userUsecase.Login(user)
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, false, "login failed", nil, err)
		return
	}

	JSONResponse(w, http.StatusOK, true, "logged in successfully", tokenInfo, nil)
	return
}

func (a *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user *model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, false, "body request is invalid", nil, err)
		return
	}

	user, err = a.userUsecase.Register(user)
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, false, "register failed", nil, err)
		return
	}

	JSONResponse(w, http.StatusOK, true, "registered successfully", user, nil)
	return
}
