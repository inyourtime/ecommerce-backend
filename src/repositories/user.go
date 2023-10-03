package repositories

import (
	"context"
	"ecommerce-backend/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRopository interface {
	Create(user models.User) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
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

func (r userRopository) FindByEmail(email string) (*models.User, error) {
	user := models.User{}
	filter := bson.D{{Key: "email", Value: email}}
	project := options.FindOne().SetProjection(bson.M{"password": 0})
	err := r.col.FindOne(context.TODO(), filter, project).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
