package mongo

import (
	"context"
	"fmt"
	"os"
	"time"

	"biz.card/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type BizCardModel struct {
	DB  *mongo.Client
	Ctx context.Context
	Log *logrus.Logger
}

const (
	conStrTemplate = "mongodb://%s:%s@localhost:27017"
)

type DBConn struct {
	Url      string
	Username string
	Password string
}

func (c DBConn) ConnectDB() *BizCardModel {

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

	return &BizCardModel{client, ctx, nil}
}

func NewBizCardModel(db *mongo.Client, ctx context.Context, log *logrus.Logger) *BizCardModel {
	return &BizCardModel{
		Log: log,
		DB:  db,
		Ctx: ctx,
	}
}

func (c *BizCardModel) Save(card *models.Bizcard) error {
	collection := c.DB.Database("bizcard").Collection("cards")
	_, err := collection.InsertOne(c.Ctx, card)

	if err != nil {
		return err
	}

	return nil
}
