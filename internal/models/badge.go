package models

import "github.com/shopspring/decimal"

type Badge struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type BadgeWithValue struct {
	Badge
	Value decimal.Decimal `json:"value"`
}
