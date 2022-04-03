package service

import (
	"context"

	"github.com/jeruktutut2/backend-mongo-user/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	Login(ctx context.Context, username string, password string)
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

func (service *UserServiceImplementation) Login(ctx context.Context, username string, password string) {
	service.UserRepository.GetUserByUsername(service.DB, ctx, username)
}
