package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/terrnit/pizza-crud-go/internal/delivery/http/pizza"
	"github.com/terrnit/pizza-crud-go/internal/repository/pizza"
	"github.com/terrnit/pizza-crud-go/internal/usecase/pizza"
)

func main() {
	// Initialize repository
	repo := pizza.NewMemoryRepository()

	// Initialize use case
	useCase := pizza.NewUseCase(repo)

	// Initialize handler
	handler := pizza.NewHandler(useCase)

	// Setup router
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/pizzas", handler.Create).Methods(http.MethodPost)
	router.HandleFunc("/pizzas/{id:[0-9]+}", handler.GetByID).Methods(http.MethodGet)
	router.HandleFunc("/pizzas/{id:[0-9]+}", handler.Update).Methods(http.MethodPut)
	router.HandleFunc("/pizzas/{id:[0-9]+}", handler.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/pizzas", handler.List).Methods(http.MethodGet)

	// Start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
