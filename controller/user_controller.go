package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/jeruktutut2/backend-mongo-user/helper"
	"github.com/jeruktutut2/backend-mongo-user/model/response"
	"github.com/jeruktutut2/backend-mongo-user/model/web"
	"github.com/jeruktutut2/backend-mongo-user/service"
	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type UserControllerImplementation struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImplementation{
		UserService: userService,
	}
}

func (controller *UserControllerImplementation) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx, cancel := context.WithTimeout(request.Context(), 10*time.Second)
	defer cancel()
	userLoginRequest := web.UserLoginRequest{}
	helper.ReadFromRequestBody(request, &userLoginRequest)
	username := userLoginRequest.Username
	password := userLoginRequest.Password
	userLoginResponse := controller.UserService.Login(ctx, username, password)
	response.ResponseHttp(writer, http.StatusOK, userLoginResponse, nil)
}
