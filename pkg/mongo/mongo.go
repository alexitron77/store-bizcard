package mongo

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	conStrTemplate = "mongodb://%s:%s@localhost:27017"
)

type MongoModel struct {
	Url      string
	Username string
	Password string
}

func (m MongoModel) ConnectDB() (*mongo.Client, context.Context) {

	connectionString := fmt.Sprintf(conStrTemplate, m.Username, m.Password)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	client.Connect(ctx)

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		fmt.Print(err)
	}

	return client, ctx
}
