package models

type Bizcard struct {
	// gorm.Model
	FirstName   string `gorm:"firstname"`
	LastName    string `gorm:"lastname"`
	Role        string `gorm:"role"`
	Company     string `gorm:"company"`
	Country     string `gorm:"country"`
	PhoneNumber string `gorm:"phone_number"`
	Website     string `gorm:"website"`
	LinkedIn    string `gorm:"url"`
}

type BizcardRepo interface {
	Save(*Bizcard) error
}
