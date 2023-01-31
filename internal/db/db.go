package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

const (
	URI = "mongodb://api:123456@mongo:27017/api"
	Database = "api"
) 

type Collection string

const (
	ProductsCollection Collection = "products"
)

var clientInstance *mongo.Client 

var mongoOnce sync.Once

var clientInstanceError error

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(URI)

		client, err := mongo.Connect(context.TODO(), clientOptions)
        if err!= nil {
            panic(err)
        }
        clientInstance = client

		clientInstanceError = err 
	})

	return clientInstance, clientInstanceError
}