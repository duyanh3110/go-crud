package models

import (
	"database/sql/driver"
	"time"
)

type TransactionStatus string

const (
	DONE    TransactionStatus = "done"
	PENDING TransactionStatus = "pending"
)

func (ts *TransactionStatus) Scan(value interface{}) error {
	*ts = TransactionStatus(value.([]byte))
	return nil
}

func (ts TransactionStatus) Value() (driver.Value, error) {
	return string(ts), nil
}

type TransactionType string

const (
	INCOME  TransactionType = "income"
	EXPENSE TransactionType = "expense"
)

func (tt *TransactionType) Scan(value interface{}) error {
	*tt = TransactionType(value.([]byte))
	return nil
}

func (tt TransactionType) Value() (driver.Value, error) {
	return string(tt), nil
}

type TransactionCategory string

const (
	WITHDRAW      TransactionCategory = "withdraw"
	ENTERTAINMENT TransactionCategory = "entertainment"
	TECHNOLOGY    TransactionCategory = "technology"
	SOFTWARE      TransactionCategory = "software"
	MOBILE        TransactionCategory = "mobile"
)

type Transaction struct {
	Base
	CardID              string
	TransactionStatus   string `sql:"type:transaction_status"`
	TransactionName     string
	BusinessName        string
	PaymentDate         time.Time
	Ammount             uint
	TransactionType     string   `sql:"type:transaction_type"`
	TransactionCategory string   `sql:"type:transaction_category"`
	Currency            Currency `sql:"type:currency;default:'EUR'"`
}
