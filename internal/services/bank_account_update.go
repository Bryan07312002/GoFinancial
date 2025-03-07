package services

import "financial/internal/db"

type UpdateBankAccount struct {
	bankAccountRepo db.BankAccountRepository
}

func NewUpdateBankAccount(
	bankAccountRepo db.BankAccountRepository) UpdateBankAccount {
	return UpdateBankAccount{bankAccountRepo}
}

type UpdateBankAccountDto struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (u *UpdateBankAccount) Run(
	bankAccountID uint, dto UpdateBankAccountDto, userID uint) error {
	bankAccount, err := u.bankAccountRepo.FindByID(bankAccountID, userID)
	if err != nil {
		return err
	}

	if dto.Name != nil {
		bankAccount.Name = *dto.Name
	}

	if dto.Description != nil {
		bankAccount.Description = *dto.Description
	}

	return u.bankAccountRepo.Update(bankAccount)
}
