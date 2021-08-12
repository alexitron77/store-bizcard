package postgres

import (
	"context"
	"fmt"
	"os"

	"biz.card/cmd/api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConn struct {
	Url      string
	Username string
	Password string
}

type DBModel struct {
	DB  *gorm.DB
	Ctx context.Context
}

func (c DBConn) ConnectDB() *DBModel {
	con_str := fmt.Sprintf("host=%s user=postgres password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", c.Url, c.Password, "postgres")
	db, err := gorm.Open(postgres.Open(con_str))

	db.AutoMigrate(&models.Bizcard{})

	if err != nil {
		os.Exit(1)
	}

	return &DBModel{db, nil}
}

func NewBizCardModel(db *gorm.DB, ctx context.Context) *DBModel {
	return &DBModel{
		DB:  db,
		Ctx: ctx,
	}
}

func (c *DBModel) Save(card *models.Bizcard) error {
	c.DB.Create(&card)
	return nil
}
