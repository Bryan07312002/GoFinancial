package models

import "github.com/shopspring/decimal"

type Item struct {
	ID uint

	TransactionID uint
	Name          string
	Value         decimal.Decimal
	Quantity      uint
}
