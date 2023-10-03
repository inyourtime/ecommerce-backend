package repositories

import (
	"context"
	"ecommerce-backend/src/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRopository interface {
	Create(user models.User) (*models.User, error)
}

type userRopository struct {
	col *mongo.Collection
}

func NewUserRepository(col *mongo.Collection) UserRopository {
	return userRopository{col: col}
}

func (r userRopository) Create(user models.User) (*models.User, error) {
	_, err := r.col.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
