package models

import "github.com/shopspring/decimal"

type Item struct {
	ID uint `json:"id"`

	TransactionID uint            `json:"transaction_id"`
	Name          string          `json:"name"`
	Value         decimal.Decimal `json:"value"`
	Quantity      uint            `json:"quantity"`
}

type ItemWithBadges struct {
	Item
	Badges []Badge `json:"badges"`
}
