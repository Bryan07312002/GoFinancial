package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type PaymentMethod string

const (
	CreditCard PaymentMethod = "credit_card"
	DebitCard  PaymentMethod = "debit_card"
	Other      PaymentMethod = "other"
)

func (p PaymentMethod) String() string {
	return string(p)
}

func (p PaymentMethod) IsValid() bool {
	switch p {
	case CreditCard, DebitCard, Other:
		return true
	}
	return false
}

func PaymentMethods() []PaymentMethod {
	return []PaymentMethod{CreditCard, DebitCard, Other}
}

type TransactionType string

const (
	Income   TransactionType = "income"
	Expense  TransactionType = "expense"
	Transfer TransactionType = "transfer"
)

func (t TransactionType) String() string {
	return string(t)
}

func (t TransactionType) IsValid() bool {
	switch t {
	case Income, Expense, Transfer:
		return true
	}
	return false
}

func TransactionTypes() []TransactionType {
	return []TransactionType{Income, Expense, Transfer}
}

type Transaction struct {
	ID            uint            `json:"id"`
	Type          TransactionType `json:"type"`
	Method        PaymentMethod   `json:"method"`
	Credit        bool            `json:"credit"`
	Value         decimal.Decimal `json:"value"`
	Date          time.Time       `json:"date"`
	CardID        *uint           `json:"card_id"`
	BankAccountID uint            `json:"bank_account_id"`
}

type TransactionWithBadges struct {
	Transaction
	Badges []Badge
}
