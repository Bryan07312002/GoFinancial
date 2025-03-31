package services

import (
	"financial/internal/db"

	"github.com/shopspring/decimal"
)

type CurrentBalance interface {
	Run(userID uint) (Balances, error)
}

type currentBalance struct {
	transactionRepo db.TransactionRepository
}

func NewCurrentBalance(transactionRepo db.TransactionRepository) CurrentBalance {
	return &currentBalance{transactionRepo}
}

type Balances struct {
	Balance decimal.Decimal `json:"balance"`
	Credit  decimal.Decimal `json:"credit"`
}

func (c *currentBalance) Run(userID uint) (Balances, error) {
	balance, credit, err := c.transactionRepo.GetCurrentBalances(userID)
	if err != nil {
		return Balances{}, err
	}

	return Balances{
		Balance: balance,
		Credit:  credit,
	}, nil
}
