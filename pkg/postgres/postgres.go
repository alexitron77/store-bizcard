package postgres

import (
	"fmt"
	"os"

	"biz.card/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresModel struct {
	Url      string
	Username string
	Password string
}

func (p PostgresModel) ConnectDB() *gorm.DB {
	con_str := fmt.Sprintf("host=%s user=postgres password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", p.Url, p.Password, "postgres")
	db, err := gorm.Open(postgres.Open(con_str))

	db.AutoMigrate(&models.Bizcard{})

	if err != nil {
		os.Exit(1)
	}
	return db
}
