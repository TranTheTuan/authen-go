package http

import (
	"encoding/json"
	"net/http"

	"github.com/TranTheTuan/authen-go/app/domain/dto"
	"github.com/TranTheTuan/authen-go/app/domain/model"
	"github.com/TranTheTuan/authen-go/app/domain/usecase"

	"github.com/gorilla/mux"
)

type AuthHandler struct {
	userUsecase   usecase.UserUsecaseInterface
	authorUsecase usecase.AuthorUsecaseInterface
}

func NewAuthHandler(userUsecase usecase.UserUsecaseInterface, authorUsecase usecase.AuthorUsecaseInterface) *AuthHandler {
	return &AuthHandler{
		userUsecase:   userUsecase,
		authorUsecase: authorUsecase,
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
}

func (a *AuthHandler) TestAuthorize(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	requestUri := r.RequestURI
	casbinUser := mux.Vars(r)["id"]

	isAuthorized, err := a.authorUsecase.Authorize(r.Context(), &dto.AuthorizeDTO{
		CasbinUser: casbinUser,
		RequestURI: requestUri,
		Method:     method,
	})
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, false, "unauthorized", nil, err)
		return
	}

	if !isAuthorized {
		JSONResponse(w, http.StatusOK, true, "unauthorized", isAuthorized, nil)
		return
	}

	JSONResponse(w, http.StatusOK, true, "authorized", isAuthorized, nil)
}
