package controller

import (
	"fmt"
	"net/http"

	"github.com/jeruktutut2/backend-mongo-user/model/request"
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
	fmt.Println("login")
	userRequestLogin := request.UserRequestLogin{}
	request.ReadFromRequestBody(r, &userRequestLogin)
	fmt.Println("request.UserRequestLogin:", userRequestLogin)
	// username := request.UserRequestLogin.Username
	// password := request.User.Password
	// controller.UserService(r.Context(), username, password)

}
