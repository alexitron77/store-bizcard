package mongo

import (
	"context"
	"fmt"
	"os"
	"time"

	"biz.card/config"
	"biz.card/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	conStrTemplate = "mongodb://%s:%s@localhost:27017"
)

type MongoRepo struct {
	Conf *config.Config
}

type DBModel struct {
	DB  *mongo.Client
	Ctx context.Context
}

type DBConn struct {
	Url      string
	Username string
	Password string
}

func (c DBConn) ConnectDB() *DBModel {

	connectionString := fmt.Sprintf(conStrTemplate, c.Username, c.Password)

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

	return &DBModel{client, ctx}
}

func NewDBRepo(config *config.Config) *MongoRepo {
	return &MongoRepo{
		Conf: config,
	}
}

func (c *MongoRepo) Save(card *models.Bizcard) error {
	collection := c.Conf.DB.Database("bizcard").Collection("cards")
	_, err := collection.InsertOne(c.Conf.Ctx, card)

	if err != nil {
		return err
	}

	return nil
}
