package models

import "time"

type Transfer struct {
	Base
	RecipientName string
	Ammount       uint
	Currency      Currency `sql:"type:currency;default:'EUR'"`
	CardID        string
	PaymentDate   time.Time
}
