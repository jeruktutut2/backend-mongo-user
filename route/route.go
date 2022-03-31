package route

import (
	"github.com/jeruktutut2/backend-mongo-user/controller"
	"github.com/julienschmidt/httprouter"
)

func UserRoute(router *httprouter.Router, userController controller.UserController) {
	router.POST("/api/v1/login", userController.Login)
}
