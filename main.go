package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jeruktutut2/backend-mongo-user/configuration"
	"github.com/jeruktutut2/backend-mongo-user/controller"
	"github.com/jeruktutut2/backend-mongo-user/repository"
	"github.com/jeruktutut2/backend-mongo-user/route"
	"github.com/jeruktutut2/backend-mongo-user/service"
	"github.com/jeruktutut2/backend-mongo-user/util"
	"github.com/julienschmidt/httprouter"
)

func main() {
	config := configuration.NewConfiguration()
	mongoDatabase, mongoDbClient := util.NewMongoConnection(config.Mongo)
	validate := validator.New()
	// fmt.Println("mongoDatabase:", mongoDatabase)

	router := httprouter.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(mongoDatabase, validate, userRepository)
	userController := controller.NewUserController(config.WebServer, userService)
	route.UserRoute(router, userController)

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(config.WebServer.Port),
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		fmt.Println("Application ready")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	util.CloseMongoDbConnection(mongoDbClient)

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %+v", err)
	}
	log.Println("Server exited properly")
}
