package repositories

import (
	"biz.card/models"
	"gorm.io/gorm"
)

type BizCardModel struct {
	DB *gorm.DB
}

func NewBizCardModel(db *gorm.DB) *BizCardModel {
	return &BizCardModel{
		DB: db,
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

	c.DB.Create(&card)
}
