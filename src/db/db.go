package db

import (
	"context"
	"ecommerce-backend/src/configs"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func DBConn(cfg *configs.Env) *mongo.Client {
	var err error
	DB, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.Db.Mongo.Uri))
	if err != nil {
		log.Fatalf("Connect to mongodb fail: %v", err)
	}

	if err = DB.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("Ping to mongodb error: %v", err)
	}

	log.Println("Mongodb has been initialize")
	return DB
}

func GetCollection(cfg *configs.Env, client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database(cfg.Db.Mongo.Database).Collection(collectionName)
}
