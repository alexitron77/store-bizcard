package repositories

import (
	"context"
	"fmt"

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

func (c *BizCardModel) Save() {
	card := &models.Bizcard{
		FirstName:   "Alexis",
		LastName:    "Tran",
		Role:        "Software Engineer",
		Company:     "thales",
		Country:     "Singapore",
		PhoneNumber: "88924600",
		Website:     "www.alexis.tran",
	}

	collection := c.DB.Database("bizcard").Collection("cards")

	_, err := collection.InsertOne(c.Ctx, card)

	if err != nil {
		fmt.Print("save", err)
	}

	// c.DB.Create(&card)
}
