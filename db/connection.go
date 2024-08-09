package db

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		// Configuração da URI do MongoDB
		uri := "mongodb://localhost:27017"

		// Configuração  do cliente
		clientOptions := options.Client().ApplyURI(uri)

		// Conexão o MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
			return
		}

		// Verificação da conexão
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
			return
		}

		clientInstance = client
	})

	return clientInstance, clientInstanceError
}
