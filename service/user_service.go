package service

import (
	"context"

	"github.com/jeruktutut2/backend-mongo-user/exception"
	"github.com/jeruktutut2/backend-mongo-user/model/response"
	"github.com/jeruktutut2/backend-mongo-user/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(ctx context.Context, username string, password string) (userLoginResponse response.UserLoginResponse)
}

type UserServiceImplementation struct {
	DB             *mongo.Database
	UserRepository repository.UserRepository
}

func NewUserService(DB *mongo.Database, userRepository repository.UserRepository) UserService {
	return &UserServiceImplementation{
		DB:             DB,
		UserRepository: userRepository,
	}
}

func (service *UserServiceImplementation) Login(ctx context.Context, username string, password string) (userLoginResponse response.UserLoginResponse) {
	user := service.UserRepository.GetUserByUsername(service.DB, ctx, username)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	exception.PanicIfError(err)
	userLoginResponse = response.ToUserLoginResponse(user.Id, user.Username)
	return userLoginResponse
}
