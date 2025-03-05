package services

import (
	"errors"
	"financial/internal/db"
)

type DeleteCard struct {
	cardRepo        db.CardRepository
	bankAccountRepo db.BankAccountRepository
}

func NewDeleteCard(
	cardRepo db.CardRepository,
	bankAccountRepo db.BankAccountRepository,
) DeleteCard {
	return DeleteCard{cardRepo, bankAccountRepo}
}

func (d *DeleteCard) Run(id, userId uint) error {
	bankAccountRepo, err := d.bankAccountRepo.FindBankAccountByCardID(id)
	if err != nil {
		return err
	}

	if bankAccountRepo.UserID != userId {
		return errors.New("cant delete card from another user")
	}

	return d.cardRepo.Delete(id)
}
