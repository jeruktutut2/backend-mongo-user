package repository

import (
	"context"

	"github.com/jeruktutut2/backend-mongo-user/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUserByUsername(DB *mongo.Database, ctx context.Context, username string) (user entity.User, err error)
}

type UserRepositoryImplementation struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImplementation{}
}

func (repository *UserRepositoryImplementation) GetUserByUsername(DB *mongo.Database, ctx context.Context, username string) (user entity.User, err error) {
	result := DB.Collection("user").FindOne(ctx, entity.User{Username: username})
	err = result.Decode(&user)
	if err != nil {
		return
	}
	return user, nil
}
