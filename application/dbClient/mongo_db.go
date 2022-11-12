package dbClient

import (
	"context"
	"sync"

	"github.com/myKemal/go_restfull_api/application/common"
	"github.com/myKemal/go_restfull_api/application/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
	Used to create a singleton object of MongoDB client.

Initialized and exposed through  GetMongoClient().
*/
var clientInstance *mongo.Client

// Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error

// Used to execute client creation procedure only once.
var mongoOnce sync.Once

type MongoClient interface {
	GetCollection() *mongo.Collection
	Disconnect()
}

type mongoClient struct {
	client *mongo.Client
	err    *error
}

func (m *mongoClient) GetCollection() *mongo.Collection {
	return m.client.Database("getircase-study").Collection("records")
}

func NewMongoClient() MongoClient {

	mongoOnce.Do(func() {

		// Set client options
		clientOptions := options.Client().ApplyURI(config.EnvMongoURI())
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			common.Logger.Error(err)
			clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			common.Logger.Error(err)
			clientInstanceError = err
		}
		clientInstance = client
	})

	return &mongoClient{client: clientInstance, err: &clientInstanceError}
}

func (m *mongoClient) Disconnect() {
	_ = m.client.Disconnect(context.Background())
}
