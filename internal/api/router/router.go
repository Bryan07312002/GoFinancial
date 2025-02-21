package routes

import (
	"financial/internal/api/handlers"
	"financial/internal/api/router/middlewares"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Router struct {
	*mux.Router
}

func NewRouter(dbCon *gorm.DB, jwtKey string) *Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", handlers.HealthHandler).Methods("GET")

	r.HandleFunc("/register",
		handlers.CreateRegisterUserHandler(dbCon)).Methods("POST")

	r.HandleFunc("/login",
		handlers.CreateLoginHandler(dbCon, jwtKey)).Methods("POST")

	protected := r.NewRoute().Subrouter()
	protected.Use(middlewares.CreateAuthMiddleware(jwtKey))

	protected.HandleFunc("/bank_accounts",
		handlers.CreateCreateBankAccountHandler(dbCon)).Methods("POST")
	protected.HandleFunc("/bank_accounts",
		handlers.CreatePaginateBankAccountHandler(dbCon)).Methods("GET")
	protected.HandleFunc("/bank_accounts/{id}",
		handlers.CreateBankAccountDelete(dbCon)).Methods("DELETE")

	protected.HandleFunc("/cards",
		handlers.CreateCreateCard(dbCon)).Methods("POST")

	protected.HandleFunc("/transactions",
		handlers.CreateCreateTransaction(dbCon)).Methods("POST")
	protected.HandleFunc("/transactions/{id}",
		handlers.CreateTransactionDelete(dbCon)).Methods("DELETE")

	protected.HandleFunc("/items",
		handlers.CreateAddItemsToTransaction(dbCon)).Methods("POST")
	protected.HandleFunc("/items/{id}",
		handlers.CreateDeleteItem(dbCon)).Methods("DELETE")

	return &Router{r}
}
