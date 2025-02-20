package services

import (
	"errors"
	"financial/internal/db"
)

type DeleteTransaction struct {
	bankAccountRepo db.BankAccountRepository
	transactionRepo db.TransactionRepository
}

func NewDeleteTransaction(
	bankAccountRepo db.BankAccountRepository,
	transactionRepo db.TransactionRepository,
) DeleteTransaction {
	return DeleteTransaction{bankAccountRepo, transactionRepo}
}

// TODO: check date of transaction and update all balances after the transaction
func (d *DeleteTransaction) Run(transactionId, userId uint) error {
	bankAccount, err := d.bankAccountRepo.FindBankAccountByTransactionID(
		transactionId)

	if err != nil {
		return err
	}

	if bankAccount.UserID != userId {
		return errors.New("cant delete another user transaction")
	}

	return d.transactionRepo.Delete(transactionId)
}
