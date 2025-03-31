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
	factory := factories.NewServiceFactory(factories.NewRepositoryFactory(dbCon))
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
	protected.HandleFunc("/transactions/{id}",
		handlers.CreateFindTransaction(dbCon)).Methods("GET")
	protected.HandleFunc("/transactions/{id}",
		handlers.CreateTransactionDelete(dbCon)).Methods("DELETE")

	protected.HandleFunc("/items",
		handlers.CreateAddItemsToTransaction(dbCon)).Methods("POST")
	protected.HandleFunc("/items/{id}",
		handlers.CreateDeleteItem(dbCon)).Methods("DELETE")

	protected.HandleFunc("/badges/most-expansive",
		handlers.CreateMostExpansiveBudgets(dbCon)).Methods("GET")
	protected.HandleFunc("/badges/{id}",
		handlers.CreateDeleteBadge(dbCon)).Methods("DELETE")
	protected.HandleFunc("/badges",
		handlers.CreateCreateBadge(dbCon)).Methods("POST")
	protected.HandleFunc("/badges",
		handlers.CreatePaginateBadge(dbCon)).Methods("GET")

	return &Router{r}
}
