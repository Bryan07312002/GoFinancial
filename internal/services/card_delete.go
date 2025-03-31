package services

import (
	"errors"
	"financial/internal/db"
)

type DeleteCard interface {
	Run(id, userId uint) error
}

type deleteCard struct {
	cardRepo        db.CardRepository
	bankAccountRepo db.BankAccountRepository
}

func NewDeleteCard(
	cardRepo db.CardRepository,
	bankAccountRepo db.BankAccountRepository,
) DeleteCard {
	return &deleteCard{cardRepo, bankAccountRepo}
}

func (d *deleteCard) Run(id, userId uint) error {
	bankAccountRepo, err := d.bankAccountRepo.FindBankAccountByCardID(id)
	if err != nil {
		return err
	}

	if bankAccountRepo.UserID != userId {
		return errors.New("cant delete card from another user")
	}

	return d.cardRepo.Delete(id)
}
