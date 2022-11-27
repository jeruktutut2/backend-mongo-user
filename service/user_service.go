package service

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/jeruktutut2/backend-mongo-user/exception"
	"github.com/jeruktutut2/backend-mongo-user/model/response"
	"github.com/jeruktutut2/backend-mongo-user/model/web"
	"github.com/jeruktutut2/backend-mongo-user/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(ctx context.Context, userLoginRequest web.UserLoginRequest) (userLoginResponse response.UserLoginResponse, err error)
}

type UserServiceImplementation struct {
	DB             *mongo.Database
	Validate       *validator.Validate
	UserRepository repository.UserRepository
}

func NewUserService(DB *mongo.Database, validate *validator.Validate, userRepository repository.UserRepository) UserService {
	return &UserServiceImplementation{
		DB:             DB,
		Validate:       validate,
		UserRepository: userRepository,
	}
}

func (service *UserServiceImplementation) Login(ctx context.Context, userLoginRequest web.UserLoginRequest) (userLoginResponse response.UserLoginResponse, err error) {
	return userLoginResponse, exception.NewBadRequestError("bad request")
	err = service.Validate.Struct(userLoginRequest)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("fieldError.ActualTag():", fieldError.ActualTag())
			fmt.Println("fieldError.Error():", fieldError.Error())
			fmt.Println("fieldError.Field():", fieldError.Field())
			fmt.Println("fieldError.Kind():", fieldError.Kind())
			fmt.Println("fieldError.Namespace():", fieldError.Namespace())
			fmt.Println("fieldError.Param():", fieldError.Param())
			fmt.Println("fieldError.StructField():", fieldError.StructField())
			fmt.Println("fieldError.StructNamespace():", fieldError.StructNamespace())
			fmt.Println("fieldError.Tag():", fieldError.Tag())
			fmt.Println("fieldError.Type():", fieldError.Type())
			fmt.Println("fieldError.Value():", fieldError.Value())
			fmt.Println("fieldError.Kind().String():", fieldError.Kind().String())
			fmt.Println("fieldError.Type().Align():", fieldError.Type().Align())
		}
	}
	user, err := service.UserRepository.GetUserByUsername(service.DB, ctx, userLoginRequest.Username)
	if err != nil {
		return userLoginResponse, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLoginRequest.Password))
	if err != nil {
		return userLoginResponse, err
	}
	userLoginResponse = response.ToUserLoginResponse(user.Id, user.Username)
	return userLoginResponse, nil
}
