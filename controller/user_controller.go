package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/jeruktutut2/backend-mongo-user/model/request"
	"github.com/jeruktutut2/backend-mongo-user/model/response"
	"github.com/jeruktutut2/backend-mongo-user/service"
	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Login(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type UserControllerImplementation struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImplementation{
		UserService: userService,
	}
}

func (controller *UserControllerImplementation) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	userRequestLogin := request.UserRequestLogin{}
	request.ReadFromRequestBody(r, &userRequestLogin)
	username := userRequestLogin.Username
	password := userRequestLogin.Password
	userLoginResponse := controller.UserService.Login(ctx, username, password)
	response.ResponseHttp(w, http.StatusOK, userLoginResponse, nil)
}
