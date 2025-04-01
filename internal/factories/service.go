package factories

import "financial/internal/services"

type ServiceFactory interface {
	// athentication service
	CreateRegisterUser() services.RegisterUser
	CreateIsAuthenticated() services.IsAuthenticated
	CreateLogin() services.Login

	// bank account service
	CreateCreateBankAccount() services.CreateBankAccount
	CreateDeleteBankAccount() services.DeleteBankAccount
	CreatePaginateBankAccounts() services.PaginateBankAccounts
	CreateUpdateBankAccount() services.UpdateBankAccount

	// card service
	CreateCreateCard() services.CreateCard
	CreateDeleteCard() services.DeleteCard

	// transaction service
	CreateRecentTransactions() services.RecentTransactions
	CreateAddItemsToTransaction() services.AddItemsToTransaction
	CreatePaginateTransaction() services.PaginateTransaction
	CreateCurrentBalance() services.CurrentBalance
	CreateCreateTransaction() services.CreateTransaction
	CreateDeleteTransaction() services.DeleteTransaction
	CreateFindTransaction() services.FindTransaction
	CreateUpdateTransaction() services.UpdateTransaction

	// item service
	CreateDeleteItem() services.DeleteItem

	// badge service
	CreateCreateBadge() services.CreateBadge
	CreateMostExpansiveBadges() services.MostExpansiveBadges
	CreateDeleteBadge() services.DeleteBadge
	CreateUpdateBadge() services.UpdateBadge
	CreatePaginateBadges() services.PaginateBadges
}

type serviceFactory struct {
	repositoryFactory RepositoryFactory
}

func NewServiceFactory(repositoryFactory RepositoryFactory) ServiceFactory {
	return &serviceFactory{repositoryFactory}
}

func (s *serviceFactory) CreateRegisterUser() services.RegisterUser {
	return services.NewRegisterUser(
		s.repositoryFactory.CreateUserRepository(),
		s.repositoryFactory.CreateHashRepository(),
	)
}

func (s *serviceFactory) CreateIsAuthenticated() services.IsAuthenticated {
	return services.NewIsAuthenticated(
		s.repositoryFactory.CreateAuthenticationRepository(),
	)
}

func (s *serviceFactory) CreateLogin() services.Login {
	return services.NewLogin(
		s.repositoryFactory.CreateUserRepository(),
		s.repositoryFactory.CreateAuthenticationRepository(),
		s.repositoryFactory.CreateHashRepository(),
	)
}

func (s *serviceFactory) CreateCreateBankAccount() services.CreateBankAccount{
	return services.NewCreateBankAccount(
		s.repositoryFactory.CreateBankAccountRepository(),
	)
}

func (s *serviceFactory) CreateDeleteBankAccount() services.DeleteBankAccount{
	return services.NewDeleteBankAccountService(
		s.repositoryFactory.CreateBankAccountRepository(),
	)
}

func (s *serviceFactory) CreatePaginateBankAccounts() services.PaginateBankAccounts{
	return services.NewPaginateBankAccounts(
		s.repositoryFactory.CreateBankAccountRepository(),
	)
}

func (s *serviceFactory) CreateUpdateBankAccount() services.UpdateBankAccount {
	return services.NewUpdateBankAccount(
		s.repositoryFactory.CreateBankAccountRepository(),
	)
}

func (s *serviceFactory) CreateCreateCard() services.CreateCard {
	return services.NewCreateCardService(
		s.repositoryFactory.CreateCardRepository(),
	)
}

func (s *serviceFactory) CreateDeleteCard() services.DeleteCard {
	return services.NewDeleteCard(
		s.repositoryFactory.CreateCardRepository(),
		s.repositoryFactory.CreateBankAccountRepository(),
	)
}

func (s *serviceFactory) CreateRecentTransactions() services.RecentTransactions {
	return services.NewRecentTransactions(
		s.repositoryFactory.CreateTransactionRepository(),
	)

}

func (s *serviceFactory) CreateAddItemsToTransaction() services.AddItemsToTransaction {
	return services.NewAddItemsToTransaction(
		s.repositoryFactory.CreateItemRepository(),
		s.repositoryFactory.CreateBadgeRepository(),
		s.repositoryFactory.CreateTransactionRepository(),
		s.repositoryFactory.CreateBankAccountRepository(),
	)
}

func (s *serviceFactory) CreatePaginateTransaction() services.PaginateTransaction {
	return services.NewPaginateTransaction(
		s.repositoryFactory.CreateTransactionRepository(),
	)
}

func (s *serviceFactory) CreateCurrentBalance() services.CurrentBalance {
	return services.NewCurrentBalance(
		s.repositoryFactory.CreateTransactionRepository(),
	)
}

func (s *serviceFactory) CreateCreateTransaction() services.CreateTransaction {
	return services.NewCreateTransactionService(
		s.repositoryFactory.CreateTransactionRepository(),
		s.repositoryFactory.CreateBankAccountRepository(),
	)
}

func (s *serviceFactory) CreateDeleteTransaction() services.DeleteTransaction {
	return services.NewDeleteTransaction(
		s.repositoryFactory.CreateBankAccountRepository(),
		s.repositoryFactory.CreateTransactionRepository(),
	)
}

func (s *serviceFactory) CreateFindTransaction() services.FindTransaction {
	return services.NewFindTransaction(
		s.repositoryFactory.CreateTransactionRepository(),
	)
}

func (s *serviceFactory) CreateUpdateTransaction() services.UpdateTransaction {
	return services.NewUpdateTransaction(
		s.repositoryFactory.CreateTransactionRepository(),
	)
}

func (s *serviceFactory) CreateDeleteItem() services.DeleteItem {
	return services.NewDeleteBadge(
		s.repositoryFactory.CreateBadgeRepository(),
	)
}

func (s *serviceFactory) CreateCreateBadge() services.CreateBadge {
	return services.NewCreateBadge(
		s.repositoryFactory.CreateBadgeRepository(),
	)
}

func (s *serviceFactory) CreateMostExpansiveBadges() services.MostExpansiveBadges {
	return services.NewMostExpansiveBadges(
		s.repositoryFactory.CreateBadgeRepository(),
	)
}

func (s *serviceFactory) CreateDeleteBadge() services.DeleteBadge {
	return services.NewDeleteBadge(
		s.repositoryFactory.CreateBadgeRepository(),
	)
}

func (s *serviceFactory) CreateUpdateBadge() services.UpdateBadge {
	return services.NewUpdateBadge(
		s.repositoryFactory.CreateBadgeRepository(),
	)
}

func (s *serviceFactory) CreatePaginateBadges() services.PaginateBadges {
	return services.NewPaginateBadges(
		s.repositoryFactory.CreateBadgeRepository(),
	)
}
