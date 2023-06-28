package dbclient

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Database *mongo.Database

func Connect() error {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	var err error
	Client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	err = Client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	Database = Client.Database(os.Getenv("DEFAULT_DB"))
	return nil
}
