package services

import (
	"financial/internal/db"
	"financial/internal/models"
)

type PaginateBankAccountsService struct {
	bankAccountRepo db.BankAccountRepository
}

func NewPaginateBankAccountsService(
	bankAccountRepo db.BankAccountRepository) PaginateBankAccountsService {
	return PaginateBankAccountsService{bankAccountRepo}
}

func (p *PaginateBankAccountsService) Run(
	paginateOpt db.PaginateOptions,
	userID uint,
) (db.PaginateResult[models.BankAccount], error) {
	return p.bankAccountRepo.PaginateFromUserID(paginateOpt, userID)
}
