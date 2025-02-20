package services

import (
	"financial/internal/db"
	"financial/internal/models"
)

type CreateBankAccountService struct {
	bankAccountRepo db.BankAccountRepository
}

func NewCreateBankAccountService(
	bankAccountRepo db.BankAccountRepository) CreateBankAccountService {
	return CreateBankAccountService{bankAccountRepo}
}

type CreateBankAccount struct {
	UserId      uint
	Name        string
	Description string
}

func (c *CreateBankAccountService) Run(create CreateBankAccount) error {
	if _, err := c.bankAccountRepo.Create(models.BankAccount{
		Name:        create.Name,
		UserID:      create.UserId,
		Description: create.Description,
	}); err != nil {
		return err
	}

	return nil
}
