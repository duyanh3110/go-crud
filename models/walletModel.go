package models

import "database/sql/driver"

type Currency string

const (
	USD Currency = "USD"
	EUR Currency = "EUR"
)

func (c *Currency) Scan(value interface{}) error {
	*c = Currency(value.([]byte))
	return nil
}

func (c Currency) Value() (driver.Value, error) {
	return string(c), nil
}

type Wallet struct {
	Base
	User     User `gorm:"foreignKey:UserID"`
	UserID   string
	Currency Currency `sql:"type:currency;default:'EUR'"`
	IsActive bool     `gorm:"default:true"`
	Cards    []Card
}
