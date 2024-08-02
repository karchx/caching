package data

import (
	"context"
	"time"

	"github.com/karchx/api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db          = "cacheService"
	collection  = "employee"
	MongoClient *mongo.Client
)

func InitializeMongoClient() error {
	var err error
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	MongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(config.MongoDBURI))
	if err != nil {
		return err
	}
	return nil
}

func getCollection() *mongo.Collection {
	return MongoClient.Database(db).Collection(collection)
}
