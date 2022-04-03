package repository

import (
	"context"
	"fmt"

	"github.com/jeruktutut2/backend-mongo-user/entity"
	"github.com/jeruktutut2/backend-mongo-user/exception"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUserByUsername(DB *mongo.Database, ctx context.Context, username string)
}

type UserRepositoryImplementation struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImplementation{}
}

func (repository *UserRepositoryImplementation) GetUserByUsername(DB *mongo.Database, ctx context.Context, username string) {
	cursor, err := DB.Collection("user").Find(ctx, entity.User{})
	defer cursor.Close(ctx)
	exception.PanicIfError(err)
	for cursor.Next(ctx) {
		user := entity.User{}
		err := cursor.Decode(&user)
		exception.PanicIfError(err)
		fmt.Println("user:", user)
	}
}
