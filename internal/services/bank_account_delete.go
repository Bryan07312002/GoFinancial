package services

import (
	"errors"
	"financial/internal/db"
)

type DeleteBankAccountService struct {
	bankAccountRepo db.BankAccountRepository
}

func NewDeleteBankAccountService(
	bankAccountRepo db.BankAccountRepository) DeleteBankAccountService {
	return DeleteBankAccountService{bankAccountRepo}
}

func (d *DeleteBankAccountService) Run(id, userId uint) error {
	bankAccount, err := d.bankAccountRepo.FindByID(id)
	if err != nil {
		return err
	}

	if bankAccount.UserID != userId {
		return errors.New("cant delete bank account from another user")
	}

	return d.bankAccountRepo.Delete(id)
}
