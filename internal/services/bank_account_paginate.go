package services

import (
	"financial/internal/db"
	"financial/internal/models"
)

type PaginateBankAccounts interface {
	Run(
		paginateOpt db.PaginateOptions,
		userID uint,
	) (db.PaginateResult[models.BankAccount], error)
}

type paginateBankAccounts struct {
	bankAccountRepo db.BankAccountRepository
}

func NewPaginateBankAccounts(
	bankAccountRepo db.BankAccountRepository) PaginateBankAccounts {
	return &paginateBankAccounts{bankAccountRepo}
}

func (p *paginateBankAccounts) Run(
	paginateOpt db.PaginateOptions,
	userID uint,
) (db.PaginateResult[models.BankAccount], error) {
	return p.bankAccountRepo.PaginateFromUserID(paginateOpt, userID)
}
