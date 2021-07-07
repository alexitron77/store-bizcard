package config

import (
	"fmt"
	"os"
	"path/filepath"

	"biz.card/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	path, err := filepath.Abs("./config/env")

	if err != nil {
		fmt.Print(err)
	}

	config := LoadConfig(path)

	con_str := fmt.Sprintf("host=%s user=postgres password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", config.DBUrl, config.DBPassword, "postgres")
	db, err := gorm.Open(postgres.Open(con_str))

	db.AutoMigrate(&models.Bizcard{})

	if err != nil {
		os.Exit(1)
	}
	return db
}
