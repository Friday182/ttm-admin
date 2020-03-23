package model

import (
	//"errors"
	// ws "github.com/friday182/ttm-go/app/api/websocket"
	"time"
)

type Mentor struct {
	ID             int `gorm:"primary_key"`
	Name           string
	Email          string `gorm:"type:varchar(100);index;unique_index"`
	Password       string `gorm:"not null"`
	Gid            string `sql:"index"`
	Role           string
	Gender         string
	Age            int
	Country        string
	City           string
	PromotionCode  string
	StudentMax     int
	StudentNum     int
	Contacts       []byte
	VerifyCode     string
	LastSigninTime time.Time
	SigninCount    int
	Token          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

func AddMentor(email string, pw string) error {
	u := Mentor{}
	u.Email = email
	u.Password = pw

	// email is the unique key in this table, db will validate it
	err := db.Save(&u).Error
	return err
}

func ValidMentor(email string, pw string) bool {
	u := Mentor{}
	if db.Where(Mentor{Email: email, Password: pw}).Find(&u).RecordNotFound() {
		return false
	}
	return true
}
