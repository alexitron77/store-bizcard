package postgres

import (
	"context"
	"fmt"
	"os"

	"biz.card/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConn struct {
	Url      string
	Username string
	Password string
}

type BizCardModel struct {
	DB  *gorm.DB
	Ctx context.Context
}

func (c DBConn) ConnectDB() *BizCardModel {
	con_str := fmt.Sprintf("host=%s user=postgres password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", c.Url, c.Password, "postgres")
	db, err := gorm.Open(postgres.Open(con_str))

	db.AutoMigrate(&models.Bizcard{})

	if err != nil {
		os.Exit(1)
	}

	return &BizCardModel{db, nil}
}

func NewBizCardModel(db *gorm.DB, ctx context.Context) *BizCardModel {
	return &BizCardModel{
		DB:  db,
		Ctx: ctx,
	}
}

func (c *BizCardModel) Save(card *models.Bizcard) error {
	c.DB.Create(&card)
	return nil
}
