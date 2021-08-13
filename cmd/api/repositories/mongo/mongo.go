package mongo

import (
	"context"
	"fmt"
	"os"
	"time"

	"biz.card/cmd/api/models"
	"biz.card/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	conStrTemplate = "mongodb://%s:%s@%s:27017"
)

type MongoRepo struct {
	Storage *config.Storage
	Config  *config.Config
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

	connectionString := fmt.Sprintf(conStrTemplate, c.Username, c.Password, c.Url)

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

func NewDBRepo(config *config.Config, storage *config.Storage) *MongoRepo {
	return &MongoRepo{
		Storage: storage,
		Config:  config,
	}
}

func (c *MongoRepo) Create(ctx context.Context, card *models.Bizcard) (string, error) {
	collection := c.Storage.DB.Database("bizcard").Collection("cards")
	result, err := collection.InsertOne(ctx, card)

	id := result.InsertedID.(primitive.ObjectID)

	if err != nil {
		return "", err
	}

	return id.Hex(), nil
}

func (c *MongoRepo) Read(ctx context.Context, id string) (models.Bizcard, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.Config.Log.Errorf(err.Error())
	}
	collection := c.Storage.DB.Database("bizcard").Collection("cards")
	cur := collection.FindOne(ctx, bson.M{"_id": objectID})

	var card models.Bizcard
	cur.Decode(&card)

	return card, nil
}

func (c *MongoRepo) ReadAll(ctx context.Context) ([]models.Bizcard, error) {
	var result []models.Bizcard

	collection := c.Storage.DB.Database("bizcard").Collection("cards")
	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		fmt.Print(err)
		return []models.Bizcard{}, err
	}

	for cur.Next(ctx) {
		var card models.Bizcard
		cur.Decode(&card)
		result = append(result, card)
	}

	return result, nil
}

func (c *MongoRepo) Update(ctx context.Context, id string, value string) {
	collection := c.Storage.DB.Database("bizcard").Collection("cards")
	objectId, _ := primitive.ObjectIDFromHex(id)
	_, err := collection.UpdateOne(ctx, bson.M{
		"_id": objectId},
		bson.D{
			{"$set", bson.D{{"card_url", value}}},
		},
	)
	if err != nil {
		c.Config.Log.Fatal("The update has failed")
	}

}
