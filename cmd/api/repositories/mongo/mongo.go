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

func NewDBRepo(config *config.Config) *MongoRepo {
	return &MongoRepo{
		Conf: config,
	}
}

func (c *MongoRepo) Create(card *models.Bizcard) (string, error) {
	collection := c.Conf.DB.Database("bizcard").Collection("cards")
	result, err := collection.InsertOne(c.Conf.Ctx, card)

	id := result.InsertedID.(primitive.ObjectID)

	if err != nil {
		return "", err
	}

	return id.Hex(), nil
}

func (c *MongoRepo) Read(id string) (models.Bizcard, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.Conf.Log.Errorf(err.Error())
	}
	collection := c.Conf.DB.Database("bizcard").Collection("cards")
	cur := collection.FindOne(c.Conf.Ctx, bson.M{"_id": objectID})

	var card models.Bizcard
	cur.Decode(&card)

	return card, nil
}

func (c *MongoRepo) ReadAll() ([]models.Bizcard, error) {
	var result []models.Bizcard

	collection := c.Conf.DB.Database("bizcard").Collection("cards")
	cur, err := collection.Find(c.Conf.Ctx, bson.D{})

	if err != nil {
		fmt.Print(err)
		return []models.Bizcard{}, err
	}

	for cur.Next(c.Conf.Ctx) {
		var card models.Bizcard
		cur.Decode(&card)
		result = append(result, card)
	}

	return result, nil
}

func (c *MongoRepo) Update(id string, value string) {
	collection := c.Conf.DB.Database("bizcard").Collection("cards")
	objectId, _ := primitive.ObjectIDFromHex(id)
	_, err := collection.UpdateOne(c.Conf.Ctx, bson.M{
		"_id": objectId},
		bson.D{
			{"$set", bson.D{{"card_url", value}}},
		},
	)
	if err != nil {
		c.Conf.Log.Fatal("The update has failed")
	}

}
