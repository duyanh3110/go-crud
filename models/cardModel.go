package models

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

type CardType string

const (
	VISA       CardType = "VISA"
	MASTERCARD CardType = "MASTERCARD"
	AMEX       CardType = "AMEX"
)

func (c *CardType) Scan(value interface{}) error {
	*c = CardType(value.([]byte))
	return nil
}

func (c CardType) Value() (driver.Value, error) {
	return string(c), nil
}

type Card struct {
	Base
	CardType     CardType `sql:"type:card_type;"`
	CardNumber   string
	Expiry       time.Time
	Ballance     uint
	Currency     Currency `sql:"type:currency;default:'EUR'"`
	WalletID     string
	Transactions []Transaction
	Transfers    []Transfer
}

func (c *Card) BeforeCreate(tx *gorm.DB) (err error) {
	today := time.Now()
	endOfMonth := today.AddDate(0, 1, -today.Day())
	c.Expiry = endOfMonth
	return
}
