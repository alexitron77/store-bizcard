package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "biz_card.db")

	if err != nil {
		os.Exit(1)
	}

	return db
}
