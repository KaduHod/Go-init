package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	URI = "mongodb://@mongo:27017/"
	Database = "products"
) 

type Collection string

const (
	ProductsCollection Collection = "products"
)

var clientInstance *mongo.Client 

var mongoOnce sync.Once

var clientInstanceError error

func GetMongoClient() (*mongo.Client, error) {

    func initClient() {
		clientOptions := options.Client().ApplyURI(URI)

		client, err := mongo.Connect(context.TODO(), clientOptions)
        if err!= nil {
            panic(err)
        }
        clientInstance = client

		clientInstanceError = err 
	}

	mongoOnce.Do(initClient)

	return clientInstance, clientInstanceError
}