package services

import (
	"financial/internal/db"

	"github.com/shopspring/decimal"
)

type CurrentBalance struct {
	transactionRepo db.TransactionRepository
}

func NewCurrentBalance(transactionRepo db.TransactionRepository) CurrentBalance {
	return CurrentBalance{transactionRepo}
}

type Balances struct {
	Balance decimal.Decimal `json:"balance"`
	Credit  decimal.Decimal `json:"credit"`
}

func (c *CurrentBalance) Run(userID uint) (Balances, error) {
	balance, credit, err := c.transactionRepo.GetCurrentBalances(userID)
	if err != nil {
		return Balances{}, err
	}

	return Balances{
		Balance: balance,
		Credit:  credit,
	}, nil
}
