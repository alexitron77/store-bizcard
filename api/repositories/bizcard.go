package repositories

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type BizCardRepo struct {
	db *gorm.DB
}

func NewBizCardRepo(db *gorm.DB) *BizCardRepo {
	return &BizCardRepo{
		db: db,
	}
}

func (r *BizCardRepo) Save() error {
	fmt.Print("Saved")
	return nil
}
