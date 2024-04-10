package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName    string
	LastName     string
	DateOfBirth  time.Time
	MobileNumber string
	Email        string `gorm:"unique"`
	Password     string
}
