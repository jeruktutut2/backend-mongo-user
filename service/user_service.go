package service

import (
	"context"

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
	err = service.Validate.Struct(userLoginRequest)
	err = exception.GetValidationErrors(err, userLoginRequest)
	if err != nil {
		return userLoginResponse, err
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
