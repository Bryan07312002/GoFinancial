package routes

import (
	"financial/internal/api/handlers"
	"financial/internal/api/router/middlewares"
	"financial/internal/factories"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Router struct {
	*mux.Router
}

func NewRouter(dbCon *gorm.DB, jwtKey string) *Router {
	factory := factories.NewServiceFactory(
		factories.NewRepositoryFactory(dbCon, jwtKey),
	)
	r := mux.NewRouter()

	r.HandleFunc("/health", handlers.HealthHandler).Methods("GET")

	r.Handle("/login", handlers.NewLoginHandler(factory)).Methods("POST")
	r.Handle("/register", handlers.NewRegisterUserHandler(factory)).Methods("POST")

	protected := r.NewRoute().Subrouter()
	protected.Use(middlewares.CreateAuthMiddleware(factory))

	protected.Handle("/bank_accounts",
		handlers.NewCreateBankAccountHandler(factory)).Methods("POST")
	protected.Handle("/bank_accounts",
		handlers.NewPaginateBankAccountHandler(factory)).Methods("GET")
	protected.Handle("/bank_accounts/{id}",
		handlers.NewDeleteBankAccountHandler(factory)).Methods("DELETE")

	protected.Handle("/cards",
		handlers.NewCreateCardHandler(factory)).Methods("POST")

	protected.Handle("/transactions",
		handlers.NewCreateTransactionHandler(factory)).Methods("POST")
	protected.Handle("/transactions/{id}",
		handlers.NewUpdateTransactionHandler(factory)).Methods("PUT")
	protected.Handle("/transactions/{id}",
		handlers.NewUpdateTransactionHandler(factory)).Methods("PATCH")
	protected.Handle("/transactions",
		handlers.NewPaginateTransaction(factory)).Methods("GET")
	protected.Handle("/transactions/recent",
		handlers.NewRecentTransactionsHandler(factory)).Methods("GET")
	protected.Handle("/transactions/balance",
		handlers.NewCurrentBalanceHandler(factory)).Methods("GET")
	protected.Handle("/transactions/{id}",
		handlers.NewFindTransactionHandler(factory)).Methods("GET")
	protected.Handle("/transactions/{id}",
		handlers.NewDeleteTransactionHandler(factory)).Methods("DELETE")

	protected.Handle("/items",
		handlers.NewAddItemsToTransactionHandler(factory)).Methods("POST")
	protected.Handle("/items/{id}",
		handlers.NewDeleteItemHandler(factory)).Methods("DELETE")

	protected.Handle("/badges/most-expansive",
		handlers.NewMostExpansiveBadgesHandler(factory)).Methods("GET")
	protected.Handle("/badges/{id}",
		handlers.NewDeleteBadgeHandler(factory)).Methods("DELETE")
	protected.Handle("/badges",
		handlers.NewCreateBadgeHandler(factory)).Methods("POST")
	protected.Handle("/badges",
		handlers.NewPaginateBadgesHandler(factory)).Methods("GET")

	return &Router{r}
}
