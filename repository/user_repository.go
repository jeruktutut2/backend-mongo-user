package repository

import (
	"context"

	"github.com/jeruktutut2/backend-mongo-user/entity"
	"github.com/jeruktutut2/backend-mongo-user/exception"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUserByUsername(DB *mongo.Database, ctx context.Context, username string) (user entity.User)
}

type UserRepositoryImplementation struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImplementation{}
}

func (repository *UserRepositoryImplementation) GetUserByUsername(DB *mongo.Database, ctx context.Context, username string) (user entity.User) {
	result := DB.Collection("user").FindOne(ctx, entity.User{Username: username})
	err := result.Decode(&user)
	exception.PanicIfError(err)
	return user
}
