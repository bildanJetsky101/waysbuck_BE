package routes

import (
	"waysbuck/handlers"
	mysql "waysbuck/pkg"
	"waysbuck/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transactions", h.FindTransactions).Methods("GET")
	r.HandleFunc("/transaction/{id}", h.GetTransaction).Methods("GET")
	r.HandleFunc("/transaction", h.CreateTransaction).Methods("POST")
	//r.HandleFunc("/transaction/{id}", h.UpdateTransaction).Methods("PATCH")

}
