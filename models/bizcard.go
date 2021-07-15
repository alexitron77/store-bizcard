package models

type Bizcard struct {
	// gorm.Model
	FirstName   string `gorm:"firstname" json:"firstname" example:"Alexis"`
	LastName    string `gorm:"lastname" json:"lastname" example:"Tran"`
	Role        string `gorm:"role" json:"role" example:"Software engineer"`
	Company     string `gorm:"company" json:"company" example:"Thales"`
	Country     string `gorm:"country" json:"country" example:"Singapore"`
	PhoneNumber string `gorm:"phone_number" json:"phone_number" example:"88924600"`
	Website     string `gorm:"website" json:"website" example:"www.alexis.tran.com"`
	LinkedIn    string `gorm:"url" json:"linked_in" example:"null"`
}

type BizcardRepo interface {
	Create(*Bizcard) error
	Read(string) ([]Bizcard, error)
}
