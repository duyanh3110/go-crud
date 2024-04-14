package models

import (
	"time"
)

type User struct {
	Base
	FirstName    string
	LastName     string
	DateOfBirth  time.Time
	MobileNumber string
	Email        string `gorm:"unique"`
	Password     string
}
