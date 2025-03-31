package services

import (
	"errors"
	"financial/internal/db"
)

type DeleteBankAccount interface {
	Run(id, userId uint) error
}

type deleteBankAccount struct {
	bankAccountRepo db.BankAccountRepository
}

func NewDeleteBankAccountService(
	bankAccountRepo db.BankAccountRepository) DeleteBankAccount {
	return &deleteBankAccount{bankAccountRepo}
}

func (d *deleteBankAccount) Run(id, userId uint) error {
	bankAccount, err := d.bankAccountRepo.FindByID(id, userId)
	if err != nil {
		return err
	}

	if bankAccount.UserID != userId {
		return errors.New("cant delete bank account from another user")
	}

	return d.bankAccountRepo.Delete(id)
}
