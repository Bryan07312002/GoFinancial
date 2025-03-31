package services

import (
	"financial/internal/db"
	"financial/internal/models"
)

type CreateBankAccount interface {
	Run(create CreateBankAccountDto) error
}

type createBankAccount struct {
	bankAccountRepo db.BankAccountRepository
}

func NewCreateBankAccount(
	bankAccountRepo db.BankAccountRepository) CreateBankAccount {
	return &createBankAccount{bankAccountRepo}
}

type CreateBankAccountDto struct {
	UserId      uint
	Name        string
	Description string
}

func (c *createBankAccount) Run(create CreateBankAccountDto) error {
	if _, err := c.bankAccountRepo.Create(models.BankAccount{
		Name:        create.Name,
		UserID:      create.UserId,
		Description: create.Description,
	}); err != nil {
		return err
	}

	return nil
}
