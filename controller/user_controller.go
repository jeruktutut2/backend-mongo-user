package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jeruktutut2/backend-mongo-user/configuration"
	"github.com/jeruktutut2/backend-mongo-user/exception"
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
	ConfigurationWebServer configuration.WebServer
	UserService            service.UserService
}

func NewUserController(configurationWebServer configuration.WebServer, userService service.UserService) UserController {
	return &UserControllerImplementation{
		ConfigurationWebServer: configurationWebServer,
		UserService:            userService,
	}
}

func (controller *UserControllerImplementation) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx, cancel := context.WithTimeout(request.Context(), time.Duration(controller.ConfigurationWebServer.Timeout)*time.Second)
	defer cancel()
	userLoginRequest := web.UserLoginRequest{}
	helper.ReadFromRequestBody(request, &userLoginRequest)
	userLoginResponse, err := controller.UserService.Login(ctx, userLoginRequest)
	if err != nil {
		exception.ErrorHandler(writer, err)
		return
	}
	fmt.Println("userLoginResponse:", userLoginResponse, err)
	response.ResponseHttp(writer, http.StatusOK, userLoginResponse, nil)
}
