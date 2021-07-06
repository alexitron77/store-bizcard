package models

type Bizcard struct {
	firstName   string
	lastName    string
	role        string
	company     string
	country     string
	phoneNumber string
	website     string
	linkedIn    string
}

type BizcardRepository interface {
	Save() error
}
