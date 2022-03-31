package controller

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Login(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type UserControllerImplementation struct {
}

func NewUserController() UserController {
	return &UserControllerImplementation{}
}

func (controller *UserControllerImplementation) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("login")
}
