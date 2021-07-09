package repositories

import (
	"context"

	"biz.card/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type BizCardModel struct {
	DB  *mongo.Client
	Ctx context.Context
}

func NewBizCardModel(db *mongo.Client, ctx context.Context) *BizCardModel {
	return &BizCardModel{
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
	// c.DB.Create(&card)
}
