package model

import (
	"time"
)

// Users
type User struct {
	ID        int `gorm:"primary_key"`
	Gid       string
	Email     string
	Password  string `gorm:"not null"`
	Username  string `gorm:"type:varchar(100);unique_index"`
	Role      string `gorm:"size:255"`
	Name      string
	Mobile    string
	Comment   string
	LastLoginTime time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
